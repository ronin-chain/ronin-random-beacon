// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LibSortValidatorsByBeaconRotatingValidatorStorage is an auto generated low-level Go binding around an user-defined struct.
type LibSortValidatorsByBeaconRotatingValidatorStorage struct {
	Cid    common.Address
	Staked *big.Int
}

// LibSortValidatorsByBeaconSortedValidatorStorage is an auto generated low-level Go binding around an user-defined struct.
type LibSortValidatorsByBeaconSortedValidatorStorage struct {
	NRV                   uint16
	NonRotatingValidators []common.Address
	RotatingValidators    []LibSortValidatorsByBeaconRotatingValidatorStorage
}

// LibSortValidatorsByBeaconUnsortedValidatorStorage is an auto generated low-level Go binding around an user-defined struct.
type LibSortValidatorsByBeaconUnsortedValidatorStorage struct {
	Cids []common.Address
}

// LibSortValidatorsByBeaconValidatorStorage is an auto generated low-level Go binding around an user-defined struct.
type LibSortValidatorsByBeaconValidatorStorage struct {
	PickAll  bool
	Sorted   LibSortValidatorsByBeaconSortedValidatorStorage
	Unsorted LibSortValidatorsByBeaconUnsortedValidatorStorage
}

// RandomRequest is an auto generated low-level Go binding around an user-defined struct.
type RandomRequest struct {
	Period            *big.Int
	PrevBeacon        *big.Int
	ChainId           *big.Int
	VerifyingContract common.Address
}

// VRFProof is an auto generated low-level Go binding around an user-defined struct.
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

// RandomBeaconCoordinatorMetaData contains all meta data concerning the RandomBeaconCoordinator contract.
var RandomBeaconCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"COOLDOWN_PERIOD_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADDITION_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PERIOD_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bulkSetValidatorThresholds\",\"inputs\":[{\"name\":\"validatorTypes\",\"type\":\"uint8[]\",\"internalType\":\"enumIRandomBeacon.ValidatorType[]\"},{\"name\":\"thresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calcKeyHash\",\"inputs\":[{\"name\":\"publicKeys\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"keyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"execRecordAndSlashUnavailability\",\"inputs\":[{\"name\":\"lastUpdatedPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashIndicator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allCids\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execRequestRandomSeedForNextPeriod\",\"inputs\":[{\"name\":\"lastUpdatedPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execWrapUpBeaconPeriod\",\"inputs\":[{\"name\":\"lastUpdatedPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allCids\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fulfillRandomSeed\",\"inputs\":[{\"name\":\"req\",\"type\":\"tuple\",\"internalType\":\"structRandomRequest\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevBeacon\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structVRF.Proof\",\"components\":[{\"name\":\"pk\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gamma\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"c\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uWitness\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"sHashWitness\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"zInv\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActivatedAtPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBeaconData\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"submissionCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContract\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"uint8\",\"internalType\":\"enumContractType\"}],\"outputs\":[{\"name\":\"contract_\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastFinalizedPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequestHash\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"reqHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSavedValidatorSet\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structLibSortValidatorsByBeacon.ValidatorStorage\",\"components\":[{\"name\":\"_pickAll\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_sorted\",\"type\":\"tuple\",\"internalType\":\"structLibSortValidatorsByBeacon.SortedValidatorStorage\",\"components\":[{\"name\":\"_nRV\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_nonRotatingValidators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_rotatingValidators\",\"type\":\"tuple[]\",\"internalType\":\"structLibSortValidatorsByBeacon.RotatingValidatorStorage[]\",\"components\":[{\"name\":\"_cid\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_staked\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]},{\"name\":\"_unsorted\",\"type\":\"tuple\",\"internalType\":\"structLibSortValidatorsByBeacon.UnsortedValidatorStorage\",\"components\":[{\"name\":\"_cids\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSelectedValidatorSet\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pickedCids\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnavailabilityCount\",\"inputs\":[{\"name\":\"consensus\",\"type\":\"address\",\"internalType\":\"TConsensus\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnavailabilityCountById\",\"inputs\":[{\"name\":\"cid\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnavailabilitySlashThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorThreshold\",\"inputs\":[{\"name\":\"validatorType\",\"type\":\"uint8\",\"internalType\":\"enumIRandomBeacon.ValidatorType\"}],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"profile\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"staking\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"trustedOrg\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validatorSet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activatedAtPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validatorTypes\",\"type\":\"uint8[]\",\"internalType\":\"enumIRandomBeacon.ValidatorType[]\"},{\"name\":\"thresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isSubmittedAt\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"consensus\",\"type\":\"address\",\"internalType\":\"TConsensus\"}],\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSubmittedAtById\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cid\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSubmittedAtByKeyHash\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"submitted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pickValidatorSetForCurrentPeriod\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pickedCids\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setContract\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"uint8\",\"internalType\":\"enumContractType\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnavailabilitySlashThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BeaconFinalized\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractUpdated\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumContractType\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Migrated\",\"inputs\":[{\"name\":\"curr\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structLibSortValidatorsByBeacon.ValidatorStorage\",\"components\":[{\"name\":\"_pickAll\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_sorted\",\"type\":\"tuple\",\"internalType\":\"structLibSortValidatorsByBeacon.SortedValidatorStorage\",\"components\":[{\"name\":\"_nRV\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_nonRotatingValidators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_rotatingValidators\",\"type\":\"tuple[]\",\"internalType\":\"structLibSortValidatorsByBeacon.RotatingValidatorStorage[]\",\"components\":[{\"name\":\"_cid\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_staked\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]},{\"name\":\"_unsorted\",\"type\":\"tuple\",\"internalType\":\"structLibSortValidatorsByBeacon.UnsortedValidatorStorage\",\"components\":[{\"name\":\"_cids\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomSeedFulfilled\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"reqHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomSeedRequested\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"reqHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"req\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRandomRequest\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevBeacon\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashUnavailabilityThresholdUpdated\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSetSaved\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"pickedAll\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"nRV\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonRotatingValidators\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"rotatingValidators\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"rotatingStakeAmounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorThresholdUpdated\",\"inputs\":[{\"name\":\"validatorType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIRandomBeacon.ValidatorType\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ErrAlreadyFinalizedBeacon\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrContractTypeNotFound\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"uint8\",\"internalType\":\"enumContractType\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidChainId\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidKeyHash\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidRandomRequest\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidThresholdConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidVerifyingContract\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actual\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ErrLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrLengthMismatch\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrNotActivated\",\"inputs\":[{\"name\":\"untilPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrNotEndedChangeKeyHashCooldown\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrNotEndedRegisterCooldown\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrNotFinalizedBeacon\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrOutOfRange\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrUnauthorized\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"expectedRole\",\"type\":\"uint8\",\"internalType\":\"enumRoleAccess\"}]},{\"type\":\"error\",\"name\":\"ErrUnauthorizedCall\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrUnexpectedInternalCall\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"expectedContractType\",\"type\":\"uint8\",\"internalType\":\"enumContractType\"},{\"name\":\"actual\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ErrZeroCodeContract\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// RandomBeaconCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomBeaconCoordinatorMetaData.ABI instead.
var RandomBeaconCoordinatorABI = RandomBeaconCoordinatorMetaData.ABI

