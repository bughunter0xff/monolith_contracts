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

// DSProxyFactoryABI is the input ABI used to generate the binding from.
const DSProxyFactoryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cache\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cache\",\"type\":\"address\"}],\"name\":\"Created\",\"type\":\"event\"}]"

// DSProxyFactoryBin is the compiled bytecode used for deploying new contracts.
const DSProxyFactoryBin = `608060405261000c61005b565b604051809103906000f080158015610028573d6000803e3d6000fd5b5060018054600160a060020a031916600160a060020a039290921691909117905534801561005557600080fd5b5061006b565b6040516102bc80610e3283390190565b610db88061007a6000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166329710388811461006657806360c7d2951461009b5780638e1a55fc146100cc578063f3701da2146100e1575b600080fd5b34801561007257600080fd5b50610087600160a060020a0360043516610102565b604080519115158252519081900360200190f35b3480156100a757600080fd5b506100b0610117565b60408051600160a060020a039092168252519081900360200190f35b3480156100d857600080fd5b506100b0610126565b3480156100ed57600080fd5b506100b0600160a060020a0360043516610136565b60006020819052908152604090205460ff1681565b600154600160a060020a031681565b600061013133610136565b905090565b600154600090600160a060020a031661014d61027f565b600160a060020a03909116815260405190819003602001906000f08015801561017a573d6000803e3d6000fd5b5060015460408051600160a060020a038085168252928316602082015281519394509185169233927f259b30ca39885c6d801a0b5dbc988640f3c25e2f37531fe138c5c5af8955d41b92908290030190a380600160a060020a03166313af4035836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082600160a060020a0316600160a060020a03168152602001915050600060405180830381600087803b15801561023f57600080fd5b505af1158015610253573d6000803e3d6000fd5b505050600160a060020a0382166000908152602081905260409020805460ff1916600117905550919050565b604051610afd80610290833901905600608060405234801561001057600080fd5b50604051602080610afd833981016040819052905160018054600160a060020a0319163390811790915590917fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a261007481640100000000610085810204565b151561007f57600080fd5b5061028f565b60006100bd337fffffffff0000000000000000000000000000000000000000000000000000000083351664010000000061016e810204565b15156100c857600080fd5b604080513480825260208201838152369383018490526004359360243593849386933393600080357fffffffff0000000000000000000000000000000000000000000000000000000016949092606082018484808284376040519201829003965090945050505050a4600160a060020a038416151561014657600080fd5b60028054600160a060020a038616600160a060020a0319909116179055600192505050919050565b6000600160a060020a03831630141561018957506001610289565b600154600160a060020a03848116911614156101a757506001610289565b600054600160a060020a031615156101c157506000610289565b60008054604080517fb7009613000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301523060248301527fffffffff00000000000000000000000000000000000000000000000000000000871660448301529151919092169263b700961392606480820193602093909283900390910190829087803b15801561025a57600080fd5b505af115801561026e573d6000803e3d6000fd5b505050506040513d602081101561028457600080fd5b505190505b92915050565b61085f8061029e6000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af4035811461008f5780631cff79cd146100b05780631f6a1eb91461011c57806360c7d295146101c95780637a9e5e4b146101fa5780638da5cb5b1461021b578063948f507614610230578063bf7e214f14610265575b005b34801561009b57600080fd5b5061008d600160a060020a036004351661027a565b60408051602060046024803582810135601f810185900485028601850190965285855261010a958335600160a060020a03169536956044949193909101919081908401838280828437509497506102f89650505050505050565b60408051918252519081900360200190f35b6040805160206004803580820135601f81018490048402850184019095528484526101a694369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506103b59650505050505050565b60408051600160a060020a03909316835260208301919091528051918290030190f35b3480156101d557600080fd5b506101de6105c5565b60408051600160a060020a039092168252519081900360200190f35b34801561020657600080fd5b5061008d600160a060020a03600435166105d4565b34801561022757600080fd5b506101de61064e565b34801561023c57600080fd5b50610251600160a060020a036004351661065d565b604080519115158252519081900360200190f35b34801561027157600080fd5b506101de61071b565b61029033600035600160e060020a03191661072a565b151561029b57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383811691909117918290556040519116907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a250565b600061031033600035600160e060020a03191661072a565b151561031b57600080fd5b60408051348082526020820183815236938301849052600435936024359384938693339360008035600160e060020a031916949092606082018484808284376040519201829003965090945050505050a4600160a060020a038516151561038157600080fd5b60206000855160208701886113885a03f460005193508015600181146103a6576103ab565b600080fd5b5050505092915050565b6002546040517f8bf4515c0000000000000000000000000000000000000000000000000000000081526020600482018181528551602484015285516000948594600160a060020a0390911693638bf4515c93899390928392604490910191908501908083838b5b8381101561043457818101518382015260200161041c565b50505050905090810190601f1680156104615780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561048057600080fd5b505af1158015610494573d6000803e3d6000fd5b505050506040513d60208110156104aa57600080fd5b50519150600160a060020a03821615156105b2576002546040517f7ed0c3b2000000000000000000000000000000000000000000000000000000008152602060048201818152875160248401528751600160a060020a0390941693637ed0c3b293899383926044909201919085019080838360005b8381101561053757818101518382015260200161051f565b50505050905090810190601f1680156105645780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561058357600080fd5b505af1158015610597573d6000803e3d6000fd5b505050506040513d60208110156105ad57600080fd5b505191505b6105bc82846102f8565b90509250929050565b600254600160a060020a031681565b6105ea33600035600160e060020a03191661072a565b15156105f557600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03838116919091178083556040519116917f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada491a250565b600154600160a060020a031681565b600061067533600035600160e060020a03191661072a565b151561068057600080fd5b60408051348082526020820183815236938301849052600435936024359384938693339360008035600160e060020a031916949092606082018484808284376040519201829003965090945050505050a4600160a060020a03841615156106e657600080fd5b60028054600160a060020a03861673ffffffffffffffffffffffffffffffffffffffff19909116179055600192505050919050565b600054600160a060020a031681565b6000600160a060020a0383163014156107455750600161082d565b600154600160a060020a03848116911614156107635750600161082d565b600054600160a060020a0316151561077d5750600061082d565b60008054604080517fb7009613000000000000000000000000000000000000000000000000000000008152600160a060020a038781166004830152306024830152600160e060020a0319871660448301529151919092169263b700961392606480820193602093909283900390910190829087803b1580156107fe57600080fd5b505af1158015610812573d6000803e3d6000fd5b505050506040513d602081101561082857600080fd5b505190505b929150505600a165627a7a72305820e4b2bff005a9953f56d84a0d3baae02d507047d280392ce3737b15466f732e8e0029a165627a7a72305820b0c918df7492a13edf063cfc1e044dc851a9ad13934c1f9143a72f57e93cf5f80029608060405234801561001057600080fd5b5061029c806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637ed0c3b281146100505780638bf4515c146100d2575b600080fd5b34801561005c57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100a994369492936024939284019190819084018382808284375094975061012b9650505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100de57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100a99436949293602493928401919081908401838280828437509497506101e79650505050505050565b6000808251602084016000f09150813b156001811461004b5750826040518082805190602001908083835b602083106101755780518252601f199092019160209182019101610156565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060009081529081905291909120805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff87161790555092949350505050565b600080826040518082805190602001908083835b6020831061021a5780518252601f1990920191602091820191016101fb565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120600090815290819052919091205473ffffffffffffffffffffffffffffffffffffffff1696955050505050505600a165627a7a7230582005add4c7dd4fc705c0ac3d441c102e581a0e9d1fb0bee9f1c77735ca20eb80c20029`

