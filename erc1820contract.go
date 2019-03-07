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

// Erc1820ABI is the input ABI used to generate the binding from.
const Erc1820ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_interfaceHash\",\"type\":\"bytes32\"},{\"name\":\"_implementer\",\"type\":\"address\"}],\"name\":\"setInterfaceImplementer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceName\",\"type\":\"string\"}],\"name\":\"interfaceHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"},{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"updateERC165Cache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_interfaceHash\",\"type\":\"bytes32\"}],\"name\":\"getInterfaceImplementer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"},{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"implementsERC165InterfaceNoCache\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"},{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"implementsERC165Interface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"interfaceHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"implementer\",\"type\":\"address\"}],\"name\":\"InterfaceImplementerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"ManagerChanged\",\"type\":\"event\"}]"

// Erc1820 is an auto generated Go binding around an Ethereum contract.
type Erc1820 struct {
	Erc1820Caller     // Read-only binding to the contract
	Erc1820Transactor // Write-only binding to the contract
	Erc1820Filterer   // Log filterer for contract events
}

// Erc1820Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc1820Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc1820Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc1820Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1820Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc1820Session struct {
	Contract     *Erc1820          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc1820CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc1820CallerSession struct {
	Contract *Erc1820Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Erc1820TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc1820TransactorSession struct {
	Contract     *Erc1820Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc1820Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc1820Raw struct {
	Contract *Erc1820 // Generic contract binding to access the raw methods on
}

// Erc1820CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc1820CallerRaw struct {
	Contract *Erc1820Caller // Generic read-only contract binding to access the raw methods on
}

// Erc1820TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc1820TransactorRaw struct {
	Contract *Erc1820Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc1820 creates a new instance of Erc1820, bound to a specific deployed contract.
func NewErc1820(address common.Address, backend bind.ContractBackend) (*Erc1820, error) {
	contract, err := bindErc1820(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc1820{Erc1820Caller: Erc1820Caller{contract: contract}, Erc1820Transactor: Erc1820Transactor{contract: contract}, Erc1820Filterer: Erc1820Filterer{contract: contract}}, nil
}

// NewErc1820Caller creates a new read-only instance of Erc1820, bound to a specific deployed contract.
func NewErc1820Caller(address common.Address, caller bind.ContractCaller) (*Erc1820Caller, error) {
	contract, err := bindErc1820(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1820Caller{contract: contract}, nil
}

// NewErc1820Transactor creates a new write-only instance of Erc1820, bound to a specific deployed contract.
func NewErc1820Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc1820Transactor, error) {
	contract, err := bindErc1820(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1820Transactor{contract: contract}, nil
}

// NewErc1820Filterer creates a new log filterer instance of Erc1820, bound to a specific deployed contract.
func NewErc1820Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc1820Filterer, error) {
	contract, err := bindErc1820(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc1820Filterer{contract: contract}, nil
}

// bindErc1820 binds a generic wrapper to an already deployed contract.
func bindErc1820(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1820ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1820 *Erc1820Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc1820.Contract.Erc1820Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1820 *Erc1820Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1820.Contract.Erc1820Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1820 *Erc1820Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1820.Contract.Erc1820Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1820 *Erc1820CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc1820.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1820 *Erc1820TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1820.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1820 *Erc1820TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1820.Contract.contract.Transact(opts, method, params...)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address _addr, bytes32 _interfaceHash) constant returns(address)
func (_Erc1820 *Erc1820Caller) GetInterfaceImplementer(opts *bind.CallOpts, _addr common.Address, _interfaceHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Erc1820.contract.Call(opts, out, "getInterfaceImplementer", _addr, _interfaceHash)
	return *ret0, err
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address _addr, bytes32 _interfaceHash) constant returns(address)
func (_Erc1820 *Erc1820Session) GetInterfaceImplementer(_addr common.Address, _interfaceHash [32]byte) (common.Address, error) {
	return _Erc1820.Contract.GetInterfaceImplementer(&_Erc1820.CallOpts, _addr, _interfaceHash)
}

// GetInterfaceImplementer is a free data retrieval call binding the contract method 0xaabbb8ca.
//
// Solidity: function getInterfaceImplementer(address _addr, bytes32 _interfaceHash) constant returns(address)
func (_Erc1820 *Erc1820CallerSession) GetInterfaceImplementer(_addr common.Address, _interfaceHash [32]byte) (common.Address, error) {
	return _Erc1820.Contract.GetInterfaceImplementer(&_Erc1820.CallOpts, _addr, _interfaceHash)
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address _addr) constant returns(address)
func (_Erc1820 *Erc1820Caller) GetManager(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Erc1820.contract.Call(opts, out, "getManager", _addr)
	return *ret0, err
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address _addr) constant returns(address)
func (_Erc1820 *Erc1820Session) GetManager(_addr common.Address) (common.Address, error) {
	return _Erc1820.Contract.GetManager(&_Erc1820.CallOpts, _addr)
}

// GetManager is a free data retrieval call binding the contract method 0x3d584063.
//
// Solidity: function getManager(address _addr) constant returns(address)
func (_Erc1820 *Erc1820CallerSession) GetManager(_addr common.Address) (common.Address, error) {
	return _Erc1820.Contract.GetManager(&_Erc1820.CallOpts, _addr)
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820 *Erc1820Caller) ImplementsERC165Interface(opts *bind.CallOpts, _contract common.Address, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Erc1820.contract.Call(opts, out, "implementsERC165Interface", _contract, _interfaceId)
	return *ret0, err
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820 *Erc1820Session) ImplementsERC165Interface(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Erc1820.Contract.ImplementsERC165Interface(&_Erc1820.CallOpts, _contract, _interfaceId)
}

// ImplementsERC165Interface is a free data retrieval call binding the contract method 0xf712f3e8.
//
// Solidity: function implementsERC165Interface(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820 *Erc1820CallerSession) ImplementsERC165Interface(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Erc1820.Contract.ImplementsERC165Interface(&_Erc1820.CallOpts, _contract, _interfaceId)
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820 *Erc1820Caller) ImplementsERC165InterfaceNoCache(opts *bind.CallOpts, _contract common.Address, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Erc1820.contract.Call(opts, out, "implementsERC165InterfaceNoCache", _contract, _interfaceId)
	return *ret0, err
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820 *Erc1820Session) ImplementsERC165InterfaceNoCache(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Erc1820.Contract.ImplementsERC165InterfaceNoCache(&_Erc1820.CallOpts, _contract, _interfaceId)
}

// ImplementsERC165InterfaceNoCache is a free data retrieval call binding the contract method 0xb7056765.
//
// Solidity: function implementsERC165InterfaceNoCache(address _contract, bytes4 _interfaceId) constant returns(bool)
func (_Erc1820 *Erc1820CallerSession) ImplementsERC165InterfaceNoCache(_contract common.Address, _interfaceId [4]byte) (bool, error) {
	return _Erc1820.Contract.ImplementsERC165InterfaceNoCache(&_Erc1820.CallOpts, _contract, _interfaceId)
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string _interfaceName) constant returns(bytes32)
func (_Erc1820 *Erc1820Caller) InterfaceHash(opts *bind.CallOpts, _interfaceName string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Erc1820.contract.Call(opts, out, "interfaceHash", _interfaceName)
	return *ret0, err
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string _interfaceName) constant returns(bytes32)
func (_Erc1820 *Erc1820Session) InterfaceHash(_interfaceName string) ([32]byte, error) {
	return _Erc1820.Contract.InterfaceHash(&_Erc1820.CallOpts, _interfaceName)
}

// InterfaceHash is a free data retrieval call binding the contract method 0x65ba36c1.
//
// Solidity: function interfaceHash(string _interfaceName) constant returns(bytes32)
func (_Erc1820 *Erc1820CallerSession) InterfaceHash(_interfaceName string) ([32]byte, error) {
	return _Erc1820.Contract.InterfaceHash(&_Erc1820.CallOpts, _interfaceName)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address _addr, bytes32 _interfaceHash, address _implementer) returns()
func (_Erc1820 *Erc1820Transactor) SetInterfaceImplementer(opts *bind.TransactOpts, _addr common.Address, _interfaceHash [32]byte, _implementer common.Address) (*types.Transaction, error) {
	return _Erc1820.contract.Transact(opts, "setInterfaceImplementer", _addr, _interfaceHash, _implementer)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address _addr, bytes32 _interfaceHash, address _implementer) returns()
func (_Erc1820 *Erc1820Session) SetInterfaceImplementer(_addr common.Address, _interfaceHash [32]byte, _implementer common.Address) (*types.Transaction, error) {
	return _Erc1820.Contract.SetInterfaceImplementer(&_Erc1820.TransactOpts, _addr, _interfaceHash, _implementer)
}

// SetInterfaceImplementer is a paid mutator transaction binding the contract method 0x29965a1d.
//
// Solidity: function setInterfaceImplementer(address _addr, bytes32 _interfaceHash, address _implementer) returns()
func (_Erc1820 *Erc1820TransactorSession) SetInterfaceImplementer(_addr common.Address, _interfaceHash [32]byte, _implementer common.Address) (*types.Transaction, error) {
	return _Erc1820.Contract.SetInterfaceImplementer(&_Erc1820.TransactOpts, _addr, _interfaceHash, _implementer)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address _addr, address _newManager) returns()
func (_Erc1820 *Erc1820Transactor) SetManager(opts *bind.TransactOpts, _addr common.Address, _newManager common.Address) (*types.Transaction, error) {
	return _Erc1820.contract.Transact(opts, "setManager", _addr, _newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address _addr, address _newManager) returns()
func (_Erc1820 *Erc1820Session) SetManager(_addr common.Address, _newManager common.Address) (*types.Transaction, error) {
	return _Erc1820.Contract.SetManager(&_Erc1820.TransactOpts, _addr, _newManager)
}

// SetManager is a paid mutator transaction binding the contract method 0x5df8122f.
//
// Solidity: function setManager(address _addr, address _newManager) returns()
func (_Erc1820 *Erc1820TransactorSession) SetManager(_addr common.Address, _newManager common.Address) (*types.Transaction, error) {
	return _Erc1820.Contract.SetManager(&_Erc1820.TransactOpts, _addr, _newManager)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address _contract, bytes4 _interfaceId) returns()
func (_Erc1820 *Erc1820Transactor) UpdateERC165Cache(opts *bind.TransactOpts, _contract common.Address, _interfaceId [4]byte) (*types.Transaction, error) {
	return _Erc1820.contract.Transact(opts, "updateERC165Cache", _contract, _interfaceId)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address _contract, bytes4 _interfaceId) returns()
func (_Erc1820 *Erc1820Session) UpdateERC165Cache(_contract common.Address, _interfaceId [4]byte) (*types.Transaction, error) {
	return _Erc1820.Contract.UpdateERC165Cache(&_Erc1820.TransactOpts, _contract, _interfaceId)
}

// UpdateERC165Cache is a paid mutator transaction binding the contract method 0xa41e7d51.
//
// Solidity: function updateERC165Cache(address _contract, bytes4 _interfaceId) returns()
func (_Erc1820 *Erc1820TransactorSession) UpdateERC165Cache(_contract common.Address, _interfaceId [4]byte) (*types.Transaction, error) {
	return _Erc1820.Contract.UpdateERC165Cache(&_Erc1820.TransactOpts, _contract, _interfaceId)
}

// Erc1820InterfaceImplementerSetIterator is returned from FilterInterfaceImplementerSet and is used to iterate over the raw logs and unpacked data for InterfaceImplementerSet events raised by the Erc1820 contract.
type Erc1820InterfaceImplementerSetIterator struct {
	Event *Erc1820InterfaceImplementerSet // Event containing the contract specifics and raw log

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
func (it *Erc1820InterfaceImplementerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1820InterfaceImplementerSet)
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
		it.Event = new(Erc1820InterfaceImplementerSet)
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
func (it *Erc1820InterfaceImplementerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1820InterfaceImplementerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1820InterfaceImplementerSet represents a InterfaceImplementerSet event raised by the Erc1820 contract.
type Erc1820InterfaceImplementerSet struct {
	Addr          common.Address
	InterfaceHash [32]byte
	Implementer   common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterfaceImplementerSet is a free log retrieval operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed addr, bytes32 indexed interfaceHash, address indexed implementer)
func (_Erc1820 *Erc1820Filterer) FilterInterfaceImplementerSet(opts *bind.FilterOpts, addr []common.Address, interfaceHash [][32]byte, implementer []common.Address) (*Erc1820InterfaceImplementerSetIterator, error) {

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

	logs, sub, err := _Erc1820.contract.FilterLogs(opts, "InterfaceImplementerSet", addrRule, interfaceHashRule, implementerRule)
	if err != nil {
		return nil, err
	}
	return &Erc1820InterfaceImplementerSetIterator{contract: _Erc1820.contract, event: "InterfaceImplementerSet", logs: logs, sub: sub}, nil
}

// WatchInterfaceImplementerSet is a free log subscription operation binding the contract event 0x93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db153.
//
// Solidity: event InterfaceImplementerSet(address indexed addr, bytes32 indexed interfaceHash, address indexed implementer)
func (_Erc1820 *Erc1820Filterer) WatchInterfaceImplementerSet(opts *bind.WatchOpts, sink chan<- *Erc1820InterfaceImplementerSet, addr []common.Address, interfaceHash [][32]byte, implementer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Erc1820.contract.WatchLogs(opts, "InterfaceImplementerSet", addrRule, interfaceHashRule, implementerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1820InterfaceImplementerSet)
				if err := _Erc1820.contract.UnpackLog(event, "InterfaceImplementerSet", log); err != nil {
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

// Erc1820ManagerChangedIterator is returned from FilterManagerChanged and is used to iterate over the raw logs and unpacked data for ManagerChanged events raised by the Erc1820 contract.
type Erc1820ManagerChangedIterator struct {
	Event *Erc1820ManagerChanged // Event containing the contract specifics and raw log

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
func (it *Erc1820ManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1820ManagerChanged)
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
		it.Event = new(Erc1820ManagerChanged)
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
func (it *Erc1820ManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1820ManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1820ManagerChanged represents a ManagerChanged event raised by the Erc1820 contract.
type Erc1820ManagerChanged struct {
	Addr       common.Address
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterManagerChanged is a free log retrieval operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed addr, address indexed newManager)
func (_Erc1820 *Erc1820Filterer) FilterManagerChanged(opts *bind.FilterOpts, addr []common.Address, newManager []common.Address) (*Erc1820ManagerChangedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Erc1820.contract.FilterLogs(opts, "ManagerChanged", addrRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &Erc1820ManagerChangedIterator{contract: _Erc1820.contract, event: "ManagerChanged", logs: logs, sub: sub}, nil
}

// WatchManagerChanged is a free log subscription operation binding the contract event 0x605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a4350.
//
// Solidity: event ManagerChanged(address indexed addr, address indexed newManager)
func (_Erc1820 *Erc1820Filterer) WatchManagerChanged(opts *bind.WatchOpts, sink chan<- *Erc1820ManagerChanged, addr []common.Address, newManager []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Erc1820.contract.WatchLogs(opts, "ManagerChanged", addrRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1820ManagerChanged)
				if err := _Erc1820.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
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