// RandomBeaconCoordinator is an auto generated Go binding around an Ethereum contract.
type RandomBeaconCoordinator struct {
	RandomBeaconCoordinatorCaller     // Read-only binding to the contract
	RandomBeaconCoordinatorTransactor // Write-only binding to the contract
	RandomBeaconCoordinatorFilterer   // Log filterer for contract events
}

// RandomBeaconCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomBeaconCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomBeaconCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomBeaconCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomBeaconCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomBeaconCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomBeaconCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomBeaconCoordinatorSession struct {
	Contract     *RandomBeaconCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RandomBeaconCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomBeaconCoordinatorCallerSession struct {
	Contract *RandomBeaconCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// RandomBeaconCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomBeaconCoordinatorTransactorSession struct {
	Contract     *RandomBeaconCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// RandomBeaconCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomBeaconCoordinatorRaw struct {
	Contract *RandomBeaconCoordinator // Generic contract binding to access the raw methods on
}

// RandomBeaconCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomBeaconCoordinatorCallerRaw struct {
	Contract *RandomBeaconCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// RandomBeaconCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomBeaconCoordinatorTransactorRaw struct {
	Contract *RandomBeaconCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomBeaconCoordinator creates a new instance of RandomBeaconCoordinator, bound to a specific deployed contract.
func NewRandomBeaconCoordinator(address common.Address, backend bind.ContractBackend) (*RandomBeaconCoordinator, error) {
	contract, err := bindRandomBeaconCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinator{RandomBeaconCoordinatorCaller: RandomBeaconCoordinatorCaller{contract: contract}, RandomBeaconCoordinatorTransactor: RandomBeaconCoordinatorTransactor{contract: contract}, RandomBeaconCoordinatorFilterer: RandomBeaconCoordinatorFilterer{contract: contract}}, nil
}

// NewRandomBeaconCoordinatorCaller creates a new read-only instance of RandomBeaconCoordinator, bound to a specific deployed contract.
func NewRandomBeaconCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*RandomBeaconCoordinatorCaller, error) {
	contract, err := bindRandomBeaconCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorCaller{contract: contract}, nil
}

// NewRandomBeaconCoordinatorTransactor creates a new write-only instance of RandomBeaconCoordinator, bound to a specific deployed contract.
func NewRandomBeaconCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomBeaconCoordinatorTransactor, error) {
	contract, err := bindRandomBeaconCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorTransactor{contract: contract}, nil
}

// NewRandomBeaconCoordinatorFilterer creates a new log filterer instance of RandomBeaconCoordinator, bound to a specific deployed contract.
func NewRandomBeaconCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomBeaconCoordinatorFilterer, error) {
	contract, err := bindRandomBeaconCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorFilterer{contract: contract}, nil
}

// bindRandomBeaconCoordinator binds a generic wrapper to an already deployed contract.
func bindRandomBeaconCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomBeaconCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomBeaconCoordinator.Contract.RandomBeaconCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.RandomBeaconCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.RandomBeaconCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomBeaconCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.contract.Transact(opts, method, params...)
}

// COOLDOWNPERIODTHRESHOLD is a free data retrieval call binding the contract method 0xe7d46997.
//
// Solidity: function COOLDOWN_PERIOD_THRESHOLD() pure returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) COOLDOWNPERIODTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "COOLDOWN_PERIOD_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COOLDOWNPERIODTHRESHOLD is a free data retrieval call binding the contract method 0xe7d46997.
//
// Solidity: function COOLDOWN_PERIOD_THRESHOLD() pure returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) COOLDOWNPERIODTHRESHOLD() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.COOLDOWNPERIODTHRESHOLD(&_RandomBeaconCoordinator.CallOpts)
}

// COOLDOWNPERIODTHRESHOLD is a free data retrieval call binding the contract method 0xe7d46997.
//
// Solidity: function COOLDOWN_PERIOD_THRESHOLD() pure returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) COOLDOWNPERIODTHRESHOLD() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.COOLDOWNPERIODTHRESHOLD(&_RandomBeaconCoordinator.CallOpts)
}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) DEFAULTADDITIONGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "DEFAULT_ADDITION_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) DEFAULTADDITIONGAS() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.DEFAULTADDITIONGAS(&_RandomBeaconCoordinator.CallOpts)
}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) DEFAULTADDITIONGAS() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.DEFAULTADDITIONGAS(&_RandomBeaconCoordinator.CallOpts)
}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) PERIODDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "PERIOD_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) PERIODDURATION() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.PERIODDURATION(&_RandomBeaconCoordinator.CallOpts)
}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) PERIODDURATION() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.PERIODDURATION(&_RandomBeaconCoordinator.CallOpts)
}

// CalcKeyHash is a free data retrieval call binding the contract method 0xe78bb627.
//
// Solidity: function calcKeyHash(uint256[2] publicKeys) pure returns(bytes32 keyHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) CalcKeyHash(opts *bind.CallOpts, publicKeys [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "calcKeyHash", publicKeys)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalcKeyHash is a free data retrieval call binding the contract method 0xe78bb627.
//
// Solidity: function calcKeyHash(uint256[2] publicKeys) pure returns(bytes32 keyHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) CalcKeyHash(publicKeys [2]*big.Int) ([32]byte, error) {
	return _RandomBeaconCoordinator.Contract.CalcKeyHash(&_RandomBeaconCoordinator.CallOpts, publicKeys)
}

// CalcKeyHash is a free data retrieval call binding the contract method 0xe78bb627.
//
// Solidity: function calcKeyHash(uint256[2] publicKeys) pure returns(bytes32 keyHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) CalcKeyHash(publicKeys [2]*big.Int) ([32]byte, error) {
	return _RandomBeaconCoordinator.Contract.CalcKeyHash(&_RandomBeaconCoordinator.CallOpts, publicKeys)
}

// GetActivatedAtPeriod is a free data retrieval call binding the contract method 0xd1bd2f01.
//
// Solidity: function getActivatedAtPeriod() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetActivatedAtPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getActivatedAtPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActivatedAtPeriod is a free data retrieval call binding the contract method 0xd1bd2f01.
//
// Solidity: function getActivatedAtPeriod() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetActivatedAtPeriod() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetActivatedAtPeriod(&_RandomBeaconCoordinator.CallOpts)
}

// GetActivatedAtPeriod is a free data retrieval call binding the contract method 0xd1bd2f01.
//
// Solidity: function getActivatedAtPeriod() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetActivatedAtPeriod() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetActivatedAtPeriod(&_RandomBeaconCoordinator.CallOpts)
}

