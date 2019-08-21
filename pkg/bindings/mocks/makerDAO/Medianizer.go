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

// MedianizerABI is the input ABI used to generate the binding from.
const MedianizerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compute\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"unset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"indexes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes12\"}],\"name\":\"values\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"min_\",\"type\":\"uint96\"}],\"name\":\"setMin\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"void\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"bytes12\"},{\"name\":\"wat\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pos\",\"type\":\"bytes12\"}],\"name\":\"unset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next_\",\"type\":\"bytes12\"}],\"name\":\"setNext\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"min\",\"outputs\":[{\"name\":\"\",\"type\":\"uint96\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// MedianizerBin is the compiled bytecode used for deploying new contracts.
var MedianizerBin = "0x606060405260058054606060020a60c060020a03196001606060020a0319909116600117166c010000000000000000000000001790555b60018054600160a060020a03191633600160a060020a03169081179091556040517fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a25b5b6110748061008e6000396000f3006060604052361561010f5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af403581146101115780631504460f1461012f57806318178358146101445780631a43c338146101565780632801617e1461017f5780632966d1b91461019d5780632db78d93146101bb5780634c8fe526146101f457806357de26a41461022157806359e02dd714610243578063651dd0de1461026c5780636ba5ef0d146102a55780637a9e5e4b146102c35780638da5cb5b146102e1578063ac4c25b21461030d578063beb38b431461031f578063bf7e214f1461034a578063e0a1fdad14610376578063f2c5925d14610395578063f8897945146103b4575bfe5b341561011957fe5b61012d600160a060020a03600435166103e0565b005b341561013757fe5b61012d600435610450565b005b341561014c57fe5b61012d6104e9565b005b341561015e57fe5b6101666104f6565b6040805192835290151560208301528051918290030190f35b341561018757fe5b61012d600160a060020a03600435166108d0565b005b34156101a557fe5b61012d600160a060020a036004351661094e565b005b34156101c357fe5b6101d7600160a060020a036004351661097d565b60408051600160a060020a03199092168252519081900360200190f35b34156101fc57fe5b6101d7610995565b60408051600160a060020a03199092168252519081900360200190f35b341561022957fe5b6102316109a1565b60408051918252519081900360200190f35b341561024b57fe5b6101666109c5565b6040805192835290151560208301528051918290030190f35b341561027457fe5b610289600160a060020a0319600435166109da565b60408051600160a060020a039092168252519081900360200190f35b34156102ad57fe5b61012d6001606060020a03600435166109f5565b005b34156102cb57fe5b61012d600160a060020a0360043516610ac6565b005b34156102e957fe5b610289610b32565b60408051600160a060020a039092168252519081900360200190f35b341561031557fe5b61012d610b41565b005b341561032757fe5b61012d600160a060020a031960043516602435600160a060020a0316610bde565b005b341561035257fe5b610289610d6c565b60408051600160a060020a039092168252519081900360200190f35b341561037e57fe5b61012d600160a060020a031960043516610d7b565b005b341561039d57fe5b61012d600160a060020a031960043516610d8a565b005b34156103bc57fe5b6103c4610e3e565b604080516001606060020a039092168252519081900360200190f35b6103fe6103f933600035600160e060020a031916610e5d565b610f63565b60018054600160a060020a031916600160a060020a0383811691909117918290556040519116907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a25b5b50565b6040805134808252602082018381523693830184905260043593602435938493869333600160a060020a03169360008035600160e060020a031916949092606082018484808284376040519201829003965090945050505050a46104b26104f6565b6001805491151560a060020a0274ff0000000000000000000000000000000000000000199092169190911790556002555b5b505050565b6104f36000610450565b5b565b60006000610502611036565b6000600060006000600060006000600060006001600560009054906101000a900460a060020a0260a060020a9004036001606060020a03166040518059106105475750595b908082528060200260200182016040525b50995060009850600197505b60055460a060020a908102046001606060020a03908116908916101561079a57600160a060020a031960a060020a890216600090815260036020526040902054600160a060020a03161561078d5760a060020a8802600160a060020a031916600090815260036020526040808220548151820183905281517f59e02dd70000000000000000000000000000000000000000000000000000000081528251600160a060020a03909216936359e02dd79360048084019491939192918390030190829087803b151561063057fe5b6102c65a03f1151561063e57fe5b50506040518051602090910151909850965050851561078d576001606060020a038916158061069057508960018a036001606060020a031681518110151561068257fe5b602090810290910101518710155b156106bc57868a8a6001606060020a03168151811015156106ad57fe5b60209081029091010152610786565b600094505b89856001606060020a03168151811015156106d857fe5b6020908102909101015187106106f3576001909401936106c1565b8893505b846001606060020a0316846001606060020a031611156107635789600185036001606060020a031681518110151561072b57fe5b906020019060200201518a856001606060020a031681518110151561074c57fe5b602090810290910101525b600019909301926106f7565b868a866001606060020a031681518110151561077b57fe5b602090810290910101525b6001909801975b5b5b600190970196610564565b6005546001606060020a036c010000000000000000000000009091048116908a1610156107cf5760025460009b509b506108c2565b60026001606060020a038a165b066001606060020a0316600014156108855789600160026001606060020a038c165b04036001606060020a031681518110151561081557fe5b6020908102909101015191508960026001606060020a038b165b046001606060020a031681518110151561084557fe5b60209081029091010151905061086c61085e8383610f74565b671bc16d674ec80000610f9d565b6fffffffffffffffffffffffffffffffff1692506108ba565b8960026001606060020a036000198c01165b046001606060020a03168151811015156108ad57fe5b9060200190602002015192505b8260019b509b505b505050505050505050509091565b60006108f06103f933600035600160e060020a031916610e5d565b610f63565b5060055460a060020a90810281900460010102610918600160a060020a031982161515610f63565b60055461092b9060a060020a0283610bde565b600580546bffffffffffffffffffffffff191660a060020a83041790555b5b5050565b600160a060020a03811660009081526004602052604081205461044c9160a060020a90910290610bde565b5b50565b60046020526000908152604090205460a060020a0281565b60055460a060020a0281565b6000600060006109af6109c5565b915091506109bc81610f63565b8192505b505090565b60025460015460a060020a900460ff165b9091565b600360205260009081526040902054600160a060020a031681565b6040805134808252602082018381523693830184905260043593602435938493869333600160a060020a03169360008035600160e060020a031916949092606082018484808284376040519201829003965090945050505050a4610a6d6103f933600035600160e060020a031916610e5d565b610f63565b6001606060020a0383161515610a835760006000fd5b6005805477ffffffffffffffffffffffff00000000000000000000000019166c010000000000000000000000006001606060020a038616021790555b5b5b505050565b610ae46103f933600035600160e060020a031916610e5d565b610f63565b60008054600160a060020a031916600160a060020a03838116919091178083556040519116917f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada491a25b5b50565b600154600160a060020a031681565b6040805134808252602082018381523693830184905260043593602435938493869333600160a060020a03169360008035600160e060020a031916949092606082018484808284376040519201829003965090945050505050a4610bb96103f933600035600160e060020a031916610e5d565b610f63565b6001805474ff0000000000000000000000000000000000000000191690555b5b5b5050565b6040805134808252602082018381523693830184905260043593602435938493869333600160a060020a03169360008035600160e060020a031916949092606082018484808284376040519201829003965090945050505050a4610c566103f933600035600160e060020a031916610e5d565b610f63565b600160a060020a031984161515610c6d5760006000fd5b600160a060020a03831615801590610cac5750600160a060020a03831660009081526004602052604090205460a060020a02600160a060020a03191615155b15610cb75760006000fd5b600160a060020a03198416600090815260036020908152604080832054600160a060020a039081168452600490925290912080546bffffffffffffffffffffffff19169055831615610d3757600160a060020a038316600090815260046020526040902080546bffffffffffffffffffffffff191660a060020a86041790555b600160a060020a031984811660009081526003602052604090208054909116600160a060020a0385161790555b5b5b50505050565b600054600160a060020a031681565b61044c816000610bde565b5b50565b6040805134808252602082018381523693830184905260043593602435938493869333600160a060020a03169360008035600160e060020a031916949092606082018484808284376040519201829003965090945050505050a4610e026103f933600035600160e060020a031916610e5d565b610f63565b600160a060020a031983161515610e195760006000fd5b600580546bffffffffffffffffffffffff191660a060020a85041790555b5b5b505050565b6005546c0100000000000000000000000090046001606060020a031681565b600030600160a060020a031683600160a060020a03161415610e8157506001610f5a565b600154600160a060020a0384811691161415610e9f57506001610f5a565b600054600160a060020a03161515610eb957506000610f5a565b6000805460408051602090810184905281517fb7009613000000000000000000000000000000000000000000000000000000008152600160a060020a0388811660048301523081166024830152600160e060020a0319881660448301529251929093169363b70096139360648082019492918390030190829087803b1515610f3d57fe5b6102c65a03f11515610f4b57fe5b5050604051519150610f5a9050565b5b5b5b92915050565b80151561044c5760006000fd5b5b50565b8082016fffffffffffffffffffffffffffffffff8084169082161015610f5a57fe5b5b92915050565b600061100b6fffffffffffffffffffffffffffffffff83166002815b046fffffffffffffffffffffffffffffffff16670de0b6b3a76400006fffffffffffffffffffffffffffffffff16866fffffffffffffffffffffffffffffffff16020181151561100557fe5b04611014565b90505b92915050565b806fffffffffffffffffffffffffffffffff8116811461103057fe5b5b919050565b604080516020810190915260008152905600a165627a7a72305820c8fbad615728222f8039cb4f94c1174b59e0f6b9c9fcbd672fe38a541f76a7c20029"

