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

// DSGuardABI is the input ABI used to generate the binding from.
const DSGuardABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"sig\",\"type\":\"bytes32\"}],\"name\":\"forbid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"bytes32\"},{\"name\":\"dst\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes32\"}],\"name\":\"forbid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ANY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src_\",\"type\":\"address\"},{\"name\":\"dst_\",\"type\":\"address\"},{\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"sig\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"bytes32\"},{\"name\":\"dst\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes32\"}],\"name\":\"LogPermit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes32\"}],\"name\":\"LogForbid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// DSGuardBin is the compiled bytecode used for deploying new contracts.
const DSGuardBin = `6060604090815260018054600160a060020a03191633600160a060020a0316908117909155907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94905160405180910390a26107b98061005f6000396000f3006060604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af403581146100a85780632bc3217d146100c957806379d88d87146100f15780637a9e5e4b1461010d5780638da5cb5b1461012c578063a8542f661461015b578063b700961314610180578063bf7e214f146101c6578063cbeea68c146101d9578063f0217ce514610201575b600080fd5b34156100b357600080fd5b6100c7600160a060020a036004351661021d565b005b34156100d457600080fd5b6100c7600160a060020a036004358116906024351660443561029c565b34156100fc57600080fd5b6100c76004356024356044356102b8565b341561011857600080fd5b6100c7600160a060020a036004351661033a565b341561013757600080fd5b61013f6103b9565b604051600160a060020a03909116815260200160405180910390f35b341561016657600080fd5b61016e6103c8565b60405190815260200160405180910390f35b341561018b57600080fd5b6101b2600160a060020a0360043581169060243516600160e060020a0319604435166103ce565b604051901515815260200160405180910390f35b34156101d157600080fd5b61013f6105d5565b34156101e457600080fd5b6100c7600160a060020a03600435811690602435166044356105e4565b341561020c57600080fd5b6100c76004356024356044356105f7565b61023333600035600160e060020a03191661067c565b151561023e57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b6102b3600160a060020a03808516908416836102b8565b505050565b6102ce33600035600160e060020a03191661067c565b15156102d957600080fd5b6000838152600260209081526040808320858452825280832084845290915290819020805460ff191690558190839085907f95ba64c95d85e67ac83a0476c4a62ac2cf8ab2d0407545b8c9d79c3eefa62829905160405180910390a4505050565b61035033600035600160e060020a03191661067c565b151561035b57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600154600160a060020a031681565b60001981565b600160a060020a038084166000818152600260209081526040808320948716808452948252808320600160e060020a03198716845290915281205490929060ff168061043d575060008281526002602090815260408083208484528252808320600019845290915290205460ff165b806104755750600082815260026020908152604080832060001984528252808320600160e060020a03198816845290915290205460ff165b806104a0575060008281526002602090815260408083206000198452825280832090915290205460ff165b806104ed575060008181527f38b5b2ceac7637132d27514ffcf440b705287635075af7b8bd5adcaa6a4cc5bb60209081526040808320600160e060020a03198816845290915290205460ff165b80610532575060008181527f38b5b2ceac7637132d27514ffcf440b705287635075af7b8bd5adcaa6a4cc5bb60209081526040808320600019845290915290205460ff165b806105755750600160e060020a0319841660009081527f47fa60fbc027ac3984ea309688a33182f4193c478b40ba8d294eb2cd3ddc4d97602052604090205460ff165b806105cb57506000196000527f47fa60fbc027ac3984ea309688a33182f4193c478b40ba8d294eb2cd3ddc4d976020527ff423d1317b37667cd26005728bffa7c8b0499e133951fcf8e814d4fc5f4c98f65460ff165b9695505050505050565b600054600160a060020a031681565b6102b3600160a060020a03808516908416835b61060d33600035600160e060020a03191661067c565b151561061857600080fd5b6000838152600260209081526040808320858452825280832084845290915290819020805460ff191660011790558190839085907f6f50375045128971c5469d343039ba7b8e30a5b190453737b28bda6f7a306771905160405180910390a4505050565b600030600160a060020a031683600160a060020a031614156106a057506001610787565b600154600160a060020a03848116911614156106be57506001610787565b600054600160a060020a031615156106d857506000610787565b60008054600160a060020a03169063b700961390859030908690604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561076a57600080fd5b6102c65a03f1151561077b57600080fd5b50505060405180519150505b929150505600a165627a7a72305820ced89fd7fadac3974d383c0d3fcef1aff9a47bd4e51dc92e06cd673dd2a87a5b0029`

