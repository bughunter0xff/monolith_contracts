// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package makerDAO

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

// ProxyRegistryABI is the input ABI used to generate the binding from.
const ProxyRegistryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxies\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ProxyRegistryBin is the compiled bytecode used for deploying new contracts.
const ProxyRegistryBin = `608060405234801561001057600080fd5b50604051602080610319833981016040525160018054600160a060020a031916600160a060020a039092169190911790556102c9806100506000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638e1a55fc811461005b578063c45527911461008c578063f3701da2146100ad575b600080fd5b34801561006757600080fd5b506100706100ce565b60408051600160a060020a039092168252519081900360200190f35b34801561009857600080fd5b50610070600160a060020a03600435166100de565b3480156100b957600080fd5b50610070600160a060020a03600435166100f9565b60006100d9336100f9565b905090565b600060208190529081526040902054600160a060020a031681565b600160a060020a0381811660009081526020819052604081205490911615806101be5750600160a060020a038083166000818152602081815260408083205481517f8da5cb5b000000000000000000000000000000000000000000000000000000008152915194951693638da5cb5b93600480840194938390030190829087803b15801561018657600080fd5b505af115801561019a573d6000803e3d6000fd5b505050506040513d60208110156101b057600080fd5b5051600160a060020a031614155b15156101c957600080fd5b600154604080517ff3701da2000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301529151919092169163f3701da29160248083019260209291908290030181600087803b15801561023157600080fd5b505af1158015610245573d6000803e3d6000fd5b505050506040513d602081101561025b57600080fd5b5051600160a060020a039283166000908152602081905260409020805473ffffffffffffffffffffffffffffffffffffffff19169382169390931790925550905600a165627a7a723058209b43552f7f5f6006f731bd096f7fc931e087019c0017b3ce0a6e77fb5c1ad5650029`

// DeployProxyRegistry deploys a new Ethereum contract, binding an instance of ProxyRegistry to it.
func DeployProxyRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, factory_ common.Address) (common.Address, *types.Transaction, *ProxyRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyRegistryBin), backend, factory_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyRegistry{ProxyRegistryCaller: ProxyRegistryCaller{contract: contract}, ProxyRegistryTransactor: ProxyRegistryTransactor{contract: contract}, ProxyRegistryFilterer: ProxyRegistryFilterer{contract: contract}}, nil
}

// ProxyRegistry is an auto generated Go binding around an Ethereum contract.
type ProxyRegistry struct {
	ProxyRegistryCaller     // Read-only binding to the contract
	ProxyRegistryTransactor // Write-only binding to the contract
	ProxyRegistryFilterer   // Log filterer for contract events
}

// ProxyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyRegistrySession struct {
	Contract     *ProxyRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyRegistryCallerSession struct {
	Contract *ProxyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProxyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyRegistryTransactorSession struct {
	Contract     *ProxyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProxyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRegistryRaw struct {
	Contract *ProxyRegistry // Generic contract binding to access the raw methods on
}

// ProxyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyRegistryCallerRaw struct {
	Contract *ProxyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyRegistryTransactorRaw struct {
	Contract *ProxyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyRegistry creates a new instance of ProxyRegistry, bound to a specific deployed contract.
func NewProxyRegistry(address common.Address, backend bind.ContractBackend) (*ProxyRegistry, error) {
	contract, err := bindProxyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyRegistry{ProxyRegistryCaller: ProxyRegistryCaller{contract: contract}, ProxyRegistryTransactor: ProxyRegistryTransactor{contract: contract}, ProxyRegistryFilterer: ProxyRegistryFilterer{contract: contract}}, nil
}

// NewProxyRegistryCaller creates a new read-only instance of ProxyRegistry, bound to a specific deployed contract.
func NewProxyRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProxyRegistryCaller, error) {
	contract, err := bindProxyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyRegistryCaller{contract: contract}, nil
}

// NewProxyRegistryTransactor creates a new write-only instance of ProxyRegistry, bound to a specific deployed contract.
func NewProxyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyRegistryTransactor, error) {
	contract, err := bindProxyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyRegistryTransactor{contract: contract}, nil
}

// NewProxyRegistryFilterer creates a new log filterer instance of ProxyRegistry, bound to a specific deployed contract.
func NewProxyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyRegistryFilterer, error) {
	contract, err := bindProxyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyRegistryFilterer{contract: contract}, nil
}

// bindProxyRegistry binds a generic wrapper to an already deployed contract.
func bindProxyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyRegistry *ProxyRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProxyRegistry.Contract.ProxyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyRegistry *ProxyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.ProxyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyRegistry *ProxyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.ProxyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyRegistry *ProxyRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProxyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyRegistry *ProxyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyRegistry *ProxyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.contract.Transact(opts, method, params...)
}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) constant returns(address)
func (_ProxyRegistry *ProxyRegistryCaller) Proxies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProxyRegistry.contract.Call(opts, out, "proxies", arg0)
	return *ret0, err
}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) constant returns(address)
func (_ProxyRegistry *ProxyRegistrySession) Proxies(arg0 common.Address) (common.Address, error) {
	return _ProxyRegistry.Contract.Proxies(&_ProxyRegistry.CallOpts, arg0)
}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) constant returns(address)
func (_ProxyRegistry *ProxyRegistryCallerSession) Proxies(arg0 common.Address) (common.Address, error) {
	return _ProxyRegistry.Contract.Proxies(&_ProxyRegistry.CallOpts, arg0)
}

// Build is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_ProxyRegistry *ProxyRegistryTransactor) Build(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ProxyRegistry.contract.Transact(opts, "build", owner)
}

// Build is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_ProxyRegistry *ProxyRegistrySession) Build(owner common.Address) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.Build(&_ProxyRegistry.TransactOpts, owner)
}

// Build is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_ProxyRegistry *ProxyRegistryTransactorSession) Build(owner common.Address) (*types.Transaction, error) {
	return _ProxyRegistry.Contract.Build(&_ProxyRegistry.TransactOpts, owner)
}
