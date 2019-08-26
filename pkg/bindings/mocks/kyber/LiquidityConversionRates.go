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

// LiquidityConversionRatesABI is the input ABI used to generate the binding from.
const LiquidityConversionRatesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxSellRateInPrecision\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"setReserveAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"abs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"qtyInwei\",\"type\":\"uint256\"}],\"name\":\"fromWeiToFp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"r\",\"type\":\"uint256\"},{\"name\":\"pMIn\",\"type\":\"uint256\"},{\"name\":\"e\",\"type\":\"uint256\"},{\"name\":\"precision\",\"type\":\"uint256\"}],\"name\":\"pE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"uint256\"},{\"name\":\"q\",\"type\":\"uint256\"},{\"name\":\"numPrecisionBits\",\"type\":\"uint256\"}],\"name\":\"ln\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxEthCapSellInFp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BIG_NUMBER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collectedFeesInTwei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rInFp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxEthCapBuyInFp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"formulaPrecision\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rInFp\",\"type\":\"uint256\"},{\"name\":\"_pMinInFp\",\"type\":\"uint256\"},{\"name\":\"_numFpBits\",\"type\":\"uint256\"},{\"name\":\"_maxCapBuyInWei\",\"type\":\"uint256\"},{\"name\":\"_maxCapSellInWei\",\"type\":\"uint256\"},{\"name\":\"_feeInBps\",\"type\":\"uint256\"},{\"name\":\"_maxTokenToEthRateInPrecision\",\"type\":\"uint256\"},{\"name\":\"_minTokenToEthRateInPrecision\",\"type\":\"uint256\"}],\"name\":\"setLiquidityParams\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"checkMultOverflow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"eInFp\",\"type\":\"uint256\"},{\"name\":\"deltaEInFp\",\"type\":\"uint256\"}],\"name\":\"buyRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"eInFp\",\"type\":\"uint256\"}],\"name\":\"sellRateZeroQuantity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"valueAfterReducingFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numFpBits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"r\",\"type\":\"uint256\"},{\"name\":\"pMIn\",\"type\":\"uint256\"},{\"name\":\"e\",\"type\":\"uint256\"},{\"name\":\"deltaE\",\"type\":\"uint256\"},{\"name\":\"precision\",\"type\":\"uint256\"}],\"name\":\"deltaTFunc\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"numPrecisionBits\",\"type\":\"uint256\"}],\"name\":\"log2ForSmallNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"uint256\"},{\"name\":\"q\",\"type\":\"uint256\"},{\"name\":\"precision\",\"type\":\"uint256\"}],\"name\":\"compactFraction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"conversionToken\",\"type\":\"address\"},{\"name\":\"buy\",\"type\":\"bool\"},{\"name\":\"qtyInSrcWei\",\"type\":\"uint256\"},{\"name\":\"eInFp\",\"type\":\"uint256\"}],\"name\":\"getRateWithE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"uint256\"},{\"name\":\"q\",\"type\":\"uint256\"}],\"name\":\"countLeadingZeros\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"eInFp\",\"type\":\"uint256\"},{\"name\":\"sellInputTokenQtyInFp\",\"type\":\"uint256\"},{\"name\":\"deltaTInFp\",\"type\":\"uint256\"}],\"name\":\"sellRate\",\"outputs\":[{\"name\":\"rateInPrecision\",\"type\":\"uint256\"},{\"name\":\"deltaEInFp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"qtyInTwei\",\"type\":\"uint256\"}],\"name\":\"fromTweiToFp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"r\",\"type\":\"uint256\"},{\"name\":\"pMIn\",\"type\":\"uint256\"},{\"name\":\"e\",\"type\":\"uint256\"},{\"name\":\"deltaT\",\"type\":\"uint256\"},{\"name\":\"precision\",\"type\":\"uint256\"},{\"name\":\"numPrecisionBits\",\"type\":\"uint256\"}],\"name\":\"deltaEFunc\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeInBps\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"uint256\"},{\"name\":\"q\",\"type\":\"uint256\"},{\"name\":\"numPrecisionBits\",\"type\":\"uint256\"}],\"name\":\"logBase2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rateInPrecision\",\"type\":\"uint256\"},{\"name\":\"buy\",\"type\":\"bool\"}],\"name\":\"rateAfterValidation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"calcCollectedFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"p\",\"type\":\"uint256\"},{\"name\":\"q\",\"type\":\"uint256\"},{\"name\":\"precision\",\"type\":\"uint256\"}],\"name\":\"exp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetCollectedFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"conversionToken\",\"type\":\"address\"},{\"name\":\"currentBlockNumber\",\"type\":\"uint256\"},{\"name\":\"buy\",\"type\":\"bool\"},{\"name\":\"qtyInSrcWei\",\"type\":\"uint256\"}],\"name\":\"getRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"conversionToken\",\"type\":\"address\"},{\"name\":\"buyAmountInTwei\",\"type\":\"int256\"},{\"name\":\"rateUpdateBlock\",\"type\":\"uint256\"},{\"name\":\"currentBlock\",\"type\":\"uint256\"}],\"name\":\"recordImbalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"eInFp\",\"type\":\"uint256\"}],\"name\":\"buyRateZeroQuantity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxBuyRateInPrecision\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minSellRateInPrecision\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pMinInFp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxQtyInFp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyRateInPrecision\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"ReserveAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"rInFp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"pMinInFp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numFpBits\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxCapBuyInFp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxEthCapSellInFp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"feeInBps\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"formulaPrecision\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxQtyInFp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxBuyRateInPrecision\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minBuyRateInPrecision\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxSellRateInPrecision\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minSellRateInPrecision\",\"type\":\"uint256\"}],\"name\":\"LiquidityParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"resetFeesInTwei\",\"type\":\"uint256\"}],\"name\":\"CollectedFeesReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"}]"