// DeployDSGuard deploys a new Ethereum contract, binding an instance of DSGuard to it.
func DeployDSGuard(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSGuard, error) {
	parsed, err := abi.JSON(strings.NewReader(DSGuardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSGuardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSGuard{DSGuardCaller: DSGuardCaller{contract: contract}, DSGuardTransactor: DSGuardTransactor{contract: contract}, DSGuardFilterer: DSGuardFilterer{contract: contract}}, nil
}

// DSGuard is an auto generated Go binding around an Ethereum contract.
type DSGuard struct {
	DSGuardCaller     // Read-only binding to the contract
	DSGuardTransactor // Write-only binding to the contract
	DSGuardFilterer   // Log filterer for contract events
}

// DSGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSGuardSession struct {
	Contract     *DSGuard          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSGuardCallerSession struct {
	Contract *DSGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DSGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSGuardTransactorSession struct {
	Contract     *DSGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DSGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSGuardRaw struct {
	Contract *DSGuard // Generic contract binding to access the raw methods on
}

// DSGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSGuardCallerRaw struct {
	Contract *DSGuardCaller // Generic read-only contract binding to access the raw methods on
}

// DSGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSGuardTransactorRaw struct {
	Contract *DSGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSGuard creates a new instance of DSGuard, bound to a specific deployed contract.
func NewDSGuard(address common.Address, backend bind.ContractBackend) (*DSGuard, error) {
	contract, err := bindDSGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSGuard{DSGuardCaller: DSGuardCaller{contract: contract}, DSGuardTransactor: DSGuardTransactor{contract: contract}, DSGuardFilterer: DSGuardFilterer{contract: contract}}, nil
}

// NewDSGuardCaller creates a new read-only instance of DSGuard, bound to a specific deployed contract.
func NewDSGuardCaller(address common.Address, caller bind.ContractCaller) (*DSGuardCaller, error) {
	contract, err := bindDSGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSGuardCaller{contract: contract}, nil
}

// NewDSGuardTransactor creates a new write-only instance of DSGuard, bound to a specific deployed contract.
func NewDSGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*DSGuardTransactor, error) {
	contract, err := bindDSGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSGuardTransactor{contract: contract}, nil
}

// NewDSGuardFilterer creates a new log filterer instance of DSGuard, bound to a specific deployed contract.
func NewDSGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*DSGuardFilterer, error) {
	contract, err := bindDSGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSGuardFilterer{contract: contract}, nil
}

// bindDSGuard binds a generic wrapper to an already deployed contract.
func bindDSGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSGuard *DSGuardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSGuard.Contract.DSGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSGuard *DSGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSGuard.Contract.DSGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSGuard *DSGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSGuard.Contract.DSGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSGuard *DSGuardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSGuard *DSGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSGuard *DSGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSGuard.Contract.contract.Transact(opts, method, params...)
}

// ANY is a free data retrieval call binding the contract method 0xa8542f66.
//
// Solidity: function ANY() constant returns(bytes32)
func (_DSGuard *DSGuardCaller) ANY(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DSGuard.contract.Call(opts, out, "ANY")
	return *ret0, err
}