// GetBeaconData is a free data retrieval call binding the contract method 0x80d12710.
//
// Solidity: function getBeaconData(uint256 period) view returns(uint256 value, bool finalized, uint256 submissionCount)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetBeaconData(opts *bind.CallOpts, period *big.Int) (struct {
	Value           *big.Int
	Finalized       bool
	SubmissionCount *big.Int
}, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getBeaconData", period)

	outstruct := new(struct {
		Value           *big.Int
		Finalized       bool
		SubmissionCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Finalized = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.SubmissionCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBeaconData is a free data retrieval call binding the contract method 0x80d12710.
//
// Solidity: function getBeaconData(uint256 period) view returns(uint256 value, bool finalized, uint256 submissionCount)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetBeaconData(period *big.Int) (struct {
	Value           *big.Int
	Finalized       bool
	SubmissionCount *big.Int
}, error) {
	return _RandomBeaconCoordinator.Contract.GetBeaconData(&_RandomBeaconCoordinator.CallOpts, period)
}

// GetBeaconData is a free data retrieval call binding the contract method 0x80d12710.
//
// Solidity: function getBeaconData(uint256 period) view returns(uint256 value, bool finalized, uint256 submissionCount)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetBeaconData(period *big.Int) (struct {
	Value           *big.Int
	Finalized       bool
	SubmissionCount *big.Int
}, error) {
	return _RandomBeaconCoordinator.Contract.GetBeaconData(&_RandomBeaconCoordinator.CallOpts, period)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetContract(opts *bind.CallOpts, contractType uint8) (common.Address, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getContract", contractType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetContract(contractType uint8) (common.Address, error) {
	return _RandomBeaconCoordinator.Contract.GetContract(&_RandomBeaconCoordinator.CallOpts, contractType)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetContract(contractType uint8) (common.Address, error) {
	return _RandomBeaconCoordinator.Contract.GetContract(&_RandomBeaconCoordinator.CallOpts, contractType)
}

// GetLastFinalizedPeriod is a free data retrieval call binding the contract method 0x11ed98f7.
//
// Solidity: function getLastFinalizedPeriod() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetLastFinalizedPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getLastFinalizedPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastFinalizedPeriod is a free data retrieval call binding the contract method 0x11ed98f7.
//
// Solidity: function getLastFinalizedPeriod() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetLastFinalizedPeriod() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetLastFinalizedPeriod(&_RandomBeaconCoordinator.CallOpts)
}

// GetLastFinalizedPeriod is a free data retrieval call binding the contract method 0x11ed98f7.
//
// Solidity: function getLastFinalizedPeriod() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetLastFinalizedPeriod() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetLastFinalizedPeriod(&_RandomBeaconCoordinator.CallOpts)
}

// GetRequestHash is a free data retrieval call binding the contract method 0x35188e31.
//
// Solidity: function getRequestHash(uint256 period) view returns(bytes32 reqHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetRequestHash(opts *bind.CallOpts, period *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getRequestHash", period)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRequestHash is a free data retrieval call binding the contract method 0x35188e31.
//
// Solidity: function getRequestHash(uint256 period) view returns(bytes32 reqHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetRequestHash(period *big.Int) ([32]byte, error) {
	return _RandomBeaconCoordinator.Contract.GetRequestHash(&_RandomBeaconCoordinator.CallOpts, period)
}

// GetRequestHash is a free data retrieval call binding the contract method 0x35188e31.
//
// Solidity: function getRequestHash(uint256 period) view returns(bytes32 reqHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetRequestHash(period *big.Int) ([32]byte, error) {
	return _RandomBeaconCoordinator.Contract.GetRequestHash(&_RandomBeaconCoordinator.CallOpts, period)
}

// GetSavedValidatorSet is a free data retrieval call binding the contract method 0xd0ddb12b.
//
// Solidity: function getSavedValidatorSet(uint256 period) view returns((bool,(uint16,address[],(address,uint96)[]),(address[])))
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetSavedValidatorSet(opts *bind.CallOpts, period *big.Int) (LibSortValidatorsByBeaconValidatorStorage, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getSavedValidatorSet", period)

	if err != nil {
		return *new(LibSortValidatorsByBeaconValidatorStorage), err
	}

	out0 := *abi.ConvertType(out[0], new(LibSortValidatorsByBeaconValidatorStorage)).(*LibSortValidatorsByBeaconValidatorStorage)

	return out0, err

}

// GetSavedValidatorSet is a free data retrieval call binding the contract method 0xd0ddb12b.
//
// Solidity: function getSavedValidatorSet(uint256 period) view returns((bool,(uint16,address[],(address,uint96)[]),(address[])))
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetSavedValidatorSet(period *big.Int) (LibSortValidatorsByBeaconValidatorStorage, error) {
	return _RandomBeaconCoordinator.Contract.GetSavedValidatorSet(&_RandomBeaconCoordinator.CallOpts, period)
}

// GetSavedValidatorSet is a free data retrieval call binding the contract method 0xd0ddb12b.
//
// Solidity: function getSavedValidatorSet(uint256 period) view returns((bool,(uint16,address[],(address,uint96)[]),(address[])))
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetSavedValidatorSet(period *big.Int) (LibSortValidatorsByBeaconValidatorStorage, error) {
	return _RandomBeaconCoordinator.Contract.GetSavedValidatorSet(&_RandomBeaconCoordinator.CallOpts, period)
}

// GetSelectedValidatorSet is a free data retrieval call binding the contract method 0x16519133.
//
// Solidity: function getSelectedValidatorSet(uint256 period, uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetSelectedValidatorSet(opts *bind.CallOpts, period *big.Int, epoch *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getSelectedValidatorSet", period, epoch)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSelectedValidatorSet is a free data retrieval call binding the contract method 0x16519133.
//
// Solidity: function getSelectedValidatorSet(uint256 period, uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetSelectedValidatorSet(period *big.Int, epoch *big.Int) ([]common.Address, error) {
	return _RandomBeaconCoordinator.Contract.GetSelectedValidatorSet(&_RandomBeaconCoordinator.CallOpts, period, epoch)
}

// GetSelectedValidatorSet is a free data retrieval call binding the contract method 0x16519133.
//
// Solidity: function getSelectedValidatorSet(uint256 period, uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetSelectedValidatorSet(period *big.Int, epoch *big.Int) ([]common.Address, error) {
	return _RandomBeaconCoordinator.Contract.GetSelectedValidatorSet(&_RandomBeaconCoordinator.CallOpts, period, epoch)
}

// GetUnavailabilityCount is a free data retrieval call binding the contract method 0x3a570332.
//
// Solidity: function getUnavailabilityCount(address consensus) view returns(uint256 count)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetUnavailabilityCount(opts *bind.CallOpts, consensus common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getUnavailabilityCount", consensus)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnavailabilityCount is a free data retrieval call binding the contract method 0x3a570332.
//
// Solidity: function getUnavailabilityCount(address consensus) view returns(uint256 count)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetUnavailabilityCount(consensus common.Address) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilityCount(&_RandomBeaconCoordinator.CallOpts, consensus)
}

// GetUnavailabilityCount is a free data retrieval call binding the contract method 0x3a570332.
//
// Solidity: function getUnavailabilityCount(address consensus) view returns(uint256 count)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetUnavailabilityCount(consensus common.Address) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilityCount(&_RandomBeaconCoordinator.CallOpts, consensus)
}

