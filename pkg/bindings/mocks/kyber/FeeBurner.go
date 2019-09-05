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

// FeeBurnerABI is the input ABI used to generate the binding from.
const FeeBurnerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserveKNCWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserveFeeToWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserveFeeToBurn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taxWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kncPerEthRatePrecision\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserveFeesInBps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taxFeeBps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletFeesInBps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"reserve\",\"type\":\"address\"},{\"name\":\"feesInBps\",\"type\":\"uint256\"},{\"name\":\"kncWallet\",\"type\":\"address\"}],\"name\":\"setReserveData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\"},{\"name\":\"feesInBps\",\"type\":\"uint256\"}],\"name\":\"setWalletFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setKNCRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_taxFeeBps\",\"type\":\"uint256\"}],\"name\":\"setTaxInBps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetwork\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"feePayedPerReserve\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\"},{\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"sendFeeToWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"knc\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_taxWallet\",\"type\":\"address\"}],\"name\":\"setTaxWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"burnReserveFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tradeWeiAmount\",\"type\":\"uint256\"},{\"name\":\"reserve\",\"type\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"handleFees\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"},{\"name\":\"_kncToken\",\"type\":\"address\"},{\"name\":\"_kyberNetwork\",\"type\":\"address\"},{\"name\":\"_initialKncToEthRatePrecision\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"feeInBps\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"kncWallet\",\"type\":\"address\"}],\"name\":\"ReserveDataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"feesInBps\",\"type\":\"uint256\"}],\"name\":\"WalletFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"feesInBps\",\"type\":\"uint256\"}],\"name\":\"TaxFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"taxWallet\",\"type\":\"address\"}],\"name\":\"TaxWalletSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ethToKncRatePrecision\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"kyberEthKnc\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"kyberKncEth\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"updater\",\"type\":\"address\"}],\"name\":\"KNCRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"walletFee\",\"type\":\"uint256\"}],\"name\":\"AssignFeeToWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"burnFee\",\"type\":\"uint256\"}],\"name\":\"AssignBurnFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"BurnAssignedFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taxWallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"SendTaxFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SendWalletFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"}]"

// FeeBurnerBin is the compiled bytecode used for deploying new contracts.
var FeeBurnerBin = "0x60606040526000600e55682086ac351052600000601155341561002157600080fd5b604051608080611db78339810160405280805191906020018051919060200180519190602001805160008054600160a060020a03191633600160a060020a039081169190911790915590925085161515905061007c57600080fd5b600160a060020a038316151561009157600080fd5b600160a060020a03821615156100a657600080fd5b8015156100b257600080fd5b60108054600160a060020a0319908116600160a060020a0394851617909155600080548216958416959095178555600f805490911693909216929092179055601155611cb390819061010490396000f3006060604052600436106101875763ffffffff60e060020a60003504166301a12fd3811461018c5780630a377f3a146101ad57806322dbf6d2146101e8578063267822471461021f57806327a099d8146102325780632b84fe83146102985780632dc0562d146102b757806330125416146102ca578063384c4d2f146102dd5780633b215823146102fc5780633ccdbb281461030f578063408ee7fe1461033857806345ab63b91461035757806346b8c49e1461037657806365dfc20f1461039f57806375829def146103c157806377f50f97146103e05780637acc8678146103f35780637c423f54146104125780637f3681f6146104255780639870d7fe146104385780639fad2dcb14610457578063ac8a584a1461046d578063b78b842d1461048c578063ce56c4541461049f578063d4fac45d146104c1578063dc93f7c9146104e6578063dd3ff4f614610505578063e61387e01461052a578063ea414b281461053d578063f6486cad1461055c578063f851a4401461057b578063fd062d3b1461058e575b600080fd5b341561019757600080fd5b6101ab600160a060020a03600435166105ca565b005b34156101b857600080fd5b6101cc600160a060020a036004351661073a565b604051600160a060020a03909116815260200160405180910390f35b34156101f357600080fd5b61020d600160a060020a0360043581169060243516610755565b60405190815260200160405180910390f35b341561022a57600080fd5b6101cc610772565b341561023d57600080fd5b610245610781565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561028457808201518382015260200161026c565b505050509050019250505060405180910390f35b34156102a357600080fd5b61020d600160a060020a03600435166107ea565b34156102c257600080fd5b6101cc6107fc565b34156102d557600080fd5b61020d61080b565b34156102e857600080fd5b61020d600160a060020a0360043516610811565b341561030757600080fd5b61020d610823565b341561031a57600080fd5b6101ab600160a060020a036004358116906024359060443516610829565b341561034357600080fd5b6101ab600160a060020a0360043516610920565b341561036257600080fd5b61020d600160a060020a0360043516610a1c565b341561038157600080fd5b6101ab600160a060020a036004358116906024359060443516610a2e565b34156103aa57600080fd5b6101ab600160a060020a0360043516602435610b10565b34156103cc57600080fd5b6101ab600160a060020a0360043516610b9f565b34156103eb57600080fd5b6101ab610c3a565b34156103fe57600080fd5b6101ab600160a060020a0360043516610cd4565b341561041d57600080fd5b610245610db6565b341561043057600080fd5b6101ab610e1c565b341561044357600080fd5b6101ab600160a060020a0360043516611038565b341561046257600080fd5b6101ab600435611108565b341561047857600080fd5b6101ab600160a060020a036004351661116c565b341561049757600080fd5b6101cc6112d8565b34156104aa57600080fd5b6101ab600435600160a060020a03602435166112e7565b34156104cc57600080fd5b61020d600160a060020a036004358116906024351661137a565b34156104f157600080fd5b61020d600160a060020a036004351661142c565b341561051057600080fd5b6101ab600160a060020a036004358116906024351661143e565b341561053557600080fd5b6101cc6115ae565b341561054857600080fd5b6101ab600160a060020a03600435166115bd565b341561056757600080fd5b6101ab600160a060020a0360043516611648565b341561058657600080fd5b6101cc6118f2565b341561059957600080fd5b6105b6600435600160a060020a0360243581169060443516611901565b604051901515815260200160405180910390f35b6000805433600160a060020a039081169116146105e657600080fd5b600160a060020a03821660009081526003602052604090205460ff16151561060d57600080fd5b50600160a060020a0381166000908152600360205260408120805460ff191690555b6005548110156107365781600160a060020a031660058281548110151561065257fe5b600091825260209091200154600160a060020a0316141561072e5760058054600019810190811061067f57fe5b60009182526020909120015460058054600160a060020a0390921691839081106106a557fe5b60009182526020909120018054600160a060020a031916600160a060020a039290921691909117905560058054906106e1906000198301611c2e565b507f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762826000604051600160a060020a039092168252151560208201526040908101905180910390a1610736565b60010161062f565b5050565b600860205260009081526040902054600160a060020a031681565b600c60209081526000928352604080842090915290825290205481565b600154600160a060020a031681565b610789611c57565b60048054806020026020016040519081016040528092919081815260200182805480156107df57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107c1575b505050505090505b90565b600a6020526000908152604090205481565b600d54600160a060020a031681565b60115481565b60076020526000908152604090205481565b600e5481565b60005433600160a060020a0390811691161461084457600080fd5b82600160a060020a031663a9059cbb828460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156108a157600080fd5b6102c65a03f115156108b257600080fd5b5050506040518051905015156108c757600080fd5b7f72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6838383604051600160a060020a03938416815260208101929092529091166040808301919091526060909101905180910390a1505050565b60005433600160a060020a0390811691161461093b57600080fd5b600160a060020a03811660009081526003602052604090205460ff161561096157600080fd5b6005546032901061097157600080fd5b7f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600360205260409020805460ff1916600190811790915560058054909181016109f08382611c2e565b5060009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055565b60096020526000908152604090205481565b600160a060020a03331660009081526002602052604090205460ff161515610a5557600080fd5b60648210610a6257600080fd5b600160a060020a0381161515610a7757600080fd5b600160a060020a0383811660009081526007602090815260408083208690556008909152908190208054600160a060020a031916928416929092179091557f999efec8241b9b7a1d9c2d2e207cde178cb3a02ca6a94d070eecb369674ead6f9084908490849051600160a060020a03938416815260208101929092529091166040808301919091526060909101905180910390a1505050565b60005433600160a060020a03908116911614610b2b57600080fd5b6127108110610b3957600080fd5b600160a060020a038216600090815260096020526040908190208290557f19f0c31fd2313f709ad6b9f15595720ff5765b72b394025288ac4f355fee0952908390839051600160a060020a03909216825260208201526040908101905180910390a15050565b60005433600160a060020a03908116911614610bba57600080fd5b600160a060020a0381161515610bcf57600080fd5b6001547f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4090600160a060020a0316604051600160a060020a03909116815260200160405180910390a160018054600160a060020a031916600160a060020a0392909216919091179055565b60015433600160a060020a03908116911614610c5557600080fd5b6001546000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed91600160a060020a039081169116604051600160a060020a039283168152911660208201526040908101905180910390a16001805460008054600160a060020a0319908116600160a060020a03841617909155169055565b60005433600160a060020a03908116911614610cef57600080fd5b600160a060020a0381161515610d0457600080fd5b7f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4081604051600160a060020a03909116815260200160405180910390a16000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed908290600160a060020a0316604051600160a060020a039283168152911660208201526040908101905180910390a160008054600160a060020a031916600160a060020a0392909216919091179055565b610dbe611c57565b60058054806020026020016040519081016040528092919081815260200182805480156107df57602002820191906000526020600020908154600160a060020a031681526001909101906020018083116107c1575050505050905090565b601054600f546000918291600160a060020a039182169163809a9e559173eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee9116670de0b6b3a7640000856040516040015260405160e060020a63ffffffff8616028152600160a060020a03938416600482015291909216602482015260448101919091526064016040805180830381600087803b1515610eaf57600080fd5b6102c65a03f11515610ec057600080fd5b50505060405180519060200180515050601054600f54919350600160a060020a039081169163809a9e55911673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee670de0b6b3a764000060006040516040015260405160e060020a63ffffffff8616028152600160a060020a03938416600482015291909216602482015260448101919091526064016040805180830381600087803b1515610f6157600080fd5b6102c65a03f11515610f7257600080fd5b505050604051805190602001805150909150506f01812f9cf7920e2b66973e200000000082820210610fa357600080fd5b6e604be73de4838ad9a5cf880000000081830211610fc057600080fd5b69d3c21bcecceda1000000821115610fd757600080fd5b60118290557fe55ada78a782c5b59f55b44255857da4f2ed737a5a94b83e9275ee710d0d48c4828083336040519384526020840192909252604080840191909152600160a060020a0390911660608301526080909101905180910390a15050565b60005433600160a060020a0390811691161461105357600080fd5b600160a060020a03811660009081526002602052604090205460ff161561107957600080fd5b6004546032901061108957600080fd5b7f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600260205260409020805460ff1916600190811790915560048054909181016109f08382611c2e565b60005433600160a060020a0390811691161461112357600080fd5b612710811061113157600080fd5b600e8190557f560f2dab6b3f89e548b63a9eabb6e43ec0e70bb81bdc69e5dc578c72bab629f58160405190815260200160405180910390a150565b6000805433600160a060020a0390811691161461118857600080fd5b600160a060020a03821660009081526002602052604090205460ff1615156111af57600080fd5b50600160a060020a0381166000908152600260205260408120805460ff191690555b6004548110156107365781600160a060020a03166004828154811015156111f457fe5b600091825260209091200154600160a060020a031614156112d05760048054600019810190811061122157fe5b60009182526020909120015460048054600160a060020a03909216918390811061124757fe5b60009182526020909120018054600160a060020a031916600160a060020a03929092169190911790556004805460001901906112839082611c2e565b507f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b826000604051600160a060020a039092168252151560208201526040908101905180910390a1610736565b6001016111d1565b601054600160a060020a031681565b60005433600160a060020a0390811691161461130257600080fd5b600160a060020a03811682156108fc0283604051600060405180830381858888f19350505050151561133357600080fd5b7fec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de8282604051918252600160a060020a031660208201526040908101905180910390a15050565b6000600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14156113b25750600160a060020a03811631611426565b82600160a060020a03166370a082318360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561140957600080fd5b6102c65a03f1151561141a57600080fd5b50505060405180519150505b92915050565b600b6020526000908152604090205481565b600160a060020a038082166000908152600c60209081526040808320938616835292905220546001811161147157600080fd5b600160a060020a038083166000818152600c602090815260408083208886168452825280832060019055600f54938352600890915280822054928416936323b872dd931691879160001987019190516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561151357600080fd5b6102c65a03f1151561152457600080fd5b50505060405180519050151561153957600080fd5b600160a060020a038083166000908152600b60205260409081902080548401600019019055908416907fb3f3e7375c0c0c4f7dd94069a5a4e68667827491318da786c818b8c7a794924e908490339051600160a060020a039283168152911660208201526040908101905180910390a2505050565b600f54600160a060020a031681565b60005433600160a060020a039081169116146115d857600080fd5b600160a060020a03811615156115ed57600080fd5b600d8054600160a060020a031916600160a060020a0383161790557f847d0f7f2b16c8dd0b72c0606e65e8bf1b624633d37905b0e08145a295ab875881604051600160a060020a03909116815260200160405180910390a150565b600160a060020a0381166000908152600a6020526040812054906002821161166f57600080fd5b600160a060020a038084166000908152600a6020526040902060019055600d54161580159061169f5750600e5415155b156117e05750600e54612710600019830191820204908190116116c157600080fd5b808203915060008111156117e057600f54600160a060020a0384811660009081526008602052604080822054600d54948416946323b872dd9491821693911691869190516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561175857600080fd5b6102c65a03f1151561176957600080fd5b50505060405180519050151561177e57600080fd5b600d54600160a060020a03808516917f540d888e67a7f36992e365be9fddab5e2fd60e27b220d330c18f04650fd562e09133911684604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a25b600f54600160a060020a0384811660009081526008602052604080822054938316936379cc6790931691600019870191516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561185857600080fd5b6102c65a03f1151561186957600080fd5b50505060405180519050151561187e57600080fd5b600160a060020a0383166000818152600b6020526040908190208054848601016000199081019091557f2f8d2d194cbe1816411754a2fc9478a11f0707da481b11cff7c69791eb877ee191339186019051600160a060020a03909216825260208201526040908101905180910390a2505050565b600054600160a060020a031681565b6000808080806b204fce5e3e2502611000000088111561192057600080fd5b600f546011546119539173eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee91600160a060020a03909116908b90611ab0565b600160a060020a038816600090815260076020526040902054909450612710908502600160a060020a0388166000908152600960205260409020549190049350612710908402049150818310156119a957600080fd5b508082036000821115611a3657600160a060020a038088166000908152600c60209081526040808320938a1683529290528190208054840190557f366bc34352215bf0bd3b527cfd6718605e1f5938777e42bcd8ed92f578368f529088908890859051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b6000811115611aa2577ff838f6ddc89706878e3c3e698e9b5cbfbf2c0e3d3dcd0bd2e00f1ccf313e01858782604051600160a060020a03909216825260208201526040908101905180910390a1600160a060020a0387166000908152600a602052604090208054820190555b506001979650505050505050565b6000611ace83611abf87611ad9565b611ac887611ad9565b85611b9d565b90505b949350505050565b600080600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415611b0a5760129150611b97565b50600160a060020a038216600090815260066020526040902054801515611b935782600160a060020a031663313ce5676000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611b7157600080fd5b6102c65a03f11515611b8257600080fd5b505050604051805190509150611b97565b8091505b50919050565b60006b204fce5e3e25026110000000851115611bb857600080fd5b69d3c21bcecceda1000000821115611bcf57600080fd5b838310611c025760128484031115611be657600080fd5b670de0b6b3a7640000858302858503600a0a025b049050611ad1565b60128385031115611c1257600080fd5b828403600a0a670de0b6b3a764000002828602811515611bfa57fe5b815481835581811511611c5257600083815260209020611c52918101908301611c69565b505050565b60206040519081016040526000815290565b6107e791905b80821115611c835760008155600101611c6f565b50905600a165627a7a723058206c2c74a742e08ed019a4d40baac87de00277b4d9751bd53311d1a2f122c8c8bc0029"

