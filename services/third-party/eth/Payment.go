// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PaymentABI is the input ABI used to generate the binding from.
const PaymentABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"trade_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"order_no\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"name\":\"Notify\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"order_no\",\"type\":\"string\"}],\"name\":\"checkPay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"order_no\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"order_no\",\"type\":\"string\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"tradesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tradesOfAt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Payment is an auto generated Go binding around an Ethereum contract.
type Payment struct {
	PaymentCaller     // Read-only binding to the contract
	PaymentTransactor // Write-only binding to the contract
	PaymentFilterer   // Log filterer for contract events
}

// PaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentSession struct {
	Contract     *Payment          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentCallerSession struct {
	Contract *PaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentTransactorSession struct {
	Contract     *PaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentRaw struct {
	Contract *Payment // Generic contract binding to access the raw methods on
}

// PaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentCallerRaw struct {
	Contract *PaymentCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentTransactorRaw struct {
	Contract *PaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayment creates a new instance of Payment, bound to a specific deployed contract.
func NewPayment(address common.Address, backend bind.ContractBackend) (*Payment, error) {
	contract, err := bindPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payment{PaymentCaller: PaymentCaller{contract: contract}, PaymentTransactor: PaymentTransactor{contract: contract}, PaymentFilterer: PaymentFilterer{contract: contract}}, nil
}

// NewPaymentCaller creates a new read-only instance of Payment, bound to a specific deployed contract.
func NewPaymentCaller(address common.Address, caller bind.ContractCaller) (*PaymentCaller, error) {
	contract, err := bindPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCaller{contract: contract}, nil
}

// NewPaymentTransactor creates a new write-only instance of Payment, bound to a specific deployed contract.
func NewPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentTransactor, error) {
	contract, err := bindPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentTransactor{contract: contract}, nil
}

// NewPaymentFilterer creates a new log filterer instance of Payment, bound to a specific deployed contract.
func NewPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentFilterer, error) {
	contract, err := bindPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentFilterer{contract: contract}, nil
}

// bindPayment binds a generic wrapper to an already deployed contract.
func bindPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.PaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transact(opts, method, params...)
}

// CheckPay is a free data retrieval call binding the contract method 0xafe92d7e.
//
// Solidity: function checkPay(string order_no) view returns(bool)
func (_Payment *PaymentCaller) CheckPay(opts *bind.CallOpts, order_no string) (bool, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "checkPay", order_no)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPay is a free data retrieval call binding the contract method 0xafe92d7e.
//
// Solidity: function checkPay(string order_no) view returns(bool)
func (_Payment *PaymentSession) CheckPay(order_no string) (bool, error) {
	return _Payment.Contract.CheckPay(&_Payment.CallOpts, order_no)
}

// CheckPay is a free data retrieval call binding the contract method 0xafe92d7e.
//
// Solidity: function checkPay(string order_no) view returns(bool)
func (_Payment *PaymentCallerSession) CheckPay(order_no string) (bool, error) {
	return _Payment.Contract.CheckPay(&_Payment.CallOpts, order_no)
}

// TradesOf is a free data retrieval call binding the contract method 0x6c081982.
//
// Solidity: function tradesOf(address buyer) view returns(uint256)
func (_Payment *PaymentCaller) TradesOf(opts *bind.CallOpts, buyer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "tradesOf", buyer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradesOf is a free data retrieval call binding the contract method 0x6c081982.
//
// Solidity: function tradesOf(address buyer) view returns(uint256)
func (_Payment *PaymentSession) TradesOf(buyer common.Address) (*big.Int, error) {
	return _Payment.Contract.TradesOf(&_Payment.CallOpts, buyer)
}

// TradesOf is a free data retrieval call binding the contract method 0x6c081982.
//
// Solidity: function tradesOf(address buyer) view returns(uint256)
func (_Payment *PaymentCallerSession) TradesOf(buyer common.Address) (*big.Int, error) {
	return _Payment.Contract.TradesOf(&_Payment.CallOpts, buyer)
}

// TradesOfAt is a free data retrieval call binding the contract method 0x31cd7e4b.
//
// Solidity: function tradesOfAt(address buyer, uint256 index) view returns(string, string, uint256 amount, uint256 date)
func (_Payment *PaymentCaller) TradesOfAt(opts *bind.CallOpts, buyer common.Address, index *big.Int) (string, string, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "tradesOfAt", buyer, index)

	if err != nil {
		return *new(string), *new(string), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// TradesOfAt is a free data retrieval call binding the contract method 0x31cd7e4b.
//
// Solidity: function tradesOfAt(address buyer, uint256 index) view returns(string, string, uint256 amount, uint256 date)
func (_Payment *PaymentSession) TradesOfAt(buyer common.Address, index *big.Int) (string, string, *big.Int, *big.Int, error) {
	return _Payment.Contract.TradesOfAt(&_Payment.CallOpts, buyer, index)
}

// TradesOfAt is a free data retrieval call binding the contract method 0x31cd7e4b.
//
// Solidity: function tradesOfAt(address buyer, uint256 index) view returns(string, string, uint256 amount, uint256 date)
func (_Payment *PaymentCallerSession) TradesOfAt(buyer common.Address, index *big.Int) (string, string, *big.Int, *big.Int, error) {
	return _Payment.Contract.TradesOfAt(&_Payment.CallOpts, buyer, index)
}

// Pay is a paid mutator transaction binding the contract method 0xf621217e.
//
// Solidity: function pay(string id, string order_no, uint256 value) payable returns()
func (_Payment *PaymentTransactor) Pay(opts *bind.TransactOpts, id string, order_no string, value *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "pay", id, order_no, value)
}

// Pay is a paid mutator transaction binding the contract method 0xf621217e.
//
// Solidity: function pay(string id, string order_no, uint256 value) payable returns()
func (_Payment *PaymentSession) Pay(id string, order_no string, value *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Pay(&_Payment.TransactOpts, id, order_no, value)
}

// Pay is a paid mutator transaction binding the contract method 0xf621217e.
//
// Solidity: function pay(string id, string order_no, uint256 value) payable returns()
func (_Payment *PaymentTransactorSession) Pay(id string, order_no string, value *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Pay(&_Payment.TransactOpts, id, order_no, value)
}

// Refund is a paid mutator transaction binding the contract method 0xfe5f2e88.
//
// Solidity: function refund(string order_no) payable returns()
func (_Payment *PaymentTransactor) Refund(opts *bind.TransactOpts, order_no string) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "refund", order_no)
}