// GetUnavailabilityCountById is a free data retrieval call binding the contract method 0x65291db7.
//
// Solidity: function getUnavailabilityCountById(address cid) view returns(uint256 count)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetUnavailabilityCountById(opts *bind.CallOpts, cid common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getUnavailabilityCountById", cid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnavailabilityCountById is a free data retrieval call binding the contract method 0x65291db7.
//
// Solidity: function getUnavailabilityCountById(address cid) view returns(uint256 count)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetUnavailabilityCountById(cid common.Address) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilityCountById(&_RandomBeaconCoordinator.CallOpts, cid)
}

// GetUnavailabilityCountById is a free data retrieval call binding the contract method 0x65291db7.
//
// Solidity: function getUnavailabilityCountById(address cid) view returns(uint256 count)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetUnavailabilityCountById(cid common.Address) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilityCountById(&_RandomBeaconCoordinator.CallOpts, cid)
}

// GetUnavailabilitySlashThreshold is a free data retrieval call binding the contract method 0x917a65bd.
//
// Solidity: function getUnavailabilitySlashThreshold() view returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetUnavailabilitySlashThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getUnavailabilitySlashThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnavailabilitySlashThreshold is a free data retrieval call binding the contract method 0x917a65bd.
//
// Solidity: function getUnavailabilitySlashThreshold() view returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetUnavailabilitySlashThreshold() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilitySlashThreshold(&_RandomBeaconCoordinator.CallOpts)
}

// GetUnavailabilitySlashThreshold is a free data retrieval call binding the contract method 0x917a65bd.
//
// Solidity: function getUnavailabilitySlashThreshold() view returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetUnavailabilitySlashThreshold() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilitySlashThreshold(&_RandomBeaconCoordinator.CallOpts)
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0x035d0062.
//
// Solidity: function getValidatorThreshold(uint8 validatorType) view returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetValidatorThreshold(opts *bind.CallOpts, validatorType uint8) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getValidatorThreshold", validatorType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0x035d0062.
//
// Solidity: function getValidatorThreshold(uint8 validatorType) view returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetValidatorThreshold(validatorType uint8) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetValidatorThreshold(&_RandomBeaconCoordinator.CallOpts, validatorType)
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0x035d0062.
//
// Solidity: function getValidatorThreshold(uint8 validatorType) view returns(uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetValidatorThreshold(validatorType uint8) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetValidatorThreshold(&_RandomBeaconCoordinator.CallOpts, validatorType)
}

// IsSubmittedAt is a free data retrieval call binding the contract method 0xe02f6271.
//
// Solidity: function isSubmittedAt(uint256 period, address consensus) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) IsSubmittedAt(opts *bind.CallOpts, period *big.Int, consensus common.Address) (bool, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "isSubmittedAt", period, consensus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubmittedAt is a free data retrieval call binding the contract method 0xe02f6271.
//
// Solidity: function isSubmittedAt(uint256 period, address consensus) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) IsSubmittedAt(period *big.Int, consensus common.Address) (bool, error) {
	return _RandomBeaconCoordinator.Contract.IsSubmittedAt(&_RandomBeaconCoordinator.CallOpts, period, consensus)
}

// IsSubmittedAt is a free data retrieval call binding the contract method 0xe02f6271.
//
// Solidity: function isSubmittedAt(uint256 period, address consensus) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) IsSubmittedAt(period *big.Int, consensus common.Address) (bool, error) {
	return _RandomBeaconCoordinator.Contract.IsSubmittedAt(&_RandomBeaconCoordinator.CallOpts, period, consensus)
}

// IsSubmittedAtById is a free data retrieval call binding the contract method 0x46337c2f.
//
// Solidity: function isSubmittedAtById(uint256 period, address cid) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) IsSubmittedAtById(opts *bind.CallOpts, period *big.Int, cid common.Address) (bool, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "isSubmittedAtById", period, cid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubmittedAtById is a free data retrieval call binding the contract method 0x46337c2f.
//
// Solidity: function isSubmittedAtById(uint256 period, address cid) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) IsSubmittedAtById(period *big.Int, cid common.Address) (bool, error) {
	return _RandomBeaconCoordinator.Contract.IsSubmittedAtById(&_RandomBeaconCoordinator.CallOpts, period, cid)
}

// IsSubmittedAtById is a free data retrieval call binding the contract method 0x46337c2f.
//
// Solidity: function isSubmittedAtById(uint256 period, address cid) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) IsSubmittedAtById(period *big.Int, cid common.Address) (bool, error) {
	return _RandomBeaconCoordinator.Contract.IsSubmittedAtById(&_RandomBeaconCoordinator.CallOpts, period, cid)
}

// IsSubmittedAtByKeyHash is a free data retrieval call binding the contract method 0xeb0645c4.
//
// Solidity: function isSubmittedAtByKeyHash(uint256 period, bytes32 keyHash) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) IsSubmittedAtByKeyHash(opts *bind.CallOpts, period *big.Int, keyHash [32]byte) (bool, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "isSubmittedAtByKeyHash", period, keyHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubmittedAtByKeyHash is a free data retrieval call binding the contract method 0xeb0645c4.
//
// Solidity: function isSubmittedAtByKeyHash(uint256 period, bytes32 keyHash) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) IsSubmittedAtByKeyHash(period *big.Int, keyHash [32]byte) (bool, error) {
	return _RandomBeaconCoordinator.Contract.IsSubmittedAtByKeyHash(&_RandomBeaconCoordinator.CallOpts, period, keyHash)
}

// IsSubmittedAtByKeyHash is a free data retrieval call binding the contract method 0xeb0645c4.
//
// Solidity: function isSubmittedAtByKeyHash(uint256 period, bytes32 keyHash) view returns(bool submitted)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) IsSubmittedAtByKeyHash(period *big.Int, keyHash [32]byte) (bool, error) {
	return _RandomBeaconCoordinator.Contract.IsSubmittedAtByKeyHash(&_RandomBeaconCoordinator.CallOpts, period, keyHash)
}

// PickValidatorSetForCurrentPeriod is a free data retrieval call binding the contract method 0x7931e9ec.
//
// Solidity: function pickValidatorSetForCurrentPeriod(uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) PickValidatorSetForCurrentPeriod(opts *bind.CallOpts, epoch *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "pickValidatorSetForCurrentPeriod", epoch)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// PickValidatorSetForCurrentPeriod is a free data retrieval call binding the contract method 0x7931e9ec.
//
// Solidity: function pickValidatorSetForCurrentPeriod(uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) PickValidatorSetForCurrentPeriod(epoch *big.Int) ([]common.Address, error) {
	return _RandomBeaconCoordinator.Contract.PickValidatorSetForCurrentPeriod(&_RandomBeaconCoordinator.CallOpts, epoch)
}

// PickValidatorSetForCurrentPeriod is a free data retrieval call binding the contract method 0x7931e9ec.
//
// Solidity: function pickValidatorSetForCurrentPeriod(uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) PickValidatorSetForCurrentPeriod(epoch *big.Int) ([]common.Address, error) {
	return _RandomBeaconCoordinator.Contract.PickValidatorSetForCurrentPeriod(&_RandomBeaconCoordinator.CallOpts, epoch)
}

