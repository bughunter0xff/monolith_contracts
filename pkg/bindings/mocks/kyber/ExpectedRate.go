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

// ExpectedRateABI is the input ABI used to generate the binding from.
const ExpectedRateABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFactor\",\"type\":\"uint256\"}],\"name\":\"setQuantityFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"worstCaseRateFactorInBps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quantityFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetwork\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"srcQty\",\"type\":\"uint256\"},{\"name\":\"usePermissionless\",\"type\":\"bool\"}],\"name\":\"getExpectedRate\",\"outputs\":[{\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"name\":\"slippageRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"setWorstCaseRateFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"knc\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"currentKncToEthRate\",\"type\":\"uint256\"}],\"name\":\"checkKncArbitrageRate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kyberNetwork\",\"type\":\"address\"},{\"name\":\"_knc\",\"type\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"oldFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"QuantityFactorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newMin\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"oldMin\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"MinSlippageFactorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"}]"

// ExpectedRateBin is the compiled bytecode used for deploying new contracts.
var ExpectedRateBin = "0x606060405260026008556032600955341561001957600080fd5b60405160608061170d83398101604052808051919060200180519190602001805160008054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061006d57600080fd5b600160a060020a038216151561008257600080fd5b600160a060020a038316151561009757600080fd5b60078054600160a060020a03948516600160a060020a0319918216179091556000805492851692821692909217909155600a805492909316911617905561162a806100e36000396000f30060606040526004361061010e5763ffffffff60e060020a60003504166301a12fd38114610113578063267822471461013457806327a099d8146101635780633ccdbb28146101c9578063408ee7fe146101f257806375829def146102115780637658c5741461023057806377f50f97146102465780637acc8678146102595780637c423f54146102785780639870d7fe1461028b5780639bc72d5f146102aa578063a7de9c63146102cf578063ac8a584a146102e2578063b78b842d14610301578063ce56c45414610314578063d38d2bea14610336578063d4fac45d1461037b578063dcb46e38146103a0578063e61387e0146103b6578063e853cda3146103c9578063f851a440146103f3575b600080fd5b341561011e57600080fd5b610132600160a060020a0360043516610406565b005b341561013f57600080fd5b610147610576565b604051600160a060020a03909116815260200160405180910390f35b341561016e57600080fd5b610176610585565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101b557808201518382015260200161019d565b505050509050019250505060405180910390f35b34156101d457600080fd5b610132600160a060020a0360043581169060243590604435166105ee565b34156101fd57600080fd5b610132600160a060020a03600435166106e5565b341561021c57600080fd5b610132600160a060020a03600435166107e1565b341561023b57600080fd5b61013260043561087c565b341561025157600080fd5b610132610908565b341561026457600080fd5b610132600160a060020a03600435166109a2565b341561028357600080fd5b610176610a84565b341561029657600080fd5b610132600160a060020a0360043516610aea565b34156102b557600080fd5b6102bd610bba565b60405190815260200160405180910390f35b34156102da57600080fd5b6102bd610bc0565b34156102ed57600080fd5b610132600160a060020a0360043516610bc6565b341561030c57600080fd5b610147610d32565b341561031f57600080fd5b610132600435600160a060020a0360243516610d41565b341561034157600080fd5b610363600160a060020a03600435811690602435166044356064351515610dd4565b60405191825260208201526040908101905180910390f35b341561038657600080fd5b6102bd600160a060020a0360043581169060243516610f13565b34156103ab57600080fd5b610132600435610fc5565b34156103c157600080fd5b610147611052565b34156103d457600080fd5b6103df600435611061565b604051901515815260200160405180910390f35b34156103fe57600080fd5b610147611100565b6000805433600160a060020a0390811691161461042257600080fd5b600160a060020a03821660009081526003602052604090205460ff16151561044957600080fd5b50600160a060020a0381166000908152600360205260408120805460ff191690555b6005548110156105725781600160a060020a031660058281548110151561048e57fe5b600091825260209091200154600160a060020a0316141561056a576005805460001981019081106104bb57fe5b60009182526020909120015460058054600160a060020a0390921691839081106104e157fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600580549061051d9060001983016115a5565b507f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762826000604051600160a060020a039092168252151560208201526040908101905180910390a1610572565b60010161046b565b5050565b600154600160a060020a031681565b61058d6115ce565b60048054806020026020016040519081016040528092919081815260200182805480156105e357602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116105c5575b505050505090505b90565b60005433600160a060020a0390811691161461060957600080fd5b82600160a060020a031663a9059cbb828460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561066657600080fd5b6102c65a03f1151561067757600080fd5b50505060405180519050151561068c57600080fd5b7f72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6838383604051600160a060020a03938416815260208101929092529091166040808301919091526060909101905180910390a1505050565b60005433600160a060020a0390811691161461070057600080fd5b600160a060020a03811660009081526003602052604090205460ff161561072657600080fd5b6005546032901061073657600080fd5b7f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600360205260409020805460ff1916600190811790915560058054909181016107b583826115a5565b5060009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a039081169116146107fc57600080fd5b600160a060020a038116151561081157600080fd5b6001547f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4090600160a060020a0316604051600160a060020a03909116815260200160405180910390a160018054600160a060020a031916600160a060020a0392909216919091179055565b600160a060020a03331660009081526002602052604090205460ff1615156108a357600080fd5b60648111156108b157600080fd5b7fd0f6fc40d497232b5aab1b7a34ea00ea45886e52d2fed39ad62af798a870fae381600854336040519283526020830191909152600160a060020a03166040808301919091526060909101905180910390a1600855565b60015433600160a060020a0390811691161461092357600080fd5b6001546000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed91600160a060020a039081169116604051600160a060020a039283168152911660208201526040908101905180910390a16001805460008054600160a060020a0319908116600160a060020a03841617909155169055565b60005433600160a060020a039081169116146109bd57600080fd5b600160a060020a03811615156109d257600080fd5b7f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4081604051600160a060020a03909116815260200160405180910390a16000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed908290600160a060020a0316604051600160a060020a039283168152911660208201526040908101905180910390a160008054600160a060020a031916600160a060020a0392909216919091179055565b610a8c6115ce565b60058054806020026020016040519081016040528092919081815260200182805480156105e357602002820191906000526020600020908154600160a060020a031681526001909101906020018083116105c5575050505050905090565b60005433600160a060020a03908116911614610b0557600080fd5b600160a060020a03811660009081526002602052604090205460ff1615610b2b57600080fd5b60045460329010610b3b57600080fd5b7f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600260205260409020805460ff1916600190811790915560048054909181016107b583826115a5565b60095481565b60085481565b6000805433600160a060020a03908116911614610be257600080fd5b600160a060020a03821660009081526002602052604090205460ff161515610c0957600080fd5b50600160a060020a0381166000908152600260205260408120805460ff191690555b6004548110156105725781600160a060020a0316600482815481101515610c4e57fe5b600091825260209091200154600160a060020a03161415610d2a57600480546000198101908110610c7b57fe5b60009182526020909120015460048054600160a060020a039092169183908110610ca157fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600480546000190190610cdd90826115a5565b507f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b826000604051600160a060020a039092168252151560208201526040908101905180910390a1610572565b600101610c2b565b600754600160a060020a031681565b60005433600160a060020a03908116911614610d5c57600080fd5b600160a060020a03811682156108fc0283604051600060405180830381858888f193505050501515610d8d57600080fd5b7fec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de8282604051918252600160a060020a031660208201526040908101905180910390a15050565b600080600080600854600014151515610dec57600080fd5b6b204fce5e3e25026110000000861115610e0557600080fd5b6008546b204fce5e3e250261100000009087021115610e2357600080fd5b851515610e2f57600195505b60009150610e3f8888888861110f565b909550935091508115610e585760009350839250610f08565b831515610e6e57610e6b8888888861122a565b93505b600a54600160a060020a038981169116148015610ea75750600160a060020a03871673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee145b8015610eba5750670de0b6b3a764000086145b15610ed257610ec884611061565b15610ed257600093505b69d3c21bcecceda1000000841115610ef05760009350839250610f08565b50600954612710908103840204808310610f08578092505b505094509492505050565b6000600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415610f4b5750600160a060020a03811631610fbf565b82600160a060020a03166370a082318360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610fa257600080fd5b6102c65a03f11515610fb357600080fd5b50505060405180519150505b92915050565b600160a060020a03331660009081526002602052604090205460ff161515610fec57600080fd5b612710811115610ffb57600080fd5b7f4357e20f1241d972328c5b3239d9ef4ac96f0f4fce8e10fd3bf9053690dad0ac81600954336040519283526020830191909152600160a060020a03166040808301919091526060909101905180910390a1600955565b600a54600160a060020a031681565b600a546000908190819061109e9073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee90600160a060020a0316670de0b6b3a76400006001610dd4565b90925090508115156110b357600192506110f9565b69d3c21bcecceda100000082118015906110d7575069d3c21bcecceda10000008411155b15156110e257600080fd5b6ec097ce7bc90715b34b9f10000000008285021192505b5050919050565b600054600160a060020a031681565b60008060008084611176576040517f66696e6442657374526174654f6e6c795065726d697373696f6e28616464726581527f73732c616464726573732c75696e743235362900000000000000000000000000602082015260330160405180910390206111ce565b6040517f66696e64426573745261746528616464726573732c616464726573732c75696e81527f7432353629000000000000000000000000000000000000000000000000000000602082015260250160405180910390205b90506111dc888888846113c4565b909450925083156111f757600193506000925082915061121f565b60085460011461121b5761121188886008548902846113c4565b909450915061121f565b8291505b509450945094915050565b6007546000908190819081908190600160a060020a0316630c235d968a73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8a8a866040516040015260405160e060020a63ffffffff8716028152600160a060020a039485166004820152929093166024830152604482015290151560648201526084016040805180830381600087803b15156112b957600080fd5b6102c65a03f115156112ca57600080fd5b505050604051805190602001805191955090935061130090508973eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8986611427565b600754909150600160a060020a0316630c235d9673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8a848a60006040516040015260405160e060020a63ffffffff8716028152600160a060020a039485166004820152929093166024830152604482015290151560648201526084016040805180830381600087803b151561138857600080fd5b6102c65a03f1151561139957600080fd5b5050506040518051906020018051670de0b6b3a76400009502949094049a9950505050505050505050565b6007546000908190600160a060020a03168160405185815288600482015287602482015286604482015260a48101604052604060648201606483865afa91506084810151604091909152925060018114610f085760019350505094509492505050565b60006114458361143687611450565b61143f87611450565b85611514565b90505b949350505050565b600080600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415611481576012915061150e565b50600160a060020a03821660009081526006602052604090205480151561150a5782600160a060020a031663313ce5676000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156114e857600080fd5b6102c65a03f115156114f957600080fd5b50505060405180519050915061150e565b8091505b50919050565b60006b204fce5e3e2502611000000085111561152f57600080fd5b69d3c21bcecceda100000082111561154657600080fd5b838310611579576012848403111561155d57600080fd5b670de0b6b3a7640000858302858503600a0a025b049050611448565b6012838503111561158957600080fd5b828403600a0a670de0b6b3a76400000282860281151561157157fe5b8154818355818115116115c9576000838152602090206115c99181019083016115e0565b505050565b60206040519081016040526000815290565b6105eb91905b808211156115fa57600081556001016115e6565b50905600a165627a7a72305820533eeb286fb9be65b1a33b676fb376d1252ea7e44cf66f70a8c2b7171d861a9a0029"

