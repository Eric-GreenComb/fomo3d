// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DiviesInterfaceABI is the input ABI used to generate the binding from.
const DiviesInterfaceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// DiviesInterfaceBin is the compiled bytecode used for deploying new contracts.
const DiviesInterfaceBin = `0x`

// DeployDiviesInterface deploys a new Ethereum contract, binding an instance of DiviesInterface to it.
func DeployDiviesInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiviesInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(DiviesInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DiviesInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiviesInterface{DiviesInterfaceCaller: DiviesInterfaceCaller{contract: contract}, DiviesInterfaceTransactor: DiviesInterfaceTransactor{contract: contract}, DiviesInterfaceFilterer: DiviesInterfaceFilterer{contract: contract}}, nil
}

// DiviesInterface is an auto generated Go binding around an Ethereum contract.
type DiviesInterface struct {
	DiviesInterfaceCaller     // Read-only binding to the contract
	DiviesInterfaceTransactor // Write-only binding to the contract
	DiviesInterfaceFilterer   // Log filterer for contract events
}

// DiviesInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiviesInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiviesInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiviesInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiviesInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiviesInterfaceSession struct {
	Contract     *DiviesInterface  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiviesInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiviesInterfaceCallerSession struct {
	Contract *DiviesInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DiviesInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiviesInterfaceTransactorSession struct {
	Contract     *DiviesInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DiviesInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiviesInterfaceRaw struct {
	Contract *DiviesInterface // Generic contract binding to access the raw methods on
}

// DiviesInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiviesInterfaceCallerRaw struct {
	Contract *DiviesInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DiviesInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiviesInterfaceTransactorRaw struct {
	Contract *DiviesInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiviesInterface creates a new instance of DiviesInterface, bound to a specific deployed contract.
func NewDiviesInterface(address common.Address, backend bind.ContractBackend) (*DiviesInterface, error) {
	contract, err := bindDiviesInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiviesInterface{DiviesInterfaceCaller: DiviesInterfaceCaller{contract: contract}, DiviesInterfaceTransactor: DiviesInterfaceTransactor{contract: contract}, DiviesInterfaceFilterer: DiviesInterfaceFilterer{contract: contract}}, nil
}

// NewDiviesInterfaceCaller creates a new read-only instance of DiviesInterface, bound to a specific deployed contract.
func NewDiviesInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DiviesInterfaceCaller, error) {
	contract, err := bindDiviesInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiviesInterfaceCaller{contract: contract}, nil
}

// NewDiviesInterfaceTransactor creates a new write-only instance of DiviesInterface, bound to a specific deployed contract.
func NewDiviesInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DiviesInterfaceTransactor, error) {
	contract, err := bindDiviesInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiviesInterfaceTransactor{contract: contract}, nil
}

// NewDiviesInterfaceFilterer creates a new log filterer instance of DiviesInterface, bound to a specific deployed contract.
func NewDiviesInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DiviesInterfaceFilterer, error) {
	contract, err := bindDiviesInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiviesInterfaceFilterer{contract: contract}, nil
}