// LiquidityConversionRatesBin is the compiled bytecode used for deploying new contracts.
var LiquidityConversionRatesBin = "0x6060604052600060115534156200001557600080fd5b60405160408062002382833981016040528080519190602001805160008054600160a060020a03191633600160a060020a0316179055915062000068905082640100000000620000df8102620011be1704565b60078054600160a060020a031916600160a060020a038381169190911791829055620000a39116640100000000620001c3810262001efd1704565b600754601290620000cb90600160a060020a031664010000000062001dea620002a382021704565b1115620000d757600080fd5b505062000386565b60005433600160a060020a03908116911614620000fb57600080fd5b600160a060020a03811615156200011157600080fd5b7f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4081604051600160a060020a03909116815260200160405180910390a16000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed908290600160a060020a0316604051600160a060020a039283168152911660208201526040908101905180910390a160008054600160a060020a031916600160a060020a0392909216919091179055565b600160a060020a03811673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14156200020b57600160a060020a038116600090815260066020526040902060129055620002a0565b80600160a060020a031663313ce5676000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156200026b57600080fd5b6102c65a03f115156200027d57600080fd5b5050506040518051600160a060020a038316600090815260066020526040902055505b50565b600080600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415620002d6576012915062000380565b50600160a060020a0382166000908152600660205260409020548015156200037c5782600160a060020a031663313ce5676000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156200035857600080fd5b6102c65a03f115156200036a57600080fd5b50505060405180519050915062000380565b8091505b50919050565b611fec80620003966000396000f3006060604052600436106102795763ffffffff60e060020a60003504166301a12fd3811461027e5780630f9b51291461029f57806314673d31146102c45780631b5ac4b5146102e35780631f05ff29146102f957806320b0961c1461030f578063267822471461032e578063275acbe31461035d578063279fe9671461037957806327a099d81461038c5780632f4fda30146103f25780633ccdbb2814610405578063402776041461042e578063408ee7fe14610441578063436f64ac14610460578063463cf7301461047357806347be7bce146104865780634857d52d146104995780635111249e146104c45780635909e897146104f1578063625cfc461461050a5780636f3d80431461052057806371f805bf1461053657806375829def1461054957806377f50f97146105685780637acc86781461057b5780637c423f541461059a57806382f19e3a146105ad5780638369ff08146105cf5780638401824f146105e8578063859618641461061c578063869d7d9314610646578063925176d61461065f578063958186031461067b5780639870d7fe14610691578063a0099b60146106b0578063a0a7299b146106d5578063a0dbde9d146106e8578063a2c99d4714610704578063a7f43acd1461071f578063aa98d57b14610732578063ac8a584a14610748578063b5debaf514610767578063b86f6aa714610783578063b8e9c22e14610796578063c6fd2103146107c0578063ce56c454146107e8578063debc74f61461080a578063e255d5ad14610820578063e570270114610833578063ec6b16ca14610846578063f0247f7814610859578063f851a4401461086c578063fbe3462c1461087f578063fc0c546a14610892575b600080fd5b341561028957600080fd5b61029d600160a060020a03600435166108a5565b005b34156102aa57600080fd5b6102b2610a15565b60405190815260200160405180910390f35b34156102cf57600080fd5b61029d600160a060020a0360043516610a1b565b34156102ee57600080fd5b6102b2600435610a91565b341561030457600080fd5b6102b2600435610aae565b341561031a57600080fd5b6102b2600435602435604435606435610ae0565b341561033957600080fd5b610341610b20565b604051600160a060020a03909116815260200160405180910390f35b341561036857600080fd5b6102b2600435602435604435610b2f565b341561038457600080fd5b6102b2610b82565b341561039757600080fd5b61039f610b88565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156103de5780820151838201526020016103c6565b505050509050019250505060405180910390f35b34156103fd57600080fd5b6102b2610bf1565b341561041057600080fd5b61029d600160a060020a036004358116906024359060443516610bf9565b341561043957600080fd5b6102b2610cf0565b341561044c57600080fd5b61029d600160a060020a0360043516610cf6565b341561046b57600080fd5b6102b2610df2565b341561047e57600080fd5b6102b2610df8565b341561049157600080fd5b6102b2610dfe565b34156104a457600080fd5b61029d60043560243560443560643560843560a43560c43560e435610e04565b34156104cf57600080fd5b6104dd600435602435610fa4565b604051901515815260200160405180910390f35b34156104fc57600080fd5b6102b2600435602435610fcf565b341561051557600080fd5b6102b2600435611020565b341561052b57600080fd5b6102b2600435611060565b341561054157600080fd5b6102b2611083565b341561055457600080fd5b61029d600160a060020a0360043516611089565b341561057357600080fd5b61029d611124565b341561058657600080fd5b61029d600160a060020a03600435166111be565b34156105a557600080fd5b61039f6112a0565b34156105b857600080fd5b6102b2600435602435604435606435608435611306565b34156105da57600080fd5b6102b26004356024356113d0565b34156105f357600080fd5b61060460043560243560443561144e565b60405191825260208201526040908101905180910390f35b341561062757600080fd5b6102b2600160a060020a03600435166024351515604435606435611493565b341561065157600080fd5b6102b260043560243561159c565b341561066a57600080fd5b610604600435602435604435611625565b341561068657600080fd5b6102b2600435611670565b341561069c57600080fd5b61029d600160a060020a03600435166116b1565b34156106bb57600080fd5b6102b260043560243560443560643560843560a435611781565b34156106e057600080fd5b6102b2611826565b34156106f357600080fd5b6102b260043560243560443561182c565b341561070f57600080fd5b6102b260043560243515156118df565b341561072a57600080fd5b61034161193e565b341561073d57600080fd5b6102b260043561194d565b341561075357600080fd5b61029d600160a060020a036004351661197d565b341561077257600080fd5b6102b2600435602435604435611ae9565b341561078e57600080fd5b61029d611bc0565b34156107a157600080fd5b6102b2600160a060020a03600435166024356044351515606435611c1c565b34156107cb57600080fd5b61029d600160a060020a0360043516602435604435606435611c80565b34156107f357600080fd5b61029d600435600160a060020a0360243516611cf0565b341561081557600080fd5b6102b2600435611d83565b341561082b57600080fd5b6102b2611dae565b341561083e57600080fd5b6102b2611db4565b341561085157600080fd5b6102b2611dba565b341561086457600080fd5b6102b2611dc0565b341561087757600080fd5b610341611dc6565b341561088a57600080fd5b6102b2611dd5565b341561089d57600080fd5b610341611ddb565b6000805433600160a060020a039081169116146108c157600080fd5b600160a060020a03821660009081526003602052604090205460ff1615156108e857600080fd5b50600160a060020a0381166000908152600360205260408120805460ff191690555b600554811015610a115781600160a060020a031660058281548110151561092d57fe5b600091825260209091200154600160a060020a03161415610a095760058054600019810190811061095a57fe5b60009182526020909120015460058054600160a060020a03909216918390811061098057fe5b60009182526020909120018054600160a060020a031916600160a060020a039290921691909117905560058054906109bc906000198301611ea4565b507f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762826000604051600160a060020a039092168252151560208201526040908101905180910390a1610a11565b60010161090a565b5050565b60145481565b60005433600160a060020a03908116911614610a3657600080fd5b60088054600160a060020a031916600160a060020a0383161790557fbd2ca09dd2b354751631db75d1a63231ec123c0d68c81928ea03d0be326c7f8881604051600160a060020a03909116815260200160405180910390a150565b600080821215610aa657506000198102610aa9565b50805b919050565b60006b204fce5e3e25026110000000821115610ac957600080fd5b600a54670de0b6b3a76400009083025b0492915050565b600080610af284870284850285611ae9565b9050610afe8186610fa4565b15610b0857600080fd5b82818602811515610b1557fe5b049695505050505050565b600154600160a060020a031681565b6000690177c17eb2ae5edd211c69021e19e0c9bab240000082610b5387878761182c565b9050610b5f8382610fa4565b15610b6957600080fd5b81818402811515610b7657fe5b04979650505050505050565b600e5481565b610b90611ecd565b6004805480602002602001604051908101604052809291908181526020018280548015610be657602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610bc8575b505050505090505b90565b60c860020a81565b60005433600160a060020a03908116911614610c1457600080fd5b82600160a060020a031663a9059cbb828460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610c7157600080fd5b6102c65a03f11515610c8257600080fd5b505050604051805190501515610c9757600080fd5b7f72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6838383604051600160a060020a03938416815260208101929092529091166040808301919091526060909101905180910390a1505050565b60115481565b60005433600160a060020a03908116911614610d1157600080fd5b600160a060020a03811660009081526003602052604090205460ff1615610d3757600080fd5b60055460329010610d4757600080fd5b7f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600360205260409020805460ff191660019081179091556005805490918101610dc68382611ea4565b5060009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055565b600b5481565b600d5481565b600a5481565b60005433600160a060020a03908116911614610e1f57600080fd5b6101008610610e2d57600080fd5b600a546b204fce5e3e25026110000000901115610e4957600080fd5b6127108310610e5757600080fd5b818110610e6357600080fd5b600b889055600c879055600286900a600a55610e8a6b204fce5e3e25026110000000610aae565b600f556009869055610e9b85610aae565b600d55610ea784610aae565b600e556010839055806ec097ce7bc90715b34b9f1000000000811515610ec957fe5b04601255816ec097ce7bc90715b34b9f1000000000811515610ee757fe5b04601381905560148390556015829055600b54600c54600954600d54600e54601054600a54600f546012547f52db0a06d138736a4425764a1f7e1b432b5ce79099d523c0c4cd01e7320aba0e998c8c6040519b8c5260208c019a909a526040808c019990995260608b019790975260808a019590955260a089019390935260c088019190915260e0870152610100860152610120850152610140840152610160830191909152610180909101905180910390a15050505050505050565b6000811515610fb557506000610fc9565b8282838502811515610fc357fe5b04141590505b92915050565b600080610fe5600b54600c548686600a54611306565b600f54909150811115610ff757600080fd5b61100081611060565b905082670de0b6b3a7640000820281151561101757fe5b04949350505050565b600080600a54670de0b6b3a7640000611041600b54600c5487600a54610ae0565b0281151561104b57fe5b04905061105781611060565b91505b50919050565b600060c860020a82111561107357600080fd5b6010546127109081038302610ad9565b60095481565b60005433600160a060020a039081169116146110a457600080fd5b600160a060020a03811615156110b957600080fd5b6001547f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4090600160a060020a0316604051600160a060020a03909116815260200160405180910390a160018054600160a060020a031916600160a060020a0392909216919091179055565b60015433600160a060020a0390811691161461113f57600080fd5b6001546000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed91600160a060020a039081169116604051600160a060020a039283168152911660208201526040908101905180910390a16001805460008054600160a060020a0319908116600160a060020a03841617909155169055565b60005433600160a060020a039081169116146111d957600080fd5b600160a060020a03811615156111ee57600080fd5b7f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4081604051600160a060020a03909116815260200160405180910390a16000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed908290600160a060020a0316604051600160a060020a039283168152911660208201526040908101905180910390a160008054600160a060020a031916600160a060020a0392909216919091179055565b6112a8611ecd565b6005805480602002602001604051908101604052809291908181526020018280548015610be657602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610bc8575050505050905090565b60008060008061131889898988610ae0565b9250828902915061132e868a0286870287611ae9565b90508481101561133d57600080fd5b61134985820386610fa4565b1561135357600080fd5b611361858683030286610fa4565b1561136b57600080fd5b61137b8586878403020286610fa4565b1561138557600080fd5b61138f8282610fa4565b1561139957600080fd5b6113a38984610fa4565b156113ad57600080fd5b8082028586878885030202028115156113c257fe5b049998505050505050505050565b600080600283810a90810281838188108015906113ed5750828811155b15156113f857600080fd5b607d871061140557600080fd5b50855b6000811115611442578388890281151561141e57fe5b04975060028204915082881061143957600288049750938101935b60001901611408565b50929695505050505050565b60008082830284101561146557508390508261148b565b611486838681151561147357fe5b04848681151561147f57fe5b048561144e565b915091505b935093915050565b6000808080806b204fce5e3e250261100000008711156114b257600080fd5b600f548611156114c157600080fd5b600754600160a060020a038a81169116146114df5760009450611590565b871561152c576114ee87610aae565b9350600d548411156115035760009450611590565b83151561151a5761151386611d83565b9050611527565b6115248685610fcf565b90505b611580565b61153587611670565b925061154083611060565b915081151561155d5761155286611020565b90506000935061156d565b611568868484611625565b945090505b600e548411156115805760009450611590565b61158a81896118df565b90508094505b50505050949350505050565b60007f800000000000000000000000000000000000000000000000000000000000000060ff5b600081126116175783828386028115156115d857fe5b04146115e95760028204915061160e565b6000828502868115156115f857fe5b0411156116075780925061161d565b6002820491505b600019016115c2565b60001992505b505092915050565b60008061163e600b54600c548786600a54600954611781565b600f5490915081111561165057600080fd5b83670de0b6b3a7640000820281151561166557fe5b049150935093915050565b60006b204fce5e3e2502611000000082111561168b57600080fd5b6007546116a090600160a060020a0316611dea565b600a0a600a548302811515610ad957fe5b60005433600160a060020a039081169116146116cc57600080fd5b600160a060020a03811660009081526002602052604090205460ff16156116f257600080fd5b6004546032901061170257600080fd5b7f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600260205260409020805460ff191660019081179091556004805490918101610dc68382611ea4565b6000806000806117938a8a8a89610ae0565b9250828a0291506117b8868884028115156117aa57fe5b048788020187880287610b2f565b90506117c48a84610fa4565b156117ce57600080fd5b6117d88687610fa4565b156117e257600080fd5b6117ec8288610fa4565b156117f657600080fd5b6118008187610fa4565b1561180a57600080fd5b8986820281151561181757fe5b049a9950505050505050505050565b60105481565b600080600283900a81808688111561184b57611848888861159c565b93505b6118558884610fa4565b1561185f57600080fd5b6118698484610fa4565b1561187357600080fd5b611881600285900a88610fa4565b1561188b57600080fd5b600284900a870288840281151561189e57fe5b0491506118ab82876113d0565b905060c860020a84840211156118c057600080fd5b60c860020a8111156118d157600080fd5b919092020195945050505050565b600080600083156118f7575050601354601254611900565b50506015546014545b8085118061190d57508185105b1561191b576000925061161d565b69d3c21bcecceda1000000851115611936576000925061161d565b84925061161d565b600854600160a060020a031681565b60006b204fce5e3e2502611000000082111561196857600080fd5b601054612710036010548302811515610ad957fe5b6000805433600160a060020a0390811691161461199957600080fd5b600160a060020a03821660009081526002602052604090205460ff1615156119c057600080fd5b50600160a060020a0381166000908152600260205260408120805460ff191690555b600454811015610a115781600160a060020a0316600482815481101515611a0557fe5b600091825260209091200154600160a060020a03161415611ae157600480546000198101908110611a3257fe5b60009182526020909120015460048054600160a060020a039092169183908110611a5857fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600480546000190190611a949082611ea4565b507f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b826000604051600160a060020a039092168252151560208201526040908101905180910390a1610a11565b6001016119e2565b6000806001808083805b611afd8489610fa4565b15611b0a57819650611bb3565b611b148386610fa4565b15611b2157819650611bb3565b848302888502811515611b3057fe5b048201915080821415611b4557819650611bb3565b5060019094019380611b57848b610fa4565b15611b6457819650611bb3565b611b6e838a610fa4565b15611b7b57819650611bb3565b611b858587610fa4565b15611b9257819650611bb3565b938502939289029291880291611ba984848a61144e565b9094509250611af3565b5050505050509392505050565b6000805433600160a060020a03908116911614611bdc57600080fd5b506011805460009091557fdeb4766cf1de6f18e3b195f199d403a02a3e09fbee1192b37d797fb300f052618160405190815260200160405180910390a150565b600080806b204fce5e3e25026110000000841115611c3957600080fd5b600854611c4f90600160a060020a031631610aae565b9150611c5d87868685611493565b905069d3c21bcecceda1000000811115611c7657600080fd5b9695505050505050565b60085433600160a060020a03908116911614611c9b57600080fd5b6000831315611cc357611cb5611cb084610a91565b61194d565b601180549091019055611cea565b612710601054611cd285610a91565b02811515611cdc57fe5b601180549290910490910190555b50505050565b60005433600160a060020a03908116911614611d0b57600080fd5b600160a060020a03811682156108fc0283604051600060405180830381858888f193505050501515611d3c57600080fd5b7fec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de8282604051918252600160a060020a031660208201526040908101905180910390a15050565b600080611d98600b54600c5485600a54610ae0565b670de0b6b3a7640000600a540281151561104b57fe5b60125481565b60155481565b600c5481565b600f5481565b600054600160a060020a031681565b60135481565b600754600160a060020a031681565b600080600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415611e1b576012915061105a565b50600160a060020a038216600090815260066020526040902054801515610fc95782600160a060020a031663313ce5676000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611e8257600080fd5b6102c65a03f11515611e9357600080fd5b50505060405180519050915061105a565b815481835581811511611ec857600083815260209020611ec8918101908301611edf565b505050565b60206040519081016040526000815290565b610bee91905b80821115611ef95760008155600101611ee5565b5090565b600160a060020a03811673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415611f4357600160a060020a038116600090815260066020526040902060129055611fbd565b80600160a060020a031663313ce5676000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611f8957600080fd5b6102c65a03f11515611f9a57600080fd5b5050506040518051600160a060020a038316600090815260066020526040902055505b505600a165627a7a72305820cded95b938ed6cd66bc063fe8a36714b7423ba61d8f5ff5dac0a1c78ea7eb5980029"