// DeployFeeBurner deploys a new Ethereum contract, binding an instance of FeeBurner to it.
func DeployFeeBurner(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _kncToken common.Address, _kyberNetwork common.Address, _initialKncToEthRatePrecision *big.Int) (common.Address, *types.Transaction, *FeeBurner, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeBurnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeeBurnerBin), backend, _admin, _kncToken, _kyberNetwork, _initialKncToEthRatePrecision)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeBurner{FeeBurnerCaller: FeeBurnerCaller{contract: contract}, FeeBurnerTransactor: FeeBurnerTransactor{contract: contract}, FeeBurnerFilterer: FeeBurnerFilterer{contract: contract}}, nil
}

// FeeBurner is an auto generated Go binding around an Ethereum contract.
type FeeBurner struct {
	FeeBurnerCaller     // Read-only binding to the contract
	FeeBurnerTransactor // Write-only binding to the contract
	FeeBurnerFilterer   // Log filterer for contract events
}

// FeeBurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeBurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeBurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeBurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeBurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeBurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeBurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeBurnerSession struct {
	Contract     *FeeBurner        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeBurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeBurnerCallerSession struct {
	Contract *FeeBurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FeeBurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeBurnerTransactorSession struct {
	Contract     *FeeBurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FeeBurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeBurnerRaw struct {
	Contract *FeeBurner // Generic contract binding to access the raw methods on
}

// FeeBurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeBurnerCallerRaw struct {
	Contract *FeeBurnerCaller // Generic read-only contract binding to access the raw methods on
}

// FeeBurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeBurnerTransactorRaw struct {
	Contract *FeeBurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeBurner creates a new instance of FeeBurner, bound to a specific deployed contract.
func NewFeeBurner(address common.Address, backend bind.ContractBackend) (*FeeBurner, error) {
	contract, err := bindFeeBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeBurner{FeeBurnerCaller: FeeBurnerCaller{contract: contract}, FeeBurnerTransactor: FeeBurnerTransactor{contract: contract}, FeeBurnerFilterer: FeeBurnerFilterer{contract: contract}}, nil
}

// NewFeeBurnerCaller creates a new read-only instance of FeeBurner, bound to a specific deployed contract.
func NewFeeBurnerCaller(address common.Address, caller bind.ContractCaller) (*FeeBurnerCaller, error) {
	contract, err := bindFeeBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeBurnerCaller{contract: contract}, nil
}

// NewFeeBurnerTransactor creates a new write-only instance of FeeBurner, bound to a specific deployed contract.
func NewFeeBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeBurnerTransactor, error) {
	contract, err := bindFeeBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeBurnerTransactor{contract: contract}, nil
}

