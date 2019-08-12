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

// MedianizerNewABI is the input ABI used to generate the binding from.
const MedianizerNewABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compute\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"unset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"name\":\"values\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min_\",\"type\":\"uint96\"}],\"name\":\"setMin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"bytes12\"},{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"bytes12\"}],\"name\":\"unset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next_\",\"type\":\"bytes12\"}],\"name\":\"setNext\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"min\",\"outputs\":[{\"name\":\"\",\"type\":\"uint96\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"LogValue\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// MedianizerNewBin is the compiled bytecode used for deploying new contracts.
const MedianizerNewBin = `60606040908152600580546c0100000000000000000000000060016001606060020a03199092168217606060020a60c060020a031916179091558054600160a060020a03191633600160a060020a0316908117909155907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94905160405180910390a26110f6806100906000396000f3006060604052600436106100f85763ffffffff60e060020a60003504166313af403581146100fd5780631504460f1461011e57806318178358146101345780631a43c338146101475780632801617e146101745780632966d1b9146101935780632db78d93146101b25780634c8fe526146101ee57806357de26a41461020157806359e02dd714610226578063651dd0de146102395780636ba5ef0d146102755780637a9e5e4b146102945780638da5cb5b146102b3578063ac4c25b2146102c6578063beb38b43146102d9578063bf7e214f14610305578063e0a1fdad14610318578063f2c5925d14610338578063f889794514610358575b600080fd5b341561010857600080fd5b61011c600160a060020a0360043516610387565b005b341561012957600080fd5b61011c6004356103f9565b341561013f57600080fd5b61011c6104c1565b341561015257600080fd5b61015a6104cd565b604051918252151560208201526040908101905180910390f35b341561017f57600080fd5b61011c600160a060020a0360043516610875565b341561019e57600080fd5b61011c600160a060020a036004351661095f565b34156101bd57600080fd5b6101d1600160a060020a0360043516610a11565b604051600160a060020a0319909116815260200160405180910390f35b34156101f957600080fd5b6101d1610a29565b341561020c57600080fd5b610214610a35565b60405190815260200160405180910390f35b341561023157600080fd5b61015a610a55565b341561024457600080fd5b610259600160a060020a031960043516610a6a565b604051600160a060020a03909116815260200160405180910390f35b341561028057600080fd5b61011c6001606060020a0360043516610a85565b341561029f57600080fd5b61011c600160a060020a0360043516610b56565b34156102be57600080fd5b610259610bc8565b34156102d157600080fd5b61011c610bd7565b34156102e457600080fd5b61011c600160a060020a031960043516602435600160a060020a0316610c70565b341561031057600080fd5b610259610df6565b341561032357600080fd5b61011c600160a060020a031960043516610e05565b341561034357600080fd5b61011c600160a060020a031960043516610e85565b341561036357600080fd5b61036b610f37565b6040516001606060020a03909116815260200160405180910390f35b61039d33600035600160e060020a031916610f56565b15156103a857600080fd5b60018054600160a060020a031916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46104576104cd565b6001805491151560a060020a0274ff00000000000000000000000000000000000000001990921691909117905560028190557f296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b5279060405190815260200160405180910390a1505050565b6104cb60006103f9565b565b6000806104d86110b8565b60008060008060008060008060006001600560009054906101000a900460a060020a0260a060020a9004036001606060020a03166040518059106105195750595b9080825280602002602001820160405250995060009850600197505b60055460a060020a908102046001606060020a03908116908916101561074657600160a060020a031960a060020a890216600090815260036020526040902054600160a060020a03161561073b57600160a060020a031960a060020a89021660009081526003602052604080822054600160a060020a0316916359e02dd79151604001526040518163ffffffff1660e060020a0281526004016040805180830381600087803b15156105e657600080fd5b6102c65a03f115156105f757600080fd5b5050506040518051906020018051905096509650851561073b576001606060020a038916158061064857508960018a036001606060020a03168151811061063a57fe5b906020019060200201518710155b1561067257868a8a6001606060020a03168151811061066357fe5b60209081029091010152610734565b600094505b89856001606060020a03168151811061068c57fe5b9060200190602002015187106106a757600190940193610677565b8893505b846001606060020a0316846001606060020a031611156107135789600185036001606060020a0316815181106106dd57fe5b906020019060200201518a856001606060020a0316815181106106fc57fe5b6020908102909101015260001993909301926106ab565b868a866001606060020a03168151811061072957fe5b602090810290910101525b6001909801975b600190970196610535565b6005546001606060020a036c010000000000000000000000009091048116908a16101561077b5760025460009b509b50610867565b60026001606060020a038a16066001606060020a03166000141561082d5789600160026001606060020a038c1604036001606060020a0316815181106107bd57fe5b9060200190602002015191508960026001606060020a038b16046001606060020a0316815181106107ea57fe5b9060200190602002015190506108266108186fffffffffffffffffffffffffffffffff80851690841661104e565b671bc16d674ec8000061105e565b925061085f565b8960026001606060020a036000198c0116046001606060020a03168151811061085257fe5b9060200190602002015192505b8260019b509b505b505050505050505050509091565b600061088d33600035600160e060020a031916610f56565b151561089857600080fd5b5060055460a060020a90810281900460010102600160a060020a0319811615156108be57fe5b600554600160a060020a0330169063beb38b439060a060020a028460405160e060020a63ffffffff8516028152600160a060020a03199092166004830152600160a060020a03166024820152604401600060405180830381600087803b151561092657600080fd5b6102c65a03f1151561093757600080fd5b5050600580546bffffffffffffffffffffffff191660a060020a909304929092179091555050565b61097533600035600160e060020a031916610f56565b151561098057600080fd5b600160a060020a0381811660009081526004602052604080822054309093169263beb38b439260a060020a909102915160e060020a63ffffffff8516028152600160a060020a03199092166004830152600160a060020a03166024820152604401600060405180830381600087803b15156109fa57600080fd5b6102c65a03f11515610a0b57600080fd5b50505050565b60046020526000908152604090205460a060020a0281565b60055460a060020a0281565b6000806000610a42610a55565b91509150801515610a4f57fe5b50919050565b60025460015460ff60a060020a909104169091565b600360205260009081526040902054600160a060020a031681565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4610af133600035600160e060020a031916610f56565b1515610afc57600080fd5b6001606060020a0383161515610b1157600080fd5b5050600580546001606060020a039092166c010000000000000000000000000277ffffffffffffffffffffffff00000000000000000000000019909216919091179055565b610b6c33600035600160e060020a031916610f56565b1515610b7757600080fd5b60008054600160a060020a031916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600154600160a060020a031681565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4610c4333600035600160e060020a031916610f56565b1515610c4e57600080fd5b50506001805474ff000000000000000000000000000000000000000019169055565b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4610cdc33600035600160e060020a031916610f56565b1515610ce757600080fd5b600160a060020a031984161515610cfd57600080fd5b600160a060020a0383161580610d395750600160a060020a03831660009081526004602052604090205460a060020a02600160a060020a031916155b1515610d4457600080fd5b600160a060020a03198416600090815260036020908152604080832054600160a060020a039081168452600490925290912080546bffffffffffffffffffffffff19169055831615610dc457600160a060020a038316600090815260046020526040902080546bffffffffffffffffffffffff191660a060020a86041790555b5050600160a060020a031991821660009081526003602052604090208054600160a060020a0390921691909216179055565b600054600160a060020a031681565b610e1b33600035600160e060020a031916610f56565b1515610e2657600080fd5b30600160a060020a031663beb38b4382600060405160e060020a63ffffffff8516028152600160a060020a03199092166004830152600160a060020a03166024820152604401600060405180830381600087803b15156109fa57600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4610ef133600035600160e060020a031916610f56565b1515610efc57600080fd5b600160a060020a031983161515610f1257600080fd5b5050600580546bffffffffffffffffffffffff191660a060020a909204919091179055565b6005546c0100000000000000000000000090046001606060020a031681565b600030600160a060020a031683600160a060020a03161415610f7a57506001611048565b600154600160a060020a0384811691161415610f9857506001611048565b600054600160a060020a03161515610fb257506000611048565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561102b57600080fd5b6102c65a03f1151561103c57600080fd5b50505060405180519150505b92915050565b8082018281101561104857600080fd5b60008161107f61107685670de0b6b3a7640000611090565b6002850461104e565b81151561108857fe5b049392505050565b60008115806110ad5750508082028282828115156110aa57fe5b04145b151561104857600080fd5b602060405190810160405260008152905600a165627a7a723058207d363fa9788a53a39f510d086d9db96eec0d7bbd40cd50acb0d47e66b0c7ac900029`

