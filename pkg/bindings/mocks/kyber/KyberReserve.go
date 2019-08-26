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

// KyberReserveABI is the input ABI used to generate the binding from.
const KyberReserveABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"enableTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"setTokenWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sanityRatesContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"approve\",\"type\":\"bool\"}],\"name\":\"approveWithdrawAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"srcToken\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"destToken\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"conversionRate\",\"type\":\"uint256\"},{\"name\":\"validate\",\"type\":\"bool\"}],\"name\":\"trade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"srcQty\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getConversionRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"dstQty\",\"type\":\"uint256\"},{\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"getSrcQty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kyberNetwork\",\"type\":\"address\"},{\"name\":\"_conversionRates\",\"type\":\"address\"},{\"name\":\"_sanityRates\",\"type\":\"address\"}],\"name\":\"setContracts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetwork\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionRatesContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tradeEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedWithdrawAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"srcQty\",\"type\":\"uint256\"},{\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"getDestQty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kyberNetwork\",\"type\":\"address\"},{\"name\":\"_ratesContract\",\"type\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"origin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"destAddress\",\"type\":\"address\"}],\"name\":\"TradeExecute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"TradeEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approve\",\"type\":\"bool\"}],\"name\":\"WithdrawAddressApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"NewTokenWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"WithdrawFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"network\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"rate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sanity\",\"type\":\"address\"}],\"name\":\"SetContractAddresses\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"}]"