// DeployLiquidityConversionRates deploys a new Ethereum contract, binding an instance of LiquidityConversionRates to it.
func DeployLiquidityConversionRates(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _token common.Address) (common.Address, *types.Transaction, *LiquidityConversionRates, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidityConversionRatesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LiquidityConversionRatesBin), backend, _admin, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LiquidityConversionRates{LiquidityConversionRatesCaller: LiquidityConversionRatesCaller{contract: contract}, LiquidityConversionRatesTransactor: LiquidityConversionRatesTransactor{contract: contract}, LiquidityConversionRatesFilterer: LiquidityConversionRatesFilterer{contract: contract}}, nil
}

// LiquidityConversionRates is an auto generated Go binding around an Ethereum contract.
type LiquidityConversionRates struct {
	LiquidityConversionRatesCaller     // Read-only binding to the contract
	LiquidityConversionRatesTransactor // Write-only binding to the contract
	LiquidityConversionRatesFilterer   // Log filterer for contract events
}

// LiquidityConversionRatesCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityConversionRatesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityConversionRatesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityConversionRatesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityConversionRatesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityConversionRatesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityConversionRatesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityConversionRatesSession struct {
	Contract     *LiquidityConversionRates // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LiquidityConversionRatesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityConversionRatesCallerSession struct {
	Contract *LiquidityConversionRatesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// LiquidityConversionRatesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityConversionRatesTransactorSession struct {
	Contract     *LiquidityConversionRatesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// LiquidityConversionRatesRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityConversionRatesRaw struct {
	Contract *LiquidityConversionRates // Generic contract binding to access the raw methods on
}

// LiquidityConversionRatesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityConversionRatesCallerRaw struct {
	Contract *LiquidityConversionRatesCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityConversionRatesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityConversionRatesTransactorRaw struct {
	Contract *LiquidityConversionRatesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityConversionRates creates a new instance of LiquidityConversionRates, bound to a specific deployed contract.
func NewLiquidityConversionRates(address common.Address, backend bind.ContractBackend) (*LiquidityConversionRates, error) {
	contract, err := bindLiquidityConversionRates(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRates{LiquidityConversionRatesCaller: LiquidityConversionRatesCaller{contract: contract}, LiquidityConversionRatesTransactor: LiquidityConversionRatesTransactor{contract: contract}, LiquidityConversionRatesFilterer: LiquidityConversionRatesFilterer{contract: contract}}, nil
}

// NewLiquidityConversionRatesCaller creates a new read-only instance of LiquidityConversionRates, bound to a specific deployed contract.
func NewLiquidityConversionRatesCaller(address common.Address, caller bind.ContractCaller) (*LiquidityConversionRatesCaller, error) {
	contract, err := bindLiquidityConversionRates(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesCaller{contract: contract}, nil
}

// NewLiquidityConversionRatesTransactor creates a new write-only instance of LiquidityConversionRates, bound to a specific deployed contract.
func NewLiquidityConversionRatesTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityConversionRatesTransactor, error) {
	contract, err := bindLiquidityConversionRates(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesTransactor{contract: contract}, nil
}

// NewLiquidityConversionRatesFilterer creates a new log filterer instance of LiquidityConversionRates, bound to a specific deployed contract.
func NewLiquidityConversionRatesFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityConversionRatesFilterer, error) {
	contract, err := bindLiquidityConversionRates(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesFilterer{contract: contract}, nil
}

// bindLiquidityConversionRates binds a generic wrapper to an already deployed contract.
func bindLiquidityConversionRates(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidityConversionRatesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityConversionRates *LiquidityConversionRatesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LiquidityConversionRates.Contract.LiquidityConversionRatesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityConversionRates *LiquidityConversionRatesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.LiquidityConversionRatesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityConversionRates *LiquidityConversionRatesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.LiquidityConversionRatesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityConversionRates *LiquidityConversionRatesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LiquidityConversionRates.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.contract.Transact(opts, method, params...)
}

// BIGNUMBER is a free data retrieval call binding the contract method 0x2f4fda30.
//
// Solidity: function BIG_NUMBER() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) BIGNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "BIG_NUMBER")
	return *ret0, err
}

// BIGNUMBER is a free data retrieval call binding the contract method 0x2f4fda30.
//
// Solidity: function BIG_NUMBER() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) BIGNUMBER() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.BIGNUMBER(&_LiquidityConversionRates.CallOpts)
}

// BIGNUMBER is a free data retrieval call binding the contract method 0x2f4fda30.
//
// Solidity: function BIG_NUMBER() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) BIGNUMBER() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.BIGNUMBER(&_LiquidityConversionRates.CallOpts)
}

// Abs is a free data retrieval call binding the contract method 0x1b5ac4b5.
//
// Solidity: function abs(int256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) Abs(opts *bind.CallOpts, val *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "abs", val)
	return *ret0, err
}

// Abs is a free data retrieval call binding the contract method 0x1b5ac4b5.
//
// Solidity: function abs(int256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) Abs(val *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.Abs(&_LiquidityConversionRates.CallOpts, val)
}

// Abs is a free data retrieval call binding the contract method 0x1b5ac4b5.
//
// Solidity: function abs(int256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) Abs(val *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.Abs(&_LiquidityConversionRates.CallOpts, val)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) Admin() (common.Address, error) {
	return _LiquidityConversionRates.Contract.Admin(&_LiquidityConversionRates.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) Admin() (common.Address, error) {
	return _LiquidityConversionRates.Contract.Admin(&_LiquidityConversionRates.CallOpts)
}

// BuyRate is a free data retrieval call binding the contract method 0x5909e897.
//
// Solidity: function buyRate(uint256 eInFp, uint256 deltaEInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) BuyRate(opts *bind.CallOpts, eInFp *big.Int, deltaEInFp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "buyRate", eInFp, deltaEInFp)
	return *ret0, err
}

// BuyRate is a free data retrieval call binding the contract method 0x5909e897.
//
// Solidity: function buyRate(uint256 eInFp, uint256 deltaEInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) BuyRate(eInFp *big.Int, deltaEInFp *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.BuyRate(&_LiquidityConversionRates.CallOpts, eInFp, deltaEInFp)
}

// BuyRate is a free data retrieval call binding the contract method 0x5909e897.
//
// Solidity: function buyRate(uint256 eInFp, uint256 deltaEInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) BuyRate(eInFp *big.Int, deltaEInFp *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.BuyRate(&_LiquidityConversionRates.CallOpts, eInFp, deltaEInFp)
}

// BuyRateZeroQuantity is a free data retrieval call binding the contract method 0xdebc74f6.
//
// Solidity: function buyRateZeroQuantity(uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) BuyRateZeroQuantity(opts *bind.CallOpts, eInFp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "buyRateZeroQuantity", eInFp)
	return *ret0, err
}

// BuyRateZeroQuantity is a free data retrieval call binding the contract method 0xdebc74f6.
//
// Solidity: function buyRateZeroQuantity(uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) BuyRateZeroQuantity(eInFp *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.BuyRateZeroQuantity(&_LiquidityConversionRates.CallOpts, eInFp)
}

