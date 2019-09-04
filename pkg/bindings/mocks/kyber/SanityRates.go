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

// SanityRatesABI is the input ABI used to generate the binding from.
const SanityRatesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reasonableDiffInBps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"srcs\",\"type\":\"address[]\"},{\"name\":\"diff\",\"type\":\"uint256[]\"}],\"name\":\"setReasonableDiff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"getSanityRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"srcs\",\"type\":\"address[]\"},{\"name\":\"rates\",\"type\":\"uint256[]\"}],\"name\":\"setSanityRates\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"}]"

// SanityRatesBin is the compiled bytecode used for deploying new contracts.
var SanityRatesBin = "0x6060604052341561000f57600080fd5b60405160208061115a8339810160405280805160008054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061005557600080fd5b60008054600160a060020a03909216600160a060020a03199092169190911790556110d5806100856000396000f3006060604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301a12fd38114610100578063267822471461012157806327a099d8146101505780633ccdbb28146101b6578063408ee7fe146101df5780635463a2e4146101fe5780635c53ec591461022f57806375829def146102be57806377f50f97146102dd5780637acc8678146102f05780637c423f541461030f5780639870d7fe14610322578063a58092b714610341578063ac8a584a14610366578063c57fbf9014610385578063ce56c454146103a4578063f5db370f146103c6578063f851a44014610455575b600080fd5b341561010b57600080fd5b61011f600160a060020a0360043516610468565b005b341561012c57600080fd5b6101346105d8565b604051600160a060020a03909116815260200160405180910390f35b341561015b57600080fd5b6101636105e7565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101a257808201518382015260200161018a565b505050509050019250505060405180910390f35b34156101c157600080fd5b61011f600160a060020a036004358116906024359060443516610650565b34156101ea57600080fd5b61011f600160a060020a0360043516610760565b341561020957600080fd5b61021d600160a060020a036004351661085c565b60405190815260200160405180910390f35b341561023a57600080fd5b61011f60046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284375094965061086e95505050505050565b34156102c957600080fd5b61011f600160a060020a0360043516610920565b34156102e857600080fd5b61011f6109bb565b34156102fb57600080fd5b61011f600160a060020a0360043516610a55565b341561031a57600080fd5b610163610b37565b341561032d57600080fd5b61011f600160a060020a0360043516610b9d565b341561034c57600080fd5b61021d600160a060020a0360043581169060243516610c6d565b341561037157600080fd5b61011f600160a060020a0360043516610d75565b341561039057600080fd5b61021d600160a060020a0360043516610ee1565b34156103af57600080fd5b61011f600435600160a060020a0360243516610ef3565b34156103d157600080fd5b61011f600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750949650610f8695505050505050565b341561046057600080fd5b610134611046565b6000805433600160a060020a0390811691161461048457600080fd5b600160a060020a03821660009081526003602052604090205460ff1615156104ab57600080fd5b50600160a060020a0381166000908152600360205260408120805460ff191690555b6005548110156105d45781600160a060020a03166005828154811015156104f057fe5b600091825260209091200154600160a060020a031614156105cc5760058054600019810190811061051d57fe5b60009182526020909120015460058054600160a060020a03909216918390811061054357fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600580549061057f906000198301611055565b507f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762826000604051600160a060020a039092168252151560208201526040908101905180910390a16105d4565b6001016104cd565b5050565b600154600160a060020a031681565b6105ef611079565b600480548060200260200160405190810160405280929190818152602001828054801561064557602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610627575b505050505090505b90565b60005433600160a060020a0390811691161461066b57600080fd5b82600160a060020a031663a9059cbb82846000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156106e157600080fd5b6102c65a03f115156106f257600080fd5b50505060405180519050151561070757600080fd5b7f72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6838383604051600160a060020a03938416815260208101929092529091166040808301919091526060909101905180910390a1505050565b60005433600160a060020a0390811691161461077b57600080fd5b600160a060020a03811660009081526003602052604090205460ff16156107a157600080fd5b600554603290106107b157600080fd5b7f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600360205260409020805460ff1916600190811790915560058054909181016108308382611055565b5060009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055565b60086020526000908152604090205481565b6000805433600160a060020a0390811691161461088a57600080fd5b815183511461089857600080fd5b5060005b825181101561091b576127108282815181106108b457fe5b9060200190602002015111156108c957600080fd5b8181815181106108d557fe5b90602001906020020151600860008584815181106108ef57fe5b90602001906020020151600160a060020a0316815260208101919091526040016000205560010161089c565b505050565b60005433600160a060020a0390811691161461093b57600080fd5b600160a060020a038116151561095057600080fd5b6001547f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4090600160a060020a0316604051600160a060020a03909116815260200160405180910390a160018054600160a060020a031916600160a060020a0392909216919091179055565b60015433600160a060020a039081169116146109d657600080fd5b6001546000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed91600160a060020a039081169116604051600160a060020a039283168152911660208201526040908101905180910390a16001805460008054600160a060020a0319908116600160a060020a03841617909155169055565b60005433600160a060020a03908116911614610a7057600080fd5b600160a060020a0381161515610a8557600080fd5b7f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4081604051600160a060020a03909116815260200160405180910390a16000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed908290600160a060020a0316604051600160a060020a039283168152911660208201526040908101905180910390a160008054600160a060020a031916600160a060020a0392909216919091179055565b610b3f611079565b600580548060200260200160405190810160405280929190818152602001828054801561064557602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610627575050505050905090565b60005433600160a060020a03908116911614610bb857600080fd5b600160a060020a03811660009081526002602052604090205460ff1615610bde57600080fd5b60045460329010610bee57600080fd5b7f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600260205260409020805460ff1916600190811790915560048054909181016108308382611055565b60008080600160a060020a03851673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14801590610cbb5750600160a060020a03841673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14155b15610cc95760009250610d6d565b600160a060020a03851673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415610d2b57600160a060020a0384166000908152600760205260409020546ec097ce7bc90715b34b9f1000000000811515610d2057fe5b049150839050610d48565b5050600160a060020a038316600090815260076020526040902054835b600160a060020a03811660009081526008602052604090205461271090810183020492505b505092915050565b6000805433600160a060020a03908116911614610d9157600080fd5b600160a060020a03821660009081526002602052604090205460ff161515610db857600080fd5b50600160a060020a0381166000908152600260205260408120805460ff191690555b6004548110156105d45781600160a060020a0316600482815481101515610dfd57fe5b600091825260209091200154600160a060020a03161415610ed957600480546000198101908110610e2a57fe5b60009182526020909120015460048054600160a060020a039092169183908110610e5057fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600480546000190190610e8c9082611055565b507f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b826000604051600160a060020a039092168252151560208201526040908101905180910390a16105d4565b600101610dda565b60076020526000908152604090205481565b60005433600160a060020a03908116911614610f0e57600080fd5b600160a060020a03811682156108fc0283604051600060405180830381858888f193505050501515610f3f57600080fd5b7fec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de8282604051918252600160a060020a031660208201526040908101905180910390a15050565b600160a060020a03331660009081526002602052604081205460ff161515610fad57600080fd5b8151835114610fbb57600080fd5b5060005b825181101561091b5769d3c21bcecceda1000000828281518110610fdf57fe5b906020019060200201511115610ff457600080fd5b81818151811061100057fe5b906020019060200201516007600085848151811061101a57fe5b90602001906020020151600160a060020a03168152602081019190915260400160002055600101610fbf565b600054600160a060020a031681565b81548183558181151161091b5760008381526020902061091b91810190830161108b565b60206040519081016040526000815290565b61064d91905b808211156110a55760008155600101611091565b50905600a165627a7a7230582057d4d9204658858f6f53390f6940db7d73aaa1975466eba07a8d4979974635cc0029"