// Refund is a paid mutator transaction binding the contract method 0xfe5f2e88.
//
// Solidity: function refund(string order_no) payable returns()
func (_Payment *PaymentSession) Refund(order_no string) (*types.Transaction, error) {
	return _Payment.Contract.Refund(&_Payment.TransactOpts, order_no)
}

// Refund is a paid mutator transaction binding the contract method 0xfe5f2e88.
//
// Solidity: function refund(string order_no) payable returns()
func (_Payment *PaymentTransactorSession) Refund(order_no string) (*types.Transaction, error) {
	return _Payment.Contract.Refund(&_Payment.TransactOpts, order_no)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_Payment *PaymentTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_Payment *PaymentSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Withdraw(&_Payment.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_Payment *PaymentTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Withdraw(&_Payment.TransactOpts, amount)
}

// PaymentNotifyIterator is returned from FilterNotify and is used to iterate over the raw logs and unpacked data for Notify events raised by the Payment contract.
type PaymentNotifyIterator struct {
	Event *PaymentNotify // Event containing the contract specifics and raw log

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
func (it *PaymentNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentNotify)
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
		it.Event = new(PaymentNotify)
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
func (it *PaymentNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentNotify represents a Notify event raised by the Payment contract.
type PaymentNotify struct {
	Addr    common.Address
	TradeId string
	OrderNo string
	Date    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNotify is a free log retrieval operation binding the contract event 0xc2a9699e623ff301b2992c4111f4c3d6e02462a5328ac118aab1c35af1be4de7.
//
// Solidity: event Notify(address addr, string trade_id, string order_no, uint256 date)
func (_Payment *PaymentFilterer) FilterNotify(opts *bind.FilterOpts) (*PaymentNotifyIterator, error) {

	logs, sub, err := _Payment.contract.FilterLogs(opts, "Notify")
	if err != nil {
		return nil, err
	}
	return &PaymentNotifyIterator{contract: _Payment.contract, event: "Notify", logs: logs, sub: sub}, nil
}

// WatchNotify is a free log subscription operation binding the contract event 0xc2a9699e623ff301b2992c4111f4c3d6e02462a5328ac118aab1c35af1be4de7.
//
// Solidity: event Notify(address addr, string trade_id, string order_no, uint256 date)
func (_Payment *PaymentFilterer) WatchNotify(opts *bind.WatchOpts, sink chan<- *PaymentNotify) (event.Subscription, error) {

	logs, sub, err := _Payment.contract.WatchLogs(opts, "Notify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentNotify)
				if err := _Payment.contract.UnpackLog(event, "Notify", log); err != nil {
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

// ParseNotify is a log parse operation binding the contract event 0xc2a9699e623ff301b2992c4111f4c3d6e02462a5328ac118aab1c35af1be4de7.
//
// Solidity: event Notify(address addr, string trade_id, string order_no, uint256 date)
func (_Payment *PaymentFilterer) ParseNotify(log types.Log) (*PaymentNotify, error) {
	event := new(PaymentNotify)
	if err := _Payment.contract.UnpackLog(event, "Notify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