// DeployExpectedRate deploys a new Ethereum contract, binding an instance of ExpectedRate to it.
func DeployExpectedRate(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetwork common.Address, _knc common.Address, _admin common.Address) (common.Address, *types.Transaction, *ExpectedRate, error) {
	parsed, err := abi.JSON(strings.NewReader(ExpectedRateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExpectedRateBin), backend, _kyberNetwork, _knc, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExpectedRate{ExpectedRateCaller: ExpectedRateCaller{contract: contract}, ExpectedRateTransactor: ExpectedRateTransactor{contract: contract}, ExpectedRateFilterer: ExpectedRateFilterer{contract: contract}}, nil
}

// ExpectedRate is an auto generated Go binding around an Ethereum contract.
type ExpectedRate struct {
	ExpectedRateCaller     // Read-only binding to the contract
	ExpectedRateTransactor // Write-only binding to the contract
	ExpectedRateFilterer   // Log filterer for contract events
}

// ExpectedRateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExpectedRateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpectedRateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExpectedRateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpectedRateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExpectedRateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpectedRateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExpectedRateSession struct {
	Contract     *ExpectedRate     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExpectedRateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExpectedRateCallerSession struct {
	Contract *ExpectedRateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExpectedRateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExpectedRateTransactorSession struct {
	Contract     *ExpectedRateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExpectedRateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExpectedRateRaw struct {
	Contract *ExpectedRate // Generic contract binding to access the raw methods on
}

// ExpectedRateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExpectedRateCallerRaw struct {
	Contract *ExpectedRateCaller // Generic read-only contract binding to access the raw methods on
}

// ExpectedRateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExpectedRateTransactorRaw struct {
	Contract *ExpectedRateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExpectedRate creates a new instance of ExpectedRate, bound to a specific deployed contract.
func NewExpectedRate(address common.Address, backend bind.ContractBackend) (*ExpectedRate, error) {
	contract, err := bindExpectedRate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExpectedRate{ExpectedRateCaller: ExpectedRateCaller{contract: contract}, ExpectedRateTransactor: ExpectedRateTransactor{contract: contract}, ExpectedRateFilterer: ExpectedRateFilterer{contract: contract}}, nil
}

// NewExpectedRateCaller creates a new read-only instance of ExpectedRate, bound to a specific deployed contract.
func NewExpectedRateCaller(address common.Address, caller bind.ContractCaller) (*ExpectedRateCaller, error) {
	contract, err := bindExpectedRate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExpectedRateCaller{contract: contract}, nil
}

// NewExpectedRateTransactor creates a new write-only instance of ExpectedRate, bound to a specific deployed contract.
func NewExpectedRateTransactor(address common.Address, transactor bind.ContractTransactor) (*ExpectedRateTransactor, error) {
	contract, err := bindExpectedRate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExpectedRateTransactor{contract: contract}, nil
}

// NewExpectedRateFilterer creates a new log filterer instance of ExpectedRate, bound to a specific deployed contract.
func NewExpectedRateFilterer(address common.Address, filterer bind.ContractFilterer) (*ExpectedRateFilterer, error) {
	contract, err := bindExpectedRate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExpectedRateFilterer{contract: contract}, nil
}

// bindExpectedRate binds a generic wrapper to an already deployed contract.
func bindExpectedRate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExpectedRateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExpectedRate *ExpectedRateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExpectedRate.Contract.ExpectedRateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExpectedRate *ExpectedRateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpectedRate.Contract.ExpectedRateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExpectedRate *ExpectedRateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExpectedRate.Contract.ExpectedRateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExpectedRate *ExpectedRateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExpectedRate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExpectedRate *ExpectedRateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpectedRate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExpectedRate *ExpectedRateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExpectedRate.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_ExpectedRate *ExpectedRateCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_ExpectedRate *ExpectedRateSession) Admin() (common.Address, error) {
	return _ExpectedRate.Contract.Admin(&_ExpectedRate.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_ExpectedRate *ExpectedRateCallerSession) Admin() (common.Address, error) {
	return _ExpectedRate.Contract.Admin(&_ExpectedRate.CallOpts)
}

// CheckKncArbitrageRate is a free data retrieval call binding the contract method 0xe853cda3.
//
// Solidity: function checkKncArbitrageRate(uint256 currentKncToEthRate) constant returns(bool)
func (_ExpectedRate *ExpectedRateCaller) CheckKncArbitrageRate(opts *bind.CallOpts, currentKncToEthRate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "checkKncArbitrageRate", currentKncToEthRate)
	return *ret0, err
}

// CheckKncArbitrageRate is a free data retrieval call binding the contract method 0xe853cda3.
//
// Solidity: function checkKncArbitrageRate(uint256 currentKncToEthRate) constant returns(bool)
func (_ExpectedRate *ExpectedRateSession) CheckKncArbitrageRate(currentKncToEthRate *big.Int) (bool, error) {
	return _ExpectedRate.Contract.CheckKncArbitrageRate(&_ExpectedRate.CallOpts, currentKncToEthRate)
}

// CheckKncArbitrageRate is a free data retrieval call binding the contract method 0xe853cda3.
//
// Solidity: function checkKncArbitrageRate(uint256 currentKncToEthRate) constant returns(bool)
func (_ExpectedRate *ExpectedRateCallerSession) CheckKncArbitrageRate(currentKncToEthRate *big.Int) (bool, error) {
	return _ExpectedRate.Contract.CheckKncArbitrageRate(&_ExpectedRate.CallOpts, currentKncToEthRate)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_ExpectedRate *ExpectedRateCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_ExpectedRate *ExpectedRateSession) GetAlerters() ([]common.Address, error) {
	return _ExpectedRate.Contract.GetAlerters(&_ExpectedRate.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_ExpectedRate *ExpectedRateCallerSession) GetAlerters() ([]common.Address, error) {
	return _ExpectedRate.Contract.GetAlerters(&_ExpectedRate.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_ExpectedRate *ExpectedRateCaller) GetBalance(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "getBalance", token, user)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_ExpectedRate *ExpectedRateSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _ExpectedRate.Contract.GetBalance(&_ExpectedRate.CallOpts, token, user)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_ExpectedRate *ExpectedRateCallerSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _ExpectedRate.Contract.GetBalance(&_ExpectedRate.CallOpts, token, user)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0xd38d2bea.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty, bool usePermissionless) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_ExpectedRate *ExpectedRateCaller) GetExpectedRate(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int, usePermissionless bool) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	ret := new(struct {
		ExpectedRate *big.Int
		SlippageRate *big.Int
	})
	out := ret
	err := _ExpectedRate.contract.Call(opts, out, "getExpectedRate", src, dest, srcQty, usePermissionless)
	return *ret, err
}