// BulkSetValidatorThresholds is a paid mutator transaction binding the contract method 0xf5a77abb.
//
// Solidity: function bulkSetValidatorThresholds(uint8[] validatorTypes, uint256[] thresholds) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) BulkSetValidatorThresholds(opts *bind.TransactOpts, validatorTypes []uint8, thresholds []*big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "bulkSetValidatorThresholds", validatorTypes, thresholds)
}

// BulkSetValidatorThresholds is a paid mutator transaction binding the contract method 0xf5a77abb.
//
// Solidity: function bulkSetValidatorThresholds(uint8[] validatorTypes, uint256[] thresholds) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) BulkSetValidatorThresholds(validatorTypes []uint8, thresholds []*big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.BulkSetValidatorThresholds(&_RandomBeaconCoordinator.TransactOpts, validatorTypes, thresholds)
}

// BulkSetValidatorThresholds is a paid mutator transaction binding the contract method 0xf5a77abb.
//
// Solidity: function bulkSetValidatorThresholds(uint8[] validatorTypes, uint256[] thresholds) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) BulkSetValidatorThresholds(validatorTypes []uint8, thresholds []*big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.BulkSetValidatorThresholds(&_RandomBeaconCoordinator.TransactOpts, validatorTypes, thresholds)
}

// ExecRecordAndSlashUnavailability is a paid mutator transaction binding the contract method 0x3ba1de37.
//
// Solidity: function execRecordAndSlashUnavailability(uint256 lastUpdatedPeriod, uint256 newPeriod, address slashIndicator, address[] allCids) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) ExecRecordAndSlashUnavailability(opts *bind.TransactOpts, lastUpdatedPeriod *big.Int, newPeriod *big.Int, slashIndicator common.Address, allCids []common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "execRecordAndSlashUnavailability", lastUpdatedPeriod, newPeriod, slashIndicator, allCids)
}

// ExecRecordAndSlashUnavailability is a paid mutator transaction binding the contract method 0x3ba1de37.
//
// Solidity: function execRecordAndSlashUnavailability(uint256 lastUpdatedPeriod, uint256 newPeriod, address slashIndicator, address[] allCids) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) ExecRecordAndSlashUnavailability(lastUpdatedPeriod *big.Int, newPeriod *big.Int, slashIndicator common.Address, allCids []common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.ExecRecordAndSlashUnavailability(&_RandomBeaconCoordinator.TransactOpts, lastUpdatedPeriod, newPeriod, slashIndicator, allCids)
}

// ExecRecordAndSlashUnavailability is a paid mutator transaction binding the contract method 0x3ba1de37.
//
// Solidity: function execRecordAndSlashUnavailability(uint256 lastUpdatedPeriod, uint256 newPeriod, address slashIndicator, address[] allCids) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) ExecRecordAndSlashUnavailability(lastUpdatedPeriod *big.Int, newPeriod *big.Int, slashIndicator common.Address, allCids []common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.ExecRecordAndSlashUnavailability(&_RandomBeaconCoordinator.TransactOpts, lastUpdatedPeriod, newPeriod, slashIndicator, allCids)
}

// ExecRequestRandomSeedForNextPeriod is a paid mutator transaction binding the contract method 0xe3a43765.
//
// Solidity: function execRequestRandomSeedForNextPeriod(uint256 lastUpdatedPeriod, uint256 newPeriod) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) ExecRequestRandomSeedForNextPeriod(opts *bind.TransactOpts, lastUpdatedPeriod *big.Int, newPeriod *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "execRequestRandomSeedForNextPeriod", lastUpdatedPeriod, newPeriod)
}

// ExecRequestRandomSeedForNextPeriod is a paid mutator transaction binding the contract method 0xe3a43765.
//
// Solidity: function execRequestRandomSeedForNextPeriod(uint256 lastUpdatedPeriod, uint256 newPeriod) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) ExecRequestRandomSeedForNextPeriod(lastUpdatedPeriod *big.Int, newPeriod *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.ExecRequestRandomSeedForNextPeriod(&_RandomBeaconCoordinator.TransactOpts, lastUpdatedPeriod, newPeriod)
}

// ExecRequestRandomSeedForNextPeriod is a paid mutator transaction binding the contract method 0xe3a43765.
//
// Solidity: function execRequestRandomSeedForNextPeriod(uint256 lastUpdatedPeriod, uint256 newPeriod) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) ExecRequestRandomSeedForNextPeriod(lastUpdatedPeriod *big.Int, newPeriod *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.ExecRequestRandomSeedForNextPeriod(&_RandomBeaconCoordinator.TransactOpts, lastUpdatedPeriod, newPeriod)
}

// ExecWrapUpBeaconPeriod is a paid mutator transaction binding the contract method 0x79ec25ae.
//
// Solidity: function execWrapUpBeaconPeriod(uint256 lastUpdatedPeriod, uint256 newPeriod, address[] allCids) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) ExecWrapUpBeaconPeriod(opts *bind.TransactOpts, lastUpdatedPeriod *big.Int, newPeriod *big.Int, allCids []common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "execWrapUpBeaconPeriod", lastUpdatedPeriod, newPeriod, allCids)
}

// ExecWrapUpBeaconPeriod is a paid mutator transaction binding the contract method 0x79ec25ae.
//
// Solidity: function execWrapUpBeaconPeriod(uint256 lastUpdatedPeriod, uint256 newPeriod, address[] allCids) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) ExecWrapUpBeaconPeriod(lastUpdatedPeriod *big.Int, newPeriod *big.Int, allCids []common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.ExecWrapUpBeaconPeriod(&_RandomBeaconCoordinator.TransactOpts, lastUpdatedPeriod, newPeriod, allCids)
}

