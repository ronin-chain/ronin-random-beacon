package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetTxOptions(privateKey string, chainId *big.Int) (*bind.TransactOpts, error) {
	if strings.HasPrefix(privateKey, "0x") {
		privateKey = privateKey[2:]
	}
	// Convert PK to ECDSA Private Key for signing txs.
	ecdsaPk, _ := crypto.HexToECDSA(privateKey)

	txOptions, err := bind.NewKeyedTransactorWithChainID(ecdsaPk, chainId)
	if err != nil {
		return nil, fmt.Errorf("Failed to create authorzied transactor: %w", err)
	}
	return txOptions, nil

}

func ConvertStringToBigInt(s string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(s, 10)
	return n
}

func Min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
