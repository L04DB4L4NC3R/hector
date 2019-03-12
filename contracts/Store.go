// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"to\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ageSuggestive\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_class\",\"type\":\"string\"},{\"name\":\"_from\",\"type\":\"string\"},{\"name\":\"_to\",\"type\":\"string\"},{\"name\":\"_timestamp\",\"type\":\"string\"},{\"name\":\"_ageSuggestive\",\"type\":\"uint8\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"class\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"from\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_class\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_from\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ageSuggestive\",\"type\":\"uint8\"}],\"name\":\"dataSetter\",\"type\":\"event\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `608060405234801561001057600080fd5b506107e6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806313151981146100675780632ffa806f146100ea578063b4d09e3f1461010e578063b80777ea14610293578063bff852fa14610316578063d5ce338914610399575b600080fd5b61006f61041c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100af578082015181840152602081019050610094565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100f26104ba565b604051808260ff1660ff16815260200191505060405180910390f35b610291600480360360a081101561012457600080fd5b810190808035906020019064010000000081111561014157600080fd5b82018360208201111561015357600080fd5b8035906020019184600183028401116401000000008311171561017557600080fd5b90919293919293908035906020019064010000000081111561019657600080fd5b8201836020820111156101a857600080fd5b803590602001918460018302840111640100000000831117156101ca57600080fd5b9091929391929390803590602001906401000000008111156101eb57600080fd5b8201836020820111156101fd57600080fd5b8035906020019184600183028401116401000000008311171561021f57600080fd5b90919293919293908035906020019064010000000081111561024057600080fd5b82018360208201111561025257600080fd5b8035906020019184600183028401116401000000008311171561027457600080fd5b9091929391929390803560ff1690602001909291905050506104cd565b005b61029b61053b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102db5780820151818401526020810190506102c0565b50505050905090810190601f1680156103085780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61031e6105d9565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561035e578082015181840152602081019050610343565b50505050905090810190601f16801561038b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103a1610677565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103e15780820151818401526020810190506103c6565b50505050905090810190601f16801561040e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104b25780601f10610487576101008083540402835291602001916104b2565b820191906000526020600020905b81548152906001019060200180831161049557829003601f168201915b505050505081565b600460009054906101000a900460ff1681565b8686600091906104de929190610715565b508484600191906104f0929190610715565b50888860039190610502929190610715565b50828260029190610514929190610715565b5080600460006101000a81548160ff021916908360ff160217905550505050505050505050565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105d15780601f106105a6576101008083540402835291602001916105d1565b820191906000526020600020905b8154815290600101906020018083116105b457829003601f168201915b505050505081565b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561066f5780601f106106445761010080835404028352916020019161066f565b820191906000526020600020905b81548152906001019060200180831161065257829003601f168201915b505050505081565b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561070d5780601f106106e25761010080835404028352916020019161070d565b820191906000526020600020905b8154815290600101906020018083116106f057829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061075657803560ff1916838001178555610784565b82800160010185558215610784579182015b82811115610783578235825591602001919060010190610768565b5b5090506107919190610795565b5090565b6107b791905b808211156107b357600081600090555060010161079b565b5090565b9056fea165627a7a72305820efda32f8c442745d41f5fd5354a5bf5a0a40b0684db9096f800174a60edf86cc0029`

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// AgeSuggestive is a free data retrieval call binding the contract method 0x2ffa806f.
//
// Solidity: function ageSuggestive() constant returns(uint8)
func (_Store *StoreCaller) AgeSuggestive(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "ageSuggestive")
	return *ret0, err
}

// AgeSuggestive is a free data retrieval call binding the contract method 0x2ffa806f.
//
// Solidity: function ageSuggestive() constant returns(uint8)
func (_Store *StoreSession) AgeSuggestive() (uint8, error) {
	return _Store.Contract.AgeSuggestive(&_Store.CallOpts)
}

// AgeSuggestive is a free data retrieval call binding the contract method 0x2ffa806f.
//
// Solidity: function ageSuggestive() constant returns(uint8)
func (_Store *StoreCallerSession) AgeSuggestive() (uint8, error) {
	return _Store.Contract.AgeSuggestive(&_Store.CallOpts)
}

// Class is a free data retrieval call binding the contract method 0xbff852fa.
//
// Solidity: function class() constant returns(string)
func (_Store *StoreCaller) Class(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "class")
	return *ret0, err
}

// Class is a free data retrieval call binding the contract method 0xbff852fa.
//
// Solidity: function class() constant returns(string)
func (_Store *StoreSession) Class() (string, error) {
	return _Store.Contract.Class(&_Store.CallOpts)
}

// Class is a free data retrieval call binding the contract method 0xbff852fa.
//
// Solidity: function class() constant returns(string)
func (_Store *StoreCallerSession) Class() (string, error) {
	return _Store.Contract.Class(&_Store.CallOpts)
}