// DeployDSProxyFactory deploys a new Ethereum contract, binding an instance of DSProxyFactory to it.
func DeployDSProxyFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSProxyFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(DSProxyFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSProxyFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSProxyFactory{DSProxyFactoryCaller: DSProxyFactoryCaller{contract: contract}, DSProxyFactoryTransactor: DSProxyFactoryTransactor{contract: contract}, DSProxyFactoryFilterer: DSProxyFactoryFilterer{contract: contract}}, nil
}

// DSProxyFactory is an auto generated Go binding around an Ethereum contract.
type DSProxyFactory struct {
	DSProxyFactoryCaller     // Read-only binding to the contract
	DSProxyFactoryTransactor // Write-only binding to the contract
	DSProxyFactoryFilterer   // Log filterer for contract events
}

// DSProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSProxyFactorySession struct {
	Contract     *DSProxyFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSProxyFactoryCallerSession struct {
	Contract *DSProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DSProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSProxyFactoryTransactorSession struct {
	Contract     *DSProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DSProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSProxyFactoryRaw struct {
	Contract *DSProxyFactory // Generic contract binding to access the raw methods on
}

// DSProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSProxyFactoryCallerRaw struct {
	Contract *DSProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DSProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSProxyFactoryTransactorRaw struct {
	Contract *DSProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSProxyFactory creates a new instance of DSProxyFactory, bound to a specific deployed contract.
func NewDSProxyFactory(address common.Address, backend bind.ContractBackend) (*DSProxyFactory, error) {
	contract, err := bindDSProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSProxyFactory{DSProxyFactoryCaller: DSProxyFactoryCaller{contract: contract}, DSProxyFactoryTransactor: DSProxyFactoryTransactor{contract: contract}, DSProxyFactoryFilterer: DSProxyFactoryFilterer{contract: contract}}, nil
}

// NewDSProxyFactoryCaller creates a new read-only instance of DSProxyFactory, bound to a specific deployed contract.
func NewDSProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*DSProxyFactoryCaller, error) {
	contract, err := bindDSProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSProxyFactoryCaller{contract: contract}, nil
}

// NewDSProxyFactoryTransactor creates a new write-only instance of DSProxyFactory, bound to a specific deployed contract.
func NewDSProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DSProxyFactoryTransactor, error) {
	contract, err := bindDSProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSProxyFactoryTransactor{contract: contract}, nil
}

// NewDSProxyFactoryFilterer creates a new log filterer instance of DSProxyFactory, bound to a specific deployed contract.
func NewDSProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DSProxyFactoryFilterer, error) {
	contract, err := bindDSProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSProxyFactoryFilterer{contract: contract}, nil
}

