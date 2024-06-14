package core

import (
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/vrfkey"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
)

type PrettyVRFProof struct {
	Pk            [2]string      `json:"pk"`
	Gamma         [2]string      `json:"gamma"`
	C             string         `json:"c"`
	S             string         `json:"s"`
	Seed          string         `json:"seed"`
	UWitness      common.Address `json:"uWitness"`
	CGammaWitness [2]string      `json:"cGammaWitness"`
	SHashWitness  [2]string      `json:"sHashWitness"`
	ZInv          string         `json:"zInv"`
}

func GenerateVRFKey() {
	key, err := vrfkey.NewV2()
	if err != nil {
		panic(err)
	}

	point, err := key.PublicKey.Point()

	if err != nil {
		panic(err)
	}
	x, y := secp256k1.Coordinates(point)
	log.Printf("Generated public key is: %s (x=%v,y=%v)", key.PublicKey, x, y)
	log.Printf("Key hash is: %v", key.PublicKey.MustHash())
	log.Printf("Secret key is: 0x%v", hex.EncodeToString(key.Raw()))
}