// ExecWrapUpBeaconPeriod is a paid mutator transaction binding the contract method 0x79ec25ae.
//
// Solidity: function execWrapUpBeaconPeriod(uint256 lastUpdatedPeriod, uint256 newPeriod, address[] allCids) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) ExecWrapUpBeaconPeriod(lastUpdatedPeriod *big.Int, newPeriod *big.Int, allCids []common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.ExecWrapUpBeaconPeriod(&_RandomBeaconCoordinator.TransactOpts, lastUpdatedPeriod, newPeriod, allCids)
}

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0xf1a7340b.
//
// Solidity: function fulfillRandomSeed((uint256,uint256,uint256,address) req, (uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) FulfillRandomSeed(opts *bind.TransactOpts, req RandomRequest, proof VRFProof) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "fulfillRandomSeed", req, proof)
}

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0xf1a7340b.
//
// Solidity: function fulfillRandomSeed((uint256,uint256,uint256,address) req, (uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) FulfillRandomSeed(req RandomRequest, proof VRFProof) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.FulfillRandomSeed(&_RandomBeaconCoordinator.TransactOpts, req, proof)
}

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0xf1a7340b.
//
// Solidity: function fulfillRandomSeed((uint256,uint256,uint256,address) req, (uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) FulfillRandomSeed(req RandomRequest, proof VRFProof) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.FulfillRandomSeed(&_RandomBeaconCoordinator.TransactOpts, req, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c7300e3.
//
// Solidity: function initialize(address profile, address staking, address trustedOrg, address validatorSet, uint256 slashThreshold, uint256 activatedAtPeriod, uint8[] validatorTypes, uint256[] thresholds) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) Initialize(opts *bind.TransactOpts, profile common.Address, staking common.Address, trustedOrg common.Address, validatorSet common.Address, slashThreshold *big.Int, activatedAtPeriod *big.Int, validatorTypes []uint8, thresholds []*big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "initialize", profile, staking, trustedOrg, validatorSet, slashThreshold, activatedAtPeriod, validatorTypes, thresholds)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c7300e3.
//
// Solidity: function initialize(address profile, address staking, address trustedOrg, address validatorSet, uint256 slashThreshold, uint256 activatedAtPeriod, uint8[] validatorTypes, uint256[] thresholds) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) Initialize(profile common.Address, staking common.Address, trustedOrg common.Address, validatorSet common.Address, slashThreshold *big.Int, activatedAtPeriod *big.Int, validatorTypes []uint8, thresholds []*big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.Initialize(&_RandomBeaconCoordinator.TransactOpts, profile, staking, trustedOrg, validatorSet, slashThreshold, activatedAtPeriod, validatorTypes, thresholds)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c7300e3.
//
// Solidity: function initialize(address profile, address staking, address trustedOrg, address validatorSet, uint256 slashThreshold, uint256 activatedAtPeriod, uint8[] validatorTypes, uint256[] thresholds) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) Initialize(profile common.Address, staking common.Address, trustedOrg common.Address, validatorSet common.Address, slashThreshold *big.Int, activatedAtPeriod *big.Int, validatorTypes []uint8, thresholds []*big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.Initialize(&_RandomBeaconCoordinator.TransactOpts, profile, staking, trustedOrg, validatorSet, slashThreshold, activatedAtPeriod, validatorTypes, thresholds)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) InitializeV2() (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.InitializeV2(&_RandomBeaconCoordinator.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) InitializeV2() (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.InitializeV2(&_RandomBeaconCoordinator.TransactOpts)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) SetContract(opts *bind.TransactOpts, contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "setContract", contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.SetContract(&_RandomBeaconCoordinator.TransactOpts, contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.SetContract(&_RandomBeaconCoordinator.TransactOpts, contractType, addr)
}

// SetUnavailabilitySlashThreshold is a paid mutator transaction binding the contract method 0x3355b0f0.
//
// Solidity: function setUnavailabilitySlashThreshold(uint256 threshold) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) SetUnavailabilitySlashThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "setUnavailabilitySlashThreshold", threshold)
}

// SetUnavailabilitySlashThreshold is a paid mutator transaction binding the contract method 0x3355b0f0.
//
// Solidity: function setUnavailabilitySlashThreshold(uint256 threshold) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) SetUnavailabilitySlashThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.SetUnavailabilitySlashThreshold(&_RandomBeaconCoordinator.TransactOpts, threshold)
}

// SetUnavailabilitySlashThreshold is a paid mutator transaction binding the contract method 0x3355b0f0.
//
// Solidity: function setUnavailabilitySlashThreshold(uint256 threshold) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) SetUnavailabilitySlashThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.SetUnavailabilitySlashThreshold(&_RandomBeaconCoordinator.TransactOpts, threshold)
}

// RandomBeaconCoordinatorBeaconFinalizedIterator is returned from FilterBeaconFinalized and is used to iterate over the raw logs and unpacked data for BeaconFinalized events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorBeaconFinalizedIterator struct {
	Event *RandomBeaconCoordinatorBeaconFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorBeaconFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorBeaconFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorBeaconFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorBeaconFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorBeaconFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorBeaconFinalized represents a BeaconFinalized event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorBeaconFinalized struct {
	Period *big.Int
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconFinalized is a free log retrieval operation binding the contract event 0xd31a863f0b5130ba75e27a3764828141ba7bc06056dfcdade31d12ab52a42784.
//
// Solidity: event BeaconFinalized(uint256 indexed period, uint256 value)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterBeaconFinalized(opts *bind.FilterOpts, period []*big.Int) (*RandomBeaconCoordinatorBeaconFinalizedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "BeaconFinalized", periodRule)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorBeaconFinalizedIterator{contract: _RandomBeaconCoordinator.contract, event: "BeaconFinalized", logs: logs, sub: sub}, nil
}