// KyberReserveBin is the compiled bytecode used for deploying new contracts.
var KyberReserveBin = "0x6060604052341561000f57600080fd5b60405160608061219083398101604052808051919060200180519190602001805160008054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061006357600080fd5b600160a060020a038216151561007857600080fd5b600160a060020a038316151561008d57600080fd5b6007805460088054600160a060020a03958616600160a060020a031991821617909155600080549486169482169490941790935560a060020a60ff02199390941691909316171674010000000000000000000000000000000000000000179055612094806100fc6000396000f3006060604052600436106101655763ffffffff60e060020a60003504166299d38681146101b957806301a12fd3146101e05780631bc7bfec14610201578063267822471461022657806327a099d8146102555780633ccdbb28146102bb578063408ee7fe146102e457806347e6924f14610303578063546dc71c1461031657806369328dec146103405780636940030f146103695780636cf698111461037c57806375829def146103a857806377f50f97146103c75780637acc8678146103da5780637c423f54146103f95780637cd442721461040c5780639870d7fe14610449578063a7fca95314610468578063a80cbac614610493578063ac8a584a146104b2578063b3066d49146104d1578063b78b842d146104fc578063ce56c4541461050f578063d5847d3314610531578063d621e81314610544578063d7b7024d14610557578063f851a4401461056d578063f8b2cb4f14610580578063fa64dffa1461059f575b7f2d0c0a8842b9944ece1495eb61121621b5e36bd6af3bba0318c695f525aef79f60008051602061204983398151915234604051600160a060020a03909216825260208201526040908101905180910390a1005b34156101c457600080fd5b6101cc6105ca565b604051901515815260200160405180910390f35b34156101eb57600080fd5b6101ff600160a060020a0360043516610658565b005b341561020c57600080fd5b6101ff600160a060020a03600435811690602435166107c8565b341561023157600080fd5b610239610874565b604051600160a060020a03909116815260200160405180910390f35b341561026057600080fd5b610268610883565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102a757808201518382015260200161028f565b505050509050019250505060405180910390f35b34156102c657600080fd5b6101ff600160a060020a0360043581169060243590604435166108eb565b34156102ef57600080fd5b6101ff600160a060020a03600435166109e2565b341561030e57600080fd5b610239610ade565b341561032157600080fd5b6101ff600160a060020a03600435811690602435166044351515610aed565b341561034b57600080fd5b6101cc600160a060020a036004358116906024359060443516610cd7565b341561037457600080fd5b6101cc610eb0565b6101cc600160a060020a03600435811690602435906044358116906064351660843560a4351515610f31565b34156103b357600080fd5b6101ff600160a060020a0360043516610f9e565b34156103d257600080fd5b6101ff611039565b34156103e557600080fd5b6101ff600160a060020a03600435166110d3565b341561040457600080fd5b6102686111b5565b341561041757600080fd5b610437600160a060020a036004358116906024351660443560643561121b565b60405190815260200160405180910390f35b341561045457600080fd5b6101ff600160a060020a0360043516611406565b341561047357600080fd5b610437600160a060020a03600435811690602435166044356064356114d6565b341561049e57600080fd5b610239600160a060020a0360043516611508565b34156104bd57600080fd5b6101ff600160a060020a0360043516611523565b34156104dc57600080fd5b6101ff600160a060020a036004358116906024358116906044351661168f565b341561050757600080fd5b610239611773565b341561051a57600080fd5b6101ff600435600160a060020a0360243516611782565b341561053c57600080fd5b610239611815565b341561054f57600080fd5b6101cc611824565b341561056257600080fd5b6101cc600435611845565b341561057857600080fd5b61023961185a565b341561058b57600080fd5b610437600160a060020a0360043516611869565b34156105aa57600080fd5b610437600160a060020a03600435811690602435166044356064356119bd565b6000805433600160a060020a039081169116146105e657600080fd5b6007805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557f7d7f00509dd73ac4449f698ae75ccc797895eff5fa9d446d3df387598a26e7356001604051901515815260200160405180910390a15060015b90565b6000805433600160a060020a0390811691161461067457600080fd5b600160a060020a03821660009081526003602052604090205460ff16151561069b57600080fd5b50600160a060020a0381166000908152600360205260408120805460ff191690555b6005548110156107c45781600160a060020a03166005828154811015156106e057fe5b600091825260209091200154600160a060020a031614156107bc5760058054600019810190811061070d57fe5b60009182526020909120015460058054600160a060020a03909216918390811061073357fe5b60009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055600580549061076f906000198301611ff4565b507f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762826000604051600160a060020a039092168252151560208201526040908101905180910390a16107c4565b6001016106bd565b5050565b60005433600160a060020a039081169116146107e357600080fd5b600160a060020a03811615156107f857600080fd5b600160a060020a038281166000908152600b6020526040908190208054600160a060020a031916928416929092179091557f81995c7b922889ac0a81e41866106d4046268ea3a9abaae9f9e080a6ce36ee7d908390839051600160a060020a039283168152911660208201526040908101905180910390a15050565b600154600160a060020a031681565b61088b612018565b60048054806020026020016040519081016040528092919081815260200182805480156108e157602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116108c3575b5050505050905090565b60005433600160a060020a0390811691161461090657600080fd5b82600160a060020a031663a9059cbb828460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561096357600080fd5b6102c65a03f1151561097457600080fd5b50505060405180519050151561098957600080fd5b7f72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6838383604051600160a060020a03938416815260208101929092529091166040808301919091526060909101905180910390a1505050565b60005433600160a060020a039081169116146109fd57600080fd5b600160a060020a03811660009081526003602052604090205460ff1615610a2357600080fd5b60055460329010610a3357600080fd5b7f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600360205260409020805460ff191660019081179091556005805490918101610ab28382611ff4565b5060009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055565b600954600160a060020a031681565b60005433600160a060020a03908116911614610b0857600080fd5b80600a600085856040516c01000000000000000000000000600160a060020a039384168102825291909216026014820152602801604051908190039020815260208101919091526040908101600020805460ff1916921515929092179091557fd5fd5351efae1f4bb760079da9f0ff9589e2c3e216337ca9d39cdff573b245c49084908490849051600160a060020a0393841681529190921660208201529015156040808301919091526060909101905180910390a1610bc7836119e4565b600160a060020a038381166000908152600b602052604090205416158015610c065750600160a060020a03831660008051602061204983398151915214155b15610cd257600160a060020a038381166000818152600b60205260408082208054600160a060020a03191630958616179055919263095ea7b39290917f800000000000000000000000000000000000000000000000000000000000000091516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610cac57600080fd5b6102c65a03f11515610cbd57600080fd5b505050604051805190501515610cd257600080fd5b505050565b600160a060020a03331660009081526002602052604081205460ff161515610cfe57600080fd5b600a600085846040516c01000000000000000000000000600160a060020a039384168102825291909216026014820152602801604051908190039020815260208101919091526040016000205460ff161515610d5957600080fd5b600160a060020a0384166000805160206120498339815191521415610dae57600160a060020a03821683156108fc0284604051600060405180830381858888f193505050501515610da957600080fd5b610e52565b600160a060020a038085166000818152600b60205260408082205492936323b872dd9316918691889190516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610e2c57600080fd5b6102c65a03f11515610e3d57600080fd5b505050604051805190501515610e5257600080fd5b7fb67719fc33c1f17d31bf3a698690d62066b1e0bae28fcd3c56cf2c015c2863d6848484604051600160a060020a03938416815260208101929092529091166040808301919091526060909101905180910390a15060019392505050565b600160a060020a03331660009081526003602052604081205460ff161515610ed757600080fd5b6007805474ff0000000000000000000000000000000000000000191690557f7d7f00509dd73ac4449f698ae75ccc797895eff5fa9d446d3df387598a26e7356000604051901515815260200160405180910390a150600190565b60075460009074010000000000000000000000000000000000000000900460ff161515610f5d57600080fd5b60075433600160a060020a03908116911614610f7857600080fd5b610f86878787878787611aa1565b1515610f9157600080fd5b5060019695505050505050565b60005433600160a060020a03908116911614610fb957600080fd5b600160a060020a0381161515610fce57600080fd5b6001547f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4090600160a060020a0316604051600160a060020a03909116815260200160405180910390a160018054600160a060020a031916600160a060020a0392909216919091179055565b60015433600160a060020a0390811691161461105457600080fd5b6001546000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed91600160a060020a039081169116604051600160a060020a039283168152911660208201526040908101905180910390a16001805460008054600160a060020a0319908116600160a060020a03841617909155169055565b60005433600160a060020a039081169116146110ee57600080fd5b600160a060020a038116151561110357600080fd5b7f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4081604051600160a060020a03909116815260200160405180910390a16000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed908290600160a060020a0316604051600160a060020a039283168152911660208201526040908101905180910390a160008054600160a060020a031916600160a060020a0392909216919091179055565b6111bd612018565b60058054806020026020016040519081016040528092919081815260200182805480156108e157602002820191906000526020600020908154600160a060020a031681526001909101906020018083116108c3575050505050905090565b600080600080600080600760149054906101000a900460ff16151561124357600095506113f9565b600080516020612049833981519152600160a060020a038b16141561126e57600193508894506112a2565b600080516020612049833981519152600160a060020a038a16141561129957600093508994506112a2565b600095506113f9565b600854600160a060020a031663b8e9c22e8689878c60006040516020015260405160e060020a63ffffffff8716028152600160a060020a0390941660048501526024840192909252151560448301526064820152608401602060405180830381600087803b151561131257600080fd5b6102c65a03f1151561132357600080fd5b50505060405180519050925061133b8a8a8a866119bd565b9150816113478a611869565b101561135657600095506113f9565b600954600160a060020a0316156113f557600954600160a060020a031663a58092b78b8b60006040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b15156113c857600080fd5b6102c65a03f115156113d957600080fd5b5050506040518051915050808311156113f557600095506113f9565b8295505b5050505050949350505050565b60005433600160a060020a0390811691161461142157600080fd5b600160a060020a03811660009081526002602052604090205460ff161561144757600080fd5b6004546032901061145757600080fd5b7f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600260205260409020805460ff191660019081179091556004805490918101610ab28382611ff4565b60008060006114e486611df6565b91506114ef87611df6565b90506114fd85828487611eb4565b979650505050505050565b600b60205260009081526040902054600160a060020a031681565b6000805433600160a060020a0390811691161461153f57600080fd5b600160a060020a03821660009081526002602052604090205460ff16151561156657600080fd5b50600160a060020a0381166000908152600260205260408120805460ff191690555b6004548110156107c45781600160a060020a03166004828154811015156115ab57fe5b600091825260209091200154600160a060020a03161415611687576004805460001981019081106115d857fe5b60009182526020909120015460048054600160a060020a0390921691839081106115fe57fe5b60009182526020909120018054600160a060020a031916600160a060020a039290921691909117905560048054600019019061163a9082611ff4565b507f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b826000604051600160a060020a039092168252151560208201526040908101905180910390a16107c4565b600101611588565b60005433600160a060020a039081169116146116aa57600080fd5b600160a060020a03831615156116bf57600080fd5b600160a060020a03821615156116d457600080fd5b60078054600160a060020a03808616600160a060020a0319928316179283905560088054868316908416179081905560098054868416941693909317928390557f7a85322644a4462d8ff5482d2a841a4d231f8cfb3c9f4a50f66f8b2bd568c31c938216929082169116604051600160a060020a03938416815291831660208301529091166040808301919091526060909101905180910390a1505050565b600754600160a060020a031681565b60005433600160a060020a0390811691161461179d57600080fd5b600160a060020a03811682156108fc0283604051600060405180830381858888f1935050505015156117ce57600080fd5b7fec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de8282604051918252600160a060020a031660208201526040908101905180910390a15050565b600854600160a060020a031681565b60075474010000000000000000000000000000000000000000900460ff1681565b600a6020526000908152604090205460ff1681565b600054600160a060020a031681565b6000808080600160a060020a038516600080516020612049833981519152141561189f5730600160a060020a03163193506119b5565b600160a060020a038086166000818152600b602052604080822054909316955090916370a08231918691516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561190c57600080fd5b6102c65a03f1151561191d57600080fd5b5050506040518051925050600160a060020a03851663dd62ed3e843060006040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561198757600080fd5b6102c65a03f1151561199857600080fd5b50505060405180519150508082106119b057806119b2565b815b93505b505050919050565b60008060006119cb86611df6565b91506119d687611df6565b90506114fd85828487611f5b565b600160a060020a0381166000805160206120498339815191521415611a2457600160a060020a038116600090815260066020526040902060129055611a9e565b80600160a060020a031663313ce5676000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611a6a57600080fd5b6102c65a03f11515611a7b57600080fd5b5050506040518051600160a060020a038316600090815260066020526040902055505b50565b6000806000808415611af55760008611611aba57600080fd5b600160a060020a038a166000805160206120498339815191521415611aea57348914611ae557600080fd5b611af5565b3415611af557600080fd5b611b018a898b896119bd565b925060008311611b1057600080fd5b600160a060020a038a166000805160206120498339815191521415611b39575086905081611b43565b5088905060001988025b600854600160a060020a031663c6fd2103838360004360405160e060020a63ffffffff8716028152600160a060020a039094166004850152602484019290925260448301526064820152608401600060405180830381600087803b1515611ba957600080fd5b6102c65a03f11515611bba57600080fd5b505050600160a060020a038a1660008051602061204983398151915214611c8057600160a060020a03808b166000818152600b60205260408082205492936323b872dd9333939116918e9190516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515611c5a57600080fd5b6102c65a03f11515611c6b57600080fd5b505050604051805190501515611c8057600080fd5b600160a060020a0388166000805160206120498339815191521415611cd557600160a060020a03871683156108fc0284604051600060405180830381858888f193505050501515611cd057600080fd5b611d79565b600160a060020a038089166000818152600b60205260408082205492936323b872dd9316918b91889190516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515611d5357600080fd5b6102c65a03f11515611d6457600080fd5b505050604051805190501515611d7957600080fd5b33600160a060020a03167fea9415385bae08fe9f6dc457b02577166790cde83bb18cc340aac6cb81b824de8b8b8b878c604051600160a060020a039586168152602081019490945291841660408085019190915260608401919091529216608082015260a001905180910390a25060019998505050505050505050565b600080600160a060020a0383166000805160206120498339815191521415611e215760129150611eae565b50600160a060020a038216600090815260066020526040902054801515611eaa5782600160a060020a031663313ce5676000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611e8857600080fd5b6102c65a03f11515611e9957600080fd5b505050604051805190509150611eae565b8091505b50919050565b600080806b204fce5e3e25026110000000871115611ed157600080fd5b69d3c21bcecceda1000000841115611ee857600080fd5b848610611f195760128587031115611eff57600080fd5b5050828403600a0a8502670de0b6b3a76400000282611f3f565b60128686031115611f2957600080fd5b5050670de0b6b3a76400008502848403600a0a83025b80600182840103811515611f4f57fe5b04979650505050505050565b60006b204fce5e3e25026110000000851115611f7657600080fd5b69d3c21bcecceda1000000821115611f8d57600080fd5b838310611fc05760128484031115611fa457600080fd5b670de0b6b3a7640000858302858503600a0a025b049050611fec565b60128385031115611fd057600080fd5b828403600a0a670de0b6b3a764000002828602811515611fb857fe5b949350505050565b815481835581811511610cd257600083815260209020610cd291810190830161202a565b60206040519081016040526000815290565b61065591905b808211156120445760008155600101612030565b50905600000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeea165627a7a723058204ae4d3f275b3a615576f40f28d26e76b684acf4a3eec80e401d199a67808667b0029"

