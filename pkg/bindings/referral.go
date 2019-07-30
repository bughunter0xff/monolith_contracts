// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ReferralABI is the input ABI used to generate the binding from.
const ReferralABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referralTokens\",\"type\":\"uint256[]\"}],\"name\":\"transferReferralBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"referralIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferReferralToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBonus\",\"type\":\"uint256\"}],\"name\":\"setReferralBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"firstOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TKNReferralBonus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"issueReferralTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintReferralTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_totalSuply\",\"type\":\"uint256\"},{\"name\":\"_TKNAddress\",\"type\":\"address\"},{\"name\":\"_referralBonus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newSupply\",\"type\":\"uint256\"}],\"name\":\"MintedReferralTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"IssuedReferralTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferredReferralBonus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newBonus\",\"type\":\"uint256\"}],\"name\":\"SetBonus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ReferralBin is the compiled bytecode used for deploying new contracts.
const ReferralBin = `608060405234801561001057600080fd5b5060405162001beb38038062001beb8339818101604052606081101561003557600080fd5b508051602082015160409092015190919033600061007b7f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b036101aa16565b6100ad7f80ac58cd000000000000000000000000000000000000000000000000000000006001600160e01b036101aa16565b600580546001600160a01b0319166001600160a01b0384161760ff60a01b1916740100000000000000000000000000000000000000008315158102919091179182905560ff91041661013657604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a15050600692909255600a80546001600160a01b0319166001600160a01b0392909216919091179055600855610278565b7fffffffff00000000000000000000000000000000000000000000000000000000808216141561023b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b61196380620002886000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c80637392b3b3116100de578063a22cb46511610097578063b88d4fde11610071578063b88d4fde14610462578063bb6acb9314610528578063c260d95d14610554578063e985e9c51461057157610173565b8063a22cb465146103fe578063abfe4c0e1461042c578063b242e5341461043457610173565b80637392b3b3146103805780637f5d6139146103885780638d75fe05146103b45780638da5cb5b146103bc5780639a037aa0146103c45780639c7e5a8f146103e157610173565b806332fad9d81161013057806332fad9d81461027257806342842e0e1461028f5780636352211e146102c557806368d27dbd146102e257806370a0823114610352578063715018a61461037857610173565b806301ffc9a714610178578063081812fc146101b3578063095ea7b3146101ec57806318160ddd1461021a5780632121dc751461023457806323b872dd1461023c575b600080fd5b61019f6004803603602081101561018e57600080fd5b50356001600160e01b03191661059f565b604080519115158252519081900360200190f35b6101d0600480360360208110156101c957600080fd5b50356105be565b604080516001600160a01b039092168252519081900360200190f35b6102186004803603604081101561020257600080fd5b506001600160a01b038135169060200135610620565b005b610222610731565b60408051918252519081900360200190f35b61019f610737565b6102186004803603606081101561025257600080fd5b506001600160a01b03813581169160208101359091169060400135610747565b61019f6004803603602081101561028857600080fd5b503561079c565b610218600480360360608110156102a557600080fd5b506001600160a01b038135811691602081013590911690604001356107b1565b6101d0600480360360208110156102db57600080fd5b50356107cc565b610218600480360360208110156102f857600080fd5b81019060208101813564010000000081111561031357600080fd5b82018360208201111561032557600080fd5b8035906020019184602083028401116401000000008311171561034757600080fd5b509092509050610826565b6102226004803603602081101561036857600080fd5b50356001600160a01b03166109be565b610218610a26565b610222610b19565b6102186004803603604081101561039e57600080fd5b506001600160a01b038135169060200135610b1f565b610222610b2e565b6101d0610b34565b610218600480360360208110156103da57600080fd5b5035610b43565b6101d0600480360360208110156103f757600080fd5b5035610bce565b6102186004803603604081101561041457600080fd5b506001600160a01b0381351690602001351515610be9565b610222610cb5565b6102186004803603604081101561044a57600080fd5b506001600160a01b0381351690602001351515610cbb565b6102186004803603608081101561047857600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156104b357600080fd5b8201836020820111156104c557600080fd5b803590602001918460018302840111640100000000831117156104e757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e6a945050505050565b6102186004803603604081101561053e57600080fd5b506001600160a01b038135169060200135610ec2565b6102186004803603602081101561056a57600080fd5b5035611057565b61019f6004803603604081101561058757600080fd5b506001600160a01b03813581169160200135166111a6565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006105c9826111d4565b6106045760405162461bcd60e51b815260040180806020018281038252602c815260200180611843602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b600061062b826107cc565b9050806001600160a01b0316836001600160a01b0316141561067e5760405162461bcd60e51b81526004018080602001828103825260218152602001806118986021913960400191505060405180910390fd5b336001600160a01b038216148061069a575061069a81336111a6565b6106d55760405162461bcd60e51b81526004018080602001828103825260388152602001806117b86038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60065481565b600554600160a01b900460ff1690565b61075133826111f1565b61078c5760405162461bcd60e51b81526004018080602001828103825260318152602001806118dc6031913960400191505060405180910390fd5b610797838383611295565b505050565b600b6020526000908152604090205460ff1681565b61079783838360405180602001604052806000815250610e6a565b6000818152600160205260408120546001600160a01b0316806108205760405162461bcd60e51b815260040180806020018281038252602981526020018061181a6029913960400191505060405180910390fd5b92915050565b61082f336113d9565b61086e576040805162461bcd60e51b8152602060048201526016602482015260008051602061176c833981519152604482015290519081900360640190fd5b60005b8181101561079757600083838381811061088757fe5b602090810292909201356000818152600b9093526040909220549192505060ff166109b5576000818152600b60209081526040808320805460ff19166001179055600a54600c835281842054600854835163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b15801561092557600080fd5b505af1158015610939573d6000803e3d6000fd5b505050506040513d602081101561094f57600080fd5b50506000818152600c60209081526040918290205460085483513381526001600160a01b039092169282019290925280830191909152905182917f40cf0b5afee61054126d1b4397c6f217e1f0c4fc65ebe84f382830b9006036fd919081900360600190a25b50600101610871565b60006001600160a01b038216610a055760405162461bcd60e51b815260040180806020018281038252602a8152602001806117f0602a913960400191505060405180910390fd5b6001600160a01b0382166000908152600360205260409020610820906113ed565b610a2f336113d9565b610a6e576040805162461bcd60e51b8152602060048201526016602482015260008051602061176c833981519152604482015290519081900360640190fd5b600554600160a01b900460ff16610acc576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600580546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b60075481565b610b2a338383611295565b5050565b60095481565b6005546001600160a01b031690565b610b4c336113d9565b610b8b576040805162461bcd60e51b8152602060048201526016602482015260008051602061176c833981519152604482015290519081900360640190fd5b6008819055604080513381526020810183905281517fce20f96597c75db9bc70de06a2bbb2ccd70ef20f0677e190486643e9c48354d3929181900390910190a150565b600c602052600090815260409020546001600160a01b031681565b6001600160a01b038216331415610c47576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b60085481565b610cc4336113d9565b610d03576040805162461bcd60e51b8152602060048201526016602482015260008051602061176c833981519152604482015290519081900360640190fd5b600554600160a01b900460ff16610d61576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b038216610da65760405162461bcd60e51b81526004018080602001828103825260238152602001806118b96023913960400191505060405180910390fd5b6005805460ff60a01b1916600160a01b8315150217905580610dff57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600554604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600580546001600160a01b0319166001600160a01b0392909216919091179055565b610e75848484610747565b610e81848484846113f1565b610ebc5760405162461bcd60e51b81526004018080602001828103825260328152602001806117166032913960400191505060405180910390fd5b50505050565b610ecb336113d9565b610f0a576040805162461bcd60e51b8152602060048201526016602482015260008051602061176c833981519152604482015290519081900360640190fd5b6005811115610f60576040805162461bcd60e51b815260206004820152601f60248201527f746f6f206d616e7920726566657272616c20746f6b656e7320676976656e2100604482015290519081900360640190fd5b600081610f6f57506001610f72565b50805b60095481600754011115610fb75760405162461bcd60e51b815260040180806020018281038252602281526020018061190d6022913960400191505060405180910390fd5b6007545b816007540181101561100157610fd2338583611295565b6000818152600c6020526040902080546001600160a01b0319166001600160a01b038616179055600101610fbb565b506007805482019055604080513381526001600160a01b038516602082015280820183905290517f933f1bfdc6b87a5cdf63eba5c3ce0355fdeecd54c1cd07f5bcd0b89f4248ccf49181900360600190a1505050565b611060336113d9565b61109f576040805162461bcd60e51b8152602060048201526016602482015260008051602061176c833981519152604482015290519081900360640190fd5b6009548082019081116110ef576040805162461bcd60e51b81526020600482015260136024820152721bdd995c999b1bddc81bdc880c081a5b9c1d5d606a1b604482015290519081900360640190fd5b60065481111561113e576040805162461bcd60e51b81526020600482015260156024820152741d1bdd185b081cdd5c1c1b1e48195e18d959591959605a1b604482015290519081900360640190fd5b6009545b8181101561115c576111543382611524565b600101611142565b506009819055604080513381526020810184905280820183905290517f70a3dc5168083fd1722327506f78f85af6d20f327c6b92dc4f5081543b03b9319181900360600190a15050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b60006111fc826111d4565b6112375760405162461bcd60e51b815260040180806020018281038252602c81526020018061178c602c913960400191505060405180910390fd5b6000611242836107cc565b9050806001600160a01b0316846001600160a01b0316148061127d5750836001600160a01b0316611272846105be565b6001600160a01b0316145b8061128d575061128d81856111a6565b949350505050565b826001600160a01b03166112a8826107cc565b6001600160a01b0316146112ed5760405162461bcd60e51b815260040180806020018281038252602981526020018061186f6029913960400191505060405180910390fd5b6001600160a01b0382166113325760405162461bcd60e51b81526004018080602001828103825260248152602001806117486024913960400191505060405180910390fd5b61133b81611655565b6001600160a01b038316600090815260036020526040902061135c90611692565b6001600160a01b038216600090815260036020526040902061137d906116a9565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6005546001600160a01b0390811691161490565b5490565b6000611405846001600160a01b03166116b2565b6114115750600161128d565b604051630a85bd0160e11b815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a169463150b7a029490938c938b938b939260a4019060208501908083838e5b8381101561148b578181015183820152602001611473565b50505050905090810190601f1680156114b85780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156114da57600080fd5b505af11580156114ee573d6000803e3d6000fd5b505050506040513d602081101561150457600080fd5b50516001600160e01b031916630a85bd0160e11b14915050949350505050565b6001600160a01b03821661157f576040805162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015290519081900360640190fd5b611588816111d4565b156115da576040805162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015290519081900360640190fd5b600081815260016020908152604080832080546001600160a01b0319166001600160a01b038716908117909155835260039091529020611619906116a9565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000818152600260205260409020546001600160a01b03161561168f57600081815260026020526040902080546001600160a01b03191690555b50565b80546116a590600163ffffffff6116b816565b9055565b80546001019055565b3b151590565b60008282111561170f576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f206164647265737373656e646572206973206e6f7420616e206f776e6572000000000000000000004552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65726f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573734552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564746f6b656e7320657863656564207468652063757272656e742073757070706c7921a265627a7a723058209006671ed6aa2e59f1a9490d1443d7de8031fa1e769b2cc9915502f2dd2f523064736f6c634300050a0032`

