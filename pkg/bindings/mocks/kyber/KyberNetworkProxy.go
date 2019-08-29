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

// KyberNetworkProxyABI is the input ABI used to generate the binding from.
const KyberNetworkProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"name\":\"walletId\",\"type\":\"address\"},{\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"tradeWithHint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxGasPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetworkContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserCapInWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapEtherToToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"srcQty\",\"type\":\"uint256\"}],\"name\":\"getExpectedRate\",\"outputs\":[{\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"name\":\"slippageRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getUserCapInTokenWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kyberNetworkContract\",\"type\":\"address\"}],\"name\":\"setKyberNetworkContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"info\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"name\":\"walletId\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"actualSrcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"actualDestAmount\",\"type\":\"uint256\"}],\"name\":\"ExecuteTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newNetworkContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldNetworkContract\",\"type\":\"address\"}],\"name\":\"KyberNetworkSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"}]"

// KyberNetworkProxyBin is the compiled bytecode used for deploying new contracts.
var KyberNetworkProxyBin = "0x6060604052341561000f57600080fd5b604051602080611a048339810160405280805160008054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061005557600080fd5b60008054600160a060020a03909216600160a060020a031990921691909117905561197f806100856000396000f3006060604052600436106101455763ffffffff60e060020a60003504166301a12fd3811461014a578063238dafe01461016b578063267822471461019257806327a099d8146101c157806329589f61146102275780633bba21dc146102ad5780633ccdbb28146102d25780633de39c11146102fb578063408ee7fe1461030e5780634f61ff8b1461032d5780636432679f146103405780637409e2eb1461035f57806375829def1461038b57806377f50f97146103aa5780637a2a0456146103bd5780637acc8678146103d45780637c423f54146103f3578063809a9e55146104065780638eaaeecf146104465780639870d7fe1461046b578063abd188a81461048a578063ac8a584a146104a9578063b64a097e146104c8578063cb3c28c7146104de578063ce56c45414610510578063d4fac45d14610532578063f851a44014610557575b600080fd5b341561015557600080fd5b610169600160a060020a036004351661056a565b005b341561017657600080fd5b61017e6106da565b604051901515815260200160405180910390f35b341561019d57600080fd5b6101a5610744565b604051600160a060020a03909116815260200160405180910390f35b34156101cc57600080fd5b6101d4610753565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102135780820151838201526020016101fb565b505050509050019250505060405180910390f35b61029b600160a060020a036004803582169160248035926044358316926064358116926084359260a4359260c43516916101049060e43590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506107bb95505050505050565b60405190815260200160405180910390f35b34156102b857600080fd5b61029b600160a060020a0360043516602435604435610ad9565b34156102dd57600080fd5b610169600160a060020a036004358116906024359060443516610b1d565b341561030657600080fd5b61029b610c14565b341561031957600080fd5b610169600160a060020a0360043516610c5e565b341561033857600080fd5b6101a5610d5a565b341561034b57600080fd5b61029b600160a060020a0360043516610d69565b341561036a57600080fd5b61029b600160a060020a036004358116906024359060443516606435610de4565b341561039657600080fd5b610169600160a060020a0360043516610e15565b34156103b557600080fd5b610169610eb0565b61029b600160a060020a0360043516602435610f4a565b34156103df57600080fd5b610169600160a060020a0360043516610f8d565b34156103fe57600080fd5b6101d461106f565b341561041157600080fd5b61042e600160a060020a03600435811690602435166044356110d5565b60405191825260208201526040908101905180910390f35b341561045157600080fd5b61029b600160a060020a0360043581169060243516611171565b341561047657600080fd5b610169600160a060020a03600435166111f7565b341561049557600080fd5b610169600160a060020a03600435166112c7565b34156104b457600080fd5b610169600160a060020a036004351661136c565b34156104d357600080fd5b61029b6004356114d8565b61029b600160a060020a03600435811690602435906044358116906064358116906084359060a4359060c4351661152b565b341561051b57600080fd5b610169600435600160a060020a0360243516611552565b341561053d57600080fd5b61029b600160a060020a03600435811690602435166115e5565b341561056257600080fd5b6101a5611696565b6000805433600160a060020a0390811691161461058657600080fd5b600160a060020a03821660009081526003602052604090205460ff1615156105ad57600080fd5b50600160a060020a0381166000908152600360205260408120805460ff191690555b6005548110156106d65781600160a060020a03166005828154811015156105f257fe5b600091825260209091200154600160a060020a031614156106ce5760058054600019810190811061061f57fe5b60009182526020909120015460058054600160a060020a03909216918390811061064557fe5b60009182526020909120018054600160a060020a031916600160a060020a039290921691909117905560058054906106819060001983016118c1565b507f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762826000604051600160a060020a039092168252151560208201526040908101905180910390a16106d6565b6001016105cf565b5050565b600754600090600160a060020a031663238dafe082604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561072457600080fd5b6102c65a03f1151561073557600080fd5b50505060405180519150505b90565b600154600160a060020a031681565b61075b6118ea565b60048054806020026020016040519081016040528092919081815260200182805480156107b157602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610793575b5050505050905090565b60006107c56118fc565b60006107cf611913565b600160a060020a038c1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14806107f8575034155b151561080357600080fd5b61080d8c336115e5565b83526108198a8a6115e5565b6020840152600160a060020a038c1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415610851573483818151019052506108e8565b600754600160a060020a03808e16916323b872dd913391168e60006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156108c257600080fd5b6102c65a03f115156108d357600080fd5b5050506040518051905015156108e857600080fd5b600754600160a060020a031663088322ef34338f8f8f8f8f8f8f8f60006040516020015260405160e060020a63ffffffff8d16028152600160a060020a03808b16600483019081528a82166024840152604483018a90528882166064840152878216608484015260a4830187905260c4830186905290841660e4830152610120610104830190815290916101240183818151815260200191508051906020019080838360005b838110156109a657808201518382015260200161098e565b50505050905090810190601f1680156109d35780820380516001836020036101000a031916815260200191505b509a50505050505050505050506020604051808303818588803b15156109f857600080fd5b6125ee5a03f11515610a0957600080fd5b5050505060405180519250610a289050835184602001518e8d8d6116a5565b905080602001518214610a3a57600080fd5b8781602001511115610a4b57600080fd5b8681604001511015610a5c57600080fd5b600160a060020a0333167f1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a398d8c84518560200151604051600160a060020a0394851681529290931660208301526040808301919091526060820192909252608001905180910390a280602001519c9b505050505050505050505050565b6000610ae36118ea565b610b14858573eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee336b204fce5e3e25026110000000886000886107bb565b95945050505050565b60005433600160a060020a03908116911614610b3857600080fd5b82600160a060020a031663a9059cbb828460006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610b9557600080fd5b6102c65a03f11515610ba657600080fd5b505050604051805190501515610bbb57600080fd5b7f72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6838383604051600160a060020a03938416815260208101929092529091166040808301919091526060909101905180910390a1505050565b600754600090600160a060020a0316633de39c1182604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561072457600080fd5b60005433600160a060020a03908116911614610c7957600080fd5b600160a060020a03811660009081526003602052604090205460ff1615610c9f57600080fd5b60055460329010610caf57600080fd5b7f5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600360205260409020805460ff191660019081179091556005805490918101610d2e83826118c1565b5060009182526020909120018054600160a060020a031916600160a060020a0392909216919091179055565b600754600160a060020a031681565b600754600090600160a060020a0316636432679f83836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610dc457600080fd5b6102c65a03f11515610dd557600080fd5b50505060405180519392505050565b6000610dee6118ea565b610e0b868686336b204fce5e3e25026110000000886000886107bb565b9695505050505050565b60005433600160a060020a03908116911614610e3057600080fd5b600160a060020a0381161515610e4557600080fd5b6001547f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4090600160a060020a0316604051600160a060020a03909116815260200160405180910390a160018054600160a060020a031916600160a060020a0392909216919091179055565b60015433600160a060020a03908116911614610ecb57600080fd5b6001546000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed91600160a060020a039081169116604051600160a060020a039283168152911660208201526040908101905180910390a16001805460008054600160a060020a0319908116600160a060020a03841617909155169055565b6000610f546118ea565b610f8573eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee3486336b204fce5e3e25026110000000886000886107bb565b949350505050565b60005433600160a060020a03908116911614610fa857600080fd5b600160a060020a0381161515610fbd57600080fd5b7f3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc4081604051600160a060020a03909116815260200160405180910390a16000547f65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed908290600160a060020a0316604051600160a060020a039283168152911660208201526040908101905180910390a160008054600160a060020a031916600160a060020a0392909216919091179055565b6110776118ea565b60058054806020026020016040519081016040528092919081815260200182805480156107b157602002820191906000526020600020908154600160a060020a03168152600190910190602001808311610793575050505050905090565b6007546000908190600160a060020a031663809a9e55868686856040516040015260405160e060020a63ffffffff8616028152600160a060020a03938416600482015291909216602482015260448101919091526064016040805180830381600087803b151561114457600080fd5b6102c65a03f1151561115557600080fd5b5050506040518051906020018051905091509150935093915050565b600754600090600160a060020a0316638eaaeecf8484846040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b15156111d457600080fd5b6102c65a03f115156111e557600080fd5b50505060405180519150505b92915050565b60005433600160a060020a0390811691161461121257600080fd5b600160a060020a03811660009081526002602052604090205460ff161561123857600080fd5b6004546032901061124857600080fd5b7f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b816001604051600160a060020a039092168252151560208201526040908101905180910390a1600160a060020a0381166000908152600260205260409020805460ff191660019081179091556004805490918101610d2e83826118c1565b60005433600160a060020a039081169116146112e257600080fd5b600160a060020a03811615156112f757600080fd5b6007547f8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02908290600160a060020a0316604051600160a060020a039283168152911660208201526040908101905180910390a160078054600160a060020a031916600160a060020a0392909216919091179055565b6000805433600160a060020a0390811691161461138857600080fd5b600160a060020a03821660009081526002602052604090205460ff1615156113af57600080fd5b50600160a060020a0381166000908152600260205260408120805460ff191690555b6004548110156106d65781600160a060020a03166004828154811015156113f457fe5b600091825260209091200154600160a060020a031614156114d05760048054600019810190811061142157fe5b60009182526020909120015460048054600160a060020a03909216918390811061144757fe5b60009182526020909120018054600160a060020a031916600160a060020a039290921691909117905560048054600019019061148390826118c1565b507f091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b826000604051600160a060020a039092168252151560208201526040908101905180910390a16106d6565b6001016113d1565b600754600090600160a060020a031663b64a097e83836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515610dc457600080fd5b60006115356118ea565b61154589898989898989886107bb565b9998505050505050505050565b60005433600160a060020a0390811691161461156d57600080fd5b600160a060020a03811682156108fc0283604051600060405180830381858888f19350505050151561159e57600080fd5b7fec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de8282604051918252600160a060020a031660208201526040908101905180910390a15050565b6000600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141561161d5750600160a060020a038116316111f1565b82600160a060020a03166370a082318360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561167457600080fd5b6102c65a03f1151561168557600080fd5b5050506040518051905090506111f1565b600054600160a060020a031681565b6116ad611913565b6000806116ba86336115e5565b91506116c685856115e5565b90508681116116d457600080fd5b8188116116e057600080fd5b8681036020840152818803835261170e835184602001516117008961171f565b6117098961171f565b611763565b604084015250909695505050505050565b600160a060020a038116600090815260066020526040812054151561174757611747826117fe565b50600160a060020a031660009081526006602052604090205490565b60006b204fce5e3e2502611000000085111561177e57600080fd5b6b204fce5e3e2502611000000084111561179757600080fd5b8282106117d257601283830311156117ae57600080fd5b84838303600a0a02670de0b6b3a764000085028115156117ca57fe5b049050610f85565b601282840311156117e257600080fd5b84828403600a0a670de0b6b3a76400008602028115156117ca57fe5b600160a060020a03811673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141561184457600160a060020a0381166000908152600660205260409020601290556118be565b80600160a060020a031663313ce5676000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561188a57600080fd5b6102c65a03f1151561189b57600080fd5b5050506040518051600160a060020a038316600090815260066020526040902055505b50565b8154818355818115116118e5576000838152602090206118e5918101908301611935565b505050565b60206040519081016040526000815290565b604080519081016040526000808252602082015290565b6060604051908101604052806000815260200160008152602001600081525090565b61074191905b8082111561194f576000815560010161193b565b50905600a165627a7a72305820703dcd60a4d7982b41131b1be3e184c25ae70373c8afc1384b15475c7ed82d1a0029"