// DeployKyberReserve deploys a new Ethereum contract, binding an instance of KyberReserve to it.
func DeployKyberReserve(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetwork common.Address, _ratesContract common.Address, _admin common.Address) (common.Address, *types.Transaction, *KyberReserve, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberReserveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KyberReserveBin), backend, _kyberNetwork, _ratesContract, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KyberReserve{KyberReserveCaller: KyberReserveCaller{contract: contract}, KyberReserveTransactor: KyberReserveTransactor{contract: contract}, KyberReserveFilterer: KyberReserveFilterer{contract: contract}}, nil
}

// KyberReserve is an auto generated Go binding around an Ethereum contract.
type KyberReserve struct {
	KyberReserveCaller     // Read-only binding to the contract
	KyberReserveTransactor // Write-only binding to the contract
	KyberReserveFilterer   // Log filterer for contract events
}

// KyberReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberReserveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberReserveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberReserveSession struct {
	Contract     *KyberReserve     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberReserveCallerSession struct {
	Contract *KyberReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KyberReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberReserveTransactorSession struct {
	Contract     *KyberReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KyberReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberReserveRaw struct {
	Contract *KyberReserve // Generic contract binding to access the raw methods on
}

// KyberReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberReserveCallerRaw struct {
	Contract *KyberReserveCaller // Generic read-only contract binding to access the raw methods on
}

// KyberReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberReserveTransactorRaw struct {
	Contract *KyberReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberReserve creates a new instance of KyberReserve, bound to a specific deployed contract.
func NewKyberReserve(address common.Address, backend bind.ContractBackend) (*KyberReserve, error) {
	contract, err := bindKyberReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberReserve{KyberReserveCaller: KyberReserveCaller{contract: contract}, KyberReserveTransactor: KyberReserveTransactor{contract: contract}, KyberReserveFilterer: KyberReserveFilterer{contract: contract}}, nil
}

// NewKyberReserveCaller creates a new read-only instance of KyberReserve, bound to a specific deployed contract.
func NewKyberReserveCaller(address common.Address, caller bind.ContractCaller) (*KyberReserveCaller, error) {
	contract, err := bindKyberReserve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberReserveCaller{contract: contract}, nil
}

// NewKyberReserveTransactor creates a new write-only instance of KyberReserve, bound to a specific deployed contract.
func NewKyberReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberReserveTransactor, error) {
	contract, err := bindKyberReserve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberReserveTransactor{contract: contract}, nil
}

// NewKyberReserveFilterer creates a new log filterer instance of KyberReserve, bound to a specific deployed contract.
func NewKyberReserveFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberReserveFilterer, error) {
	contract, err := bindKyberReserve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberReserveFilterer{contract: contract}, nil
}

// bindKyberReserve binds a generic wrapper to an already deployed contract.
func bindKyberReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberReserveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberReserve *KyberReserveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberReserve.Contract.KyberReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberReserve *KyberReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberReserve.Contract.KyberReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberReserve *KyberReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberReserve.Contract.KyberReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberReserve *KyberReserveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberReserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberReserve *KyberReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberReserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberReserve *KyberReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberReserve.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberReserve *KyberReserveCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberReserve *KyberReserveSession) Admin() (common.Address, error) {
	return _KyberReserve.Contract.Admin(&_KyberReserve.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberReserve *KyberReserveCallerSession) Admin() (common.Address, error) {
	return _KyberReserve.Contract.Admin(&_KyberReserve.CallOpts)
}

// ApprovedWithdrawAddresses is a free data retrieval call binding the contract method 0xd7b7024d.
//
// Solidity: function approvedWithdrawAddresses(bytes32 ) constant returns(bool)
func (_KyberReserve *KyberReserveCaller) ApprovedWithdrawAddresses(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "approvedWithdrawAddresses", arg0)
	return *ret0, err
}

// ApprovedWithdrawAddresses is a free data retrieval call binding the contract method 0xd7b7024d.
//
// Solidity: function approvedWithdrawAddresses(bytes32 ) constant returns(bool)
func (_KyberReserve *KyberReserveSession) ApprovedWithdrawAddresses(arg0 [32]byte) (bool, error) {
	return _KyberReserve.Contract.ApprovedWithdrawAddresses(&_KyberReserve.CallOpts, arg0)
}

// ApprovedWithdrawAddresses is a free data retrieval call binding the contract method 0xd7b7024d.
//
// Solidity: function approvedWithdrawAddresses(bytes32 ) constant returns(bool)
func (_KyberReserve *KyberReserveCallerSession) ApprovedWithdrawAddresses(arg0 [32]byte) (bool, error) {
	return _KyberReserve.Contract.ApprovedWithdrawAddresses(&_KyberReserve.CallOpts, arg0)
}

// ConversionRatesContract is a free data retrieval call binding the contract method 0xd5847d33.
//
// Solidity: function conversionRatesContract() constant returns(address)
func (_KyberReserve *KyberReserveCaller) ConversionRatesContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "conversionRatesContract")
	return *ret0, err
}

// ConversionRatesContract is a free data retrieval call binding the contract method 0xd5847d33.
//
// Solidity: function conversionRatesContract() constant returns(address)
func (_KyberReserve *KyberReserveSession) ConversionRatesContract() (common.Address, error) {
	return _KyberReserve.Contract.ConversionRatesContract(&_KyberReserve.CallOpts)
}