// ANY is a free data retrieval call binding the contract method 0xa8542f66.
//
// Solidity: function ANY() constant returns(bytes32)
func (_DSGuard *DSGuardSession) ANY() ([32]byte, error) {
	return _DSGuard.Contract.ANY(&_DSGuard.CallOpts)
}

// ANY is a free data retrieval call binding the contract method 0xa8542f66.
//
// Solidity: function ANY() constant returns(bytes32)
func (_DSGuard *DSGuardCallerSession) ANY() ([32]byte, error) {
	return _DSGuard.Contract.ANY(&_DSGuard.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSGuard *DSGuardCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGuard.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSGuard *DSGuardSession) Authority() (common.Address, error) {
	return _DSGuard.Contract.Authority(&_DSGuard.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSGuard *DSGuardCallerSession) Authority() (common.Address, error) {
	return _DSGuard.Contract.Authority(&_DSGuard.CallOpts)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address src_, address dst_, bytes4 sig) constant returns(bool)
func (_DSGuard *DSGuardCaller) CanCall(opts *bind.CallOpts, src_ common.Address, dst_ common.Address, sig [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DSGuard.contract.Call(opts, out, "canCall", src_, dst_, sig)
	return *ret0, err
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address src_, address dst_, bytes4 sig) constant returns(bool)
func (_DSGuard *DSGuardSession) CanCall(src_ common.Address, dst_ common.Address, sig [4]byte) (bool, error) {
	return _DSGuard.Contract.CanCall(&_DSGuard.CallOpts, src_, dst_, sig)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address src_, address dst_, bytes4 sig) constant returns(bool)
func (_DSGuard *DSGuardCallerSession) CanCall(src_ common.Address, dst_ common.Address, sig [4]byte) (bool, error) {
	return _DSGuard.Contract.CanCall(&_DSGuard.CallOpts, src_, dst_, sig)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSGuard *DSGuardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGuard.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSGuard *DSGuardSession) Owner() (common.Address, error) {
	return _DSGuard.Contract.Owner(&_DSGuard.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSGuard *DSGuardCallerSession) Owner() (common.Address, error) {
	return _DSGuard.Contract.Owner(&_DSGuard.CallOpts)
}

// Forbid is a paid mutator transaction binding the contract method 0x79d88d87.
//
// Solidity: function forbid(bytes32 src, bytes32 dst, bytes32 sig) returns()
func (_DSGuard *DSGuardTransactor) Forbid(opts *bind.TransactOpts, src [32]byte, dst [32]byte, sig [32]byte) (*types.Transaction, error) {
	return _DSGuard.contract.Transact(opts, "forbid", src, dst, sig)
}

// Forbid is a paid mutator transaction binding the contract method 0x79d88d87.
//
// Solidity: function forbid(bytes32 src, bytes32 dst, bytes32 sig) returns()
func (_DSGuard *DSGuardSession) Forbid(src [32]byte, dst [32]byte, sig [32]byte) (*types.Transaction, error) {
	return _DSGuard.Contract.Forbid(&_DSGuard.TransactOpts, src, dst, sig)
}

// Forbid is a paid mutator transaction binding the contract method 0x79d88d87.
//
// Solidity: function forbid(bytes32 src, bytes32 dst, bytes32 sig) returns()
func (_DSGuard *DSGuardTransactorSession) Forbid(src [32]byte, dst [32]byte, sig [32]byte) (*types.Transaction, error) {
	return _DSGuard.Contract.Forbid(&_DSGuard.TransactOpts, src, dst, sig)
}

// Permit is a paid mutator transaction binding the contract method 0xf0217ce5.
//
// Solidity: function permit(bytes32 src, bytes32 dst, bytes32 sig) returns()
func (_DSGuard *DSGuardTransactor) Permit(opts *bind.TransactOpts, src [32]byte, dst [32]byte, sig [32]byte) (*types.Transaction, error) {
	return _DSGuard.contract.Transact(opts, "permit", src, dst, sig)
}

// Permit is a paid mutator transaction binding the contract method 0xf0217ce5.
//
// Solidity: function permit(bytes32 src, bytes32 dst, bytes32 sig) returns()
func (_DSGuard *DSGuardSession) Permit(src [32]byte, dst [32]byte, sig [32]byte) (*types.Transaction, error) {
	return _DSGuard.Contract.Permit(&_DSGuard.TransactOpts, src, dst, sig)
}

// Permit is a paid mutator transaction binding the contract method 0xf0217ce5.
//
// Solidity: function permit(bytes32 src, bytes32 dst, bytes32 sig) returns()
func (_DSGuard *DSGuardTransactorSession) Permit(src [32]byte, dst [32]byte, sig [32]byte) (*types.Transaction, error) {
	return _DSGuard.Contract.Permit(&_DSGuard.TransactOpts, src, dst, sig)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_DSGuard *DSGuardTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _DSGuard.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_DSGuard *DSGuardSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSGuard.Contract.SetAuthority(&_DSGuard.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_DSGuard *DSGuardTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSGuard.Contract.SetAuthority(&_DSGuard.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_DSGuard *DSGuardTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _DSGuard.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_DSGuard *DSGuardSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSGuard.Contract.SetOwner(&_DSGuard.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_DSGuard *DSGuardTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSGuard.Contract.SetOwner(&_DSGuard.TransactOpts, owner_)
}

// DSGuardLogForbidIterator is returned from FilterLogForbid and is used to iterate over the raw logs and unpacked data for LogForbid events raised by the DSGuard contract.
type DSGuardLogForbidIterator struct {
	Event *DSGuardLogForbid // Event containing the contract specifics and raw log

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
func (it *DSGuardLogForbidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSGuardLogForbid)
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
		it.Event = new(DSGuardLogForbid)
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
func (it *DSGuardLogForbidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSGuardLogForbidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSGuardLogForbid represents a LogForbid event raised by the DSGuard contract.
type DSGuardLogForbid struct {
	Src [32]byte
	Dst [32]byte
	Sig [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogForbid is a free log retrieval operation binding the contract event 0x95ba64c95d85e67ac83a0476c4a62ac2cf8ab2d0407545b8c9d79c3eefa62829.
//
// Solidity: event LogForbid(bytes32 indexed src, bytes32 indexed dst, bytes32 indexed sig)
func (_DSGuard *DSGuardFilterer) FilterLogForbid(opts *bind.FilterOpts, src [][32]byte, dst [][32]byte, sig [][32]byte) (*DSGuardLogForbidIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}
	var sigRule []interface{}
	for _, sigItem := range sig {
		sigRule = append(sigRule, sigItem)
	}

	logs, sub, err := _DSGuard.contract.FilterLogs(opts, "LogForbid", srcRule, dstRule, sigRule)
	if err != nil {
		return nil, err
	}
	return &DSGuardLogForbidIterator{contract: _DSGuard.contract, event: "LogForbid", logs: logs, sub: sub}, nil
}

// WatchLogForbid is a free log subscription operation binding the contract event 0x95ba64c95d85e67ac83a0476c4a62ac2cf8ab2d0407545b8c9d79c3eefa62829.
//
// Solidity: event LogForbid(bytes32 indexed src, bytes32 indexed dst, bytes32 indexed sig)
func (_DSGuard *DSGuardFilterer) WatchLogForbid(opts *bind.WatchOpts, sink chan<- *DSGuardLogForbid, src [][32]byte, dst [][32]byte, sig [][32]byte) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}
	var sigRule []interface{}
	for _, sigItem := range sig {
		sigRule = append(sigRule, sigItem)
	}

	logs, sub, err := _DSGuard.contract.WatchLogs(opts, "LogForbid", srcRule, dstRule, sigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSGuardLogForbid)
				if err := _DSGuard.contract.UnpackLog(event, "LogForbid", log); err != nil {
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

// DSGuardLogPermitIterator is returned from FilterLogPermit and is used to iterate over the raw logs and unpacked data for LogPermit events raised by the DSGuard contract.
type DSGuardLogPermitIterator struct {
	Event *DSGuardLogPermit // Event containing the contract specifics and raw log

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
func (it *DSGuardLogPermitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSGuardLogPermit)
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
		it.Event = new(DSGuardLogPermit)
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
func (it *DSGuardLogPermitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSGuardLogPermitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSGuardLogPermit represents a LogPermit event raised by the DSGuard contract.
type DSGuardLogPermit struct {
	Src [32]byte
	Dst [32]byte
	Sig [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogPermit is a free log retrieval operation binding the contract event 0x6f50375045128971c5469d343039ba7b8e30a5b190453737b28bda6f7a306771.
//
// Solidity: event LogPermit(bytes32 indexed src, bytes32 indexed dst, bytes32 indexed sig)
func (_DSGuard *DSGuardFilterer) FilterLogPermit(opts *bind.FilterOpts, src [][32]byte, dst [][32]byte, sig [][32]byte) (*DSGuardLogPermitIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}
	var sigRule []interface{}
	for _, sigItem := range sig {
		sigRule = append(sigRule, sigItem)
	}

	logs, sub, err := _DSGuard.contract.FilterLogs(opts, "LogPermit", srcRule, dstRule, sigRule)
	if err != nil {
		return nil, err
	}
	return &DSGuardLogPermitIterator{contract: _DSGuard.contract, event: "LogPermit", logs: logs, sub: sub}, nil
}

// WatchLogPermit is a free log subscription operation binding the contract event 0x6f50375045128971c5469d343039ba7b8e30a5b190453737b28bda6f7a306771.
//
// Solidity: event LogPermit(bytes32 indexed src, bytes32 indexed dst, bytes32 indexed sig)
func (_DSGuard *DSGuardFilterer) WatchLogPermit(opts *bind.WatchOpts, sink chan<- *DSGuardLogPermit, src [][32]byte, dst [][32]byte, sig [][32]byte) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}
	var sigRule []interface{}
	for _, sigItem := range sig {
		sigRule = append(sigRule, sigItem)
	}

	logs, sub, err := _DSGuard.contract.WatchLogs(opts, "LogPermit", srcRule, dstRule, sigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSGuardLogPermit)
				if err := _DSGuard.contract.UnpackLog(event, "LogPermit", log); err != nil {
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

// DSGuardLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the DSGuard contract.
type DSGuardLogSetAuthorityIterator struct {
	Event *DSGuardLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *DSGuardLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSGuardLogSetAuthority)
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
		it.Event = new(DSGuardLogSetAuthority)
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
func (it *DSGuardLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSGuardLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSGuardLogSetAuthority represents a LogSetAuthority event raised by the DSGuard contract.
type DSGuardLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_DSGuard *DSGuardFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*DSGuardLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _DSGuard.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &DSGuardLogSetAuthorityIterator{contract: _DSGuard.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_DSGuard *DSGuardFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *DSGuardLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _DSGuard.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSGuardLogSetAuthority)
				if err := _DSGuard.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// DSGuardLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the DSGuard contract.
type DSGuardLogSetOwnerIterator struct {
	Event *DSGuardLogSetOwner // Event containing the contract specifics and raw log

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
func (it *DSGuardLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSGuardLogSetOwner)
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
		it.Event = new(DSGuardLogSetOwner)
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
func (it *DSGuardLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSGuardLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSGuardLogSetOwner represents a LogSetOwner event raised by the DSGuard contract.
type DSGuardLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_DSGuard *DSGuardFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*DSGuardLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSGuard.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DSGuardLogSetOwnerIterator{contract: _DSGuard.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_DSGuard *DSGuardFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *DSGuardLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSGuard.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSGuardLogSetOwner)
				if err := _DSGuard.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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
