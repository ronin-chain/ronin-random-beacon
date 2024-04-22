package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ronin-chain/ronin-random-beacon/pkg/config"
	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/ronin-chain/ronin-random-beacon/pkg/core"
	"github.com/ronin-chain/ronin-random-beacon/pkg/dbutils"
	"github.com/ronin-chain/ronin-random-beacon/pkg/listener"
	"github.com/ronin-chain/ronin-random-beacon/pkg/task"
	"github.com/ronin-chain/ronin-random-beacon/pkg/utils"
	"github.com/urfave/cli"
)

const (
	creatingTaskInterval   = 60               // 1 minute
	sendingTaskInterval    = 60 * 5           // 5 minutes
	cleaningTaskInterval   = 60 * 60 * 24 * 7 // 1 week
	trackerInterval        = 3                // 3s
	reorgCheckerInterval   = 60 * 5           // 5 minutes
	taskSyncerSentInterval = 60               // 1 minutes

	creatingTaskBatchSize   = 5
	sendingTaskBatchSize    = 5
	reorgTaskBatchSize      = 5
	taskSyncerSentBatchSize = 5

	chainInfo = "chainInfo"
	blockSize = 100
)

var (
	ConfigFileFlag = cli.StringFlag{
		Name:  "config-file",
		Usage: "Explicitly defines the path, name and extension of the yaml config file",
		Value: "./config.yaml",
	}
	PeriodFlag = cli.StringFlag{
		Name:     "period",
		Usage:    "Period of the random beacon",
		Required: true,
	}
	PrevBeaconFlag = cli.StringFlag{
		Name:     "prev-beacon",
		Usage:    "Previous beacon",
		Required: true,
	}
	OutputPathFlag = cli.StringFlag{
		Name:  "output-path",
		Usage: "Output path for the proof file",
	}
	GenKeyCommand = cli.Command{
		Action:      generateVRFKey,
		Name:        "generate-key",
		Usage:       "Generate secret key and verification key",
		Category:    "Random Beacon Commands",
		Description: "The generate-key command generate the pair of secret key and verification key",
	}
	StartCommand = cli.Command{
		Action:      start,
		Name:        "start",
		Usage:       "Start the random beacon",
		Category:    "Random Beacon Commands",
		Description: "The start tracking event from Coodirnator contract and submit the proof to the contract",
		Flags: []cli.Flag{
			ConfigFileFlag,
		},
	}
	RandomCommand = cli.Command{
		Action:      random,
		Name:        "random",
		Usage:       "Generate a random number",
		Category:    "Random Beacon Commands",
		Description: "Random based on the input flag and write out the proof file ./random-result.json",
		Flags: []cli.Flag{
			ConfigFileFlag,
			PeriodFlag,
			PrevBeaconFlag,
			OutputPathFlag,
		},
	}
)

func generateVRFKey(ctx *cli.Context) {
	core.GenerateVRFKey()
}

func start(cli *cli.Context) {
	config := config.LoadConfig(cli.String(ConfigFileFlag.GetName()))
	dbClient := dbutils.Open(config.DatabaseURL)
	dbutils.MigrateDB(dbClient, context.Background())
	ethClient, err := ethclient.Dial(config.RpcEndpoint)

	if err != nil {
		panic(fmt.Errorf("Failed to connect to the Ethereum client: %w", err))
	}

	// Chain listener
	chainListener, chainProcessor := listener.NewChainListener(chainInfo, dbClient, ethClient, contract.Coordinator, config.SkynetEndpoint, trackerInterval)
	// // Random request listener
	tracker := listener.NewRandomRequestListener(chainInfo, &chainListener, contract.Coordinator.Contract, dbClient, config.CoordinatorStartingBlock, blockSize, trackerInterval)
	// Produce task based on tracked RandomRequest.
	taskCreator := task.NewTaskCreator(dbClient, config, creatingTaskBatchSize, creatingTaskInterval)
	taskSender := task.NewTaskSender(dbClient, sendingTaskBatchSize, contract.Coordinator, chainProcessor, sendingTaskInterval)

	taskSyncerSent := task.NewTaskSyncerSent(dbClient, taskSyncerSentBatchSize, ethClient, chainProcessor, taskSyncerSentInterval)
	taskReorgChecker := task.NewTaskReorgChecker(dbClient, ethClient, contract.Coordinator, chainProcessor, reorgTaskBatchSize, reorgCheckerInterval)
	taskCleaner := task.NewTaskCleaner(dbClient, cleaningTaskInterval)

	// Schedule pending tasks to fullfilRandomSeeds onchain.
	chainListener.Start()
	log.Info("[Main] Waiting 2 second before staring trackers.")
	time.Sleep(2 * time.Second)
	tracker.Start()
	taskCreator.Start()
	taskSender.Start()
	taskSyncerSent.Start()
	taskReorgChecker.Start()

	// Handle grateful shutdown.
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChannel
		log.Info("Stopping... need to stop all listeners and processors, close all connections and write file  to disk before shutting down.")
		tracker.Stop()
		chainListener.Stop()
		// taskProcessor.Stop()
		taskCreator.Stop()
		taskSender.Stop()
		taskSyncerSent.Stop()
		taskReorgChecker.Stop()
		taskCleaner.Stop()
		ethClient.Close()
		contract.Client.Close()
		dbClient.Close()
	}()
	chainListener.Wait()

}
func random(ctx *cli.Context) {
	fmt.Println("aaa")
	config := config.LoadConfig(ctx.String(ConfigFileFlag.GetName()))
	period := ctx.String(PeriodFlag.GetName())
	prevBeacon := ctx.String(PrevBeaconFlag.GetName())
	log.Info("Loaded config ", "period", period, "Prev Beacon", prevBeacon)
	log.Info("Loaded config", config.PrivateKey)
	res, err := core.GenerateProofResponse(config.VRFConfig.Key, utils.ConvertStringToBigInt(period), utils.ConvertStringToBigInt(prevBeacon))
	if err != nil {
		panic(err)
	}
	log.Info("Output", "res", res)

	outputPath := ctx.String(OutputPathFlag.GetName())
	if outputPath != "" {
		proof := core.PrettyVRFProof{
			[2]string{res.Proof.Pk[0].String(), res.Proof.Pk[1].String()},
			[2]string{res.Proof.Gamma[0].String(), res.Proof.Gamma[1].String()},
			res.Proof.C.String(),
			res.Proof.S.String(),
			res.Proof.Seed.String(),
			res.Proof.UWitness,
			[2]string{res.Proof.CGammaWitness[0].String(), res.Proof.CGammaWitness[1].String()},
			[2]string{res.Proof.SHashWitness[0].String(), res.Proof.SHashWitness[1].String()},
			res.Proof.ZInv.String(),
		}
		file, _ := json.MarshalIndent(proof, "", " ")
		_ = ioutil.WriteFile(outputPath+"/random-result.json", file, 0644)
	}

}
func main() {
	app := cli.NewApp()
	app.Name = "Ronin Random Beacon"
	app.Commands = []cli.Command{
		GenKeyCommand,
		StartCommand,
		RandomCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		log.Error("Failed for starting", err)
	}
}