// ConversionRatesContract is a free data retrieval call binding the contract method 0xd5847d33.
//
// Solidity: function conversionRatesContract() constant returns(address)
func (_KyberReserve *KyberReserveCallerSession) ConversionRatesContract() (common.Address, error) {
	return _KyberReserve.Contract.ConversionRatesContract(&_KyberReserve.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberReserve *KyberReserveCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberReserve *KyberReserveSession) GetAlerters() ([]common.Address, error) {
	return _KyberReserve.Contract.GetAlerters(&_KyberReserve.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberReserve *KyberReserveCallerSession) GetAlerters() ([]common.Address, error) {
	return _KyberReserve.Contract.GetAlerters(&_KyberReserve.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) constant returns(uint256)
func (_KyberReserve *KyberReserveCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "getBalance", token)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) constant returns(uint256)
func (_KyberReserve *KyberReserveSession) GetBalance(token common.Address) (*big.Int, error) {
	return _KyberReserve.Contract.GetBalance(&_KyberReserve.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) constant returns(uint256)
func (_KyberReserve *KyberReserveCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _KyberReserve.Contract.GetBalance(&_KyberReserve.CallOpts, token)
}

// GetConversionRate is a free data retrieval call binding the contract method 0x7cd44272.
//
// Solidity: function getConversionRate(address src, address dest, uint256 srcQty, uint256 blockNumber) constant returns(uint256)
func (_KyberReserve *KyberReserveCaller) GetConversionRate(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "getConversionRate", src, dest, srcQty, blockNumber)
	return *ret0, err
}

// GetConversionRate is a free data retrieval call binding the contract method 0x7cd44272.
//
// Solidity: function getConversionRate(address src, address dest, uint256 srcQty, uint256 blockNumber) constant returns(uint256)
func (_KyberReserve *KyberReserveSession) GetConversionRate(src common.Address, dest common.Address, srcQty *big.Int, blockNumber *big.Int) (*big.Int, error) {
	return _KyberReserve.Contract.GetConversionRate(&_KyberReserve.CallOpts, src, dest, srcQty, blockNumber)
}

// GetConversionRate is a free data retrieval call binding the contract method 0x7cd44272.
//
// Solidity: function getConversionRate(address src, address dest, uint256 srcQty, uint256 blockNumber) constant returns(uint256)
func (_KyberReserve *KyberReserveCallerSession) GetConversionRate(src common.Address, dest common.Address, srcQty *big.Int, blockNumber *big.Int) (*big.Int, error) {
	return _KyberReserve.Contract.GetConversionRate(&_KyberReserve.CallOpts, src, dest, srcQty, blockNumber)
}

// GetDestQty is a free data retrieval call binding the contract method 0xfa64dffa.
//
// Solidity: function getDestQty(address src, address dest, uint256 srcQty, uint256 rate) constant returns(uint256)
func (_KyberReserve *KyberReserveCaller) GetDestQty(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int, rate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "getDestQty", src, dest, srcQty, rate)
	return *ret0, err
}

// GetDestQty is a free data retrieval call binding the contract method 0xfa64dffa.
//
// Solidity: function getDestQty(address src, address dest, uint256 srcQty, uint256 rate) constant returns(uint256)
func (_KyberReserve *KyberReserveSession) GetDestQty(src common.Address, dest common.Address, srcQty *big.Int, rate *big.Int) (*big.Int, error) {
	return _KyberReserve.Contract.GetDestQty(&_KyberReserve.CallOpts, src, dest, srcQty, rate)
}

// GetDestQty is a free data retrieval call binding the contract method 0xfa64dffa.
//
// Solidity: function getDestQty(address src, address dest, uint256 srcQty, uint256 rate) constant returns(uint256)
func (_KyberReserve *KyberReserveCallerSession) GetDestQty(src common.Address, dest common.Address, srcQty *big.Int, rate *big.Int) (*big.Int, error) {
	return _KyberReserve.Contract.GetDestQty(&_KyberReserve.CallOpts, src, dest, srcQty, rate)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberReserve *KyberReserveCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberReserve *KyberReserveSession) GetOperators() ([]common.Address, error) {
	return _KyberReserve.Contract.GetOperators(&_KyberReserve.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberReserve *KyberReserveCallerSession) GetOperators() ([]common.Address, error) {
	return _KyberReserve.Contract.GetOperators(&_KyberReserve.CallOpts)
}

// GetSrcQty is a free data retrieval call binding the contract method 0xa7fca953.
//
// Solidity: function getSrcQty(address src, address dest, uint256 dstQty, uint256 rate) constant returns(uint256)
func (_KyberReserve *KyberReserveCaller) GetSrcQty(opts *bind.CallOpts, src common.Address, dest common.Address, dstQty *big.Int, rate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "getSrcQty", src, dest, dstQty, rate)
	return *ret0, err
}

// GetSrcQty is a free data retrieval call binding the contract method 0xa7fca953.
//
// Solidity: function getSrcQty(address src, address dest, uint256 dstQty, uint256 rate) constant returns(uint256)
func (_KyberReserve *KyberReserveSession) GetSrcQty(src common.Address, dest common.Address, dstQty *big.Int, rate *big.Int) (*big.Int, error) {
	return _KyberReserve.Contract.GetSrcQty(&_KyberReserve.CallOpts, src, dest, dstQty, rate)
}

// GetSrcQty is a free data retrieval call binding the contract method 0xa7fca953.
//
// Solidity: function getSrcQty(address src, address dest, uint256 dstQty, uint256 rate) constant returns(uint256)
func (_KyberReserve *KyberReserveCallerSession) GetSrcQty(src common.Address, dest common.Address, dstQty *big.Int, rate *big.Int) (*big.Int, error) {
	return _KyberReserve.Contract.GetSrcQty(&_KyberReserve.CallOpts, src, dest, dstQty, rate)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_KyberReserve *KyberReserveCaller) KyberNetwork(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "kyberNetwork")
	return *ret0, err
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_KyberReserve *KyberReserveSession) KyberNetwork() (common.Address, error) {
	return _KyberReserve.Contract.KyberNetwork(&_KyberReserve.CallOpts)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_KyberReserve *KyberReserveCallerSession) KyberNetwork() (common.Address, error) {
	return _KyberReserve.Contract.KyberNetwork(&_KyberReserve.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberReserve *KyberReserveCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberReserve *KyberReserveSession) PendingAdmin() (common.Address, error) {
	return _KyberReserve.Contract.PendingAdmin(&_KyberReserve.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberReserve *KyberReserveCallerSession) PendingAdmin() (common.Address, error) {
	return _KyberReserve.Contract.PendingAdmin(&_KyberReserve.CallOpts)
}

// SanityRatesContract is a free data retrieval call binding the contract method 0x47e6924f.
//
// Solidity: function sanityRatesContract() constant returns(address)
func (_KyberReserve *KyberReserveCaller) SanityRatesContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "sanityRatesContract")
	return *ret0, err
}

// SanityRatesContract is a free data retrieval call binding the contract method 0x47e6924f.
//
// Solidity: function sanityRatesContract() constant returns(address)
func (_KyberReserve *KyberReserveSession) SanityRatesContract() (common.Address, error) {
	return _KyberReserve.Contract.SanityRatesContract(&_KyberReserve.CallOpts)
}

// SanityRatesContract is a free data retrieval call binding the contract method 0x47e6924f.
//
// Solidity: function sanityRatesContract() constant returns(address)
func (_KyberReserve *KyberReserveCallerSession) SanityRatesContract() (common.Address, error) {
	return _KyberReserve.Contract.SanityRatesContract(&_KyberReserve.CallOpts)
}

// TokenWallet is a free data retrieval call binding the contract method 0xa80cbac6.
//
// Solidity: function tokenWallet(address ) constant returns(address)
func (_KyberReserve *KyberReserveCaller) TokenWallet(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "tokenWallet", arg0)
	return *ret0, err
}

// TokenWallet is a free data retrieval call binding the contract method 0xa80cbac6.
//
// Solidity: function tokenWallet(address ) constant returns(address)
func (_KyberReserve *KyberReserveSession) TokenWallet(arg0 common.Address) (common.Address, error) {
	return _KyberReserve.Contract.TokenWallet(&_KyberReserve.CallOpts, arg0)
}

// TokenWallet is a free data retrieval call binding the contract method 0xa80cbac6.
//
// Solidity: function tokenWallet(address ) constant returns(address)
func (_KyberReserve *KyberReserveCallerSession) TokenWallet(arg0 common.Address) (common.Address, error) {
	return _KyberReserve.Contract.TokenWallet(&_KyberReserve.CallOpts, arg0)
}

// TradeEnabled is a free data retrieval call binding the contract method 0xd621e813.
//
// Solidity: function tradeEnabled() constant returns(bool)
func (_KyberReserve *KyberReserveCaller) TradeEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KyberReserve.contract.Call(opts, out, "tradeEnabled")
	return *ret0, err
}

// TradeEnabled is a free data retrieval call binding the contract method 0xd621e813.
//
// Solidity: function tradeEnabled() constant returns(bool)
func (_KyberReserve *KyberReserveSession) TradeEnabled() (bool, error) {
	return _KyberReserve.Contract.TradeEnabled(&_KyberReserve.CallOpts)
}

