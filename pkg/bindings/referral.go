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
const ReferralABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBonus\",\"type\":\"uint256\"}],\"name\":\"setBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"referralIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"transferReferralToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"issueReferralTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referralTokens\",\"type\":\"uint256[]\"}],\"name\":\"transferBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintReferralTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TKNBonus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_totalSuply\",\"type\":\"uint256\"},{\"name\":\"_TKNAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newSupply\",\"type\":\"uint256\"}],\"name\":\"MintedReferralTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ReferralBin is the compiled bytecode used for deploying new contracts.
const ReferralBin = `608060405234801561001057600080fd5b50604051611a46380380611a468339818101604052604081101561003357600080fd5b5080516020909101513360006100717f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0361019f16565b6100a37f80ac58cd000000000000000000000000000000000000000000000000000000006001600160e01b0361019f16565b600580546001600160a01b0319166001600160a01b0384161760ff60a01b1916740100000000000000000000000000000000000000008315158102919091179182905560ff91041661012c57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a15050600691909155600a80546001600160a01b0319166001600160a01b0390921691909117815560085561026d565b7fffffffff00000000000000000000000000000000000000000000000000000000808216141561023057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b6117ca8061027c6000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c80637392b3b3116100c3578063b88d4fde1161007c578063b88d4fde1461038a578063bb6acb9314610450578063c12712e81461047c578063c260d95d146104ec578063c35305b114610509578063e985e9c5146105115761014d565b80637392b3b3146102ea5780637f5d6139146102f25780638d75fe051461031e5780638da5cb5b14610326578063a22cb4651461032e578063b242e5341461035c5761014d565b80632121dc75116101155780632121dc751461022b57806323b872dd1461023357806342842e0e146102695780636352211e1461029f57806370a08231146102bc578063715018a6146102e25761014d565b806301ffc9a714610152578063081812fc1461018d578063095ea7b3146101c65780630b98f975146101f457806318160ddd14610211575b600080fd5b6101796004803603602081101561016857600080fd5b50356001600160e01b03191661053f565b604080519115158252519081900360200190f35b6101aa600480360360208110156101a357600080fd5b503561055e565b604080516001600160a01b039092168252519081900360200190f35b6101f2600480360360408110156101dc57600080fd5b506001600160a01b0381351690602001356105c0565b005b6101f26004803603602081101561020a57600080fd5b50356106d1565b61021961071e565b60408051918252519081900360200190f35b610179610724565b6101f26004803603606081101561024957600080fd5b506001600160a01b03813581169160208101359091169060400135610734565b6101f26004803603606081101561027f57600080fd5b506001600160a01b03813581169160208101359091169060400135610789565b6101aa600480360360208110156102b557600080fd5b50356107a4565b610219600480360360208110156102d257600080fd5b50356001600160a01b03166107fe565b6101f2610866565b610219610959565b6101f26004803603604081101561030857600080fd5b506001600160a01b03813516906020013561095f565b61021961096e565b6101aa610974565b6101f26004803603604081101561034457600080fd5b506001600160a01b0381351690602001351515610983565b6101f26004803603604081101561037257600080fd5b506001600160a01b0381351690602001351515610a4f565b6101f2600480360360808110156103a057600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156103db57600080fd5b8201836020820111156103ed57600080fd5b8035906020019184600183028401116401000000008311171561040f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610bfe945050505050565b6101f26004803603604081101561046657600080fd5b506001600160a01b038135169060200135610c56565b6101f26004803603602081101561049257600080fd5b8101906020810181356401000000008111156104ad57600080fd5b8201836020820111156104bf57600080fd5b803590602001918460208302840111640100000000831117156104e157600080fd5b509092509050610da5565b6101f26004803603602081101561050257600080fd5b5035610eda565b610219611029565b6101796004803603604081101561052757600080fd5b506001600160a01b038135811691602001351661102f565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006105698261105d565b6105a45760405162461bcd60e51b815260040180806020018281038252602c8152602001806116cc602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006105cb826107a4565b9050806001600160a01b0316836001600160a01b0316141561061e5760405162461bcd60e51b81526004018080602001828103825260218152602001806117216021913960400191505060405180910390fd5b336001600160a01b038216148061063a575061063a813361102f565b6106755760405162461bcd60e51b81526004018080602001828103825260388152602001806116416038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6106da3361107a565b610719576040805162461bcd60e51b815260206004820152601660248201526000805160206115f5833981519152604482015290519081900360640190fd5b600855565b60065481565b600554600160a01b900460ff1690565b61073e338261108e565b6107795760405162461bcd60e51b81526004018080602001828103825260318152602001806117656031913960400191505060405180910390fd5b610784838383611132565b505050565b61078483838360405180602001604052806000815250610bfe565b6000818152600160205260408120546001600160a01b0316806107f85760405162461bcd60e51b81526004018080602001828103825260298152602001806116a36029913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166108455760405162461bcd60e51b815260040180806020018281038252602a815260200180611679602a913960400191505060405180910390fd5b6001600160a01b03821660009081526003602052604090206107f890611276565b61086f3361107a565b6108ae576040805162461bcd60e51b815260206004820152601660248201526000805160206115f5833981519152604482015290519081900360640190fd5b600554600160a01b900460ff1661090c576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600580546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b60075481565b61096a338383611132565b5050565b60095481565b6005546001600160a01b031690565b6001600160a01b0382163314156109e1576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b3360008181526004602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b610a583361107a565b610a97576040805162461bcd60e51b815260206004820152601660248201526000805160206115f5833981519152604482015290519081900360640190fd5b600554600160a01b900460ff16610af5576040805162461bcd60e51b815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b038216610b3a5760405162461bcd60e51b81526004018080602001828103825260238152602001806117426023913960400191505060405180910390fd5b6005805460ff60a01b1916600160a01b8315150217905580610b9357604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600554604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600580546001600160a01b0319166001600160a01b0392909216919091179055565b610c09848484610734565b610c158484848461127a565b610c505760405162461bcd60e51b815260040180806020018281038252603281526020018061159f6032913960400191505060405180910390fd5b50505050565b610c5f3361107a565b610c9e576040805162461bcd60e51b815260206004820152601660248201526000805160206115f5833981519152604482015290519081900360640190fd5b600654816007540110610cf8576040805162461bcd60e51b815260206004820181905260248201527f746f6b656e7320657863656564656420746865206d61632073757070706c7921604482015290519081900360640190fd5b6005811115610d4e576040805162461bcd60e51b815260206004820152601f60248201527f746f6f206d616e7920726566657272616c20746f6b656e7320676976656e2100604482015290519081900360640190fd5b6007545b8160075401811015610d9857610d69338483611132565b6000818152600c6020526040902080546001600160a01b0319166001600160a01b038516179055600101610d52565b5060078054909101905550565b610dae3361107a565b610ded576040805162461bcd60e51b815260206004820152601660248201526000805160206115f5833981519152604482015290519081900360640190fd5b60005b81811015610784576000838383818110610e0657fe5b602090810292909201356000818152600b9093526040909220549192505060ff16610ed1576000818152600b60209081526040808320805460ff19166001179055600a54600c835281842054600854835163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152925191169363a9059cbb93604480850194919392918390030190829087803b158015610ea457600080fd5b505af1158015610eb8573d6000803e3d6000fd5b505050506040513d6020811015610ece57600080fd5b50505b50600101610df0565b610ee33361107a565b610f22576040805162461bcd60e51b815260206004820152601660248201526000805160206115f5833981519152604482015290519081900360640190fd5b600954808201908111610f72576040805162461bcd60e51b81526020600482015260136024820152721bdd995c999b1bddc81bdc880c081a5b9c1d5d606a1b604482015290519081900360640190fd5b600654811115610fc1576040805162461bcd60e51b81526020600482015260156024820152741d1bdd185b081cdd5c1c1b1e48195e18d959591959605a1b604482015290519081900360640190fd5b6009545b81811015610fdf57610fd733826113ad565b600101610fc5565b506009819055604080513381526020810184905280820183905290517f70a3dc5168083fd1722327506f78f85af6d20f327c6b92dc4f5081543b03b9319181900360600190a15050565b60085481565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b6005546001600160a01b0390811691161490565b60006110998261105d565b6110d45760405162461bcd60e51b815260040180806020018281038252602c815260200180611615602c913960400191505060405180910390fd5b60006110df836107a4565b9050806001600160a01b0316846001600160a01b0316148061111a5750836001600160a01b031661110f8461055e565b6001600160a01b0316145b8061112a575061112a818561102f565b949350505050565b826001600160a01b0316611145826107a4565b6001600160a01b03161461118a5760405162461bcd60e51b81526004018080602001828103825260298152602001806116f86029913960400191505060405180910390fd5b6001600160a01b0382166111cf5760405162461bcd60e51b81526004018080602001828103825260248152602001806115d16024913960400191505060405180910390fd5b6111d8816114de565b6001600160a01b03831660009081526003602052604090206111f99061151b565b6001600160a01b038216600090815260036020526040902061121a90611532565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b5490565b600061128e846001600160a01b031661153b565b61129a5750600161112a565b604051630a85bd0160e11b815233600482018181526001600160a01b03888116602485015260448401879052608060648501908152865160848601528651600095928a169463150b7a029490938c938b938b939260a4019060208501908083838e5b838110156113145781810151838201526020016112fc565b50505050905090810190601f1680156113415780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561136357600080fd5b505af1158015611377573d6000803e3d6000fd5b505050506040513d602081101561138d57600080fd5b50516001600160e01b031916630a85bd0160e11b14915050949350505050565b6001600160a01b038216611408576040805162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015290519081900360640190fd5b6114118161105d565b15611463576040805162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015290519081900360640190fd5b600081815260016020908152604080832080546001600160a01b0319166001600160a01b0387169081179091558352600390915290206114a290611532565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000818152600260205260409020546001600160a01b03161561151857600081815260026020526040902080546001600160a01b03191690555b50565b805461152e90600163ffffffff61154116565b9055565b80546001019055565b3b151590565b600082821115611598576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f206164647265737373656e646572206973206e6f7420616e206f776e6572000000000000000000004552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65726f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573734552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a265627a7a72305820d6f8026c3190a30d9a2a44c87abd856f1528d8866e67cf4b5bf69a3c1f1ed03164736f6c634300050a0032`