// DeploySanityRates deploys a new Ethereum contract, binding an instance of SanityRates to it.
func DeploySanityRates(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *SanityRates, error) {
	parsed, err := abi.JSON(strings.NewReader(SanityRatesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SanityRatesBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SanityRates{SanityRatesCaller: SanityRatesCaller{contract: contract}, SanityRatesTransactor: SanityRatesTransactor{contract: contract}, SanityRatesFilterer: SanityRatesFilterer{contract: contract}}, nil
}

// SanityRates is an auto generated Go binding around an Ethereum contract.
type SanityRates struct {
	SanityRatesCaller     // Read-only binding to the contract
	SanityRatesTransactor // Write-only binding to the contract
	SanityRatesFilterer   // Log filterer for contract events
}

// SanityRatesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SanityRatesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SanityRatesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SanityRatesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SanityRatesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SanityRatesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SanityRatesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SanityRatesSession struct {
	Contract     *SanityRates      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SanityRatesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SanityRatesCallerSession struct {
	Contract *SanityRatesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SanityRatesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SanityRatesTransactorSession struct {
	Contract     *SanityRatesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SanityRatesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SanityRatesRaw struct {
	Contract *SanityRates // Generic contract binding to access the raw methods on
}

// SanityRatesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SanityRatesCallerRaw struct {
	Contract *SanityRatesCaller // Generic read-only contract binding to access the raw methods on
}

// SanityRatesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SanityRatesTransactorRaw struct {
	Contract *SanityRatesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSanityRates creates a new instance of SanityRates, bound to a specific deployed contract.
func NewSanityRates(address common.Address, backend bind.ContractBackend) (*SanityRates, error) {
	contract, err := bindSanityRates(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SanityRates{SanityRatesCaller: SanityRatesCaller{contract: contract}, SanityRatesTransactor: SanityRatesTransactor{contract: contract}, SanityRatesFilterer: SanityRatesFilterer{contract: contract}}, nil
}

// NewSanityRatesCaller creates a new read-only instance of SanityRates, bound to a specific deployed contract.
func NewSanityRatesCaller(address common.Address, caller bind.ContractCaller) (*SanityRatesCaller, error) {
	contract, err := bindSanityRates(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SanityRatesCaller{contract: contract}, nil
}

// NewSanityRatesTransactor creates a new write-only instance of SanityRates, bound to a specific deployed contract.
func NewSanityRatesTransactor(address common.Address, transactor bind.ContractTransactor) (*SanityRatesTransactor, error) {
	contract, err := bindSanityRates(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SanityRatesTransactor{contract: contract}, nil
}

// NewSanityRatesFilterer creates a new log filterer instance of SanityRates, bound to a specific deployed contract.
func NewSanityRatesFilterer(address common.Address, filterer bind.ContractFilterer) (*SanityRatesFilterer, error) {
	contract, err := bindSanityRates(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SanityRatesFilterer{contract: contract}, nil
}

// bindSanityRates binds a generic wrapper to an already deployed contract.
func bindSanityRates(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SanityRatesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SanityRates *SanityRatesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SanityRates.Contract.SanityRatesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SanityRates *SanityRatesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SanityRates.Contract.SanityRatesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SanityRates *SanityRatesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SanityRates.Contract.SanityRatesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SanityRates *SanityRatesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SanityRates.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SanityRates *SanityRatesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SanityRates.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SanityRates *SanityRatesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SanityRates.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_SanityRates *SanityRatesCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SanityRates.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_SanityRates *SanityRatesSession) Admin() (common.Address, error) {
	return _SanityRates.Contract.Admin(&_SanityRates.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_SanityRates *SanityRatesCallerSession) Admin() (common.Address, error) {
	return _SanityRates.Contract.Admin(&_SanityRates.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_SanityRates *SanityRatesCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _SanityRates.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_SanityRates *SanityRatesSession) GetAlerters() ([]common.Address, error) {
	return _SanityRates.Contract.GetAlerters(&_SanityRates.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_SanityRates *SanityRatesCallerSession) GetAlerters() ([]common.Address, error) {
	return _SanityRates.Contract.GetAlerters(&_SanityRates.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_SanityRates *SanityRatesCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _SanityRates.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_SanityRates *SanityRatesSession) GetOperators() ([]common.Address, error) {
	return _SanityRates.Contract.GetOperators(&_SanityRates.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_SanityRates *SanityRatesCallerSession) GetOperators() ([]common.Address, error) {
	return _SanityRates.Contract.GetOperators(&_SanityRates.CallOpts)
}

// GetSanityRate is a free data retrieval call binding the contract method 0xa58092b7.
//
// Solidity: function getSanityRate(address src, address dest) constant returns(uint256)
func (_SanityRates *SanityRatesCaller) GetSanityRate(opts *bind.CallOpts, src common.Address, dest common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SanityRates.contract.Call(opts, out, "getSanityRate", src, dest)
	return *ret0, err
}

// GetSanityRate is a free data retrieval call binding the contract method 0xa58092b7.
//
// Solidity: function getSanityRate(address src, address dest) constant returns(uint256)
func (_SanityRates *SanityRatesSession) GetSanityRate(src common.Address, dest common.Address) (*big.Int, error) {
	return _SanityRates.Contract.GetSanityRate(&_SanityRates.CallOpts, src, dest)
}

// GetSanityRate is a free data retrieval call binding the contract method 0xa58092b7.
//
// Solidity: function getSanityRate(address src, address dest) constant returns(uint256)
func (_SanityRates *SanityRatesCallerSession) GetSanityRate(src common.Address, dest common.Address) (*big.Int, error) {
	return _SanityRates.Contract.GetSanityRate(&_SanityRates.CallOpts, src, dest)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_SanityRates *SanityRatesCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SanityRates.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_SanityRates *SanityRatesSession) PendingAdmin() (common.Address, error) {
	return _SanityRates.Contract.PendingAdmin(&_SanityRates.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_SanityRates *SanityRatesCallerSession) PendingAdmin() (common.Address, error) {
	return _SanityRates.Contract.PendingAdmin(&_SanityRates.CallOpts)
}

// ReasonableDiffInBps is a free data retrieval call binding the contract method 0x5463a2e4.
//
// Solidity: function reasonableDiffInBps(address ) constant returns(uint256)
func (_SanityRates *SanityRatesCaller) ReasonableDiffInBps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SanityRates.contract.Call(opts, out, "reasonableDiffInBps", arg0)
	return *ret0, err
}

// ReasonableDiffInBps is a free data retrieval call binding the contract method 0x5463a2e4.
//
// Solidity: function reasonableDiffInBps(address ) constant returns(uint256)
func (_SanityRates *SanityRatesSession) ReasonableDiffInBps(arg0 common.Address) (*big.Int, error) {
	return _SanityRates.Contract.ReasonableDiffInBps(&_SanityRates.CallOpts, arg0)
}

// ReasonableDiffInBps is a free data retrieval call binding the contract method 0x5463a2e4.
//
// Solidity: function reasonableDiffInBps(address ) constant returns(uint256)
func (_SanityRates *SanityRatesCallerSession) ReasonableDiffInBps(arg0 common.Address) (*big.Int, error) {
	return _SanityRates.Contract.ReasonableDiffInBps(&_SanityRates.CallOpts, arg0)
}

// TokenRate is a free data retrieval call binding the contract method 0xc57fbf90.
//
// Solidity: function tokenRate(address ) constant returns(uint256)
func (_SanityRates *SanityRatesCaller) TokenRate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SanityRates.contract.Call(opts, out, "tokenRate", arg0)
	return *ret0, err
}

// TokenRate is a free data retrieval call binding the contract method 0xc57fbf90.
//
// Solidity: function tokenRate(address ) constant returns(uint256)
func (_SanityRates *SanityRatesSession) TokenRate(arg0 common.Address) (*big.Int, error) {
	return _SanityRates.Contract.TokenRate(&_SanityRates.CallOpts, arg0)
}

// TokenRate is a free data retrieval call binding the contract method 0xc57fbf90.
//
// Solidity: function tokenRate(address ) constant returns(uint256)
func (_SanityRates *SanityRatesCallerSession) TokenRate(arg0 common.Address) (*big.Int, error) {
	return _SanityRates.Contract.TokenRate(&_SanityRates.CallOpts, arg0)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_SanityRates *SanityRatesTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_SanityRates *SanityRatesSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.AddAlerter(&_SanityRates.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_SanityRates *SanityRatesTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.AddAlerter(&_SanityRates.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_SanityRates *SanityRatesTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_SanityRates *SanityRatesSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.AddOperator(&_SanityRates.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_SanityRates *SanityRatesTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.AddOperator(&_SanityRates.TransactOpts, newOperator)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SanityRates *SanityRatesTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SanityRates *SanityRatesSession) ClaimAdmin() (*types.Transaction, error) {
	return _SanityRates.Contract.ClaimAdmin(&_SanityRates.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_SanityRates *SanityRatesTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _SanityRates.Contract.ClaimAdmin(&_SanityRates.TransactOpts)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_SanityRates *SanityRatesTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_SanityRates *SanityRatesSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.RemoveAlerter(&_SanityRates.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_SanityRates *SanityRatesTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.RemoveAlerter(&_SanityRates.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SanityRates *SanityRatesTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SanityRates *SanityRatesSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.RemoveOperator(&_SanityRates.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SanityRates *SanityRatesTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.RemoveOperator(&_SanityRates.TransactOpts, operator)
}

// SetReasonableDiff is a paid mutator transaction binding the contract method 0x5c53ec59.
//
// Solidity: function setReasonableDiff(address[] srcs, uint256[] diff) returns()
func (_SanityRates *SanityRatesTransactor) SetReasonableDiff(opts *bind.TransactOpts, srcs []common.Address, diff []*big.Int) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "setReasonableDiff", srcs, diff)
}

// SetReasonableDiff is a paid mutator transaction binding the contract method 0x5c53ec59.
//
// Solidity: function setReasonableDiff(address[] srcs, uint256[] diff) returns()
func (_SanityRates *SanityRatesSession) SetReasonableDiff(srcs []common.Address, diff []*big.Int) (*types.Transaction, error) {
	return _SanityRates.Contract.SetReasonableDiff(&_SanityRates.TransactOpts, srcs, diff)
}

// SetReasonableDiff is a paid mutator transaction binding the contract method 0x5c53ec59.
//
// Solidity: function setReasonableDiff(address[] srcs, uint256[] diff) returns()
func (_SanityRates *SanityRatesTransactorSession) SetReasonableDiff(srcs []common.Address, diff []*big.Int) (*types.Transaction, error) {
	return _SanityRates.Contract.SetReasonableDiff(&_SanityRates.TransactOpts, srcs, diff)
}

// SetSanityRates is a paid mutator transaction binding the contract method 0xf5db370f.
//
// Solidity: function setSanityRates(address[] srcs, uint256[] rates) returns()
func (_SanityRates *SanityRatesTransactor) SetSanityRates(opts *bind.TransactOpts, srcs []common.Address, rates []*big.Int) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "setSanityRates", srcs, rates)
}

// SetSanityRates is a paid mutator transaction binding the contract method 0xf5db370f.
//
// Solidity: function setSanityRates(address[] srcs, uint256[] rates) returns()
func (_SanityRates *SanityRatesSession) SetSanityRates(srcs []common.Address, rates []*big.Int) (*types.Transaction, error) {
	return _SanityRates.Contract.SetSanityRates(&_SanityRates.TransactOpts, srcs, rates)
}

// SetSanityRates is a paid mutator transaction binding the contract method 0xf5db370f.
//
// Solidity: function setSanityRates(address[] srcs, uint256[] rates) returns()
func (_SanityRates *SanityRatesTransactorSession) SetSanityRates(srcs []common.Address, rates []*big.Int) (*types.Transaction, error) {
	return _SanityRates.Contract.SetSanityRates(&_SanityRates.TransactOpts, srcs, rates)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SanityRates *SanityRatesTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SanityRates *SanityRatesSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.TransferAdmin(&_SanityRates.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SanityRates *SanityRatesTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.TransferAdmin(&_SanityRates.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SanityRates *SanityRatesTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SanityRates *SanityRatesSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.TransferAdminQuickly(&_SanityRates.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_SanityRates *SanityRatesTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.TransferAdminQuickly(&_SanityRates.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_SanityRates *SanityRatesTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_SanityRates *SanityRatesSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.WithdrawEther(&_SanityRates.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_SanityRates *SanityRatesTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.WithdrawEther(&_SanityRates.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SanityRates *SanityRatesTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SanityRates.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SanityRates *SanityRatesSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.WithdrawToken(&_SanityRates.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_SanityRates *SanityRatesTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _SanityRates.Contract.WithdrawToken(&_SanityRates.TransactOpts, token, amount, sendTo)
}

// SanityRatesAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the SanityRates contract.
type SanityRatesAdminClaimedIterator struct {
	Event *SanityRatesAdminClaimed // Event containing the contract specifics and raw log

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
func (it *SanityRatesAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SanityRatesAdminClaimed)
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
		it.Event = new(SanityRatesAdminClaimed)
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
func (it *SanityRatesAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SanityRatesAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SanityRatesAdminClaimed represents a AdminClaimed event raised by the SanityRates contract.
type SanityRatesAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_SanityRates *SanityRatesFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*SanityRatesAdminClaimedIterator, error) {

	logs, sub, err := _SanityRates.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &SanityRatesAdminClaimedIterator{contract: _SanityRates.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_SanityRates *SanityRatesFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *SanityRatesAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _SanityRates.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SanityRatesAdminClaimed)
				if err := _SanityRates.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_SanityRates *SanityRatesFilterer) ParseAdminClaimed(log types.Log) (*SanityRatesAdminClaimed, error) {
	event := new(SanityRatesAdminClaimed)
	if err := _SanityRates.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SanityRatesAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the SanityRates contract.
type SanityRatesAlerterAddedIterator struct {
	Event *SanityRatesAlerterAdded // Event containing the contract specifics and raw log

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
func (it *SanityRatesAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SanityRatesAlerterAdded)
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
		it.Event = new(SanityRatesAlerterAdded)
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
func (it *SanityRatesAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SanityRatesAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SanityRatesAlerterAdded represents a AlerterAdded event raised by the SanityRates contract.
type SanityRatesAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_SanityRates *SanityRatesFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*SanityRatesAlerterAddedIterator, error) {

	logs, sub, err := _SanityRates.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &SanityRatesAlerterAddedIterator{contract: _SanityRates.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_SanityRates *SanityRatesFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *SanityRatesAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _SanityRates.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SanityRatesAlerterAdded)
				if err := _SanityRates.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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
func (_SanityRates *SanityRatesFilterer) ParseAlerterAdded(log types.Log) (*SanityRatesAlerterAdded, error) {
	event := new(SanityRatesAlerterAdded)
	if err := _SanityRates.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SanityRatesEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the SanityRates contract.
type SanityRatesEtherWithdrawIterator struct {
	Event *SanityRatesEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *SanityRatesEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SanityRatesEtherWithdraw)
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
		it.Event = new(SanityRatesEtherWithdraw)
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
func (it *SanityRatesEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SanityRatesEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SanityRatesEtherWithdraw represents a EtherWithdraw event raised by the SanityRates contract.
type SanityRatesEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_SanityRates *SanityRatesFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*SanityRatesEtherWithdrawIterator, error) {

	logs, sub, err := _SanityRates.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &SanityRatesEtherWithdrawIterator{contract: _SanityRates.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_SanityRates *SanityRatesFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *SanityRatesEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _SanityRates.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SanityRatesEtherWithdraw)
				if err := _SanityRates.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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
func (_SanityRates *SanityRatesFilterer) ParseEtherWithdraw(log types.Log) (*SanityRatesEtherWithdraw, error) {
	event := new(SanityRatesEtherWithdraw)
	if err := _SanityRates.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SanityRatesOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the SanityRates contract.
type SanityRatesOperatorAddedIterator struct {
	Event *SanityRatesOperatorAdded // Event containing the contract specifics and raw log

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
func (it *SanityRatesOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SanityRatesOperatorAdded)
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
		it.Event = new(SanityRatesOperatorAdded)
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
func (it *SanityRatesOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SanityRatesOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SanityRatesOperatorAdded represents a OperatorAdded event raised by the SanityRates contract.
type SanityRatesOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_SanityRates *SanityRatesFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*SanityRatesOperatorAddedIterator, error) {

	logs, sub, err := _SanityRates.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &SanityRatesOperatorAddedIterator{contract: _SanityRates.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_SanityRates *SanityRatesFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *SanityRatesOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _SanityRates.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SanityRatesOperatorAdded)
				if err := _SanityRates.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
func (_SanityRates *SanityRatesFilterer) ParseOperatorAdded(log types.Log) (*SanityRatesOperatorAdded, error) {
	event := new(SanityRatesOperatorAdded)
	if err := _SanityRates.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SanityRatesTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the SanityRates contract.
type SanityRatesTokenWithdrawIterator struct {
	Event *SanityRatesTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *SanityRatesTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SanityRatesTokenWithdraw)
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
		it.Event = new(SanityRatesTokenWithdraw)
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
func (it *SanityRatesTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SanityRatesTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SanityRatesTokenWithdraw represents a TokenWithdraw event raised by the SanityRates contract.
type SanityRatesTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_SanityRates *SanityRatesFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*SanityRatesTokenWithdrawIterator, error) {

	logs, sub, err := _SanityRates.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &SanityRatesTokenWithdrawIterator{contract: _SanityRates.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_SanityRates *SanityRatesFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *SanityRatesTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _SanityRates.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SanityRatesTokenWithdraw)
				if err := _SanityRates.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_SanityRates *SanityRatesFilterer) ParseTokenWithdraw(log types.Log) (*SanityRatesTokenWithdraw, error) {
	event := new(SanityRatesTokenWithdraw)
	if err := _SanityRates.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SanityRatesTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the SanityRates contract.
type SanityRatesTransferAdminPendingIterator struct {
	Event *SanityRatesTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *SanityRatesTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SanityRatesTransferAdminPending)
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
		it.Event = new(SanityRatesTransferAdminPending)
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
func (it *SanityRatesTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SanityRatesTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SanityRatesTransferAdminPending represents a TransferAdminPending event raised by the SanityRates contract.
type SanityRatesTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_SanityRates *SanityRatesFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*SanityRatesTransferAdminPendingIterator, error) {

	logs, sub, err := _SanityRates.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &SanityRatesTransferAdminPendingIterator{contract: _SanityRates.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_SanityRates *SanityRatesFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *SanityRatesTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _SanityRates.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SanityRatesTransferAdminPending)
				if err := _SanityRates.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_SanityRates *SanityRatesFilterer) ParseTransferAdminPending(log types.Log) (*SanityRatesTransferAdminPending, error) {
	event := new(SanityRatesTransferAdminPending)
	if err := _SanityRates.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}