// TradeEnabled is a free data retrieval call binding the contract method 0xd621e813.
//
// Solidity: function tradeEnabled() constant returns(bool)
func (_KyberReserve *KyberReserveCallerSession) TradeEnabled() (bool, error) {
	return _KyberReserve.Contract.TradeEnabled(&_KyberReserve.CallOpts)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberReserve *KyberReserveTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberReserve *KyberReserveSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.AddAlerter(&_KyberReserve.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberReserve *KyberReserveTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.AddAlerter(&_KyberReserve.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberReserve *KyberReserveTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberReserve *KyberReserveSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.AddOperator(&_KyberReserve.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberReserve *KyberReserveTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.AddOperator(&_KyberReserve.TransactOpts, newOperator)
}

// ApproveWithdrawAddress is a paid mutator transaction binding the contract method 0x546dc71c.
//
// Solidity: function approveWithdrawAddress(address token, address addr, bool approve) returns()
func (_KyberReserve *KyberReserveTransactor) ApproveWithdrawAddress(opts *bind.TransactOpts, token common.Address, addr common.Address, approve bool) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "approveWithdrawAddress", token, addr, approve)
}

// ApproveWithdrawAddress is a paid mutator transaction binding the contract method 0x546dc71c.
//
// Solidity: function approveWithdrawAddress(address token, address addr, bool approve) returns()
func (_KyberReserve *KyberReserveSession) ApproveWithdrawAddress(token common.Address, addr common.Address, approve bool) (*types.Transaction, error) {
	return _KyberReserve.Contract.ApproveWithdrawAddress(&_KyberReserve.TransactOpts, token, addr, approve)
}

// ApproveWithdrawAddress is a paid mutator transaction binding the contract method 0x546dc71c.
//
// Solidity: function approveWithdrawAddress(address token, address addr, bool approve) returns()
func (_KyberReserve *KyberReserveTransactorSession) ApproveWithdrawAddress(token common.Address, addr common.Address, approve bool) (*types.Transaction, error) {
	return _KyberReserve.Contract.ApproveWithdrawAddress(&_KyberReserve.TransactOpts, token, addr, approve)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberReserve *KyberReserveTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberReserve *KyberReserveSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberReserve.Contract.ClaimAdmin(&_KyberReserve.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberReserve *KyberReserveTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberReserve.Contract.ClaimAdmin(&_KyberReserve.TransactOpts)
}

// DisableTrade is a paid mutator transaction binding the contract method 0x6940030f.
//
// Solidity: function disableTrade() returns(bool)
func (_KyberReserve *KyberReserveTransactor) DisableTrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "disableTrade")
}

// DisableTrade is a paid mutator transaction binding the contract method 0x6940030f.
//
// Solidity: function disableTrade() returns(bool)
func (_KyberReserve *KyberReserveSession) DisableTrade() (*types.Transaction, error) {
	return _KyberReserve.Contract.DisableTrade(&_KyberReserve.TransactOpts)
}

// DisableTrade is a paid mutator transaction binding the contract method 0x6940030f.
//
// Solidity: function disableTrade() returns(bool)
func (_KyberReserve *KyberReserveTransactorSession) DisableTrade() (*types.Transaction, error) {
	return _KyberReserve.Contract.DisableTrade(&_KyberReserve.TransactOpts)
}

// EnableTrade is a paid mutator transaction binding the contract method 0x0099d386.
//
// Solidity: function enableTrade() returns(bool)
func (_KyberReserve *KyberReserveTransactor) EnableTrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "enableTrade")
}

// EnableTrade is a paid mutator transaction binding the contract method 0x0099d386.
//
// Solidity: function enableTrade() returns(bool)
func (_KyberReserve *KyberReserveSession) EnableTrade() (*types.Transaction, error) {
	return _KyberReserve.Contract.EnableTrade(&_KyberReserve.TransactOpts)
}

// EnableTrade is a paid mutator transaction binding the contract method 0x0099d386.
//
// Solidity: function enableTrade() returns(bool)
func (_KyberReserve *KyberReserveTransactorSession) EnableTrade() (*types.Transaction, error) {
	return _KyberReserve.Contract.EnableTrade(&_KyberReserve.TransactOpts)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberReserve *KyberReserveTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberReserve *KyberReserveSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.RemoveAlerter(&_KyberReserve.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberReserve *KyberReserveTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.RemoveAlerter(&_KyberReserve.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberReserve *KyberReserveTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberReserve *KyberReserveSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.RemoveOperator(&_KyberReserve.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberReserve *KyberReserveTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.RemoveOperator(&_KyberReserve.TransactOpts, operator)
}

// SetContracts is a paid mutator transaction binding the contract method 0xb3066d49.
//
// Solidity: function setContracts(address _kyberNetwork, address _conversionRates, address _sanityRates) returns()
func (_KyberReserve *KyberReserveTransactor) SetContracts(opts *bind.TransactOpts, _kyberNetwork common.Address, _conversionRates common.Address, _sanityRates common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "setContracts", _kyberNetwork, _conversionRates, _sanityRates)
}

// SetContracts is a paid mutator transaction binding the contract method 0xb3066d49.
//
// Solidity: function setContracts(address _kyberNetwork, address _conversionRates, address _sanityRates) returns()
func (_KyberReserve *KyberReserveSession) SetContracts(_kyberNetwork common.Address, _conversionRates common.Address, _sanityRates common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.SetContracts(&_KyberReserve.TransactOpts, _kyberNetwork, _conversionRates, _sanityRates)
}

// SetContracts is a paid mutator transaction binding the contract method 0xb3066d49.
//
// Solidity: function setContracts(address _kyberNetwork, address _conversionRates, address _sanityRates) returns()
func (_KyberReserve *KyberReserveTransactorSession) SetContracts(_kyberNetwork common.Address, _conversionRates common.Address, _sanityRates common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.SetContracts(&_KyberReserve.TransactOpts, _kyberNetwork, _conversionRates, _sanityRates)
}

// SetTokenWallet is a paid mutator transaction binding the contract method 0x1bc7bfec.
//
// Solidity: function setTokenWallet(address token, address wallet) returns()
func (_KyberReserve *KyberReserveTransactor) SetTokenWallet(opts *bind.TransactOpts, token common.Address, wallet common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "setTokenWallet", token, wallet)
}

// SetTokenWallet is a paid mutator transaction binding the contract method 0x1bc7bfec.
//
// Solidity: function setTokenWallet(address token, address wallet) returns()
func (_KyberReserve *KyberReserveSession) SetTokenWallet(token common.Address, wallet common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.SetTokenWallet(&_KyberReserve.TransactOpts, token, wallet)
}

// SetTokenWallet is a paid mutator transaction binding the contract method 0x1bc7bfec.
//
// Solidity: function setTokenWallet(address token, address wallet) returns()
func (_KyberReserve *KyberReserveTransactorSession) SetTokenWallet(token common.Address, wallet common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.SetTokenWallet(&_KyberReserve.TransactOpts, token, wallet)
}

// Trade is a paid mutator transaction binding the contract method 0x6cf69811.
//
// Solidity: function trade(address srcToken, uint256 srcAmount, address destToken, address destAddress, uint256 conversionRate, bool validate) returns(bool)
func (_KyberReserve *KyberReserveTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, srcAmount *big.Int, destToken common.Address, destAddress common.Address, conversionRate *big.Int, validate bool) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "trade", srcToken, srcAmount, destToken, destAddress, conversionRate, validate)
}

// Trade is a paid mutator transaction binding the contract method 0x6cf69811.
//
// Solidity: function trade(address srcToken, uint256 srcAmount, address destToken, address destAddress, uint256 conversionRate, bool validate) returns(bool)
func (_KyberReserve *KyberReserveSession) Trade(srcToken common.Address, srcAmount *big.Int, destToken common.Address, destAddress common.Address, conversionRate *big.Int, validate bool) (*types.Transaction, error) {
	return _KyberReserve.Contract.Trade(&_KyberReserve.TransactOpts, srcToken, srcAmount, destToken, destAddress, conversionRate, validate)
}

