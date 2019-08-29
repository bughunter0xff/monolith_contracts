// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kyber

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

// KyberNetworkProxyABI is the input ABI used to generate the binding from.
const KyberNetworkProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"trader\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"name\":\"walletId\",\"type\":\"address\"},{\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"tradeWithHint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"name\":\"walletId\",\"type\":\"address\"},{\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"tradeWithHint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxGasPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetworkContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserCapInWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapEtherToToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"srcQty\",\"type\":\"uint256\"}],\"name\":\"getExpectedRate\",\"outputs\":[{\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"name\":\"slippageRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getUserCapInTokenWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kyberNetworkContract\",\"type\":\"address\"}],\"name\":\"setKyberNetworkContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"info\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"name\":\"walletId\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"actualSrcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"actualDestAmount\",\"type\":\"uint256\"}],\"name\":\"ExecuteTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newNetworkContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldNetworkContract\",\"type\":\"address\"}],\"name\":\"KyberNetworkSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"}]"

// KyberNetworkProxy is an auto generated Go binding around an Ethereum contract.
type KyberNetworkProxy struct {
	KyberNetworkProxyCaller     // Read-only binding to the contract
	KyberNetworkProxyTransactor // Write-only binding to the contract
	KyberNetworkProxyFilterer   // Log filterer for contract events
}

// KyberNetworkProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberNetworkProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberNetworkProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberNetworkProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberNetworkProxySession struct {
	Contract     *KyberNetworkProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KyberNetworkProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberNetworkProxyCallerSession struct {
	Contract *KyberNetworkProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KyberNetworkProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberNetworkProxyTransactorSession struct {
	Contract     *KyberNetworkProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KyberNetworkProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberNetworkProxyRaw struct {
	Contract *KyberNetworkProxy // Generic contract binding to access the raw methods on
}

// KyberNetworkProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberNetworkProxyCallerRaw struct {
	Contract *KyberNetworkProxyCaller // Generic read-only contract binding to access the raw methods on
}

// KyberNetworkProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberNetworkProxyTransactorRaw struct {
	Contract *KyberNetworkProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberNetworkProxy creates a new instance of KyberNetworkProxy, bound to a specific deployed contract.
func NewKyberNetworkProxy(address common.Address, backend bind.ContractBackend) (*KyberNetworkProxy, error) {
	contract, err := bindKyberNetworkProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxy{KyberNetworkProxyCaller: KyberNetworkProxyCaller{contract: contract}, KyberNetworkProxyTransactor: KyberNetworkProxyTransactor{contract: contract}, KyberNetworkProxyFilterer: KyberNetworkProxyFilterer{contract: contract}}, nil
}

// NewKyberNetworkProxyCaller creates a new read-only instance of KyberNetworkProxy, bound to a specific deployed contract.
func NewKyberNetworkProxyCaller(address common.Address, caller bind.ContractCaller) (*KyberNetworkProxyCaller, error) {
	contract, err := bindKyberNetworkProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyCaller{contract: contract}, nil
}

// NewKyberNetworkProxyTransactor creates a new write-only instance of KyberNetworkProxy, bound to a specific deployed contract.
func NewKyberNetworkProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberNetworkProxyTransactor, error) {
	contract, err := bindKyberNetworkProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyTransactor{contract: contract}, nil
}

// NewKyberNetworkProxyFilterer creates a new log filterer instance of KyberNetworkProxy, bound to a specific deployed contract.
func NewKyberNetworkProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberNetworkProxyFilterer, error) {
	contract, err := bindKyberNetworkProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyFilterer{contract: contract}, nil
}