// GetExpectedRate is a free data retrieval call binding the contract method 0xd38d2bea.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty, bool usePermissionless) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_ExpectedRate *ExpectedRateSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int, usePermissionless bool) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _ExpectedRate.Contract.GetExpectedRate(&_ExpectedRate.CallOpts, src, dest, srcQty, usePermissionless)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0xd38d2bea.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty, bool usePermissionless) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_ExpectedRate *ExpectedRateCallerSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int, usePermissionless bool) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _ExpectedRate.Contract.GetExpectedRate(&_ExpectedRate.CallOpts, src, dest, srcQty, usePermissionless)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_ExpectedRate *ExpectedRateCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_ExpectedRate *ExpectedRateSession) GetOperators() ([]common.Address, error) {
	return _ExpectedRate.Contract.GetOperators(&_ExpectedRate.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_ExpectedRate *ExpectedRateCallerSession) GetOperators() ([]common.Address, error) {
	return _ExpectedRate.Contract.GetOperators(&_ExpectedRate.CallOpts)
}

// Knc is a free data retrieval call binding the contract method 0xe61387e0.
//
// Solidity: function knc() constant returns(address)
func (_ExpectedRate *ExpectedRateCaller) Knc(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "knc")
	return *ret0, err
}

// Knc is a free data retrieval call binding the contract method 0xe61387e0.
//
// Solidity: function knc() constant returns(address)
func (_ExpectedRate *ExpectedRateSession) Knc() (common.Address, error) {
	return _ExpectedRate.Contract.Knc(&_ExpectedRate.CallOpts)
}