// Trade is a paid mutator transaction binding the contract method 0x6cf69811.
//
// Solidity: function trade(address srcToken, uint256 srcAmount, address destToken, address destAddress, uint256 conversionRate, bool validate) returns(bool)
func (_KyberReserve *KyberReserveTransactorSession) Trade(srcToken common.Address, srcAmount *big.Int, destToken common.Address, destAddress common.Address, conversionRate *big.Int, validate bool) (*types.Transaction, error) {
	return _KyberReserve.Contract.Trade(&_KyberReserve.TransactOpts, srcToken, srcAmount, destToken, destAddress, conversionRate, validate)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberReserve *KyberReserveTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberReserve *KyberReserveSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.TransferAdmin(&_KyberReserve.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberReserve *KyberReserveTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.TransferAdmin(&_KyberReserve.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberReserve *KyberReserveTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberReserve *KyberReserveSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.TransferAdminQuickly(&_KyberReserve.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberReserve *KyberReserveTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.TransferAdminQuickly(&_KyberReserve.TransactOpts, newAdmin)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token, uint256 amount, address destination) returns(bool)
func (_KyberReserve *KyberReserveTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "withdraw", token, amount, destination)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token, uint256 amount, address destination) returns(bool)
func (_KyberReserve *KyberReserveSession) Withdraw(token common.Address, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.Withdraw(&_KyberReserve.TransactOpts, token, amount, destination)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address token, uint256 amount, address destination) returns(bool)
func (_KyberReserve *KyberReserveTransactorSession) Withdraw(token common.Address, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.Withdraw(&_KyberReserve.TransactOpts, token, amount, destination)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberReserve *KyberReserveTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberReserve *KyberReserveSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.WithdrawEther(&_KyberReserve.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberReserve *KyberReserveTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.WithdrawEther(&_KyberReserve.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberReserve *KyberReserveTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberReserve.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberReserve *KyberReserveSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.WithdrawToken(&_KyberReserve.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberReserve *KyberReserveTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberReserve.Contract.WithdrawToken(&_KyberReserve.TransactOpts, token, amount, sendTo)
}

// KyberReserveAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the KyberReserve contract.
type KyberReserveAdminClaimedIterator struct {
	Event *KyberReserveAdminClaimed // Event containing the contract specifics and raw log

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
func (it *KyberReserveAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveAdminClaimed)
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
		it.Event = new(KyberReserveAdminClaimed)
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
func (it *KyberReserveAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveAdminClaimed represents a AdminClaimed event raised by the KyberReserve contract.
type KyberReserveAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberReserve *KyberReserveFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*KyberReserveAdminClaimedIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &KyberReserveAdminClaimedIterator{contract: _KyberReserve.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberReserve *KyberReserveFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *KyberReserveAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveAdminClaimed)
				if err := _KyberReserve.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_KyberReserve *KyberReserveFilterer) ParseAdminClaimed(log types.Log) (*KyberReserveAdminClaimed, error) {
	event := new(KyberReserveAdminClaimed)
	if err := _KyberReserve.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the KyberReserve contract.
type KyberReserveAlerterAddedIterator struct {
	Event *KyberReserveAlerterAdded // Event containing the contract specifics and raw log

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
func (it *KyberReserveAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveAlerterAdded)
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
		it.Event = new(KyberReserveAlerterAdded)
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
func (it *KyberReserveAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveAlerterAdded represents a AlerterAdded event raised by the KyberReserve contract.
type KyberReserveAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberReserve *KyberReserveFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*KyberReserveAlerterAddedIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &KyberReserveAlerterAddedIterator{contract: _KyberReserve.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberReserve *KyberReserveFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *KyberReserveAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveAlerterAdded)
				if err := _KyberReserve.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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
func (_KyberReserve *KyberReserveFilterer) ParseAlerterAdded(log types.Log) (*KyberReserveAlerterAdded, error) {
	event := new(KyberReserveAlerterAdded)
	if err := _KyberReserve.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveDepositTokenIterator is returned from FilterDepositToken and is used to iterate over the raw logs and unpacked data for DepositToken events raised by the KyberReserve contract.
type KyberReserveDepositTokenIterator struct {
	Event *KyberReserveDepositToken // Event containing the contract specifics and raw log

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
func (it *KyberReserveDepositTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveDepositToken)
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
		it.Event = new(KyberReserveDepositToken)
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
func (it *KyberReserveDepositTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveDepositTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveDepositToken represents a DepositToken event raised by the KyberReserve contract.
type KyberReserveDepositToken struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositToken is a free log retrieval operation binding the contract event 0x2d0c0a8842b9944ece1495eb61121621b5e36bd6af3bba0318c695f525aef79f.
//
// Solidity: event DepositToken(address token, uint256 amount)
func (_KyberReserve *KyberReserveFilterer) FilterDepositToken(opts *bind.FilterOpts) (*KyberReserveDepositTokenIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "DepositToken")
	if err != nil {
		return nil, err
	}
	return &KyberReserveDepositTokenIterator{contract: _KyberReserve.contract, event: "DepositToken", logs: logs, sub: sub}, nil
}

// WatchDepositToken is a free log subscription operation binding the contract event 0x2d0c0a8842b9944ece1495eb61121621b5e36bd6af3bba0318c695f525aef79f.
//
// Solidity: event DepositToken(address token, uint256 amount)
func (_KyberReserve *KyberReserveFilterer) WatchDepositToken(opts *bind.WatchOpts, sink chan<- *KyberReserveDepositToken) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "DepositToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveDepositToken)
				if err := _KyberReserve.contract.UnpackLog(event, "DepositToken", log); err != nil {
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

// ParseDepositToken is a log parse operation binding the contract event 0x2d0c0a8842b9944ece1495eb61121621b5e36bd6af3bba0318c695f525aef79f.
//
// Solidity: event DepositToken(address token, uint256 amount)
func (_KyberReserve *KyberReserveFilterer) ParseDepositToken(log types.Log) (*KyberReserveDepositToken, error) {
	event := new(KyberReserveDepositToken)
	if err := _KyberReserve.contract.UnpackLog(event, "DepositToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the KyberReserve contract.
type KyberReserveEtherWithdrawIterator struct {
	Event *KyberReserveEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberReserveEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveEtherWithdraw)
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
		it.Event = new(KyberReserveEtherWithdraw)
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
func (it *KyberReserveEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveEtherWithdraw represents a EtherWithdraw event raised by the KyberReserve contract.
type KyberReserveEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberReserve *KyberReserveFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*KyberReserveEtherWithdrawIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberReserveEtherWithdrawIterator{contract: _KyberReserve.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberReserve *KyberReserveFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *KyberReserveEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveEtherWithdraw)
				if err := _KyberReserve.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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
func (_KyberReserve *KyberReserveFilterer) ParseEtherWithdraw(log types.Log) (*KyberReserveEtherWithdraw, error) {
	event := new(KyberReserveEtherWithdraw)
	if err := _KyberReserve.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveNewTokenWalletIterator is returned from FilterNewTokenWallet and is used to iterate over the raw logs and unpacked data for NewTokenWallet events raised by the KyberReserve contract.
type KyberReserveNewTokenWalletIterator struct {
	Event *KyberReserveNewTokenWallet // Event containing the contract specifics and raw log

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
func (it *KyberReserveNewTokenWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveNewTokenWallet)
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
		it.Event = new(KyberReserveNewTokenWallet)
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
func (it *KyberReserveNewTokenWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveNewTokenWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveNewTokenWallet represents a NewTokenWallet event raised by the KyberReserve contract.
type KyberReserveNewTokenWallet struct {
	Token  common.Address
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewTokenWallet is a free log retrieval operation binding the contract event 0x81995c7b922889ac0a81e41866106d4046268ea3a9abaae9f9e080a6ce36ee7d.
//
// Solidity: event NewTokenWallet(address token, address wallet)
func (_KyberReserve *KyberReserveFilterer) FilterNewTokenWallet(opts *bind.FilterOpts) (*KyberReserveNewTokenWalletIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "NewTokenWallet")
	if err != nil {
		return nil, err
	}
	return &KyberReserveNewTokenWalletIterator{contract: _KyberReserve.contract, event: "NewTokenWallet", logs: logs, sub: sub}, nil
}

// WatchNewTokenWallet is a free log subscription operation binding the contract event 0x81995c7b922889ac0a81e41866106d4046268ea3a9abaae9f9e080a6ce36ee7d.
//
// Solidity: event NewTokenWallet(address token, address wallet)
func (_KyberReserve *KyberReserveFilterer) WatchNewTokenWallet(opts *bind.WatchOpts, sink chan<- *KyberReserveNewTokenWallet) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "NewTokenWallet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveNewTokenWallet)
				if err := _KyberReserve.contract.UnpackLog(event, "NewTokenWallet", log); err != nil {
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

// ParseNewTokenWallet is a log parse operation binding the contract event 0x81995c7b922889ac0a81e41866106d4046268ea3a9abaae9f9e080a6ce36ee7d.
//
// Solidity: event NewTokenWallet(address token, address wallet)
func (_KyberReserve *KyberReserveFilterer) ParseNewTokenWallet(log types.Log) (*KyberReserveNewTokenWallet, error) {
	event := new(KyberReserveNewTokenWallet)
	if err := _KyberReserve.contract.UnpackLog(event, "NewTokenWallet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the KyberReserve contract.
type KyberReserveOperatorAddedIterator struct {
	Event *KyberReserveOperatorAdded // Event containing the contract specifics and raw log

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
func (it *KyberReserveOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveOperatorAdded)
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
		it.Event = new(KyberReserveOperatorAdded)
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
func (it *KyberReserveOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveOperatorAdded represents a OperatorAdded event raised by the KyberReserve contract.
type KyberReserveOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberReserve *KyberReserveFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*KyberReserveOperatorAddedIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &KyberReserveOperatorAddedIterator{contract: _KyberReserve.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberReserve *KyberReserveFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *KyberReserveOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveOperatorAdded)
				if err := _KyberReserve.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
func (_KyberReserve *KyberReserveFilterer) ParseOperatorAdded(log types.Log) (*KyberReserveOperatorAdded, error) {
	event := new(KyberReserveOperatorAdded)
	if err := _KyberReserve.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveSetContractAddressesIterator is returned from FilterSetContractAddresses and is used to iterate over the raw logs and unpacked data for SetContractAddresses events raised by the KyberReserve contract.
type KyberReserveSetContractAddressesIterator struct {
	Event *KyberReserveSetContractAddresses // Event containing the contract specifics and raw log

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
func (it *KyberReserveSetContractAddressesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveSetContractAddresses)
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
		it.Event = new(KyberReserveSetContractAddresses)
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
func (it *KyberReserveSetContractAddressesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveSetContractAddressesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveSetContractAddresses represents a SetContractAddresses event raised by the KyberReserve contract.
type KyberReserveSetContractAddresses struct {
	Network common.Address
	Rate    common.Address
	Sanity  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetContractAddresses is a free log retrieval operation binding the contract event 0x7a85322644a4462d8ff5482d2a841a4d231f8cfb3c9f4a50f66f8b2bd568c31c.
//
// Solidity: event SetContractAddresses(address network, address rate, address sanity)
func (_KyberReserve *KyberReserveFilterer) FilterSetContractAddresses(opts *bind.FilterOpts) (*KyberReserveSetContractAddressesIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "SetContractAddresses")
	if err != nil {
		return nil, err
	}
	return &KyberReserveSetContractAddressesIterator{contract: _KyberReserve.contract, event: "SetContractAddresses", logs: logs, sub: sub}, nil
}

// WatchSetContractAddresses is a free log subscription operation binding the contract event 0x7a85322644a4462d8ff5482d2a841a4d231f8cfb3c9f4a50f66f8b2bd568c31c.
//
// Solidity: event SetContractAddresses(address network, address rate, address sanity)
func (_KyberReserve *KyberReserveFilterer) WatchSetContractAddresses(opts *bind.WatchOpts, sink chan<- *KyberReserveSetContractAddresses) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "SetContractAddresses")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveSetContractAddresses)
				if err := _KyberReserve.contract.UnpackLog(event, "SetContractAddresses", log); err != nil {
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

// ParseSetContractAddresses is a log parse operation binding the contract event 0x7a85322644a4462d8ff5482d2a841a4d231f8cfb3c9f4a50f66f8b2bd568c31c.
//
// Solidity: event SetContractAddresses(address network, address rate, address sanity)
func (_KyberReserve *KyberReserveFilterer) ParseSetContractAddresses(log types.Log) (*KyberReserveSetContractAddresses, error) {
	event := new(KyberReserveSetContractAddresses)
	if err := _KyberReserve.contract.UnpackLog(event, "SetContractAddresses", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the KyberReserve contract.
type KyberReserveTokenWithdrawIterator struct {
	Event *KyberReserveTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberReserveTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveTokenWithdraw)
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
		it.Event = new(KyberReserveTokenWithdraw)
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
func (it *KyberReserveTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveTokenWithdraw represents a TokenWithdraw event raised by the KyberReserve contract.
type KyberReserveTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberReserve *KyberReserveFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*KyberReserveTokenWithdrawIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberReserveTokenWithdrawIterator{contract: _KyberReserve.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberReserve *KyberReserveFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *KyberReserveTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveTokenWithdraw)
				if err := _KyberReserve.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_KyberReserve *KyberReserveFilterer) ParseTokenWithdraw(log types.Log) (*KyberReserveTokenWithdraw, error) {
	event := new(KyberReserveTokenWithdraw)
	if err := _KyberReserve.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveTradeEnabledIterator is returned from FilterTradeEnabled and is used to iterate over the raw logs and unpacked data for TradeEnabled events raised by the KyberReserve contract.
type KyberReserveTradeEnabledIterator struct {
	Event *KyberReserveTradeEnabled // Event containing the contract specifics and raw log

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
func (it *KyberReserveTradeEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveTradeEnabled)
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
		it.Event = new(KyberReserveTradeEnabled)
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
func (it *KyberReserveTradeEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveTradeEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveTradeEnabled represents a TradeEnabled event raised by the KyberReserve contract.
type KyberReserveTradeEnabled struct {
	Enable bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradeEnabled is a free log retrieval operation binding the contract event 0x7d7f00509dd73ac4449f698ae75ccc797895eff5fa9d446d3df387598a26e735.
//
// Solidity: event TradeEnabled(bool enable)
func (_KyberReserve *KyberReserveFilterer) FilterTradeEnabled(opts *bind.FilterOpts) (*KyberReserveTradeEnabledIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "TradeEnabled")
	if err != nil {
		return nil, err
	}
	return &KyberReserveTradeEnabledIterator{contract: _KyberReserve.contract, event: "TradeEnabled", logs: logs, sub: sub}, nil
}

// WatchTradeEnabled is a free log subscription operation binding the contract event 0x7d7f00509dd73ac4449f698ae75ccc797895eff5fa9d446d3df387598a26e735.
//
// Solidity: event TradeEnabled(bool enable)
func (_KyberReserve *KyberReserveFilterer) WatchTradeEnabled(opts *bind.WatchOpts, sink chan<- *KyberReserveTradeEnabled) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "TradeEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveTradeEnabled)
				if err := _KyberReserve.contract.UnpackLog(event, "TradeEnabled", log); err != nil {
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

// ParseTradeEnabled is a log parse operation binding the contract event 0x7d7f00509dd73ac4449f698ae75ccc797895eff5fa9d446d3df387598a26e735.
//
// Solidity: event TradeEnabled(bool enable)
func (_KyberReserve *KyberReserveFilterer) ParseTradeEnabled(log types.Log) (*KyberReserveTradeEnabled, error) {
	event := new(KyberReserveTradeEnabled)
	if err := _KyberReserve.contract.UnpackLog(event, "TradeEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveTradeExecuteIterator is returned from FilterTradeExecute and is used to iterate over the raw logs and unpacked data for TradeExecute events raised by the KyberReserve contract.
type KyberReserveTradeExecuteIterator struct {
	Event *KyberReserveTradeExecute // Event containing the contract specifics and raw log

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
func (it *KyberReserveTradeExecuteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveTradeExecute)
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
		it.Event = new(KyberReserveTradeExecute)
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
func (it *KyberReserveTradeExecuteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveTradeExecuteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveTradeExecute represents a TradeExecute event raised by the KyberReserve contract.
type KyberReserveTradeExecute struct {
	Origin      common.Address
	Src         common.Address
	SrcAmount   *big.Int
	DestToken   common.Address
	DestAmount  *big.Int
	DestAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTradeExecute is a free log retrieval operation binding the contract event 0xea9415385bae08fe9f6dc457b02577166790cde83bb18cc340aac6cb81b824de.
//
// Solidity: event TradeExecute(address indexed origin, address src, uint256 srcAmount, address destToken, uint256 destAmount, address destAddress)
func (_KyberReserve *KyberReserveFilterer) FilterTradeExecute(opts *bind.FilterOpts, origin []common.Address) (*KyberReserveTradeExecuteIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "TradeExecute", originRule)
	if err != nil {
		return nil, err
	}
	return &KyberReserveTradeExecuteIterator{contract: _KyberReserve.contract, event: "TradeExecute", logs: logs, sub: sub}, nil
}

// WatchTradeExecute is a free log subscription operation binding the contract event 0xea9415385bae08fe9f6dc457b02577166790cde83bb18cc340aac6cb81b824de.
//
// Solidity: event TradeExecute(address indexed origin, address src, uint256 srcAmount, address destToken, uint256 destAmount, address destAddress)
func (_KyberReserve *KyberReserveFilterer) WatchTradeExecute(opts *bind.WatchOpts, sink chan<- *KyberReserveTradeExecute, origin []common.Address) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "TradeExecute", originRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveTradeExecute)
				if err := _KyberReserve.contract.UnpackLog(event, "TradeExecute", log); err != nil {
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

// ParseTradeExecute is a log parse operation binding the contract event 0xea9415385bae08fe9f6dc457b02577166790cde83bb18cc340aac6cb81b824de.
//
// Solidity: event TradeExecute(address indexed origin, address src, uint256 srcAmount, address destToken, uint256 destAmount, address destAddress)
func (_KyberReserve *KyberReserveFilterer) ParseTradeExecute(log types.Log) (*KyberReserveTradeExecute, error) {
	event := new(KyberReserveTradeExecute)
	if err := _KyberReserve.contract.UnpackLog(event, "TradeExecute", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the KyberReserve contract.
type KyberReserveTransferAdminPendingIterator struct {
	Event *KyberReserveTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *KyberReserveTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveTransferAdminPending)
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
		it.Event = new(KyberReserveTransferAdminPending)
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
func (it *KyberReserveTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveTransferAdminPending represents a TransferAdminPending event raised by the KyberReserve contract.
type KyberReserveTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberReserve *KyberReserveFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*KyberReserveTransferAdminPendingIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &KyberReserveTransferAdminPendingIterator{contract: _KyberReserve.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberReserve *KyberReserveFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *KyberReserveTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveTransferAdminPending)
				if err := _KyberReserve.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_KyberReserve *KyberReserveFilterer) ParseTransferAdminPending(log types.Log) (*KyberReserveTransferAdminPending, error) {
	event := new(KyberReserveTransferAdminPending)
	if err := _KyberReserve.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveWithdrawAddressApprovedIterator is returned from FilterWithdrawAddressApproved and is used to iterate over the raw logs and unpacked data for WithdrawAddressApproved events raised by the KyberReserve contract.
type KyberReserveWithdrawAddressApprovedIterator struct {
	Event *KyberReserveWithdrawAddressApproved // Event containing the contract specifics and raw log

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
func (it *KyberReserveWithdrawAddressApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveWithdrawAddressApproved)
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
		it.Event = new(KyberReserveWithdrawAddressApproved)
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
func (it *KyberReserveWithdrawAddressApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveWithdrawAddressApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveWithdrawAddressApproved represents a WithdrawAddressApproved event raised by the KyberReserve contract.
type KyberReserveWithdrawAddressApproved struct {
	Token   common.Address
	Addr    common.Address
	Approve bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAddressApproved is a free log retrieval operation binding the contract event 0xd5fd5351efae1f4bb760079da9f0ff9589e2c3e216337ca9d39cdff573b245c4.
//
// Solidity: event WithdrawAddressApproved(address token, address addr, bool approve)
func (_KyberReserve *KyberReserveFilterer) FilterWithdrawAddressApproved(opts *bind.FilterOpts) (*KyberReserveWithdrawAddressApprovedIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "WithdrawAddressApproved")
	if err != nil {
		return nil, err
	}
	return &KyberReserveWithdrawAddressApprovedIterator{contract: _KyberReserve.contract, event: "WithdrawAddressApproved", logs: logs, sub: sub}, nil
}

// WatchWithdrawAddressApproved is a free log subscription operation binding the contract event 0xd5fd5351efae1f4bb760079da9f0ff9589e2c3e216337ca9d39cdff573b245c4.
//
// Solidity: event WithdrawAddressApproved(address token, address addr, bool approve)
func (_KyberReserve *KyberReserveFilterer) WatchWithdrawAddressApproved(opts *bind.WatchOpts, sink chan<- *KyberReserveWithdrawAddressApproved) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "WithdrawAddressApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveWithdrawAddressApproved)
				if err := _KyberReserve.contract.UnpackLog(event, "WithdrawAddressApproved", log); err != nil {
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

// ParseWithdrawAddressApproved is a log parse operation binding the contract event 0xd5fd5351efae1f4bb760079da9f0ff9589e2c3e216337ca9d39cdff573b245c4.
//
// Solidity: event WithdrawAddressApproved(address token, address addr, bool approve)
func (_KyberReserve *KyberReserveFilterer) ParseWithdrawAddressApproved(log types.Log) (*KyberReserveWithdrawAddressApproved, error) {
	event := new(KyberReserveWithdrawAddressApproved)
	if err := _KyberReserve.contract.UnpackLog(event, "WithdrawAddressApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberReserveWithdrawFundsIterator is returned from FilterWithdrawFunds and is used to iterate over the raw logs and unpacked data for WithdrawFunds events raised by the KyberReserve contract.
type KyberReserveWithdrawFundsIterator struct {
	Event *KyberReserveWithdrawFunds // Event containing the contract specifics and raw log

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
func (it *KyberReserveWithdrawFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberReserveWithdrawFunds)
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
		it.Event = new(KyberReserveWithdrawFunds)
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
func (it *KyberReserveWithdrawFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberReserveWithdrawFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberReserveWithdrawFunds represents a WithdrawFunds event raised by the KyberReserve contract.
type KyberReserveWithdrawFunds struct {
	Token       common.Address
	Amount      *big.Int
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFunds is a free log retrieval operation binding the contract event 0xb67719fc33c1f17d31bf3a698690d62066b1e0bae28fcd3c56cf2c015c2863d6.
//
// Solidity: event WithdrawFunds(address token, uint256 amount, address destination)
func (_KyberReserve *KyberReserveFilterer) FilterWithdrawFunds(opts *bind.FilterOpts) (*KyberReserveWithdrawFundsIterator, error) {

	logs, sub, err := _KyberReserve.contract.FilterLogs(opts, "WithdrawFunds")
	if err != nil {
		return nil, err
	}
	return &KyberReserveWithdrawFundsIterator{contract: _KyberReserve.contract, event: "WithdrawFunds", logs: logs, sub: sub}, nil
}

// WatchWithdrawFunds is a free log subscription operation binding the contract event 0xb67719fc33c1f17d31bf3a698690d62066b1e0bae28fcd3c56cf2c015c2863d6.
//
// Solidity: event WithdrawFunds(address token, uint256 amount, address destination)
func (_KyberReserve *KyberReserveFilterer) WatchWithdrawFunds(opts *bind.WatchOpts, sink chan<- *KyberReserveWithdrawFunds) (event.Subscription, error) {

	logs, sub, err := _KyberReserve.contract.WatchLogs(opts, "WithdrawFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberReserveWithdrawFunds)
				if err := _KyberReserve.contract.UnpackLog(event, "WithdrawFunds", log); err != nil {
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

// ParseWithdrawFunds is a log parse operation binding the contract event 0xb67719fc33c1f17d31bf3a698690d62066b1e0bae28fcd3c56cf2c015c2863d6.
//
// Solidity: event WithdrawFunds(address token, uint256 amount, address destination)
func (_KyberReserve *KyberReserveFilterer) ParseWithdrawFunds(log types.Log) (*KyberReserveWithdrawFunds, error) {
	event := new(KyberReserveWithdrawFunds)
	if err := _KyberReserve.contract.UnpackLog(event, "WithdrawFunds", log); err != nil {
		return nil, err
	}
	return event, nil
}