// BuyRateZeroQuantity is a free data retrieval call binding the contract method 0xdebc74f6.
//
// Solidity: function buyRateZeroQuantity(uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) BuyRateZeroQuantity(eInFp *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.BuyRateZeroQuantity(&_LiquidityConversionRates.CallOpts, eInFp)
}

// CalcCollectedFee is a free data retrieval call binding the contract method 0xaa98d57b.
//
// Solidity: function calcCollectedFee(uint256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) CalcCollectedFee(opts *bind.CallOpts, val *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "calcCollectedFee", val)
	return *ret0, err
}

// CalcCollectedFee is a free data retrieval call binding the contract method 0xaa98d57b.
//
// Solidity: function calcCollectedFee(uint256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) CalcCollectedFee(val *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.CalcCollectedFee(&_LiquidityConversionRates.CallOpts, val)
}

// CalcCollectedFee is a free data retrieval call binding the contract method 0xaa98d57b.
//
// Solidity: function calcCollectedFee(uint256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) CalcCollectedFee(val *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.CalcCollectedFee(&_LiquidityConversionRates.CallOpts, val)
}

// CheckMultOverflow is a free data retrieval call binding the contract method 0x5111249e.
//
// Solidity: function checkMultOverflow(uint256 x, uint256 y) constant returns(bool)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) CheckMultOverflow(opts *bind.CallOpts, x *big.Int, y *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "checkMultOverflow", x, y)
	return *ret0, err
}

// CheckMultOverflow is a free data retrieval call binding the contract method 0x5111249e.
//
// Solidity: function checkMultOverflow(uint256 x, uint256 y) constant returns(bool)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) CheckMultOverflow(x *big.Int, y *big.Int) (bool, error) {
	return _LiquidityConversionRates.Contract.CheckMultOverflow(&_LiquidityConversionRates.CallOpts, x, y)
}

// CheckMultOverflow is a free data retrieval call binding the contract method 0x5111249e.
//
// Solidity: function checkMultOverflow(uint256 x, uint256 y) constant returns(bool)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) CheckMultOverflow(x *big.Int, y *big.Int) (bool, error) {
	return _LiquidityConversionRates.Contract.CheckMultOverflow(&_LiquidityConversionRates.CallOpts, x, y)
}

// CollectedFeesInTwei is a free data retrieval call binding the contract method 0x40277604.
//
// Solidity: function collectedFeesInTwei() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) CollectedFeesInTwei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "collectedFeesInTwei")
	return *ret0, err
}

// CollectedFeesInTwei is a free data retrieval call binding the contract method 0x40277604.
//
// Solidity: function collectedFeesInTwei() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) CollectedFeesInTwei() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.CollectedFeesInTwei(&_LiquidityConversionRates.CallOpts)
}

// CollectedFeesInTwei is a free data retrieval call binding the contract method 0x40277604.
//
// Solidity: function collectedFeesInTwei() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) CollectedFeesInTwei() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.CollectedFeesInTwei(&_LiquidityConversionRates.CallOpts)
}

// CompactFraction is a free data retrieval call binding the contract method 0x8401824f.
//
// Solidity: function compactFraction(uint256 p, uint256 q, uint256 precision) constant returns(uint256, uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) CompactFraction(opts *bind.CallOpts, p *big.Int, q *big.Int, precision *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _LiquidityConversionRates.contract.Call(opts, out, "compactFraction", p, q, precision)
	return *ret0, *ret1, err
}

// CompactFraction is a free data retrieval call binding the contract method 0x8401824f.
//
// Solidity: function compactFraction(uint256 p, uint256 q, uint256 precision) constant returns(uint256, uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) CompactFraction(p *big.Int, q *big.Int, precision *big.Int) (*big.Int, *big.Int, error) {
	return _LiquidityConversionRates.Contract.CompactFraction(&_LiquidityConversionRates.CallOpts, p, q, precision)
}

// CompactFraction is a free data retrieval call binding the contract method 0x8401824f.
//
// Solidity: function compactFraction(uint256 p, uint256 q, uint256 precision) constant returns(uint256, uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) CompactFraction(p *big.Int, q *big.Int, precision *big.Int) (*big.Int, *big.Int, error) {
	return _LiquidityConversionRates.Contract.CompactFraction(&_LiquidityConversionRates.CallOpts, p, q, precision)
}

// CountLeadingZeros is a free data retrieval call binding the contract method 0x869d7d93.
//
// Solidity: function countLeadingZeros(uint256 p, uint256 q) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) CountLeadingZeros(opts *bind.CallOpts, p *big.Int, q *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "countLeadingZeros", p, q)
	return *ret0, err
}

// CountLeadingZeros is a free data retrieval call binding the contract method 0x869d7d93.
//
// Solidity: function countLeadingZeros(uint256 p, uint256 q) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) CountLeadingZeros(p *big.Int, q *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.CountLeadingZeros(&_LiquidityConversionRates.CallOpts, p, q)
}

// CountLeadingZeros is a free data retrieval call binding the contract method 0x869d7d93.
//
// Solidity: function countLeadingZeros(uint256 p, uint256 q) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) CountLeadingZeros(p *big.Int, q *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.CountLeadingZeros(&_LiquidityConversionRates.CallOpts, p, q)
}

// DeltaEFunc is a free data retrieval call binding the contract method 0xa0099b60.
//
// Solidity: function deltaEFunc(uint256 r, uint256 pMIn, uint256 e, uint256 deltaT, uint256 precision, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) DeltaEFunc(opts *bind.CallOpts, r *big.Int, pMIn *big.Int, e *big.Int, deltaT *big.Int, precision *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "deltaEFunc", r, pMIn, e, deltaT, precision, numPrecisionBits)
	return *ret0, err
}

// DeltaEFunc is a free data retrieval call binding the contract method 0xa0099b60.
//
// Solidity: function deltaEFunc(uint256 r, uint256 pMIn, uint256 e, uint256 deltaT, uint256 precision, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) DeltaEFunc(r *big.Int, pMIn *big.Int, e *big.Int, deltaT *big.Int, precision *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.DeltaEFunc(&_LiquidityConversionRates.CallOpts, r, pMIn, e, deltaT, precision, numPrecisionBits)
}

// DeltaEFunc is a free data retrieval call binding the contract method 0xa0099b60.
//
// Solidity: function deltaEFunc(uint256 r, uint256 pMIn, uint256 e, uint256 deltaT, uint256 precision, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) DeltaEFunc(r *big.Int, pMIn *big.Int, e *big.Int, deltaT *big.Int, precision *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.DeltaEFunc(&_LiquidityConversionRates.CallOpts, r, pMIn, e, deltaT, precision, numPrecisionBits)
}

// DeltaTFunc is a free data retrieval call binding the contract method 0x82f19e3a.
//
// Solidity: function deltaTFunc(uint256 r, uint256 pMIn, uint256 e, uint256 deltaE, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) DeltaTFunc(opts *bind.CallOpts, r *big.Int, pMIn *big.Int, e *big.Int, deltaE *big.Int, precision *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "deltaTFunc", r, pMIn, e, deltaE, precision)
	return *ret0, err
}

// DeltaTFunc is a free data retrieval call binding the contract method 0x82f19e3a.
//
// Solidity: function deltaTFunc(uint256 r, uint256 pMIn, uint256 e, uint256 deltaE, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) DeltaTFunc(r *big.Int, pMIn *big.Int, e *big.Int, deltaE *big.Int, precision *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.DeltaTFunc(&_LiquidityConversionRates.CallOpts, r, pMIn, e, deltaE, precision)
}

// DeltaTFunc is a free data retrieval call binding the contract method 0x82f19e3a.
//
// Solidity: function deltaTFunc(uint256 r, uint256 pMIn, uint256 e, uint256 deltaE, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) DeltaTFunc(r *big.Int, pMIn *big.Int, e *big.Int, deltaE *big.Int, precision *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.DeltaTFunc(&_LiquidityConversionRates.CallOpts, r, pMIn, e, deltaE, precision)
}

// Exp is a free data retrieval call binding the contract method 0xb5debaf5.
//
// Solidity: function exp(uint256 p, uint256 q, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) Exp(opts *bind.CallOpts, p *big.Int, q *big.Int, precision *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "exp", p, q, precision)
	return *ret0, err
}

// Exp is a free data retrieval call binding the contract method 0xb5debaf5.
//
// Solidity: function exp(uint256 p, uint256 q, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) Exp(p *big.Int, q *big.Int, precision *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.Exp(&_LiquidityConversionRates.CallOpts, p, q, precision)
}

// Exp is a free data retrieval call binding the contract method 0xb5debaf5.
//
// Solidity: function exp(uint256 p, uint256 q, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) Exp(p *big.Int, q *big.Int, precision *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.Exp(&_LiquidityConversionRates.CallOpts, p, q, precision)
}

// FeeInBps is a free data retrieval call binding the contract method 0xa0a7299b.
//
// Solidity: function feeInBps() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) FeeInBps(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "feeInBps")
	return *ret0, err
}

// FeeInBps is a free data retrieval call binding the contract method 0xa0a7299b.
//
// Solidity: function feeInBps() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) FeeInBps() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.FeeInBps(&_LiquidityConversionRates.CallOpts)
}

// FeeInBps is a free data retrieval call binding the contract method 0xa0a7299b.
//
// Solidity: function feeInBps() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) FeeInBps() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.FeeInBps(&_LiquidityConversionRates.CallOpts)
}

// FormulaPrecision is a free data retrieval call binding the contract method 0x47be7bce.
//
// Solidity: function formulaPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) FormulaPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "formulaPrecision")
	return *ret0, err
}

// FormulaPrecision is a free data retrieval call binding the contract method 0x47be7bce.
//
// Solidity: function formulaPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) FormulaPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.FormulaPrecision(&_LiquidityConversionRates.CallOpts)
}

// FormulaPrecision is a free data retrieval call binding the contract method 0x47be7bce.
//
// Solidity: function formulaPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) FormulaPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.FormulaPrecision(&_LiquidityConversionRates.CallOpts)
}

// FromTweiToFp is a free data retrieval call binding the contract method 0x95818603.
//
// Solidity: function fromTweiToFp(uint256 qtyInTwei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) FromTweiToFp(opts *bind.CallOpts, qtyInTwei *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "fromTweiToFp", qtyInTwei)
	return *ret0, err
}

