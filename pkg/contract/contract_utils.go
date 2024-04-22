package contract

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ronin-chain/ronin-random-beacon/pkg/utils"
)

var (
	Client      *ethclient.Client
	Coordinator *RandomBeaconCoordinatorSession
)

func CreateCoordinatorContract(client *ethclient.Client, coordinatorAddr string, privateKey string, chainId *big.Int) (*RandomBeaconCoordinatorSession, error) {
	contract, err := NewRandomBeaconCoordinator(common.HexToAddress(coordinatorAddr), client)
	if err != nil {
		return nil, err
	}
	txOptions, err := utils.GetTxOptions(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	coordinator := RandomBeaconCoordinatorSession{
		Contract:     contract,
		CallOpts:     bind.CallOpts{Pending: true},
		TransactOpts: *txOptions,
	}
	return &coordinator, nil
}

func Init(endpoint string, coordinatorAddr string, privateKey string, chainId *big.Int) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		panic(fmt.Errorf("Failed to connect to the Ethereum client: %w", err))
	}
	Client = client

	contract, err := CreateCoordinatorContract(client, coordinatorAddr, privateKey, chainId)
	if err != nil {
		panic(fmt.Errorf("Failed to create Axie Release contract instance: %w", err))
	}

	Coordinator = contract
}