// bindKyberNetworkProxy binds a generic wrapper to an already deployed contract.
func bindKyberNetworkProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberNetworkProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberNetworkProxy *KyberNetworkProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberNetworkProxy.Contract.KyberNetworkProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberNetworkProxy *KyberNetworkProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.KyberNetworkProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberNetworkProxy *KyberNetworkProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.KyberNetworkProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberNetworkProxy *KyberNetworkProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberNetworkProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberNetworkProxy *KyberNetworkProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberNetworkProxy *KyberNetworkProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxySession) Admin() (common.Address, error) {
	return _KyberNetworkProxy.Contract.Admin(&_KyberNetworkProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) Admin() (common.Address, error) {
	return _KyberNetworkProxy.Contract.Admin(&_KyberNetworkProxy.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "enabled")
	return *ret0, err
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetworkProxy *KyberNetworkProxySession) Enabled() (bool, error) {
	return _KyberNetworkProxy.Contract.Enabled(&_KyberNetworkProxy.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) Enabled() (bool, error) {
	return _KyberNetworkProxy.Contract.Enabled(&_KyberNetworkProxy.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxySession) GetAlerters() ([]common.Address, error) {
	return _KyberNetworkProxy.Contract.GetAlerters(&_KyberNetworkProxy.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetAlerters() ([]common.Address, error) {
	return _KyberNetworkProxy.Contract.GetAlerters(&_KyberNetworkProxy.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetBalance(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getBalance", token, user)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetBalance(&_KyberNetworkProxy.CallOpts, token, user)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetBalance(&_KyberNetworkProxy.CallOpts, token, user)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetExpectedRate(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	ret := new(struct {
		ExpectedRate *big.Int
		SlippageRate *big.Int
	})
	out := ret
	err := _KyberNetworkProxy.contract.Call(opts, out, "getExpectedRate", src, dest, srcQty)
	return *ret, err
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetworkProxy *KyberNetworkProxySession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetworkProxy.Contract.GetExpectedRate(&_KyberNetworkProxy.CallOpts, src, dest, srcQty)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetworkProxy.Contract.GetExpectedRate(&_KyberNetworkProxy.CallOpts, src, dest, srcQty)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxySession) GetOperators() ([]common.Address, error) {
	return _KyberNetworkProxy.Contract.GetOperators(&_KyberNetworkProxy.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetOperators() ([]common.Address, error) {
	return _KyberNetworkProxy.Contract.GetOperators(&_KyberNetworkProxy.CallOpts)
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetUserCapInTokenWei(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getUserCapInTokenWei", user, token)
	return *ret0, err
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) GetUserCapInTokenWei(user common.Address, token common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetUserCapInTokenWei(&_KyberNetworkProxy.CallOpts, user, token)
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetUserCapInTokenWei(user common.Address, token common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetUserCapInTokenWei(&_KyberNetworkProxy.CallOpts, user, token)
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetUserCapInWei(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getUserCapInWei", user)
	return *ret0, err
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) GetUserCapInWei(user common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetUserCapInWei(&_KyberNetworkProxy.CallOpts, user)
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetUserCapInWei(user common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetUserCapInWei(&_KyberNetworkProxy.CallOpts, user)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) Info(opts *bind.CallOpts, field [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "info", field)
	return *ret0, err
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) Info(field [32]byte) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.Info(&_KyberNetworkProxy.CallOpts, field)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) Info(field [32]byte) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.Info(&_KyberNetworkProxy.CallOpts, field)
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) KyberNetworkContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "kyberNetworkContract")
	return *ret0, err
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxySession) KyberNetworkContract() (common.Address, error) {
	return _KyberNetworkProxy.Contract.KyberNetworkContract(&_KyberNetworkProxy.CallOpts)
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) KyberNetworkContract() (common.Address, error) {
	return _KyberNetworkProxy.Contract.KyberNetworkContract(&_KyberNetworkProxy.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "maxGasPrice")
	return *ret0, err
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) MaxGasPrice() (*big.Int, error) {
	return _KyberNetworkProxy.Contract.MaxGasPrice(&_KyberNetworkProxy.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) MaxGasPrice() (*big.Int, error) {
	return _KyberNetworkProxy.Contract.MaxGasPrice(&_KyberNetworkProxy.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxySession) PendingAdmin() (common.Address, error) {
	return _KyberNetworkProxy.Contract.PendingAdmin(&_KyberNetworkProxy.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) PendingAdmin() (common.Address, error) {
	return _KyberNetworkProxy.Contract.PendingAdmin(&_KyberNetworkProxy.CallOpts)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.AddAlerter(&_KyberNetworkProxy.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.AddAlerter(&_KyberNetworkProxy.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.AddOperator(&_KyberNetworkProxy.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.AddOperator(&_KyberNetworkProxy.TransactOpts, newOperator)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.ClaimAdmin(&_KyberNetworkProxy.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.ClaimAdmin(&_KyberNetworkProxy.TransactOpts)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.RemoveAlerter(&_KyberNetworkProxy.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.RemoveAlerter(&_KyberNetworkProxy.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.RemoveOperator(&_KyberNetworkProxy.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.RemoveOperator(&_KyberNetworkProxy.TransactOpts, operator)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) SetKyberNetworkContract(opts *bind.TransactOpts, _kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "setKyberNetworkContract", _kyberNetworkContract)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) SetKyberNetworkContract(_kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SetKyberNetworkContract(&_KyberNetworkProxy.TransactOpts, _kyberNetworkContract)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) SetKyberNetworkContract(_kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SetKyberNetworkContract(&_KyberNetworkProxy.TransactOpts, _kyberNetworkContract)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) SwapEtherToToken(opts *bind.TransactOpts, token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "swapEtherToToken", token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapEtherToToken(&_KyberNetworkProxy.TransactOpts, token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapEtherToToken(&_KyberNetworkProxy.TransactOpts, token, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) SwapTokenToEther(opts *bind.TransactOpts, token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "swapTokenToEther", token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapTokenToEther(&_KyberNetworkProxy.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapTokenToEther(&_KyberNetworkProxy.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) SwapTokenToToken(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "swapTokenToToken", src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapTokenToToken(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapTokenToToken(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) Trade(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "trade", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.Trade(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.Trade(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x088322ef.
//
// Solidity: function tradeWithHint(address trader, address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) TradeWithHint(opts *bind.TransactOpts, trader common.Address, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "tradeWithHint", trader, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x088322ef.
//
// Solidity: function tradeWithHint(address trader, address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) TradeWithHint(trader common.Address, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TradeWithHint(&_KyberNetworkProxy.TransactOpts, trader, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x088322ef.
//
// Solidity: function tradeWithHint(address trader, address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) TradeWithHint(trader common.Address, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TradeWithHint(&_KyberNetworkProxy.TransactOpts, trader, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint0 is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) TradeWithHint0(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "tradeWithHint0", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint0 is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) TradeWithHint0(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TradeWithHint0(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint0 is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) TradeWithHint0(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TradeWithHint0(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TransferAdmin(&_KyberNetworkProxy.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TransferAdmin(&_KyberNetworkProxy.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TransferAdminQuickly(&_KyberNetworkProxy.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TransferAdminQuickly(&_KyberNetworkProxy.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.WithdrawEther(&_KyberNetworkProxy.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.WithdrawEther(&_KyberNetworkProxy.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.WithdrawToken(&_KyberNetworkProxy.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.WithdrawToken(&_KyberNetworkProxy.TransactOpts, token, amount, sendTo)
}

// KyberNetworkProxyAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyAdminClaimedIterator struct {
	Event *KyberNetworkProxyAdminClaimed // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyAdminClaimed)
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
		it.Event = new(KyberNetworkProxyAdminClaimed)
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
func (it *KyberNetworkProxyAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyAdminClaimed represents a AdminClaimed event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*KyberNetworkProxyAdminClaimedIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyAdminClaimedIterator{contract: _KyberNetworkProxy.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyAdminClaimed)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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

// ParseAdminClaimed is a log parse operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseAdminClaimed(log types.Log) (*KyberNetworkProxyAdminClaimed, error) {
	event := new(KyberNetworkProxyAdminClaimed)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyAlerterAddedIterator struct {
	Event *KyberNetworkProxyAlerterAdded // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyAlerterAdded)
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
		it.Event = new(KyberNetworkProxyAlerterAdded)
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
func (it *KyberNetworkProxyAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyAlerterAdded represents a AlerterAdded event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*KyberNetworkProxyAlerterAddedIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyAlerterAddedIterator{contract: _KyberNetworkProxy.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyAlerterAdded)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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

// ParseAlerterAdded is a log parse operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseAlerterAdded(log types.Log) (*KyberNetworkProxyAlerterAdded, error) {
	event := new(KyberNetworkProxyAlerterAdded)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyEtherWithdrawIterator struct {
	Event *KyberNetworkProxyEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyEtherWithdraw)
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
		it.Event = new(KyberNetworkProxyEtherWithdraw)
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
func (it *KyberNetworkProxyEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyEtherWithdraw represents a EtherWithdraw event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*KyberNetworkProxyEtherWithdrawIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyEtherWithdrawIterator{contract: _KyberNetworkProxy.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyEtherWithdraw)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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

// ParseEtherWithdraw is a log parse operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseEtherWithdraw(log types.Log) (*KyberNetworkProxyEtherWithdraw, error) {
	event := new(KyberNetworkProxyEtherWithdraw)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyExecuteTradeIterator is returned from FilterExecuteTrade and is used to iterate over the raw logs and unpacked data for ExecuteTrade events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyExecuteTradeIterator struct {
	Event *KyberNetworkProxyExecuteTrade // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyExecuteTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyExecuteTrade)
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
		it.Event = new(KyberNetworkProxyExecuteTrade)
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
func (it *KyberNetworkProxyExecuteTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyExecuteTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyExecuteTrade represents a ExecuteTrade event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyExecuteTrade struct {
	Trader           common.Address
	Src              common.Address
	Dest             common.Address
	ActualSrcAmount  *big.Int
	ActualDestAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecuteTrade is a free log retrieval operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterExecuteTrade(opts *bind.FilterOpts, trader []common.Address) (*KyberNetworkProxyExecuteTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "ExecuteTrade", traderRule)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyExecuteTradeIterator{contract: _KyberNetworkProxy.contract, event: "ExecuteTrade", logs: logs, sub: sub}, nil
}

// WatchExecuteTrade is a free log subscription operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchExecuteTrade(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyExecuteTrade, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "ExecuteTrade", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyExecuteTrade)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "ExecuteTrade", log); err != nil {
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

// ParseExecuteTrade is a log parse operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseExecuteTrade(log types.Log) (*KyberNetworkProxyExecuteTrade, error) {
	event := new(KyberNetworkProxyExecuteTrade)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "ExecuteTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyKyberNetworkSetIterator is returned from FilterKyberNetworkSet and is used to iterate over the raw logs and unpacked data for KyberNetworkSet events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyKyberNetworkSetIterator struct {
	Event *KyberNetworkProxyKyberNetworkSet // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyKyberNetworkSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyKyberNetworkSet)
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
		it.Event = new(KyberNetworkProxyKyberNetworkSet)
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
func (it *KyberNetworkProxyKyberNetworkSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyKyberNetworkSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyKyberNetworkSet represents a KyberNetworkSet event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyKyberNetworkSet struct {
	NewNetworkContract common.Address
	OldNetworkContract common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterKyberNetworkSet is a free log retrieval operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterKyberNetworkSet(opts *bind.FilterOpts) (*KyberNetworkProxyKyberNetworkSetIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "KyberNetworkSet")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyKyberNetworkSetIterator{contract: _KyberNetworkProxy.contract, event: "KyberNetworkSet", logs: logs, sub: sub}, nil
}

// WatchKyberNetworkSet is a free log subscription operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchKyberNetworkSet(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyKyberNetworkSet) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "KyberNetworkSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyKyberNetworkSet)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "KyberNetworkSet", log); err != nil {
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

// ParseKyberNetworkSet is a log parse operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseKyberNetworkSet(log types.Log) (*KyberNetworkProxyKyberNetworkSet, error) {
	event := new(KyberNetworkProxyKyberNetworkSet)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "KyberNetworkSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyOperatorAddedIterator struct {
	Event *KyberNetworkProxyOperatorAdded // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyOperatorAdded)
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
		it.Event = new(KyberNetworkProxyOperatorAdded)
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
func (it *KyberNetworkProxyOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyOperatorAdded represents a OperatorAdded event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*KyberNetworkProxyOperatorAddedIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyOperatorAddedIterator{contract: _KyberNetworkProxy.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyOperatorAdded)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseOperatorAdded(log types.Log) (*KyberNetworkProxyOperatorAdded, error) {
	event := new(KyberNetworkProxyOperatorAdded)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyTokenWithdrawIterator struct {
	Event *KyberNetworkProxyTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyTokenWithdraw)
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
		it.Event = new(KyberNetworkProxyTokenWithdraw)
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
func (it *KyberNetworkProxyTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyTokenWithdraw represents a TokenWithdraw event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*KyberNetworkProxyTokenWithdrawIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyTokenWithdrawIterator{contract: _KyberNetworkProxy.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyTokenWithdraw)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseTokenWithdraw(log types.Log) (*KyberNetworkProxyTokenWithdraw, error) {
	event := new(KyberNetworkProxyTokenWithdraw)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyTransferAdminPendingIterator struct {
	Event *KyberNetworkProxyTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyTransferAdminPending)
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
		it.Event = new(KyberNetworkProxyTransferAdminPending)
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
func (it *KyberNetworkProxyTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyTransferAdminPending represents a TransferAdminPending event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*KyberNetworkProxyTransferAdminPendingIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyTransferAdminPendingIterator{contract: _KyberNetworkProxy.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyTransferAdminPending)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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

// ParseTransferAdminPending is a log parse operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseTransferAdminPending(log types.Log) (*KyberNetworkProxyTransferAdminPending, error) {
	event := new(KyberNetworkProxyTransferAdminPending)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}