// Knc is a free data retrieval call binding the contract method 0xe61387e0.
//
// Solidity: function knc() constant returns(address)
func (_ExpectedRate *ExpectedRateCallerSession) Knc() (common.Address, error) {
	return _ExpectedRate.Contract.Knc(&_ExpectedRate.CallOpts)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_ExpectedRate *ExpectedRateCaller) KyberNetwork(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "kyberNetwork")
	return *ret0, err
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_ExpectedRate *ExpectedRateSession) KyberNetwork() (common.Address, error) {
	return _ExpectedRate.Contract.KyberNetwork(&_ExpectedRate.CallOpts)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_ExpectedRate *ExpectedRateCallerSession) KyberNetwork() (common.Address, error) {
	return _ExpectedRate.Contract.KyberNetwork(&_ExpectedRate.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_ExpectedRate *ExpectedRateCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_ExpectedRate *ExpectedRateSession) PendingAdmin() (common.Address, error) {
	return _ExpectedRate.Contract.PendingAdmin(&_ExpectedRate.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_ExpectedRate *ExpectedRateCallerSession) PendingAdmin() (common.Address, error) {
	return _ExpectedRate.Contract.PendingAdmin(&_ExpectedRate.CallOpts)
}

// QuantityFactor is a free data retrieval call binding the contract method 0xa7de9c63.
//
// Solidity: function quantityFactor() constant returns(uint256)
func (_ExpectedRate *ExpectedRateCaller) QuantityFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "quantityFactor")
	return *ret0, err
}

