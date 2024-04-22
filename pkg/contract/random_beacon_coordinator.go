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

// RandomRequest is an auto generated low-level Go binding around an user-defined struct.
type RandomRequest struct {
	Period     *big.Int
	PrevBeacon *big.Int
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
	ABI: "[{\"type\":\"function\",\"name\":\"bulkSetValidatorThresholds\",\"inputs\":[{\"name\":\"validatorTypes\",\"type\":\"uint8[]\",\"internalType\":\"enumIRandomBeacon.ValidatorType[]\"},{\"name\":\"thresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fulfillRandomSeed\",\"inputs\":[{\"name\":\"req\",\"type\":\"tuple\",\"internalType\":\"structRandomRequest\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevBeacon\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structVRF.Proof\",\"components\":[{\"name\":\"pk\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gamma\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"c\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uWitness\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"sHashWitness\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"zInv\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActivatedAtPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBeacon\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnavailabilityCount\",\"inputs\":[{\"name\":\"cid\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnavailabilitySlashThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorThreshold\",\"inputs\":[{\"name\":\"validatorType\",\"type\":\"uint8\",\"internalType\":\"enumIRandomBeacon.ValidatorType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSubmittedAt\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onWrapUpEpoch\",\"inputs\":[{\"name\":\"lastPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pickValidatorSet\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pickedCids\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUnavailabilitySlashThreshold\",\"inputs\":[{\"name\":\"slashThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BeaconFinalized\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomSeedFulfilled\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"reqHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomSeedRequested\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"reqHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"req\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRandomRequest\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevBeacon\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashUnavailabilityThresholdUpdated\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorThresholdUpdated\",\"inputs\":[{\"name\":\"validatorType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumIRandomBeacon.ValidatorType\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ErrAlreadySubmitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrBeaconAlreadyFinalized\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrBeaconNotFinalized\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidKeyHash\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidRandomRequest\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ErrKeyHashChangeCooldownNotEnded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrNotActivated\",\"inputs\":[{\"name\":\"untilPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrRegisterCoolDownNotEnded\",\"inputs\":[]}]",
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

// GetBeacon is a free data retrieval call binding the contract method 0x90caecc2.
//
// Solidity: function getBeacon(uint256 period) view returns(uint256 value, bool finalized)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetBeacon(opts *bind.CallOpts, period *big.Int) (struct {
	Value     *big.Int
	Finalized bool
}, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getBeacon", period)

	outstruct := new(struct {
		Value     *big.Int
		Finalized bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Finalized = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetBeacon is a free data retrieval call binding the contract method 0x90caecc2.
//
// Solidity: function getBeacon(uint256 period) view returns(uint256 value, bool finalized)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetBeacon(period *big.Int) (struct {
	Value     *big.Int
	Finalized bool
}, error) {
	return _RandomBeaconCoordinator.Contract.GetBeacon(&_RandomBeaconCoordinator.CallOpts, period)
}

// GetBeacon is a free data retrieval call binding the contract method 0x90caecc2.
//
// Solidity: function getBeacon(uint256 period) view returns(uint256 value, bool finalized)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetBeacon(period *big.Int) (struct {
	Value     *big.Int
	Finalized bool
}, error) {
	return _RandomBeaconCoordinator.Contract.GetBeacon(&_RandomBeaconCoordinator.CallOpts, period)
}

// GetUnavailabilityCount is a free data retrieval call binding the contract method 0x3a570332.
//
// Solidity: function getUnavailabilityCount(address cid) view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) GetUnavailabilityCount(opts *bind.CallOpts, cid common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "getUnavailabilityCount", cid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnavailabilityCount is a free data retrieval call binding the contract method 0x3a570332.
//
// Solidity: function getUnavailabilityCount(address cid) view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetUnavailabilityCount(cid common.Address) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilityCount(&_RandomBeaconCoordinator.CallOpts, cid)
}

// GetUnavailabilityCount is a free data retrieval call binding the contract method 0x3a570332.
//
// Solidity: function getUnavailabilityCount(address cid) view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetUnavailabilityCount(cid common.Address) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilityCount(&_RandomBeaconCoordinator.CallOpts, cid)
}