// DeployReferral deploys a new Ethereum contract, binding an instance of Referral to it.
func DeployReferral(auth *bind.TransactOpts, backend bind.ContractBackend, _totalSuply *big.Int, _TKNAddress common.Address, _referralBonus *big.Int) (common.Address, *types.Transaction, *Referral, error) {
	parsed, err := abi.JSON(strings.NewReader(ReferralABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReferralBin), backend, _totalSuply, _TKNAddress, _referralBonus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Referral{ReferralCaller: ReferralCaller{contract: contract}, ReferralTransactor: ReferralTransactor{contract: contract}, ReferralFilterer: ReferralFilterer{contract: contract}}, nil
}

// Referral is an auto generated Go binding around an Ethereum contract.
type Referral struct {
	ReferralCaller     // Read-only binding to the contract
	ReferralTransactor // Write-only binding to the contract
	ReferralFilterer   // Log filterer for contract events
}

// ReferralCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReferralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReferralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReferralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReferralSession struct {
	Contract     *Referral         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReferralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReferralCallerSession struct {
	Contract *ReferralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ReferralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReferralTransactorSession struct {
	Contract     *ReferralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReferralRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReferralRaw struct {
	Contract *Referral // Generic contract binding to access the raw methods on
}

// ReferralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReferralCallerRaw struct {
	Contract *ReferralCaller // Generic read-only contract binding to access the raw methods on
}

// ReferralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReferralTransactorRaw struct {
	Contract *ReferralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReferral creates a new instance of Referral, bound to a specific deployed contract.
func NewReferral(address common.Address, backend bind.ContractBackend) (*Referral, error) {
	contract, err := bindReferral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Referral{ReferralCaller: ReferralCaller{contract: contract}, ReferralTransactor: ReferralTransactor{contract: contract}, ReferralFilterer: ReferralFilterer{contract: contract}}, nil
}

// NewReferralCaller creates a new read-only instance of Referral, bound to a specific deployed contract.
func NewReferralCaller(address common.Address, caller bind.ContractCaller) (*ReferralCaller, error) {
	contract, err := bindReferral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralCaller{contract: contract}, nil
}

// NewReferralTransactor creates a new write-only instance of Referral, bound to a specific deployed contract.
func NewReferralTransactor(address common.Address, transactor bind.ContractTransactor) (*ReferralTransactor, error) {
	contract, err := bindReferral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralTransactor{contract: contract}, nil
}

// NewReferralFilterer creates a new log filterer instance of Referral, bound to a specific deployed contract.
func NewReferralFilterer(address common.Address, filterer bind.ContractFilterer) (*ReferralFilterer, error) {
	contract, err := bindReferral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReferralFilterer{contract: contract}, nil
}

// bindReferral binds a generic wrapper to an already deployed contract.
func bindReferral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReferralABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Referral *ReferralRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Referral.Contract.ReferralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Referral *ReferralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Referral.Contract.ReferralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Referral *ReferralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Referral.Contract.ReferralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Referral *ReferralCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Referral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Referral *ReferralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Referral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Referral *ReferralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Referral.Contract.contract.Transact(opts, method, params...)
}

// TKNReferralBonus is a free data retrieval call binding the contract method 0xabfe4c0e.
//
// Solidity: function TKNReferralBonus() constant returns(uint256)
func (_Referral *ReferralCaller) TKNReferralBonus(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "TKNReferralBonus")
	return *ret0, err
}