// bindDSProxyFactory binds a generic wrapper to an already deployed contract.
func bindDSProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSProxyFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSProxyFactory *DSProxyFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSProxyFactory.Contract.DSProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSProxyFactory *DSProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSProxyFactory.Contract.DSProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSProxyFactory *DSProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSProxyFactory.Contract.DSProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSProxyFactory *DSProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSProxyFactory *DSProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSProxyFactory *DSProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_DSProxyFactory *DSProxyFactoryCaller) Cache(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSProxyFactory.contract.Call(opts, out, "cache")
	return *ret0, err
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_DSProxyFactory *DSProxyFactorySession) Cache() (common.Address, error) {
	return _DSProxyFactory.Contract.Cache(&_DSProxyFactory.CallOpts)
}

// Cache is a free data retrieval call binding the contract method 0x60c7d295.
//
// Solidity: function cache() constant returns(address)
func (_DSProxyFactory *DSProxyFactoryCallerSession) Cache() (common.Address, error) {
	return _DSProxyFactory.Contract.Cache(&_DSProxyFactory.CallOpts)
}

// IsProxy is a free data retrieval call binding the contract method 0x29710388.
//
// Solidity: function isProxy(address ) constant returns(bool)
func (_DSProxyFactory *DSProxyFactoryCaller) IsProxy(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DSProxyFactory.contract.Call(opts, out, "isProxy", arg0)
	return *ret0, err
}

// IsProxy is a free data retrieval call binding the contract method 0x29710388.
//
// Solidity: function isProxy(address ) constant returns(bool)
func (_DSProxyFactory *DSProxyFactorySession) IsProxy(arg0 common.Address) (bool, error) {
	return _DSProxyFactory.Contract.IsProxy(&_DSProxyFactory.CallOpts, arg0)
}

// IsProxy is a free data retrieval call binding the contract method 0x29710388.
//
// Solidity: function isProxy(address ) constant returns(bool)
func (_DSProxyFactory *DSProxyFactoryCallerSession) IsProxy(arg0 common.Address) (bool, error) {
	return _DSProxyFactory.Contract.IsProxy(&_DSProxyFactory.CallOpts, arg0)
}

// Build is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_DSProxyFactory *DSProxyFactoryTransactor) Build(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _DSProxyFactory.contract.Transact(opts, "build", owner)
}

// Build is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_DSProxyFactory *DSProxyFactorySession) Build(owner common.Address) (*types.Transaction, error) {
	return _DSProxyFactory.Contract.Build(&_DSProxyFactory.TransactOpts, owner)
}

// Build is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_DSProxyFactory *DSProxyFactoryTransactorSession) Build(owner common.Address) (*types.Transaction, error) {
	return _DSProxyFactory.Contract.Build(&_DSProxyFactory.TransactOpts, owner)
}

// DSProxyFactoryCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the DSProxyFactory contract.
type DSProxyFactoryCreatedIterator struct {
	Event *DSProxyFactoryCreated // Event containing the contract specifics and raw log

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
func (it *DSProxyFactoryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSProxyFactoryCreated)
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
		it.Event = new(DSProxyFactoryCreated)
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
func (it *DSProxyFactoryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSProxyFactoryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSProxyFactoryCreated represents a Created event raised by the DSProxyFactory contract.
type DSProxyFactoryCreated struct {
	Sender common.Address
	Owner  common.Address
	Proxy  common.Address
	Cache  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x259b30ca39885c6d801a0b5dbc988640f3c25e2f37531fe138c5c5af8955d41b.
//
// Solidity: event Created(address indexed sender, address indexed owner, address proxy, address cache)
func (_DSProxyFactory *DSProxyFactoryFilterer) FilterCreated(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*DSProxyFactoryCreatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSProxyFactory.contract.FilterLogs(opts, "Created", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &DSProxyFactoryCreatedIterator{contract: _DSProxyFactory.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x259b30ca39885c6d801a0b5dbc988640f3c25e2f37531fe138c5c5af8955d41b.
//
// Solidity: event Created(address indexed sender, address indexed owner, address proxy, address cache)
func (_DSProxyFactory *DSProxyFactoryFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *DSProxyFactoryCreated, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSProxyFactory.contract.WatchLogs(opts, "Created", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSProxyFactoryCreated)
				if err := _DSProxyFactory.contract.UnpackLog(event, "Created", log); err != nil {
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
