package tests

import (
	"math/big"
	"testing"

	"github.com/ronin-chain/ronin-random-beacon/pkg/utils"
)

func TestConvertStringToBigInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected *big.Int
	}{
		{"1234567890", big.NewInt(1234567890)},
		{"0", big.NewInt(0)},
		{"-1234567890", big.NewInt(-1234567890)},
	}

	for _, tc := range testCases {
		result := utils.ConvertStringToBigInt(tc.input)
		if result.Cmp(tc.expected) != 0 {
			t.Errorf("ConvertStringToBigInt(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		a, b, expected uint64
	}{
		{1, 2, 1},
		{2, 1, 1},
		{2, 2, 2},
	}

	for _, tc := range testCases {
		result := utils.Min(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Min(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		a, b, expected uint64
	}{
		{1, 2, 2},
		{2, 1, 2},
		{2, 2, 2},
	}

	for _, tc := range testCases {
		result := utils.Max(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Max(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expected)
		}
	}
}