// From is a free data retrieval call binding the contract method 0xd5ce3389.
//
// Solidity: function from() constant returns(string)
func (_Store *StoreCaller) From(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "from")
	return *ret0, err
}

// From is a free data retrieval call binding the contract method 0xd5ce3389.
//
// Solidity: function from() constant returns(string)
func (_Store *StoreSession) From() (string, error) {
	return _Store.Contract.From(&_Store.CallOpts)
}

// From is a free data retrieval call binding the contract method 0xd5ce3389.
//
// Solidity: function from() constant returns(string)
func (_Store *StoreCallerSession) From() (string, error) {
	return _Store.Contract.From(&_Store.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() constant returns(string)
func (_Store *StoreCaller) Timestamp(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "timestamp")
	return *ret0, err
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() constant returns(string)
func (_Store *StoreSession) Timestamp() (string, error) {
	return _Store.Contract.Timestamp(&_Store.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() constant returns(string)
func (_Store *StoreCallerSession) Timestamp() (string, error) {
	return _Store.Contract.Timestamp(&_Store.CallOpts)
}

// To is a free data retrieval call binding the contract method 0x13151981.
//
// Solidity: function to() constant returns(string)
func (_Store *StoreCaller) To(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "to")
	return *ret0, err
}

// To is a free data retrieval call binding the contract method 0x13151981.
//
// Solidity: function to() constant returns(string)
func (_Store *StoreSession) To() (string, error) {
	return _Store.Contract.To(&_Store.CallOpts)
}

// To is a free data retrieval call binding the contract method 0x13151981.
//
// Solidity: function to() constant returns(string)
func (_Store *StoreCallerSession) To() (string, error) {
	return _Store.Contract.To(&_Store.CallOpts)
}

// SetData is a paid mutator transaction binding the contract method 0xb4d09e3f.
//
// Solidity: function setData(string _class, string _from, string _to, string _timestamp, uint8 _ageSuggestive) returns()
func (_Store *StoreTransactor) SetData(opts *bind.TransactOpts, _class string, _from string, _to string, _timestamp string, _ageSuggestive uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setData", _class, _from, _to, _timestamp, _ageSuggestive)
}

// SetData is a paid mutator transaction binding the contract method 0xb4d09e3f.
//
// Solidity: function setData(string _class, string _from, string _to, string _timestamp, uint8 _ageSuggestive) returns()
func (_Store *StoreSession) SetData(_class string, _from string, _to string, _timestamp string, _ageSuggestive uint8) (*types.Transaction, error) {
	return _Store.Contract.SetData(&_Store.TransactOpts, _class, _from, _to, _timestamp, _ageSuggestive)
}

// SetData is a paid mutator transaction binding the contract method 0xb4d09e3f.
//
// Solidity: function setData(string _class, string _from, string _to, string _timestamp, uint8 _ageSuggestive) returns()
func (_Store *StoreTransactorSession) SetData(_class string, _from string, _to string, _timestamp string, _ageSuggestive uint8) (*types.Transaction, error) {
	return _Store.Contract.SetData(&_Store.TransactOpts, _class, _from, _to, _timestamp, _ageSuggestive)
}

// StoreDataSetterIterator is returned from FilterDataSetter and is used to iterate over the raw logs and unpacked data for DataSetter events raised by the Store contract.
type StoreDataSetterIterator struct {
	Event *StoreDataSetter // Event containing the contract specifics and raw log

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
func (it *StoreDataSetterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDataSetter)
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
		it.Event = new(StoreDataSetter)
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
func (it *StoreDataSetterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDataSetterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDataSetter represents a DataSetter event raised by the Store contract.
type StoreDataSetter struct {
	Class         string
	From          string
	To            string
	Timestamp     string
	AgeSuggestive uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDataSetter is a free log retrieval operation binding the contract event 0x782d13adf76114d530de3a833e1ad406312093bf892d9798348cabda6b939158.
//
// Solidity: event dataSetter(string _class, string _from, string _to, string _timestamp, uint8 _ageSuggestive)
func (_Store *StoreFilterer) FilterDataSetter(opts *bind.FilterOpts) (*StoreDataSetterIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "dataSetter")
	if err != nil {
		return nil, err
	}
	return &StoreDataSetterIterator{contract: _Store.contract, event: "dataSetter", logs: logs, sub: sub}, nil
}

// WatchDataSetter is a free log subscription operation binding the contract event 0x782d13adf76114d530de3a833e1ad406312093bf892d9798348cabda6b939158.
//
// Solidity: event dataSetter(string _class, string _from, string _to, string _timestamp, uint8 _ageSuggestive)
func (_Store *StoreFilterer) WatchDataSetter(opts *bind.WatchOpts, sink chan<- *StoreDataSetter) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "dataSetter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDataSetter)
				if err := _Store.contract.UnpackLog(event, "dataSetter", log); err != nil {
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
