// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Erc1820RegistryABI is the input ABI used to generate the binding from.
const Erc1820RegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_interfaceHash\",\"type\":\"bytes32\"},{\"name\":\"_implementer\",\"type\":\"address\"}],\"name\":\"setInterfaceImplementer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceName\",\"type\":\"string\"}],\"name\":\"interfaceHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"},{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"updateERC165Cache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_interfaceHash\",\"type\":\"bytes32\"}],\"name\":\"getInterfaceImplementer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"},{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"implementsERC165InterfaceNoCache\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"},{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"implementsERC165Interface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"interfaceHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceImplementerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"ManagerChanged\",\"type\":\"event\"}]"

// Erc1820Registry is an auto generated Go binding around an Ethereum contract.
type Erc1820Registry struct {
	Erc1820RegistryCaller     // Read-only binding to the contract
	Erc1820RegistryTransactor // Write-only binding to the contract
	Erc1820RegistryFilterer   // Log filterer for contract events
}

// Erc1820RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc1820RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc1820RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc1820RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc1820RegistrySession struct {
	Contract     *Erc1820Registry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc1820RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc1820RegistryCallerSession struct {
	Contract *Erc1820RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Erc1820RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc1820RegistryTransactorSession struct {
	Contract     *Erc1820RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Erc1820RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc1820RegistryRaw struct {
	Contract *Erc1820Registry // Generic contract binding to access the raw methods on
}

// Erc1820RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc1820RegistryCallerRaw struct {
	Contract *Erc1820RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// Erc1820RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc1820RegistryTransactorRaw struct {
	Contract *Erc1820RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc1820Registry creates a new instance of Erc1820Registry, bound to a specific deployed contract.
func NewErc1820Registry(address common.Address, backend bind.ContractBackend) (*Erc1820Registry, error) {
	contract, err := bindErc1820Registry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc1820Registry{Erc1820RegistryCaller: Erc1820RegistryCaller{contract: contract}, Erc1820RegistryTransactor: Erc1820RegistryTransactor{contract: contract}, Erc1820RegistryFilterer: Erc1820RegistryFilterer{contract: contract}}, nil
}

// NewErc1820RegistryCaller creates a new read-only instance of Erc1820Registry, bound to a specific deployed contract.
func NewErc1820RegistryCaller(address common.Address, caller bind.ContractCaller) (*Erc1820RegistryCaller, error) {
	contract, err := bindErc1820Registry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1820RegistryCaller{contract: contract}, nil
}

// NewErc1820RegistryTransactor creates a new write-only instance of Erc1820Registry, bound to a specific deployed contract.
func NewErc1820RegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc1820RegistryTransactor, error) {
	contract, err := bindErc1820Registry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1820RegistryTransactor{contract: contract}, nil
}

// NewErc1820RegistryFilterer creates a new log filterer instance of Erc1820Registry, bound to a specific deployed contract.
func NewErc1820RegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc1820RegistryFilterer, error) {
	contract, err := bindErc1820Registry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc1820RegistryFilterer{contract: contract}, nil
}

// bindErc1820Registry binds a generic wrapper to an already deployed contract.
func bindErc1820Registry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1820RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1820Registry *Erc1820RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc1820Registry.Contract.Erc1820RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1820Registry *Erc1820RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.Erc1820RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1820Registry *Erc1820RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.Erc1820RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1820Registry *Erc1820RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc1820Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1820Registry *Erc1820RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1820Registry *Erc1820RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.contract.Transact(opts, method, params...)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address _addr, bytes32 _interfaceHash) constant returns(address)
func (_Erc1820Registry *Erc1820RegistryCaller) GetInterfaceImplementer(opts *bind.CallOpts, _addr common.Address, _interfaceHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Erc1820Registry.contract.Call(opts, out, "getInterfaceImplementer", _addr, _interfaceHash)
	return *ret0, err
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address _addr, bytes32 _interfaceHash) constant returns(address)
func (_Erc1820Registry *Erc1820RegistrySession) GetInterfaceImplementer(_addr common.Address, _interfaceHash [32]byte) (common.Address, error) {
	return _Erc1820Registry.Contract.GetInterfaceImplementer(&_Erc1820Registry.CallOpts, _addr, _interfaceHash)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address _addr, bytes32 _interfaceHash) constant returns(address)