// FromTweiToFp is a free data retrieval call binding the contract method 0x95818603.
//
// Solidity: function fromTweiToFp(uint256 qtyInTwei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) FromTweiToFp(qtyInTwei *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.FromTweiToFp(&_LiquidityConversionRates.CallOpts, qtyInTwei)
}

// FromTweiToFp is a free data retrieval call binding the contract method 0x95818603.
//
// Solidity: function fromTweiToFp(uint256 qtyInTwei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) FromTweiToFp(qtyInTwei *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.FromTweiToFp(&_LiquidityConversionRates.CallOpts, qtyInTwei)
}

// FromWeiToFp is a free data retrieval call binding the contract method 0x1f05ff29.
//
// Solidity: function fromWeiToFp(uint256 qtyInwei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) FromWeiToFp(opts *bind.CallOpts, qtyInwei *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "fromWeiToFp", qtyInwei)
	return *ret0, err
}

// FromWeiToFp is a free data retrieval call binding the contract method 0x1f05ff29.
//
// Solidity: function fromWeiToFp(uint256 qtyInwei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) FromWeiToFp(qtyInwei *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.FromWeiToFp(&_LiquidityConversionRates.CallOpts, qtyInwei)
}

// FromWeiToFp is a free data retrieval call binding the contract method 0x1f05ff29.
//
// Solidity: function fromWeiToFp(uint256 qtyInwei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) FromWeiToFp(qtyInwei *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.FromWeiToFp(&_LiquidityConversionRates.CallOpts, qtyInwei)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_LiquidityConversionRates *LiquidityConversionRatesSession) GetAlerters() ([]common.Address, error) {
	return _LiquidityConversionRates.Contract.GetAlerters(&_LiquidityConversionRates.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) GetAlerters() ([]common.Address, error) {
	return _LiquidityConversionRates.Contract.GetAlerters(&_LiquidityConversionRates.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_LiquidityConversionRates *LiquidityConversionRatesSession) GetOperators() ([]common.Address, error) {
	return _LiquidityConversionRates.Contract.GetOperators(&_LiquidityConversionRates.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) GetOperators() ([]common.Address, error) {
	return _LiquidityConversionRates.Contract.GetOperators(&_LiquidityConversionRates.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0xb8e9c22e.
//
// Solidity: function getRate(address conversionToken, uint256 currentBlockNumber, bool buy, uint256 qtyInSrcWei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) GetRate(opts *bind.CallOpts, conversionToken common.Address, currentBlockNumber *big.Int, buy bool, qtyInSrcWei *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "getRate", conversionToken, currentBlockNumber, buy, qtyInSrcWei)
	return *ret0, err
}

// GetRate is a free data retrieval call binding the contract method 0xb8e9c22e.
//
// Solidity: function getRate(address conversionToken, uint256 currentBlockNumber, bool buy, uint256 qtyInSrcWei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) GetRate(conversionToken common.Address, currentBlockNumber *big.Int, buy bool, qtyInSrcWei *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.GetRate(&_LiquidityConversionRates.CallOpts, conversionToken, currentBlockNumber, buy, qtyInSrcWei)
}

// GetRate is a free data retrieval call binding the contract method 0xb8e9c22e.
//
// Solidity: function getRate(address conversionToken, uint256 currentBlockNumber, bool buy, uint256 qtyInSrcWei) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) GetRate(conversionToken common.Address, currentBlockNumber *big.Int, buy bool, qtyInSrcWei *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.GetRate(&_LiquidityConversionRates.CallOpts, conversionToken, currentBlockNumber, buy, qtyInSrcWei)
}

// GetRateWithE is a free data retrieval call binding the contract method 0x85961864.
//
// Solidity: function getRateWithE(address conversionToken, bool buy, uint256 qtyInSrcWei, uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) GetRateWithE(opts *bind.CallOpts, conversionToken common.Address, buy bool, qtyInSrcWei *big.Int, eInFp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "getRateWithE", conversionToken, buy, qtyInSrcWei, eInFp)
	return *ret0, err
}

// GetRateWithE is a free data retrieval call binding the contract method 0x85961864.
//
// Solidity: function getRateWithE(address conversionToken, bool buy, uint256 qtyInSrcWei, uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) GetRateWithE(conversionToken common.Address, buy bool, qtyInSrcWei *big.Int, eInFp *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.GetRateWithE(&_LiquidityConversionRates.CallOpts, conversionToken, buy, qtyInSrcWei, eInFp)
}

// GetRateWithE is a free data retrieval call binding the contract method 0x85961864.
//
// Solidity: function getRateWithE(address conversionToken, bool buy, uint256 qtyInSrcWei, uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) GetRateWithE(conversionToken common.Address, buy bool, qtyInSrcWei *big.Int, eInFp *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.GetRateWithE(&_LiquidityConversionRates.CallOpts, conversionToken, buy, qtyInSrcWei, eInFp)
}

// Ln is a free data retrieval call binding the contract method 0x275acbe3.
//
// Solidity: function ln(uint256 p, uint256 q, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) Ln(opts *bind.CallOpts, p *big.Int, q *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "ln", p, q, numPrecisionBits)
	return *ret0, err
}

// Ln is a free data retrieval call binding the contract method 0x275acbe3.
//
// Solidity: function ln(uint256 p, uint256 q, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) Ln(p *big.Int, q *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.Ln(&_LiquidityConversionRates.CallOpts, p, q, numPrecisionBits)
}

// Ln is a free data retrieval call binding the contract method 0x275acbe3.
//
// Solidity: function ln(uint256 p, uint256 q, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) Ln(p *big.Int, q *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.Ln(&_LiquidityConversionRates.CallOpts, p, q, numPrecisionBits)
}

// Log2ForSmallNumber is a free data retrieval call binding the contract method 0x8369ff08.
//
// Solidity: function log2ForSmallNumber(uint256 x, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) Log2ForSmallNumber(opts *bind.CallOpts, x *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "log2ForSmallNumber", x, numPrecisionBits)
	return *ret0, err
}

// Log2ForSmallNumber is a free data retrieval call binding the contract method 0x8369ff08.
//
// Solidity: function log2ForSmallNumber(uint256 x, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) Log2ForSmallNumber(x *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.Log2ForSmallNumber(&_LiquidityConversionRates.CallOpts, x, numPrecisionBits)
}

// Log2ForSmallNumber is a free data retrieval call binding the contract method 0x8369ff08.
//
// Solidity: function log2ForSmallNumber(uint256 x, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) Log2ForSmallNumber(x *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.Log2ForSmallNumber(&_LiquidityConversionRates.CallOpts, x, numPrecisionBits)
}

// LogBase2 is a free data retrieval call binding the contract method 0xa0dbde9d.
//
// Solidity: function logBase2(uint256 p, uint256 q, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) LogBase2(opts *bind.CallOpts, p *big.Int, q *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "logBase2", p, q, numPrecisionBits)
	return *ret0, err
}

// LogBase2 is a free data retrieval call binding the contract method 0xa0dbde9d.
//
// Solidity: function logBase2(uint256 p, uint256 q, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) LogBase2(p *big.Int, q *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.LogBase2(&_LiquidityConversionRates.CallOpts, p, q, numPrecisionBits)
}

// LogBase2 is a free data retrieval call binding the contract method 0xa0dbde9d.
//
// Solidity: function logBase2(uint256 p, uint256 q, uint256 numPrecisionBits) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) LogBase2(p *big.Int, q *big.Int, numPrecisionBits *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.LogBase2(&_LiquidityConversionRates.CallOpts, p, q, numPrecisionBits)
}

// MaxBuyRateInPrecision is a free data retrieval call binding the contract method 0xe255d5ad.
//
// Solidity: function maxBuyRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) MaxBuyRateInPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "maxBuyRateInPrecision")
	return *ret0, err
}

// MaxBuyRateInPrecision is a free data retrieval call binding the contract method 0xe255d5ad.
//
// Solidity: function maxBuyRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) MaxBuyRateInPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxBuyRateInPrecision(&_LiquidityConversionRates.CallOpts)
}

// MaxBuyRateInPrecision is a free data retrieval call binding the contract method 0xe255d5ad.
//
// Solidity: function maxBuyRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) MaxBuyRateInPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxBuyRateInPrecision(&_LiquidityConversionRates.CallOpts)
}

// MaxEthCapBuyInFp is a free data retrieval call binding the contract method 0x463cf730.
//
// Solidity: function maxEthCapBuyInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) MaxEthCapBuyInFp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "maxEthCapBuyInFp")
	return *ret0, err
}

// MaxEthCapBuyInFp is a free data retrieval call binding the contract method 0x463cf730.
//
// Solidity: function maxEthCapBuyInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) MaxEthCapBuyInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxEthCapBuyInFp(&_LiquidityConversionRates.CallOpts)
}

// MaxEthCapBuyInFp is a free data retrieval call binding the contract method 0x463cf730.
//
// Solidity: function maxEthCapBuyInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) MaxEthCapBuyInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxEthCapBuyInFp(&_LiquidityConversionRates.CallOpts)
}

// MaxEthCapSellInFp is a free data retrieval call binding the contract method 0x279fe967.
//
// Solidity: function maxEthCapSellInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) MaxEthCapSellInFp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "maxEthCapSellInFp")
	return *ret0, err
}

// MaxEthCapSellInFp is a free data retrieval call binding the contract method 0x279fe967.
//
// Solidity: function maxEthCapSellInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) MaxEthCapSellInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxEthCapSellInFp(&_LiquidityConversionRates.CallOpts)
}

// MaxEthCapSellInFp is a free data retrieval call binding the contract method 0x279fe967.
//
// Solidity: function maxEthCapSellInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) MaxEthCapSellInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxEthCapSellInFp(&_LiquidityConversionRates.CallOpts)
}

// MaxQtyInFp is a free data retrieval call binding the contract method 0xf0247f78.
//
// Solidity: function maxQtyInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) MaxQtyInFp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "maxQtyInFp")
	return *ret0, err
}

// MaxQtyInFp is a free data retrieval call binding the contract method 0xf0247f78.
//
// Solidity: function maxQtyInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) MaxQtyInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxQtyInFp(&_LiquidityConversionRates.CallOpts)
}