// bindDiviesInterface binds a generic wrapper to an already deployed contract.
func bindDiviesInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiviesInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiviesInterface *DiviesInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DiviesInterface.Contract.DiviesInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiviesInterface *DiviesInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiviesInterface.Contract.DiviesInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiviesInterface *DiviesInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiviesInterface.Contract.DiviesInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiviesInterface *DiviesInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DiviesInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiviesInterface *DiviesInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiviesInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiviesInterface *DiviesInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiviesInterface.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_DiviesInterface *DiviesInterfaceTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiviesInterface.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_DiviesInterface *DiviesInterfaceSession) Deposit() (*types.Transaction, error) {
	return _DiviesInterface.Contract.Deposit(&_DiviesInterface.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_DiviesInterface *DiviesInterfaceTransactorSession) Deposit() (*types.Transaction, error) {
	return _DiviesInterface.Contract.Deposit(&_DiviesInterface.TransactOpts)
}

// F3DKeysCalcLongABI is the input ABI used to generate the binding from.
const F3DKeysCalcLongABI = "[]"

// F3DKeysCalcLongBin is the compiled bytecode used for deploying new contracts.
const F3DKeysCalcLongBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058209583c248e562bf792e36efd16ef207935b244274634b8df7e90f0f9e5984d1280029`

// DeployF3DKeysCalcLong deploys a new Ethereum contract, binding an instance of F3DKeysCalcLong to it.
func DeployF3DKeysCalcLong(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *F3DKeysCalcLong, error) {
	parsed, err := abi.JSON(strings.NewReader(F3DKeysCalcLongABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(F3DKeysCalcLongBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &F3DKeysCalcLong{F3DKeysCalcLongCaller: F3DKeysCalcLongCaller{contract: contract}, F3DKeysCalcLongTransactor: F3DKeysCalcLongTransactor{contract: contract}, F3DKeysCalcLongFilterer: F3DKeysCalcLongFilterer{contract: contract}}, nil
}

// F3DKeysCalcLong is an auto generated Go binding around an Ethereum contract.
type F3DKeysCalcLong struct {
	F3DKeysCalcLongCaller     // Read-only binding to the contract
	F3DKeysCalcLongTransactor // Write-only binding to the contract
	F3DKeysCalcLongFilterer   // Log filterer for contract events
}

// F3DKeysCalcLongCaller is an auto generated read-only Go binding around an Ethereum contract.
type F3DKeysCalcLongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DKeysCalcLongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type F3DKeysCalcLongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DKeysCalcLongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type F3DKeysCalcLongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DKeysCalcLongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type F3DKeysCalcLongSession struct {
	Contract     *F3DKeysCalcLong  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// F3DKeysCalcLongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type F3DKeysCalcLongCallerSession struct {
	Contract *F3DKeysCalcLongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// F3DKeysCalcLongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type F3DKeysCalcLongTransactorSession struct {
	Contract     *F3DKeysCalcLongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// F3DKeysCalcLongRaw is an auto generated low-level Go binding around an Ethereum contract.
type F3DKeysCalcLongRaw struct {
	Contract *F3DKeysCalcLong // Generic contract binding to access the raw methods on
}

// F3DKeysCalcLongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type F3DKeysCalcLongCallerRaw struct {
	Contract *F3DKeysCalcLongCaller // Generic read-only contract binding to access the raw methods on
}

// F3DKeysCalcLongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type F3DKeysCalcLongTransactorRaw struct {
	Contract *F3DKeysCalcLongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewF3DKeysCalcLong creates a new instance of F3DKeysCalcLong, bound to a specific deployed contract.
func NewF3DKeysCalcLong(address common.Address, backend bind.ContractBackend) (*F3DKeysCalcLong, error) {
	contract, err := bindF3DKeysCalcLong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &F3DKeysCalcLong{F3DKeysCalcLongCaller: F3DKeysCalcLongCaller{contract: contract}, F3DKeysCalcLongTransactor: F3DKeysCalcLongTransactor{contract: contract}, F3DKeysCalcLongFilterer: F3DKeysCalcLongFilterer{contract: contract}}, nil
}

// NewF3DKeysCalcLongCaller creates a new read-only instance of F3DKeysCalcLong, bound to a specific deployed contract.
func NewF3DKeysCalcLongCaller(address common.Address, caller bind.ContractCaller) (*F3DKeysCalcLongCaller, error) {
	contract, err := bindF3DKeysCalcLong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &F3DKeysCalcLongCaller{contract: contract}, nil
}

// NewF3DKeysCalcLongTransactor creates a new write-only instance of F3DKeysCalcLong, bound to a specific deployed contract.
func NewF3DKeysCalcLongTransactor(address common.Address, transactor bind.ContractTransactor) (*F3DKeysCalcLongTransactor, error) {
	contract, err := bindF3DKeysCalcLong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &F3DKeysCalcLongTransactor{contract: contract}, nil
}

// NewF3DKeysCalcLongFilterer creates a new log filterer instance of F3DKeysCalcLong, bound to a specific deployed contract.
func NewF3DKeysCalcLongFilterer(address common.Address, filterer bind.ContractFilterer) (*F3DKeysCalcLongFilterer, error) {
	contract, err := bindF3DKeysCalcLong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &F3DKeysCalcLongFilterer{contract: contract}, nil
}

// bindF3DKeysCalcLong binds a generic wrapper to an already deployed contract.
func bindF3DKeysCalcLong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(F3DKeysCalcLongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_F3DKeysCalcLong *F3DKeysCalcLongRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _F3DKeysCalcLong.Contract.F3DKeysCalcLongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_F3DKeysCalcLong *F3DKeysCalcLongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3DKeysCalcLong.Contract.F3DKeysCalcLongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_F3DKeysCalcLong *F3DKeysCalcLongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _F3DKeysCalcLong.Contract.F3DKeysCalcLongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_F3DKeysCalcLong *F3DKeysCalcLongCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _F3DKeysCalcLong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_F3DKeysCalcLong *F3DKeysCalcLongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3DKeysCalcLong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_F3DKeysCalcLong *F3DKeysCalcLongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _F3DKeysCalcLong.Contract.contract.Transact(opts, method, params...)
}

// F3DdatasetsABI is the input ABI used to generate the binding from.
const F3DdatasetsABI = "[]"

// F3DdatasetsBin is the compiled bytecode used for deploying new contracts.
const F3DdatasetsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c03ece80255992ad5687c4da22184702571b915791a0d24cce827249c1b410d30029`

// DeployF3Ddatasets deploys a new Ethereum contract, binding an instance of F3Ddatasets to it.
func DeployF3Ddatasets(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *F3Ddatasets, error) {
	parsed, err := abi.JSON(strings.NewReader(F3DdatasetsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(F3DdatasetsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &F3Ddatasets{F3DdatasetsCaller: F3DdatasetsCaller{contract: contract}, F3DdatasetsTransactor: F3DdatasetsTransactor{contract: contract}, F3DdatasetsFilterer: F3DdatasetsFilterer{contract: contract}}, nil
}

// F3Ddatasets is an auto generated Go binding around an Ethereum contract.
type F3Ddatasets struct {
	F3DdatasetsCaller     // Read-only binding to the contract
	F3DdatasetsTransactor // Write-only binding to the contract
	F3DdatasetsFilterer   // Log filterer for contract events
}

// F3DdatasetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type F3DdatasetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DdatasetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type F3DdatasetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DdatasetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type F3DdatasetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DdatasetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type F3DdatasetsSession struct {
	Contract     *F3Ddatasets      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// F3DdatasetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type F3DdatasetsCallerSession struct {
	Contract *F3DdatasetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// F3DdatasetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type F3DdatasetsTransactorSession struct {
	Contract     *F3DdatasetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// F3DdatasetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type F3DdatasetsRaw struct {
	Contract *F3Ddatasets // Generic contract binding to access the raw methods on
}

// F3DdatasetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type F3DdatasetsCallerRaw struct {
	Contract *F3DdatasetsCaller // Generic read-only contract binding to access the raw methods on
}

// F3DdatasetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type F3DdatasetsTransactorRaw struct {
	Contract *F3DdatasetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewF3Ddatasets creates a new instance of F3Ddatasets, bound to a specific deployed contract.
func NewF3Ddatasets(address common.Address, backend bind.ContractBackend) (*F3Ddatasets, error) {
	contract, err := bindF3Ddatasets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &F3Ddatasets{F3DdatasetsCaller: F3DdatasetsCaller{contract: contract}, F3DdatasetsTransactor: F3DdatasetsTransactor{contract: contract}, F3DdatasetsFilterer: F3DdatasetsFilterer{contract: contract}}, nil
}

// NewF3DdatasetsCaller creates a new read-only instance of F3Ddatasets, bound to a specific deployed contract.
func NewF3DdatasetsCaller(address common.Address, caller bind.ContractCaller) (*F3DdatasetsCaller, error) {
	contract, err := bindF3Ddatasets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &F3DdatasetsCaller{contract: contract}, nil
}

// NewF3DdatasetsTransactor creates a new write-only instance of F3Ddatasets, bound to a specific deployed contract.
func NewF3DdatasetsTransactor(address common.Address, transactor bind.ContractTransactor) (*F3DdatasetsTransactor, error) {
	contract, err := bindF3Ddatasets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &F3DdatasetsTransactor{contract: contract}, nil
}

// NewF3DdatasetsFilterer creates a new log filterer instance of F3Ddatasets, bound to a specific deployed contract.
func NewF3DdatasetsFilterer(address common.Address, filterer bind.ContractFilterer) (*F3DdatasetsFilterer, error) {
	contract, err := bindF3Ddatasets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &F3DdatasetsFilterer{contract: contract}, nil
}

// bindF3Ddatasets binds a generic wrapper to an already deployed contract.
func bindF3Ddatasets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(F3DdatasetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_F3Ddatasets *F3DdatasetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _F3Ddatasets.Contract.F3DdatasetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_F3Ddatasets *F3DdatasetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3Ddatasets.Contract.F3DdatasetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_F3Ddatasets *F3DdatasetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _F3Ddatasets.Contract.F3DdatasetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_F3Ddatasets *F3DdatasetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _F3Ddatasets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_F3Ddatasets *F3DdatasetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3Ddatasets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_F3Ddatasets *F3DdatasetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _F3Ddatasets.Contract.contract.Transact(opts, method, params...)
}

// F3DeventsABI is the input ABI used to generate the binding from.
const F3DeventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isNewPlayer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onNewName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"keysBought\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"potAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"airDropPot\",\"type\":\"uint256\"}],\"name\":\"onEndTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onWithdrawAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onReLoadAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"buyerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onAffiliatePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountAddedToPot\",\"type\":\"uint256\"}],\"name\":\"onPotSwapDeposit\",\"type\":\"event\"}]"

// F3DeventsBin is the compiled bytecode used for deploying new contracts.
const F3DeventsBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a7230582057eca8d45b01f45cb9d78c6d6702e0dac7cd312a3fad38faa04efc547dd189810029`

// DeployF3Devents deploys a new Ethereum contract, binding an instance of F3Devents to it.
func DeployF3Devents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *F3Devents, error) {
	parsed, err := abi.JSON(strings.NewReader(F3DeventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(F3DeventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &F3Devents{F3DeventsCaller: F3DeventsCaller{contract: contract}, F3DeventsTransactor: F3DeventsTransactor{contract: contract}, F3DeventsFilterer: F3DeventsFilterer{contract: contract}}, nil
}

// F3Devents is an auto generated Go binding around an Ethereum contract.
type F3Devents struct {
	F3DeventsCaller     // Read-only binding to the contract
	F3DeventsTransactor // Write-only binding to the contract
	F3DeventsFilterer   // Log filterer for contract events
}

// F3DeventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type F3DeventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DeventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type F3DeventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DeventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type F3DeventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DeventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type F3DeventsSession struct {
	Contract     *F3Devents        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// F3DeventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type F3DeventsCallerSession struct {
	Contract *F3DeventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// F3DeventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type F3DeventsTransactorSession struct {
	Contract     *F3DeventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// F3DeventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type F3DeventsRaw struct {
	Contract *F3Devents // Generic contract binding to access the raw methods on
}

// F3DeventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type F3DeventsCallerRaw struct {
	Contract *F3DeventsCaller // Generic read-only contract binding to access the raw methods on
}

// F3DeventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type F3DeventsTransactorRaw struct {
	Contract *F3DeventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewF3Devents creates a new instance of F3Devents, bound to a specific deployed contract.
func NewF3Devents(address common.Address, backend bind.ContractBackend) (*F3Devents, error) {
	contract, err := bindF3Devents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &F3Devents{F3DeventsCaller: F3DeventsCaller{contract: contract}, F3DeventsTransactor: F3DeventsTransactor{contract: contract}, F3DeventsFilterer: F3DeventsFilterer{contract: contract}}, nil
}

// NewF3DeventsCaller creates a new read-only instance of F3Devents, bound to a specific deployed contract.
func NewF3DeventsCaller(address common.Address, caller bind.ContractCaller) (*F3DeventsCaller, error) {
	contract, err := bindF3Devents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &F3DeventsCaller{contract: contract}, nil
}

// NewF3DeventsTransactor creates a new write-only instance of F3Devents, bound to a specific deployed contract.
func NewF3DeventsTransactor(address common.Address, transactor bind.ContractTransactor) (*F3DeventsTransactor, error) {
	contract, err := bindF3Devents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &F3DeventsTransactor{contract: contract}, nil
}

// NewF3DeventsFilterer creates a new log filterer instance of F3Devents, bound to a specific deployed contract.
func NewF3DeventsFilterer(address common.Address, filterer bind.ContractFilterer) (*F3DeventsFilterer, error) {
	contract, err := bindF3Devents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &F3DeventsFilterer{contract: contract}, nil
}

// bindF3Devents binds a generic wrapper to an already deployed contract.
func bindF3Devents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(F3DeventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_F3Devents *F3DeventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _F3Devents.Contract.F3DeventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_F3Devents *F3DeventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3Devents.Contract.F3DeventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_F3Devents *F3DeventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _F3Devents.Contract.F3DeventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_F3Devents *F3DeventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _F3Devents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_F3Devents *F3DeventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3Devents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_F3Devents *F3DeventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _F3Devents.Contract.contract.Transact(opts, method, params...)
}

// F3DeventsOnAffiliatePayoutIterator is returned from FilterOnAffiliatePayout and is used to iterate over the raw logs and unpacked data for OnAffiliatePayout events raised by the F3Devents contract.
type F3DeventsOnAffiliatePayoutIterator struct {
	Event *F3DeventsOnAffiliatePayout // Event containing the contract specifics and raw log

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
func (it *F3DeventsOnAffiliatePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(F3DeventsOnAffiliatePayout)
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
		it.Event = new(F3DeventsOnAffiliatePayout)
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
func (it *F3DeventsOnAffiliatePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *F3DeventsOnAffiliatePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// F3DeventsOnAffiliatePayout represents a OnAffiliatePayout event raised by the F3Devents contract.
type F3DeventsOnAffiliatePayout struct {
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	RoundID          *big.Int
	BuyerID          *big.Int
	Amount           *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnAffiliatePayout is a free log retrieval operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: e onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_F3Devents *F3DeventsFilterer) FilterOnAffiliatePayout(opts *bind.FilterOpts, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (*F3DeventsOnAffiliatePayoutIterator, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _F3Devents.contract.FilterLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return &F3DeventsOnAffiliatePayoutIterator{contract: _F3Devents.contract, event: "onAffiliatePayout", logs: logs, sub: sub}, nil
}

// WatchOnAffiliatePayout is a free log subscription operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: e onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_F3Devents *F3DeventsFilterer) WatchOnAffiliatePayout(opts *bind.WatchOpts, sink chan<- *F3DeventsOnAffiliatePayout, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (event.Subscription, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _F3Devents.contract.WatchLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(F3DeventsOnAffiliatePayout)
				if err := _F3Devents.contract.UnpackLog(event, "onAffiliatePayout", log); err != nil {
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

// F3DeventsOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the F3Devents contract.
type F3DeventsOnBuyAndDistributeIterator struct {
	Event *F3DeventsOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *F3DeventsOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(F3DeventsOnBuyAndDistribute)
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
		it.Event = new(F3DeventsOnBuyAndDistribute)
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
func (it *F3DeventsOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *F3DeventsOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// F3DeventsOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the F3Devents contract.
type F3DeventsOnBuyAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EthIn          *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: e onBuyAndDistribute(playerAddress address, playerName bytes32, ethIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_F3Devents *F3DeventsFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*F3DeventsOnBuyAndDistributeIterator, error) {

	logs, sub, err := _F3Devents.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &F3DeventsOnBuyAndDistributeIterator{contract: _F3Devents.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: e onBuyAndDistribute(playerAddress address, playerName bytes32, ethIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_F3Devents *F3DeventsFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *F3DeventsOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _F3Devents.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(F3DeventsOnBuyAndDistribute)
				if err := _F3Devents.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// F3DeventsOnEndTxIterator is returned from FilterOnEndTx and is used to iterate over the raw logs and unpacked data for OnEndTx events raised by the F3Devents contract.
type F3DeventsOnEndTxIterator struct {
	Event *F3DeventsOnEndTx // Event containing the contract specifics and raw log

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
func (it *F3DeventsOnEndTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(F3DeventsOnEndTx)
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
		it.Event = new(F3DeventsOnEndTx)
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
func (it *F3DeventsOnEndTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *F3DeventsOnEndTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// F3DeventsOnEndTx represents a OnEndTx event raised by the F3Devents contract.
type F3DeventsOnEndTx struct {
	CompressedData *big.Int
	CompressedIDs  *big.Int
	PlayerName     [32]byte
	PlayerAddress  common.Address
	EthIn          *big.Int
	KeysBought     *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	PotAmount      *big.Int
	AirDropPot     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnEndTx is a free log retrieval operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: e onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, ethIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_F3Devents *F3DeventsFilterer) FilterOnEndTx(opts *bind.FilterOpts) (*F3DeventsOnEndTxIterator, error) {

	logs, sub, err := _F3Devents.contract.FilterLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return &F3DeventsOnEndTxIterator{contract: _F3Devents.contract, event: "onEndTx", logs: logs, sub: sub}, nil
}

// WatchOnEndTx is a free log subscription operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: e onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, ethIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_F3Devents *F3DeventsFilterer) WatchOnEndTx(opts *bind.WatchOpts, sink chan<- *F3DeventsOnEndTx) (event.Subscription, error) {

	logs, sub, err := _F3Devents.contract.WatchLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(F3DeventsOnEndTx)
				if err := _F3Devents.contract.UnpackLog(event, "onEndTx", log); err != nil {
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

// F3DeventsOnNewNameIterator is returned from FilterOnNewName and is used to iterate over the raw logs and unpacked data for OnNewName events raised by the F3Devents contract.
type F3DeventsOnNewNameIterator struct {
	Event *F3DeventsOnNewName // Event containing the contract specifics and raw log

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
func (it *F3DeventsOnNewNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(F3DeventsOnNewName)
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
		it.Event = new(F3DeventsOnNewName)
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
func (it *F3DeventsOnNewNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *F3DeventsOnNewNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// F3DeventsOnNewName represents a OnNewName event raised by the F3Devents contract.
type F3DeventsOnNewName struct {
	PlayerID         *big.Int
	PlayerAddress    common.Address
	PlayerName       [32]byte
	IsNewPlayer      bool
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	AmountPaid       *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnNewName is a free log retrieval operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: e onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_F3Devents *F3DeventsFilterer) FilterOnNewName(opts *bind.FilterOpts, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (*F3DeventsOnNewNameIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _F3Devents.contract.FilterLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return &F3DeventsOnNewNameIterator{contract: _F3Devents.contract, event: "onNewName", logs: logs, sub: sub}, nil
}

// WatchOnNewName is a free log subscription operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: e onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_F3Devents *F3DeventsFilterer) WatchOnNewName(opts *bind.WatchOpts, sink chan<- *F3DeventsOnNewName, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _F3Devents.contract.WatchLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(F3DeventsOnNewName)
				if err := _F3Devents.contract.UnpackLog(event, "onNewName", log); err != nil {
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

// F3DeventsOnPotSwapDepositIterator is returned from FilterOnPotSwapDeposit and is used to iterate over the raw logs and unpacked data for OnPotSwapDeposit events raised by the F3Devents contract.
type F3DeventsOnPotSwapDepositIterator struct {
	Event *F3DeventsOnPotSwapDeposit // Event containing the contract specifics and raw log

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
func (it *F3DeventsOnPotSwapDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(F3DeventsOnPotSwapDeposit)
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
		it.Event = new(F3DeventsOnPotSwapDeposit)
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
func (it *F3DeventsOnPotSwapDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *F3DeventsOnPotSwapDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// F3DeventsOnPotSwapDeposit represents a OnPotSwapDeposit event raised by the F3Devents contract.
type F3DeventsOnPotSwapDeposit struct {
	RoundID          *big.Int
	AmountAddedToPot *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnPotSwapDeposit is a free log retrieval operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: e onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_F3Devents *F3DeventsFilterer) FilterOnPotSwapDeposit(opts *bind.FilterOpts) (*F3DeventsOnPotSwapDepositIterator, error) {

	logs, sub, err := _F3Devents.contract.FilterLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return &F3DeventsOnPotSwapDepositIterator{contract: _F3Devents.contract, event: "onPotSwapDeposit", logs: logs, sub: sub}, nil
}

// WatchOnPotSwapDeposit is a free log subscription operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: e onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_F3Devents *F3DeventsFilterer) WatchOnPotSwapDeposit(opts *bind.WatchOpts, sink chan<- *F3DeventsOnPotSwapDeposit) (event.Subscription, error) {

	logs, sub, err := _F3Devents.contract.WatchLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(F3DeventsOnPotSwapDeposit)
				if err := _F3Devents.contract.UnpackLog(event, "onPotSwapDeposit", log); err != nil {
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

// F3DeventsOnReLoadAndDistributeIterator is returned from FilterOnReLoadAndDistribute and is used to iterate over the raw logs and unpacked data for OnReLoadAndDistribute events raised by the F3Devents contract.
type F3DeventsOnReLoadAndDistributeIterator struct {
	Event *F3DeventsOnReLoadAndDistribute // Event containing the contract specifics and raw log

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
func (it *F3DeventsOnReLoadAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(F3DeventsOnReLoadAndDistribute)
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
		it.Event = new(F3DeventsOnReLoadAndDistribute)
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
func (it *F3DeventsOnReLoadAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *F3DeventsOnReLoadAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// F3DeventsOnReLoadAndDistribute represents a OnReLoadAndDistribute event raised by the F3Devents contract.
type F3DeventsOnReLoadAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnReLoadAndDistribute is a free log retrieval operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: e onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_F3Devents *F3DeventsFilterer) FilterOnReLoadAndDistribute(opts *bind.FilterOpts) (*F3DeventsOnReLoadAndDistributeIterator, error) {

	logs, sub, err := _F3Devents.contract.FilterLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return &F3DeventsOnReLoadAndDistributeIterator{contract: _F3Devents.contract, event: "onReLoadAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnReLoadAndDistribute is a free log subscription operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: e onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_F3Devents *F3DeventsFilterer) WatchOnReLoadAndDistribute(opts *bind.WatchOpts, sink chan<- *F3DeventsOnReLoadAndDistribute) (event.Subscription, error) {

	logs, sub, err := _F3Devents.contract.WatchLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(F3DeventsOnReLoadAndDistribute)
				if err := _F3Devents.contract.UnpackLog(event, "onReLoadAndDistribute", log); err != nil {
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

// F3DeventsOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the F3Devents contract.
type F3DeventsOnWithdrawIterator struct {
	Event *F3DeventsOnWithdraw // Event containing the contract specifics and raw log

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
func (it *F3DeventsOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(F3DeventsOnWithdraw)
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
		it.Event = new(F3DeventsOnWithdraw)
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
func (it *F3DeventsOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *F3DeventsOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// F3DeventsOnWithdraw represents a OnWithdraw event raised by the F3Devents contract.
type F3DeventsOnWithdraw struct {
	PlayerID      *big.Int
	PlayerAddress common.Address
	PlayerName    [32]byte
	EthOut        *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: e onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, ethOut uint256, timeStamp uint256)
func (_F3Devents *F3DeventsFilterer) FilterOnWithdraw(opts *bind.FilterOpts, playerID []*big.Int) (*F3DeventsOnWithdrawIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _F3Devents.contract.FilterLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return &F3DeventsOnWithdrawIterator{contract: _F3Devents.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: e onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, ethOut uint256, timeStamp uint256)
func (_F3Devents *F3DeventsFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *F3DeventsOnWithdraw, playerID []*big.Int) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _F3Devents.contract.WatchLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(F3DeventsOnWithdraw)
				if err := _F3Devents.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// F3DeventsOnWithdrawAndDistributeIterator is returned from FilterOnWithdrawAndDistribute and is used to iterate over the raw logs and unpacked data for OnWithdrawAndDistribute events raised by the F3Devents contract.
type F3DeventsOnWithdrawAndDistributeIterator struct {
	Event *F3DeventsOnWithdrawAndDistribute // Event containing the contract specifics and raw log

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
func (it *F3DeventsOnWithdrawAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(F3DeventsOnWithdrawAndDistribute)
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
		it.Event = new(F3DeventsOnWithdrawAndDistribute)
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
func (it *F3DeventsOnWithdrawAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *F3DeventsOnWithdrawAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// F3DeventsOnWithdrawAndDistribute represents a OnWithdrawAndDistribute event raised by the F3Devents contract.
type F3DeventsOnWithdrawAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EthOut         *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnWithdrawAndDistribute is a free log retrieval operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: e onWithdrawAndDistribute(playerAddress address, playerName bytes32, ethOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_F3Devents *F3DeventsFilterer) FilterOnWithdrawAndDistribute(opts *bind.FilterOpts) (*F3DeventsOnWithdrawAndDistributeIterator, error) {

	logs, sub, err := _F3Devents.contract.FilterLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return &F3DeventsOnWithdrawAndDistributeIterator{contract: _F3Devents.contract, event: "onWithdrawAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnWithdrawAndDistribute is a free log subscription operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: e onWithdrawAndDistribute(playerAddress address, playerName bytes32, ethOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_F3Devents *F3DeventsFilterer) WatchOnWithdrawAndDistribute(opts *bind.WatchOpts, sink chan<- *F3DeventsOnWithdrawAndDistribute) (event.Subscription, error) {

	logs, sub, err := _F3Devents.contract.WatchLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(F3DeventsOnWithdrawAndDistribute)
				if err := _F3Devents.contract.UnpackLog(event, "onWithdrawAndDistribute", log); err != nil {
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

// F3DexternalSettingsInterfaceABI is the input ABI used to generate the binding from.
const F3DexternalSettingsInterfaceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getLongGap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getFastGap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getFastExtra\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getLongExtra\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// F3DexternalSettingsInterfaceBin is the compiled bytecode used for deploying new contracts.
const F3DexternalSettingsInterfaceBin = `0x`

// DeployF3DexternalSettingsInterface deploys a new Ethereum contract, binding an instance of F3DexternalSettingsInterface to it.
func DeployF3DexternalSettingsInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *F3DexternalSettingsInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(F3DexternalSettingsInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(F3DexternalSettingsInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &F3DexternalSettingsInterface{F3DexternalSettingsInterfaceCaller: F3DexternalSettingsInterfaceCaller{contract: contract}, F3DexternalSettingsInterfaceTransactor: F3DexternalSettingsInterfaceTransactor{contract: contract}, F3DexternalSettingsInterfaceFilterer: F3DexternalSettingsInterfaceFilterer{contract: contract}}, nil
}

// F3DexternalSettingsInterface is an auto generated Go binding around an Ethereum contract.
type F3DexternalSettingsInterface struct {
	F3DexternalSettingsInterfaceCaller     // Read-only binding to the contract
	F3DexternalSettingsInterfaceTransactor // Write-only binding to the contract
	F3DexternalSettingsInterfaceFilterer   // Log filterer for contract events
}

// F3DexternalSettingsInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type F3DexternalSettingsInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DexternalSettingsInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type F3DexternalSettingsInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DexternalSettingsInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type F3DexternalSettingsInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// F3DexternalSettingsInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type F3DexternalSettingsInterfaceSession struct {
	Contract     *F3DexternalSettingsInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// F3DexternalSettingsInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type F3DexternalSettingsInterfaceCallerSession struct {
	Contract *F3DexternalSettingsInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// F3DexternalSettingsInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type F3DexternalSettingsInterfaceTransactorSession struct {
	Contract     *F3DexternalSettingsInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// F3DexternalSettingsInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type F3DexternalSettingsInterfaceRaw struct {
	Contract *F3DexternalSettingsInterface // Generic contract binding to access the raw methods on
}

// F3DexternalSettingsInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type F3DexternalSettingsInterfaceCallerRaw struct {
	Contract *F3DexternalSettingsInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// F3DexternalSettingsInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type F3DexternalSettingsInterfaceTransactorRaw struct {
	Contract *F3DexternalSettingsInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewF3DexternalSettingsInterface creates a new instance of F3DexternalSettingsInterface, bound to a specific deployed contract.
func NewF3DexternalSettingsInterface(address common.Address, backend bind.ContractBackend) (*F3DexternalSettingsInterface, error) {
	contract, err := bindF3DexternalSettingsInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &F3DexternalSettingsInterface{F3DexternalSettingsInterfaceCaller: F3DexternalSettingsInterfaceCaller{contract: contract}, F3DexternalSettingsInterfaceTransactor: F3DexternalSettingsInterfaceTransactor{contract: contract}, F3DexternalSettingsInterfaceFilterer: F3DexternalSettingsInterfaceFilterer{contract: contract}}, nil
}

// NewF3DexternalSettingsInterfaceCaller creates a new read-only instance of F3DexternalSettingsInterface, bound to a specific deployed contract.
func NewF3DexternalSettingsInterfaceCaller(address common.Address, caller bind.ContractCaller) (*F3DexternalSettingsInterfaceCaller, error) {
	contract, err := bindF3DexternalSettingsInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &F3DexternalSettingsInterfaceCaller{contract: contract}, nil
}

// NewF3DexternalSettingsInterfaceTransactor creates a new write-only instance of F3DexternalSettingsInterface, bound to a specific deployed contract.
func NewF3DexternalSettingsInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*F3DexternalSettingsInterfaceTransactor, error) {
	contract, err := bindF3DexternalSettingsInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &F3DexternalSettingsInterfaceTransactor{contract: contract}, nil
}

// NewF3DexternalSettingsInterfaceFilterer creates a new log filterer instance of F3DexternalSettingsInterface, bound to a specific deployed contract.
func NewF3DexternalSettingsInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*F3DexternalSettingsInterfaceFilterer, error) {
	contract, err := bindF3DexternalSettingsInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &F3DexternalSettingsInterfaceFilterer{contract: contract}, nil
}

// bindF3DexternalSettingsInterface binds a generic wrapper to an already deployed contract.
func bindF3DexternalSettingsInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(F3DexternalSettingsInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _F3DexternalSettingsInterface.Contract.F3DexternalSettingsInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.F3DexternalSettingsInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.F3DexternalSettingsInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _F3DexternalSettingsInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.contract.Transact(opts, method, params...)
}

// GetFastExtra is a paid mutator transaction binding the contract method 0x1a9be331.
//
// Solidity: function getFastExtra() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactor) GetFastExtra(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.contract.Transact(opts, "getFastExtra")
}

// GetFastExtra is a paid mutator transaction binding the contract method 0x1a9be331.
//
// Solidity: function getFastExtra() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceSession) GetFastExtra() (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.GetFastExtra(&_F3DexternalSettingsInterface.TransactOpts)
}

// GetFastExtra is a paid mutator transaction binding the contract method 0x1a9be331.
//
// Solidity: function getFastExtra() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactorSession) GetFastExtra() (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.GetFastExtra(&_F3DexternalSettingsInterface.TransactOpts)
}

// GetFastGap is a paid mutator transaction binding the contract method 0x18d0376c.
//
// Solidity: function getFastGap() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactor) GetFastGap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.contract.Transact(opts, "getFastGap")
}

// GetFastGap is a paid mutator transaction binding the contract method 0x18d0376c.
//
// Solidity: function getFastGap() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceSession) GetFastGap() (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.GetFastGap(&_F3DexternalSettingsInterface.TransactOpts)
}

// GetFastGap is a paid mutator transaction binding the contract method 0x18d0376c.
//
// Solidity: function getFastGap() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactorSession) GetFastGap() (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.GetFastGap(&_F3DexternalSettingsInterface.TransactOpts)
}

// GetLongExtra is a paid mutator transaction binding the contract method 0x4ccbe888.
//
// Solidity: function getLongExtra() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactor) GetLongExtra(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.contract.Transact(opts, "getLongExtra")
}

// GetLongExtra is a paid mutator transaction binding the contract method 0x4ccbe888.
//
// Solidity: function getLongExtra() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceSession) GetLongExtra() (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.GetLongExtra(&_F3DexternalSettingsInterface.TransactOpts)
}

// GetLongExtra is a paid mutator transaction binding the contract method 0x4ccbe888.
//
// Solidity: function getLongExtra() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactorSession) GetLongExtra() (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.GetLongExtra(&_F3DexternalSettingsInterface.TransactOpts)
}

// GetLongGap is a paid mutator transaction binding the contract method 0x114719c5.
//
// Solidity: function getLongGap() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactor) GetLongGap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.contract.Transact(opts, "getLongGap")
}

// GetLongGap is a paid mutator transaction binding the contract method 0x114719c5.
//
// Solidity: function getLongGap() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceSession) GetLongGap() (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.GetLongGap(&_F3DexternalSettingsInterface.TransactOpts)
}

// GetLongGap is a paid mutator transaction binding the contract method 0x114719c5.
//
// Solidity: function getLongGap() returns(uint256)
func (_F3DexternalSettingsInterface *F3DexternalSettingsInterfaceTransactorSession) GetLongGap() (*types.Transaction, error) {
	return _F3DexternalSettingsInterface.Contract.GetLongGap(&_F3DexternalSettingsInterface.TransactOpts)
}

// FoMo3DlongABI is the input ABI used to generate the binding from.
const FoMo3DlongABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBuyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_team\",\"type\":\"uint256\"},{\"name\":\"_eth\",\"type\":\"uint256\"}],\"name\":\"reLoadXname\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"pIDxAddr_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"airDropTracker_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round_\",\"outputs\":[{\"name\":\"plyr\",\"type\":\"uint256\"},{\"name\":\"team\",\"type\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\"},{\"name\":\"strt\",\"type\":\"uint256\"},{\"name\":\"keys\",\"type\":\"uint256\"},{\"name\":\"eth\",\"type\":\"uint256\"},{\"name\":\"pot\",\"type\":\"uint256\"},{\"name\":\"mask\",\"type\":\"uint256\"},{\"name\":\"ico\",\"type\":\"uint256\"},{\"name\":\"icoGen\",\"type\":\"uint256\"},{\"name\":\"icoAvg\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"plyrNames_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fees_\",\"outputs\":[{\"name\":\"gen\",\"type\":\"uint256\"},{\"name\":\"p3d\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pIDxName_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_team\",\"type\":\"uint256\"},{\"name\":\"_eth\",\"type\":\"uint256\"}],\"name\":\"reLoadXid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXaddr\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"},{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_laff\",\"type\":\"uint256\"}],\"name\":\"receivePlayerInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rndTmEth_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rID_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerVaults\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXname\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentRoundInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_team\",\"type\":\"uint256\"},{\"name\":\"_eth\",\"type\":\"uint256\"}],\"name\":\"reLoadXaddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_team\",\"type\":\"uint256\"}],\"name\":\"buyXid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"receivePlayerNameList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nameString\",\"type\":\"string\"},{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXID\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_team\",\"type\":\"uint256\"}],\"name\":\"buyXaddr\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plyrRnds_\",\"outputs\":[{\"name\":\"eth\",\"type\":\"uint256\"},{\"name\":\"keys\",\"type\":\"uint256\"},{\"name\":\"mask\",\"type\":\"uint256\"},{\"name\":\"ico\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_team\",\"type\":\"uint256\"}],\"name\":\"buyXname\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_otherF3D\",\"type\":\"address\"}],\"name\":\"setOtherFomo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"potSplit_\",\"outputs\":[{\"name\":\"gen\",\"type\":\"uint256\"},{\"name\":\"p3d\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTimeLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_rID\",\"type\":\"uint256\"},{\"name\":\"_eth\",\"type\":\"uint256\"}],\"name\":\"calcKeysReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_keys\",\"type\":\"uint256\"}],\"name\":\"iWantXKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activated_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"airDropPot_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plyr_\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"win\",\"type\":\"uint256\"},{\"name\":\"gen\",\"type\":\"uint256\"},{\"name\":\"aff\",\"type\":\"uint256\"},{\"name\":\"lrnd\",\"type\":\"uint256\"},{\"name\":\"laff\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"potSwap\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPlayerInfoByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isNewPlayer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onNewName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"keysBought\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"potAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"airDropPot\",\"type\":\"uint256\"}],\"name\":\"onEndTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onWithdrawAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onReLoadAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"buyerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onAffiliatePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountAddedToPot\",\"type\":\"uint256\"}],\"name\":\"onPotSwapDeposit\",\"type\":\"event\"}]"

// FoMo3DlongBin is the compiled bytecode used for deploying new contracts.
const FoMo3DlongBin = `0x6080604052601e6001556301e133806002556000600455600f805460ff1916905534801561002c57600080fd5b50604080518082018252601e808252600660208084019182526000808052600d80835294517f81955a0a11e65eac625c29e8882660bae4e165a75d72780094acae8ece9a29ee5591517f81955a0a11e65eac625c29e8882660bae4e165a75d72780094acae8ece9a29ef5584518086018652602b808252818301848152600180865287855292517ffd54ff1ed53f34a900b24c5ba64f85761163b5d82d98a47b9bd80e45466993c555517ffd54ff1ed53f34a900b24c5ba64f85761163b5d82d98a47b9bd80e45466993c6558651808801885260388152600a818501818152600280885289875292517f10a81eed9d63d16face5e76357905348e6253d3394086026bb2bf2145d7cc24955517f10a81eed9d63d16face5e76357905348e6253d3394086026bb2bf2145d7cc24a558851808a018a529283526008838601908152600380885298865292517f26b4a10d0f0b04925c23bd4480ee147c916e5e87a7d68206a533dad160ac81e25591517f26b4a10d0f0b04925c23bd4480ee147c916e5e87a7d68206a533dad160ac81e35587518089018952600f8152808501838152868052600e80875291517fe710864318d4a32f37d6ce54cb3fadbef648dd12d8dbdf53973564d56b7f881c55517fe710864318d4a32f37d6ce54cb3fadbef648dd12d8dbdf53973564d56b7f881d558851808a018a5260198152808601878152948752818652517fa7c5ba7114a813b50159add3a36832908dc83db71d0b9a24c2ad0f83be9582075592517fa7c5ba7114a813b50159add3a36832908dc83db71d0b9a24c2ad0f83be95820855875180890189526014808252818601908152918652838552517f9adb202b1492743bc00c81d33cdc6423fa8c79109027eb6a845391e8fc1f048155517f9adb202b1492743bc00c81d33cdc6423fa8c79109027eb6a845391e8fc1f0482558651808801909752938652858201938452939091529190915290517fe0283e559c29e31ee7f56467acc9dd307779c843a883aeeb3bf5c6128c90814455517fe0283e559c29e31ee7f56467acc9dd307779c843a883aeeb3bf5c6128c908145556155d380620003486000396000f3006080604052600436106101b35763ffffffff60e060020a600035041663018a25e8811461035157806306fdde0314610378578063079ce327146104025780630f15f4c01461042257806310f01eba1461043757806311a09ae71461045857806324c33d331461046d5780632660316e146104e45780632ce21999146105135780632e19ebdc14610544578063349cdcac1461055c5780633ccfd60b1461057a5780633ddd46981461058f57806349cc635d146105eb5780635893d48114610615578063624ae5c0146106305780636306643414610645578063685ffd831461067b578063747dff42146106ce57806382bfc739146107595780638f38f309146107805780638f7140ea1461078e578063921dec21146107a957806395d89b41146107fc57806398a0871d14610811578063a2bccae914610828578063a65b37a114610869578063b483c05414610877578063c519500e14610898578063c7e284b8146108b0578063ce89c80c146108c5578063cf808000146108e0578063d53b2679146108f8578063d87574e01461090d578063de7874f314610922578063ed78cf4a1461097c578063ee0b5d8b14610984575b6101bb6154ae565b600f5460009060ff16151560011461021f576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155488339815191526044820152600080516020615508833981519152606482015290519081900360840190fd5b33803b8015610266576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b34633b9aca008110156102be576040805160e560020a62461bcd0281526020600482015260216024820152600080516020615528833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af680000081111561030e576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020615568833981519152604482015290519081900360640190fd5b610317856109dd565b3360009081526006602081815260408084205480855260089092529092200154919650945061034a908590600288610c91565b5050505050005b34801561035d57600080fd5b50610366610f22565b60408051918252519081900360200190f35b34801561038457600080fd5b5061038d610fe7565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103c75781810151838201526020016103af565b50505050905090810190601f1680156103f45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040e57600080fd5b5061042060043560243560443561101e565b005b34801561042e57600080fd5b5061042061122a565b34801561044357600080fd5b50610366600160a060020a03600435166113e8565b34801561046457600080fd5b506103666113fa565b34801561047957600080fd5b50610485600435611400565b604080519c8d5260208d019b909b528b8b019990995296151560608b015260808a019590955260a089019390935260c088019190915260e087015261010086015261012085015261014084015261016083015251908190036101800190f35b3480156104f057600080fd5b506104ff600435602435611463565b604080519115158252519081900360200190f35b34801561051f57600080fd5b5061052b600435611483565b6040805192835260208301919091528051918290030190f35b34801561055057600080fd5b5061036660043561149c565b34801561056857600080fd5b506104206004356024356044356114ae565b34801561058657600080fd5b50610420611694565b6040805160206004803580820135601f810184900484028501840190955284845261042094369492936024939284019190819084018382808284375094975050600160a060020a03853516955050505050602001351515611a15565b3480156105f757600080fd5b50610420600435600160a060020a0360243516604435606435611bce565b34801561062157600080fd5b50610366600435602435611dbf565b34801561063c57600080fd5b50610366611ddc565b34801561065157600080fd5b5061065d600435611de2565b60408051938452602084019290925282820152519081900360600190f35b6040805160206004803580820135601f8101849004840285018401909552848452610420943694929360249392840191908190840183828082843750949750508435955050505050602001351515611f88565b3480156106da57600080fd5b506106e3612068565b604080519e8f5260208f019d909d528d8d019b909b5260608d019990995260808c019790975260a08b019590955260c08a0193909352600160a060020a0390911660e08901526101008801526101208701526101408601526101608501526101808401526101a083015251908190036101c00190f35b34801561076557600080fd5b50610420600160a060020a0360043516602435604435612266565b610420600435602435612460565b34801561079a57600080fd5b50610420600435602435612647565b6040805160206004803580820135601f8101849004840285018401909552848452610420943694929360249392840191908190840183828082843750949750508435955050505050602001351515612724565b34801561080857600080fd5b5061038d612804565b610420600160a060020a036004351660243561283b565b34801561083457600080fd5b50610843600435602435612a50565b604080519485526020850193909352838301919091526060830152519081900360800190f35b610420600435602435612a82565b34801561088357600080fd5b50610420600160a060020a0360043516612c7f565b3480156108a457600080fd5b5061052b600435612d9a565b3480156108bc57600080fd5b50610366612db3565b3480156108d157600080fd5b50610366600435602435612e42565b3480156108ec57600080fd5b50610366600435612eea565b34801561090457600080fd5b506104ff612f9c565b34801561091957600080fd5b50610366612fa5565b34801561092e57600080fd5b5061093a600435612fab565b60408051600160a060020a0390981688526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b610420612ff2565b34801561099057600080fd5b506109a5600160a060020a036004351661306f565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b6109e56154ae565b336000908152600660205260408120549080821515610c8857604080517fe56556a9000000000000000000000000000000000000000000000000000000008152336004820152905173cfdf4905f02d4def863a6b459e95089967713d6a9163e56556a99160248083019260209291908290030181600087803b158015610a6a57600080fd5b505af1158015610a7e573d6000803e3d6000fd5b505050506040513d6020811015610a9457600080fd5b5051604080517f82e37b2c00000000000000000000000000000000000000000000000000000000815260048101839052905191945073cfdf4905f02d4def863a6b459e95089967713d6a916382e37b2c916024808201926020929091908290030181600087803b158015610b0757600080fd5b505af1158015610b1b573d6000803e3d6000fd5b505050506040513d6020811015610b3157600080fd5b5051604080517fe3c08adf00000000000000000000000000000000000000000000000000000000815260048101869052905191935073cfdf4905f02d4def863a6b459e95089967713d6a9163e3c08adf916024808201926020929091908290030181600087803b158015610ba457600080fd5b505af1158015610bb8573d6000803e3d6000fd5b505050506040513d6020811015610bce57600080fd5b505133600081815260066020908152604080832088905587835260089091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905590508115610c57576000828152600760209081526040808320869055858352600882528083206001908101869055600a8352818420868552909252909120805460ff191690911790555b8015801590610c665750828114155b15610c805760008381526008602052604090206006018190555b845160010185525b50929392505050565b6005546002546000828152600b602052604090206004015442910181118015610cfc57506000828152600b602052604090206002015481111580610cfc57506000828152600b602052604090206002015481118015610cfc57506000828152600b6020526040902054155b15610d6b57341515610d58576040805160e560020a62461bcd02815260206004820152601060248201527f73656e64696e6720302065746865722e00000000000000000000000000000000604482015290519081900360640190fd5b610d66828734888888613144565b610f1a565b6000828152600b602052604090206002015481118015610d9d57506000828152600b602052604090206003015460ff16155b15610ee5576000828152600b60205260409020600301805460ff19166001179055610dc783613695565b925080670de0b6b3a764000002836000015101836000018181525050858360200151018360200181815250507fa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a3360086000898152602001908152602001600020600101543486600001518760200151886040015189606001518a608001518b60a001518c60c001518d60e00151604051808c600160a060020a0316600160a060020a031681526020018b600019166000191681526020018a815260200189815260200188815260200187600160a060020a0316600160a060020a0316815260200186600019166000191681526020018581526020018481526020018381526020018281526020019b50505050505050505050505060405180910390a15b600086815260086020526040902060030154610f07903463ffffffff613b0816565b6000878152600860205260409020600301555b505050505050565b6005546002546000828152600b602052604081206004015490929142910181118015610f9057506000828152600b602052604090206002015481111580610f9057506000828152600b602052604090206002015481118015610f9057506000828152600b6020526040902054155b15610fd8576000828152600b6020526040902060050154610fd190670de0b6b3a764000090610fc5908263ffffffff613b0816565b9063ffffffff613b6916565b9250610fe2565b6544364c5bb00092505b505090565b60408051808201909152601481527f466f4d6f3444204c6f6e67204f6666696369616c000000000000000000000000602082015281565b6110266154ae565b600f54600090819060ff16151560011461108c576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155488339815191526044820152600080516020615508833981519152606482015290519081900360840190fd5b33803b80156110d3576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b85633b9aca0081101561112b576040805160e560020a62461bcd0281526020600482015260216024820152600080516020615528833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af680000081111561117b576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020615568833981519152604482015290519081900360640190fd5b3360009081526006602052604090205494508815806111aa575060008581526008602052604090206001015489145b156111c8576000858152600860205260409020600601549350611207565b60008981526007602090815260408083205488845260089092529091206006015490945084146112075760008581526008602052604090206006018490555b61121088613b96565b975061121f85858a8a8a613bbb565b505050505050505050565b73dbeb69c655b666b3e17b8061df7ea4cc2399df1133148061125f5750733d3b33b8f50ab9e8f5a9ff369853f0e638450adb33145b15156112b5576040805160e560020a62461bcd02815260206004820152601b60248201527f6f6e6c79207465616d206a7573742063616e2061637469766174650000000000604482015290519081900360640190fd5b600054600160a060020a03161515611317576040805160e560020a62461bcd02815260206004820152601f60248201527f6d757374206c696e6b20746f206f7468657220466f4d6f334420666972737400604482015290519081900360640190fd5b600f5460ff1615611372576040805160e560020a62461bcd02815260206004820152601860248201527f666f6d6f336420616c7265616479206163746976617465640000000000000000604482015290519081900360640190fd5b600f805460ff1916600190811790915560058190556002548154600092909252600b602052429091019081037f72c6bfb7988af3a1efa6568f02a999bc52252641c659d85961ca3d372b57d5d355610e10017f72c6bfb7988af3a1efa6568f02a999bc52252641c659d85961ca3d372b57d5d155565b60066020526000908152604090205481565b60045481565b600b60208190526000918252604090912080546001820154600283015460038401546004850154600586015460068701546007880154600889015460098a0154600a8b01549a909b0154989a9799969860ff90961697949693959294919390928c565b600a60209081526000928352604080842090915290825290205460ff1681565b600d602052600090815260409020805460019091015482565b60076020526000908152604090205481565b6114b66154ae565b600f5460009060ff16151560011461151a576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155488339815191526044820152600080516020615508833981519152606482015290519081900360840190fd5b33803b8015611561576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b84633b9aca008110156115b9576040805160e560020a62461bcd0281526020600482015260216024820152600080516020615528833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af6800000811115611609576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020615568833981519152604482015290519081900360640190fd5b33600090815260066020526040902054935087158061162757508388145b15611645576000848152600860205260409020600601549750611672565b60008481526008602052604090206006015488146116725760008481526008602052604090206006018890555b61167b87613b96565b965061168a8489898989613bbb565b5050505050505050565b6000806000806116a26154ae565b600f5460ff161515600114611703576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155488339815191526044820152600080516020615508833981519152606482015290519081900360840190fd5b33803b801561174a576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b60055433600090815260066020908152604080832054848452600b9092529091206002015491985042975095508611801561179757506000878152600b602052604090206003015460ff16155b80156117b057506000878152600b602052604090205415155b15611956576000878152600b60205260409020600301805460ff191660011790556117da83613695565b92506117e585613dd7565b9350600084111561183657600085815260086020526040808220549051600160a060020a039091169186156108fc02918791818181858888f19350505050158015611834573d6000803e3d6000fd5b505b85670de0b6b3a764000002836000015101836000018181525050848360200151018360200181815250507f0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc3360086000888152602001908152602001600020600101548686600001518760200151886040015189606001518a608001518b60a001518c60c001518d60e00151604051808c600160a060020a0316600160a060020a031681526020018b600019166000191681526020018a815260200189815260200188815260200187600160a060020a0316600160a060020a0316815260200186600019166000191681526020018581526020018481526020018381526020018281526020019b50505050505050505050505060405180910390a1611a0c565b61195f85613dd7565b935060008411156119b057600085815260086020526040808220549051600160a060020a039091169186156108fc02918791818181858888f193505050501580156119ae573d6000803e3d6000fd5b505b6000858152600860209081526040918290206001015482513381529182015280820186905260608101889052905186917f8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a919081900360800190a25b50505050505050565b6000808080808033803b8015611a63576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b611a6c8b613e5e565b604080517faa4d490b000000000000000000000000000000000000000000000000000000008152336004820181905260248201849052600160a060020a038e1660448301528c151560648301528251939b50995034985073cfdf4905f02d4def863a6b459e95089967713d6a9263aa4d490b928a926084808201939182900301818588803b158015611afd57600080fd5b505af1158015611b11573d6000803e3d6000fd5b50505050506040513d6040811015611b2857600080fd5b508051602091820151600160a060020a03808b1660008181526006865260408082205485835260088852918190208054600190910154825188151581529889018790529416878201526060870193909352608086018c90524260a0870152915193995091975095508a92909186917fdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442919081900360c00190a45050505050505050505050565b3373cfdf4905f02d4def863a6b459e95089967713d6a14611c5f576040805160e560020a62461bcd02815260206004820152602760248201527f796f7572206e6f7420706c617965724e616d657320636f6e74726163742e2e2e60448201527f20686d6d6d2e2e00000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0383166000908152600660205260409020548414611c9a57600160a060020a03831660009081526006602052604090208490555b6000828152600760205260409020548414611cc15760008281526007602052604090208490555b600084815260086020526040902054600160a060020a03848116911614611d17576000848152600860205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161790555b6000848152600860205260409020600101548214611d445760008481526008602052604090206001018290555b6000848152600860205260409020600601548114611d715760008481526008602052604090206006018190555b6000848152600a6020908152604080832085845290915290205460ff161515611db9576000848152600a602090815260408083208584529091529020805460ff191660011790555b50505050565b600c60209081526000928352604080842090915290825290205481565b60055481565b6005546000818152600b60205260408120600201549091829182919042118015611e1e57506000818152600b602052604090206003015460ff16155b8015611e3757506000818152600b602052604090205415155b15611f58576000818152600b6020526040902054851415611f1c576000818152600b6020526040902060070154611ea590606490611e7c90603063ffffffff61467116565b811515611e8557fe5b60008881526008602052604090206002015491900463ffffffff613b0816565b6000868152600960209081526040808320858452909152902060020154611efe90611ee090611ed489866146e8565b9063ffffffff6147b616565b6000888152600860205260409020600301549063ffffffff613b0816565b60008781526008602052604090206004015491955093509150611f80565b600085815260086020908152604080832060029081015460098452828520868652909352922090910154611efe90611ee090611ed489866146e8565b60008581526008602052604090206002810154600590910154611efe90611ee0908890614816565b509193909250565b6000808080808033803b8015611fd6576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b611fdf8b613e5e565b604080517f745ea0c1000000000000000000000000000000000000000000000000000000008152336004820181905260248201849052604482018e90528c151560648301528251939b50995034985073cfdf4905f02d4def863a6b459e95089967713d6a9263745ea0c1928a926084808201939182900301818588803b158015611afd57600080fd5b60008060008060008060008060008060008060008060006005549050600b60008281526020019081526020016000206009015481600b600084815260200190815260200160002060050154600b600085815260200190815260200160002060020154600b600086815260200190815260200160002060040154600b600087815260200190815260200160002060070154600b600088815260200190815260200160002060000154600a02600b6000898152602001908152602001600020600101540160086000600b60008b815260200190815260200160002060000154815260200190815260200160002060000160009054906101000a9004600160a060020a031660086000600b60008c815260200190815260200160002060000154815260200190815260200160002060010154600c60008b8152602001908152602001600020600080815260200190815260200160002054600c60008c815260200190815260200160002060006001815260200190815260200160002054600c60008d815260200190815260200160002060006002815260200190815260200160002054600c60008e8152602001908152602001600020600060038152602001908152602001600020546003546103e802600454019e509e509e509e509e509e509e509e509e509e509e509e509e509e5050909192939495969798999a9b9c9d565b61226e6154ae565b600f54600090819060ff1615156001146122d4576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155488339815191526044820152600080516020615508833981519152606482015290519081900360840190fd5b33803b801561231b576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b85633b9aca00811015612373576040805160e560020a62461bcd0281526020600482015260216024820152600080516020615528833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af68000008111156123c3576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020615568833981519152604482015290519081900360640190fd5b336000908152600660205260409020549450600160a060020a03891615806123f35750600160a060020a03891633145b15612411576000858152600860205260409020600601549350611207565b600160a060020a03891660009081526006602081815260408084205489855260089092529092200154909450841461120757600085815260086020526040902060060184905561121088613b96565b6124686154ae565b600f5460009060ff1615156001146124cc576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155488339815191526044820152600080516020615508833981519152606482015290519081900360840190fd5b33803b8015612513576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b34633b9aca0081101561256b576040805160e560020a62461bcd0281526020600482015260216024820152600080516020615528833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af68000008111156125bb576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020615568833981519152604482015290519081900360640190fd5b6125c4856109dd565b3360009081526006602052604090205490955093508615806125e557508387145b15612603576000848152600860205260409020600601549650612630565b60008481526008602052604090206006015487146126305760008481526008602052604090206006018790555b61263986613b96565b9550611a0c84888888610c91565b3373cfdf4905f02d4def863a6b459e95089967713d6a146126d8576040805160e560020a62461bcd02815260206004820152602760248201527f796f7572206e6f7420706c617965724e616d657320636f6e74726163742e2e2e60448201527f20686d6d6d2e2e00000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000828152600a6020908152604080832084845290915290205460ff161515612720576000828152600a602090815260408083208484529091529020805460ff191660011790555b5050565b6000808080808033803b8015612772576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b61277b8b613e5e565b604080517fc0942dfd000000000000000000000000000000000000000000000000000000008152336004820181905260248201849052604482018e90528c151560648301528251939b50995034985073cfdf4905f02d4def863a6b459e95089967713d6a9263c0942dfd928a926084808201939182900301818588803b158015611afd57600080fd5b60408051808201909152600381527f4634440000000000000000000000000000000000000000000000000000000000602082015281565b6128436154ae565b600f54600090819060ff1615156001146128a9576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155488339815191526044820152600080516020615508833981519152606482015290519081900360840190fd5b33803b80156128f0576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b34633b9aca00811015612948576040805160e560020a62461bcd0281526020600482015260216024820152600080516020615528833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af6800000811115612998576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020615568833981519152604482015290519081900360640190fd5b6129a1866109dd565b336000908152600660205260409020549096509450600160a060020a03881615806129d45750600160a060020a03881633145b156129f2576000858152600860205260409020600601549350612a39565b600160a060020a038816600090815260066020818152604080842054898552600890925290922001549094508414612a395760008581526008602052604090206006018490555b612a4287613b96565b965061168a85858989610c91565b600960209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b612a8a6154ae565b600f54600090819060ff161515600114612af0576040805160e560020a62461bcd02815260206004820152602960248201526000805160206155488339815191526044820152600080516020615508833981519152606482015290519081900360840190fd5b33803b8015612b37576040805160e560020a62461bcd0281526020600482015260116024820152600080516020615588833981519152604482015290519081900360640190fd5b34633b9aca00811015612b8f576040805160e560020a62461bcd0281526020600482015260216024820152600080516020615528833981519152604482015260f860020a607902606482015290519081900360840190fd5b69152d02c7e14af6800000811115612bdf576040805160e560020a62461bcd02815260206004820152600e6024820152600080516020615568833981519152604482015290519081900360640190fd5b612be8866109dd565b336000908152600660205260409020549096509450871580612c1a575060008581526008602052604090206001015488145b15612c38576000858152600860205260409020600601549350612a39565b6000888152600760209081526040808320548884526008909252909120600601549094508414612a39576000858152600860205260409020600601849055612a4287613b96565b73dbeb69c655b666b3e17b8061df7ea4cc2399df11331480612cb45750733d3b33b8f50ab9e8f5a9ff369853f0e638450adb33145b1515612d0a576040805160e560020a62461bcd02815260206004820152601b60248201527f6f6e6c79207465616d206a7573742063616e2061637469766174650000000000604482015290519081900360640190fd5b600054600160a060020a031615612d6b576040805160e560020a62461bcd02815260206004820152601f60248201527f73696c6c79206465762c20796f7520616c726561647920646964207468617400604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600e602052600090815260409020805460019091015482565b6005546000818152600b60205260408120600201549091904290811015612e39576002546000838152600b602052604090206004015401811115612e13576000828152600b6020526040902060020154610fd1908263ffffffff6147b616565b6002546000838152600b6020526040902060040154610fd191018263ffffffff6147b616565b60009250610fe2565b6002546000838152600b6020526040812060040154909142910181118015612eac57506000848152600b602052604090206002015481111580612eac57506000848152600b602052604090206002015481118015612eac57506000848152600b6020526040902054155b15612eda576000848152600b6020526040902060060154612ed3908463ffffffff61487316565b9150612ee3565b612ed383614894565b5092915050565b6005546002546000828152600b602052604081206004015490929142910181118015612f5857506000828152600b602052604090206002015481111580612f5857506000828152600b602052604090206002015481118015612f5857506000828152600b6020526040902054155b15612f8c576000828152600b6020526040902060050154612f85908590610fc5908263ffffffff613b0816565b9250612f95565b612f858461490c565b5050919050565b600f5460ff1681565b60035481565b6008602052600090815260409020805460018201546002830154600384015460048501546005860154600690960154600160a060020a039095169593949293919290919087565b6005546001016000818152600b602052604090206007015461301a903463ffffffff613b0816565b6000828152600b6020908152604091829020600701929092558051838152349281019290925280517f74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c9281900390910190a150565b6000806000806000806000806000600554915050600160a060020a038916600090815260066020908152604080832054808452600880845282852060018082015460098752858820898952875294872001549583905293526002830154600590930154909384939091613105906130e7908690614816565b6000878152600860205260409020600301549063ffffffff613b0816565b600095865260086020908152604080882060040154600983528189209989529890915290952054939e929d50909b509950919750919550909350915050565b6000858152600960209081526040808320898452909152812060010154819081908190151561317a576131778986614979565b94505b60008a8152600b602052604090206006015468056bc75e2d631000001180156131d4575060008981526009602090815260408083208d8452909152902054670de0b6b3a7640000906131d2908a63ffffffff613b0816565b115b1561325b5760008981526009602090815260408083208d845290915290205461320c90670de0b6b3a76400009063ffffffff6147b616565b935061321e888563ffffffff6147b616565b60008a815260086020526040902060030154909350613243908463ffffffff613b0816565b60008a81526008602052604090206003015592965086925b633b9aca008811156136895760008a8152600b6020526040902060060154613289908963ffffffff61487316565b9150670de0b6b3a76400008210613300576132a4828b6149d8565b60008a8152600b602052604090205489146132cb5760008a8152600b602052604090208990555b60008a8152600b602052604090206001015486146132f85760008a8152600b602052604090206001018690555b845160640185525b67016345785d8a0000881061354057600480546001019055613320614ab6565b15156001141561354057678ac7230489e8000088106133c15760035460649061335090604b63ffffffff61467116565b81151561335957fe5b60008b8152600860205260409020600201549190049150613380908263ffffffff613b0816565b60008a8152600860205260409020600201556003546133a5908263ffffffff6147b616565b60035584516d0eca8847c4129106ce8300000000018552613515565b670de0b6b3a764000088101580156133e05750678ac7230489e8000088105b1561346d576003546064906133fc90603263ffffffff61467116565b81151561340557fe5b60008b815260086020526040902060020154919004915061342c908263ffffffff613b0816565b60008a815260086020526040902060020155600354613451908263ffffffff6147b616565b60035584516d09dc5ada82b70b59df0200000000018552613515565b67016345785d8a0000881015801561348c5750670de0b6b3a764000088105b15613515576003546064906134a890601963ffffffff61467116565b8115156134b157fe5b60008b81526008602052604090206002015491900491506134d8908263ffffffff613b0816565b60008a8152600860205260409020600201556003546134fd908263ffffffff6147b616565b60035584516d0eca8847c4129106ce83000000000185525b84516d314dc6448d9338c15b0a000000008202016c7e37be2022c0914b268000000001855260006004555b60045485516103e890910201855260008981526009602090815260408083208d845290915290206001015461357c90839063ffffffff613b0816565b60008a81526009602090815260408083208e845290915290206001810191909155546135a9908990613b08565b60008a81526009602090815260408083208e8452825280832093909355600b905220600501546135e090839063ffffffff613b0816565b60008b8152600b6020526040902060058101919091556006015461360b90899063ffffffff613b0816565b60008b8152600b6020908152604080832060060193909355600c81528282208983529052205461364290899063ffffffff613b0816565b60008b8152600c602090815260408083208a845290915290205561366a8a8a8a8a8a8a614ccd565b945061367a8a8a8a89868a614fa3565b945061368989878a8589615111565b50505050505050505050565b61369d6154ae565b6005546000818152600b6020526040812080546001820154600790920154909280808080808060646136d689603063ffffffff61467116565b8115156136df57fe5b04965060328860008b8152600e6020526040902054919004965060649061370d908a9063ffffffff61467116565b81151561371657fe5b60008b8152600e60205260409020600101549190049550606490613741908a9063ffffffff61467116565b81151561374a57fe5b04935061376584611ed487818a818e8e63ffffffff6147b616565b60008c8152600b602052604090206005015490935061379286670de0b6b3a764000063ffffffff61467116565b81151561379b57fe5b60008d8152600b602052604090206005015491900492506137e990670de0b6b3a7640000906137d190859063ffffffff61467116565b8115156137da57fe5b8791900463ffffffff6147b616565b9050600081111561381957613804858263ffffffff6147b616565b9450613816838263ffffffff613b0816565b92505b60008a81526008602052604090206002015461383c90889063ffffffff613b0816565b60008b8152600860205260408082206002019290925581517f6465706f736974282900000000000000000000000000000000000000000000008152825190819003600901812063ffffffff60e060020a918290049081169091028252925173548e2295fc38b69000ff43a730933919b08c256293928a9260048082019391829003018185885af1935050505015156138e5576138de848763ffffffff613b0816565b9350600095505b60008b8152600b602052604090206008015461390890839063ffffffff613b0816565b60008c8152600b60205260408120600801919091558411156139905773e7d5f7a1afbfeb44894c1114fb021cab7e0367fd600160a060020a031663d0e30db0856040518263ffffffff1660e060020a0281526004016000604051808303818588803b15801561397657600080fd5b505af115801561398a573d6000803e3d6000fd5b50505050505b600b60008c815260200190815260200160002060020154620f4240028d60000151018d60000181815250508867016345785d8a0000028a6a52b7d2dcc80cd2e4000000028e6020015101018d6020018181525050600860008b815260200190815260200160002060000160009054906101000a9004600160a060020a03168d60400190600160a060020a03169081600160a060020a031681525050600860008b8152602001908152602001600020600101548d606001906000191690816000191681525050868d6080018181525050848d60e0018181525050838d60c0018181525050828d60a00181815250506005600081548092919060010191905055508a806001019b505042600b60008d815260200190815260200160002060040181905550613ad9600254613acd610e1042613b0890919063ffffffff16565b9063ffffffff613b0816565b60008c8152600b6020526040902060028101919091556007018390558c9b505050505050505050505050919050565b81810182811015613b63576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820616464206661696c656400000000000000000000000000604482015290519081900360640190fd5b92915050565b6000613b8f613b86613b81858563ffffffff6147b616565b61490c565b611ed48561490c565b9392505050565b600080821080613ba65750600382115b15613bb357506002613bb6565b50805b919050565b6005546002546000828152600b602052604090206004015442910181118015613c2657506000828152600b602052604090206002015481111580613c2657506000828152600b602052604090206002015481118015613c2657506000828152600b6020526040902054155b15613c5d57613c3884611ed489613dd7565b600088815260086020526040902060030155613c58828886898988613144565b611a0c565b6000828152600b602052604090206002015481118015613c8f57506000828152600b602052604090206003015460ff16155b15611a0c576000828152600b60205260409020600301805460ff19166001179055613cb983613695565b925080670de0b6b3a764000002836000015101836000018181525050868360200151018360200181815250507f88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd33600860008a815260200190815260200160002060010154856000015186602001518760400151886060015189608001518a60a001518b60c001518c60e00151604051808b600160a060020a0316600160a060020a031681526020018a6000191660001916815260200189815260200188815260200187600160a060020a0316600160a060020a0316815260200186600019166000191681526020018581526020018481526020018381526020018281526020019a505050505050505050505060405180910390a150505050505050565b6000818152600860205260408120600501548190613df690849061527f565b600083815260086020526040902060048101546003820154600290920154613e2892613acd919063ffffffff613b0816565b90506000811115613e545760008381526008602052604081206002810182905560038101829055600401555b8091505b50919050565b8051600090829082808060208411801590613e795750600084115b1515613ef5576040805160e560020a62461bcd02815260206004820152602a60248201527f737472696e67206d757374206265206265747765656e203120616e642033322060448201527f6368617261637465727300000000000000000000000000000000000000000000606482015290519081900360840190fd5b846000815181101515613f0457fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214158015613f6b57508460018503815181101515613f4357fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214155b1515613fe7576040805160e560020a62461bcd02815260206004820152602560248201527f737472696e672063616e6e6f74207374617274206f7220656e6420776974682060448201527f7370616365000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b846000815181101515613ff657fe5b90602001015160f860020a900460f860020a02600160f860020a031916603060f860020a0214156141395784600181518110151561403057fe5b90602001015160f860020a900460f860020a02600160f860020a031916607860f860020a02141515156140ad576040805160e560020a62461bcd02815260206004820152601b60248201527f737472696e672063616e6e6f7420737461727420776974682030780000000000604482015290519081900360640190fd5b8460018151811015156140bc57fe5b90602001015160f860020a900460f860020a02600160f860020a031916605860f860020a0214151515614139576040805160e560020a62461bcd02815260206004820152601b60248201527f737472696e672063616e6e6f7420737461727420776974682030580000000000604482015290519081900360640190fd5b600091505b838210156146095784517f40000000000000000000000000000000000000000000000000000000000000009086908490811061417657fe5b90602001015160f860020a900460f860020a02600160f860020a0319161180156141ea575084517f5b00000000000000000000000000000000000000000000000000000000000000908690849081106141cb57fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b156142575784828151811015156141fd57fe5b90602001015160f860020a900460f860020a0260f860020a900460200160f860020a02858381518110151561422e57fe5b906020010190600160f860020a031916908160001a90535082151561425257600192505b6145fe565b848281518110151561426557fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a021480614335575084517f6000000000000000000000000000000000000000000000000000000000000000908690849081106142c157fe5b90602001015160f860020a900460f860020a02600160f860020a031916118015614335575084517f7b000000000000000000000000000000000000000000000000000000000000009086908490811061431657fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b806143df575084517f2f000000000000000000000000000000000000000000000000000000000000009086908490811061436b57fe5b90602001015160f860020a900460f860020a02600160f860020a0319161180156143df575084517f3a00000000000000000000000000000000000000000000000000000000000000908690849081106143c057fe5b90602001015160f860020a900460f860020a02600160f860020a031916105b151561445b576040805160e560020a62461bcd02815260206004820152602260248201527f737472696e6720636f6e7461696e7320696e76616c696420636861726163746560448201527f7273000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b848281518110151561446957fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214156145485784826001018151811015156144a557fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a0214151515614548576040805160e560020a62461bcd02815260206004820152602860248201527f737472696e672063616e6e6f7420636f6e7461696e20636f6e7365637574697660448201527f6520737061636573000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b821580156145f4575084517f30000000000000000000000000000000000000000000000000000000000000009086908490811061458157fe5b90602001015160f860020a900460f860020a02600160f860020a03191610806145f4575084517f3900000000000000000000000000000000000000000000000000000000000000908690849081106145d557fe5b90602001015160f860020a900460f860020a02600160f860020a031916115b156145fe57600192505b60019091019061413e565b600183151514614663576040805160e560020a62461bcd02815260206004820152601d60248201527f737472696e672063616e6e6f74206265206f6e6c79206e756d62657273000000604482015290519081900360640190fd5b505050506020015192915050565b600082151561468257506000613b63565b5081810281838281151561469257fe5b0414613b63576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d617468206d756c206661696c656400000000000000000000000000604482015290519081900360640190fd5b60008281526009602090815260408083208484528252808320600190810154600b8085528386206005810154938101548752600e8652938620548787529452600790920154670de0b6b3a7640000936147a5939261479992909161477091879160649161475a9163ffffffff61467116565b81151561476357fe5b049063ffffffff61467116565b81151561477957fe5b6000888152600b602052604090206008015491900463ffffffff613b0816565b9063ffffffff61467116565b8115156147ae57fe5b049392505050565b600082821115614810576040805160e560020a62461bcd02815260206004820152601360248201527f536166654d61746820737562206661696c656400000000000000000000000000604482015290519081900360640190fd5b50900390565b600082815260096020908152604080832084845282528083206002810154600190910154600b90935290832060080154613b8f92670de0b6b3a76400009161485d91614671565b81151561486657fe5b049063ffffffff6147b616565b6000613b8f61488184614894565b611ed4614894868663ffffffff613b0816565b60006309502f906148fc6d03b2a1d15167e7c5699bfde00000611ed46148f77a0dac7055469777a6122ee4310dd6c14410500f2904840000000000613acd6b01027e72f1f12813088000006147998a670de0b6b3a764000063ffffffff61467116565b615316565b81151561490557fe5b0492915050565b600061491f670de0b6b3a7640000615369565b6148fc600261495261493f86670de0b6b3a764000063ffffffff61467116565b65886c8f6730709063ffffffff61467116565b81151561495b57fe5b04613acd61496886615369565b6304a817c89063ffffffff61467116565b6149816154ae565b600083815260086020526040902060050154156149b5576000838152600860205260409020600501546149b590849061527f565b506005805460009384526008602052604090932001919091558051600a01815290565b6000818152600b602052604081206002015442919082118015614a0757506000838152600b6020526040902054155b15614a2b57614a2482613acd601e670de0b6b3a764000088614763565b9050614a58565b6000838152600b6020526040902060020154614a5590613acd601e670de0b6b3a764000088614763565b90505b614a6b620151808363ffffffff613b0816565b811015614a8b576000838152600b60205260409020600201819055611db9565b614a9e620151808363ffffffff613b0816565b6000848152600b602052604090206002015550505050565b600080614c2743613acd42336040516020018082600160a060020a0316600160a060020a03166c010000000000000000000000000281526014019150506040516020818303038152906040526040518082805190602001908083835b60208310614b315780518252601f199092019160209182019101614b12565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912092505050811515614b6757fe5b04613acd45613acd42416040516020018082600160a060020a0316600160a060020a03166c010000000000000000000000000281526014019150506040516020818303038152906040526040518082805190602001908083835b60208310614be05780518252601f199092019160209182019101614bc1565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912092505050811515614c1657fe5b04613acd424463ffffffff613b0816565b604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b60208310614c755780518252601f199092019160209182019101614c56565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912060045490945092506103e89150839050046103e80282031015614cc45760019150614cc9565b600091505b5090565b614cd56154ae565b600080808060328904935073548e2295fc38b69000ff43a730933919b08c2562600160a060020a03168460405180807f6465706f736974282900000000000000000000000000000000000000000000008152506009019050604051809103902060e060020a9004906040518263ffffffff1660e060020a02815260040160006040518083038185885af193505050501515614d705760009392505b60008054604080517fed78cf4a000000000000000000000000000000000000000000000000000000008152905160648d049550600160a060020a039092169263ed78cf4a928692600480820193929182900301818588803b158015614dd457600080fd5b505af1158015614de8573d6000803e3d6000fd5b5050505050600a89811515614df957fe5b049050898814158015614e1c575060008881526008602052604090206001015415155b15614ebc57600088815260086020526040902060040154614e4490829063ffffffff613b0816565b600089815260086020908152604091829020600481019390935582546001909301548251600160a060020a03909416845290830152818101839052426060830152518b918d918b917f590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331919081900360800190a4614ec0565b8092505b6000878152600d6020526040902060010154614f0290606490614eea908c9063ffffffff61467116565b811515614ef357fe5b8591900463ffffffff613b0816565b92506000831115614f945773e7d5f7a1afbfeb44894c1114fb021cab7e0367fd600160a060020a031663d0e30db0846040518263ffffffff1660e060020a0281526004016000604051808303818588803b158015614f5f57600080fd5b505af1158015614f73573d6000803e3d6000fd5b50505060c0880151614f8e925085915063ffffffff613b0816565b60c08701525b50939998505050505050505050565b614fab6154ae565b6000848152600d6020526040812054819081908190606490614fd4908b9063ffffffff61467116565b811515614fdd57fe5b049350606489049250614ffb83600354613b0890919063ffffffff16565b6003556000888152600d602052604090206001015461506c9061505f9060649061502c908d9063ffffffff61467116565b81151561503557fe5b0460646150498d600e63ffffffff61467116565b81151561505257fe5b049063ffffffff613b0816565b8a9063ffffffff6147b616565b985061507e898563ffffffff6147b616565b915061508c8b8b868a615375565b905060008111156150aa576150a7848263ffffffff6147b616565b93505b60008b8152600b60205260409020600701546150d090613acd848463ffffffff613b0816565b60008c8152600b602052604090206007015560e08601516150f890859063ffffffff613b0816565b60e0870152506101008501525091979650505050505050565b836c01431e0fae6d7217caa00000000242670de0b6b3a76400000282600001510101816000018181525050600554751aba4714957d300d0e549208b31adb100000000000000285826020015101018160200181815250507f500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c7468160000151826020015160086000898152602001908152602001600020600101543387878760400151886060015189608001518a60a001518b60c001518c60e001518d6101000151600354604051808f81526020018e81526020018d600019166000191681526020018c600160a060020a0316600160a060020a031681526020018b81526020018a815260200189600160a060020a0316600160a060020a0316815260200188600019166000191681526020018781526020018681526020018581526020018481526020018381526020018281526020019e50505050505050505050505050505060405180910390a15050505050565b600061528b8383614816565b90506000811115615311576000838152600860205260409020600301546152b990829063ffffffff613b0816565b60008481526008602090815260408083206003019390935560098152828220858352905220600201546152f390829063ffffffff613b0816565b60008481526009602090815260408083208684529091529020600201555b505050565b6000806002615326846001613b08565b81151561532f57fe5b0490508291505b81811015613e58578091506002615358828581151561535157fe5b0483613b08565b81151561536157fe5b049050615336565b6000613b638283614671565b6000848152600b6020526040812060050154819081906153a386670de0b6b3a764000063ffffffff61467116565b8115156153ac57fe5b6000898152600b602052604090206008015491900492506153d490839063ffffffff613b0816565b6000888152600b6020526040902060080155670de0b6b3a76400006153ff838663ffffffff61467116565b81151561540857fe5b60008881526009602090815260408083208c8452825280832060020154600b9092529091206008015492909104925061545b91613acd908490670de0b6b3a76400009061485d908a63ffffffff61467116565b60008781526009602090815260408083208b8452825280832060020193909355600b905220600501546154a390670de0b6b3a7640000906137d190859063ffffffff61467116565b979650505050505050565b6101206040519081016040528060008152602001600081526020016000600160a060020a03168152602001600080191681526020016000815260200160008152602001600081526020016000815260200160008152509056006e20646973636f72640000000000000000000000000000000000000000000000706f636b6574206c696e743a206e6f7420612076616c69642063757272656e63697473206e6f74207265616479207965742e2020636865636b203f65746120696e6f20766974616c696b2c206e6f000000000000000000000000000000000000736f7272792068756d616e73206f6e6c79000000000000000000000000000000a165627a7a72305820f7bdae8b7fab78b446dbff75ce2b466fb4d9ac1824a95923fe8d18cea6b9a41a0029`

// DeployFoMo3Dlong deploys a new Ethereum contract, binding an instance of FoMo3Dlong to it.
func DeployFoMo3Dlong(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FoMo3Dlong, error) {
	parsed, err := abi.JSON(strings.NewReader(FoMo3DlongABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FoMo3DlongBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FoMo3Dlong{FoMo3DlongCaller: FoMo3DlongCaller{contract: contract}, FoMo3DlongTransactor: FoMo3DlongTransactor{contract: contract}, FoMo3DlongFilterer: FoMo3DlongFilterer{contract: contract}}, nil
}

// FoMo3Dlong is an auto generated Go binding around an Ethereum contract.
type FoMo3Dlong struct {
	FoMo3DlongCaller     // Read-only binding to the contract
	FoMo3DlongTransactor // Write-only binding to the contract
	FoMo3DlongFilterer   // Log filterer for contract events
}

// FoMo3DlongCaller is an auto generated read-only Go binding around an Ethereum contract.
type FoMo3DlongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoMo3DlongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FoMo3DlongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoMo3DlongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FoMo3DlongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoMo3DlongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FoMo3DlongSession struct {
	Contract     *FoMo3Dlong       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FoMo3DlongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FoMo3DlongCallerSession struct {
	Contract *FoMo3DlongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FoMo3DlongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FoMo3DlongTransactorSession struct {
	Contract     *FoMo3DlongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FoMo3DlongRaw is an auto generated low-level Go binding around an Ethereum contract.
type FoMo3DlongRaw struct {
	Contract *FoMo3Dlong // Generic contract binding to access the raw methods on
}

// FoMo3DlongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FoMo3DlongCallerRaw struct {
	Contract *FoMo3DlongCaller // Generic read-only contract binding to access the raw methods on
}

// FoMo3DlongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FoMo3DlongTransactorRaw struct {
	Contract *FoMo3DlongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoMo3Dlong creates a new instance of FoMo3Dlong, bound to a specific deployed contract.
func NewFoMo3Dlong(address common.Address, backend bind.ContractBackend) (*FoMo3Dlong, error) {
	contract, err := bindFoMo3Dlong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FoMo3Dlong{FoMo3DlongCaller: FoMo3DlongCaller{contract: contract}, FoMo3DlongTransactor: FoMo3DlongTransactor{contract: contract}, FoMo3DlongFilterer: FoMo3DlongFilterer{contract: contract}}, nil
}

// NewFoMo3DlongCaller creates a new read-only instance of FoMo3Dlong, bound to a specific deployed contract.
func NewFoMo3DlongCaller(address common.Address, caller bind.ContractCaller) (*FoMo3DlongCaller, error) {
	contract, err := bindFoMo3Dlong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongCaller{contract: contract}, nil
}

// NewFoMo3DlongTransactor creates a new write-only instance of FoMo3Dlong, bound to a specific deployed contract.
func NewFoMo3DlongTransactor(address common.Address, transactor bind.ContractTransactor) (*FoMo3DlongTransactor, error) {
	contract, err := bindFoMo3Dlong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongTransactor{contract: contract}, nil
}

// NewFoMo3DlongFilterer creates a new log filterer instance of FoMo3Dlong, bound to a specific deployed contract.
func NewFoMo3DlongFilterer(address common.Address, filterer bind.ContractFilterer) (*FoMo3DlongFilterer, error) {
	contract, err := bindFoMo3Dlong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongFilterer{contract: contract}, nil
}

// bindFoMo3Dlong binds a generic wrapper to an already deployed contract.
func bindFoMo3Dlong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FoMo3DlongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FoMo3Dlong *FoMo3DlongRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FoMo3Dlong.Contract.FoMo3DlongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FoMo3Dlong *FoMo3DlongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.FoMo3DlongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FoMo3Dlong *FoMo3DlongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.FoMo3DlongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FoMo3Dlong *FoMo3DlongCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FoMo3Dlong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FoMo3Dlong *FoMo3DlongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FoMo3Dlong *FoMo3DlongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.contract.Transact(opts, method, params...)
}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() constant returns(bool)
func (_FoMo3Dlong *FoMo3DlongCaller) Activated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "activated_")
	return *ret0, err
}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() constant returns(bool)
func (_FoMo3Dlong *FoMo3DlongSession) Activated() (bool, error) {
	return _FoMo3Dlong.Contract.Activated(&_FoMo3Dlong.CallOpts)
}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() constant returns(bool)
func (_FoMo3Dlong *FoMo3DlongCallerSession) Activated() (bool, error) {
	return _FoMo3Dlong.Contract.Activated(&_FoMo3Dlong.CallOpts)
}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) AirDropPot(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "airDropPot_")
	return *ret0, err
}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) AirDropPot() (*big.Int, error) {
	return _FoMo3Dlong.Contract.AirDropPot(&_FoMo3Dlong.CallOpts)
}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) AirDropPot() (*big.Int, error) {
	return _FoMo3Dlong.Contract.AirDropPot(&_FoMo3Dlong.CallOpts)
}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) AirDropTracker(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "airDropTracker_")
	return *ret0, err
}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) AirDropTracker() (*big.Int, error) {
	return _FoMo3Dlong.Contract.AirDropTracker(&_FoMo3Dlong.CallOpts)
}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) AirDropTracker() (*big.Int, error) {
	return _FoMo3Dlong.Contract.AirDropTracker(&_FoMo3Dlong.CallOpts)
}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(_rID uint256, _eth uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) CalcKeysReceived(opts *bind.CallOpts, _rID *big.Int, _eth *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "calcKeysReceived", _rID, _eth)
	return *ret0, err
}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(_rID uint256, _eth uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) CalcKeysReceived(_rID *big.Int, _eth *big.Int) (*big.Int, error) {
	return _FoMo3Dlong.Contract.CalcKeysReceived(&_FoMo3Dlong.CallOpts, _rID, _eth)
}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(_rID uint256, _eth uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) CalcKeysReceived(_rID *big.Int, _eth *big.Int) (*big.Int, error) {
	return _FoMo3Dlong.Contract.CalcKeysReceived(&_FoMo3Dlong.CallOpts, _rID, _eth)
}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_( uint256) constant returns(gen uint256, p3d uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) Fees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	ret := new(struct {
		Gen *big.Int
		P3d *big.Int
	})
	out := ret
	err := _FoMo3Dlong.contract.Call(opts, out, "fees_", arg0)
	return *ret, err
}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_( uint256) constant returns(gen uint256, p3d uint256)
func (_FoMo3Dlong *FoMo3DlongSession) Fees(arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	return _FoMo3Dlong.Contract.Fees(&_FoMo3Dlong.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_( uint256) constant returns(gen uint256, p3d uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) Fees(arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	return _FoMo3Dlong.Contract.Fees(&_FoMo3Dlong.CallOpts, arg0)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) GetBuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "getBuyPrice")
	return *ret0, err
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) GetBuyPrice() (*big.Int, error) {
	return _FoMo3Dlong.Contract.GetBuyPrice(&_FoMo3Dlong.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) GetBuyPrice() (*big.Int, error) {
	return _FoMo3Dlong.Contract.GetBuyPrice(&_FoMo3Dlong.CallOpts)
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) GetCurrentRoundInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0  = new(*big.Int)
		ret1  = new(*big.Int)
		ret2  = new(*big.Int)
		ret3  = new(*big.Int)
		ret4  = new(*big.Int)
		ret5  = new(*big.Int)
		ret6  = new(*big.Int)
		ret7  = new(common.Address)
		ret8  = new([32]byte)
		ret9  = new(*big.Int)
		ret10 = new(*big.Int)
		ret11 = new(*big.Int)
		ret12 = new(*big.Int)
		ret13 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
		ret9,
		ret10,
		ret11,
		ret12,
		ret13,
	}
	err := _FoMo3Dlong.contract.Call(opts, out, "getCurrentRoundInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, *ret9, *ret10, *ret11, *ret12, *ret13, err
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongSession) GetCurrentRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FoMo3Dlong.Contract.GetCurrentRoundInfo(&_FoMo3Dlong.CallOpts)
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) GetCurrentRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FoMo3Dlong.Contract.GetCurrentRoundInfo(&_FoMo3Dlong.CallOpts)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(_addr address) constant returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) GetPlayerInfoByAddress(opts *bind.CallOpts, _addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _FoMo3Dlong.contract.Call(opts, out, "getPlayerInfoByAddress", _addr)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(_addr address) constant returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongSession) GetPlayerInfoByAddress(_addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FoMo3Dlong.Contract.GetPlayerInfoByAddress(&_FoMo3Dlong.CallOpts, _addr)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(_addr address) constant returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) GetPlayerInfoByAddress(_addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FoMo3Dlong.Contract.GetPlayerInfoByAddress(&_FoMo3Dlong.CallOpts, _addr)
}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(_pID uint256) constant returns(uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) GetPlayerVaults(opts *bind.CallOpts, _pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _FoMo3Dlong.contract.Call(opts, out, "getPlayerVaults", _pID)
	return *ret0, *ret1, *ret2, err
}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(_pID uint256) constant returns(uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongSession) GetPlayerVaults(_pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _FoMo3Dlong.Contract.GetPlayerVaults(&_FoMo3Dlong.CallOpts, _pID)
}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(_pID uint256) constant returns(uint256, uint256, uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) GetPlayerVaults(_pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _FoMo3Dlong.Contract.GetPlayerVaults(&_FoMo3Dlong.CallOpts, _pID)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) GetTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "getTimeLeft")
	return *ret0, err
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) GetTimeLeft() (*big.Int, error) {
	return _FoMo3Dlong.Contract.GetTimeLeft(&_FoMo3Dlong.CallOpts)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) GetTimeLeft() (*big.Int, error) {
	return _FoMo3Dlong.Contract.GetTimeLeft(&_FoMo3Dlong.CallOpts)
}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(_keys uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) IWantXKeys(opts *bind.CallOpts, _keys *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "iWantXKeys", _keys)
	return *ret0, err
}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(_keys uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) IWantXKeys(_keys *big.Int) (*big.Int, error) {
	return _FoMo3Dlong.Contract.IWantXKeys(&_FoMo3Dlong.CallOpts, _keys)
}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(_keys uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) IWantXKeys(_keys *big.Int) (*big.Int, error) {
	return _FoMo3Dlong.Contract.IWantXKeys(&_FoMo3Dlong.CallOpts, _keys)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FoMo3Dlong *FoMo3DlongCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FoMo3Dlong *FoMo3DlongSession) Name() (string, error) {
	return _FoMo3Dlong.Contract.Name(&_FoMo3Dlong.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FoMo3Dlong *FoMo3DlongCallerSession) Name() (string, error) {
	return _FoMo3Dlong.Contract.Name(&_FoMo3Dlong.CallOpts)
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) PIDxAddr(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "pIDxAddr_", arg0)
	return *ret0, err
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) PIDxAddr(arg0 common.Address) (*big.Int, error) {
	return _FoMo3Dlong.Contract.PIDxAddr(&_FoMo3Dlong.CallOpts, arg0)
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_( address) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) PIDxAddr(arg0 common.Address) (*big.Int, error) {
	return _FoMo3Dlong.Contract.PIDxAddr(&_FoMo3Dlong.CallOpts, arg0)
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) PIDxName(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "pIDxName_", arg0)
	return *ret0, err
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) PIDxName(arg0 [32]byte) (*big.Int, error) {
	return _FoMo3Dlong.Contract.PIDxName(&_FoMo3Dlong.CallOpts, arg0)
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_( bytes32) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) PIDxName(arg0 [32]byte) (*big.Int, error) {
	return _FoMo3Dlong.Contract.PIDxName(&_FoMo3Dlong.CallOpts, arg0)
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_FoMo3Dlong *FoMo3DlongCaller) PlyrNames(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "plyrNames_", arg0, arg1)
	return *ret0, err
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_FoMo3Dlong *FoMo3DlongSession) PlyrNames(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _FoMo3Dlong.Contract.PlyrNames(&_FoMo3Dlong.CallOpts, arg0, arg1)
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_( uint256,  bytes32) constant returns(bool)
func (_FoMo3Dlong *FoMo3DlongCallerSession) PlyrNames(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _FoMo3Dlong.Contract.PlyrNames(&_FoMo3Dlong.CallOpts, arg0, arg1)
}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_( uint256,  uint256) constant returns(eth uint256, keys uint256, mask uint256, ico uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) PlyrRnds(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Eth  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	ret := new(struct {
		Eth  *big.Int
		Keys *big.Int
		Mask *big.Int
		Ico  *big.Int
	})
	out := ret
	err := _FoMo3Dlong.contract.Call(opts, out, "plyrRnds_", arg0, arg1)
	return *ret, err
}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_( uint256,  uint256) constant returns(eth uint256, keys uint256, mask uint256, ico uint256)
func (_FoMo3Dlong *FoMo3DlongSession) PlyrRnds(arg0 *big.Int, arg1 *big.Int) (struct {
	Eth  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	return _FoMo3Dlong.Contract.PlyrRnds(&_FoMo3Dlong.CallOpts, arg0, arg1)
}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_( uint256,  uint256) constant returns(eth uint256, keys uint256, mask uint256, ico uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) PlyrRnds(arg0 *big.Int, arg1 *big.Int) (struct {
	Eth  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	return _FoMo3Dlong.Contract.PlyrRnds(&_FoMo3Dlong.CallOpts, arg0, arg1)
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, win uint256, gen uint256, aff uint256, lrnd uint256, laff uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) Plyr(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	ret := new(struct {
		Addr common.Address
		Name [32]byte
		Win  *big.Int
		Gen  *big.Int
		Aff  *big.Int
		Lrnd *big.Int
		Laff *big.Int
	})
	out := ret
	err := _FoMo3Dlong.contract.Call(opts, out, "plyr_", arg0)
	return *ret, err
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, win uint256, gen uint256, aff uint256, lrnd uint256, laff uint256)
func (_FoMo3Dlong *FoMo3DlongSession) Plyr(arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	return _FoMo3Dlong.Contract.Plyr(&_FoMo3Dlong.CallOpts, arg0)
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_( uint256) constant returns(addr address, name bytes32, win uint256, gen uint256, aff uint256, lrnd uint256, laff uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) Plyr(arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	return _FoMo3Dlong.Contract.Plyr(&_FoMo3Dlong.CallOpts, arg0)
}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_( uint256) constant returns(gen uint256, p3d uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) PotSplit(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	ret := new(struct {
		Gen *big.Int
		P3d *big.Int
	})
	out := ret
	err := _FoMo3Dlong.contract.Call(opts, out, "potSplit_", arg0)
	return *ret, err
}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_( uint256) constant returns(gen uint256, p3d uint256)
func (_FoMo3Dlong *FoMo3DlongSession) PotSplit(arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	return _FoMo3Dlong.Contract.PotSplit(&_FoMo3Dlong.CallOpts, arg0)
}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_( uint256) constant returns(gen uint256, p3d uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) PotSplit(arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	return _FoMo3Dlong.Contract.PotSplit(&_FoMo3Dlong.CallOpts, arg0)
}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) RID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "rID_")
	return *ret0, err
}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) RID() (*big.Int, error) {
	return _FoMo3Dlong.Contract.RID(&_FoMo3Dlong.CallOpts)
}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) RID() (*big.Int, error) {
	return _FoMo3Dlong.Contract.RID(&_FoMo3Dlong.CallOpts)
}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_( uint256,  uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) RndTmEth(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "rndTmEth_", arg0, arg1)
	return *ret0, err
}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_( uint256,  uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongSession) RndTmEth(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _FoMo3Dlong.Contract.RndTmEth(&_FoMo3Dlong.CallOpts, arg0, arg1)
}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_( uint256,  uint256) constant returns(uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) RndTmEth(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _FoMo3Dlong.Contract.RndTmEth(&_FoMo3Dlong.CallOpts, arg0, arg1)
}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_( uint256) constant returns(plyr uint256, team uint256, end uint256, ended bool, strt uint256, keys uint256, eth uint256, pot uint256, mask uint256, ico uint256, icoGen uint256, icoAvg uint256)
func (_FoMo3Dlong *FoMo3DlongCaller) Round(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Eth    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	ret := new(struct {
		Plyr   *big.Int
		Team   *big.Int
		End    *big.Int
		Ended  bool
		Strt   *big.Int
		Keys   *big.Int
		Eth    *big.Int
		Pot    *big.Int
		Mask   *big.Int
		Ico    *big.Int
		IcoGen *big.Int
		IcoAvg *big.Int
	})
	out := ret
	err := _FoMo3Dlong.contract.Call(opts, out, "round_", arg0)
	return *ret, err
}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_( uint256) constant returns(plyr uint256, team uint256, end uint256, ended bool, strt uint256, keys uint256, eth uint256, pot uint256, mask uint256, ico uint256, icoGen uint256, icoAvg uint256)
func (_FoMo3Dlong *FoMo3DlongSession) Round(arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Eth    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	return _FoMo3Dlong.Contract.Round(&_FoMo3Dlong.CallOpts, arg0)
}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_( uint256) constant returns(plyr uint256, team uint256, end uint256, ended bool, strt uint256, keys uint256, eth uint256, pot uint256, mask uint256, ico uint256, icoGen uint256, icoAvg uint256)
func (_FoMo3Dlong *FoMo3DlongCallerSession) Round(arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Eth    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	return _FoMo3Dlong.Contract.Round(&_FoMo3Dlong.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FoMo3Dlong *FoMo3DlongCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FoMo3Dlong.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FoMo3Dlong *FoMo3DlongSession) Symbol() (string, error) {
	return _FoMo3Dlong.Contract.Symbol(&_FoMo3Dlong.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FoMo3Dlong *FoMo3DlongCallerSession) Symbol() (string, error) {
	return _FoMo3Dlong.Contract.Symbol(&_FoMo3Dlong.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FoMo3Dlong *FoMo3DlongSession) Activate() (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.Activate(&_FoMo3Dlong.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) Activate() (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.Activate(&_FoMo3Dlong.TransactOpts)
}

// BuyXaddr is a paid mutator transaction binding the contract method 0x98a0871d.
//
// Solidity: function buyXaddr(_affCode address, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) BuyXaddr(opts *bind.TransactOpts, _affCode common.Address, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "buyXaddr", _affCode, _team)
}

// BuyXaddr is a paid mutator transaction binding the contract method 0x98a0871d.
//
// Solidity: function buyXaddr(_affCode address, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongSession) BuyXaddr(_affCode common.Address, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.BuyXaddr(&_FoMo3Dlong.TransactOpts, _affCode, _team)
}

// BuyXaddr is a paid mutator transaction binding the contract method 0x98a0871d.
//
// Solidity: function buyXaddr(_affCode address, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) BuyXaddr(_affCode common.Address, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.BuyXaddr(&_FoMo3Dlong.TransactOpts, _affCode, _team)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(_affCode uint256, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) BuyXid(opts *bind.TransactOpts, _affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "buyXid", _affCode, _team)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(_affCode uint256, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongSession) BuyXid(_affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.BuyXid(&_FoMo3Dlong.TransactOpts, _affCode, _team)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(_affCode uint256, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) BuyXid(_affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.BuyXid(&_FoMo3Dlong.TransactOpts, _affCode, _team)
}

// BuyXname is a paid mutator transaction binding the contract method 0xa65b37a1.
//
// Solidity: function buyXname(_affCode bytes32, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) BuyXname(opts *bind.TransactOpts, _affCode [32]byte, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "buyXname", _affCode, _team)
}

// BuyXname is a paid mutator transaction binding the contract method 0xa65b37a1.
//
// Solidity: function buyXname(_affCode bytes32, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongSession) BuyXname(_affCode [32]byte, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.BuyXname(&_FoMo3Dlong.TransactOpts, _affCode, _team)
}

// BuyXname is a paid mutator transaction binding the contract method 0xa65b37a1.
//
// Solidity: function buyXname(_affCode bytes32, _team uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) BuyXname(_affCode [32]byte, _team *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.BuyXname(&_FoMo3Dlong.TransactOpts, _affCode, _team)
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) PotSwap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "potSwap")
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_FoMo3Dlong *FoMo3DlongSession) PotSwap() (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.PotSwap(&_FoMo3Dlong.TransactOpts)
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) PotSwap() (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.PotSwap(&_FoMo3Dlong.TransactOpts)
}

// ReLoadXaddr is a paid mutator transaction binding the contract method 0x82bfc739.
//
// Solidity: function reLoadXaddr(_affCode address, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) ReLoadXaddr(opts *bind.TransactOpts, _affCode common.Address, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "reLoadXaddr", _affCode, _team, _eth)
}

// ReLoadXaddr is a paid mutator transaction binding the contract method 0x82bfc739.
//
// Solidity: function reLoadXaddr(_affCode address, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongSession) ReLoadXaddr(_affCode common.Address, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReLoadXaddr(&_FoMo3Dlong.TransactOpts, _affCode, _team, _eth)
}

// ReLoadXaddr is a paid mutator transaction binding the contract method 0x82bfc739.
//
// Solidity: function reLoadXaddr(_affCode address, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) ReLoadXaddr(_affCode common.Address, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReLoadXaddr(&_FoMo3Dlong.TransactOpts, _affCode, _team, _eth)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(_affCode uint256, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) ReLoadXid(opts *bind.TransactOpts, _affCode *big.Int, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "reLoadXid", _affCode, _team, _eth)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(_affCode uint256, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongSession) ReLoadXid(_affCode *big.Int, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReLoadXid(&_FoMo3Dlong.TransactOpts, _affCode, _team, _eth)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(_affCode uint256, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) ReLoadXid(_affCode *big.Int, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReLoadXid(&_FoMo3Dlong.TransactOpts, _affCode, _team, _eth)
}

// ReLoadXname is a paid mutator transaction binding the contract method 0x079ce327.
//
// Solidity: function reLoadXname(_affCode bytes32, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) ReLoadXname(opts *bind.TransactOpts, _affCode [32]byte, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "reLoadXname", _affCode, _team, _eth)
}

// ReLoadXname is a paid mutator transaction binding the contract method 0x079ce327.
//
// Solidity: function reLoadXname(_affCode bytes32, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongSession) ReLoadXname(_affCode [32]byte, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReLoadXname(&_FoMo3Dlong.TransactOpts, _affCode, _team, _eth)
}

// ReLoadXname is a paid mutator transaction binding the contract method 0x079ce327.
//
// Solidity: function reLoadXname(_affCode bytes32, _team uint256, _eth uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) ReLoadXname(_affCode [32]byte, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReLoadXname(&_FoMo3Dlong.TransactOpts, _affCode, _team, _eth)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) ReceivePlayerInfo(opts *bind.TransactOpts, _pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "receivePlayerInfo", _pID, _addr, _name, _laff)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_FoMo3Dlong *FoMo3DlongSession) ReceivePlayerInfo(_pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReceivePlayerInfo(&_FoMo3Dlong.TransactOpts, _pID, _addr, _name, _laff)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(_pID uint256, _addr address, _name bytes32, _laff uint256) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) ReceivePlayerInfo(_pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReceivePlayerInfo(&_FoMo3Dlong.TransactOpts, _pID, _addr, _name, _laff)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) ReceivePlayerNameList(opts *bind.TransactOpts, _pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "receivePlayerNameList", _pID, _name)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_FoMo3Dlong *FoMo3DlongSession) ReceivePlayerNameList(_pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReceivePlayerNameList(&_FoMo3Dlong.TransactOpts, _pID, _name)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(_pID uint256, _name bytes32) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) ReceivePlayerNameList(_pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.ReceivePlayerNameList(&_FoMo3Dlong.TransactOpts, _pID, _name)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) RegisterNameXID(opts *bind.TransactOpts, _nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "registerNameXID", _nameString, _affCode, _all)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongSession) RegisterNameXID(_nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.RegisterNameXID(&_FoMo3Dlong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXID is a paid mutator transaction binding the contract method 0x921dec21.
//
// Solidity: function registerNameXID(_nameString string, _affCode uint256, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) RegisterNameXID(_nameString string, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.RegisterNameXID(&_FoMo3Dlong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) RegisterNameXaddr(opts *bind.TransactOpts, _nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "registerNameXaddr", _nameString, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongSession) RegisterNameXaddr(_nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.RegisterNameXaddr(&_FoMo3Dlong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXaddr is a paid mutator transaction binding the contract method 0x3ddd4698.
//
// Solidity: function registerNameXaddr(_nameString string, _affCode address, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) RegisterNameXaddr(_nameString string, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.RegisterNameXaddr(&_FoMo3Dlong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) RegisterNameXname(opts *bind.TransactOpts, _nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "registerNameXname", _nameString, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongSession) RegisterNameXname(_nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.RegisterNameXname(&_FoMo3Dlong.TransactOpts, _nameString, _affCode, _all)
}

// RegisterNameXname is a paid mutator transaction binding the contract method 0x685ffd83.
//
// Solidity: function registerNameXname(_nameString string, _affCode bytes32, _all bool) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) RegisterNameXname(_nameString string, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.RegisterNameXname(&_FoMo3Dlong.TransactOpts, _nameString, _affCode, _all)
}

// SetOtherFomo is a paid mutator transaction binding the contract method 0xb483c054.
//
// Solidity: function setOtherFomo(_otherF3D address) returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) SetOtherFomo(opts *bind.TransactOpts, _otherF3D common.Address) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "setOtherFomo", _otherF3D)
}

// SetOtherFomo is a paid mutator transaction binding the contract method 0xb483c054.
//
// Solidity: function setOtherFomo(_otherF3D address) returns()
func (_FoMo3Dlong *FoMo3DlongSession) SetOtherFomo(_otherF3D common.Address) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.SetOtherFomo(&_FoMo3Dlong.TransactOpts, _otherF3D)
}

// SetOtherFomo is a paid mutator transaction binding the contract method 0xb483c054.
//
// Solidity: function setOtherFomo(_otherF3D address) returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) SetOtherFomo(_otherF3D common.Address) (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.SetOtherFomo(&_FoMo3Dlong.TransactOpts, _otherF3D)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FoMo3Dlong *FoMo3DlongTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FoMo3Dlong.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FoMo3Dlong *FoMo3DlongSession) Withdraw() (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.Withdraw(&_FoMo3Dlong.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FoMo3Dlong *FoMo3DlongTransactorSession) Withdraw() (*types.Transaction, error) {
	return _FoMo3Dlong.Contract.Withdraw(&_FoMo3Dlong.TransactOpts)
}

// FoMo3DlongOnAffiliatePayoutIterator is returned from FilterOnAffiliatePayout and is used to iterate over the raw logs and unpacked data for OnAffiliatePayout events raised by the FoMo3Dlong contract.
type FoMo3DlongOnAffiliatePayoutIterator struct {
	Event *FoMo3DlongOnAffiliatePayout // Event containing the contract specifics and raw log

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
func (it *FoMo3DlongOnAffiliatePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoMo3DlongOnAffiliatePayout)
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
		it.Event = new(FoMo3DlongOnAffiliatePayout)
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
func (it *FoMo3DlongOnAffiliatePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoMo3DlongOnAffiliatePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoMo3DlongOnAffiliatePayout represents a OnAffiliatePayout event raised by the FoMo3Dlong contract.
type FoMo3DlongOnAffiliatePayout struct {
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	RoundID          *big.Int
	BuyerID          *big.Int
	Amount           *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnAffiliatePayout is a free log retrieval operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: e onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) FilterOnAffiliatePayout(opts *bind.FilterOpts, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (*FoMo3DlongOnAffiliatePayoutIterator, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _FoMo3Dlong.contract.FilterLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongOnAffiliatePayoutIterator{contract: _FoMo3Dlong.contract, event: "onAffiliatePayout", logs: logs, sub: sub}, nil
}

// WatchOnAffiliatePayout is a free log subscription operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: e onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) WatchOnAffiliatePayout(opts *bind.WatchOpts, sink chan<- *FoMo3DlongOnAffiliatePayout, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (event.Subscription, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _FoMo3Dlong.contract.WatchLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoMo3DlongOnAffiliatePayout)
				if err := _FoMo3Dlong.contract.UnpackLog(event, "onAffiliatePayout", log); err != nil {
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

// FoMo3DlongOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the FoMo3Dlong contract.
type FoMo3DlongOnBuyAndDistributeIterator struct {
	Event *FoMo3DlongOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *FoMo3DlongOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoMo3DlongOnBuyAndDistribute)
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
		it.Event = new(FoMo3DlongOnBuyAndDistribute)
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
func (it *FoMo3DlongOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoMo3DlongOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoMo3DlongOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the FoMo3Dlong contract.
type FoMo3DlongOnBuyAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EthIn          *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: e onBuyAndDistribute(playerAddress address, playerName bytes32, ethIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*FoMo3DlongOnBuyAndDistributeIterator, error) {

	logs, sub, err := _FoMo3Dlong.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongOnBuyAndDistributeIterator{contract: _FoMo3Dlong.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: e onBuyAndDistribute(playerAddress address, playerName bytes32, ethIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *FoMo3DlongOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FoMo3Dlong.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoMo3DlongOnBuyAndDistribute)
				if err := _FoMo3Dlong.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// FoMo3DlongOnEndTxIterator is returned from FilterOnEndTx and is used to iterate over the raw logs and unpacked data for OnEndTx events raised by the FoMo3Dlong contract.
type FoMo3DlongOnEndTxIterator struct {
	Event *FoMo3DlongOnEndTx // Event containing the contract specifics and raw log

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
func (it *FoMo3DlongOnEndTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoMo3DlongOnEndTx)
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
		it.Event = new(FoMo3DlongOnEndTx)
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
func (it *FoMo3DlongOnEndTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoMo3DlongOnEndTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoMo3DlongOnEndTx represents a OnEndTx event raised by the FoMo3Dlong contract.
type FoMo3DlongOnEndTx struct {
	CompressedData *big.Int
	CompressedIDs  *big.Int
	PlayerName     [32]byte
	PlayerAddress  common.Address
	EthIn          *big.Int
	KeysBought     *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	PotAmount      *big.Int
	AirDropPot     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnEndTx is a free log retrieval operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: e onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, ethIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) FilterOnEndTx(opts *bind.FilterOpts) (*FoMo3DlongOnEndTxIterator, error) {

	logs, sub, err := _FoMo3Dlong.contract.FilterLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongOnEndTxIterator{contract: _FoMo3Dlong.contract, event: "onEndTx", logs: logs, sub: sub}, nil
}

// WatchOnEndTx is a free log subscription operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: e onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, ethIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) WatchOnEndTx(opts *bind.WatchOpts, sink chan<- *FoMo3DlongOnEndTx) (event.Subscription, error) {

	logs, sub, err := _FoMo3Dlong.contract.WatchLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoMo3DlongOnEndTx)
				if err := _FoMo3Dlong.contract.UnpackLog(event, "onEndTx", log); err != nil {
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

// FoMo3DlongOnNewNameIterator is returned from FilterOnNewName and is used to iterate over the raw logs and unpacked data for OnNewName events raised by the FoMo3Dlong contract.
type FoMo3DlongOnNewNameIterator struct {
	Event *FoMo3DlongOnNewName // Event containing the contract specifics and raw log

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
func (it *FoMo3DlongOnNewNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoMo3DlongOnNewName)
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
		it.Event = new(FoMo3DlongOnNewName)
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
func (it *FoMo3DlongOnNewNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoMo3DlongOnNewNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoMo3DlongOnNewName represents a OnNewName event raised by the FoMo3Dlong contract.
type FoMo3DlongOnNewName struct {
	PlayerID         *big.Int
	PlayerAddress    common.Address
	PlayerName       [32]byte
	IsNewPlayer      bool
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	AmountPaid       *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnNewName is a free log retrieval operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: e onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) FilterOnNewName(opts *bind.FilterOpts, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (*FoMo3DlongOnNewNameIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _FoMo3Dlong.contract.FilterLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongOnNewNameIterator{contract: _FoMo3Dlong.contract, event: "onNewName", logs: logs, sub: sub}, nil
}

// WatchOnNewName is a free log subscription operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: e onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) WatchOnNewName(opts *bind.WatchOpts, sink chan<- *FoMo3DlongOnNewName, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _FoMo3Dlong.contract.WatchLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoMo3DlongOnNewName)
				if err := _FoMo3Dlong.contract.UnpackLog(event, "onNewName", log); err != nil {
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

// FoMo3DlongOnPotSwapDepositIterator is returned from FilterOnPotSwapDeposit and is used to iterate over the raw logs and unpacked data for OnPotSwapDeposit events raised by the FoMo3Dlong contract.
type FoMo3DlongOnPotSwapDepositIterator struct {
	Event *FoMo3DlongOnPotSwapDeposit // Event containing the contract specifics and raw log

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
func (it *FoMo3DlongOnPotSwapDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoMo3DlongOnPotSwapDeposit)
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
		it.Event = new(FoMo3DlongOnPotSwapDeposit)
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
func (it *FoMo3DlongOnPotSwapDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoMo3DlongOnPotSwapDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoMo3DlongOnPotSwapDeposit represents a OnPotSwapDeposit event raised by the FoMo3Dlong contract.
type FoMo3DlongOnPotSwapDeposit struct {
	RoundID          *big.Int
	AmountAddedToPot *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnPotSwapDeposit is a free log retrieval operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: e onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) FilterOnPotSwapDeposit(opts *bind.FilterOpts) (*FoMo3DlongOnPotSwapDepositIterator, error) {

	logs, sub, err := _FoMo3Dlong.contract.FilterLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongOnPotSwapDepositIterator{contract: _FoMo3Dlong.contract, event: "onPotSwapDeposit", logs: logs, sub: sub}, nil
}

// WatchOnPotSwapDeposit is a free log subscription operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: e onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) WatchOnPotSwapDeposit(opts *bind.WatchOpts, sink chan<- *FoMo3DlongOnPotSwapDeposit) (event.Subscription, error) {

	logs, sub, err := _FoMo3Dlong.contract.WatchLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoMo3DlongOnPotSwapDeposit)
				if err := _FoMo3Dlong.contract.UnpackLog(event, "onPotSwapDeposit", log); err != nil {
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

// FoMo3DlongOnReLoadAndDistributeIterator is returned from FilterOnReLoadAndDistribute and is used to iterate over the raw logs and unpacked data for OnReLoadAndDistribute events raised by the FoMo3Dlong contract.
type FoMo3DlongOnReLoadAndDistributeIterator struct {
	Event *FoMo3DlongOnReLoadAndDistribute // Event containing the contract specifics and raw log

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
func (it *FoMo3DlongOnReLoadAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoMo3DlongOnReLoadAndDistribute)
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
		it.Event = new(FoMo3DlongOnReLoadAndDistribute)
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
func (it *FoMo3DlongOnReLoadAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoMo3DlongOnReLoadAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoMo3DlongOnReLoadAndDistribute represents a OnReLoadAndDistribute event raised by the FoMo3Dlong contract.
type FoMo3DlongOnReLoadAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnReLoadAndDistribute is a free log retrieval operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: e onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) FilterOnReLoadAndDistribute(opts *bind.FilterOpts) (*FoMo3DlongOnReLoadAndDistributeIterator, error) {

	logs, sub, err := _FoMo3Dlong.contract.FilterLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongOnReLoadAndDistributeIterator{contract: _FoMo3Dlong.contract, event: "onReLoadAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnReLoadAndDistribute is a free log subscription operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: e onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) WatchOnReLoadAndDistribute(opts *bind.WatchOpts, sink chan<- *FoMo3DlongOnReLoadAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FoMo3Dlong.contract.WatchLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoMo3DlongOnReLoadAndDistribute)
				if err := _FoMo3Dlong.contract.UnpackLog(event, "onReLoadAndDistribute", log); err != nil {
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

// FoMo3DlongOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the FoMo3Dlong contract.
type FoMo3DlongOnWithdrawIterator struct {
	Event *FoMo3DlongOnWithdraw // Event containing the contract specifics and raw log

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
func (it *FoMo3DlongOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoMo3DlongOnWithdraw)
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
		it.Event = new(FoMo3DlongOnWithdraw)
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
func (it *FoMo3DlongOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoMo3DlongOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoMo3DlongOnWithdraw represents a OnWithdraw event raised by the FoMo3Dlong contract.
type FoMo3DlongOnWithdraw struct {
	PlayerID      *big.Int
	PlayerAddress common.Address
	PlayerName    [32]byte
	EthOut        *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: e onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, ethOut uint256, timeStamp uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) FilterOnWithdraw(opts *bind.FilterOpts, playerID []*big.Int) (*FoMo3DlongOnWithdrawIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _FoMo3Dlong.contract.FilterLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongOnWithdrawIterator{contract: _FoMo3Dlong.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: e onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, ethOut uint256, timeStamp uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *FoMo3DlongOnWithdraw, playerID []*big.Int) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _FoMo3Dlong.contract.WatchLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoMo3DlongOnWithdraw)
				if err := _FoMo3Dlong.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// FoMo3DlongOnWithdrawAndDistributeIterator is returned from FilterOnWithdrawAndDistribute and is used to iterate over the raw logs and unpacked data for OnWithdrawAndDistribute events raised by the FoMo3Dlong contract.
type FoMo3DlongOnWithdrawAndDistributeIterator struct {
	Event *FoMo3DlongOnWithdrawAndDistribute // Event containing the contract specifics and raw log

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
func (it *FoMo3DlongOnWithdrawAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoMo3DlongOnWithdrawAndDistribute)
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
		it.Event = new(FoMo3DlongOnWithdrawAndDistribute)
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
func (it *FoMo3DlongOnWithdrawAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoMo3DlongOnWithdrawAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoMo3DlongOnWithdrawAndDistribute represents a OnWithdrawAndDistribute event raised by the FoMo3Dlong contract.
type FoMo3DlongOnWithdrawAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EthOut         *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnWithdrawAndDistribute is a free log retrieval operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: e onWithdrawAndDistribute(playerAddress address, playerName bytes32, ethOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) FilterOnWithdrawAndDistribute(opts *bind.FilterOpts) (*FoMo3DlongOnWithdrawAndDistributeIterator, error) {

	logs, sub, err := _FoMo3Dlong.contract.FilterLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return &FoMo3DlongOnWithdrawAndDistributeIterator{contract: _FoMo3Dlong.contract, event: "onWithdrawAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnWithdrawAndDistribute is a free log subscription operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: e onWithdrawAndDistribute(playerAddress address, playerName bytes32, ethOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_FoMo3Dlong *FoMo3DlongFilterer) WatchOnWithdrawAndDistribute(opts *bind.WatchOpts, sink chan<- *FoMo3DlongOnWithdrawAndDistribute) (event.Subscription, error) {

	logs, sub, err := _FoMo3Dlong.contract.WatchLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoMo3DlongOnWithdrawAndDistribute)
				if err := _FoMo3Dlong.contract.UnpackLog(event, "onWithdrawAndDistribute", log); err != nil {
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

// JIincForwarderInterfaceABI is the input ABI used to generate the binding from.
const JIincForwarderInterfaceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"cancelMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_firstCorpBank\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCorpBank\",\"type\":\"address\"}],\"name\":\"startMigration\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// JIincForwarderInterfaceBin is the compiled bytecode used for deploying new contracts.
const JIincForwarderInterfaceBin = `0x`

// DeployJIincForwarderInterface deploys a new Ethereum contract, binding an instance of JIincForwarderInterface to it.
func DeployJIincForwarderInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JIincForwarderInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(JIincForwarderInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JIincForwarderInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JIincForwarderInterface{JIincForwarderInterfaceCaller: JIincForwarderInterfaceCaller{contract: contract}, JIincForwarderInterfaceTransactor: JIincForwarderInterfaceTransactor{contract: contract}, JIincForwarderInterfaceFilterer: JIincForwarderInterfaceFilterer{contract: contract}}, nil
}

// JIincForwarderInterface is an auto generated Go binding around an Ethereum contract.
type JIincForwarderInterface struct {
	JIincForwarderInterfaceCaller     // Read-only binding to the contract
	JIincForwarderInterfaceTransactor // Write-only binding to the contract
	JIincForwarderInterfaceFilterer   // Log filterer for contract events
}

// JIincForwarderInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type JIincForwarderInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincForwarderInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JIincForwarderInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincForwarderInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JIincForwarderInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JIincForwarderInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JIincForwarderInterfaceSession struct {
	Contract     *JIincForwarderInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// JIincForwarderInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JIincForwarderInterfaceCallerSession struct {
	Contract *JIincForwarderInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// JIincForwarderInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JIincForwarderInterfaceTransactorSession struct {
	Contract     *JIincForwarderInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// JIincForwarderInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type JIincForwarderInterfaceRaw struct {
	Contract *JIincForwarderInterface // Generic contract binding to access the raw methods on
}

// JIincForwarderInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JIincForwarderInterfaceCallerRaw struct {
	Contract *JIincForwarderInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// JIincForwarderInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JIincForwarderInterfaceTransactorRaw struct {
	Contract *JIincForwarderInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJIincForwarderInterface creates a new instance of JIincForwarderInterface, bound to a specific deployed contract.
func NewJIincForwarderInterface(address common.Address, backend bind.ContractBackend) (*JIincForwarderInterface, error) {
	contract, err := bindJIincForwarderInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JIincForwarderInterface{JIincForwarderInterfaceCaller: JIincForwarderInterfaceCaller{contract: contract}, JIincForwarderInterfaceTransactor: JIincForwarderInterfaceTransactor{contract: contract}, JIincForwarderInterfaceFilterer: JIincForwarderInterfaceFilterer{contract: contract}}, nil
}

// NewJIincForwarderInterfaceCaller creates a new read-only instance of JIincForwarderInterface, bound to a specific deployed contract.
func NewJIincForwarderInterfaceCaller(address common.Address, caller bind.ContractCaller) (*JIincForwarderInterfaceCaller, error) {
	contract, err := bindJIincForwarderInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JIincForwarderInterfaceCaller{contract: contract}, nil
}

// NewJIincForwarderInterfaceTransactor creates a new write-only instance of JIincForwarderInterface, bound to a specific deployed contract.
func NewJIincForwarderInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*JIincForwarderInterfaceTransactor, error) {
	contract, err := bindJIincForwarderInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JIincForwarderInterfaceTransactor{contract: contract}, nil
}

// NewJIincForwarderInterfaceFilterer creates a new log filterer instance of JIincForwarderInterface, bound to a specific deployed contract.
func NewJIincForwarderInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*JIincForwarderInterfaceFilterer, error) {
	contract, err := bindJIincForwarderInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JIincForwarderInterfaceFilterer{contract: contract}, nil
}

// bindJIincForwarderInterface binds a generic wrapper to an already deployed contract.
func bindJIincForwarderInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JIincForwarderInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JIincForwarderInterface *JIincForwarderInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JIincForwarderInterface.Contract.JIincForwarderInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JIincForwarderInterface *JIincForwarderInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.JIincForwarderInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JIincForwarderInterface *JIincForwarderInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.JIincForwarderInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JIincForwarderInterface *JIincForwarderInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JIincForwarderInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.contract.Transact(opts, method, params...)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceCaller) Status(opts *bind.CallOpts) (common.Address, common.Address, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _JIincForwarderInterface.contract.Call(opts, out, "status")
	return *ret0, *ret1, *ret2, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceSession) Status() (common.Address, common.Address, bool, error) {
	return _JIincForwarderInterface.Contract.Status(&_JIincForwarderInterface.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(address, address, bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceCallerSession) Status() (common.Address, common.Address, bool, error) {
	return _JIincForwarderInterface.Contract.Status(&_JIincForwarderInterface.CallOpts)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactor) CancelMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarderInterface.contract.Transact(opts, "cancelMigration")
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceSession) CancelMigration() (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.CancelMigration(&_JIincForwarderInterface.TransactOpts)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactorSession) CancelMigration() (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.CancelMigration(&_JIincForwarderInterface.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarderInterface.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceSession) Deposit() (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.Deposit(&_JIincForwarderInterface.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactorSession) Deposit() (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.Deposit(&_JIincForwarderInterface.TransactOpts)
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactor) FinishMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JIincForwarderInterface.contract.Transact(opts, "finishMigration")
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceSession) FinishMigration() (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.FinishMigration(&_JIincForwarderInterface.TransactOpts)
}

// FinishMigration is a paid mutator transaction binding the contract method 0x88d761f2.
//
// Solidity: function finishMigration() returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactorSession) FinishMigration() (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.FinishMigration(&_JIincForwarderInterface.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactor) Setup(opts *bind.TransactOpts, _firstCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarderInterface.contract.Transact(opts, "setup", _firstCorpBank)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_JIincForwarderInterface *JIincForwarderInterfaceSession) Setup(_firstCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.Setup(&_JIincForwarderInterface.TransactOpts, _firstCorpBank)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(_firstCorpBank address) returns()
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactorSession) Setup(_firstCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.Setup(&_JIincForwarderInterface.TransactOpts, _firstCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactor) StartMigration(opts *bind.TransactOpts, _newCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarderInterface.contract.Transact(opts, "startMigration", _newCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceSession) StartMigration(_newCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.StartMigration(&_JIincForwarderInterface.TransactOpts, _newCorpBank)
}

// StartMigration is a paid mutator transaction binding the contract method 0xa0f52da0.
//
// Solidity: function startMigration(_newCorpBank address) returns(bool)
func (_JIincForwarderInterface *JIincForwarderInterfaceTransactorSession) StartMigration(_newCorpBank common.Address) (*types.Transaction, error) {
	return _JIincForwarderInterface.Contract.StartMigration(&_JIincForwarderInterface.TransactOpts, _newCorpBank)
}

// NameFilterABI is the input ABI used to generate the binding from.
const NameFilterABI = "[]"

// NameFilterBin is the compiled bytecode used for deploying new contracts.
const NameFilterBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206228e06a0123e3acdeb152888cc0da1d07638f1ccfbdf84b31d3e6b25de1f1ec0029`

// DeployNameFilter deploys a new Ethereum contract, binding an instance of NameFilter to it.
func DeployNameFilter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NameFilter, error) {
	parsed, err := abi.JSON(strings.NewReader(NameFilterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NameFilterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NameFilter{NameFilterCaller: NameFilterCaller{contract: contract}, NameFilterTransactor: NameFilterTransactor{contract: contract}, NameFilterFilterer: NameFilterFilterer{contract: contract}}, nil
}

// NameFilter is an auto generated Go binding around an Ethereum contract.
type NameFilter struct {
	NameFilterCaller     // Read-only binding to the contract
	NameFilterTransactor // Write-only binding to the contract
	NameFilterFilterer   // Log filterer for contract events
}

// NameFilterCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameFilterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameFilterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameFilterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameFilterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NameFilterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameFilterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameFilterSession struct {
	Contract     *NameFilter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameFilterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameFilterCallerSession struct {
	Contract *NameFilterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NameFilterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameFilterTransactorSession struct {
	Contract     *NameFilterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NameFilterRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameFilterRaw struct {
	Contract *NameFilter // Generic contract binding to access the raw methods on
}

// NameFilterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameFilterCallerRaw struct {
	Contract *NameFilterCaller // Generic read-only contract binding to access the raw methods on
}

// NameFilterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameFilterTransactorRaw struct {
	Contract *NameFilterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameFilter creates a new instance of NameFilter, bound to a specific deployed contract.
func NewNameFilter(address common.Address, backend bind.ContractBackend) (*NameFilter, error) {
	contract, err := bindNameFilter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameFilter{NameFilterCaller: NameFilterCaller{contract: contract}, NameFilterTransactor: NameFilterTransactor{contract: contract}, NameFilterFilterer: NameFilterFilterer{contract: contract}}, nil
}

// NewNameFilterCaller creates a new read-only instance of NameFilter, bound to a specific deployed contract.
func NewNameFilterCaller(address common.Address, caller bind.ContractCaller) (*NameFilterCaller, error) {
	contract, err := bindNameFilter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameFilterCaller{contract: contract}, nil
}

// NewNameFilterTransactor creates a new write-only instance of NameFilter, bound to a specific deployed contract.
func NewNameFilterTransactor(address common.Address, transactor bind.ContractTransactor) (*NameFilterTransactor, error) {
	contract, err := bindNameFilter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameFilterTransactor{contract: contract}, nil
}

// NewNameFilterFilterer creates a new log filterer instance of NameFilter, bound to a specific deployed contract.
func NewNameFilterFilterer(address common.Address, filterer bind.ContractFilterer) (*NameFilterFilterer, error) {
	contract, err := bindNameFilter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameFilterFilterer{contract: contract}, nil
}

// bindNameFilter binds a generic wrapper to an already deployed contract.
func bindNameFilter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameFilterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameFilter *NameFilterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameFilter.Contract.NameFilterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameFilter *NameFilterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameFilter.Contract.NameFilterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameFilter *NameFilterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameFilter.Contract.NameFilterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameFilter *NameFilterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameFilter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameFilter *NameFilterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameFilter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameFilter *NameFilterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameFilter.Contract.contract.Transact(opts, method, params...)
}

// PlayerBookInterfaceABI is the input ABI used to generate the binding from.
const PlayerBookInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getNameFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"bytes32\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXnameFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"address\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXaddrFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_affCode\",\"type\":\"uint256\"},{\"name\":\"_all\",\"type\":\"bool\"}],\"name\":\"registerNameXIDFromDapp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerLAff\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPlayerID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PlayerBookInterfaceBin is the compiled bytecode used for deploying new contracts.
const PlayerBookInterfaceBin = `0x`

// DeployPlayerBookInterface deploys a new Ethereum contract, binding an instance of PlayerBookInterface to it.
func DeployPlayerBookInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PlayerBookInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayerBookInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PlayerBookInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PlayerBookInterface{PlayerBookInterfaceCaller: PlayerBookInterfaceCaller{contract: contract}, PlayerBookInterfaceTransactor: PlayerBookInterfaceTransactor{contract: contract}, PlayerBookInterfaceFilterer: PlayerBookInterfaceFilterer{contract: contract}}, nil
}

// PlayerBookInterface is an auto generated Go binding around an Ethereum contract.
type PlayerBookInterface struct {
	PlayerBookInterfaceCaller     // Read-only binding to the contract
	PlayerBookInterfaceTransactor // Write-only binding to the contract
	PlayerBookInterfaceFilterer   // Log filterer for contract events
}

// PlayerBookInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlayerBookInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlayerBookInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlayerBookInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayerBookInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlayerBookInterfaceSession struct {
	Contract     *PlayerBookInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PlayerBookInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlayerBookInterfaceCallerSession struct {
	Contract *PlayerBookInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PlayerBookInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlayerBookInterfaceTransactorSession struct {
	Contract     *PlayerBookInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PlayerBookInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlayerBookInterfaceRaw struct {
	Contract *PlayerBookInterface // Generic contract binding to access the raw methods on
}

// PlayerBookInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlayerBookInterfaceCallerRaw struct {
	Contract *PlayerBookInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PlayerBookInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlayerBookInterfaceTransactorRaw struct {
	Contract *PlayerBookInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlayerBookInterface creates a new instance of PlayerBookInterface, bound to a specific deployed contract.
func NewPlayerBookInterface(address common.Address, backend bind.ContractBackend) (*PlayerBookInterface, error) {
	contract, err := bindPlayerBookInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlayerBookInterface{PlayerBookInterfaceCaller: PlayerBookInterfaceCaller{contract: contract}, PlayerBookInterfaceTransactor: PlayerBookInterfaceTransactor{contract: contract}, PlayerBookInterfaceFilterer: PlayerBookInterfaceFilterer{contract: contract}}, nil
}

// NewPlayerBookInterfaceCaller creates a new read-only instance of PlayerBookInterface, bound to a specific deployed contract.
func NewPlayerBookInterfaceCaller(address common.Address, caller bind.ContractCaller) (*PlayerBookInterfaceCaller, error) {
	contract, err := bindPlayerBookInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlayerBookInterfaceCaller{contract: contract}, nil
}

// NewPlayerBookInterfaceTransactor creates a new write-only instance of PlayerBookInterface, bound to a specific deployed contract.
func NewPlayerBookInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PlayerBookInterfaceTransactor, error) {
	contract, err := bindPlayerBookInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlayerBookInterfaceTransactor{contract: contract}, nil
}

// NewPlayerBookInterfaceFilterer creates a new log filterer instance of PlayerBookInterface, bound to a specific deployed contract.
func NewPlayerBookInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PlayerBookInterfaceFilterer, error) {
	contract, err := bindPlayerBookInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlayerBookInterfaceFilterer{contract: contract}, nil
}

// bindPlayerBookInterface binds a generic wrapper to an already deployed contract.
func bindPlayerBookInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayerBookInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlayerBookInterface *PlayerBookInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlayerBookInterface.Contract.PlayerBookInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlayerBookInterface *PlayerBookInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.PlayerBookInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlayerBookInterface *PlayerBookInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.PlayerBookInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlayerBookInterface *PlayerBookInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlayerBookInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlayerBookInterface *PlayerBookInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlayerBookInterface *PlayerBookInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.contract.Transact(opts, method, params...)
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceCaller) GetNameFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBookInterface.contract.Call(opts, out, "getNameFee")
	return *ret0, err
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetNameFee() (*big.Int, error) {
	return _PlayerBookInterface.Contract.GetNameFee(&_PlayerBookInterface.CallOpts)
}

// GetNameFee is a free data retrieval call binding the contract method 0x2614195f.
//
// Solidity: function getNameFee() constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceCallerSession) GetNameFee() (*big.Int, error) {
	return _PlayerBookInterface.Contract.GetNameFee(&_PlayerBookInterface.CallOpts)
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBookInterface *PlayerBookInterfaceCaller) GetPlayerAddr(opts *bind.CallOpts, _pID *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlayerBookInterface.contract.Call(opts, out, "getPlayerAddr", _pID)
	return *ret0, err
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetPlayerAddr(_pID *big.Int) (common.Address, error) {
	return _PlayerBookInterface.Contract.GetPlayerAddr(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerAddr is a free data retrieval call binding the contract method 0x4d0d35ff.
//
// Solidity: function getPlayerAddr(_pID uint256) constant returns(address)
func (_PlayerBookInterface *PlayerBookInterfaceCallerSession) GetPlayerAddr(_pID *big.Int) (common.Address, error) {
	return _PlayerBookInterface.Contract.GetPlayerAddr(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceCaller) GetPlayerLAff(opts *bind.CallOpts, _pID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlayerBookInterface.contract.Call(opts, out, "getPlayerLAff", _pID)
	return *ret0, err
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetPlayerLAff(_pID *big.Int) (*big.Int, error) {
	return _PlayerBookInterface.Contract.GetPlayerLAff(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerLAff is a free data retrieval call binding the contract method 0xe3c08adf.
//
// Solidity: function getPlayerLAff(_pID uint256) constant returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceCallerSession) GetPlayerLAff(_pID *big.Int) (*big.Int, error) {
	return _PlayerBookInterface.Contract.GetPlayerLAff(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBookInterface *PlayerBookInterfaceCaller) GetPlayerName(opts *bind.CallOpts, _pID *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PlayerBookInterface.contract.Call(opts, out, "getPlayerName", _pID)
	return *ret0, err
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetPlayerName(_pID *big.Int) ([32]byte, error) {
	return _PlayerBookInterface.Contract.GetPlayerName(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerName is a free data retrieval call binding the contract method 0x82e37b2c.
//
// Solidity: function getPlayerName(_pID uint256) constant returns(bytes32)
func (_PlayerBookInterface *PlayerBookInterfaceCallerSession) GetPlayerName(_pID *big.Int) ([32]byte, error) {
	return _PlayerBookInterface.Contract.GetPlayerName(&_PlayerBookInterface.CallOpts, _pID)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactor) GetPlayerID(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _PlayerBookInterface.contract.Transact(opts, "getPlayerID", _addr)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) GetPlayerID(_addr common.Address) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.GetPlayerID(&_PlayerBookInterface.TransactOpts, _addr)
}

// GetPlayerID is a paid mutator transaction binding the contract method 0xe56556a9.
//
// Solidity: function getPlayerID(_addr address) returns(uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactorSession) GetPlayerID(_addr common.Address) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.GetPlayerID(&_PlayerBookInterface.TransactOpts, _addr)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactor) RegisterNameXIDFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.contract.Transact(opts, "registerNameXIDFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) RegisterNameXIDFromDapp(_addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXIDFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXIDFromDapp is a paid mutator transaction binding the contract method 0xc0942dfd.
//
// Solidity: function registerNameXIDFromDapp(_addr address, _name bytes32, _affCode uint256, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactorSession) RegisterNameXIDFromDapp(_addr common.Address, _name [32]byte, _affCode *big.Int, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXIDFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactor) RegisterNameXaddrFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.contract.Transact(opts, "registerNameXaddrFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) RegisterNameXaddrFromDapp(_addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXaddrFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXaddrFromDapp is a paid mutator transaction binding the contract method 0xaa4d490b.
//
// Solidity: function registerNameXaddrFromDapp(_addr address, _name bytes32, _affCode address, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactorSession) RegisterNameXaddrFromDapp(_addr common.Address, _name [32]byte, _affCode common.Address, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXaddrFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactor) RegisterNameXnameFromDapp(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.contract.Transact(opts, "registerNameXnameFromDapp", _addr, _name, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceSession) RegisterNameXnameFromDapp(_addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXnameFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// RegisterNameXnameFromDapp is a paid mutator transaction binding the contract method 0x745ea0c1.
//
// Solidity: function registerNameXnameFromDapp(_addr address, _name bytes32, _affCode bytes32, _all bool) returns(bool, uint256)
func (_PlayerBookInterface *PlayerBookInterfaceTransactorSession) RegisterNameXnameFromDapp(_addr common.Address, _name [32]byte, _affCode [32]byte, _all bool) (*types.Transaction, error) {
	return _PlayerBookInterface.Contract.RegisterNameXnameFromDapp(&_PlayerBookInterface.TransactOpts, _addr, _name, _affCode, _all)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e3da4033667851c2cde6084b45114e732ff794e19f06fc997a80a3acd2332ad70029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ModularLongABI is the input ABI used to generate the binding from.
const ModularLongABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isNewPlayer\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onNewName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"keysBought\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"potAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"airDropPot\",\"type\":\"uint256\"}],\"name\":\"onEndTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethOut\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onWithdrawAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onReLoadAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"buyerID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onAffiliatePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amountAddedToPot\",\"type\":\"uint256\"}],\"name\":\"onPotSwapDeposit\",\"type\":\"event\"}]"

// ModularLongBin is the compiled bytecode used for deploying new contracts.
const ModularLongBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820bda9e6fb3be31d8b3f17f10e0ab67b324973ad615cec6609bd45facd7875fdd60029`

// DeployModularLong deploys a new Ethereum contract, binding an instance of ModularLong to it.
func DeployModularLong(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModularLong, error) {
	parsed, err := abi.JSON(strings.NewReader(ModularLongABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ModularLongBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModularLong{ModularLongCaller: ModularLongCaller{contract: contract}, ModularLongTransactor: ModularLongTransactor{contract: contract}, ModularLongFilterer: ModularLongFilterer{contract: contract}}, nil
}

// ModularLong is an auto generated Go binding around an Ethereum contract.
type ModularLong struct {
	ModularLongCaller     // Read-only binding to the contract
	ModularLongTransactor // Write-only binding to the contract
	ModularLongFilterer   // Log filterer for contract events
}

// ModularLongCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModularLongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularLongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModularLongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularLongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModularLongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularLongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModularLongSession struct {
	Contract     *ModularLong      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModularLongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModularLongCallerSession struct {
	Contract *ModularLongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ModularLongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModularLongTransactorSession struct {
	Contract     *ModularLongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ModularLongRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModularLongRaw struct {
	Contract *ModularLong // Generic contract binding to access the raw methods on
}

// ModularLongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModularLongCallerRaw struct {
	Contract *ModularLongCaller // Generic read-only contract binding to access the raw methods on
}

// ModularLongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModularLongTransactorRaw struct {
	Contract *ModularLongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModularLong creates a new instance of ModularLong, bound to a specific deployed contract.
func NewModularLong(address common.Address, backend bind.ContractBackend) (*ModularLong, error) {
	contract, err := bindModularLong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModularLong{ModularLongCaller: ModularLongCaller{contract: contract}, ModularLongTransactor: ModularLongTransactor{contract: contract}, ModularLongFilterer: ModularLongFilterer{contract: contract}}, nil
}

// NewModularLongCaller creates a new read-only instance of ModularLong, bound to a specific deployed contract.
func NewModularLongCaller(address common.Address, caller bind.ContractCaller) (*ModularLongCaller, error) {
	contract, err := bindModularLong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModularLongCaller{contract: contract}, nil
}

// NewModularLongTransactor creates a new write-only instance of ModularLong, bound to a specific deployed contract.
func NewModularLongTransactor(address common.Address, transactor bind.ContractTransactor) (*ModularLongTransactor, error) {
	contract, err := bindModularLong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModularLongTransactor{contract: contract}, nil
}

// NewModularLongFilterer creates a new log filterer instance of ModularLong, bound to a specific deployed contract.
func NewModularLongFilterer(address common.Address, filterer bind.ContractFilterer) (*ModularLongFilterer, error) {
	contract, err := bindModularLong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModularLongFilterer{contract: contract}, nil
}

// bindModularLong binds a generic wrapper to an already deployed contract.
func bindModularLong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModularLongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModularLong *ModularLongRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ModularLong.Contract.ModularLongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModularLong *ModularLongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModularLong.Contract.ModularLongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModularLong *ModularLongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModularLong.Contract.ModularLongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModularLong *ModularLongCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ModularLong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModularLong *ModularLongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModularLong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModularLong *ModularLongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModularLong.Contract.contract.Transact(opts, method, params...)
}

// ModularLongOnAffiliatePayoutIterator is returned from FilterOnAffiliatePayout and is used to iterate over the raw logs and unpacked data for OnAffiliatePayout events raised by the ModularLong contract.
type ModularLongOnAffiliatePayoutIterator struct {
	Event *ModularLongOnAffiliatePayout // Event containing the contract specifics and raw log

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
func (it *ModularLongOnAffiliatePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnAffiliatePayout)
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
		it.Event = new(ModularLongOnAffiliatePayout)
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
func (it *ModularLongOnAffiliatePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnAffiliatePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnAffiliatePayout represents a OnAffiliatePayout event raised by the ModularLong contract.
type ModularLongOnAffiliatePayout struct {
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	RoundID          *big.Int
	BuyerID          *big.Int
	Amount           *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnAffiliatePayout is a free log retrieval operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: e onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) FilterOnAffiliatePayout(opts *bind.FilterOpts, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (*ModularLongOnAffiliatePayoutIterator, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return &ModularLongOnAffiliatePayoutIterator{contract: _ModularLong.contract, event: "onAffiliatePayout", logs: logs, sub: sub}, nil
}

// WatchOnAffiliatePayout is a free log subscription operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: e onAffiliatePayout(affiliateID indexed uint256, affiliateAddress address, affiliateName bytes32, roundID indexed uint256, buyerID indexed uint256, amount uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) WatchOnAffiliatePayout(opts *bind.WatchOpts, sink chan<- *ModularLongOnAffiliatePayout, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (event.Subscription, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnAffiliatePayout)
				if err := _ModularLong.contract.UnpackLog(event, "onAffiliatePayout", log); err != nil {
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

// ModularLongOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the ModularLong contract.
type ModularLongOnBuyAndDistributeIterator struct {
	Event *ModularLongOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *ModularLongOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnBuyAndDistribute)
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
		it.Event = new(ModularLongOnBuyAndDistribute)
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
func (it *ModularLongOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the ModularLong contract.
type ModularLongOnBuyAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EthIn          *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: e onBuyAndDistribute(playerAddress address, playerName bytes32, ethIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*ModularLongOnBuyAndDistributeIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnBuyAndDistributeIterator{contract: _ModularLong.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: e onBuyAndDistribute(playerAddress address, playerName bytes32, ethIn uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *ModularLongOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnBuyAndDistribute)
				if err := _ModularLong.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// ModularLongOnEndTxIterator is returned from FilterOnEndTx and is used to iterate over the raw logs and unpacked data for OnEndTx events raised by the ModularLong contract.
type ModularLongOnEndTxIterator struct {
	Event *ModularLongOnEndTx // Event containing the contract specifics and raw log

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
func (it *ModularLongOnEndTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnEndTx)
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
		it.Event = new(ModularLongOnEndTx)
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
func (it *ModularLongOnEndTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnEndTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnEndTx represents a OnEndTx event raised by the ModularLong contract.
type ModularLongOnEndTx struct {
	CompressedData *big.Int
	CompressedIDs  *big.Int
	PlayerName     [32]byte
	PlayerAddress  common.Address
	EthIn          *big.Int
	KeysBought     *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	PotAmount      *big.Int
	AirDropPot     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnEndTx is a free log retrieval operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: e onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, ethIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_ModularLong *ModularLongFilterer) FilterOnEndTx(opts *bind.FilterOpts) (*ModularLongOnEndTxIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnEndTxIterator{contract: _ModularLong.contract, event: "onEndTx", logs: logs, sub: sub}, nil
}

// WatchOnEndTx is a free log subscription operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: e onEndTx(compressedData uint256, compressedIDs uint256, playerName bytes32, playerAddress address, ethIn uint256, keysBought uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256, potAmount uint256, airDropPot uint256)
func (_ModularLong *ModularLongFilterer) WatchOnEndTx(opts *bind.WatchOpts, sink chan<- *ModularLongOnEndTx) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnEndTx)
				if err := _ModularLong.contract.UnpackLog(event, "onEndTx", log); err != nil {
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

// ModularLongOnNewNameIterator is returned from FilterOnNewName and is used to iterate over the raw logs and unpacked data for OnNewName events raised by the ModularLong contract.
type ModularLongOnNewNameIterator struct {
	Event *ModularLongOnNewName // Event containing the contract specifics and raw log

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
func (it *ModularLongOnNewNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnNewName)
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
		it.Event = new(ModularLongOnNewName)
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
func (it *ModularLongOnNewNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnNewNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnNewName represents a OnNewName event raised by the ModularLong contract.
type ModularLongOnNewName struct {
	PlayerID         *big.Int
	PlayerAddress    common.Address
	PlayerName       [32]byte
	IsNewPlayer      bool
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	AmountPaid       *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnNewName is a free log retrieval operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: e onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) FilterOnNewName(opts *bind.FilterOpts, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (*ModularLongOnNewNameIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return &ModularLongOnNewNameIterator{contract: _ModularLong.contract, event: "onNewName", logs: logs, sub: sub}, nil
}

// WatchOnNewName is a free log subscription operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: e onNewName(playerID indexed uint256, playerAddress indexed address, playerName indexed bytes32, isNewPlayer bool, affiliateID uint256, affiliateAddress address, affiliateName bytes32, amountPaid uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) WatchOnNewName(opts *bind.WatchOpts, sink chan<- *ModularLongOnNewName, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnNewName)
				if err := _ModularLong.contract.UnpackLog(event, "onNewName", log); err != nil {
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

// ModularLongOnPotSwapDepositIterator is returned from FilterOnPotSwapDeposit and is used to iterate over the raw logs and unpacked data for OnPotSwapDeposit events raised by the ModularLong contract.
type ModularLongOnPotSwapDepositIterator struct {
	Event *ModularLongOnPotSwapDeposit // Event containing the contract specifics and raw log

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
func (it *ModularLongOnPotSwapDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnPotSwapDeposit)
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
		it.Event = new(ModularLongOnPotSwapDeposit)
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
func (it *ModularLongOnPotSwapDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnPotSwapDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnPotSwapDeposit represents a OnPotSwapDeposit event raised by the ModularLong contract.
type ModularLongOnPotSwapDeposit struct {
	RoundID          *big.Int
	AmountAddedToPot *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnPotSwapDeposit is a free log retrieval operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: e onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_ModularLong *ModularLongFilterer) FilterOnPotSwapDeposit(opts *bind.FilterOpts) (*ModularLongOnPotSwapDepositIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnPotSwapDepositIterator{contract: _ModularLong.contract, event: "onPotSwapDeposit", logs: logs, sub: sub}, nil
}

// WatchOnPotSwapDeposit is a free log subscription operation binding the contract event 0x74b1d2f771e0eff1b2c36c38499febdbea80fe4013bdace4fc4b653322c2895c.
//
// Solidity: e onPotSwapDeposit(roundID uint256, amountAddedToPot uint256)
func (_ModularLong *ModularLongFilterer) WatchOnPotSwapDeposit(opts *bind.WatchOpts, sink chan<- *ModularLongOnPotSwapDeposit) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onPotSwapDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnPotSwapDeposit)
				if err := _ModularLong.contract.UnpackLog(event, "onPotSwapDeposit", log); err != nil {
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

// ModularLongOnReLoadAndDistributeIterator is returned from FilterOnReLoadAndDistribute and is used to iterate over the raw logs and unpacked data for OnReLoadAndDistribute events raised by the ModularLong contract.
type ModularLongOnReLoadAndDistributeIterator struct {
	Event *ModularLongOnReLoadAndDistribute // Event containing the contract specifics and raw log

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
func (it *ModularLongOnReLoadAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnReLoadAndDistribute)
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
		it.Event = new(ModularLongOnReLoadAndDistribute)
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
func (it *ModularLongOnReLoadAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnReLoadAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnReLoadAndDistribute represents a OnReLoadAndDistribute event raised by the ModularLong contract.
type ModularLongOnReLoadAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnReLoadAndDistribute is a free log retrieval operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: e onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) FilterOnReLoadAndDistribute(opts *bind.FilterOpts) (*ModularLongOnReLoadAndDistributeIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnReLoadAndDistributeIterator{contract: _ModularLong.contract, event: "onReLoadAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnReLoadAndDistribute is a free log subscription operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: e onReLoadAndDistribute(playerAddress address, playerName bytes32, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) WatchOnReLoadAndDistribute(opts *bind.WatchOpts, sink chan<- *ModularLongOnReLoadAndDistribute) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnReLoadAndDistribute)
				if err := _ModularLong.contract.UnpackLog(event, "onReLoadAndDistribute", log); err != nil {
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

// ModularLongOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the ModularLong contract.
type ModularLongOnWithdrawIterator struct {
	Event *ModularLongOnWithdraw // Event containing the contract specifics and raw log

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
func (it *ModularLongOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnWithdraw)
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
		it.Event = new(ModularLongOnWithdraw)
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
func (it *ModularLongOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnWithdraw represents a OnWithdraw event raised by the ModularLong contract.
type ModularLongOnWithdraw struct {
	PlayerID      *big.Int
	PlayerAddress common.Address
	PlayerName    [32]byte
	EthOut        *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: e onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, ethOut uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) FilterOnWithdraw(opts *bind.FilterOpts, playerID []*big.Int) (*ModularLongOnWithdrawIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return &ModularLongOnWithdrawIterator{contract: _ModularLong.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: e onWithdraw(playerID indexed uint256, playerAddress address, playerName bytes32, ethOut uint256, timeStamp uint256)
func (_ModularLong *ModularLongFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *ModularLongOnWithdraw, playerID []*big.Int) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnWithdraw)
				if err := _ModularLong.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// ModularLongOnWithdrawAndDistributeIterator is returned from FilterOnWithdrawAndDistribute and is used to iterate over the raw logs and unpacked data for OnWithdrawAndDistribute events raised by the ModularLong contract.
type ModularLongOnWithdrawAndDistributeIterator struct {
	Event *ModularLongOnWithdrawAndDistribute // Event containing the contract specifics and raw log

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
func (it *ModularLongOnWithdrawAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularLongOnWithdrawAndDistribute)
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
		it.Event = new(ModularLongOnWithdrawAndDistribute)
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
func (it *ModularLongOnWithdrawAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularLongOnWithdrawAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularLongOnWithdrawAndDistribute represents a OnWithdrawAndDistribute event raised by the ModularLong contract.
type ModularLongOnWithdrawAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EthOut         *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnWithdrawAndDistribute is a free log retrieval operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: e onWithdrawAndDistribute(playerAddress address, playerName bytes32, ethOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) FilterOnWithdrawAndDistribute(opts *bind.FilterOpts) (*ModularLongOnWithdrawAndDistributeIterator, error) {

	logs, sub, err := _ModularLong.contract.FilterLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return &ModularLongOnWithdrawAndDistributeIterator{contract: _ModularLong.contract, event: "onWithdrawAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnWithdrawAndDistribute is a free log subscription operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: e onWithdrawAndDistribute(playerAddress address, playerName bytes32, ethOut uint256, compressedData uint256, compressedIDs uint256, winnerAddr address, winnerName bytes32, amountWon uint256, newPot uint256, P3DAmount uint256, genAmount uint256)
func (_ModularLong *ModularLongFilterer) WatchOnWithdrawAndDistribute(opts *bind.WatchOpts, sink chan<- *ModularLongOnWithdrawAndDistribute) (event.Subscription, error) {

	logs, sub, err := _ModularLong.contract.WatchLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularLongOnWithdrawAndDistribute)
				if err := _ModularLong.contract.UnpackLog(event, "onWithdrawAndDistribute", log); err != nil {
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

// OtherFoMo3DABI is the input ABI used to generate the binding from.
const OtherFoMo3DABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"potSwap\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// OtherFoMo3DBin is the compiled bytecode used for deploying new contracts.
const OtherFoMo3DBin = `0x`

// DeployOtherFoMo3D deploys a new Ethereum contract, binding an instance of OtherFoMo3D to it.
func DeployOtherFoMo3D(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OtherFoMo3D, error) {
	parsed, err := abi.JSON(strings.NewReader(OtherFoMo3DABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OtherFoMo3DBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OtherFoMo3D{OtherFoMo3DCaller: OtherFoMo3DCaller{contract: contract}, OtherFoMo3DTransactor: OtherFoMo3DTransactor{contract: contract}, OtherFoMo3DFilterer: OtherFoMo3DFilterer{contract: contract}}, nil
}

// OtherFoMo3D is an auto generated Go binding around an Ethereum contract.
type OtherFoMo3D struct {
	OtherFoMo3DCaller     // Read-only binding to the contract
	OtherFoMo3DTransactor // Write-only binding to the contract
	OtherFoMo3DFilterer   // Log filterer for contract events
}

// OtherFoMo3DCaller is an auto generated read-only Go binding around an Ethereum contract.
type OtherFoMo3DCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherFoMo3DTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OtherFoMo3DTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherFoMo3DFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OtherFoMo3DFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtherFoMo3DSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OtherFoMo3DSession struct {
	Contract     *OtherFoMo3D      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtherFoMo3DCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OtherFoMo3DCallerSession struct {
	Contract *OtherFoMo3DCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OtherFoMo3DTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OtherFoMo3DTransactorSession struct {
	Contract     *OtherFoMo3DTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OtherFoMo3DRaw is an auto generated low-level Go binding around an Ethereum contract.
type OtherFoMo3DRaw struct {
	Contract *OtherFoMo3D // Generic contract binding to access the raw methods on
}

// OtherFoMo3DCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OtherFoMo3DCallerRaw struct {
	Contract *OtherFoMo3DCaller // Generic read-only contract binding to access the raw methods on
}

// OtherFoMo3DTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OtherFoMo3DTransactorRaw struct {
	Contract *OtherFoMo3DTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOtherFoMo3D creates a new instance of OtherFoMo3D, bound to a specific deployed contract.
func NewOtherFoMo3D(address common.Address, backend bind.ContractBackend) (*OtherFoMo3D, error) {
	contract, err := bindOtherFoMo3D(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OtherFoMo3D{OtherFoMo3DCaller: OtherFoMo3DCaller{contract: contract}, OtherFoMo3DTransactor: OtherFoMo3DTransactor{contract: contract}, OtherFoMo3DFilterer: OtherFoMo3DFilterer{contract: contract}}, nil
}

// NewOtherFoMo3DCaller creates a new read-only instance of OtherFoMo3D, bound to a specific deployed contract.
func NewOtherFoMo3DCaller(address common.Address, caller bind.ContractCaller) (*OtherFoMo3DCaller, error) {
	contract, err := bindOtherFoMo3D(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OtherFoMo3DCaller{contract: contract}, nil
}

// NewOtherFoMo3DTransactor creates a new write-only instance of OtherFoMo3D, bound to a specific deployed contract.
func NewOtherFoMo3DTransactor(address common.Address, transactor bind.ContractTransactor) (*OtherFoMo3DTransactor, error) {
	contract, err := bindOtherFoMo3D(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OtherFoMo3DTransactor{contract: contract}, nil
}

// NewOtherFoMo3DFilterer creates a new log filterer instance of OtherFoMo3D, bound to a specific deployed contract.
func NewOtherFoMo3DFilterer(address common.Address, filterer bind.ContractFilterer) (*OtherFoMo3DFilterer, error) {
	contract, err := bindOtherFoMo3D(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OtherFoMo3DFilterer{contract: contract}, nil
}

// bindOtherFoMo3D binds a generic wrapper to an already deployed contract.
func bindOtherFoMo3D(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OtherFoMo3DABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtherFoMo3D *OtherFoMo3DRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OtherFoMo3D.Contract.OtherFoMo3DCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtherFoMo3D *OtherFoMo3DRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtherFoMo3D.Contract.OtherFoMo3DTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtherFoMo3D *OtherFoMo3DRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtherFoMo3D.Contract.OtherFoMo3DTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtherFoMo3D *OtherFoMo3DCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OtherFoMo3D.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtherFoMo3D *OtherFoMo3DTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtherFoMo3D.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtherFoMo3D *OtherFoMo3DTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtherFoMo3D.Contract.contract.Transact(opts, method, params...)
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_OtherFoMo3D *OtherFoMo3DTransactor) PotSwap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtherFoMo3D.contract.Transact(opts, "potSwap")
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_OtherFoMo3D *OtherFoMo3DSession) PotSwap() (*types.Transaction, error) {
	return _OtherFoMo3D.Contract.PotSwap(&_OtherFoMo3D.TransactOpts)
}

// PotSwap is a paid mutator transaction binding the contract method 0xed78cf4a.
//
// Solidity: function potSwap() returns()
func (_OtherFoMo3D *OtherFoMo3DTransactorSession) PotSwap() (*types.Transaction, error) {
	return _OtherFoMo3D.Contract.PotSwap(&_OtherFoMo3D.TransactOpts)
}