// WatchBeaconFinalized is a free log subscription operation binding the contract event 0xd31a863f0b5130ba75e27a3764828141ba7bc06056dfcdade31d12ab52a42784.
//
// Solidity: event BeaconFinalized(uint256 indexed period, uint256 value)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchBeaconFinalized(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorBeaconFinalized, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "BeaconFinalized", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorBeaconFinalized)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "BeaconFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconFinalized is a log parse operation binding the contract event 0xd31a863f0b5130ba75e27a3764828141ba7bc06056dfcdade31d12ab52a42784.
//
// Solidity: event BeaconFinalized(uint256 indexed period, uint256 value)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseBeaconFinalized(log types.Log) (*RandomBeaconCoordinatorBeaconFinalized, error) {
	event := new(RandomBeaconCoordinatorBeaconFinalized)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "BeaconFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomBeaconCoordinatorContractUpdatedIterator is returned from FilterContractUpdated and is used to iterate over the raw logs and unpacked data for ContractUpdated events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorContractUpdatedIterator struct {
	Event *RandomBeaconCoordinatorContractUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorContractUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorContractUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorContractUpdated represents a ContractUpdated event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorContractUpdated struct {
	ContractType uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractUpdated is a free log retrieval operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterContractUpdated(opts *bind.FilterOpts, contractType []uint8, addr []common.Address) (*RandomBeaconCoordinatorContractUpdatedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorContractUpdatedIterator{contract: _RandomBeaconCoordinator.contract, event: "ContractUpdated", logs: logs, sub: sub}, nil
}

// WatchContractUpdated is a free log subscription operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchContractUpdated(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorContractUpdated, contractType []uint8, addr []common.Address) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorContractUpdated)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractUpdated is a log parse operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseContractUpdated(log types.Log) (*RandomBeaconCoordinatorContractUpdated, error) {
	event := new(RandomBeaconCoordinatorContractUpdated)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomBeaconCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorInitializedIterator struct {
	Event *RandomBeaconCoordinatorInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorInitialized represents a Initialized event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*RandomBeaconCoordinatorInitializedIterator, error) {

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorInitializedIterator{contract: _RandomBeaconCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorInitialized)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseInitialized(log types.Log) (*RandomBeaconCoordinatorInitialized, error) {
	event := new(RandomBeaconCoordinatorInitialized)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomBeaconCoordinatorMigratedIterator is returned from FilterMigrated and is used to iterate over the raw logs and unpacked data for Migrated events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorMigratedIterator struct {
	Event *RandomBeaconCoordinatorMigrated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorMigrated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorMigrated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorMigrated represents a Migrated event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorMigrated struct {
	Curr LibSortValidatorsByBeaconValidatorStorage
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMigrated is a free log retrieval operation binding the contract event 0x858c70d183bcd79a024da80c93a49f57863e2473168f1bf7d405c20f21c986e4.
//
// Solidity: event Migrated((bool,(uint16,address[],(address,uint96)[]),(address[])) curr)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterMigrated(opts *bind.FilterOpts) (*RandomBeaconCoordinatorMigratedIterator, error) {

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorMigratedIterator{contract: _RandomBeaconCoordinator.contract, event: "Migrated", logs: logs, sub: sub}, nil
}

// WatchMigrated is a free log subscription operation binding the contract event 0x858c70d183bcd79a024da80c93a49f57863e2473168f1bf7d405c20f21c986e4.
//
// Solidity: event Migrated((bool,(uint16,address[],(address,uint96)[]),(address[])) curr)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchMigrated(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorMigrated) (event.Subscription, error) {

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "Migrated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorMigrated)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "Migrated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMigrated is a log parse operation binding the contract event 0x858c70d183bcd79a024da80c93a49f57863e2473168f1bf7d405c20f21c986e4.
//
// Solidity: event Migrated((bool,(uint16,address[],(address,uint96)[]),(address[])) curr)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseMigrated(log types.Log) (*RandomBeaconCoordinatorMigrated, error) {
	event := new(RandomBeaconCoordinatorMigrated)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "Migrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomBeaconCoordinatorRandomSeedFulfilledIterator is returned from FilterRandomSeedFulfilled and is used to iterate over the raw logs and unpacked data for RandomSeedFulfilled events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorRandomSeedFulfilledIterator struct {
	Event *RandomBeaconCoordinatorRandomSeedFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorRandomSeedFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorRandomSeedFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorRandomSeedFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorRandomSeedFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorRandomSeedFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorRandomSeedFulfilled represents a RandomSeedFulfilled event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorRandomSeedFulfilled struct {
	By      common.Address
	Period  *big.Int
	ReqHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRandomSeedFulfilled is a free log retrieval operation binding the contract event 0x9c42f0e233770417c75c168e2ec97a2db8e34b756d12f8e002e90d144f4a16cc.
//
// Solidity: event RandomSeedFulfilled(address indexed by, uint256 indexed period, bytes32 indexed reqHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterRandomSeedFulfilled(opts *bind.FilterOpts, by []common.Address, period []*big.Int, reqHash [][32]byte) (*RandomBeaconCoordinatorRandomSeedFulfilledIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var reqHashRule []interface{}
	for _, reqHashItem := range reqHash {
		reqHashRule = append(reqHashRule, reqHashItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "RandomSeedFulfilled", byRule, periodRule, reqHashRule)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorRandomSeedFulfilledIterator{contract: _RandomBeaconCoordinator.contract, event: "RandomSeedFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomSeedFulfilled is a free log subscription operation binding the contract event 0x9c42f0e233770417c75c168e2ec97a2db8e34b756d12f8e002e90d144f4a16cc.
//
// Solidity: event RandomSeedFulfilled(address indexed by, uint256 indexed period, bytes32 indexed reqHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchRandomSeedFulfilled(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorRandomSeedFulfilled, by []common.Address, period []*big.Int, reqHash [][32]byte) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var reqHashRule []interface{}
	for _, reqHashItem := range reqHash {
		reqHashRule = append(reqHashRule, reqHashItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "RandomSeedFulfilled", byRule, periodRule, reqHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorRandomSeedFulfilled)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "RandomSeedFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomSeedFulfilled is a log parse operation binding the contract event 0x9c42f0e233770417c75c168e2ec97a2db8e34b756d12f8e002e90d144f4a16cc.
//
// Solidity: event RandomSeedFulfilled(address indexed by, uint256 indexed period, bytes32 indexed reqHash)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseRandomSeedFulfilled(log types.Log) (*RandomBeaconCoordinatorRandomSeedFulfilled, error) {
	event := new(RandomBeaconCoordinatorRandomSeedFulfilled)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "RandomSeedFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomBeaconCoordinatorRandomSeedRequestedIterator is returned from FilterRandomSeedRequested and is used to iterate over the raw logs and unpacked data for RandomSeedRequested events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorRandomSeedRequestedIterator struct {
	Event *RandomBeaconCoordinatorRandomSeedRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorRandomSeedRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorRandomSeedRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorRandomSeedRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorRandomSeedRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorRandomSeedRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorRandomSeedRequested represents a RandomSeedRequested event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorRandomSeedRequested struct {
	Period  *big.Int
	ReqHash [32]byte
	Req     RandomRequest
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRandomSeedRequested is a free log retrieval operation binding the contract event 0x6c8f5422289494c55b959cb4d0900ec451d35b3bdf06e9d29cd80bfab61b2438.
//
// Solidity: event RandomSeedRequested(uint256 indexed period, bytes32 indexed reqHash, (uint256,uint256,uint256,address) req)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterRandomSeedRequested(opts *bind.FilterOpts, period []*big.Int, reqHash [][32]byte) (*RandomBeaconCoordinatorRandomSeedRequestedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var reqHashRule []interface{}
	for _, reqHashItem := range reqHash {
		reqHashRule = append(reqHashRule, reqHashItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "RandomSeedRequested", periodRule, reqHashRule)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorRandomSeedRequestedIterator{contract: _RandomBeaconCoordinator.contract, event: "RandomSeedRequested", logs: logs, sub: sub}, nil
}

// WatchRandomSeedRequested is a free log subscription operation binding the contract event 0x6c8f5422289494c55b959cb4d0900ec451d35b3bdf06e9d29cd80bfab61b2438.
//
// Solidity: event RandomSeedRequested(uint256 indexed period, bytes32 indexed reqHash, (uint256,uint256,uint256,address) req)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchRandomSeedRequested(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorRandomSeedRequested, period []*big.Int, reqHash [][32]byte) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var reqHashRule []interface{}
	for _, reqHashItem := range reqHash {
		reqHashRule = append(reqHashRule, reqHashItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "RandomSeedRequested", periodRule, reqHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorRandomSeedRequested)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "RandomSeedRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomSeedRequested is a log parse operation binding the contract event 0x6c8f5422289494c55b959cb4d0900ec451d35b3bdf06e9d29cd80bfab61b2438.
//
// Solidity: event RandomSeedRequested(uint256 indexed period, bytes32 indexed reqHash, (uint256,uint256,uint256,address) req)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseRandomSeedRequested(log types.Log) (*RandomBeaconCoordinatorRandomSeedRequested, error) {
	event := new(RandomBeaconCoordinatorRandomSeedRequested)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "RandomSeedRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdatedIterator is returned from FilterSlashUnavailabilityThresholdUpdated and is used to iterate over the raw logs and unpacked data for SlashUnavailabilityThresholdUpdated events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdatedIterator struct {
	Event *RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated represents a SlashUnavailabilityThresholdUpdated event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSlashUnavailabilityThresholdUpdated is a free log retrieval operation binding the contract event 0xc642d5e4298fca4d2285021791ffdf4afce31ec672299930b7ada0047b15d243.
//
// Solidity: event SlashUnavailabilityThresholdUpdated(uint256 value)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterSlashUnavailabilityThresholdUpdated(opts *bind.FilterOpts) (*RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdatedIterator, error) {

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "SlashUnavailabilityThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdatedIterator{contract: _RandomBeaconCoordinator.contract, event: "SlashUnavailabilityThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashUnavailabilityThresholdUpdated is a free log subscription operation binding the contract event 0xc642d5e4298fca4d2285021791ffdf4afce31ec672299930b7ada0047b15d243.
//
// Solidity: event SlashUnavailabilityThresholdUpdated(uint256 value)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchSlashUnavailabilityThresholdUpdated(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "SlashUnavailabilityThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "SlashUnavailabilityThresholdUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashUnavailabilityThresholdUpdated is a log parse operation binding the contract event 0xc642d5e4298fca4d2285021791ffdf4afce31ec672299930b7ada0047b15d243.
//
// Solidity: event SlashUnavailabilityThresholdUpdated(uint256 value)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseSlashUnavailabilityThresholdUpdated(log types.Log) (*RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated, error) {
	event := new(RandomBeaconCoordinatorSlashUnavailabilityThresholdUpdated)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "SlashUnavailabilityThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomBeaconCoordinatorValidatorSetSavedIterator is returned from FilterValidatorSetSaved and is used to iterate over the raw logs and unpacked data for ValidatorSetSaved events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorValidatorSetSavedIterator struct {
	Event *RandomBeaconCoordinatorValidatorSetSaved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorValidatorSetSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorValidatorSetSaved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorValidatorSetSaved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorValidatorSetSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorValidatorSetSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorValidatorSetSaved represents a ValidatorSetSaved event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorValidatorSetSaved struct {
	Period                *big.Int
	PickedAll             bool
	NRV                   *big.Int
	NonRotatingValidators []common.Address
	RotatingValidators    []common.Address
	RotatingStakeAmounts  []*big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetSaved is a free log retrieval operation binding the contract event 0x32c9c468454fcf853e08d256000b08598539a799e849e0e7d8309c15eeb53ec3.
//
// Solidity: event ValidatorSetSaved(uint256 indexed period, bool pickedAll, uint256 nRV, address[] nonRotatingValidators, address[] rotatingValidators, uint256[] rotatingStakeAmounts)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterValidatorSetSaved(opts *bind.FilterOpts, period []*big.Int) (*RandomBeaconCoordinatorValidatorSetSavedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "ValidatorSetSaved", periodRule)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorValidatorSetSavedIterator{contract: _RandomBeaconCoordinator.contract, event: "ValidatorSetSaved", logs: logs, sub: sub}, nil
}

// WatchValidatorSetSaved is a free log subscription operation binding the contract event 0x32c9c468454fcf853e08d256000b08598539a799e849e0e7d8309c15eeb53ec3.
//
// Solidity: event ValidatorSetSaved(uint256 indexed period, bool pickedAll, uint256 nRV, address[] nonRotatingValidators, address[] rotatingValidators, uint256[] rotatingStakeAmounts)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchValidatorSetSaved(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorValidatorSetSaved, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "ValidatorSetSaved", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorValidatorSetSaved)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "ValidatorSetSaved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSetSaved is a log parse operation binding the contract event 0x32c9c468454fcf853e08d256000b08598539a799e849e0e7d8309c15eeb53ec3.
//
// Solidity: event ValidatorSetSaved(uint256 indexed period, bool pickedAll, uint256 nRV, address[] nonRotatingValidators, address[] rotatingValidators, uint256[] rotatingStakeAmounts)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseValidatorSetSaved(log types.Log) (*RandomBeaconCoordinatorValidatorSetSaved, error) {
	event := new(RandomBeaconCoordinatorValidatorSetSaved)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "ValidatorSetSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomBeaconCoordinatorValidatorThresholdUpdatedIterator is returned from FilterValidatorThresholdUpdated and is used to iterate over the raw logs and unpacked data for ValidatorThresholdUpdated events raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorValidatorThresholdUpdatedIterator struct {
	Event *RandomBeaconCoordinatorValidatorThresholdUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RandomBeaconCoordinatorValidatorThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomBeaconCoordinatorValidatorThresholdUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RandomBeaconCoordinatorValidatorThresholdUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RandomBeaconCoordinatorValidatorThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomBeaconCoordinatorValidatorThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomBeaconCoordinatorValidatorThresholdUpdated represents a ValidatorThresholdUpdated event raised by the RandomBeaconCoordinator contract.
type RandomBeaconCoordinatorValidatorThresholdUpdated struct {
	ValidatorType uint8
	Threshold     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorThresholdUpdated is a free log retrieval operation binding the contract event 0x57e533b2b2703988ec2b03256b82a8f952d21aecdcb8644c531a302021008906.
//
// Solidity: event ValidatorThresholdUpdated(uint8 indexed validatorType, uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) FilterValidatorThresholdUpdated(opts *bind.FilterOpts, validatorType []uint8) (*RandomBeaconCoordinatorValidatorThresholdUpdatedIterator, error) {

	var validatorTypeRule []interface{}
	for _, validatorTypeItem := range validatorType {
		validatorTypeRule = append(validatorTypeRule, validatorTypeItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.FilterLogs(opts, "ValidatorThresholdUpdated", validatorTypeRule)
	if err != nil {
		return nil, err
	}
	return &RandomBeaconCoordinatorValidatorThresholdUpdatedIterator{contract: _RandomBeaconCoordinator.contract, event: "ValidatorThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorThresholdUpdated is a free log subscription operation binding the contract event 0x57e533b2b2703988ec2b03256b82a8f952d21aecdcb8644c531a302021008906.
//
// Solidity: event ValidatorThresholdUpdated(uint8 indexed validatorType, uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) WatchValidatorThresholdUpdated(opts *bind.WatchOpts, sink chan<- *RandomBeaconCoordinatorValidatorThresholdUpdated, validatorType []uint8) (event.Subscription, error) {

	var validatorTypeRule []interface{}
	for _, validatorTypeItem := range validatorType {
		validatorTypeRule = append(validatorTypeRule, validatorTypeItem)
	}

	logs, sub, err := _RandomBeaconCoordinator.contract.WatchLogs(opts, "ValidatorThresholdUpdated", validatorTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomBeaconCoordinatorValidatorThresholdUpdated)
				if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "ValidatorThresholdUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorThresholdUpdated is a log parse operation binding the contract event 0x57e533b2b2703988ec2b03256b82a8f952d21aecdcb8644c531a302021008906.
//
// Solidity: event ValidatorThresholdUpdated(uint8 indexed validatorType, uint256 threshold)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorFilterer) ParseValidatorThresholdUpdated(log types.Log) (*RandomBeaconCoordinatorValidatorThresholdUpdated, error) {
	event := new(RandomBeaconCoordinatorValidatorThresholdUpdated)
	if err := _RandomBeaconCoordinator.contract.UnpackLog(event, "ValidatorThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