// DeployReferral deploys a new Ethereum contract, binding an instance of Referral to it.
func DeployReferral(auth *bind.TransactOpts, backend bind.ContractBackend, _totalSuply *big.Int, _TKNAddress common.Address) (common.Address, *types.Transaction, *Referral, error) {
	parsed, err := abi.JSON(strings.NewReader(ReferralABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReferralBin), backend, _totalSuply, _TKNAddress)
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

// TKNBonus is a free data retrieval call binding the contract method 0xc35305b1.
//
// Solidity: function TKNBonus() constant returns(uint256)
func (_Referral *ReferralCaller) TKNBonus(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Referral.contract.Call(opts, out, "TKNBonus")
	return *ret0, err
}

// TKNBonus is a free data retrieval call binding the contract method 0xc35305b1.
//
// Solidity: function TKNBonus() constant returns(uint256)
func (_Referral *ReferralSession) TKNBonus() (*big.Int, error) {
	return _Referral.Contract.TKNBonus(&_Referral.CallOpts)
}

// TKNBonus is a free data retrieval call binding the contract method 0xc35305b1.
//
// Solidity: function TKNBonus() constant returns(uint256)
func (_Referral *ReferralCallerSession) TKNBonus() (*big.Int, error) {
	return _Referral.Contract.TKNBonus(&_Referral.CallOpts)
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
// Solidity: function issueReferralTokens(address _to, uint256 _count) returns()
func (_Referral *ReferralTransactor) IssueReferralTokens(opts *bind.TransactOpts, _to common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "issueReferralTokens", _to, _count)
}

// IssueReferralTokens is a paid mutator transaction binding the contract method 0xbb6acb93.
//
// Solidity: function issueReferralTokens(address _to, uint256 _count) returns()
func (_Referral *ReferralSession) IssueReferralTokens(_to common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.IssueReferralTokens(&_Referral.TransactOpts, _to, _count)
}

// IssueReferralTokens is a paid mutator transaction binding the contract method 0xbb6acb93.
//
// Solidity: function issueReferralTokens(address _to, uint256 _count) returns()
func (_Referral *ReferralTransactorSession) IssueReferralTokens(_to common.Address, _count *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.IssueReferralTokens(&_Referral.TransactOpts, _to, _count)
}

// MintReferralTokens is a paid mutator transaction binding the contract method 0xc260d95d.
//
// Solidity: function mintReferralTokens(uint256 amount) returns()
func (_Referral *ReferralTransactor) MintReferralTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "mintReferralTokens", amount)
}

