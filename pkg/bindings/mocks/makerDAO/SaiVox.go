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

// SaiVoxABI is the input ABI used to generate the binding from.
const SaiVoxABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"prod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"era\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"how\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"par\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ray\",\"type\":\"uint256\"}],\"name\":\"tell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"way\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"param\",\"type\":\"bytes32\"},{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"mold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fix\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ray\",\"type\":\"uint256\"}],\"name\":\"tune\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tau\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"par_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// SaiVoxBin is the compiled bytecode used for deploying new contracts.
var SaiVoxBin = "0x6060604052341561000f57600080fd5b6040516020806109b98339810160405280805160018054600160a060020a03191633600160a060020a031690811790915590925090507fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a2600481905560028190556b033b2e3c9fd0803ce80000006003556100a06401000000006103a56100a982021704565b600655506100ad565b4290565b6108fd806100bc6000396000f3006060604052600436106100cf5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630302c68881146100d457806313af4035146100e9578063143e55e0146101085780633a4a42331461012d578063495d32cb1461014057806355deb8fc146101535780635d6542af146101695780637a9e5e4b1461017c5780638da5cb5b1461019b57806392b0d721146101ca578063a551878e146101e3578063becda0ea146101f6578063bf7e214f1461020c578063cfc4af551461021f575b600080fd5b34156100df57600080fd5b6100e7610232565b005b34156100f457600080fd5b6100e7600160a060020a0360043516610326565b341561011357600080fd5b61011b6103a5565b60405190815260200160405180910390f35b341561013857600080fd5b61011b6103a9565b341561014b57600080fd5b61011b6103af565b341561015e57600080fd5b6100e76004356103c0565b341561017457600080fd5b61011b61043e565b341561018757600080fd5b6100e7600160a060020a036004351661044f565b34156101a657600080fd5b6101ae6104ce565b604051600160a060020a03909116815260200160405180910390f35b34156101d557600080fd5b6100e76004356024356104dd565b34156101ee57600080fd5b61011b610584565b341561020157600080fd5b6100e760043561058a565b341561021757600080fd5b6101ae610608565b341561022a57600080fd5b61011b610617565b6000806004356024358082600160a060020a033316600160e060020a031986351634873660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46006546102936103a5565b0393508315156102a257610320565b6102aa6103a5565b6006556003546b033b2e3c9fd0803ce8000000146102dd576102d96002546102d46003548761061d565b61067f565b6002555b60055415156102eb57610320565b8360055402925061031c60025460045410610309578360000361030b565b835b6103166003546106c2565b01610714565b6003555b50505050565b61033c33600035600160e060020a031916610765565b151561034757600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b4290565b60055481565b60006103b9610232565b5060025490565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a461042c33600035600160e060020a031916610765565b151561043757600080fd5b5050600455565b6000610448610232565b5060035490565b61046533600035600160e060020a031916610765565b151561047057600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600154600160a060020a031681565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a461054933600035600160e060020a031916610765565b151561055457600080fd5b7f776179000000000000000000000000000000000000000000000000000000000084141561032057505060035550565b60045481565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46105f633600035600160e060020a031916610765565b151561060157600080fd5b5050600555565b600054600160a060020a031681565b60065481565b600060028206151561063b576b033b2e3c9fd0803ce800000061063d565b825b90506002820491505b811561067957610656838461067f565b9250600282061561066e5761066b818461067f565b90505b600282049150610646565b92915050565b60006b033b2e3c9fd0803ce80000006106b161069b8585610875565b60026b033b2e3c9fd0803ce80000005b0461089d565b8115156106ba57fe5b049392505050565b60006b033b2e3c9fd0803ce8000000821015610701576106ee6b033b2e3c9fd0803ce8000000836108ad565b6b033b2e3c9fd0803ce800000003610679565b506b033b2e3c9fd0803ce7ffffff190190565b60008082600f0b12156107505761074b6b033b2e3c9fd0803ce800000083600003600f0b6b033b2e3c9fd0803ce8000000016108ad565b610679565b50600f0b6b033b2e3c9fd0803ce80000000190565b600030600160a060020a031683600160a060020a0316141561078957506001610679565b600154600160a060020a03848116911614156107a757506001610679565b600054600160a060020a031615156107c157506000610679565b60008054600160a060020a03169063b700961390859030908690604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561085357600080fd5b6102c65a03f1151561086457600080fd5b505050604051805190509050610679565b600081158061089257505080820282828281151561088f57fe5b04145b151561067957600080fd5b8082018281101561067957600080fd5b6000816106b16108c9856b033b2e3c9fd0803ce8000000610875565b6002856106ab5600a165627a7a723058208d7656ba2bb4b14e487cd38a7cee380c6b119e5ea35fd750561db93b80b2f7950029"