// GetUnavailabilitySlashThreshold is a free data retrieval call binding the contract method 0x917a65bd.
//
// Solidity: function getUnavailabilitySlashThreshold() view returns(uint256)
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
// Solidity: function getUnavailabilitySlashThreshold() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetUnavailabilitySlashThreshold() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilitySlashThreshold(&_RandomBeaconCoordinator.CallOpts)
}

// GetUnavailabilitySlashThreshold is a free data retrieval call binding the contract method 0x917a65bd.
//
// Solidity: function getUnavailabilitySlashThreshold() view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetUnavailabilitySlashThreshold() (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetUnavailabilitySlashThreshold(&_RandomBeaconCoordinator.CallOpts)
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0x035d0062.
//
// Solidity: function getValidatorThreshold(uint8 validatorType) view returns(uint256)
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
// Solidity: function getValidatorThreshold(uint8 validatorType) view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) GetValidatorThreshold(validatorType uint8) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetValidatorThreshold(&_RandomBeaconCoordinator.CallOpts, validatorType)
}

// GetValidatorThreshold is a free data retrieval call binding the contract method 0x035d0062.
//
// Solidity: function getValidatorThreshold(uint8 validatorType) view returns(uint256)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) GetValidatorThreshold(validatorType uint8) (*big.Int, error) {
	return _RandomBeaconCoordinator.Contract.GetValidatorThreshold(&_RandomBeaconCoordinator.CallOpts, validatorType)
}

// IsSubmittedAt is a free data retrieval call binding the contract method 0xe02f6271.
//
// Solidity: function isSubmittedAt(uint256 period, address oracle) view returns(bool)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) IsSubmittedAt(opts *bind.CallOpts, period *big.Int, oracle common.Address) (bool, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "isSubmittedAt", period, oracle)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubmittedAt is a free data retrieval call binding the contract method 0xe02f6271.
//
// Solidity: function isSubmittedAt(uint256 period, address oracle) view returns(bool)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) IsSubmittedAt(period *big.Int, oracle common.Address) (bool, error) {
	return _RandomBeaconCoordinator.Contract.IsSubmittedAt(&_RandomBeaconCoordinator.CallOpts, period, oracle)
}

// IsSubmittedAt is a free data retrieval call binding the contract method 0xe02f6271.
//
// Solidity: function isSubmittedAt(uint256 period, address oracle) view returns(bool)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) IsSubmittedAt(period *big.Int, oracle common.Address) (bool, error) {
	return _RandomBeaconCoordinator.Contract.IsSubmittedAt(&_RandomBeaconCoordinator.CallOpts, period, oracle)
}

// PickValidatorSet is a free data retrieval call binding the contract method 0xe6c933a4.
//
// Solidity: function pickValidatorSet(uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCaller) PickValidatorSet(opts *bind.CallOpts, epoch *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RandomBeaconCoordinator.contract.Call(opts, &out, "pickValidatorSet", epoch)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// PickValidatorSet is a free data retrieval call binding the contract method 0xe6c933a4.
//
// Solidity: function pickValidatorSet(uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) PickValidatorSet(epoch *big.Int) ([]common.Address, error) {
	return _RandomBeaconCoordinator.Contract.PickValidatorSet(&_RandomBeaconCoordinator.CallOpts, epoch)
}

// PickValidatorSet is a free data retrieval call binding the contract method 0xe6c933a4.
//
// Solidity: function pickValidatorSet(uint256 epoch) view returns(address[] pickedCids)
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorCallerSession) PickValidatorSet(epoch *big.Int) ([]common.Address, error) {
	return _RandomBeaconCoordinator.Contract.PickValidatorSet(&_RandomBeaconCoordinator.CallOpts, epoch)
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

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0x39ea0148.
//
// Solidity: function fulfillRandomSeed((uint256,uint256) req, (uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) FulfillRandomSeed(opts *bind.TransactOpts, req RandomRequest, proof VRFProof) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "fulfillRandomSeed", req, proof)
}

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0x39ea0148.
//
// Solidity: function fulfillRandomSeed((uint256,uint256) req, (uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) FulfillRandomSeed(req RandomRequest, proof VRFProof) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.FulfillRandomSeed(&_RandomBeaconCoordinator.TransactOpts, req, proof)
}

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0x39ea0148.
//
// Solidity: function fulfillRandomSeed((uint256,uint256) req, (uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) FulfillRandomSeed(req RandomRequest, proof VRFProof) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.FulfillRandomSeed(&_RandomBeaconCoordinator.TransactOpts, req, proof)
}