func (_Erc1820Registry *Erc1820RegistryCallerSession) GetInterfaceImplementer(_addr common.Address, _interfaceHash [32]byte) (common.Address, error) {
	return _Erc1820Registry.Contract.GetInterfaceImplementer(&_Erc1820Registry.CallOpts, _addr, _interfaceHash)
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address _addr) constant returns(address)
func (_Erc1820Registry *Erc1820RegistryCaller) GetManager(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Erc1820Registry.contract.Call(opts, out, "getManager", _addr)
	return *ret0, err
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address _addr) constant returns(address)
func (_Erc1820Registry *Erc1820RegistrySession) GetManager(_addr common.Address) (common.Address, error) {
	return _Erc1820Registry.Contract.GetManager(&_Erc1820Registry.CallOpts, _addr)
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address _addr) constant returns(address)
func (_Erc1820Registry *Erc1820RegistryCallerSession) GetManager(_addr common.Address) (common.Address, error) {
	return _Erc1820Registry.Contract.GetManager(&_Erc1820Registry.CallOpts, _addr)
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820Registry *Erc1820RegistryCaller) ImplementsERC165Interface(opts *bind.CallOpts, _contract common.Address, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Erc1820Registry.contract.Call(opts, out, "implementsERC165Interface", _contract, _interfaceId)
	return *ret0, err
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820Registry *Erc1820RegistrySession) ImplementsERC165Interface(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Erc1820Registry.Contract.ImplementsERC165Interface(&_Erc1820Registry.CallOpts, _contract, _interfaceId)
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820Registry *Erc1820RegistryCallerSession) ImplementsERC165Interface(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Erc1820Registry.Contract.ImplementsERC165Interface(&_Erc1820Registry.CallOpts, _contract, _interfaceId)
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820Registry *Erc1820RegistryCaller) ImplementsERC165InterfaceNoCache(opts *bind.CallOpts, _contract common.Address, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Erc1820Registry.contract.Call(opts, out, "implementsERC165InterfaceNoCache", _contract, _interfaceId)
	return *ret0, err
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820Registry *Erc1820RegistrySession) ImplementsERC165InterfaceNoCache(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Erc1820Registry.Contract.ImplementsERC165InterfaceNoCache(&_Erc1820Registry.CallOpts, _contract, _interfaceId)
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820Registry *Erc1820RegistryCallerSession) ImplementsERC165InterfaceNoCache(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Erc1820Registry.Contract.ImplementsERC165InterfaceNoCache(&_Erc1820Registry.CallOpts, _contract, _interfaceId)
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string _interfaceName) constant returns(bytes32)
func (_Erc1820Registry *Erc1820RegistryCaller) InterfaceHash(opts *bind.CallOpts, _interfaceName string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Erc1820Registry.contract.Call(opts, out, "interfaceHash", _interfaceName)
	return *ret0, err
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string _interfaceName) constant returns(bytes32)
func (_Erc1820Registry *Erc1820RegistrySession) InterfaceHash(_interfaceName string) ([32]byte, error) {
	return _Erc1820Registry.Contract.InterfaceHash(&_Erc1820Registry.CallOpts, _interfaceName)
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string _interfaceName) constant returns(bytes32)
func (_Erc1820Registry *Erc1820RegistryCallerSession) InterfaceHash(_interfaceName string) ([32]byte, error) {
	return _Erc1820Registry.Contract.InterfaceHash(&_Erc1820Registry.CallOpts, _interfaceName)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address _addr, bytes32 _interfaceHash, address _implementer) returns()
