package tests

import (
	"math/big"
	"strings"
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

func TestParseBlockNumberWithOptionalDelay(t *testing.T) {
	const (
		bufferBlock = 2
	)
	testCases := []struct {
		input        string
		buffer       uint64
		expected     uint64
		errorMessage string
	}{
		{"0x0", bufferBlock, 0, "invalid hexa string"},
		{"0x11", bufferBlock, 15, ""},
		{"0x1234", bufferBlock, 4658, ""},
		{"  0x1234  ", bufferBlock, 4658, ""},
		{" 1234", bufferBlock, 4658, ""},
		{"0x1234", 0, 4660, ""},
		{"0x1", bufferBlock, 0, "invalid hexa string"},
		{"0x", bufferBlock, 0, "invalid hexa string"},
		{"    ", bufferBlock, 0, "invalid hexa string"},
		{" ", bufferBlock, 0, "invalid hexa string"},
		{"", bufferBlock, 0, "invalid hexa string"},
	}

	for _, tc := range testCases {
		result, err := utils.ParseBlockNumberWithOptionalDelay(tc.input, tc.buffer)
		if result != tc.expected {
			t.Errorf("ConvertHexaStringToIntWithBuffer(%q, %v) = %v; want %v", tc.input, tc.buffer, result, tc.expected)
		}

		if tc.errorMessage == "" {
			if err != nil {
				t.Errorf("ConvertHexaStringToIntWithBuffer(%q, %v) got error %v; want nil ", tc.input, tc.buffer, err)
			}
		} else {
			if !strings.Contains(err.Error(), tc.errorMessage) {
				t.Errorf("ConvertHexaStringToIntWithBuffer(%q, %v) got error %v; want %v ", tc.input, tc.buffer, err, tc.errorMessage)
			}
		}
	}
}
