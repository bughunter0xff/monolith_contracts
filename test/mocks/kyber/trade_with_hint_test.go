package kyber_test

import (
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/tokencard/ethertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
)

var _ = Describe("kyber trade with hint", func() {

    BeforeEach(func() {
        minConversionRate := new(big.Int)
        minConversionRate.SetString("384252315663000000000",10)
        maxDestAmount := new(big.Int)
        maxDestAmount.SetString("8000000000000000000000000000000000000000000000000000000000000000",16)
        hint := common.Hex2Bytes("5045524d")
        tx, err := KyberNetworkProxy.TradeWithHint(BankAccount.TransactOpts(ethertest.WithValue(EthToWei(1))),
        common.HexToAddress("0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"),
        EthToWei(1),
        TKNBurnerAddress,
        WalletAddress,
        maxDestAmount,
        minConversionRate,
        Owner.Address(),
        hint,
        )
        Expect(err).ToNot(HaveOccurred())
        Backend.Commit()
        Expect(isSuccessful(tx)).To(BeTrue())
    })

    It("should increase TKN balance of the Wallet address", func() {
        b, err := TKNBurner.BalanceOf(nil, WalletAddress)
        Expect(err).ToNot(HaveOccurred())
        Expect(b.String()).ToNot(Equal("0"))
    })




})

const TRADEWITHHINTABI = `[
    {
        "constant": false,
        "inputs": [
            {
                "name": "src",
                "type": "address"
            },
            {
                "name": "srcAmount",
                "type": "uint256"
            },
            {
                "name": "dest",
                "type": "address"
            },
            {
                "name": "destAddress",
                "type": "address"
            },
            {
                "name": "maxDestAmount",
                "type": "uint256"
            },
            {
                "name": "minConversionRate",
                "type": "uint256"
            },
            {
                "name": "walletId",
                "type": "address"
            },
            {
                "name": "hint",
                "type": "bytes"
            }
        ],
        "name": "tradeWithHint",
        "outputs": [
            {
                "name": "",
                "type": "uint256"
            }
        ],
        "payable": true,
        "stateMutability": "payable",
        "type": "function"
    }
]`