// DeployMedianizerNew deploys a new Ethereum contract, binding an instance of MedianizerNew to it.
func DeployMedianizerNew(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MedianizerNew, error) {
	parsed, err := abi.JSON(strings.NewReader(MedianizerNewABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MedianizerNewBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MedianizerNew{MedianizerNewCaller: MedianizerNewCaller{contract: contract}, MedianizerNewTransactor: MedianizerNewTransactor{contract: contract}, MedianizerNewFilterer: MedianizerNewFilterer{contract: contract}}, nil
}

// MedianizerNew is an auto generated Go binding around an Ethereum contract.
type MedianizerNew struct {
	MedianizerNewCaller     // Read-only binding to the contract
	MedianizerNewTransactor // Write-only binding to the contract
	MedianizerNewFilterer   // Log filterer for contract events
}

// MedianizerNewCaller is an auto generated read-only Go binding around an Ethereum contract.
type MedianizerNewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerNewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MedianizerNewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerNewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MedianizerNewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerNewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MedianizerNewSession struct {
	Contract     *MedianizerNew    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedianizerNewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MedianizerNewCallerSession struct {
	Contract *MedianizerNewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MedianizerNewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MedianizerNewTransactorSession struct {
	Contract     *MedianizerNewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MedianizerNewRaw is an auto generated low-level Go binding around an Ethereum contract.
type MedianizerNewRaw struct {
	Contract *MedianizerNew // Generic contract binding to access the raw methods on
}

// MedianizerNewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MedianizerNewCallerRaw struct {
	Contract *MedianizerNewCaller // Generic read-only contract binding to access the raw methods on
}

// MedianizerNewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MedianizerNewTransactorRaw struct {
	Contract *MedianizerNewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMedianizerNew creates a new instance of MedianizerNew, bound to a specific deployed contract.
func NewMedianizerNew(address common.Address, backend bind.ContractBackend) (*MedianizerNew, error) {
	contract, err := bindMedianizerNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MedianizerNew{MedianizerNewCaller: MedianizerNewCaller{contract: contract}, MedianizerNewTransactor: MedianizerNewTransactor{contract: contract}, MedianizerNewFilterer: MedianizerNewFilterer{contract: contract}}, nil
}

// NewMedianizerNewCaller creates a new read-only instance of MedianizerNew, bound to a specific deployed contract.
func NewMedianizerNewCaller(address common.Address, caller bind.ContractCaller) (*MedianizerNewCaller, error) {
	contract, err := bindMedianizerNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MedianizerNewCaller{contract: contract}, nil
}

// NewMedianizerNewTransactor creates a new write-only instance of MedianizerNew, bound to a specific deployed contract.
func NewMedianizerNewTransactor(address common.Address, transactor bind.ContractTransactor) (*MedianizerNewTransactor, error) {
	contract, err := bindMedianizerNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MedianizerNewTransactor{contract: contract}, nil
}

// NewMedianizerNewFilterer creates a new log filterer instance of MedianizerNew, bound to a specific deployed contract.
func NewMedianizerNewFilterer(address common.Address, filterer bind.ContractFilterer) (*MedianizerNewFilterer, error) {
	contract, err := bindMedianizerNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MedianizerNewFilterer{contract: contract}, nil
}

// bindMedianizerNew binds a generic wrapper to an already deployed contract.
func bindMedianizerNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MedianizerNewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MedianizerNew *MedianizerNewRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MedianizerNew.Contract.MedianizerNewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MedianizerNew *MedianizerNewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MedianizerNew.Contract.MedianizerNewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MedianizerNew *MedianizerNewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MedianizerNew.Contract.MedianizerNewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MedianizerNew *MedianizerNewCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MedianizerNew.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MedianizerNew *MedianizerNewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MedianizerNew.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MedianizerNew *MedianizerNewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MedianizerNew.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_MedianizerNew *MedianizerNewCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MedianizerNew.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_MedianizerNew *MedianizerNewSession) Authority() (common.Address, error) {
	return _MedianizerNew.Contract.Authority(&_MedianizerNew.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_MedianizerNew *MedianizerNewCallerSession) Authority() (common.Address, error) {
	return _MedianizerNew.Contract.Authority(&_MedianizerNew.CallOpts)
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_MedianizerNew *MedianizerNewCaller) Compute(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MedianizerNew.contract.Call(opts, out, "compute")
	return *ret0, *ret1, err
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_MedianizerNew *MedianizerNewSession) Compute() ([32]byte, bool, error) {
	return _MedianizerNew.Contract.Compute(&_MedianizerNew.CallOpts)
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_MedianizerNew *MedianizerNewCallerSession) Compute() ([32]byte, bool, error) {
	return _MedianizerNew.Contract.Compute(&_MedianizerNew.CallOpts)
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes(address ) constant returns(bytes12)
func (_MedianizerNew *MedianizerNewCaller) Indexes(opts *bind.CallOpts, arg0 common.Address) ([12]byte, error) {
	var (
		ret0 = new([12]byte)
	)
	out := ret0
	err := _MedianizerNew.contract.Call(opts, out, "indexes", arg0)
	return *ret0, err
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes(address ) constant returns(bytes12)
func (_MedianizerNew *MedianizerNewSession) Indexes(arg0 common.Address) ([12]byte, error) {
	return _MedianizerNew.Contract.Indexes(&_MedianizerNew.CallOpts, arg0)
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes(address ) constant returns(bytes12)
func (_MedianizerNew *MedianizerNewCallerSession) Indexes(arg0 common.Address) ([12]byte, error) {
	return _MedianizerNew.Contract.Indexes(&_MedianizerNew.CallOpts, arg0)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_MedianizerNew *MedianizerNewCaller) Min(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MedianizerNew.contract.Call(opts, out, "min")
	return *ret0, err
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_MedianizerNew *MedianizerNewSession) Min() (*big.Int, error) {
	return _MedianizerNew.Contract.Min(&_MedianizerNew.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_MedianizerNew *MedianizerNewCallerSession) Min() (*big.Int, error) {
	return _MedianizerNew.Contract.Min(&_MedianizerNew.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_MedianizerNew *MedianizerNewCaller) Next(opts *bind.CallOpts) ([12]byte, error) {
	var (
		ret0 = new([12]byte)
	)
	out := ret0
	err := _MedianizerNew.contract.Call(opts, out, "next")
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_MedianizerNew *MedianizerNewSession) Next() ([12]byte, error) {
	return _MedianizerNew.Contract.Next(&_MedianizerNew.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_MedianizerNew *MedianizerNewCallerSession) Next() ([12]byte, error) {
	return _MedianizerNew.Contract.Next(&_MedianizerNew.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MedianizerNew *MedianizerNewCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MedianizerNew.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MedianizerNew *MedianizerNewSession) Owner() (common.Address, error) {
	return _MedianizerNew.Contract.Owner(&_MedianizerNew.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MedianizerNew *MedianizerNewCallerSession) Owner() (common.Address, error) {
	return _MedianizerNew.Contract.Owner(&_MedianizerNew.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_MedianizerNew *MedianizerNewCaller) Peek(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MedianizerNew.contract.Call(opts, out, "peek")
	return *ret0, *ret1, err
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_MedianizerNew *MedianizerNewSession) Peek() ([32]byte, bool, error) {
	return _MedianizerNew.Contract.Peek(&_MedianizerNew.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_MedianizerNew *MedianizerNewCallerSession) Peek() ([32]byte, bool, error) {
	return _MedianizerNew.Contract.Peek(&_MedianizerNew.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_MedianizerNew *MedianizerNewCaller) Read(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MedianizerNew.contract.Call(opts, out, "read")
	return *ret0, err
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_MedianizerNew *MedianizerNewSession) Read() ([32]byte, error) {
	return _MedianizerNew.Contract.Read(&_MedianizerNew.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_MedianizerNew *MedianizerNewCallerSession) Read() ([32]byte, error) {
	return _MedianizerNew.Contract.Read(&_MedianizerNew.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values(bytes12 ) constant returns(address)
func (_MedianizerNew *MedianizerNewCaller) Values(opts *bind.CallOpts, arg0 [12]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MedianizerNew.contract.Call(opts, out, "values", arg0)
	return *ret0, err
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values(bytes12 ) constant returns(address)
func (_MedianizerNew *MedianizerNewSession) Values(arg0 [12]byte) (common.Address, error) {
	return _MedianizerNew.Contract.Values(&_MedianizerNew.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values(bytes12 ) constant returns(address)
func (_MedianizerNew *MedianizerNewCallerSession) Values(arg0 [12]byte) (common.Address, error) {
	return _MedianizerNew.Contract.Values(&_MedianizerNew.CallOpts, arg0)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_MedianizerNew *MedianizerNewTransactor) Poke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MedianizerNew.contract.Transact(opts, "poke")
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_MedianizerNew *MedianizerNewSession) Poke() (*types.Transaction, error) {
	return _MedianizerNew.Contract.Poke(&_MedianizerNew.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_MedianizerNew *MedianizerNewTransactorSession) Poke() (*types.Transaction, error) {
	return _MedianizerNew.Contract.Poke(&_MedianizerNew.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(bytes12 pos, address wat) returns()
func (_MedianizerNew *MedianizerNewTransactor) Set(opts *bind.TransactOpts, pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _MedianizerNew.contract.Transact(opts, "set", pos, wat)
}

// Set is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(bytes12 pos, address wat) returns()
func (_MedianizerNew *MedianizerNewSession) Set(pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _MedianizerNew.Contract.Set(&_MedianizerNew.TransactOpts, pos, wat)
}

// Set is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(bytes12 pos, address wat) returns()
func (_MedianizerNew *MedianizerNewTransactorSession) Set(pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _MedianizerNew.Contract.Set(&_MedianizerNew.TransactOpts, pos, wat)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MedianizerNew *MedianizerNewTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _MedianizerNew.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MedianizerNew *MedianizerNewSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _MedianizerNew.Contract.SetAuthority(&_MedianizerNew.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MedianizerNew *MedianizerNewTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _MedianizerNew.Contract.SetAuthority(&_MedianizerNew.TransactOpts, authority_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(uint96 min_) returns()
func (_MedianizerNew *MedianizerNewTransactor) SetMin(opts *bind.TransactOpts, min_ *big.Int) (*types.Transaction, error) {
	return _MedianizerNew.contract.Transact(opts, "setMin", min_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(uint96 min_) returns()
func (_MedianizerNew *MedianizerNewSession) SetMin(min_ *big.Int) (*types.Transaction, error) {
	return _MedianizerNew.Contract.SetMin(&_MedianizerNew.TransactOpts, min_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(uint96 min_) returns()
func (_MedianizerNew *MedianizerNewTransactorSession) SetMin(min_ *big.Int) (*types.Transaction, error) {
	return _MedianizerNew.Contract.SetMin(&_MedianizerNew.TransactOpts, min_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(bytes12 next_) returns()
func (_MedianizerNew *MedianizerNewTransactor) SetNext(opts *bind.TransactOpts, next_ [12]byte) (*types.Transaction, error) {
	return _MedianizerNew.contract.Transact(opts, "setNext", next_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(bytes12 next_) returns()
func (_MedianizerNew *MedianizerNewSession) SetNext(next_ [12]byte) (*types.Transaction, error) {
	return _MedianizerNew.Contract.SetNext(&_MedianizerNew.TransactOpts, next_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(bytes12 next_) returns()
func (_MedianizerNew *MedianizerNewTransactorSession) SetNext(next_ [12]byte) (*types.Transaction, error) {
	return _MedianizerNew.Contract.SetNext(&_MedianizerNew.TransactOpts, next_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MedianizerNew *MedianizerNewTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _MedianizerNew.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MedianizerNew *MedianizerNewSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _MedianizerNew.Contract.SetOwner(&_MedianizerNew.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MedianizerNew *MedianizerNewTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _MedianizerNew.Contract.SetOwner(&_MedianizerNew.TransactOpts, owner_)
}

// Unset is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(bytes12 pos) returns()
func (_MedianizerNew *MedianizerNewTransactor) Unset(opts *bind.TransactOpts, pos [12]byte) (*types.Transaction, error) {
	return _MedianizerNew.contract.Transact(opts, "unset", pos)
}

// Unset is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(bytes12 pos) returns()
func (_MedianizerNew *MedianizerNewSession) Unset(pos [12]byte) (*types.Transaction, error) {
	return _MedianizerNew.Contract.Unset(&_MedianizerNew.TransactOpts, pos)
}

// Unset is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(bytes12 pos) returns()
func (_MedianizerNew *MedianizerNewTransactorSession) Unset(pos [12]byte) (*types.Transaction, error) {
	return _MedianizerNew.Contract.Unset(&_MedianizerNew.TransactOpts, pos)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_MedianizerNew *MedianizerNewTransactor) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MedianizerNew.contract.Transact(opts, "void")
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_MedianizerNew *MedianizerNewSession) Void() (*types.Transaction, error) {
	return _MedianizerNew.Contract.Void(&_MedianizerNew.TransactOpts)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_MedianizerNew *MedianizerNewTransactorSession) Void() (*types.Transaction, error) {
	return _MedianizerNew.Contract.Void(&_MedianizerNew.TransactOpts)
}

// MedianizerNewLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the MedianizerNew contract.
type MedianizerNewLogSetAuthorityIterator struct {
	Event *MedianizerNewLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *MedianizerNewLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerNewLogSetAuthority)
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
		it.Event = new(MedianizerNewLogSetAuthority)
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
func (it *MedianizerNewLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerNewLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerNewLogSetAuthority represents a LogSetAuthority event raised by the MedianizerNew contract.
type MedianizerNewLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MedianizerNew *MedianizerNewFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*MedianizerNewLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MedianizerNew.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &MedianizerNewLogSetAuthorityIterator{contract: _MedianizerNew.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MedianizerNew *MedianizerNewFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *MedianizerNewLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MedianizerNew.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerNewLogSetAuthority)
				if err := _MedianizerNew.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// MedianizerNewLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the MedianizerNew contract.
type MedianizerNewLogSetOwnerIterator struct {
	Event *MedianizerNewLogSetOwner // Event containing the contract specifics and raw log

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
func (it *MedianizerNewLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerNewLogSetOwner)
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
		it.Event = new(MedianizerNewLogSetOwner)
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
func (it *MedianizerNewLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerNewLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerNewLogSetOwner represents a LogSetOwner event raised by the MedianizerNew contract.
type MedianizerNewLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MedianizerNew *MedianizerNewFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*MedianizerNewLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MedianizerNew.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MedianizerNewLogSetOwnerIterator{contract: _MedianizerNew.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MedianizerNew *MedianizerNewFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *MedianizerNewLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MedianizerNew.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerNewLogSetOwner)
				if err := _MedianizerNew.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// MedianizerNewLogValueIterator is returned from FilterLogValue and is used to iterate over the raw logs and unpacked data for LogValue events raised by the MedianizerNew contract.
type MedianizerNewLogValueIterator struct {
	Event *MedianizerNewLogValue // Event containing the contract specifics and raw log

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
func (it *MedianizerNewLogValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerNewLogValue)
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
		it.Event = new(MedianizerNewLogValue)
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
func (it *MedianizerNewLogValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerNewLogValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerNewLogValue represents a LogValue event raised by the MedianizerNew contract.
type MedianizerNewLogValue struct {
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogValue is a free log retrieval operation binding the contract event 0x296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b527.
//
// Solidity: event LogValue(bytes32 val)
func (_MedianizerNew *MedianizerNewFilterer) FilterLogValue(opts *bind.FilterOpts) (*MedianizerNewLogValueIterator, error) {

	logs, sub, err := _MedianizerNew.contract.FilterLogs(opts, "LogValue")
	if err != nil {
		return nil, err
	}
	return &MedianizerNewLogValueIterator{contract: _MedianizerNew.contract, event: "LogValue", logs: logs, sub: sub}, nil
}

// WatchLogValue is a free log subscription operation binding the contract event 0x296ba4ca62c6c21c95e828080cb8aec7481b71390585605300a8a76f9e95b527.
//
// Solidity: event LogValue(bytes32 val)
func (_MedianizerNew *MedianizerNewFilterer) WatchLogValue(opts *bind.WatchOpts, sink chan<- *MedianizerNewLogValue) (event.Subscription, error) {

	logs, sub, err := _MedianizerNew.contract.WatchLogs(opts, "LogValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerNewLogValue)
				if err := _MedianizerNew.contract.UnpackLog(event, "LogValue", log); err != nil {
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
