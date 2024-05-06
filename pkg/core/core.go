package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ronin-chain/ronin-random-beacon/pkg/contract"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/vrfkey"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	ChainLinkVrfProof "github.com/smartcontractkit/chainlink/core/services/vrf/proof"
	"github.com/smartcontractkit/chainlink/core/utils"
)

type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

type RandomWordsResponse struct {
	Proof         contract.VRFProof
	RandomRequest contract.RandomRequest
}

var (
	Uint256, _ = abi.NewType("uint256", "", nil)
	Address, _ = abi.NewType("address", "", nil)
	Bytes32, _ = abi.NewType("bytes32", "", nil)
)

var seedArguments = abi.Arguments{
	{
		Name: "period",
		Type: Uint256,
	},
	{
		Name: "prevBeacon",
		Type: Uint256,
	},
	{
		Name: "keyHash",
		Type: Bytes32,
	},
}

func getProofSeed(period *big.Int, prevBeacon *big.Int, keyHash common.Hash) (*big.Int, error) {
	encodedData, err := seedArguments.Pack(
		period,
		prevBeacon,
		keyHash,
	)

	if err != nil {
		return nil, err
	}
	return utils.MustHash(string(encodedData)).Big(), nil

}

func GenerateProofResponse(key vrfkey.KeyV2, period *big.Int, prevBeacon *big.Int, keyHash common.Hash) (*RandomWordsResponse, error) {
	// Generate a seed
	seed, err := getProofSeed(period, prevBeacon, keyHash)

	if err != nil {
		return nil, err

	}
	// Generate a proof
	proof, err := key.GenerateProof(seed)

	if err != nil {
		return nil, err
	}

	solidityProof, err := ChainLinkVrfProof.SolidityPrecalculations(&proof)

	if err != nil {
		return nil, err
	}

	// Set seed to the proof
	solidityProof.P.Seed = seed

	// Get x, y of the point in the Curve
	x, y := secp256k1.Coordinates(solidityProof.P.PublicKey)
	gx, gy := secp256k1.Coordinates(solidityProof.P.Gamma)
	cgx, cgy := secp256k1.Coordinates(solidityProof.CGammaWitness)
	shx, shy := secp256k1.Coordinates(solidityProof.SHashWitness)

	return &RandomWordsResponse{
		Proof: contract.VRFProof{
			Pk:            [2]*big.Int{x, y},
			Gamma:         [2]*big.Int{gx, gy},
			C:             solidityProof.P.C,
			S:             solidityProof.P.S,
			Seed:          seed,
			UWitness:      solidityProof.UWitness,
			CGammaWitness: [2]*big.Int{cgx, cgy},
			SHashWitness:  [2]*big.Int{shx, shy},
			ZInv:          solidityProof.ZInv,
		},
		RandomRequest: contract.RandomRequest{
			Period:     period,
			PrevBeacon: prevBeacon,
		},
	}, nil

}