// MaxQtyInFp is a free data retrieval call binding the contract method 0xf0247f78.
//
// Solidity: function maxQtyInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) MaxQtyInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxQtyInFp(&_LiquidityConversionRates.CallOpts)
}

// MaxSellRateInPrecision is a free data retrieval call binding the contract method 0x0f9b5129.
//
// Solidity: function maxSellRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) MaxSellRateInPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "maxSellRateInPrecision")
	return *ret0, err
}

// MaxSellRateInPrecision is a free data retrieval call binding the contract method 0x0f9b5129.
//
// Solidity: function maxSellRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) MaxSellRateInPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxSellRateInPrecision(&_LiquidityConversionRates.CallOpts)
}

// MaxSellRateInPrecision is a free data retrieval call binding the contract method 0x0f9b5129.
//
// Solidity: function maxSellRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) MaxSellRateInPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MaxSellRateInPrecision(&_LiquidityConversionRates.CallOpts)
}

// MinBuyRateInPrecision is a free data retrieval call binding the contract method 0xfbe3462c.
//
// Solidity: function minBuyRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) MinBuyRateInPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "minBuyRateInPrecision")
	return *ret0, err
}

// MinBuyRateInPrecision is a free data retrieval call binding the contract method 0xfbe3462c.
//
// Solidity: function minBuyRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) MinBuyRateInPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MinBuyRateInPrecision(&_LiquidityConversionRates.CallOpts)
}

// MinBuyRateInPrecision is a free data retrieval call binding the contract method 0xfbe3462c.
//
// Solidity: function minBuyRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) MinBuyRateInPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MinBuyRateInPrecision(&_LiquidityConversionRates.CallOpts)
}

// MinSellRateInPrecision is a free data retrieval call binding the contract method 0xe5702701.
//
// Solidity: function minSellRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) MinSellRateInPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "minSellRateInPrecision")
	return *ret0, err
}

// MinSellRateInPrecision is a free data retrieval call binding the contract method 0xe5702701.
//
// Solidity: function minSellRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) MinSellRateInPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MinSellRateInPrecision(&_LiquidityConversionRates.CallOpts)
}

// MinSellRateInPrecision is a free data retrieval call binding the contract method 0xe5702701.
//
// Solidity: function minSellRateInPrecision() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) MinSellRateInPrecision() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.MinSellRateInPrecision(&_LiquidityConversionRates.CallOpts)
}

// NumFpBits is a free data retrieval call binding the contract method 0x71f805bf.
//
// Solidity: function numFpBits() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) NumFpBits(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "numFpBits")
	return *ret0, err
}

// NumFpBits is a free data retrieval call binding the contract method 0x71f805bf.
//
// Solidity: function numFpBits() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) NumFpBits() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.NumFpBits(&_LiquidityConversionRates.CallOpts)
}

// NumFpBits is a free data retrieval call binding the contract method 0x71f805bf.
//
// Solidity: function numFpBits() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) NumFpBits() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.NumFpBits(&_LiquidityConversionRates.CallOpts)
}

// PE is a free data retrieval call binding the contract method 0x20b0961c.
//
// Solidity: function pE(uint256 r, uint256 pMIn, uint256 e, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) PE(opts *bind.CallOpts, r *big.Int, pMIn *big.Int, e *big.Int, precision *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "pE", r, pMIn, e, precision)
	return *ret0, err
}

// PE is a free data retrieval call binding the contract method 0x20b0961c.
//
// Solidity: function pE(uint256 r, uint256 pMIn, uint256 e, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) PE(r *big.Int, pMIn *big.Int, e *big.Int, precision *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.PE(&_LiquidityConversionRates.CallOpts, r, pMIn, e, precision)
}

// PE is a free data retrieval call binding the contract method 0x20b0961c.
//
// Solidity: function pE(uint256 r, uint256 pMIn, uint256 e, uint256 precision) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) PE(r *big.Int, pMIn *big.Int, e *big.Int, precision *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.PE(&_LiquidityConversionRates.CallOpts, r, pMIn, e, precision)
}

// PMinInFp is a free data retrieval call binding the contract method 0xec6b16ca.
//
// Solidity: function pMinInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) PMinInFp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "pMinInFp")
	return *ret0, err
}

// PMinInFp is a free data retrieval call binding the contract method 0xec6b16ca.
//
// Solidity: function pMinInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) PMinInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.PMinInFp(&_LiquidityConversionRates.CallOpts)
}

// PMinInFp is a free data retrieval call binding the contract method 0xec6b16ca.
//
// Solidity: function pMinInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) PMinInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.PMinInFp(&_LiquidityConversionRates.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) PendingAdmin() (common.Address, error) {
	return _LiquidityConversionRates.Contract.PendingAdmin(&_LiquidityConversionRates.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) PendingAdmin() (common.Address, error) {
	return _LiquidityConversionRates.Contract.PendingAdmin(&_LiquidityConversionRates.CallOpts)
}

// RInFp is a free data retrieval call binding the contract method 0x436f64ac.
//
// Solidity: function rInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) RInFp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "rInFp")
	return *ret0, err
}

// RInFp is a free data retrieval call binding the contract method 0x436f64ac.
//
// Solidity: function rInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) RInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.RInFp(&_LiquidityConversionRates.CallOpts)
}

// RInFp is a free data retrieval call binding the contract method 0x436f64ac.
//
// Solidity: function rInFp() constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) RInFp() (*big.Int, error) {
	return _LiquidityConversionRates.Contract.RInFp(&_LiquidityConversionRates.CallOpts)
}

// RateAfterValidation is a free data retrieval call binding the contract method 0xa2c99d47.
//
// Solidity: function rateAfterValidation(uint256 rateInPrecision, bool buy) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) RateAfterValidation(opts *bind.CallOpts, rateInPrecision *big.Int, buy bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "rateAfterValidation", rateInPrecision, buy)
	return *ret0, err
}

// RateAfterValidation is a free data retrieval call binding the contract method 0xa2c99d47.
//
// Solidity: function rateAfterValidation(uint256 rateInPrecision, bool buy) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) RateAfterValidation(rateInPrecision *big.Int, buy bool) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.RateAfterValidation(&_LiquidityConversionRates.CallOpts, rateInPrecision, buy)
}

// RateAfterValidation is a free data retrieval call binding the contract method 0xa2c99d47.
//
// Solidity: function rateAfterValidation(uint256 rateInPrecision, bool buy) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) RateAfterValidation(rateInPrecision *big.Int, buy bool) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.RateAfterValidation(&_LiquidityConversionRates.CallOpts, rateInPrecision, buy)
}

// ReserveContract is a free data retrieval call binding the contract method 0xa7f43acd.
//
// Solidity: function reserveContract() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) ReserveContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "reserveContract")
	return *ret0, err
}

// ReserveContract is a free data retrieval call binding the contract method 0xa7f43acd.
//
// Solidity: function reserveContract() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) ReserveContract() (common.Address, error) {
	return _LiquidityConversionRates.Contract.ReserveContract(&_LiquidityConversionRates.CallOpts)
}

// ReserveContract is a free data retrieval call binding the contract method 0xa7f43acd.
//
// Solidity: function reserveContract() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) ReserveContract() (common.Address, error) {
	return _LiquidityConversionRates.Contract.ReserveContract(&_LiquidityConversionRates.CallOpts)
}

// SellRate is a free data retrieval call binding the contract method 0x925176d6.
//
// Solidity: function sellRate(uint256 eInFp, uint256 sellInputTokenQtyInFp, uint256 deltaTInFp) constant returns(uint256 rateInPrecision, uint256 deltaEInFp)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) SellRate(opts *bind.CallOpts, eInFp *big.Int, sellInputTokenQtyInFp *big.Int, deltaTInFp *big.Int) (struct {
	RateInPrecision *big.Int
	DeltaEInFp      *big.Int
}, error) {
	ret := new(struct {
		RateInPrecision *big.Int
		DeltaEInFp      *big.Int
	})
	out := ret
	err := _LiquidityConversionRates.contract.Call(opts, out, "sellRate", eInFp, sellInputTokenQtyInFp, deltaTInFp)
	return *ret, err
}

// SellRate is a free data retrieval call binding the contract method 0x925176d6.
//
// Solidity: function sellRate(uint256 eInFp, uint256 sellInputTokenQtyInFp, uint256 deltaTInFp) constant returns(uint256 rateInPrecision, uint256 deltaEInFp)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) SellRate(eInFp *big.Int, sellInputTokenQtyInFp *big.Int, deltaTInFp *big.Int) (struct {
	RateInPrecision *big.Int
	DeltaEInFp      *big.Int
}, error) {
	return _LiquidityConversionRates.Contract.SellRate(&_LiquidityConversionRates.CallOpts, eInFp, sellInputTokenQtyInFp, deltaTInFp)
}

// SellRate is a free data retrieval call binding the contract method 0x925176d6.
//
// Solidity: function sellRate(uint256 eInFp, uint256 sellInputTokenQtyInFp, uint256 deltaTInFp) constant returns(uint256 rateInPrecision, uint256 deltaEInFp)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) SellRate(eInFp *big.Int, sellInputTokenQtyInFp *big.Int, deltaTInFp *big.Int) (struct {
	RateInPrecision *big.Int
	DeltaEInFp      *big.Int
}, error) {
	return _LiquidityConversionRates.Contract.SellRate(&_LiquidityConversionRates.CallOpts, eInFp, sellInputTokenQtyInFp, deltaTInFp)
}

// SellRateZeroQuantity is a free data retrieval call binding the contract method 0x625cfc46.
//
// Solidity: function sellRateZeroQuantity(uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) SellRateZeroQuantity(opts *bind.CallOpts, eInFp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "sellRateZeroQuantity", eInFp)
	return *ret0, err
}

// SellRateZeroQuantity is a free data retrieval call binding the contract method 0x625cfc46.
//
// Solidity: function sellRateZeroQuantity(uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) SellRateZeroQuantity(eInFp *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.SellRateZeroQuantity(&_LiquidityConversionRates.CallOpts, eInFp)
}