// MintReferralTokens is a paid mutator transaction binding the contract method 0xc260d95d.
//
// Solidity: function mintReferralTokens(uint256 amount) returns()
func (_Referral *ReferralSession) MintReferralTokens(amount *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.MintReferralTokens(&_Referral.TransactOpts, amount)
}

// MintReferralTokens is a paid mutator transaction binding the contract method 0xc260d95d.
//
// Solidity: function mintReferralTokens(uint256 amount) returns()
func (_Referral *ReferralTransactorSession) MintReferralTokens(amount *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.MintReferralTokens(&_Referral.TransactOpts, amount)
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

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(uint256 newBonus) returns()
func (_Referral *ReferralTransactor) SetBonus(opts *bind.TransactOpts, newBonus *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "setBonus", newBonus)
}

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(uint256 newBonus) returns()
func (_Referral *ReferralSession) SetBonus(newBonus *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.SetBonus(&_Referral.TransactOpts, newBonus)
}

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(uint256 newBonus) returns()
func (_Referral *ReferralTransactorSession) SetBonus(newBonus *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.SetBonus(&_Referral.TransactOpts, newBonus)
}

// TransferBonus is a paid mutator transaction binding the contract method 0xc12712e8.
//
// Solidity: function transferBonus(uint256[] _referralTokens) returns()
func (_Referral *ReferralTransactor) TransferBonus(opts *bind.TransactOpts, _referralTokens []*big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "transferBonus", _referralTokens)
}

