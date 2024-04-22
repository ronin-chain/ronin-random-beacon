package tests

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"testing"

	"log"

	"github.com/ronin-chain/ronin-random-beacon/pkg/config"
	"github.com/ronin-chain/ronin-random-beacon/pkg/core"
)

var curve = elliptic.P256()

func generatePrivateKey() string {
	// Generate a random private key
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Fatal("Failed to generate private key: %w", err)
		return ""
	}
	// Serialize the private key
	privateKeyBytes := privateKey.D.Bytes()
	// Convert the private key to a hex string
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	return privateKeyHex
}

func TestGenerateProofResponse(t *testing.T) {
	// Replace with your actual key and period and prevBeacon values
	c := config.AppConfig{}
	c.SecretKey = generatePrivateKey()
	vrfkey, _ := c.VRFKey()
	period := big.NewInt(456)
	prevBeacon := big.NewInt(789)

	_, err := core.GenerateProofResponse(vrfkey, period, prevBeacon)
	if err != nil {
		t.Errorf("GenerateProofResponse returned an error: %v", err)
	}

	// Add assertions here to verify the behavior of GenerateProofResponse.
	// For example, you might check that the correct proof response was generated.
}
