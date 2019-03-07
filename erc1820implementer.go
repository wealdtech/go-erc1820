// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1820

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Erc1820ImplementerABI is the input ABI used to generate the binding from.
const Erc1820ImplementerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceHash\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"canImplementInterfaceForAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Erc1820Implementer is an auto generated Go binding around an Ethereum contract.
type Erc1820Implementer struct {
	Erc1820ImplementerCaller     // Read-only binding to the contract
	Erc1820ImplementerTransactor // Write-only binding to the contract
	Erc1820ImplementerFilterer   // Log filterer for contract events
}

// Erc1820ImplementerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc1820ImplementerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820ImplementerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc1820ImplementerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820ImplementerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc1820ImplementerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820ImplementerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc1820ImplementerSession struct {
	Contract     *Erc1820Implementer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Erc1820ImplementerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc1820ImplementerCallerSession struct {
	Contract *Erc1820ImplementerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Erc1820ImplementerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc1820ImplementerTransactorSession struct {
	Contract     *Erc1820ImplementerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Erc1820ImplementerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc1820ImplementerRaw struct {
	Contract *Erc1820Implementer // Generic contract binding to access the raw methods on
}

// Erc1820ImplementerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc1820ImplementerCallerRaw struct {
	Contract *Erc1820ImplementerCaller // Generic read-only contract binding to access the raw methods on
}

// Erc1820ImplementerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc1820ImplementerTransactorRaw struct {
	Contract *Erc1820ImplementerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc1820Implementer creates a new instance of Erc1820Implementer, bound to a specific deployed contract.
func NewErc1820Implementer(address common.Address, backend bind.ContractBackend) (*Erc1820Implementer, error) {
	contract, err := bindErc1820Implementer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc1820Implementer{Erc1820ImplementerCaller: Erc1820ImplementerCaller{contract: contract}, Erc1820ImplementerTransactor: Erc1820ImplementerTransactor{contract: contract}, Erc1820ImplementerFilterer: Erc1820ImplementerFilterer{contract: contract}}, nil
}

// NewErc1820ImplementerCaller creates a new read-only instance of Erc1820Implementer, bound to a specific deployed contract.
func NewErc1820ImplementerCaller(address common.Address, caller bind.ContractCaller) (*Erc1820ImplementerCaller, error) {
	contract, err := bindErc1820Implementer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1820ImplementerCaller{contract: contract}, nil
}

// NewErc1820ImplementerTransactor creates a new write-only instance of Erc1820Implementer, bound to a specific deployed contract.
func NewErc1820ImplementerTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc1820ImplementerTransactor, error) {
	contract, err := bindErc1820Implementer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1820ImplementerTransactor{contract: contract}, nil
}

// NewErc1820ImplementerFilterer creates a new log filterer instance of Erc1820Implementer, bound to a specific deployed contract.
func NewErc1820ImplementerFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc1820ImplementerFilterer, error) {
	contract, err := bindErc1820Implementer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc1820ImplementerFilterer{contract: contract}, nil
}

// bindErc1820Implementer binds a generic wrapper to an already deployed contract.
func bindErc1820Implementer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1820ImplementerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1820Implementer *Erc1820ImplementerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc1820Implementer.Contract.Erc1820ImplementerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1820Implementer *Erc1820ImplementerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1820Implementer.Contract.Erc1820ImplementerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1820Implementer *Erc1820ImplementerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1820Implementer.Contract.Erc1820ImplementerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1820Implementer *Erc1820ImplementerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc1820Implementer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1820Implementer *Erc1820ImplementerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1820Implementer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1820Implementer *Erc1820ImplementerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1820Implementer.Contract.contract.Transact(opts, method, params...)
}

// CanImplementInterfaceForAddress is a free data retrieval call binding the contract method 0x249cb3fa.
//
// Solidity: function canImplementInterfaceForAddress(bytes32 interfaceHash, address addr) constant returns(bytes32)
func (_Erc1820Implementer *Erc1820ImplementerCaller) CanImplementInterfaceForAddress(opts *bind.CallOpts, interfaceHash [32]byte, addr common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Erc1820Implementer.contract.Call(opts, out, "canImplementInterfaceForAddress", interfaceHash, addr)
	return *ret0, err
}

// CanImplementInterfaceForAddress is a free data retrieval call binding the contract method 0x249cb3fa.
//
// Solidity: function canImplementInterfaceForAddress(bytes32 interfaceHash, address addr) constant returns(bytes32)
func (_Erc1820Implementer *Erc1820ImplementerSession) CanImplementInterfaceForAddress(interfaceHash [32]byte, addr common.Address) ([32]byte, error) {
	return _Erc1820Implementer.Contract.CanImplementInterfaceForAddress(&_Erc1820Implementer.CallOpts, interfaceHash, addr)
}

// CanImplementInterfaceForAddress is a free data retrieval call binding the contract method 0x249cb3fa.
//
// Solidity: function canImplementInterfaceForAddress(bytes32 interfaceHash, address addr) constant returns(bytes32)
func (_Erc1820Implementer *Erc1820ImplementerCallerSession) CanImplementInterfaceForAddress(interfaceHash [32]byte, addr common.Address) ([32]byte, error) {
	return _Erc1820Implementer.Contract.CanImplementInterfaceForAddress(&_Erc1820Implementer.CallOpts, interfaceHash, addr)
}