// DeployKyberNetworkProxy deploys a new Ethereum contract, binding an instance of KyberNetworkProxy to it.
func DeployKyberNetworkProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *KyberNetworkProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberNetworkProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KyberNetworkProxyBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KyberNetworkProxy{KyberNetworkProxyCaller: KyberNetworkProxyCaller{contract: contract}, KyberNetworkProxyTransactor: KyberNetworkProxyTransactor{contract: contract}, KyberNetworkProxyFilterer: KyberNetworkProxyFilterer{contract: contract}}, nil
}

// KyberNetworkProxy is an auto generated Go binding around an Ethereum contract.
type KyberNetworkProxy struct {
	KyberNetworkProxyCaller     // Read-only binding to the contract
	KyberNetworkProxyTransactor // Write-only binding to the contract
	KyberNetworkProxyFilterer   // Log filterer for contract events
}

// KyberNetworkProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberNetworkProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberNetworkProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberNetworkProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberNetworkProxySession struct {
	Contract     *KyberNetworkProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KyberNetworkProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberNetworkProxyCallerSession struct {
	Contract *KyberNetworkProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KyberNetworkProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberNetworkProxyTransactorSession struct {
	Contract     *KyberNetworkProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KyberNetworkProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberNetworkProxyRaw struct {
	Contract *KyberNetworkProxy // Generic contract binding to access the raw methods on
}

// KyberNetworkProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberNetworkProxyCallerRaw struct {
	Contract *KyberNetworkProxyCaller // Generic read-only contract binding to access the raw methods on
}

// KyberNetworkProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberNetworkProxyTransactorRaw struct {
	Contract *KyberNetworkProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberNetworkProxy creates a new instance of KyberNetworkProxy, bound to a specific deployed contract.
func NewKyberNetworkProxy(address common.Address, backend bind.ContractBackend) (*KyberNetworkProxy, error) {
	contract, err := bindKyberNetworkProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxy{KyberNetworkProxyCaller: KyberNetworkProxyCaller{contract: contract}, KyberNetworkProxyTransactor: KyberNetworkProxyTransactor{contract: contract}, KyberNetworkProxyFilterer: KyberNetworkProxyFilterer{contract: contract}}, nil
}

// NewKyberNetworkProxyCaller creates a new read-only instance of KyberNetworkProxy, bound to a specific deployed contract.
func NewKyberNetworkProxyCaller(address common.Address, caller bind.ContractCaller) (*KyberNetworkProxyCaller, error) {
	contract, err := bindKyberNetworkProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyCaller{contract: contract}, nil
}

// NewKyberNetworkProxyTransactor creates a new write-only instance of KyberNetworkProxy, bound to a specific deployed contract.
func NewKyberNetworkProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberNetworkProxyTransactor, error) {
	contract, err := bindKyberNetworkProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyTransactor{contract: contract}, nil
}