// TKNReferralBonus is a free data retrieval call binding the contract method 0xabfe4c0e.
//
// Solidity: function TKNReferralBonus() constant returns(uint256)
func (_Referral *ReferralSession) TKNReferralBonus() (*big.Int, error) {
	return _Referral.Contract.TKNReferralBonus(&_Referral.CallOpts)
}

// TKNReferralBonus is a free data retrieval call binding the contract method 0xabfe4c0e.
//
// Solidity: function TKNReferralBonus() constant returns(uint256)
func (_Referral *ReferralCallerSession) TKNReferralBonus() (*big.Int, error) {
	return _Referral.Contract.TKNReferralBonus(&_Referral.CallOpts)
}

// Activated is a free data retrieval call binding the contract method 0x32fad9d8.
//
// Solidity: function activated(uint256 ) constant returns(bool)
func (_Referral *ReferralCaller) Activated(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "activated", arg0)
	return *ret0, err
}

// Activated is a free data retrieval call binding the contract method 0x32fad9d8.
//
// Solidity: function activated(uint256 ) constant returns(bool)
func (_Referral *ReferralSession) Activated(arg0 *big.Int) (bool, error) {
	return _Referral.Contract.Activated(&_Referral.CallOpts, arg0)
}

// Activated is a free data retrieval call binding the contract method 0x32fad9d8.
//
// Solidity: function activated(uint256 ) constant returns(bool)
func (_Referral *ReferralCallerSession) Activated(arg0 *big.Int) (bool, error) {
	return _Referral.Contract.Activated(&_Referral.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Referral *ReferralCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Referral *ReferralSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Referral.Contract.BalanceOf(&_Referral.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Referral *ReferralCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Referral.Contract.BalanceOf(&_Referral.CallOpts, owner)
}

// FirstOwner is a free data retrieval call binding the contract method 0x9c7e5a8f.
//
// Solidity: function firstOwner(uint256 ) constant returns(address)
func (_Referral *ReferralCaller) FirstOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "firstOwner", arg0)
	return *ret0, err
}