func (_Erc1820Registry *Erc1820RegistryTransactor) SetInterfaceImplementer(opts *bind.TransactOpts, _addr common.Address, _interfaceHash [32]byte, _implementer common.Address) (*types.Transaction, error) {
	return _Erc1820Registry.contract.Transact(opts, "setInterfaceImplementer", _addr, _interfaceHash, _implementer)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address _addr, bytes32 _interfaceHash, address _implementer) returns()
func (_Erc1820Registry *Erc1820RegistrySession) SetInterfaceImplementer(_addr common.Address, _interfaceHash [32]byte, _implementer common.Address) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.SetInterfaceImplementer(&_Erc1820Registry.TransactOpts, _addr, _interfaceHash, _implementer)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address _addr, bytes32 _interfaceHash, address _implementer) returns()
func (_Erc1820Registry *Erc1820RegistryTransactorSession) SetInterfaceImplementer(_addr common.Address, _interfaceHash [32]byte, _implementer common.Address) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.SetInterfaceImplementer(&_Erc1820Registry.TransactOpts, _addr, _interfaceHash, _implementer)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address _addr, address _newManager) returns()
func (_Erc1820Registry *Erc1820RegistryTransactor) SetManager(opts *bind.TransactOpts, _addr common.Address, _newManager common.Address) (*types.Transaction, error) {
	return _Erc1820Registry.contract.Transact(opts, "setManager", _addr, _newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address _addr, address _newManager) returns()
func (_Erc1820Registry *Erc1820RegistrySession) SetManager(_addr common.Address, _newManager common.Address) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.SetManager(&_Erc1820Registry.TransactOpts, _addr, _newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address _addr, address _newManager) returns()
func (_Erc1820Registry *Erc1820RegistryTransactorSession) SetManager(_addr common.Address, _newManager common.Address) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.SetManager(&_Erc1820Registry.TransactOpts, _addr, _newManager)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address _contract, bytes4 _interfaceId) returns()
func (_Erc1820Registry *Erc1820RegistryTransactor) UpdateERC165Cache(opts *bind.TransactOpts, _contract common.Address, _interfaceId [4]byte) (*types.Transaction, error) {
	return _Erc1820Registry.contract.Transact(opts, "updateERC165Cache", _contract, _interfaceId)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address _contract, bytes4 _interfaceId) returns()
func (_Erc1820Registry *Erc1820RegistrySession) UpdateERC165Cache(_contract common.Address, _interfaceId [4]byte) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.UpdateERC165Cache(&_Erc1820Registry.TransactOpts, _contract, _interfaceId)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address _contract, bytes4 _interfaceId) returns()
func (_Erc1820Registry *Erc1820RegistryTransactorSession) UpdateERC165Cache(_contract common.Address, _interfaceId [4]byte) (*types.Transaction, error) {
	return _Erc1820Registry.Contract.UpdateERC165Cache(&_Erc1820Registry.TransactOpts, _contract, _interfaceId)
}