// QuantityFactor is a free data retrieval call binding the contract method 0xa7de9c63.
//
// Solidity: function quantityFactor() constant returns(uint256)
func (_ExpectedRate *ExpectedRateSession) QuantityFactor() (*big.Int, error) {
	return _ExpectedRate.Contract.QuantityFactor(&_ExpectedRate.CallOpts)
}

// QuantityFactor is a free data retrieval call binding the contract method 0xa7de9c63.
//
// Solidity: function quantityFactor() constant returns(uint256)
func (_ExpectedRate *ExpectedRateCallerSession) QuantityFactor() (*big.Int, error) {
	return _ExpectedRate.Contract.QuantityFactor(&_ExpectedRate.CallOpts)
}

// WorstCaseRateFactorInBps is a free data retrieval call binding the contract method 0x9bc72d5f.
//
// Solidity: function worstCaseRateFactorInBps() constant returns(uint256)
func (_ExpectedRate *ExpectedRateCaller) WorstCaseRateFactorInBps(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExpectedRate.contract.Call(opts, out, "worstCaseRateFactorInBps")
	return *ret0, err
}

// WorstCaseRateFactorInBps is a free data retrieval call binding the contract method 0x9bc72d5f.
//
// Solidity: function worstCaseRateFactorInBps() constant returns(uint256)
func (_ExpectedRate *ExpectedRateSession) WorstCaseRateFactorInBps() (*big.Int, error) {
	return _ExpectedRate.Contract.WorstCaseRateFactorInBps(&_ExpectedRate.CallOpts)
}

