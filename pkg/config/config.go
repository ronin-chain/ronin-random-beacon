package config

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/ronin-chain/ronin-random-beacon/pkg/utils"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/vrfkey"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/spf13/viper"
)

const (
	FailedMaxAttempts = 10
)

type VRFConfig struct {
	Key     vrfkey.KeyV2
	KeyHash common.Hash
}

type AppConfig struct {
	PrivateKey               string         `mapstructure:"private_key"`
	SecretKey                string         `mapstructure:"secret_key"`
	Verbosity                int            `mapstructure:"verbosity"`
	DatabaseURL              string         `mapstructure:"database_url"`
	SkynetEndpoint           string         `mapstructure:"skynet_endpoint"`
	CoordinatorAddress       string         `mapstructure:"coordinator_address"`
	CoordinatorStartingBlock uint64         `mapstructure:"coordinator_starting_block"`
	RpcEndpoint              string         `mapstructure:"rpc_endpoint"`
	ChainID                  int64          `mapstructure:"chain_id"`
	OracleAddress            common.Address `mapstructure:"oracle_address,omitempty"`
	VRFConfig                VRFConfig      `mapstructure:"vrf,omitempty"`
}

func (c *AppConfig) VRFKey() (vrfkey.KeyV2, common.Hash) {
	if strings.HasPrefix(c.SecretKey, "0x") {
		c.SecretKey = c.SecretKey[2:]
	}
	byteSK, err := hex.DecodeString(c.SecretKey)
	if err != nil {
		log.Error("Invalid secret key")
		panic(err)
	}
	VRFKey := vrfkey.Raw(byteSK).Key()
	publicKey, err := secp256k1.NewPublicKeyFromBytes(VRFKey.PublicKey[:])
	if err != nil {
		log.Error("Invalid PK")
		panic(err)
	}
	return VRFKey, publicKey.MustHash()
}

func (c *AppConfig) SetVRFConfig() {
	vrfkey, keyHash := c.VRFKey()
	c.VRFConfig = VRFConfig{
		Key:     vrfkey,
		KeyHash: keyHash,
	}
}

func SetLogger(config *AppConfig) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(config.Verbosity))
	log.Root().SetHandler(glogger)
}

func LoadConfig(configFilePath string) *AppConfig {
	// get the type from Path's extension, default is yaml
	configType := "yaml"
	exts := strings.Split(configFilePath, ".")
	if len(exts) > 0 {
		ext := strings.ToLower(exts[len(exts)-1])
		if ext == "json" {
			configType = ext
		}
	}
	viper.SetConfigName("config")
	viper.SetConfigType(configType)
	viper.AddConfigPath(configFilePath)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.SetEnvPrefix("RONIN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.BindEnv("private_key", "PRIVATE_KEY")
	viper.BindEnv("secret_key", "SECRET_KEY")
	viper.BindEnv("database_url", "DATABASE_URL")

	config := &AppConfig{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	SetLogger(config)
	config.SetVRFConfig()
	txOptions, err := utils.GetTxOptions(config.PrivateKey, big.NewInt(config.ChainID))

	if err != nil {
		panic(fmt.Errorf("Failed to create authorized transactor: %w", err))
	}
	contract.Init(config.RpcEndpoint, config.CoordinatorAddress, config.PrivateKey, big.NewInt(config.ChainID))
	config.OracleAddress = txOptions.From
	return config
}