// Erc1820RegistryInterfaceImplementerSetIterator is returned from FilterInterfaceImplementerSet and is used to iterate over the raw logs and unpacked data for InterfaceImplementerSet events raised by the Erc1820Registry contract.
type Erc1820RegistryInterfaceImplementerSetIterator struct {
	Event *Erc1820RegistryInterfaceImplementerSet // Event containing the contract specifics and raw log

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
func (it *Erc1820RegistryInterfaceImplementerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1820RegistryInterfaceImplementerSet)
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
		it.Event = new(Erc1820RegistryInterfaceImplementerSet)
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
func (it *Erc1820RegistryInterfaceImplementerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1820RegistryInterfaceImplementerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1820RegistryInterfaceImplementerSet represents a InterfaceImplementerSet event raised by the Erc1820Registry contract.
type Erc1820RegistryInterfaceImplementerSet struct {
	Addr          common.Address
	InterfaceHash [32]byte
	Implementer   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterfaceImplementerSet is a free log retrieval operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed addr, bytes32 indexed interfaceHash, address indexed implementer)
func (_Erc1820Registry *Erc1820RegistryFilterer) FilterInterfaceImplementerSet(opts *bind.FilterOpts, addr []common.Address, interfaceHash [][32]byte, implementer []common.Address) (*Erc1820RegistryInterfaceImplementerSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var interfaceHashRule []interface{}
	for _, interfaceHashItem := range interfaceHash {
		interfaceHashRule = append(interfaceHashRule, interfaceHashItem)
	}
	var implementerRule []interface{}
	for _, implementerItem := range implementer {
		implementerRule = append(implementerRule, implementerItem)
	}

	logs, sub, err := _Erc1820Registry.contract.FilterLogs(opts, "InterfaceImplementerSet", addrRule, interfaceHashRule, implementerRule)
	if err != nil {
		return nil, err
	}
	return &Erc1820RegistryInterfaceImplementerSetIterator{contract: _Erc1820Registry.contract, event: "InterfaceImplementerSet", logs: logs, sub: sub}, nil
}

// WatchInterfaceImplementerSet is a free log subscription operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed addr, bytes32 indexed interfaceHash, address indexed implementer)
func (_Erc1820Registry *Erc1820RegistryFilterer) WatchInterfaceImplementerSet(opts *bind.WatchOpts, sink chan<- *Erc1820RegistryInterfaceImplementerSet, addr []common.Address, interfaceHash [][32]byte, implementer []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var interfaceHashRule []interface{}
	for _, interfaceHashItem := range interfaceHash {
		interfaceHashRule = append(interfaceHashRule, interfaceHashItem)
	}
	var implementerRule []interface{}
	for _, implementerItem := range implementer {
		implementerRule = append(implementerRule, implementerItem)
	}

	logs, sub, err := _Erc1820Registry.contract.WatchLogs(opts, "InterfaceImplementerSet", addrRule, interfaceHashRule, implementerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1820RegistryInterfaceImplementerSet)
				if err := _Erc1820Registry.contract.UnpackLog(event, "InterfaceImplementerSet", log); err != nil {
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

// Erc1820RegistryManagerChangedIterator is returned from FilterManagerChanged and is used to iterate over the raw logs and unpacked data for ManagerChanged events raised by the Erc1820Registry contract.
type Erc1820RegistryManagerChangedIterator struct {
	Event *Erc1820RegistryManagerChanged // Event containing the contract specifics and raw log

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
func (it *Erc1820RegistryManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1820RegistryManagerChanged)
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
		it.Event = new(Erc1820RegistryManagerChanged)
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
func (it *Erc1820RegistryManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1820RegistryManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1820RegistryManagerChanged represents a ManagerChanged event raised by the Erc1820Registry contract.
type Erc1820RegistryManagerChanged struct {
	Addr       common.Address
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterManagerChanged is a free log retrieval operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed addr, address indexed newManager)
func (_Erc1820Registry *Erc1820RegistryFilterer) FilterManagerChanged(opts *bind.FilterOpts, addr []common.Address, newManager []common.Address) (*Erc1820RegistryManagerChangedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Erc1820Registry.contract.FilterLogs(opts, "ManagerChanged", addrRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &Erc1820RegistryManagerChangedIterator{contract: _Erc1820Registry.contract, event: "ManagerChanged", logs: logs, sub: sub}, nil
}

// WatchManagerChanged is a free log subscription operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed addr, address indexed newManager)
func (_Erc1820Registry *Erc1820RegistryFilterer) WatchManagerChanged(opts *bind.WatchOpts, sink chan<- *Erc1820RegistryManagerChanged, addr []common.Address, newManager []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Erc1820Registry.contract.WatchLogs(opts, "ManagerChanged", addrRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1820RegistryManagerChanged)
				if err := _Erc1820Registry.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
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