// NewKyberNetworkProxyFilterer creates a new log filterer instance of KyberNetworkProxy, bound to a specific deployed contract.
func NewKyberNetworkProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberNetworkProxyFilterer, error) {
	contract, err := bindKyberNetworkProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyFilterer{contract: contract}, nil
}

// bindKyberNetworkProxy binds a generic wrapper to an already deployed contract.
func bindKyberNetworkProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberNetworkProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberNetworkProxy *KyberNetworkProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberNetworkProxy.Contract.KyberNetworkProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberNetworkProxy *KyberNetworkProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.KyberNetworkProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberNetworkProxy *KyberNetworkProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.KyberNetworkProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberNetworkProxy *KyberNetworkProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberNetworkProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberNetworkProxy *KyberNetworkProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberNetworkProxy *KyberNetworkProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxySession) Admin() (common.Address, error) {
	return _KyberNetworkProxy.Contract.Admin(&_KyberNetworkProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) Admin() (common.Address, error) {
	return _KyberNetworkProxy.Contract.Admin(&_KyberNetworkProxy.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "enabled")
	return *ret0, err
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetworkProxy *KyberNetworkProxySession) Enabled() (bool, error) {
	return _KyberNetworkProxy.Contract.Enabled(&_KyberNetworkProxy.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) Enabled() (bool, error) {
	return _KyberNetworkProxy.Contract.Enabled(&_KyberNetworkProxy.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxySession) GetAlerters() ([]common.Address, error) {
	return _KyberNetworkProxy.Contract.GetAlerters(&_KyberNetworkProxy.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetAlerters() ([]common.Address, error) {
	return _KyberNetworkProxy.Contract.GetAlerters(&_KyberNetworkProxy.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetBalance(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getBalance", token, user)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetBalance(&_KyberNetworkProxy.CallOpts, token, user)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetBalance(&_KyberNetworkProxy.CallOpts, token, user)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetExpectedRate(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	ret := new(struct {
		ExpectedRate *big.Int
		SlippageRate *big.Int
	})
	out := ret
	err := _KyberNetworkProxy.contract.Call(opts, out, "getExpectedRate", src, dest, srcQty)
	return *ret, err
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetworkProxy *KyberNetworkProxySession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetworkProxy.Contract.GetExpectedRate(&_KyberNetworkProxy.CallOpts, src, dest, srcQty)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetworkProxy.Contract.GetExpectedRate(&_KyberNetworkProxy.CallOpts, src, dest, srcQty)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxySession) GetOperators() ([]common.Address, error) {
	return _KyberNetworkProxy.Contract.GetOperators(&_KyberNetworkProxy.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetOperators() ([]common.Address, error) {
	return _KyberNetworkProxy.Contract.GetOperators(&_KyberNetworkProxy.CallOpts)
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetUserCapInTokenWei(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getUserCapInTokenWei", user, token)
	return *ret0, err
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) GetUserCapInTokenWei(user common.Address, token common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetUserCapInTokenWei(&_KyberNetworkProxy.CallOpts, user, token)
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetUserCapInTokenWei(user common.Address, token common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetUserCapInTokenWei(&_KyberNetworkProxy.CallOpts, user, token)
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) GetUserCapInWei(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "getUserCapInWei", user)
	return *ret0, err
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) GetUserCapInWei(user common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetUserCapInWei(&_KyberNetworkProxy.CallOpts, user)
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) GetUserCapInWei(user common.Address) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.GetUserCapInWei(&_KyberNetworkProxy.CallOpts, user)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) Info(opts *bind.CallOpts, field [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "info", field)
	return *ret0, err
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) Info(field [32]byte) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.Info(&_KyberNetworkProxy.CallOpts, field)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) Info(field [32]byte) (*big.Int, error) {
	return _KyberNetworkProxy.Contract.Info(&_KyberNetworkProxy.CallOpts, field)
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) KyberNetworkContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "kyberNetworkContract")
	return *ret0, err
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxySession) KyberNetworkContract() (common.Address, error) {
	return _KyberNetworkProxy.Contract.KyberNetworkContract(&_KyberNetworkProxy.CallOpts)
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) KyberNetworkContract() (common.Address, error) {
	return _KyberNetworkProxy.Contract.KyberNetworkContract(&_KyberNetworkProxy.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "maxGasPrice")
	return *ret0, err
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) MaxGasPrice() (*big.Int, error) {
	return _KyberNetworkProxy.Contract.MaxGasPrice(&_KyberNetworkProxy.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) MaxGasPrice() (*big.Int, error) {
	return _KyberNetworkProxy.Contract.MaxGasPrice(&_KyberNetworkProxy.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetworkProxy.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxySession) PendingAdmin() (common.Address, error) {
	return _KyberNetworkProxy.Contract.PendingAdmin(&_KyberNetworkProxy.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetworkProxy *KyberNetworkProxyCallerSession) PendingAdmin() (common.Address, error) {
	return _KyberNetworkProxy.Contract.PendingAdmin(&_KyberNetworkProxy.CallOpts)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.AddAlerter(&_KyberNetworkProxy.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.AddAlerter(&_KyberNetworkProxy.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.AddOperator(&_KyberNetworkProxy.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.AddOperator(&_KyberNetworkProxy.TransactOpts, newOperator)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.ClaimAdmin(&_KyberNetworkProxy.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.ClaimAdmin(&_KyberNetworkProxy.TransactOpts)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.RemoveAlerter(&_KyberNetworkProxy.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.RemoveAlerter(&_KyberNetworkProxy.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.RemoveOperator(&_KyberNetworkProxy.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.RemoveOperator(&_KyberNetworkProxy.TransactOpts, operator)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) SetKyberNetworkContract(opts *bind.TransactOpts, _kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "setKyberNetworkContract", _kyberNetworkContract)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) SetKyberNetworkContract(_kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SetKyberNetworkContract(&_KyberNetworkProxy.TransactOpts, _kyberNetworkContract)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) SetKyberNetworkContract(_kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SetKyberNetworkContract(&_KyberNetworkProxy.TransactOpts, _kyberNetworkContract)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) SwapEtherToToken(opts *bind.TransactOpts, token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "swapEtherToToken", token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapEtherToToken(&_KyberNetworkProxy.TransactOpts, token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapEtherToToken(&_KyberNetworkProxy.TransactOpts, token, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) SwapTokenToEther(opts *bind.TransactOpts, token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "swapTokenToEther", token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapTokenToEther(&_KyberNetworkProxy.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapTokenToEther(&_KyberNetworkProxy.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) SwapTokenToToken(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "swapTokenToToken", src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapTokenToToken(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.SwapTokenToToken(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) Trade(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "trade", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.Trade(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.Trade(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) TradeWithHint(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "tradeWithHint", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxySession) TradeWithHint(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TradeWithHint(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) TradeWithHint(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TradeWithHint(&_KyberNetworkProxy.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TransferAdmin(&_KyberNetworkProxy.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TransferAdmin(&_KyberNetworkProxy.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TransferAdminQuickly(&_KyberNetworkProxy.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.TransferAdminQuickly(&_KyberNetworkProxy.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.WithdrawEther(&_KyberNetworkProxy.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.WithdrawEther(&_KyberNetworkProxy.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxySession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.WithdrawToken(&_KyberNetworkProxy.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetworkProxy *KyberNetworkProxyTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetworkProxy.Contract.WithdrawToken(&_KyberNetworkProxy.TransactOpts, token, amount, sendTo)
}

// KyberNetworkProxyAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyAdminClaimedIterator struct {
	Event *KyberNetworkProxyAdminClaimed // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyAdminClaimed)
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
		it.Event = new(KyberNetworkProxyAdminClaimed)
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
func (it *KyberNetworkProxyAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyAdminClaimed represents a AdminClaimed event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*KyberNetworkProxyAdminClaimedIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyAdminClaimedIterator{contract: _KyberNetworkProxy.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyAdminClaimed)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseAdminClaimed(log types.Log) (*KyberNetworkProxyAdminClaimed, error) {
	event := new(KyberNetworkProxyAdminClaimed)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyAlerterAddedIterator struct {
	Event *KyberNetworkProxyAlerterAdded // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyAlerterAdded)
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
		it.Event = new(KyberNetworkProxyAlerterAdded)
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
func (it *KyberNetworkProxyAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyAlerterAdded represents a AlerterAdded event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*KyberNetworkProxyAlerterAddedIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyAlerterAddedIterator{contract: _KyberNetworkProxy.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyAlerterAdded)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseAlerterAdded(log types.Log) (*KyberNetworkProxyAlerterAdded, error) {
	event := new(KyberNetworkProxyAlerterAdded)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyEtherWithdrawIterator struct {
	Event *KyberNetworkProxyEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyEtherWithdraw)
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
		it.Event = new(KyberNetworkProxyEtherWithdraw)
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
func (it *KyberNetworkProxyEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyEtherWithdraw represents a EtherWithdraw event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*KyberNetworkProxyEtherWithdrawIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyEtherWithdrawIterator{contract: _KyberNetworkProxy.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyEtherWithdraw)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseEtherWithdraw(log types.Log) (*KyberNetworkProxyEtherWithdraw, error) {
	event := new(KyberNetworkProxyEtherWithdraw)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyExecuteTradeIterator is returned from FilterExecuteTrade and is used to iterate over the raw logs and unpacked data for ExecuteTrade events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyExecuteTradeIterator struct {
	Event *KyberNetworkProxyExecuteTrade // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyExecuteTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyExecuteTrade)
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
		it.Event = new(KyberNetworkProxyExecuteTrade)
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
func (it *KyberNetworkProxyExecuteTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyExecuteTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyExecuteTrade represents a ExecuteTrade event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyExecuteTrade struct {
	Trader           common.Address
	Src              common.Address
	Dest             common.Address
	ActualSrcAmount  *big.Int
	ActualDestAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecuteTrade is a free log retrieval operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterExecuteTrade(opts *bind.FilterOpts, trader []common.Address) (*KyberNetworkProxyExecuteTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "ExecuteTrade", traderRule)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyExecuteTradeIterator{contract: _KyberNetworkProxy.contract, event: "ExecuteTrade", logs: logs, sub: sub}, nil
}

// WatchExecuteTrade is a free log subscription operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchExecuteTrade(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyExecuteTrade, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "ExecuteTrade", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyExecuteTrade)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "ExecuteTrade", log); err != nil {
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

// ParseExecuteTrade is a log parse operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseExecuteTrade(log types.Log) (*KyberNetworkProxyExecuteTrade, error) {
	event := new(KyberNetworkProxyExecuteTrade)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "ExecuteTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyKyberNetworkSetIterator is returned from FilterKyberNetworkSet and is used to iterate over the raw logs and unpacked data for KyberNetworkSet events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyKyberNetworkSetIterator struct {
	Event *KyberNetworkProxyKyberNetworkSet // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyKyberNetworkSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyKyberNetworkSet)
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
		it.Event = new(KyberNetworkProxyKyberNetworkSet)
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
func (it *KyberNetworkProxyKyberNetworkSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyKyberNetworkSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyKyberNetworkSet represents a KyberNetworkSet event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyKyberNetworkSet struct {
	NewNetworkContract common.Address
	OldNetworkContract common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterKyberNetworkSet is a free log retrieval operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterKyberNetworkSet(opts *bind.FilterOpts) (*KyberNetworkProxyKyberNetworkSetIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "KyberNetworkSet")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyKyberNetworkSetIterator{contract: _KyberNetworkProxy.contract, event: "KyberNetworkSet", logs: logs, sub: sub}, nil
}

// WatchKyberNetworkSet is a free log subscription operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchKyberNetworkSet(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyKyberNetworkSet) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "KyberNetworkSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyKyberNetworkSet)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "KyberNetworkSet", log); err != nil {
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

// ParseKyberNetworkSet is a log parse operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseKyberNetworkSet(log types.Log) (*KyberNetworkProxyKyberNetworkSet, error) {
	event := new(KyberNetworkProxyKyberNetworkSet)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "KyberNetworkSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyOperatorAddedIterator struct {
	Event *KyberNetworkProxyOperatorAdded // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyOperatorAdded)
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
		it.Event = new(KyberNetworkProxyOperatorAdded)
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
func (it *KyberNetworkProxyOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyOperatorAdded represents a OperatorAdded event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*KyberNetworkProxyOperatorAddedIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyOperatorAddedIterator{contract: _KyberNetworkProxy.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyOperatorAdded)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseOperatorAdded(log types.Log) (*KyberNetworkProxyOperatorAdded, error) {
	event := new(KyberNetworkProxyOperatorAdded)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyTokenWithdrawIterator struct {
	Event *KyberNetworkProxyTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyTokenWithdraw)
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
		it.Event = new(KyberNetworkProxyTokenWithdraw)
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
func (it *KyberNetworkProxyTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyTokenWithdraw represents a TokenWithdraw event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*KyberNetworkProxyTokenWithdrawIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyTokenWithdrawIterator{contract: _KyberNetworkProxy.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyTokenWithdraw)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseTokenWithdraw(log types.Log) (*KyberNetworkProxyTokenWithdraw, error) {
	event := new(KyberNetworkProxyTokenWithdraw)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkProxyTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the KyberNetworkProxy contract.
type KyberNetworkProxyTransferAdminPendingIterator struct {
	Event *KyberNetworkProxyTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *KyberNetworkProxyTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkProxyTransferAdminPending)
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
		it.Event = new(KyberNetworkProxyTransferAdminPending)
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
func (it *KyberNetworkProxyTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkProxyTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkProxyTransferAdminPending represents a TransferAdminPending event raised by the KyberNetworkProxy contract.
type KyberNetworkProxyTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*KyberNetworkProxyTransferAdminPendingIterator, error) {

	logs, sub, err := _KyberNetworkProxy.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkProxyTransferAdminPendingIterator{contract: _KyberNetworkProxy.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *KyberNetworkProxyTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _KyberNetworkProxy.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkProxyTransferAdminPending)
				if err := _KyberNetworkProxy.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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
func (_KyberNetworkProxy *KyberNetworkProxyFilterer) ParseTransferAdminPending(log types.Log) (*KyberNetworkProxyTransferAdminPending, error) {
	event := new(KyberNetworkProxyTransferAdminPending)
	if err := _KyberNetworkProxy.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}