// DeploySaiVox deploys a new Ethereum contract, binding an instance of SaiVox to it.
func DeploySaiVox(auth *bind.TransactOpts, backend bind.ContractBackend, par_ *big.Int) (common.Address, *types.Transaction, *SaiVox, error) {
	parsed, err := abi.JSON(strings.NewReader(SaiVoxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaiVoxBin), backend, par_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SaiVox{SaiVoxCaller: SaiVoxCaller{contract: contract}, SaiVoxTransactor: SaiVoxTransactor{contract: contract}, SaiVoxFilterer: SaiVoxFilterer{contract: contract}}, nil
}

// SaiVox is an auto generated Go binding around an Ethereum contract.
type SaiVox struct {
	SaiVoxCaller     // Read-only binding to the contract
	SaiVoxTransactor // Write-only binding to the contract
	SaiVoxFilterer   // Log filterer for contract events
}

// SaiVoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaiVoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiVoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaiVoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiVoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaiVoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaiVoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaiVoxSession struct {
	Contract     *SaiVox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaiVoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaiVoxCallerSession struct {
	Contract *SaiVoxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SaiVoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaiVoxTransactorSession struct {
	Contract     *SaiVoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaiVoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaiVoxRaw struct {
	Contract *SaiVox // Generic contract binding to access the raw methods on
}

// SaiVoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaiVoxCallerRaw struct {
	Contract *SaiVoxCaller // Generic read-only contract binding to access the raw methods on
}

// SaiVoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaiVoxTransactorRaw struct {
	Contract *SaiVoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaiVox creates a new instance of SaiVox, bound to a specific deployed contract.
func NewSaiVox(address common.Address, backend bind.ContractBackend) (*SaiVox, error) {
	contract, err := bindSaiVox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SaiVox{SaiVoxCaller: SaiVoxCaller{contract: contract}, SaiVoxTransactor: SaiVoxTransactor{contract: contract}, SaiVoxFilterer: SaiVoxFilterer{contract: contract}}, nil
}

// NewSaiVoxCaller creates a new read-only instance of SaiVox, bound to a specific deployed contract.
func NewSaiVoxCaller(address common.Address, caller bind.ContractCaller) (*SaiVoxCaller, error) {
	contract, err := bindSaiVox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaiVoxCaller{contract: contract}, nil
}

// NewSaiVoxTransactor creates a new write-only instance of SaiVox, bound to a specific deployed contract.
func NewSaiVoxTransactor(address common.Address, transactor bind.ContractTransactor) (*SaiVoxTransactor, error) {
	contract, err := bindSaiVox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaiVoxTransactor{contract: contract}, nil
}

// NewSaiVoxFilterer creates a new log filterer instance of SaiVox, bound to a specific deployed contract.
func NewSaiVoxFilterer(address common.Address, filterer bind.ContractFilterer) (*SaiVoxFilterer, error) {
	contract, err := bindSaiVox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaiVoxFilterer{contract: contract}, nil
}

// bindSaiVox binds a generic wrapper to an already deployed contract.
func bindSaiVox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaiVoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaiVox *SaiVoxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaiVox.Contract.SaiVoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaiVox *SaiVoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaiVox.Contract.SaiVoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaiVox *SaiVoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaiVox.Contract.SaiVoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaiVox *SaiVoxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaiVox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaiVox *SaiVoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaiVox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaiVox *SaiVoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaiVox.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SaiVox *SaiVoxCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaiVox.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SaiVox *SaiVoxSession) Authority() (common.Address, error) {
	return _SaiVox.Contract.Authority(&_SaiVox.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SaiVox *SaiVoxCallerSession) Authority() (common.Address, error) {
	return _SaiVox.Contract.Authority(&_SaiVox.CallOpts)
}

// Era is a free data retrieval call binding the contract method 0x143e55e0.
//
// Solidity: function era() constant returns(uint256)
func (_SaiVox *SaiVoxCaller) Era(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaiVox.contract.Call(opts, out, "era")
	return *ret0, err
}

// Era is a free data retrieval call binding the contract method 0x143e55e0.
//
// Solidity: function era() constant returns(uint256)
func (_SaiVox *SaiVoxSession) Era() (*big.Int, error) {
	return _SaiVox.Contract.Era(&_SaiVox.CallOpts)
}

// Era is a free data retrieval call binding the contract method 0x143e55e0.
//
// Solidity: function era() constant returns(uint256)
func (_SaiVox *SaiVoxCallerSession) Era() (*big.Int, error) {
	return _SaiVox.Contract.Era(&_SaiVox.CallOpts)
}

// Fix is a free data retrieval call binding the contract method 0xa551878e.
//
// Solidity: function fix() constant returns(uint256)
func (_SaiVox *SaiVoxCaller) Fix(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaiVox.contract.Call(opts, out, "fix")
	return *ret0, err
}

// Fix is a free data retrieval call binding the contract method 0xa551878e.
//
// Solidity: function fix() constant returns(uint256)
func (_SaiVox *SaiVoxSession) Fix() (*big.Int, error) {
	return _SaiVox.Contract.Fix(&_SaiVox.CallOpts)
}

// Fix is a free data retrieval call binding the contract method 0xa551878e.
//
// Solidity: function fix() constant returns(uint256)
func (_SaiVox *SaiVoxCallerSession) Fix() (*big.Int, error) {
	return _SaiVox.Contract.Fix(&_SaiVox.CallOpts)
}

// How is a free data retrieval call binding the contract method 0x3a4a4233.
//
// Solidity: function how() constant returns(uint256)
func (_SaiVox *SaiVoxCaller) How(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaiVox.contract.Call(opts, out, "how")
	return *ret0, err
}

// How is a free data retrieval call binding the contract method 0x3a4a4233.
//
// Solidity: function how() constant returns(uint256)
func (_SaiVox *SaiVoxSession) How() (*big.Int, error) {
	return _SaiVox.Contract.How(&_SaiVox.CallOpts)
}

// How is a free data retrieval call binding the contract method 0x3a4a4233.
//
// Solidity: function how() constant returns(uint256)
func (_SaiVox *SaiVoxCallerSession) How() (*big.Int, error) {
	return _SaiVox.Contract.How(&_SaiVox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaiVox *SaiVoxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaiVox.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaiVox *SaiVoxSession) Owner() (common.Address, error) {
	return _SaiVox.Contract.Owner(&_SaiVox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SaiVox *SaiVoxCallerSession) Owner() (common.Address, error) {
	return _SaiVox.Contract.Owner(&_SaiVox.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint256)
func (_SaiVox *SaiVoxCaller) Tau(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaiVox.contract.Call(opts, out, "tau")
	return *ret0, err
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint256)
func (_SaiVox *SaiVoxSession) Tau() (*big.Int, error) {
	return _SaiVox.Contract.Tau(&_SaiVox.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() constant returns(uint256)
func (_SaiVox *SaiVoxCallerSession) Tau() (*big.Int, error) {
	return _SaiVox.Contract.Tau(&_SaiVox.CallOpts)
}

// Mold is a paid mutator transaction binding the contract method 0x92b0d721.
//
// Solidity: function mold(bytes32 param, uint256 val) returns()
func (_SaiVox *SaiVoxTransactor) Mold(opts *bind.TransactOpts, param [32]byte, val *big.Int) (*types.Transaction, error) {
	return _SaiVox.contract.Transact(opts, "mold", param, val)
}

// Mold is a paid mutator transaction binding the contract method 0x92b0d721.
//
// Solidity: function mold(bytes32 param, uint256 val) returns()
func (_SaiVox *SaiVoxSession) Mold(param [32]byte, val *big.Int) (*types.Transaction, error) {
	return _SaiVox.Contract.Mold(&_SaiVox.TransactOpts, param, val)
}

// Mold is a paid mutator transaction binding the contract method 0x92b0d721.
//
// Solidity: function mold(bytes32 param, uint256 val) returns()
func (_SaiVox *SaiVoxTransactorSession) Mold(param [32]byte, val *big.Int) (*types.Transaction, error) {
	return _SaiVox.Contract.Mold(&_SaiVox.TransactOpts, param, val)
}

// Par is a paid mutator transaction binding the contract method 0x495d32cb.
//
// Solidity: function par() returns(uint256)
func (_SaiVox *SaiVoxTransactor) Par(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaiVox.contract.Transact(opts, "par")
}

// Par is a paid mutator transaction binding the contract method 0x495d32cb.
//
// Solidity: function par() returns(uint256)
func (_SaiVox *SaiVoxSession) Par() (*types.Transaction, error) {
	return _SaiVox.Contract.Par(&_SaiVox.TransactOpts)
}

// Par is a paid mutator transaction binding the contract method 0x495d32cb.
//
// Solidity: function par() returns(uint256)
func (_SaiVox *SaiVoxTransactorSession) Par() (*types.Transaction, error) {
	return _SaiVox.Contract.Par(&_SaiVox.TransactOpts)
}

// Prod is a paid mutator transaction binding the contract method 0x0302c688.
//
// Solidity: function prod() returns()
func (_SaiVox *SaiVoxTransactor) Prod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaiVox.contract.Transact(opts, "prod")
}

// Prod is a paid mutator transaction binding the contract method 0x0302c688.
//
// Solidity: function prod() returns()
func (_SaiVox *SaiVoxSession) Prod() (*types.Transaction, error) {
	return _SaiVox.Contract.Prod(&_SaiVox.TransactOpts)
}

// Prod is a paid mutator transaction binding the contract method 0x0302c688.
//
// Solidity: function prod() returns()
func (_SaiVox *SaiVoxTransactorSession) Prod() (*types.Transaction, error) {
	return _SaiVox.Contract.Prod(&_SaiVox.TransactOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_SaiVox *SaiVoxTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _SaiVox.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_SaiVox *SaiVoxSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _SaiVox.Contract.SetAuthority(&_SaiVox.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_SaiVox *SaiVoxTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _SaiVox.Contract.SetAuthority(&_SaiVox.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_SaiVox *SaiVoxTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _SaiVox.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_SaiVox *SaiVoxSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SaiVox.Contract.SetOwner(&_SaiVox.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_SaiVox *SaiVoxTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SaiVox.Contract.SetOwner(&_SaiVox.TransactOpts, owner_)
}

// Tell is a paid mutator transaction binding the contract method 0x55deb8fc.
//
// Solidity: function tell(uint256 ray) returns()
func (_SaiVox *SaiVoxTransactor) Tell(opts *bind.TransactOpts, ray *big.Int) (*types.Transaction, error) {
	return _SaiVox.contract.Transact(opts, "tell", ray)
}

// Tell is a paid mutator transaction binding the contract method 0x55deb8fc.
//
// Solidity: function tell(uint256 ray) returns()
func (_SaiVox *SaiVoxSession) Tell(ray *big.Int) (*types.Transaction, error) {
	return _SaiVox.Contract.Tell(&_SaiVox.TransactOpts, ray)
}

// Tell is a paid mutator transaction binding the contract method 0x55deb8fc.
//
// Solidity: function tell(uint256 ray) returns()
func (_SaiVox *SaiVoxTransactorSession) Tell(ray *big.Int) (*types.Transaction, error) {
	return _SaiVox.Contract.Tell(&_SaiVox.TransactOpts, ray)
}

// Tune is a paid mutator transaction binding the contract method 0xbecda0ea.
//
// Solidity: function tune(uint256 ray) returns()
func (_SaiVox *SaiVoxTransactor) Tune(opts *bind.TransactOpts, ray *big.Int) (*types.Transaction, error) {
	return _SaiVox.contract.Transact(opts, "tune", ray)
}

// Tune is a paid mutator transaction binding the contract method 0xbecda0ea.
//
// Solidity: function tune(uint256 ray) returns()
func (_SaiVox *SaiVoxSession) Tune(ray *big.Int) (*types.Transaction, error) {
	return _SaiVox.Contract.Tune(&_SaiVox.TransactOpts, ray)
}

// Tune is a paid mutator transaction binding the contract method 0xbecda0ea.
//
// Solidity: function tune(uint256 ray) returns()
func (_SaiVox *SaiVoxTransactorSession) Tune(ray *big.Int) (*types.Transaction, error) {
	return _SaiVox.Contract.Tune(&_SaiVox.TransactOpts, ray)
}

// Way is a paid mutator transaction binding the contract method 0x5d6542af.
//
// Solidity: function way() returns(uint256)
func (_SaiVox *SaiVoxTransactor) Way(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaiVox.contract.Transact(opts, "way")
}

// Way is a paid mutator transaction binding the contract method 0x5d6542af.
//
// Solidity: function way() returns(uint256)
func (_SaiVox *SaiVoxSession) Way() (*types.Transaction, error) {
	return _SaiVox.Contract.Way(&_SaiVox.TransactOpts)
}

// Way is a paid mutator transaction binding the contract method 0x5d6542af.
//
// Solidity: function way() returns(uint256)
func (_SaiVox *SaiVoxTransactorSession) Way() (*types.Transaction, error) {
	return _SaiVox.Contract.Way(&_SaiVox.TransactOpts)
}

// SaiVoxLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the SaiVox contract.
type SaiVoxLogSetAuthorityIterator struct {
	Event *SaiVoxLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *SaiVoxLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaiVoxLogSetAuthority)
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
		it.Event = new(SaiVoxLogSetAuthority)
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
func (it *SaiVoxLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaiVoxLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaiVoxLogSetAuthority represents a LogSetAuthority event raised by the SaiVox contract.
type SaiVoxLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_SaiVox *SaiVoxFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*SaiVoxLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _SaiVox.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &SaiVoxLogSetAuthorityIterator{contract: _SaiVox.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_SaiVox *SaiVoxFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *SaiVoxLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _SaiVox.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaiVoxLogSetAuthority)
				if err := _SaiVox.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_SaiVox *SaiVoxFilterer) ParseLogSetAuthority(log types.Log) (*SaiVoxLogSetAuthority, error) {
	event := new(SaiVoxLogSetAuthority)
	if err := _SaiVox.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SaiVoxLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the SaiVox contract.
type SaiVoxLogSetOwnerIterator struct {
	Event *SaiVoxLogSetOwner // Event containing the contract specifics and raw log

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
func (it *SaiVoxLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaiVoxLogSetOwner)
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
		it.Event = new(SaiVoxLogSetOwner)
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
func (it *SaiVoxLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaiVoxLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaiVoxLogSetOwner represents a LogSetOwner event raised by the SaiVox contract.
type SaiVoxLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_SaiVox *SaiVoxFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*SaiVoxLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SaiVox.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SaiVoxLogSetOwnerIterator{contract: _SaiVox.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_SaiVox *SaiVoxFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *SaiVoxLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SaiVox.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaiVoxLogSetOwner)
				if err := _SaiVox.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_SaiVox *SaiVoxFilterer) ParseLogSetOwner(log types.Log) (*SaiVoxLogSetOwner, error) {
	event := new(SaiVoxLogSetOwner)
	if err := _SaiVox.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}