// WorstCaseRateFactorInBps is a free data retrieval call binding the contract method 0x9bc72d5f.
//
// Solidity: function worstCaseRateFactorInBps() constant returns(uint256)
func (_ExpectedRate *ExpectedRateCallerSession) WorstCaseRateFactorInBps() (*big.Int, error) {
	return _ExpectedRate.Contract.WorstCaseRateFactorInBps(&_ExpectedRate.CallOpts)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_ExpectedRate *ExpectedRateTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_ExpectedRate *ExpectedRateSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.AddAlerter(&_ExpectedRate.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.AddAlerter(&_ExpectedRate.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_ExpectedRate *ExpectedRateTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_ExpectedRate *ExpectedRateSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.AddOperator(&_ExpectedRate.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.AddOperator(&_ExpectedRate.TransactOpts, newOperator)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_ExpectedRate *ExpectedRateTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_ExpectedRate *ExpectedRateSession) ClaimAdmin() (*types.Transaction, error) {
	return _ExpectedRate.Contract.ClaimAdmin(&_ExpectedRate.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_ExpectedRate *ExpectedRateTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _ExpectedRate.Contract.ClaimAdmin(&_ExpectedRate.TransactOpts)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_ExpectedRate *ExpectedRateTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_ExpectedRate *ExpectedRateSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.RemoveAlerter(&_ExpectedRate.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.RemoveAlerter(&_ExpectedRate.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_ExpectedRate *ExpectedRateTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_ExpectedRate *ExpectedRateSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.RemoveOperator(&_ExpectedRate.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.RemoveOperator(&_ExpectedRate.TransactOpts, operator)
}

// SetQuantityFactor is a paid mutator transaction binding the contract method 0x7658c574.
//
// Solidity: function setQuantityFactor(uint256 newFactor) returns()
func (_ExpectedRate *ExpectedRateTransactor) SetQuantityFactor(opts *bind.TransactOpts, newFactor *big.Int) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "setQuantityFactor", newFactor)
}

// SetQuantityFactor is a paid mutator transaction binding the contract method 0x7658c574.
//
// Solidity: function setQuantityFactor(uint256 newFactor) returns()
func (_ExpectedRate *ExpectedRateSession) SetQuantityFactor(newFactor *big.Int) (*types.Transaction, error) {
	return _ExpectedRate.Contract.SetQuantityFactor(&_ExpectedRate.TransactOpts, newFactor)
}

// SetQuantityFactor is a paid mutator transaction binding the contract method 0x7658c574.
//
// Solidity: function setQuantityFactor(uint256 newFactor) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) SetQuantityFactor(newFactor *big.Int) (*types.Transaction, error) {
	return _ExpectedRate.Contract.SetQuantityFactor(&_ExpectedRate.TransactOpts, newFactor)
}

// SetWorstCaseRateFactor is a paid mutator transaction binding the contract method 0xdcb46e38.
//
// Solidity: function setWorstCaseRateFactor(uint256 bps) returns()
func (_ExpectedRate *ExpectedRateTransactor) SetWorstCaseRateFactor(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "setWorstCaseRateFactor", bps)
}

// SetWorstCaseRateFactor is a paid mutator transaction binding the contract method 0xdcb46e38.
//
// Solidity: function setWorstCaseRateFactor(uint256 bps) returns()
func (_ExpectedRate *ExpectedRateSession) SetWorstCaseRateFactor(bps *big.Int) (*types.Transaction, error) {
	return _ExpectedRate.Contract.SetWorstCaseRateFactor(&_ExpectedRate.TransactOpts, bps)
}

// SetWorstCaseRateFactor is a paid mutator transaction binding the contract method 0xdcb46e38.
//
// Solidity: function setWorstCaseRateFactor(uint256 bps) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) SetWorstCaseRateFactor(bps *big.Int) (*types.Transaction, error) {
	return _ExpectedRate.Contract.SetWorstCaseRateFactor(&_ExpectedRate.TransactOpts, bps)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_ExpectedRate *ExpectedRateTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_ExpectedRate *ExpectedRateSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.TransferAdmin(&_ExpectedRate.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.TransferAdmin(&_ExpectedRate.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_ExpectedRate *ExpectedRateTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_ExpectedRate *ExpectedRateSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.TransferAdminQuickly(&_ExpectedRate.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.TransferAdminQuickly(&_ExpectedRate.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_ExpectedRate *ExpectedRateTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_ExpectedRate *ExpectedRateSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.WithdrawEther(&_ExpectedRate.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.WithdrawEther(&_ExpectedRate.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_ExpectedRate *ExpectedRateTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _ExpectedRate.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_ExpectedRate *ExpectedRateSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.WithdrawToken(&_ExpectedRate.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_ExpectedRate *ExpectedRateTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _ExpectedRate.Contract.WithdrawToken(&_ExpectedRate.TransactOpts, token, amount, sendTo)
}

// ExpectedRateAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the ExpectedRate contract.
type ExpectedRateAdminClaimedIterator struct {
	Event *ExpectedRateAdminClaimed // Event containing the contract specifics and raw log

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
func (it *ExpectedRateAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpectedRateAdminClaimed)
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
		it.Event = new(ExpectedRateAdminClaimed)
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
func (it *ExpectedRateAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpectedRateAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpectedRateAdminClaimed represents a AdminClaimed event raised by the ExpectedRate contract.
type ExpectedRateAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_ExpectedRate *ExpectedRateFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*ExpectedRateAdminClaimedIterator, error) {

	logs, sub, err := _ExpectedRate.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &ExpectedRateAdminClaimedIterator{contract: _ExpectedRate.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_ExpectedRate *ExpectedRateFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *ExpectedRateAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _ExpectedRate.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpectedRateAdminClaimed)
				if err := _ExpectedRate.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_ExpectedRate *ExpectedRateFilterer) ParseAdminClaimed(log types.Log) (*ExpectedRateAdminClaimed, error) {
	event := new(ExpectedRateAdminClaimed)
	if err := _ExpectedRate.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExpectedRateAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the ExpectedRate contract.
type ExpectedRateAlerterAddedIterator struct {
	Event *ExpectedRateAlerterAdded // Event containing the contract specifics and raw log

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
func (it *ExpectedRateAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpectedRateAlerterAdded)
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
		it.Event = new(ExpectedRateAlerterAdded)
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
func (it *ExpectedRateAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpectedRateAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpectedRateAlerterAdded represents a AlerterAdded event raised by the ExpectedRate contract.
type ExpectedRateAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_ExpectedRate *ExpectedRateFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*ExpectedRateAlerterAddedIterator, error) {

	logs, sub, err := _ExpectedRate.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &ExpectedRateAlerterAddedIterator{contract: _ExpectedRate.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_ExpectedRate *ExpectedRateFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *ExpectedRateAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _ExpectedRate.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpectedRateAlerterAdded)
				if err := _ExpectedRate.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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
func (_ExpectedRate *ExpectedRateFilterer) ParseAlerterAdded(log types.Log) (*ExpectedRateAlerterAdded, error) {
	event := new(ExpectedRateAlerterAdded)
	if err := _ExpectedRate.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExpectedRateEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the ExpectedRate contract.
type ExpectedRateEtherWithdrawIterator struct {
	Event *ExpectedRateEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *ExpectedRateEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpectedRateEtherWithdraw)
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
		it.Event = new(ExpectedRateEtherWithdraw)
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
func (it *ExpectedRateEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpectedRateEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpectedRateEtherWithdraw represents a EtherWithdraw event raised by the ExpectedRate contract.
type ExpectedRateEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_ExpectedRate *ExpectedRateFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*ExpectedRateEtherWithdrawIterator, error) {

	logs, sub, err := _ExpectedRate.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &ExpectedRateEtherWithdrawIterator{contract: _ExpectedRate.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_ExpectedRate *ExpectedRateFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *ExpectedRateEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _ExpectedRate.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpectedRateEtherWithdraw)
				if err := _ExpectedRate.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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
func (_ExpectedRate *ExpectedRateFilterer) ParseEtherWithdraw(log types.Log) (*ExpectedRateEtherWithdraw, error) {
	event := new(ExpectedRateEtherWithdraw)
	if err := _ExpectedRate.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExpectedRateMinSlippageFactorSetIterator is returned from FilterMinSlippageFactorSet and is used to iterate over the raw logs and unpacked data for MinSlippageFactorSet events raised by the ExpectedRate contract.
type ExpectedRateMinSlippageFactorSetIterator struct {
	Event *ExpectedRateMinSlippageFactorSet // Event containing the contract specifics and raw log

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
func (it *ExpectedRateMinSlippageFactorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpectedRateMinSlippageFactorSet)
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
		it.Event = new(ExpectedRateMinSlippageFactorSet)
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
func (it *ExpectedRateMinSlippageFactorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpectedRateMinSlippageFactorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpectedRateMinSlippageFactorSet represents a MinSlippageFactorSet event raised by the ExpectedRate contract.
type ExpectedRateMinSlippageFactorSet struct {
	NewMin *big.Int
	OldMin *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinSlippageFactorSet is a free log retrieval operation binding the contract event 0x4357e20f1241d972328c5b3239d9ef4ac96f0f4fce8e10fd3bf9053690dad0ac.
//
// Solidity: event MinSlippageFactorSet(uint256 newMin, uint256 oldMin, address sender)
func (_ExpectedRate *ExpectedRateFilterer) FilterMinSlippageFactorSet(opts *bind.FilterOpts) (*ExpectedRateMinSlippageFactorSetIterator, error) {

	logs, sub, err := _ExpectedRate.contract.FilterLogs(opts, "MinSlippageFactorSet")
	if err != nil {
		return nil, err
	}
	return &ExpectedRateMinSlippageFactorSetIterator{contract: _ExpectedRate.contract, event: "MinSlippageFactorSet", logs: logs, sub: sub}, nil
}

// WatchMinSlippageFactorSet is a free log subscription operation binding the contract event 0x4357e20f1241d972328c5b3239d9ef4ac96f0f4fce8e10fd3bf9053690dad0ac.
//
// Solidity: event MinSlippageFactorSet(uint256 newMin, uint256 oldMin, address sender)
func (_ExpectedRate *ExpectedRateFilterer) WatchMinSlippageFactorSet(opts *bind.WatchOpts, sink chan<- *ExpectedRateMinSlippageFactorSet) (event.Subscription, error) {

	logs, sub, err := _ExpectedRate.contract.WatchLogs(opts, "MinSlippageFactorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpectedRateMinSlippageFactorSet)
				if err := _ExpectedRate.contract.UnpackLog(event, "MinSlippageFactorSet", log); err != nil {
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

// ParseMinSlippageFactorSet is a log parse operation binding the contract event 0x4357e20f1241d972328c5b3239d9ef4ac96f0f4fce8e10fd3bf9053690dad0ac.
//
// Solidity: event MinSlippageFactorSet(uint256 newMin, uint256 oldMin, address sender)
func (_ExpectedRate *ExpectedRateFilterer) ParseMinSlippageFactorSet(log types.Log) (*ExpectedRateMinSlippageFactorSet, error) {
	event := new(ExpectedRateMinSlippageFactorSet)
	if err := _ExpectedRate.contract.UnpackLog(event, "MinSlippageFactorSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExpectedRateOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the ExpectedRate contract.
type ExpectedRateOperatorAddedIterator struct {
	Event *ExpectedRateOperatorAdded // Event containing the contract specifics and raw log

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
func (it *ExpectedRateOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpectedRateOperatorAdded)
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
		it.Event = new(ExpectedRateOperatorAdded)
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
func (it *ExpectedRateOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpectedRateOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpectedRateOperatorAdded represents a OperatorAdded event raised by the ExpectedRate contract.
type ExpectedRateOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_ExpectedRate *ExpectedRateFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*ExpectedRateOperatorAddedIterator, error) {

	logs, sub, err := _ExpectedRate.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &ExpectedRateOperatorAddedIterator{contract: _ExpectedRate.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_ExpectedRate *ExpectedRateFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *ExpectedRateOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _ExpectedRate.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpectedRateOperatorAdded)
				if err := _ExpectedRate.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
func (_ExpectedRate *ExpectedRateFilterer) ParseOperatorAdded(log types.Log) (*ExpectedRateOperatorAdded, error) {
	event := new(ExpectedRateOperatorAdded)
	if err := _ExpectedRate.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExpectedRateQuantityFactorSetIterator is returned from FilterQuantityFactorSet and is used to iterate over the raw logs and unpacked data for QuantityFactorSet events raised by the ExpectedRate contract.
type ExpectedRateQuantityFactorSetIterator struct {
	Event *ExpectedRateQuantityFactorSet // Event containing the contract specifics and raw log

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
func (it *ExpectedRateQuantityFactorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpectedRateQuantityFactorSet)
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
		it.Event = new(ExpectedRateQuantityFactorSet)
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
func (it *ExpectedRateQuantityFactorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpectedRateQuantityFactorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpectedRateQuantityFactorSet represents a QuantityFactorSet event raised by the ExpectedRate contract.
type ExpectedRateQuantityFactorSet struct {
	NewFactor *big.Int
	OldFactor *big.Int
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuantityFactorSet is a free log retrieval operation binding the contract event 0xd0f6fc40d497232b5aab1b7a34ea00ea45886e52d2fed39ad62af798a870fae3.
//
// Solidity: event QuantityFactorSet(uint256 newFactor, uint256 oldFactor, address sender)
func (_ExpectedRate *ExpectedRateFilterer) FilterQuantityFactorSet(opts *bind.FilterOpts) (*ExpectedRateQuantityFactorSetIterator, error) {

	logs, sub, err := _ExpectedRate.contract.FilterLogs(opts, "QuantityFactorSet")
	if err != nil {
		return nil, err
	}
	return &ExpectedRateQuantityFactorSetIterator{contract: _ExpectedRate.contract, event: "QuantityFactorSet", logs: logs, sub: sub}, nil
}

// WatchQuantityFactorSet is a free log subscription operation binding the contract event 0xd0f6fc40d497232b5aab1b7a34ea00ea45886e52d2fed39ad62af798a870fae3.
//
// Solidity: event QuantityFactorSet(uint256 newFactor, uint256 oldFactor, address sender)
func (_ExpectedRate *ExpectedRateFilterer) WatchQuantityFactorSet(opts *bind.WatchOpts, sink chan<- *ExpectedRateQuantityFactorSet) (event.Subscription, error) {

	logs, sub, err := _ExpectedRate.contract.WatchLogs(opts, "QuantityFactorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpectedRateQuantityFactorSet)
				if err := _ExpectedRate.contract.UnpackLog(event, "QuantityFactorSet", log); err != nil {
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

// ParseQuantityFactorSet is a log parse operation binding the contract event 0xd0f6fc40d497232b5aab1b7a34ea00ea45886e52d2fed39ad62af798a870fae3.
//
// Solidity: event QuantityFactorSet(uint256 newFactor, uint256 oldFactor, address sender)
func (_ExpectedRate *ExpectedRateFilterer) ParseQuantityFactorSet(log types.Log) (*ExpectedRateQuantityFactorSet, error) {
	event := new(ExpectedRateQuantityFactorSet)
	if err := _ExpectedRate.contract.UnpackLog(event, "QuantityFactorSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExpectedRateTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the ExpectedRate contract.
type ExpectedRateTokenWithdrawIterator struct {
	Event *ExpectedRateTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *ExpectedRateTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpectedRateTokenWithdraw)
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
		it.Event = new(ExpectedRateTokenWithdraw)
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
func (it *ExpectedRateTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpectedRateTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpectedRateTokenWithdraw represents a TokenWithdraw event raised by the ExpectedRate contract.
type ExpectedRateTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_ExpectedRate *ExpectedRateFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*ExpectedRateTokenWithdrawIterator, error) {

	logs, sub, err := _ExpectedRate.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &ExpectedRateTokenWithdrawIterator{contract: _ExpectedRate.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_ExpectedRate *ExpectedRateFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *ExpectedRateTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _ExpectedRate.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpectedRateTokenWithdraw)
				if err := _ExpectedRate.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_ExpectedRate *ExpectedRateFilterer) ParseTokenWithdraw(log types.Log) (*ExpectedRateTokenWithdraw, error) {
	event := new(ExpectedRateTokenWithdraw)
	if err := _ExpectedRate.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExpectedRateTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the ExpectedRate contract.
type ExpectedRateTransferAdminPendingIterator struct {
	Event *ExpectedRateTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *ExpectedRateTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpectedRateTransferAdminPending)
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
		it.Event = new(ExpectedRateTransferAdminPending)
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
func (it *ExpectedRateTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpectedRateTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpectedRateTransferAdminPending represents a TransferAdminPending event raised by the ExpectedRate contract.
type ExpectedRateTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_ExpectedRate *ExpectedRateFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*ExpectedRateTransferAdminPendingIterator, error) {

	logs, sub, err := _ExpectedRate.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &ExpectedRateTransferAdminPendingIterator{contract: _ExpectedRate.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_ExpectedRate *ExpectedRateFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *ExpectedRateTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _ExpectedRate.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpectedRateTransferAdminPending)
				if err := _ExpectedRate.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_ExpectedRate *ExpectedRateFilterer) ParseTransferAdminPending(log types.Log) (*ExpectedRateTransferAdminPending, error) {
	event := new(ExpectedRateTransferAdminPending)
	if err := _ExpectedRate.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}