// NewFeeBurnerFilterer creates a new log filterer instance of FeeBurner, bound to a specific deployed contract.
func NewFeeBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeBurnerFilterer, error) {
	contract, err := bindFeeBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeBurnerFilterer{contract: contract}, nil
}

// bindFeeBurner binds a generic wrapper to an already deployed contract.
func bindFeeBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeBurnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeBurner *FeeBurnerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeBurner.Contract.FeeBurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeBurner *FeeBurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeBurner.Contract.FeeBurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeBurner *FeeBurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeBurner.Contract.FeeBurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeBurner *FeeBurnerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeBurner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeBurner *FeeBurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeBurner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeBurner *FeeBurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeBurner.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_FeeBurner *FeeBurnerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_FeeBurner *FeeBurnerSession) Admin() (common.Address, error) {
	return _FeeBurner.Contract.Admin(&_FeeBurner.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_FeeBurner *FeeBurnerCallerSession) Admin() (common.Address, error) {
	return _FeeBurner.Contract.Admin(&_FeeBurner.CallOpts)
}

// FeePayedPerReserve is a free data retrieval call binding the contract method 0xdc93f7c9.
//
// Solidity: function feePayedPerReserve(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCaller) FeePayedPerReserve(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "feePayedPerReserve", arg0)
	return *ret0, err
}

// FeePayedPerReserve is a free data retrieval call binding the contract method 0xdc93f7c9.
//
// Solidity: function feePayedPerReserve(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerSession) FeePayedPerReserve(arg0 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.FeePayedPerReserve(&_FeeBurner.CallOpts, arg0)
}

// FeePayedPerReserve is a free data retrieval call binding the contract method 0xdc93f7c9.
//
// Solidity: function feePayedPerReserve(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCallerSession) FeePayedPerReserve(arg0 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.FeePayedPerReserve(&_FeeBurner.CallOpts, arg0)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_FeeBurner *FeeBurnerCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_FeeBurner *FeeBurnerSession) GetAlerters() ([]common.Address, error) {
	return _FeeBurner.Contract.GetAlerters(&_FeeBurner.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_FeeBurner *FeeBurnerCallerSession) GetAlerters() ([]common.Address, error) {
	return _FeeBurner.Contract.GetAlerters(&_FeeBurner.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_FeeBurner *FeeBurnerCaller) GetBalance(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "getBalance", token, user)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_FeeBurner *FeeBurnerSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.GetBalance(&_FeeBurner.CallOpts, token, user)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_FeeBurner *FeeBurnerCallerSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.GetBalance(&_FeeBurner.CallOpts, token, user)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_FeeBurner *FeeBurnerCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_FeeBurner *FeeBurnerSession) GetOperators() ([]common.Address, error) {
	return _FeeBurner.Contract.GetOperators(&_FeeBurner.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_FeeBurner *FeeBurnerCallerSession) GetOperators() ([]common.Address, error) {
	return _FeeBurner.Contract.GetOperators(&_FeeBurner.CallOpts)
}

// Knc is a free data retrieval call binding the contract method 0xe61387e0.
//
// Solidity: function knc() constant returns(address)
func (_FeeBurner *FeeBurnerCaller) Knc(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "knc")
	return *ret0, err
}

// Knc is a free data retrieval call binding the contract method 0xe61387e0.
//
// Solidity: function knc() constant returns(address)
func (_FeeBurner *FeeBurnerSession) Knc() (common.Address, error) {
	return _FeeBurner.Contract.Knc(&_FeeBurner.CallOpts)
}

// Knc is a free data retrieval call binding the contract method 0xe61387e0.
//
// Solidity: function knc() constant returns(address)
func (_FeeBurner *FeeBurnerCallerSession) Knc() (common.Address, error) {
	return _FeeBurner.Contract.Knc(&_FeeBurner.CallOpts)
}

// KncPerEthRatePrecision is a free data retrieval call binding the contract method 0x30125416.
//
// Solidity: function kncPerEthRatePrecision() constant returns(uint256)
func (_FeeBurner *FeeBurnerCaller) KncPerEthRatePrecision(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "kncPerEthRatePrecision")
	return *ret0, err
}

// KncPerEthRatePrecision is a free data retrieval call binding the contract method 0x30125416.
//
// Solidity: function kncPerEthRatePrecision() constant returns(uint256)
func (_FeeBurner *FeeBurnerSession) KncPerEthRatePrecision() (*big.Int, error) {
	return _FeeBurner.Contract.KncPerEthRatePrecision(&_FeeBurner.CallOpts)
}

// KncPerEthRatePrecision is a free data retrieval call binding the contract method 0x30125416.
//
// Solidity: function kncPerEthRatePrecision() constant returns(uint256)
func (_FeeBurner *FeeBurnerCallerSession) KncPerEthRatePrecision() (*big.Int, error) {
	return _FeeBurner.Contract.KncPerEthRatePrecision(&_FeeBurner.CallOpts)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_FeeBurner *FeeBurnerCaller) KyberNetwork(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "kyberNetwork")
	return *ret0, err
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_FeeBurner *FeeBurnerSession) KyberNetwork() (common.Address, error) {
	return _FeeBurner.Contract.KyberNetwork(&_FeeBurner.CallOpts)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_FeeBurner *FeeBurnerCallerSession) KyberNetwork() (common.Address, error) {
	return _FeeBurner.Contract.KyberNetwork(&_FeeBurner.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_FeeBurner *FeeBurnerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_FeeBurner *FeeBurnerSession) PendingAdmin() (common.Address, error) {
	return _FeeBurner.Contract.PendingAdmin(&_FeeBurner.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_FeeBurner *FeeBurnerCallerSession) PendingAdmin() (common.Address, error) {
	return _FeeBurner.Contract.PendingAdmin(&_FeeBurner.CallOpts)
}

// ReserveFeeToBurn is a free data retrieval call binding the contract method 0x2b84fe83.
//
// Solidity: function reserveFeeToBurn(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCaller) ReserveFeeToBurn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "reserveFeeToBurn", arg0)
	return *ret0, err
}

// ReserveFeeToBurn is a free data retrieval call binding the contract method 0x2b84fe83.
//
// Solidity: function reserveFeeToBurn(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerSession) ReserveFeeToBurn(arg0 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.ReserveFeeToBurn(&_FeeBurner.CallOpts, arg0)
}

// ReserveFeeToBurn is a free data retrieval call binding the contract method 0x2b84fe83.
//
// Solidity: function reserveFeeToBurn(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCallerSession) ReserveFeeToBurn(arg0 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.ReserveFeeToBurn(&_FeeBurner.CallOpts, arg0)
}

// ReserveFeeToWallet is a free data retrieval call binding the contract method 0x22dbf6d2.
//
// Solidity: function reserveFeeToWallet(address , address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCaller) ReserveFeeToWallet(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "reserveFeeToWallet", arg0, arg1)
	return *ret0, err
}

// ReserveFeeToWallet is a free data retrieval call binding the contract method 0x22dbf6d2.
//
// Solidity: function reserveFeeToWallet(address , address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerSession) ReserveFeeToWallet(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.ReserveFeeToWallet(&_FeeBurner.CallOpts, arg0, arg1)
}

// ReserveFeeToWallet is a free data retrieval call binding the contract method 0x22dbf6d2.
//
// Solidity: function reserveFeeToWallet(address , address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCallerSession) ReserveFeeToWallet(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.ReserveFeeToWallet(&_FeeBurner.CallOpts, arg0, arg1)
}

// ReserveFeesInBps is a free data retrieval call binding the contract method 0x384c4d2f.
//
// Solidity: function reserveFeesInBps(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCaller) ReserveFeesInBps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "reserveFeesInBps", arg0)
	return *ret0, err
}

// ReserveFeesInBps is a free data retrieval call binding the contract method 0x384c4d2f.
//
// Solidity: function reserveFeesInBps(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerSession) ReserveFeesInBps(arg0 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.ReserveFeesInBps(&_FeeBurner.CallOpts, arg0)
}

// ReserveFeesInBps is a free data retrieval call binding the contract method 0x384c4d2f.
//
// Solidity: function reserveFeesInBps(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCallerSession) ReserveFeesInBps(arg0 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.ReserveFeesInBps(&_FeeBurner.CallOpts, arg0)
}

// ReserveKNCWallet is a free data retrieval call binding the contract method 0x0a377f3a.
//
// Solidity: function reserveKNCWallet(address ) constant returns(address)
func (_FeeBurner *FeeBurnerCaller) ReserveKNCWallet(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "reserveKNCWallet", arg0)
	return *ret0, err
}