// DeployMedianizer deploys a new Ethereum contract, binding an instance of Medianizer to it.
func DeployMedianizer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Medianizer, error) {
	parsed, err := abi.JSON(strings.NewReader(MedianizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MedianizerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Medianizer{MedianizerCaller: MedianizerCaller{contract: contract}, MedianizerTransactor: MedianizerTransactor{contract: contract}, MedianizerFilterer: MedianizerFilterer{contract: contract}}, nil
}

// Medianizer is an auto generated Go binding around an Ethereum contract.
type Medianizer struct {
	MedianizerCaller     // Read-only binding to the contract
	MedianizerTransactor // Write-only binding to the contract
	MedianizerFilterer   // Log filterer for contract events
}

// MedianizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MedianizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MedianizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MedianizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MedianizerSession struct {
	Contract     *Medianizer       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedianizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MedianizerCallerSession struct {
	Contract *MedianizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MedianizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MedianizerTransactorSession struct {
	Contract     *MedianizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MedianizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MedianizerRaw struct {
	Contract *Medianizer // Generic contract binding to access the raw methods on
}

// MedianizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MedianizerCallerRaw struct {
	Contract *MedianizerCaller // Generic read-only contract binding to access the raw methods on
}

// MedianizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MedianizerTransactorRaw struct {
	Contract *MedianizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMedianizer creates a new instance of Medianizer, bound to a specific deployed contract.
func NewMedianizer(address common.Address, backend bind.ContractBackend) (*Medianizer, error) {
	contract, err := bindMedianizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Medianizer{MedianizerCaller: MedianizerCaller{contract: contract}, MedianizerTransactor: MedianizerTransactor{contract: contract}, MedianizerFilterer: MedianizerFilterer{contract: contract}}, nil
}

// NewMedianizerCaller creates a new read-only instance of Medianizer, bound to a specific deployed contract.
func NewMedianizerCaller(address common.Address, caller bind.ContractCaller) (*MedianizerCaller, error) {
	contract, err := bindMedianizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MedianizerCaller{contract: contract}, nil
}

// NewMedianizerTransactor creates a new write-only instance of Medianizer, bound to a specific deployed contract.
func NewMedianizerTransactor(address common.Address, transactor bind.ContractTransactor) (*MedianizerTransactor, error) {
	contract, err := bindMedianizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MedianizerTransactor{contract: contract}, nil
}

// NewMedianizerFilterer creates a new log filterer instance of Medianizer, bound to a specific deployed contract.
func NewMedianizerFilterer(address common.Address, filterer bind.ContractFilterer) (*MedianizerFilterer, error) {
	contract, err := bindMedianizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MedianizerFilterer{contract: contract}, nil
}

// bindMedianizer binds a generic wrapper to an already deployed contract.
func bindMedianizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MedianizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Medianizer *MedianizerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Medianizer.Contract.MedianizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Medianizer *MedianizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.Contract.MedianizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Medianizer *MedianizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Medianizer.Contract.MedianizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Medianizer *MedianizerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Medianizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Medianizer *MedianizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Medianizer *MedianizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Medianizer.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Medianizer *MedianizerCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Medianizer *MedianizerSession) Authority() (common.Address, error) {
	return _Medianizer.Contract.Authority(&_Medianizer.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Medianizer *MedianizerCallerSession) Authority() (common.Address, error) {
	return _Medianizer.Contract.Authority(&_Medianizer.CallOpts)
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_Medianizer *MedianizerCaller) Compute(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Medianizer.contract.Call(opts, out, "compute")
	return *ret0, *ret1, err
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_Medianizer *MedianizerSession) Compute() ([32]byte, bool, error) {
	return _Medianizer.Contract.Compute(&_Medianizer.CallOpts)
}

// Compute is a free data retrieval call binding the contract method 0x1a43c338.
//
// Solidity: function compute() constant returns(bytes32, bool)
func (_Medianizer *MedianizerCallerSession) Compute() ([32]byte, bool, error) {
	return _Medianizer.Contract.Compute(&_Medianizer.CallOpts)
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes(address ) constant returns(bytes12)
func (_Medianizer *MedianizerCaller) Indexes(opts *bind.CallOpts, arg0 common.Address) ([12]byte, error) {
	var (
		ret0 = new([12]byte)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "indexes", arg0)
	return *ret0, err
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes(address ) constant returns(bytes12)
func (_Medianizer *MedianizerSession) Indexes(arg0 common.Address) ([12]byte, error) {
	return _Medianizer.Contract.Indexes(&_Medianizer.CallOpts, arg0)
}

// Indexes is a free data retrieval call binding the contract method 0x2db78d93.
//
// Solidity: function indexes(address ) constant returns(bytes12)
func (_Medianizer *MedianizerCallerSession) Indexes(arg0 common.Address) ([12]byte, error) {
	return _Medianizer.Contract.Indexes(&_Medianizer.CallOpts, arg0)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_Medianizer *MedianizerCaller) Min(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "min")
	return *ret0, err
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_Medianizer *MedianizerSession) Min() (*big.Int, error) {
	return _Medianizer.Contract.Min(&_Medianizer.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() constant returns(uint96)
func (_Medianizer *MedianizerCallerSession) Min() (*big.Int, error) {
	return _Medianizer.Contract.Min(&_Medianizer.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_Medianizer *MedianizerCaller) Next(opts *bind.CallOpts) ([12]byte, error) {
	var (
		ret0 = new([12]byte)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "next")
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_Medianizer *MedianizerSession) Next() ([12]byte, error) {
	return _Medianizer.Contract.Next(&_Medianizer.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4c8fe526.
//
// Solidity: function next() constant returns(bytes12)
func (_Medianizer *MedianizerCallerSession) Next() ([12]byte, error) {
	return _Medianizer.Contract.Next(&_Medianizer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Medianizer *MedianizerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Medianizer *MedianizerSession) Owner() (common.Address, error) {
	return _Medianizer.Contract.Owner(&_Medianizer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Medianizer *MedianizerCallerSession) Owner() (common.Address, error) {
	return _Medianizer.Contract.Owner(&_Medianizer.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_Medianizer *MedianizerCaller) Peek(opts *bind.CallOpts) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Medianizer.contract.Call(opts, out, "peek")
	return *ret0, *ret1, err
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_Medianizer *MedianizerSession) Peek() ([32]byte, bool, error) {
	return _Medianizer.Contract.Peek(&_Medianizer.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32, bool)
func (_Medianizer *MedianizerCallerSession) Peek() ([32]byte, bool, error) {
	return _Medianizer.Contract.Peek(&_Medianizer.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_Medianizer *MedianizerCaller) Read(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "read")
	return *ret0, err
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_Medianizer *MedianizerSession) Read() ([32]byte, error) {
	return _Medianizer.Contract.Read(&_Medianizer.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32)
func (_Medianizer *MedianizerCallerSession) Read() ([32]byte, error) {
	return _Medianizer.Contract.Read(&_Medianizer.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values(bytes12 ) constant returns(address)
func (_Medianizer *MedianizerCaller) Values(opts *bind.CallOpts, arg0 [12]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Medianizer.contract.Call(opts, out, "values", arg0)
	return *ret0, err
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values(bytes12 ) constant returns(address)
func (_Medianizer *MedianizerSession) Values(arg0 [12]byte) (common.Address, error) {
	return _Medianizer.Contract.Values(&_Medianizer.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x651dd0de.
//
// Solidity: function values(bytes12 ) constant returns(address)
func (_Medianizer *MedianizerCallerSession) Values(arg0 [12]byte) (common.Address, error) {
	return _Medianizer.Contract.Values(&_Medianizer.CallOpts, arg0)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ) returns()
func (_Medianizer *MedianizerTransactor) Poke(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "poke", arg0)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ) returns()
func (_Medianizer *MedianizerSession) Poke(arg0 [32]byte) (*types.Transaction, error) {
	return _Medianizer.Contract.Poke(&_Medianizer.TransactOpts, arg0)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ) returns()
func (_Medianizer *MedianizerTransactorSession) Poke(arg0 [32]byte) (*types.Transaction, error) {
	return _Medianizer.Contract.Poke(&_Medianizer.TransactOpts, arg0)
}

// Poke0 is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Medianizer *MedianizerTransactor) Poke0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "poke0")
}

// Poke0 is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Medianizer *MedianizerSession) Poke0() (*types.Transaction, error) {
	return _Medianizer.Contract.Poke0(&_Medianizer.TransactOpts)
}

// Poke0 is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Medianizer *MedianizerTransactorSession) Poke0() (*types.Transaction, error) {
	return _Medianizer.Contract.Poke0(&_Medianizer.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address wat) returns()
func (_Medianizer *MedianizerTransactor) Set(opts *bind.TransactOpts, wat common.Address) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "set", wat)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address wat) returns()
func (_Medianizer *MedianizerSession) Set(wat common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.Set(&_Medianizer.TransactOpts, wat)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address wat) returns()
func (_Medianizer *MedianizerTransactorSession) Set(wat common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.Set(&_Medianizer.TransactOpts, wat)
}

// Set0 is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(bytes12 pos, address wat) returns()
func (_Medianizer *MedianizerTransactor) Set0(opts *bind.TransactOpts, pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "set0", pos, wat)
}

// Set0 is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(bytes12 pos, address wat) returns()
func (_Medianizer *MedianizerSession) Set0(pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.Set0(&_Medianizer.TransactOpts, pos, wat)
}

// Set0 is a paid mutator transaction binding the contract method 0xbeb38b43.
//
// Solidity: function set(bytes12 pos, address wat) returns()
func (_Medianizer *MedianizerTransactorSession) Set0(pos [12]byte, wat common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.Set0(&_Medianizer.TransactOpts, pos, wat)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Medianizer *MedianizerTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Medianizer *MedianizerSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.SetAuthority(&_Medianizer.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Medianizer *MedianizerTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.SetAuthority(&_Medianizer.TransactOpts, authority_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(uint96 min_) returns()
func (_Medianizer *MedianizerTransactor) SetMin(opts *bind.TransactOpts, min_ *big.Int) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "setMin", min_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(uint96 min_) returns()
func (_Medianizer *MedianizerSession) SetMin(min_ *big.Int) (*types.Transaction, error) {
	return _Medianizer.Contract.SetMin(&_Medianizer.TransactOpts, min_)
}

// SetMin is a paid mutator transaction binding the contract method 0x6ba5ef0d.
//
// Solidity: function setMin(uint96 min_) returns()
func (_Medianizer *MedianizerTransactorSession) SetMin(min_ *big.Int) (*types.Transaction, error) {
	return _Medianizer.Contract.SetMin(&_Medianizer.TransactOpts, min_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(bytes12 next_) returns()
func (_Medianizer *MedianizerTransactor) SetNext(opts *bind.TransactOpts, next_ [12]byte) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "setNext", next_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(bytes12 next_) returns()
func (_Medianizer *MedianizerSession) SetNext(next_ [12]byte) (*types.Transaction, error) {
	return _Medianizer.Contract.SetNext(&_Medianizer.TransactOpts, next_)
}

// SetNext is a paid mutator transaction binding the contract method 0xf2c5925d.
//
// Solidity: function setNext(bytes12 next_) returns()
func (_Medianizer *MedianizerTransactorSession) SetNext(next_ [12]byte) (*types.Transaction, error) {
	return _Medianizer.Contract.SetNext(&_Medianizer.TransactOpts, next_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Medianizer *MedianizerTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Medianizer *MedianizerSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.SetOwner(&_Medianizer.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Medianizer *MedianizerTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.SetOwner(&_Medianizer.TransactOpts, owner_)
}

// Unset is a paid mutator transaction binding the contract method 0x2966d1b9.
//
// Solidity: function unset(address wat) returns()
func (_Medianizer *MedianizerTransactor) Unset(opts *bind.TransactOpts, wat common.Address) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "unset", wat)
}

// Unset is a paid mutator transaction binding the contract method 0x2966d1b9.
//
// Solidity: function unset(address wat) returns()
func (_Medianizer *MedianizerSession) Unset(wat common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.Unset(&_Medianizer.TransactOpts, wat)
}

// Unset is a paid mutator transaction binding the contract method 0x2966d1b9.
//
// Solidity: function unset(address wat) returns()
func (_Medianizer *MedianizerTransactorSession) Unset(wat common.Address) (*types.Transaction, error) {
	return _Medianizer.Contract.Unset(&_Medianizer.TransactOpts, wat)
}

// Unset0 is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(bytes12 pos) returns()
func (_Medianizer *MedianizerTransactor) Unset0(opts *bind.TransactOpts, pos [12]byte) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "unset0", pos)
}

// Unset0 is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(bytes12 pos) returns()
func (_Medianizer *MedianizerSession) Unset0(pos [12]byte) (*types.Transaction, error) {
	return _Medianizer.Contract.Unset0(&_Medianizer.TransactOpts, pos)
}

// Unset0 is a paid mutator transaction binding the contract method 0xe0a1fdad.
//
// Solidity: function unset(bytes12 pos) returns()
func (_Medianizer *MedianizerTransactorSession) Unset0(pos [12]byte) (*types.Transaction, error) {
	return _Medianizer.Contract.Unset0(&_Medianizer.TransactOpts, pos)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Medianizer *MedianizerTransactor) Void(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medianizer.contract.Transact(opts, "void")
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Medianizer *MedianizerSession) Void() (*types.Transaction, error) {
	return _Medianizer.Contract.Void(&_Medianizer.TransactOpts)
}

// Void is a paid mutator transaction binding the contract method 0xac4c25b2.
//
// Solidity: function void() returns()
func (_Medianizer *MedianizerTransactorSession) Void() (*types.Transaction, error) {
	return _Medianizer.Contract.Void(&_Medianizer.TransactOpts)
}

// MedianizerLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Medianizer contract.
type MedianizerLogSetAuthorityIterator struct {
	Event *MedianizerLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *MedianizerLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerLogSetAuthority)
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
		it.Event = new(MedianizerLogSetAuthority)
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
func (it *MedianizerLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerLogSetAuthority represents a LogSetAuthority event raised by the Medianizer contract.
type MedianizerLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Medianizer *MedianizerFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*MedianizerLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Medianizer.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &MedianizerLogSetAuthorityIterator{contract: _Medianizer.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Medianizer *MedianizerFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *MedianizerLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Medianizer.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerLogSetAuthority)
				if err := _Medianizer.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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
func (_Medianizer *MedianizerFilterer) ParseLogSetAuthority(log types.Log) (*MedianizerLogSetAuthority, error) {
	event := new(MedianizerLogSetAuthority)
	if err := _Medianizer.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MedianizerLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Medianizer contract.
type MedianizerLogSetOwnerIterator struct {
	Event *MedianizerLogSetOwner // Event containing the contract specifics and raw log

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
func (it *MedianizerLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MedianizerLogSetOwner)
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
		it.Event = new(MedianizerLogSetOwner)
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
func (it *MedianizerLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MedianizerLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MedianizerLogSetOwner represents a LogSetOwner event raised by the Medianizer contract.
type MedianizerLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Medianizer *MedianizerFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*MedianizerLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Medianizer.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MedianizerLogSetOwnerIterator{contract: _Medianizer.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Medianizer *MedianizerFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *MedianizerLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Medianizer.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MedianizerLogSetOwner)
				if err := _Medianizer.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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
func (_Medianizer *MedianizerFilterer) ParseLogSetOwner(log types.Log) (*MedianizerLogSetOwner, error) {
	event := new(MedianizerLogSetOwner)
	if err := _Medianizer.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}