// SellRateZeroQuantity is a free data retrieval call binding the contract method 0x625cfc46.
//
// Solidity: function sellRateZeroQuantity(uint256 eInFp) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) SellRateZeroQuantity(eInFp *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.SellRateZeroQuantity(&_LiquidityConversionRates.CallOpts, eInFp)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) Token() (common.Address, error) {
	return _LiquidityConversionRates.Contract.Token(&_LiquidityConversionRates.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) Token() (common.Address, error) {
	return _LiquidityConversionRates.Contract.Token(&_LiquidityConversionRates.CallOpts)
}

// ValueAfterReducingFee is a free data retrieval call binding the contract method 0x6f3d8043.
//
// Solidity: function valueAfterReducingFee(uint256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCaller) ValueAfterReducingFee(opts *bind.CallOpts, val *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityConversionRates.contract.Call(opts, out, "valueAfterReducingFee", val)
	return *ret0, err
}

// ValueAfterReducingFee is a free data retrieval call binding the contract method 0x6f3d8043.
//
// Solidity: function valueAfterReducingFee(uint256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesSession) ValueAfterReducingFee(val *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.ValueAfterReducingFee(&_LiquidityConversionRates.CallOpts, val)
}

// ValueAfterReducingFee is a free data retrieval call binding the contract method 0x6f3d8043.
//
// Solidity: function valueAfterReducingFee(uint256 val) constant returns(uint256)
func (_LiquidityConversionRates *LiquidityConversionRatesCallerSession) ValueAfterReducingFee(val *big.Int) (*big.Int, error) {
	return _LiquidityConversionRates.Contract.ValueAfterReducingFee(&_LiquidityConversionRates.CallOpts, val)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.AddAlerter(&_LiquidityConversionRates.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.AddAlerter(&_LiquidityConversionRates.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.AddOperator(&_LiquidityConversionRates.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.AddOperator(&_LiquidityConversionRates.TransactOpts, newOperator)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) ClaimAdmin() (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.ClaimAdmin(&_LiquidityConversionRates.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.ClaimAdmin(&_LiquidityConversionRates.TransactOpts)
}

// RecordImbalance is a paid mutator transaction binding the contract method 0xc6fd2103.
//
// Solidity: function recordImbalance(address conversionToken, int256 buyAmountInTwei, uint256 rateUpdateBlock, uint256 currentBlock) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) RecordImbalance(opts *bind.TransactOpts, conversionToken common.Address, buyAmountInTwei *big.Int, rateUpdateBlock *big.Int, currentBlock *big.Int) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "recordImbalance", conversionToken, buyAmountInTwei, rateUpdateBlock, currentBlock)
}

// RecordImbalance is a paid mutator transaction binding the contract method 0xc6fd2103.
//
// Solidity: function recordImbalance(address conversionToken, int256 buyAmountInTwei, uint256 rateUpdateBlock, uint256 currentBlock) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) RecordImbalance(conversionToken common.Address, buyAmountInTwei *big.Int, rateUpdateBlock *big.Int, currentBlock *big.Int) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.RecordImbalance(&_LiquidityConversionRates.TransactOpts, conversionToken, buyAmountInTwei, rateUpdateBlock, currentBlock)
}

// RecordImbalance is a paid mutator transaction binding the contract method 0xc6fd2103.
//
// Solidity: function recordImbalance(address conversionToken, int256 buyAmountInTwei, uint256 rateUpdateBlock, uint256 currentBlock) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) RecordImbalance(conversionToken common.Address, buyAmountInTwei *big.Int, rateUpdateBlock *big.Int, currentBlock *big.Int) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.RecordImbalance(&_LiquidityConversionRates.TransactOpts, conversionToken, buyAmountInTwei, rateUpdateBlock, currentBlock)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.RemoveAlerter(&_LiquidityConversionRates.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.RemoveAlerter(&_LiquidityConversionRates.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.RemoveOperator(&_LiquidityConversionRates.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.RemoveOperator(&_LiquidityConversionRates.TransactOpts, operator)
}

// ResetCollectedFees is a paid mutator transaction binding the contract method 0xb86f6aa7.
//
// Solidity: function resetCollectedFees() returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) ResetCollectedFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "resetCollectedFees")
}

// ResetCollectedFees is a paid mutator transaction binding the contract method 0xb86f6aa7.
//
// Solidity: function resetCollectedFees() returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) ResetCollectedFees() (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.ResetCollectedFees(&_LiquidityConversionRates.TransactOpts)
}

// ResetCollectedFees is a paid mutator transaction binding the contract method 0xb86f6aa7.
//
// Solidity: function resetCollectedFees() returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) ResetCollectedFees() (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.ResetCollectedFees(&_LiquidityConversionRates.TransactOpts)
}

// SetLiquidityParams is a paid mutator transaction binding the contract method 0x4857d52d.
//
// Solidity: function setLiquidityParams(uint256 _rInFp, uint256 _pMinInFp, uint256 _numFpBits, uint256 _maxCapBuyInWei, uint256 _maxCapSellInWei, uint256 _feeInBps, uint256 _maxTokenToEthRateInPrecision, uint256 _minTokenToEthRateInPrecision) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) SetLiquidityParams(opts *bind.TransactOpts, _rInFp *big.Int, _pMinInFp *big.Int, _numFpBits *big.Int, _maxCapBuyInWei *big.Int, _maxCapSellInWei *big.Int, _feeInBps *big.Int, _maxTokenToEthRateInPrecision *big.Int, _minTokenToEthRateInPrecision *big.Int) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "setLiquidityParams", _rInFp, _pMinInFp, _numFpBits, _maxCapBuyInWei, _maxCapSellInWei, _feeInBps, _maxTokenToEthRateInPrecision, _minTokenToEthRateInPrecision)
}

// SetLiquidityParams is a paid mutator transaction binding the contract method 0x4857d52d.
//
// Solidity: function setLiquidityParams(uint256 _rInFp, uint256 _pMinInFp, uint256 _numFpBits, uint256 _maxCapBuyInWei, uint256 _maxCapSellInWei, uint256 _feeInBps, uint256 _maxTokenToEthRateInPrecision, uint256 _minTokenToEthRateInPrecision) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) SetLiquidityParams(_rInFp *big.Int, _pMinInFp *big.Int, _numFpBits *big.Int, _maxCapBuyInWei *big.Int, _maxCapSellInWei *big.Int, _feeInBps *big.Int, _maxTokenToEthRateInPrecision *big.Int, _minTokenToEthRateInPrecision *big.Int) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.SetLiquidityParams(&_LiquidityConversionRates.TransactOpts, _rInFp, _pMinInFp, _numFpBits, _maxCapBuyInWei, _maxCapSellInWei, _feeInBps, _maxTokenToEthRateInPrecision, _minTokenToEthRateInPrecision)
}

// SetLiquidityParams is a paid mutator transaction binding the contract method 0x4857d52d.
//
// Solidity: function setLiquidityParams(uint256 _rInFp, uint256 _pMinInFp, uint256 _numFpBits, uint256 _maxCapBuyInWei, uint256 _maxCapSellInWei, uint256 _feeInBps, uint256 _maxTokenToEthRateInPrecision, uint256 _minTokenToEthRateInPrecision) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) SetLiquidityParams(_rInFp *big.Int, _pMinInFp *big.Int, _numFpBits *big.Int, _maxCapBuyInWei *big.Int, _maxCapSellInWei *big.Int, _feeInBps *big.Int, _maxTokenToEthRateInPrecision *big.Int, _minTokenToEthRateInPrecision *big.Int) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.SetLiquidityParams(&_LiquidityConversionRates.TransactOpts, _rInFp, _pMinInFp, _numFpBits, _maxCapBuyInWei, _maxCapSellInWei, _feeInBps, _maxTokenToEthRateInPrecision, _minTokenToEthRateInPrecision)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(address reserve) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) SetReserveAddress(opts *bind.TransactOpts, reserve common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "setReserveAddress", reserve)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(address reserve) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) SetReserveAddress(reserve common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.SetReserveAddress(&_LiquidityConversionRates.TransactOpts, reserve)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(address reserve) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) SetReserveAddress(reserve common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.SetReserveAddress(&_LiquidityConversionRates.TransactOpts, reserve)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.TransferAdmin(&_LiquidityConversionRates.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.TransferAdmin(&_LiquidityConversionRates.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.TransferAdminQuickly(&_LiquidityConversionRates.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.TransferAdminQuickly(&_LiquidityConversionRates.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.WithdrawEther(&_LiquidityConversionRates.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.WithdrawEther(&_LiquidityConversionRates.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.WithdrawToken(&_LiquidityConversionRates.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_LiquidityConversionRates *LiquidityConversionRatesTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _LiquidityConversionRates.Contract.WithdrawToken(&_LiquidityConversionRates.TransactOpts, token, amount, sendTo)
}

// LiquidityConversionRatesAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesAdminClaimedIterator struct {
	Event *LiquidityConversionRatesAdminClaimed // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesAdminClaimed)
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
		it.Event = new(LiquidityConversionRatesAdminClaimed)
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
func (it *LiquidityConversionRatesAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesAdminClaimed represents a AdminClaimed event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*LiquidityConversionRatesAdminClaimedIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesAdminClaimedIterator{contract: _LiquidityConversionRates.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesAdminClaimed)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseAdminClaimed(log types.Log) (*LiquidityConversionRatesAdminClaimed, error) {
	event := new(LiquidityConversionRatesAdminClaimed)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityConversionRatesAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesAlerterAddedIterator struct {
	Event *LiquidityConversionRatesAlerterAdded // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesAlerterAdded)
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
		it.Event = new(LiquidityConversionRatesAlerterAdded)
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
func (it *LiquidityConversionRatesAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesAlerterAdded represents a AlerterAdded event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*LiquidityConversionRatesAlerterAddedIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesAlerterAddedIterator{contract: _LiquidityConversionRates.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesAlerterAdded)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseAlerterAdded(log types.Log) (*LiquidityConversionRatesAlerterAdded, error) {
	event := new(LiquidityConversionRatesAlerterAdded)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityConversionRatesCollectedFeesResetIterator is returned from FilterCollectedFeesReset and is used to iterate over the raw logs and unpacked data for CollectedFeesReset events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesCollectedFeesResetIterator struct {
	Event *LiquidityConversionRatesCollectedFeesReset // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesCollectedFeesResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesCollectedFeesReset)
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
		it.Event = new(LiquidityConversionRatesCollectedFeesReset)
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
func (it *LiquidityConversionRatesCollectedFeesResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesCollectedFeesResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesCollectedFeesReset represents a CollectedFeesReset event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesCollectedFeesReset struct {
	ResetFeesInTwei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollectedFeesReset is a free log retrieval operation binding the contract event 0xdeb4766cf1de6f18e3b195f199d403a02a3e09fbee1192b37d797fb300f05261.
//
// Solidity: event CollectedFeesReset(uint256 resetFeesInTwei)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterCollectedFeesReset(opts *bind.FilterOpts) (*LiquidityConversionRatesCollectedFeesResetIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "CollectedFeesReset")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesCollectedFeesResetIterator{contract: _LiquidityConversionRates.contract, event: "CollectedFeesReset", logs: logs, sub: sub}, nil
}

// WatchCollectedFeesReset is a free log subscription operation binding the contract event 0xdeb4766cf1de6f18e3b195f199d403a02a3e09fbee1192b37d797fb300f05261.
//
// Solidity: event CollectedFeesReset(uint256 resetFeesInTwei)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchCollectedFeesReset(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesCollectedFeesReset) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "CollectedFeesReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesCollectedFeesReset)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "CollectedFeesReset", log); err != nil {
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

// ParseCollectedFeesReset is a log parse operation binding the contract event 0xdeb4766cf1de6f18e3b195f199d403a02a3e09fbee1192b37d797fb300f05261.
//
// Solidity: event CollectedFeesReset(uint256 resetFeesInTwei)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseCollectedFeesReset(log types.Log) (*LiquidityConversionRatesCollectedFeesReset, error) {
	event := new(LiquidityConversionRatesCollectedFeesReset)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "CollectedFeesReset", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityConversionRatesEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesEtherWithdrawIterator struct {
	Event *LiquidityConversionRatesEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesEtherWithdraw)
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
		it.Event = new(LiquidityConversionRatesEtherWithdraw)
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
func (it *LiquidityConversionRatesEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesEtherWithdraw represents a EtherWithdraw event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*LiquidityConversionRatesEtherWithdrawIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesEtherWithdrawIterator{contract: _LiquidityConversionRates.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesEtherWithdraw)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseEtherWithdraw(log types.Log) (*LiquidityConversionRatesEtherWithdraw, error) {
	event := new(LiquidityConversionRatesEtherWithdraw)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityConversionRatesLiquidityParamsSetIterator is returned from FilterLiquidityParamsSet and is used to iterate over the raw logs and unpacked data for LiquidityParamsSet events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesLiquidityParamsSetIterator struct {
	Event *LiquidityConversionRatesLiquidityParamsSet // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesLiquidityParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesLiquidityParamsSet)
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
		it.Event = new(LiquidityConversionRatesLiquidityParamsSet)
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
func (it *LiquidityConversionRatesLiquidityParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesLiquidityParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesLiquidityParamsSet represents a LiquidityParamsSet event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesLiquidityParamsSet struct {
	RInFp                  *big.Int
	PMinInFp               *big.Int
	NumFpBits              *big.Int
	MaxCapBuyInFp          *big.Int
	MaxEthCapSellInFp      *big.Int
	FeeInBps               *big.Int
	FormulaPrecision       *big.Int
	MaxQtyInFp             *big.Int
	MaxBuyRateInPrecision  *big.Int
	MinBuyRateInPrecision  *big.Int
	MaxSellRateInPrecision *big.Int
	MinSellRateInPrecision *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterLiquidityParamsSet is a free log retrieval operation binding the contract event 0x52db0a06d138736a4425764a1f7e1b432b5ce79099d523c0c4cd01e7320aba0e.
//
// Solidity: event LiquidityParamsSet(uint256 rInFp, uint256 pMinInFp, uint256 numFpBits, uint256 maxCapBuyInFp, uint256 maxEthCapSellInFp, uint256 feeInBps, uint256 formulaPrecision, uint256 maxQtyInFp, uint256 maxBuyRateInPrecision, uint256 minBuyRateInPrecision, uint256 maxSellRateInPrecision, uint256 minSellRateInPrecision)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterLiquidityParamsSet(opts *bind.FilterOpts) (*LiquidityConversionRatesLiquidityParamsSetIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "LiquidityParamsSet")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesLiquidityParamsSetIterator{contract: _LiquidityConversionRates.contract, event: "LiquidityParamsSet", logs: logs, sub: sub}, nil
}

// WatchLiquidityParamsSet is a free log subscription operation binding the contract event 0x52db0a06d138736a4425764a1f7e1b432b5ce79099d523c0c4cd01e7320aba0e.
//
// Solidity: event LiquidityParamsSet(uint256 rInFp, uint256 pMinInFp, uint256 numFpBits, uint256 maxCapBuyInFp, uint256 maxEthCapSellInFp, uint256 feeInBps, uint256 formulaPrecision, uint256 maxQtyInFp, uint256 maxBuyRateInPrecision, uint256 minBuyRateInPrecision, uint256 maxSellRateInPrecision, uint256 minSellRateInPrecision)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchLiquidityParamsSet(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesLiquidityParamsSet) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "LiquidityParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesLiquidityParamsSet)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "LiquidityParamsSet", log); err != nil {
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

// ParseLiquidityParamsSet is a log parse operation binding the contract event 0x52db0a06d138736a4425764a1f7e1b432b5ce79099d523c0c4cd01e7320aba0e.
//
// Solidity: event LiquidityParamsSet(uint256 rInFp, uint256 pMinInFp, uint256 numFpBits, uint256 maxCapBuyInFp, uint256 maxEthCapSellInFp, uint256 feeInBps, uint256 formulaPrecision, uint256 maxQtyInFp, uint256 maxBuyRateInPrecision, uint256 minBuyRateInPrecision, uint256 maxSellRateInPrecision, uint256 minSellRateInPrecision)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseLiquidityParamsSet(log types.Log) (*LiquidityConversionRatesLiquidityParamsSet, error) {
	event := new(LiquidityConversionRatesLiquidityParamsSet)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "LiquidityParamsSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityConversionRatesOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesOperatorAddedIterator struct {
	Event *LiquidityConversionRatesOperatorAdded // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesOperatorAdded)
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
		it.Event = new(LiquidityConversionRatesOperatorAdded)
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
func (it *LiquidityConversionRatesOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesOperatorAdded represents a OperatorAdded event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*LiquidityConversionRatesOperatorAddedIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesOperatorAddedIterator{contract: _LiquidityConversionRates.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesOperatorAdded)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseOperatorAdded(log types.Log) (*LiquidityConversionRatesOperatorAdded, error) {
	event := new(LiquidityConversionRatesOperatorAdded)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityConversionRatesReserveAddressSetIterator is returned from FilterReserveAddressSet and is used to iterate over the raw logs and unpacked data for ReserveAddressSet events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesReserveAddressSetIterator struct {
	Event *LiquidityConversionRatesReserveAddressSet // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesReserveAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesReserveAddressSet)
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
		it.Event = new(LiquidityConversionRatesReserveAddressSet)
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
func (it *LiquidityConversionRatesReserveAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesReserveAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesReserveAddressSet represents a ReserveAddressSet event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesReserveAddressSet struct {
	Reserve common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveAddressSet is a free log retrieval operation binding the contract event 0xbd2ca09dd2b354751631db75d1a63231ec123c0d68c81928ea03d0be326c7f88.
//
// Solidity: event ReserveAddressSet(address reserve)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterReserveAddressSet(opts *bind.FilterOpts) (*LiquidityConversionRatesReserveAddressSetIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "ReserveAddressSet")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesReserveAddressSetIterator{contract: _LiquidityConversionRates.contract, event: "ReserveAddressSet", logs: logs, sub: sub}, nil
}

// WatchReserveAddressSet is a free log subscription operation binding the contract event 0xbd2ca09dd2b354751631db75d1a63231ec123c0d68c81928ea03d0be326c7f88.
//
// Solidity: event ReserveAddressSet(address reserve)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchReserveAddressSet(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesReserveAddressSet) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "ReserveAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesReserveAddressSet)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "ReserveAddressSet", log); err != nil {
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

// ParseReserveAddressSet is a log parse operation binding the contract event 0xbd2ca09dd2b354751631db75d1a63231ec123c0d68c81928ea03d0be326c7f88.
//
// Solidity: event ReserveAddressSet(address reserve)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseReserveAddressSet(log types.Log) (*LiquidityConversionRatesReserveAddressSet, error) {
	event := new(LiquidityConversionRatesReserveAddressSet)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "ReserveAddressSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityConversionRatesTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesTokenWithdrawIterator struct {
	Event *LiquidityConversionRatesTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesTokenWithdraw)
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
		it.Event = new(LiquidityConversionRatesTokenWithdraw)
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
func (it *LiquidityConversionRatesTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesTokenWithdraw represents a TokenWithdraw event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*LiquidityConversionRatesTokenWithdrawIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesTokenWithdrawIterator{contract: _LiquidityConversionRates.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesTokenWithdraw)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseTokenWithdraw(log types.Log) (*LiquidityConversionRatesTokenWithdraw, error) {
	event := new(LiquidityConversionRatesTokenWithdraw)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityConversionRatesTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesTransferAdminPendingIterator struct {
	Event *LiquidityConversionRatesTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *LiquidityConversionRatesTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityConversionRatesTransferAdminPending)
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
		it.Event = new(LiquidityConversionRatesTransferAdminPending)
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
func (it *LiquidityConversionRatesTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityConversionRatesTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityConversionRatesTransferAdminPending represents a TransferAdminPending event raised by the LiquidityConversionRates contract.
type LiquidityConversionRatesTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*LiquidityConversionRatesTransferAdminPendingIterator, error) {

	logs, sub, err := _LiquidityConversionRates.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &LiquidityConversionRatesTransferAdminPendingIterator{contract: _LiquidityConversionRates.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *LiquidityConversionRatesTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _LiquidityConversionRates.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityConversionRatesTransferAdminPending)
				if err := _LiquidityConversionRates.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_LiquidityConversionRates *LiquidityConversionRatesFilterer) ParseTransferAdminPending(log types.Log) (*LiquidityConversionRatesTransferAdminPending, error) {
	event := new(LiquidityConversionRatesTransferAdminPending)
	if err := _LiquidityConversionRates.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}