// ReserveKNCWallet is a free data retrieval call binding the contract method 0x0a377f3a.
//
// Solidity: function reserveKNCWallet(address ) constant returns(address)
func (_FeeBurner *FeeBurnerSession) ReserveKNCWallet(arg0 common.Address) (common.Address, error) {
	return _FeeBurner.Contract.ReserveKNCWallet(&_FeeBurner.CallOpts, arg0)
}

// ReserveKNCWallet is a free data retrieval call binding the contract method 0x0a377f3a.
//
// Solidity: function reserveKNCWallet(address ) constant returns(address)
func (_FeeBurner *FeeBurnerCallerSession) ReserveKNCWallet(arg0 common.Address) (common.Address, error) {
	return _FeeBurner.Contract.ReserveKNCWallet(&_FeeBurner.CallOpts, arg0)
}

// TaxFeeBps is a free data retrieval call binding the contract method 0x3b215823.
//
// Solidity: function taxFeeBps() constant returns(uint256)
func (_FeeBurner *FeeBurnerCaller) TaxFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "taxFeeBps")
	return *ret0, err
}

// TaxFeeBps is a free data retrieval call binding the contract method 0x3b215823.
//
// Solidity: function taxFeeBps() constant returns(uint256)
func (_FeeBurner *FeeBurnerSession) TaxFeeBps() (*big.Int, error) {
	return _FeeBurner.Contract.TaxFeeBps(&_FeeBurner.CallOpts)
}

// TaxFeeBps is a free data retrieval call binding the contract method 0x3b215823.
//
// Solidity: function taxFeeBps() constant returns(uint256)
func (_FeeBurner *FeeBurnerCallerSession) TaxFeeBps() (*big.Int, error) {
	return _FeeBurner.Contract.TaxFeeBps(&_FeeBurner.CallOpts)
}

// TaxWallet is a free data retrieval call binding the contract method 0x2dc0562d.
//
// Solidity: function taxWallet() constant returns(address)
func (_FeeBurner *FeeBurnerCaller) TaxWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "taxWallet")
	return *ret0, err
}

// TaxWallet is a free data retrieval call binding the contract method 0x2dc0562d.
//
// Solidity: function taxWallet() constant returns(address)
func (_FeeBurner *FeeBurnerSession) TaxWallet() (common.Address, error) {
	return _FeeBurner.Contract.TaxWallet(&_FeeBurner.CallOpts)
}

// TaxWallet is a free data retrieval call binding the contract method 0x2dc0562d.
//
// Solidity: function taxWallet() constant returns(address)
func (_FeeBurner *FeeBurnerCallerSession) TaxWallet() (common.Address, error) {
	return _FeeBurner.Contract.TaxWallet(&_FeeBurner.CallOpts)
}

// WalletFeesInBps is a free data retrieval call binding the contract method 0x45ab63b9.
//
// Solidity: function walletFeesInBps(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCaller) WalletFeesInBps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FeeBurner.contract.Call(opts, out, "walletFeesInBps", arg0)
	return *ret0, err
}

// WalletFeesInBps is a free data retrieval call binding the contract method 0x45ab63b9.
//
// Solidity: function walletFeesInBps(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerSession) WalletFeesInBps(arg0 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.WalletFeesInBps(&_FeeBurner.CallOpts, arg0)
}

// WalletFeesInBps is a free data retrieval call binding the contract method 0x45ab63b9.
//
// Solidity: function walletFeesInBps(address ) constant returns(uint256)
func (_FeeBurner *FeeBurnerCallerSession) WalletFeesInBps(arg0 common.Address) (*big.Int, error) {
	return _FeeBurner.Contract.WalletFeesInBps(&_FeeBurner.CallOpts, arg0)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_FeeBurner *FeeBurnerTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_FeeBurner *FeeBurnerSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.AddAlerter(&_FeeBurner.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_FeeBurner *FeeBurnerTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.AddAlerter(&_FeeBurner.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_FeeBurner *FeeBurnerTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_FeeBurner *FeeBurnerSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.AddOperator(&_FeeBurner.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_FeeBurner *FeeBurnerTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.AddOperator(&_FeeBurner.TransactOpts, newOperator)
}

// BurnReserveFees is a paid mutator transaction binding the contract method 0xf6486cad.
//
// Solidity: function burnReserveFees(address reserve) returns()
func (_FeeBurner *FeeBurnerTransactor) BurnReserveFees(opts *bind.TransactOpts, reserve common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "burnReserveFees", reserve)
}

// BurnReserveFees is a paid mutator transaction binding the contract method 0xf6486cad.
//
// Solidity: function burnReserveFees(address reserve) returns()
func (_FeeBurner *FeeBurnerSession) BurnReserveFees(reserve common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.BurnReserveFees(&_FeeBurner.TransactOpts, reserve)
}

// BurnReserveFees is a paid mutator transaction binding the contract method 0xf6486cad.
//
// Solidity: function burnReserveFees(address reserve) returns()
func (_FeeBurner *FeeBurnerTransactorSession) BurnReserveFees(reserve common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.BurnReserveFees(&_FeeBurner.TransactOpts, reserve)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_FeeBurner *FeeBurnerTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_FeeBurner *FeeBurnerSession) ClaimAdmin() (*types.Transaction, error) {
	return _FeeBurner.Contract.ClaimAdmin(&_FeeBurner.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_FeeBurner *FeeBurnerTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _FeeBurner.Contract.ClaimAdmin(&_FeeBurner.TransactOpts)
}

// HandleFees is a paid mutator transaction binding the contract method 0xfd062d3b.
//
// Solidity: function handleFees(uint256 tradeWeiAmount, address reserve, address wallet) returns(bool)
func (_FeeBurner *FeeBurnerTransactor) HandleFees(opts *bind.TransactOpts, tradeWeiAmount *big.Int, reserve common.Address, wallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "handleFees", tradeWeiAmount, reserve, wallet)
}

// HandleFees is a paid mutator transaction binding the contract method 0xfd062d3b.
//
// Solidity: function handleFees(uint256 tradeWeiAmount, address reserve, address wallet) returns(bool)
func (_FeeBurner *FeeBurnerSession) HandleFees(tradeWeiAmount *big.Int, reserve common.Address, wallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.HandleFees(&_FeeBurner.TransactOpts, tradeWeiAmount, reserve, wallet)
}

// HandleFees is a paid mutator transaction binding the contract method 0xfd062d3b.
//
// Solidity: function handleFees(uint256 tradeWeiAmount, address reserve, address wallet) returns(bool)
func (_FeeBurner *FeeBurnerTransactorSession) HandleFees(tradeWeiAmount *big.Int, reserve common.Address, wallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.HandleFees(&_FeeBurner.TransactOpts, tradeWeiAmount, reserve, wallet)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_FeeBurner *FeeBurnerTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_FeeBurner *FeeBurnerSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.RemoveAlerter(&_FeeBurner.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_FeeBurner *FeeBurnerTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.RemoveAlerter(&_FeeBurner.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_FeeBurner *FeeBurnerTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_FeeBurner *FeeBurnerSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.RemoveOperator(&_FeeBurner.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_FeeBurner *FeeBurnerTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.RemoveOperator(&_FeeBurner.TransactOpts, operator)
}

// SendFeeToWallet is a paid mutator transaction binding the contract method 0xdd3ff4f6.
//
// Solidity: function sendFeeToWallet(address wallet, address reserve) returns()
func (_FeeBurner *FeeBurnerTransactor) SendFeeToWallet(opts *bind.TransactOpts, wallet common.Address, reserve common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "sendFeeToWallet", wallet, reserve)
}

// SendFeeToWallet is a paid mutator transaction binding the contract method 0xdd3ff4f6.
//
// Solidity: function sendFeeToWallet(address wallet, address reserve) returns()
func (_FeeBurner *FeeBurnerSession) SendFeeToWallet(wallet common.Address, reserve common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.SendFeeToWallet(&_FeeBurner.TransactOpts, wallet, reserve)
}

// SendFeeToWallet is a paid mutator transaction binding the contract method 0xdd3ff4f6.
//
// Solidity: function sendFeeToWallet(address wallet, address reserve) returns()
func (_FeeBurner *FeeBurnerTransactorSession) SendFeeToWallet(wallet common.Address, reserve common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.SendFeeToWallet(&_FeeBurner.TransactOpts, wallet, reserve)
}

// SetKNCRate is a paid mutator transaction binding the contract method 0x7f3681f6.
//
// Solidity: function setKNCRate() returns()
func (_FeeBurner *FeeBurnerTransactor) SetKNCRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "setKNCRate")
}

// SetKNCRate is a paid mutator transaction binding the contract method 0x7f3681f6.
//
// Solidity: function setKNCRate() returns()
func (_FeeBurner *FeeBurnerSession) SetKNCRate() (*types.Transaction, error) {
	return _FeeBurner.Contract.SetKNCRate(&_FeeBurner.TransactOpts)
}

// SetKNCRate is a paid mutator transaction binding the contract method 0x7f3681f6.
//
// Solidity: function setKNCRate() returns()
func (_FeeBurner *FeeBurnerTransactorSession) SetKNCRate() (*types.Transaction, error) {
	return _FeeBurner.Contract.SetKNCRate(&_FeeBurner.TransactOpts)
}

// SetReserveData is a paid mutator transaction binding the contract method 0x46b8c49e.
//
// Solidity: function setReserveData(address reserve, uint256 feesInBps, address kncWallet) returns()
func (_FeeBurner *FeeBurnerTransactor) SetReserveData(opts *bind.TransactOpts, reserve common.Address, feesInBps *big.Int, kncWallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "setReserveData", reserve, feesInBps, kncWallet)
}

// SetReserveData is a paid mutator transaction binding the contract method 0x46b8c49e.
//
// Solidity: function setReserveData(address reserve, uint256 feesInBps, address kncWallet) returns()
func (_FeeBurner *FeeBurnerSession) SetReserveData(reserve common.Address, feesInBps *big.Int, kncWallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.SetReserveData(&_FeeBurner.TransactOpts, reserve, feesInBps, kncWallet)
}

// SetReserveData is a paid mutator transaction binding the contract method 0x46b8c49e.
//
// Solidity: function setReserveData(address reserve, uint256 feesInBps, address kncWallet) returns()
func (_FeeBurner *FeeBurnerTransactorSession) SetReserveData(reserve common.Address, feesInBps *big.Int, kncWallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.SetReserveData(&_FeeBurner.TransactOpts, reserve, feesInBps, kncWallet)
}

// SetTaxInBps is a paid mutator transaction binding the contract method 0x9fad2dcb.
//
// Solidity: function setTaxInBps(uint256 _taxFeeBps) returns()
func (_FeeBurner *FeeBurnerTransactor) SetTaxInBps(opts *bind.TransactOpts, _taxFeeBps *big.Int) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "setTaxInBps", _taxFeeBps)
}