// OnWrapUpEpoch is a paid mutator transaction binding the contract method 0x8180c739.
//
// Solidity: function onWrapUpEpoch(uint256 lastPeriod, uint256 newPeriod) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) OnWrapUpEpoch(opts *bind.TransactOpts, lastPeriod *big.Int, newPeriod *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "onWrapUpEpoch", lastPeriod, newPeriod)
}

// OnWrapUpEpoch is a paid mutator transaction binding the contract method 0x8180c739.
//
// Solidity: function onWrapUpEpoch(uint256 lastPeriod, uint256 newPeriod) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) OnWrapUpEpoch(lastPeriod *big.Int, newPeriod *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.OnWrapUpEpoch(&_RandomBeaconCoordinator.TransactOpts, lastPeriod, newPeriod)
}

// OnWrapUpEpoch is a paid mutator transaction binding the contract method 0x8180c739.
//
// Solidity: function onWrapUpEpoch(uint256 lastPeriod, uint256 newPeriod) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) OnWrapUpEpoch(lastPeriod *big.Int, newPeriod *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.OnWrapUpEpoch(&_RandomBeaconCoordinator.TransactOpts, lastPeriod, newPeriod)
}

// SetUnavailabilitySlashThreshold is a paid mutator transaction binding the contract method 0x3355b0f0.
//
// Solidity: function setUnavailabilitySlashThreshold(uint256 slashThreshold) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactor) SetUnavailabilitySlashThreshold(opts *bind.TransactOpts, slashThreshold *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.contract.Transact(opts, "setUnavailabilitySlashThreshold", slashThreshold)
}

// SetUnavailabilitySlashThreshold is a paid mutator transaction binding the contract method 0x3355b0f0.
//
// Solidity: function setUnavailabilitySlashThreshold(uint256 slashThreshold) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorSession) SetUnavailabilitySlashThreshold(slashThreshold *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.SetUnavailabilitySlashThreshold(&_RandomBeaconCoordinator.TransactOpts, slashThreshold)
}

// SetUnavailabilitySlashThreshold is a paid mutator transaction binding the contract method 0x3355b0f0.
//
// Solidity: function setUnavailabilitySlashThreshold(uint256 slashThreshold) returns()
func (_RandomBeaconCoordinator *RandomBeaconCoordinatorTransactorSession) SetUnavailabilitySlashThreshold(slashThreshold *big.Int) (*types.Transaction, error) {
	return _RandomBeaconCoordinator.Contract.SetUnavailabilitySlashThreshold(&_RandomBeaconCoordinator.TransactOpts, slashThreshold)
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

// FilterRandomSeedRequested is a free log retrieval operation binding the contract event 0xd5609ba18c38af4207dc13061bcbfc3fe68cbe32e5e55d664b27097501bb5f7d.
//
// Solidity: event RandomSeedRequested(uint256 indexed period, bytes32 indexed reqHash, (uint256,uint256) req)
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

// WatchRandomSeedRequested is a free log subscription operation binding the contract event 0xd5609ba18c38af4207dc13061bcbfc3fe68cbe32e5e55d664b27097501bb5f7d.
//
// Solidity: event RandomSeedRequested(uint256 indexed period, bytes32 indexed reqHash, (uint256,uint256) req)
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

// ParseRandomSeedRequested is a log parse operation binding the contract event 0xd5609ba18c38af4207dc13061bcbfc3fe68cbe32e5e55d664b27097501bb5f7d.
//
// Solidity: event RandomSeedRequested(uint256 indexed period, bytes32 indexed reqHash, (uint256,uint256) req)
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