// FirstOwner is a free data retrieval call binding the contract method 0x9c7e5a8f.
//
// Solidity: function firstOwner(uint256 ) constant returns(address)
func (_Referral *ReferralSession) FirstOwner(arg0 *big.Int) (common.Address, error) {
	return _Referral.Contract.FirstOwner(&_Referral.CallOpts, arg0)
}

// FirstOwner is a free data retrieval call binding the contract method 0x9c7e5a8f.
//
// Solidity: function firstOwner(uint256 ) constant returns(address)
func (_Referral *ReferralCallerSession) FirstOwner(arg0 *big.Int) (common.Address, error) {
	return _Referral.Contract.FirstOwner(&_Referral.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_Referral *ReferralCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_Referral *ReferralSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Referral.Contract.GetApproved(&_Referral.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_Referral *ReferralCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Referral.Contract.GetApproved(&_Referral.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_Referral *ReferralCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_Referral *ReferralSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Referral.Contract.IsApprovedForAll(&_Referral.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_Referral *ReferralCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Referral.Contract.IsApprovedForAll(&_Referral.CallOpts, owner, operator)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Referral *ReferralCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Referral *ReferralSession) IsTransferable() (bool, error) {
	return _Referral.Contract.IsTransferable(&_Referral.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_Referral *ReferralCallerSession) IsTransferable() (bool, error) {
	return _Referral.Contract.IsTransferable(&_Referral.CallOpts)
}

// MintedTokens is a free data retrieval call binding the contract method 0x8d75fe05.
//
// Solidity: function mintedTokens() constant returns(uint256)
func (_Referral *ReferralCaller) MintedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "mintedTokens")
	return *ret0, err
}

// MintedTokens is a free data retrieval call binding the contract method 0x8d75fe05.
//
// Solidity: function mintedTokens() constant returns(uint256)
func (_Referral *ReferralSession) MintedTokens() (*big.Int, error) {
	return _Referral.Contract.MintedTokens(&_Referral.CallOpts)
}

// MintedTokens is a free data retrieval call binding the contract method 0x8d75fe05.
//
// Solidity: function mintedTokens() constant returns(uint256)
func (_Referral *ReferralCallerSession) MintedTokens() (*big.Int, error) {
	return _Referral.Contract.MintedTokens(&_Referral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Referral *ReferralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Referral *ReferralSession) Owner() (common.Address, error) {
	return _Referral.Contract.Owner(&_Referral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Referral *ReferralCallerSession) Owner() (common.Address, error) {
	return _Referral.Contract.Owner(&_Referral.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_Referral *ReferralCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_Referral *ReferralSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Referral.Contract.OwnerOf(&_Referral.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_Referral *ReferralCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Referral.Contract.OwnerOf(&_Referral.CallOpts, tokenId)
}

// ReferralIndex is a free data retrieval call binding the contract method 0x7392b3b3.
//
// Solidity: function referralIndex() constant returns(uint256)
func (_Referral *ReferralCaller) ReferralIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "referralIndex")
	return *ret0, err
}

// ReferralIndex is a free data retrieval call binding the contract method 0x7392b3b3.
//
// Solidity: function referralIndex() constant returns(uint256)
func (_Referral *ReferralSession) ReferralIndex() (*big.Int, error) {
	return _Referral.Contract.ReferralIndex(&_Referral.CallOpts)
}

// ReferralIndex is a free data retrieval call binding the contract method 0x7392b3b3.
//
// Solidity: function referralIndex() constant returns(uint256)
func (_Referral *ReferralCallerSession) ReferralIndex() (*big.Int, error) {
	return _Referral.Contract.ReferralIndex(&_Referral.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Referral *ReferralCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Referral *ReferralSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Referral.Contract.SupportsInterface(&_Referral.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Referral *ReferralCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Referral.Contract.SupportsInterface(&_Referral.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Referral *ReferralCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Referral *ReferralSession) TotalSupply() (*big.Int, error) {
	return _Referral.Contract.TotalSupply(&_Referral.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Referral *ReferralCallerSession) TotalSupply() (*big.Int, error) {
	return _Referral.Contract.TotalSupply(&_Referral.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Referral *ReferralTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Referral *ReferralSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.Approve(&_Referral.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Referral *ReferralTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.Approve(&_Referral.TransactOpts, to, tokenId)
}

// IssueReferralTokens is a paid mutator transaction binding the contract method 0xbb6acb93.
//
// Solidity: function issueReferralTokens(address _to, uint256 _amount) returns()
func (_Referral *ReferralTransactor) IssueReferralTokens(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "issueReferralTokens", _to, _amount)
}

// IssueReferralTokens is a paid mutator transaction binding the contract method 0xbb6acb93.
//
// Solidity: function issueReferralTokens(address _to, uint256 _amount) returns()
func (_Referral *ReferralSession) IssueReferralTokens(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.IssueReferralTokens(&_Referral.TransactOpts, _to, _amount)
}

// IssueReferralTokens is a paid mutator transaction binding the contract method 0xbb6acb93.
//
// Solidity: function issueReferralTokens(address _to, uint256 _amount) returns()
func (_Referral *ReferralTransactorSession) IssueReferralTokens(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.IssueReferralTokens(&_Referral.TransactOpts, _to, _amount)
}

// MintReferralTokens is a paid mutator transaction binding the contract method 0xc260d95d.
//
// Solidity: function mintReferralTokens(uint256 _amount) returns()
func (_Referral *ReferralTransactor) MintReferralTokens(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "mintReferralTokens", _amount)
}

// MintReferralTokens is a paid mutator transaction binding the contract method 0xc260d95d.
//
// Solidity: function mintReferralTokens(uint256 _amount) returns()
func (_Referral *ReferralSession) MintReferralTokens(_amount *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.MintReferralTokens(&_Referral.TransactOpts, _amount)
}

// MintReferralTokens is a paid mutator transaction binding the contract method 0xc260d95d.
//
// Solidity: function mintReferralTokens(uint256 _amount) returns()
func (_Referral *ReferralTransactorSession) MintReferralTokens(_amount *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.MintReferralTokens(&_Referral.TransactOpts, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Referral *ReferralTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Referral *ReferralSession) RenounceOwnership() (*types.Transaction, error) {
	return _Referral.Contract.RenounceOwnership(&_Referral.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Referral *ReferralTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Referral.Contract.RenounceOwnership(&_Referral.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Referral *ReferralTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "safeTransferFrom", from, to, tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Referral *ReferralSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Referral.Contract.SafeTransferFrom(&_Referral.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Referral *ReferralTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Referral.Contract.SafeTransferFrom(&_Referral.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Referral *ReferralTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Referral *ReferralSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _Referral.Contract.SetApprovalForAll(&_Referral.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Referral *ReferralTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _Referral.Contract.SetApprovalForAll(&_Referral.TransactOpts, to, approved)
}

// SetReferralBonus is a paid mutator transaction binding the contract method 0x9a037aa0.
//
// Solidity: function setReferralBonus(uint256 newBonus) returns()
func (_Referral *ReferralTransactor) SetReferralBonus(opts *bind.TransactOpts, newBonus *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "setReferralBonus", newBonus)
}

// SetReferralBonus is a paid mutator transaction binding the contract method 0x9a037aa0.
//
// Solidity: function setReferralBonus(uint256 newBonus) returns()
func (_Referral *ReferralSession) SetReferralBonus(newBonus *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.SetReferralBonus(&_Referral.TransactOpts, newBonus)
}

// SetReferralBonus is a paid mutator transaction binding the contract method 0x9a037aa0.
//
// Solidity: function setReferralBonus(uint256 newBonus) returns()
func (_Referral *ReferralTransactorSession) SetReferralBonus(newBonus *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.SetReferralBonus(&_Referral.TransactOpts, newBonus)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Referral *ReferralTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Referral *ReferralSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferFrom(&_Referral.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Referral *ReferralTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferFrom(&_Referral.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Referral *ReferralTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Referral *ReferralSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Referral.Contract.TransferOwnership(&_Referral.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_Referral *ReferralTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _Referral.Contract.TransferOwnership(&_Referral.TransactOpts, _account, _transferable)
}

// TransferReferralBonus is a paid mutator transaction binding the contract method 0x68d27dbd.
//
// Solidity: function transferReferralBonus(uint256[] _referralTokens) returns()
func (_Referral *ReferralTransactor) TransferReferralBonus(opts *bind.TransactOpts, _referralTokens []*big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "transferReferralBonus", _referralTokens)
}

// TransferReferralBonus is a paid mutator transaction binding the contract method 0x68d27dbd.
//
// Solidity: function transferReferralBonus(uint256[] _referralTokens) returns()
func (_Referral *ReferralSession) TransferReferralBonus(_referralTokens []*big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferReferralBonus(&_Referral.TransactOpts, _referralTokens)
}

// TransferReferralBonus is a paid mutator transaction binding the contract method 0x68d27dbd.
//
// Solidity: function transferReferralBonus(uint256[] _referralTokens) returns()
func (_Referral *ReferralTransactorSession) TransferReferralBonus(_referralTokens []*big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferReferralBonus(&_Referral.TransactOpts, _referralTokens)
}

// TransferReferralToken is a paid mutator transaction binding the contract method 0x7f5d6139.
//
// Solidity: function transferReferralToken(address _to, uint256 _tokenId) returns()
func (_Referral *ReferralTransactor) TransferReferralToken(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "transferReferralToken", _to, _tokenId)
}

// TransferReferralToken is a paid mutator transaction binding the contract method 0x7f5d6139.
//
// Solidity: function transferReferralToken(address _to, uint256 _tokenId) returns()
func (_Referral *ReferralSession) TransferReferralToken(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferReferralToken(&_Referral.TransactOpts, _to, _tokenId)
}

// TransferReferralToken is a paid mutator transaction binding the contract method 0x7f5d6139.
//
// Solidity: function transferReferralToken(address _to, uint256 _tokenId) returns()
func (_Referral *ReferralTransactorSession) TransferReferralToken(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferReferralToken(&_Referral.TransactOpts, _to, _tokenId)
}

// ReferralApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Referral contract.
type ReferralApprovalIterator struct {
	Event *ReferralApproval // Event containing the contract specifics and raw log

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
func (it *ReferralApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralApproval)
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
		it.Event = new(ReferralApproval)
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
func (it *ReferralApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralApproval represents a Approval event raised by the Referral contract.
type ReferralApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Referral *ReferralFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ReferralApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Referral.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ReferralApprovalIterator{contract: _Referral.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Referral *ReferralFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ReferralApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Referral.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralApproval)
				if err := _Referral.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ReferralApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Referral contract.
type ReferralApprovalForAllIterator struct {
	Event *ReferralApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ReferralApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralApprovalForAll)
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
		it.Event = new(ReferralApprovalForAll)
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
func (it *ReferralApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralApprovalForAll represents a ApprovalForAll event raised by the Referral contract.
type ReferralApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Referral *ReferralFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ReferralApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Referral.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ReferralApprovalForAllIterator{contract: _Referral.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Referral *ReferralFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ReferralApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Referral.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralApprovalForAll)
				if err := _Referral.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ReferralIssuedReferralTokensIterator is returned from FilterIssuedReferralTokens and is used to iterate over the raw logs and unpacked data for IssuedReferralTokens events raised by the Referral contract.
type ReferralIssuedReferralTokensIterator struct {
	Event *ReferralIssuedReferralTokens // Event containing the contract specifics and raw log

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
func (it *ReferralIssuedReferralTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralIssuedReferralTokens)
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
		it.Event = new(ReferralIssuedReferralTokens)
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
func (it *ReferralIssuedReferralTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralIssuedReferralTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralIssuedReferralTokens represents a IssuedReferralTokens event raised by the Referral contract.
type ReferralIssuedReferralTokens struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuedReferralTokens is a free log retrieval operation binding the contract event 0x933f1bfdc6b87a5cdf63eba5c3ce0355fdeecd54c1cd07f5bcd0b89f4248ccf4.
//
// Solidity: event IssuedReferralTokens(address _from, address _to, uint256 _amount)
func (_Referral *ReferralFilterer) FilterIssuedReferralTokens(opts *bind.FilterOpts) (*ReferralIssuedReferralTokensIterator, error) {

	logs, sub, err := _Referral.contract.FilterLogs(opts, "IssuedReferralTokens")
	if err != nil {
		return nil, err
	}
	return &ReferralIssuedReferralTokensIterator{contract: _Referral.contract, event: "IssuedReferralTokens", logs: logs, sub: sub}, nil
}

// WatchIssuedReferralTokens is a free log subscription operation binding the contract event 0x933f1bfdc6b87a5cdf63eba5c3ce0355fdeecd54c1cd07f5bcd0b89f4248ccf4.
//
// Solidity: event IssuedReferralTokens(address _from, address _to, uint256 _amount)
func (_Referral *ReferralFilterer) WatchIssuedReferralTokens(opts *bind.WatchOpts, sink chan<- *ReferralIssuedReferralTokens) (event.Subscription, error) {

	logs, sub, err := _Referral.contract.WatchLogs(opts, "IssuedReferralTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralIssuedReferralTokens)
				if err := _Referral.contract.UnpackLog(event, "IssuedReferralTokens", log); err != nil {
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

// ReferralLockedOwnershipIterator is returned from FilterLockedOwnership and is used to iterate over the raw logs and unpacked data for LockedOwnership events raised by the Referral contract.
type ReferralLockedOwnershipIterator struct {
	Event *ReferralLockedOwnership // Event containing the contract specifics and raw log

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
func (it *ReferralLockedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralLockedOwnership)
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
		it.Event = new(ReferralLockedOwnership)
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
func (it *ReferralLockedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralLockedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralLockedOwnership represents a LockedOwnership event raised by the Referral contract.
type ReferralLockedOwnership struct {
	Locked common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedOwnership is a free log retrieval operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Referral *ReferralFilterer) FilterLockedOwnership(opts *bind.FilterOpts) (*ReferralLockedOwnershipIterator, error) {

	logs, sub, err := _Referral.contract.FilterLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return &ReferralLockedOwnershipIterator{contract: _Referral.contract, event: "LockedOwnership", logs: logs, sub: sub}, nil
}

// WatchLockedOwnership is a free log subscription operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_Referral *ReferralFilterer) WatchLockedOwnership(opts *bind.WatchOpts, sink chan<- *ReferralLockedOwnership) (event.Subscription, error) {

	logs, sub, err := _Referral.contract.WatchLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralLockedOwnership)
				if err := _Referral.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
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

// ReferralMintedReferralTokensIterator is returned from FilterMintedReferralTokens and is used to iterate over the raw logs and unpacked data for MintedReferralTokens events raised by the Referral contract.
type ReferralMintedReferralTokensIterator struct {
	Event *ReferralMintedReferralTokens // Event containing the contract specifics and raw log

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
func (it *ReferralMintedReferralTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralMintedReferralTokens)
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
		it.Event = new(ReferralMintedReferralTokens)
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
func (it *ReferralMintedReferralTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralMintedReferralTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralMintedReferralTokens represents a MintedReferralTokens event raised by the Referral contract.
type ReferralMintedReferralTokens struct {
	From      common.Address
	Amount    *big.Int
	NewSupply *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintedReferralTokens is a free log retrieval operation binding the contract event 0x70a3dc5168083fd1722327506f78f85af6d20f327c6b92dc4f5081543b03b931.
//
// Solidity: event MintedReferralTokens(address _from, uint256 _amount, uint256 _newSupply)
func (_Referral *ReferralFilterer) FilterMintedReferralTokens(opts *bind.FilterOpts) (*ReferralMintedReferralTokensIterator, error) {

	logs, sub, err := _Referral.contract.FilterLogs(opts, "MintedReferralTokens")
	if err != nil {
		return nil, err
	}
	return &ReferralMintedReferralTokensIterator{contract: _Referral.contract, event: "MintedReferralTokens", logs: logs, sub: sub}, nil
}

// WatchMintedReferralTokens is a free log subscription operation binding the contract event 0x70a3dc5168083fd1722327506f78f85af6d20f327c6b92dc4f5081543b03b931.
//
// Solidity: event MintedReferralTokens(address _from, uint256 _amount, uint256 _newSupply)
func (_Referral *ReferralFilterer) WatchMintedReferralTokens(opts *bind.WatchOpts, sink chan<- *ReferralMintedReferralTokens) (event.Subscription, error) {

	logs, sub, err := _Referral.contract.WatchLogs(opts, "MintedReferralTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralMintedReferralTokens)
				if err := _Referral.contract.UnpackLog(event, "MintedReferralTokens", log); err != nil {
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

// ReferralSetBonusIterator is returned from FilterSetBonus and is used to iterate over the raw logs and unpacked data for SetBonus events raised by the Referral contract.
type ReferralSetBonusIterator struct {
	Event *ReferralSetBonus // Event containing the contract specifics and raw log

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
func (it *ReferralSetBonusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralSetBonus)
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
		it.Event = new(ReferralSetBonus)
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
func (it *ReferralSetBonusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralSetBonusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralSetBonus represents a SetBonus event raised by the Referral contract.
type ReferralSetBonus struct {
	From     common.Address
	NewBonus *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetBonus is a free log retrieval operation binding the contract event 0xce20f96597c75db9bc70de06a2bbb2ccd70ef20f0677e190486643e9c48354d3.
//
// Solidity: event SetBonus(address _from, uint256 _newBonus)
func (_Referral *ReferralFilterer) FilterSetBonus(opts *bind.FilterOpts) (*ReferralSetBonusIterator, error) {

	logs, sub, err := _Referral.contract.FilterLogs(opts, "SetBonus")
	if err != nil {
		return nil, err
	}
	return &ReferralSetBonusIterator{contract: _Referral.contract, event: "SetBonus", logs: logs, sub: sub}, nil
}

// WatchSetBonus is a free log subscription operation binding the contract event 0xce20f96597c75db9bc70de06a2bbb2ccd70ef20f0677e190486643e9c48354d3.
//
// Solidity: event SetBonus(address _from, uint256 _newBonus)
func (_Referral *ReferralFilterer) WatchSetBonus(opts *bind.WatchOpts, sink chan<- *ReferralSetBonus) (event.Subscription, error) {

	logs, sub, err := _Referral.contract.WatchLogs(opts, "SetBonus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralSetBonus)
				if err := _Referral.contract.UnpackLog(event, "SetBonus", log); err != nil {
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

// ReferralTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Referral contract.
type ReferralTransferIterator struct {
	Event *ReferralTransfer // Event containing the contract specifics and raw log

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
func (it *ReferralTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralTransfer)
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
		it.Event = new(ReferralTransfer)
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
func (it *ReferralTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralTransfer represents a Transfer event raised by the Referral contract.
type ReferralTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Referral *ReferralFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ReferralTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Referral.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ReferralTransferIterator{contract: _Referral.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Referral *ReferralFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ReferralTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Referral.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralTransfer)
				if err := _Referral.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ReferralTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the Referral contract.
type ReferralTransferredOwnershipIterator struct {
	Event *ReferralTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *ReferralTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralTransferredOwnership)
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
		it.Event = new(ReferralTransferredOwnership)
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
func (it *ReferralTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralTransferredOwnership represents a TransferredOwnership event raised by the Referral contract.
type ReferralTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Referral *ReferralFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*ReferralTransferredOwnershipIterator, error) {

	logs, sub, err := _Referral.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &ReferralTransferredOwnershipIterator{contract: _Referral.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_Referral *ReferralFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *ReferralTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _Referral.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralTransferredOwnership)
				if err := _Referral.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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

// ReferralTransferredReferralBonusIterator is returned from FilterTransferredReferralBonus and is used to iterate over the raw logs and unpacked data for TransferredReferralBonus events raised by the Referral contract.
type ReferralTransferredReferralBonusIterator struct {
	Event *ReferralTransferredReferralBonus // Event containing the contract specifics and raw log

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
func (it *ReferralTransferredReferralBonusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralTransferredReferralBonus)
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
		it.Event = new(ReferralTransferredReferralBonus)
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
func (it *ReferralTransferredReferralBonusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralTransferredReferralBonusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralTransferredReferralBonus represents a TransferredReferralBonus event raised by the Referral contract.
type ReferralTransferredReferralBonus struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferredReferralBonus is a free log retrieval operation binding the contract event 0x40cf0b5afee61054126d1b4397c6f217e1f0c4fc65ebe84f382830b9006036fd.
//
// Solidity: event TransferredReferralBonus(address _from, address _to, uint256 indexed _tokenId, uint256 _amount)
func (_Referral *ReferralFilterer) FilterTransferredReferralBonus(opts *bind.FilterOpts, _tokenId []*big.Int) (*ReferralTransferredReferralBonusIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Referral.contract.FilterLogs(opts, "TransferredReferralBonus", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ReferralTransferredReferralBonusIterator{contract: _Referral.contract, event: "TransferredReferralBonus", logs: logs, sub: sub}, nil
}

// WatchTransferredReferralBonus is a free log subscription operation binding the contract event 0x40cf0b5afee61054126d1b4397c6f217e1f0c4fc65ebe84f382830b9006036fd.
//
// Solidity: event TransferredReferralBonus(address _from, address _to, uint256 indexed _tokenId, uint256 _amount)
func (_Referral *ReferralFilterer) WatchTransferredReferralBonus(opts *bind.WatchOpts, sink chan<- *ReferralTransferredReferralBonus, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Referral.contract.WatchLogs(opts, "TransferredReferralBonus", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralTransferredReferralBonus)
				if err := _Referral.contract.UnpackLog(event, "TransferredReferralBonus", log); err != nil {
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