// SetTaxInBps is a paid mutator transaction binding the contract method 0x9fad2dcb.
//
// Solidity: function setTaxInBps(uint256 _taxFeeBps) returns()
func (_FeeBurner *FeeBurnerSession) SetTaxInBps(_taxFeeBps *big.Int) (*types.Transaction, error) {
	return _FeeBurner.Contract.SetTaxInBps(&_FeeBurner.TransactOpts, _taxFeeBps)
}

// SetTaxInBps is a paid mutator transaction binding the contract method 0x9fad2dcb.
//
// Solidity: function setTaxInBps(uint256 _taxFeeBps) returns()
func (_FeeBurner *FeeBurnerTransactorSession) SetTaxInBps(_taxFeeBps *big.Int) (*types.Transaction, error) {
	return _FeeBurner.Contract.SetTaxInBps(&_FeeBurner.TransactOpts, _taxFeeBps)
}

// SetTaxWallet is a paid mutator transaction binding the contract method 0xea414b28.
//
// Solidity: function setTaxWallet(address _taxWallet) returns()
func (_FeeBurner *FeeBurnerTransactor) SetTaxWallet(opts *bind.TransactOpts, _taxWallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "setTaxWallet", _taxWallet)
}

// SetTaxWallet is a paid mutator transaction binding the contract method 0xea414b28.
//
// Solidity: function setTaxWallet(address _taxWallet) returns()
func (_FeeBurner *FeeBurnerSession) SetTaxWallet(_taxWallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.SetTaxWallet(&_FeeBurner.TransactOpts, _taxWallet)
}

// SetTaxWallet is a paid mutator transaction binding the contract method 0xea414b28.
//
// Solidity: function setTaxWallet(address _taxWallet) returns()
func (_FeeBurner *FeeBurnerTransactorSession) SetTaxWallet(_taxWallet common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.SetTaxWallet(&_FeeBurner.TransactOpts, _taxWallet)
}

// SetWalletFees is a paid mutator transaction binding the contract method 0x65dfc20f.
//
// Solidity: function setWalletFees(address wallet, uint256 feesInBps) returns()
func (_FeeBurner *FeeBurnerTransactor) SetWalletFees(opts *bind.TransactOpts, wallet common.Address, feesInBps *big.Int) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "setWalletFees", wallet, feesInBps)
}

// SetWalletFees is a paid mutator transaction binding the contract method 0x65dfc20f.
//
// Solidity: function setWalletFees(address wallet, uint256 feesInBps) returns()
func (_FeeBurner *FeeBurnerSession) SetWalletFees(wallet common.Address, feesInBps *big.Int) (*types.Transaction, error) {
	return _FeeBurner.Contract.SetWalletFees(&_FeeBurner.TransactOpts, wallet, feesInBps)
}