// TransferBonus is a paid mutator transaction binding the contract method 0xc12712e8.
//
// Solidity: function transferBonus(uint256[] _referralTokens) returns()
func (_Referral *ReferralSession) TransferBonus(_referralTokens []*big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferBonus(&_Referral.TransactOpts, _referralTokens)
}

// TransferBonus is a paid mutator transaction binding the contract method 0xc12712e8.
//
// Solidity: function transferBonus(uint256[] _referralTokens) returns()
func (_Referral *ReferralTransactorSession) TransferBonus(_referralTokens []*big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferBonus(&_Referral.TransactOpts, _referralTokens)
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

// TransferReferralToken is a paid mutator transaction binding the contract method 0x7f5d6139.
//
// Solidity: function transferReferralToken(address _to, uint256 _tokenid) returns()
func (_Referral *ReferralTransactor) TransferReferralToken(opts *bind.TransactOpts, _to common.Address, _tokenid *big.Int) (*types.Transaction, error) {
	return _Referral.contract.Transact(opts, "transferReferralToken", _to, _tokenid)
}

// TransferReferralToken is a paid mutator transaction binding the contract method 0x7f5d6139.
//
// Solidity: function transferReferralToken(address _to, uint256 _tokenid) returns()
func (_Referral *ReferralSession) TransferReferralToken(_to common.Address, _tokenid *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferReferralToken(&_Referral.TransactOpts, _to, _tokenid)
}

// TransferReferralToken is a paid mutator transaction binding the contract method 0x7f5d6139.
//
// Solidity: function transferReferralToken(address _to, uint256 _tokenid) returns()
func (_Referral *ReferralTransactorSession) TransferReferralToken(_to common.Address, _tokenid *big.Int) (*types.Transaction, error) {
	return _Referral.Contract.TransferReferralToken(&_Referral.TransactOpts, _to, _tokenid)
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
// Solidity: event MintedReferralTokens(address from, uint256 amount, uint256 newSupply)
func (_Referral *ReferralFilterer) FilterMintedReferralTokens(opts *bind.FilterOpts) (*ReferralMintedReferralTokensIterator, error) {

	logs, sub, err := _Referral.contract.FilterLogs(opts, "MintedReferralTokens")
	if err != nil {
		return nil, err
	}
	return &ReferralMintedReferralTokensIterator{contract: _Referral.contract, event: "MintedReferralTokens", logs: logs, sub: sub}, nil
}

// WatchMintedReferralTokens is a free log subscription operation binding the contract event 0x70a3dc5168083fd1722327506f78f85af6d20f327c6b92dc4f5081543b03b931.
//
// Solidity: event MintedReferralTokens(address from, uint256 amount, uint256 newSupply)
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