// SetWalletFees is a paid mutator transaction binding the contract method 0x65dfc20f.
//
// Solidity: function setWalletFees(address wallet, uint256 feesInBps) returns()
func (_FeeBurner *FeeBurnerTransactorSession) SetWalletFees(wallet common.Address, feesInBps *big.Int) (*types.Transaction, error) {
	return _FeeBurner.Contract.SetWalletFees(&_FeeBurner.TransactOpts, wallet, feesInBps)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_FeeBurner *FeeBurnerTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_FeeBurner *FeeBurnerSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.TransferAdmin(&_FeeBurner.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_FeeBurner *FeeBurnerTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.TransferAdmin(&_FeeBurner.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_FeeBurner *FeeBurnerTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_FeeBurner *FeeBurnerSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.TransferAdminQuickly(&_FeeBurner.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_FeeBurner *FeeBurnerTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.TransferAdminQuickly(&_FeeBurner.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_FeeBurner *FeeBurnerTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_FeeBurner *FeeBurnerSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.WithdrawEther(&_FeeBurner.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_FeeBurner *FeeBurnerTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.WithdrawEther(&_FeeBurner.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_FeeBurner *FeeBurnerTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FeeBurner.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_FeeBurner *FeeBurnerSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.WithdrawToken(&_FeeBurner.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_FeeBurner *FeeBurnerTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _FeeBurner.Contract.WithdrawToken(&_FeeBurner.TransactOpts, token, amount, sendTo)
}

// FeeBurnerAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the FeeBurner contract.
type FeeBurnerAdminClaimedIterator struct {
	Event *FeeBurnerAdminClaimed // Event containing the contract specifics and raw log

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
func (it *FeeBurnerAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerAdminClaimed)
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
		it.Event = new(FeeBurnerAdminClaimed)
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
func (it *FeeBurnerAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerAdminClaimed represents a AdminClaimed event raised by the FeeBurner contract.
type FeeBurnerAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_FeeBurner *FeeBurnerFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*FeeBurnerAdminClaimedIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerAdminClaimedIterator{contract: _FeeBurner.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_FeeBurner *FeeBurnerFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *FeeBurnerAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerAdminClaimed)
				if err := _FeeBurner.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_FeeBurner *FeeBurnerFilterer) ParseAdminClaimed(log types.Log) (*FeeBurnerAdminClaimed, error) {
	event := new(FeeBurnerAdminClaimed)
	if err := _FeeBurner.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the FeeBurner contract.
type FeeBurnerAlerterAddedIterator struct {
	Event *FeeBurnerAlerterAdded // Event containing the contract specifics and raw log

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
func (it *FeeBurnerAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerAlerterAdded)
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
		it.Event = new(FeeBurnerAlerterAdded)
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
func (it *FeeBurnerAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerAlerterAdded represents a AlerterAdded event raised by the FeeBurner contract.
type FeeBurnerAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_FeeBurner *FeeBurnerFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*FeeBurnerAlerterAddedIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerAlerterAddedIterator{contract: _FeeBurner.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_FeeBurner *FeeBurnerFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *FeeBurnerAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerAlerterAdded)
				if err := _FeeBurner.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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
func (_FeeBurner *FeeBurnerFilterer) ParseAlerterAdded(log types.Log) (*FeeBurnerAlerterAdded, error) {
	event := new(FeeBurnerAlerterAdded)
	if err := _FeeBurner.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerAssignBurnFeesIterator is returned from FilterAssignBurnFees and is used to iterate over the raw logs and unpacked data for AssignBurnFees events raised by the FeeBurner contract.
type FeeBurnerAssignBurnFeesIterator struct {
	Event *FeeBurnerAssignBurnFees // Event containing the contract specifics and raw log

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
func (it *FeeBurnerAssignBurnFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerAssignBurnFees)
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
		it.Event = new(FeeBurnerAssignBurnFees)
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
func (it *FeeBurnerAssignBurnFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerAssignBurnFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerAssignBurnFees represents a AssignBurnFees event raised by the FeeBurner contract.
type FeeBurnerAssignBurnFees struct {
	Reserve common.Address
	BurnFee *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssignBurnFees is a free log retrieval operation binding the contract event 0xf838f6ddc89706878e3c3e698e9b5cbfbf2c0e3d3dcd0bd2e00f1ccf313e0185.
//
// Solidity: event AssignBurnFees(address reserve, uint256 burnFee)
func (_FeeBurner *FeeBurnerFilterer) FilterAssignBurnFees(opts *bind.FilterOpts) (*FeeBurnerAssignBurnFeesIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "AssignBurnFees")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerAssignBurnFeesIterator{contract: _FeeBurner.contract, event: "AssignBurnFees", logs: logs, sub: sub}, nil
}

// WatchAssignBurnFees is a free log subscription operation binding the contract event 0xf838f6ddc89706878e3c3e698e9b5cbfbf2c0e3d3dcd0bd2e00f1ccf313e0185.
//
// Solidity: event AssignBurnFees(address reserve, uint256 burnFee)
func (_FeeBurner *FeeBurnerFilterer) WatchAssignBurnFees(opts *bind.WatchOpts, sink chan<- *FeeBurnerAssignBurnFees) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "AssignBurnFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerAssignBurnFees)
				if err := _FeeBurner.contract.UnpackLog(event, "AssignBurnFees", log); err != nil {
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

// ParseAssignBurnFees is a log parse operation binding the contract event 0xf838f6ddc89706878e3c3e698e9b5cbfbf2c0e3d3dcd0bd2e00f1ccf313e0185.
//
// Solidity: event AssignBurnFees(address reserve, uint256 burnFee)
func (_FeeBurner *FeeBurnerFilterer) ParseAssignBurnFees(log types.Log) (*FeeBurnerAssignBurnFees, error) {
	event := new(FeeBurnerAssignBurnFees)
	if err := _FeeBurner.contract.UnpackLog(event, "AssignBurnFees", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerAssignFeeToWalletIterator is returned from FilterAssignFeeToWallet and is used to iterate over the raw logs and unpacked data for AssignFeeToWallet events raised by the FeeBurner contract.
type FeeBurnerAssignFeeToWalletIterator struct {
	Event *FeeBurnerAssignFeeToWallet // Event containing the contract specifics and raw log

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
func (it *FeeBurnerAssignFeeToWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerAssignFeeToWallet)
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
		it.Event = new(FeeBurnerAssignFeeToWallet)
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
func (it *FeeBurnerAssignFeeToWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerAssignFeeToWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerAssignFeeToWallet represents a AssignFeeToWallet event raised by the FeeBurner contract.
type FeeBurnerAssignFeeToWallet struct {
	Reserve   common.Address
	Wallet    common.Address
	WalletFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssignFeeToWallet is a free log retrieval operation binding the contract event 0x366bc34352215bf0bd3b527cfd6718605e1f5938777e42bcd8ed92f578368f52.
//
// Solidity: event AssignFeeToWallet(address reserve, address wallet, uint256 walletFee)
func (_FeeBurner *FeeBurnerFilterer) FilterAssignFeeToWallet(opts *bind.FilterOpts) (*FeeBurnerAssignFeeToWalletIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "AssignFeeToWallet")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerAssignFeeToWalletIterator{contract: _FeeBurner.contract, event: "AssignFeeToWallet", logs: logs, sub: sub}, nil
}

// WatchAssignFeeToWallet is a free log subscription operation binding the contract event 0x366bc34352215bf0bd3b527cfd6718605e1f5938777e42bcd8ed92f578368f52.
//
// Solidity: event AssignFeeToWallet(address reserve, address wallet, uint256 walletFee)
func (_FeeBurner *FeeBurnerFilterer) WatchAssignFeeToWallet(opts *bind.WatchOpts, sink chan<- *FeeBurnerAssignFeeToWallet) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "AssignFeeToWallet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerAssignFeeToWallet)
				if err := _FeeBurner.contract.UnpackLog(event, "AssignFeeToWallet", log); err != nil {
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

// ParseAssignFeeToWallet is a log parse operation binding the contract event 0x366bc34352215bf0bd3b527cfd6718605e1f5938777e42bcd8ed92f578368f52.
//
// Solidity: event AssignFeeToWallet(address reserve, address wallet, uint256 walletFee)
func (_FeeBurner *FeeBurnerFilterer) ParseAssignFeeToWallet(log types.Log) (*FeeBurnerAssignFeeToWallet, error) {
	event := new(FeeBurnerAssignFeeToWallet)
	if err := _FeeBurner.contract.UnpackLog(event, "AssignFeeToWallet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerBurnAssignedFeesIterator is returned from FilterBurnAssignedFees and is used to iterate over the raw logs and unpacked data for BurnAssignedFees events raised by the FeeBurner contract.
type FeeBurnerBurnAssignedFeesIterator struct {
	Event *FeeBurnerBurnAssignedFees // Event containing the contract specifics and raw log

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
func (it *FeeBurnerBurnAssignedFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerBurnAssignedFees)
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
		it.Event = new(FeeBurnerBurnAssignedFees)
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
func (it *FeeBurnerBurnAssignedFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerBurnAssignedFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerBurnAssignedFees represents a BurnAssignedFees event raised by the FeeBurner contract.
type FeeBurnerBurnAssignedFees struct {
	Reserve  common.Address
	Sender   common.Address
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnAssignedFees is a free log retrieval operation binding the contract event 0x2f8d2d194cbe1816411754a2fc9478a11f0707da481b11cff7c69791eb877ee1.
//
// Solidity: event BurnAssignedFees(address indexed reserve, address sender, uint256 quantity)
func (_FeeBurner *FeeBurnerFilterer) FilterBurnAssignedFees(opts *bind.FilterOpts, reserve []common.Address) (*FeeBurnerBurnAssignedFeesIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "BurnAssignedFees", reserveRule)
	if err != nil {
		return nil, err
	}
	return &FeeBurnerBurnAssignedFeesIterator{contract: _FeeBurner.contract, event: "BurnAssignedFees", logs: logs, sub: sub}, nil
}

// WatchBurnAssignedFees is a free log subscription operation binding the contract event 0x2f8d2d194cbe1816411754a2fc9478a11f0707da481b11cff7c69791eb877ee1.
//
// Solidity: event BurnAssignedFees(address indexed reserve, address sender, uint256 quantity)
func (_FeeBurner *FeeBurnerFilterer) WatchBurnAssignedFees(opts *bind.WatchOpts, sink chan<- *FeeBurnerBurnAssignedFees, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "BurnAssignedFees", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerBurnAssignedFees)
				if err := _FeeBurner.contract.UnpackLog(event, "BurnAssignedFees", log); err != nil {
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

// ParseBurnAssignedFees is a log parse operation binding the contract event 0x2f8d2d194cbe1816411754a2fc9478a11f0707da481b11cff7c69791eb877ee1.
//
// Solidity: event BurnAssignedFees(address indexed reserve, address sender, uint256 quantity)
func (_FeeBurner *FeeBurnerFilterer) ParseBurnAssignedFees(log types.Log) (*FeeBurnerBurnAssignedFees, error) {
	event := new(FeeBurnerBurnAssignedFees)
	if err := _FeeBurner.contract.UnpackLog(event, "BurnAssignedFees", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the FeeBurner contract.
type FeeBurnerEtherWithdrawIterator struct {
	Event *FeeBurnerEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *FeeBurnerEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerEtherWithdraw)
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
		it.Event = new(FeeBurnerEtherWithdraw)
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
func (it *FeeBurnerEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerEtherWithdraw represents a EtherWithdraw event raised by the FeeBurner contract.
type FeeBurnerEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_FeeBurner *FeeBurnerFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*FeeBurnerEtherWithdrawIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerEtherWithdrawIterator{contract: _FeeBurner.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_FeeBurner *FeeBurnerFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *FeeBurnerEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerEtherWithdraw)
				if err := _FeeBurner.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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
func (_FeeBurner *FeeBurnerFilterer) ParseEtherWithdraw(log types.Log) (*FeeBurnerEtherWithdraw, error) {
	event := new(FeeBurnerEtherWithdraw)
	if err := _FeeBurner.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerKNCRateSetIterator is returned from FilterKNCRateSet and is used to iterate over the raw logs and unpacked data for KNCRateSet events raised by the FeeBurner contract.
type FeeBurnerKNCRateSetIterator struct {
	Event *FeeBurnerKNCRateSet // Event containing the contract specifics and raw log

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
func (it *FeeBurnerKNCRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerKNCRateSet)
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
		it.Event = new(FeeBurnerKNCRateSet)
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
func (it *FeeBurnerKNCRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerKNCRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerKNCRateSet represents a KNCRateSet event raised by the FeeBurner contract.
type FeeBurnerKNCRateSet struct {
	EthToKncRatePrecision *big.Int
	KyberEthKnc           *big.Int
	KyberKncEth           *big.Int
	Updater               common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterKNCRateSet is a free log retrieval operation binding the contract event 0xe55ada78a782c5b59f55b44255857da4f2ed737a5a94b83e9275ee710d0d48c4.
//
// Solidity: event KNCRateSet(uint256 ethToKncRatePrecision, uint256 kyberEthKnc, uint256 kyberKncEth, address updater)
func (_FeeBurner *FeeBurnerFilterer) FilterKNCRateSet(opts *bind.FilterOpts) (*FeeBurnerKNCRateSetIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "KNCRateSet")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerKNCRateSetIterator{contract: _FeeBurner.contract, event: "KNCRateSet", logs: logs, sub: sub}, nil
}

// WatchKNCRateSet is a free log subscription operation binding the contract event 0xe55ada78a782c5b59f55b44255857da4f2ed737a5a94b83e9275ee710d0d48c4.
//
// Solidity: event KNCRateSet(uint256 ethToKncRatePrecision, uint256 kyberEthKnc, uint256 kyberKncEth, address updater)
func (_FeeBurner *FeeBurnerFilterer) WatchKNCRateSet(opts *bind.WatchOpts, sink chan<- *FeeBurnerKNCRateSet) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "KNCRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerKNCRateSet)
				if err := _FeeBurner.contract.UnpackLog(event, "KNCRateSet", log); err != nil {
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

// ParseKNCRateSet is a log parse operation binding the contract event 0xe55ada78a782c5b59f55b44255857da4f2ed737a5a94b83e9275ee710d0d48c4.
//
// Solidity: event KNCRateSet(uint256 ethToKncRatePrecision, uint256 kyberEthKnc, uint256 kyberKncEth, address updater)
func (_FeeBurner *FeeBurnerFilterer) ParseKNCRateSet(log types.Log) (*FeeBurnerKNCRateSet, error) {
	event := new(FeeBurnerKNCRateSet)
	if err := _FeeBurner.contract.UnpackLog(event, "KNCRateSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the FeeBurner contract.
type FeeBurnerOperatorAddedIterator struct {
	Event *FeeBurnerOperatorAdded // Event containing the contract specifics and raw log

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
func (it *FeeBurnerOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerOperatorAdded)
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
		it.Event = new(FeeBurnerOperatorAdded)
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
func (it *FeeBurnerOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerOperatorAdded represents a OperatorAdded event raised by the FeeBurner contract.
type FeeBurnerOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_FeeBurner *FeeBurnerFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*FeeBurnerOperatorAddedIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerOperatorAddedIterator{contract: _FeeBurner.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_FeeBurner *FeeBurnerFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *FeeBurnerOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerOperatorAdded)
				if err := _FeeBurner.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
func (_FeeBurner *FeeBurnerFilterer) ParseOperatorAdded(log types.Log) (*FeeBurnerOperatorAdded, error) {
	event := new(FeeBurnerOperatorAdded)
	if err := _FeeBurner.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerReserveDataSetIterator is returned from FilterReserveDataSet and is used to iterate over the raw logs and unpacked data for ReserveDataSet events raised by the FeeBurner contract.
type FeeBurnerReserveDataSetIterator struct {
	Event *FeeBurnerReserveDataSet // Event containing the contract specifics and raw log

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
func (it *FeeBurnerReserveDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerReserveDataSet)
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
		it.Event = new(FeeBurnerReserveDataSet)
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
func (it *FeeBurnerReserveDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerReserveDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerReserveDataSet represents a ReserveDataSet event raised by the FeeBurner contract.
type FeeBurnerReserveDataSet struct {
	Reserve   common.Address
	FeeInBps  *big.Int
	KncWallet common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReserveDataSet is a free log retrieval operation binding the contract event 0x999efec8241b9b7a1d9c2d2e207cde178cb3a02ca6a94d070eecb369674ead6f.
//
// Solidity: event ReserveDataSet(address reserve, uint256 feeInBps, address kncWallet)
func (_FeeBurner *FeeBurnerFilterer) FilterReserveDataSet(opts *bind.FilterOpts) (*FeeBurnerReserveDataSetIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "ReserveDataSet")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerReserveDataSetIterator{contract: _FeeBurner.contract, event: "ReserveDataSet", logs: logs, sub: sub}, nil
}

// WatchReserveDataSet is a free log subscription operation binding the contract event 0x999efec8241b9b7a1d9c2d2e207cde178cb3a02ca6a94d070eecb369674ead6f.
//
// Solidity: event ReserveDataSet(address reserve, uint256 feeInBps, address kncWallet)
func (_FeeBurner *FeeBurnerFilterer) WatchReserveDataSet(opts *bind.WatchOpts, sink chan<- *FeeBurnerReserveDataSet) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "ReserveDataSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerReserveDataSet)
				if err := _FeeBurner.contract.UnpackLog(event, "ReserveDataSet", log); err != nil {
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

// ParseReserveDataSet is a log parse operation binding the contract event 0x999efec8241b9b7a1d9c2d2e207cde178cb3a02ca6a94d070eecb369674ead6f.
//
// Solidity: event ReserveDataSet(address reserve, uint256 feeInBps, address kncWallet)
func (_FeeBurner *FeeBurnerFilterer) ParseReserveDataSet(log types.Log) (*FeeBurnerReserveDataSet, error) {
	event := new(FeeBurnerReserveDataSet)
	if err := _FeeBurner.contract.UnpackLog(event, "ReserveDataSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerSendTaxFeeIterator is returned from FilterSendTaxFee and is used to iterate over the raw logs and unpacked data for SendTaxFee events raised by the FeeBurner contract.
type FeeBurnerSendTaxFeeIterator struct {
	Event *FeeBurnerSendTaxFee // Event containing the contract specifics and raw log

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
func (it *FeeBurnerSendTaxFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerSendTaxFee)
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
		it.Event = new(FeeBurnerSendTaxFee)
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
func (it *FeeBurnerSendTaxFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerSendTaxFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerSendTaxFee represents a SendTaxFee event raised by the FeeBurner contract.
type FeeBurnerSendTaxFee struct {
	Reserve   common.Address
	Sender    common.Address
	TaxWallet common.Address
	Quantity  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendTaxFee is a free log retrieval operation binding the contract event 0x540d888e67a7f36992e365be9fddab5e2fd60e27b220d330c18f04650fd562e0.
//
// Solidity: event SendTaxFee(address indexed reserve, address sender, address taxWallet, uint256 quantity)
func (_FeeBurner *FeeBurnerFilterer) FilterSendTaxFee(opts *bind.FilterOpts, reserve []common.Address) (*FeeBurnerSendTaxFeeIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "SendTaxFee", reserveRule)
	if err != nil {
		return nil, err
	}
	return &FeeBurnerSendTaxFeeIterator{contract: _FeeBurner.contract, event: "SendTaxFee", logs: logs, sub: sub}, nil
}

// WatchSendTaxFee is a free log subscription operation binding the contract event 0x540d888e67a7f36992e365be9fddab5e2fd60e27b220d330c18f04650fd562e0.
//
// Solidity: event SendTaxFee(address indexed reserve, address sender, address taxWallet, uint256 quantity)
func (_FeeBurner *FeeBurnerFilterer) WatchSendTaxFee(opts *bind.WatchOpts, sink chan<- *FeeBurnerSendTaxFee, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "SendTaxFee", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerSendTaxFee)
				if err := _FeeBurner.contract.UnpackLog(event, "SendTaxFee", log); err != nil {
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

// ParseSendTaxFee is a log parse operation binding the contract event 0x540d888e67a7f36992e365be9fddab5e2fd60e27b220d330c18f04650fd562e0.
//
// Solidity: event SendTaxFee(address indexed reserve, address sender, address taxWallet, uint256 quantity)
func (_FeeBurner *FeeBurnerFilterer) ParseSendTaxFee(log types.Log) (*FeeBurnerSendTaxFee, error) {
	event := new(FeeBurnerSendTaxFee)
	if err := _FeeBurner.contract.UnpackLog(event, "SendTaxFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerSendWalletFeesIterator is returned from FilterSendWalletFees and is used to iterate over the raw logs and unpacked data for SendWalletFees events raised by the FeeBurner contract.
type FeeBurnerSendWalletFeesIterator struct {
	Event *FeeBurnerSendWalletFees // Event containing the contract specifics and raw log

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
func (it *FeeBurnerSendWalletFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerSendWalletFees)
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
		it.Event = new(FeeBurnerSendWalletFees)
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
func (it *FeeBurnerSendWalletFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerSendWalletFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerSendWalletFees represents a SendWalletFees event raised by the FeeBurner contract.
type FeeBurnerSendWalletFees struct {
	Wallet  common.Address
	Reserve common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSendWalletFees is a free log retrieval operation binding the contract event 0xb3f3e7375c0c0c4f7dd94069a5a4e68667827491318da786c818b8c7a794924e.
//
// Solidity: event SendWalletFees(address indexed wallet, address reserve, address sender)
func (_FeeBurner *FeeBurnerFilterer) FilterSendWalletFees(opts *bind.FilterOpts, wallet []common.Address) (*FeeBurnerSendWalletFeesIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "SendWalletFees", walletRule)
	if err != nil {
		return nil, err
	}
	return &FeeBurnerSendWalletFeesIterator{contract: _FeeBurner.contract, event: "SendWalletFees", logs: logs, sub: sub}, nil
}

// WatchSendWalletFees is a free log subscription operation binding the contract event 0xb3f3e7375c0c0c4f7dd94069a5a4e68667827491318da786c818b8c7a794924e.
//
// Solidity: event SendWalletFees(address indexed wallet, address reserve, address sender)
func (_FeeBurner *FeeBurnerFilterer) WatchSendWalletFees(opts *bind.WatchOpts, sink chan<- *FeeBurnerSendWalletFees, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "SendWalletFees", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerSendWalletFees)
				if err := _FeeBurner.contract.UnpackLog(event, "SendWalletFees", log); err != nil {
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

// ParseSendWalletFees is a log parse operation binding the contract event 0xb3f3e7375c0c0c4f7dd94069a5a4e68667827491318da786c818b8c7a794924e.
//
// Solidity: event SendWalletFees(address indexed wallet, address reserve, address sender)
func (_FeeBurner *FeeBurnerFilterer) ParseSendWalletFees(log types.Log) (*FeeBurnerSendWalletFees, error) {
	event := new(FeeBurnerSendWalletFees)
	if err := _FeeBurner.contract.UnpackLog(event, "SendWalletFees", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerTaxFeesSetIterator is returned from FilterTaxFeesSet and is used to iterate over the raw logs and unpacked data for TaxFeesSet events raised by the FeeBurner contract.
type FeeBurnerTaxFeesSetIterator struct {
	Event *FeeBurnerTaxFeesSet // Event containing the contract specifics and raw log

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
func (it *FeeBurnerTaxFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerTaxFeesSet)
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
		it.Event = new(FeeBurnerTaxFeesSet)
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
func (it *FeeBurnerTaxFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerTaxFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerTaxFeesSet represents a TaxFeesSet event raised by the FeeBurner contract.
type FeeBurnerTaxFeesSet struct {
	FeesInBps *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaxFeesSet is a free log retrieval operation binding the contract event 0x560f2dab6b3f89e548b63a9eabb6e43ec0e70bb81bdc69e5dc578c72bab629f5.
//
// Solidity: event TaxFeesSet(uint256 feesInBps)
func (_FeeBurner *FeeBurnerFilterer) FilterTaxFeesSet(opts *bind.FilterOpts) (*FeeBurnerTaxFeesSetIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "TaxFeesSet")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerTaxFeesSetIterator{contract: _FeeBurner.contract, event: "TaxFeesSet", logs: logs, sub: sub}, nil
}

// WatchTaxFeesSet is a free log subscription operation binding the contract event 0x560f2dab6b3f89e548b63a9eabb6e43ec0e70bb81bdc69e5dc578c72bab629f5.
//
// Solidity: event TaxFeesSet(uint256 feesInBps)
func (_FeeBurner *FeeBurnerFilterer) WatchTaxFeesSet(opts *bind.WatchOpts, sink chan<- *FeeBurnerTaxFeesSet) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "TaxFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerTaxFeesSet)
				if err := _FeeBurner.contract.UnpackLog(event, "TaxFeesSet", log); err != nil {
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

// ParseTaxFeesSet is a log parse operation binding the contract event 0x560f2dab6b3f89e548b63a9eabb6e43ec0e70bb81bdc69e5dc578c72bab629f5.
//
// Solidity: event TaxFeesSet(uint256 feesInBps)
func (_FeeBurner *FeeBurnerFilterer) ParseTaxFeesSet(log types.Log) (*FeeBurnerTaxFeesSet, error) {
	event := new(FeeBurnerTaxFeesSet)
	if err := _FeeBurner.contract.UnpackLog(event, "TaxFeesSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerTaxWalletSetIterator is returned from FilterTaxWalletSet and is used to iterate over the raw logs and unpacked data for TaxWalletSet events raised by the FeeBurner contract.
type FeeBurnerTaxWalletSetIterator struct {
	Event *FeeBurnerTaxWalletSet // Event containing the contract specifics and raw log

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
func (it *FeeBurnerTaxWalletSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerTaxWalletSet)
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
		it.Event = new(FeeBurnerTaxWalletSet)
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
func (it *FeeBurnerTaxWalletSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerTaxWalletSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerTaxWalletSet represents a TaxWalletSet event raised by the FeeBurner contract.
type FeeBurnerTaxWalletSet struct {
	TaxWallet common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaxWalletSet is a free log retrieval operation binding the contract event 0x847d0f7f2b16c8dd0b72c0606e65e8bf1b624633d37905b0e08145a295ab8758.
//
// Solidity: event TaxWalletSet(address taxWallet)
func (_FeeBurner *FeeBurnerFilterer) FilterTaxWalletSet(opts *bind.FilterOpts) (*FeeBurnerTaxWalletSetIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "TaxWalletSet")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerTaxWalletSetIterator{contract: _FeeBurner.contract, event: "TaxWalletSet", logs: logs, sub: sub}, nil
}

// WatchTaxWalletSet is a free log subscription operation binding the contract event 0x847d0f7f2b16c8dd0b72c0606e65e8bf1b624633d37905b0e08145a295ab8758.
//
// Solidity: event TaxWalletSet(address taxWallet)
func (_FeeBurner *FeeBurnerFilterer) WatchTaxWalletSet(opts *bind.WatchOpts, sink chan<- *FeeBurnerTaxWalletSet) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "TaxWalletSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerTaxWalletSet)
				if err := _FeeBurner.contract.UnpackLog(event, "TaxWalletSet", log); err != nil {
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

// ParseTaxWalletSet is a log parse operation binding the contract event 0x847d0f7f2b16c8dd0b72c0606e65e8bf1b624633d37905b0e08145a295ab8758.
//
// Solidity: event TaxWalletSet(address taxWallet)
func (_FeeBurner *FeeBurnerFilterer) ParseTaxWalletSet(log types.Log) (*FeeBurnerTaxWalletSet, error) {
	event := new(FeeBurnerTaxWalletSet)
	if err := _FeeBurner.contract.UnpackLog(event, "TaxWalletSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the FeeBurner contract.
type FeeBurnerTokenWithdrawIterator struct {
	Event *FeeBurnerTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *FeeBurnerTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerTokenWithdraw)
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
		it.Event = new(FeeBurnerTokenWithdraw)
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
func (it *FeeBurnerTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerTokenWithdraw represents a TokenWithdraw event raised by the FeeBurner contract.
type FeeBurnerTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_FeeBurner *FeeBurnerFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*FeeBurnerTokenWithdrawIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerTokenWithdrawIterator{contract: _FeeBurner.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_FeeBurner *FeeBurnerFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *FeeBurnerTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerTokenWithdraw)
				if err := _FeeBurner.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_FeeBurner *FeeBurnerFilterer) ParseTokenWithdraw(log types.Log) (*FeeBurnerTokenWithdraw, error) {
	event := new(FeeBurnerTokenWithdraw)
	if err := _FeeBurner.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the FeeBurner contract.
type FeeBurnerTransferAdminPendingIterator struct {
	Event *FeeBurnerTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *FeeBurnerTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerTransferAdminPending)
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
		it.Event = new(FeeBurnerTransferAdminPending)
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
func (it *FeeBurnerTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerTransferAdminPending represents a TransferAdminPending event raised by the FeeBurner contract.
type FeeBurnerTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_FeeBurner *FeeBurnerFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*FeeBurnerTransferAdminPendingIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerTransferAdminPendingIterator{contract: _FeeBurner.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_FeeBurner *FeeBurnerFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *FeeBurnerTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerTransferAdminPending)
				if err := _FeeBurner.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_FeeBurner *FeeBurnerFilterer) ParseTransferAdminPending(log types.Log) (*FeeBurnerTransferAdminPending, error) {
	event := new(FeeBurnerTransferAdminPending)
	if err := _FeeBurner.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FeeBurnerWalletFeesSetIterator is returned from FilterWalletFeesSet and is used to iterate over the raw logs and unpacked data for WalletFeesSet events raised by the FeeBurner contract.
type FeeBurnerWalletFeesSetIterator struct {
	Event *FeeBurnerWalletFeesSet // Event containing the contract specifics and raw log

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
func (it *FeeBurnerWalletFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeBurnerWalletFeesSet)
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
		it.Event = new(FeeBurnerWalletFeesSet)
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
func (it *FeeBurnerWalletFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeBurnerWalletFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeBurnerWalletFeesSet represents a WalletFeesSet event raised by the FeeBurner contract.
type FeeBurnerWalletFeesSet struct {
	Wallet    common.Address
	FeesInBps *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletFeesSet is a free log retrieval operation binding the contract event 0x19f0c31fd2313f709ad6b9f15595720ff5765b72b394025288ac4f355fee0952.
//
// Solidity: event WalletFeesSet(address wallet, uint256 feesInBps)
func (_FeeBurner *FeeBurnerFilterer) FilterWalletFeesSet(opts *bind.FilterOpts) (*FeeBurnerWalletFeesSetIterator, error) {

	logs, sub, err := _FeeBurner.contract.FilterLogs(opts, "WalletFeesSet")
	if err != nil {
		return nil, err
	}
	return &FeeBurnerWalletFeesSetIterator{contract: _FeeBurner.contract, event: "WalletFeesSet", logs: logs, sub: sub}, nil
}

// WatchWalletFeesSet is a free log subscription operation binding the contract event 0x19f0c31fd2313f709ad6b9f15595720ff5765b72b394025288ac4f355fee0952.
//
// Solidity: event WalletFeesSet(address wallet, uint256 feesInBps)
func (_FeeBurner *FeeBurnerFilterer) WatchWalletFeesSet(opts *bind.WatchOpts, sink chan<- *FeeBurnerWalletFeesSet) (event.Subscription, error) {

	logs, sub, err := _FeeBurner.contract.WatchLogs(opts, "WalletFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeBurnerWalletFeesSet)
				if err := _FeeBurner.contract.UnpackLog(event, "WalletFeesSet", log); err != nil {
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

// ParseWalletFeesSet is a log parse operation binding the contract event 0x19f0c31fd2313f709ad6b9f15595720ff5765b72b394025288ac4f355fee0952.
//
// Solidity: event WalletFeesSet(address wallet, uint256 feesInBps)
func (_FeeBurner *FeeBurnerFilterer) ParseWalletFeesSet(log types.Log) (*FeeBurnerWalletFeesSet, error) {
	event := new(FeeBurnerWalletFeesSet)
	if err := _FeeBurner.contract.UnpackLog(event, "WalletFeesSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
