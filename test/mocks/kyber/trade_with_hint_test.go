package kyber_test

import (
    "math/big"
    "strings"
    "context"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
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

        a, err := abi.JSON(strings.NewReader(TRADEWITHHINTABI))
        Expect(err).ToNot(HaveOccurred())
        data, err := a.Pack("tradeWithHint",
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

        tx, err := Wallet.ExecuteTransaction(Owner.TransactOpts(), KyberNetworkProxyAddress, EthToWei(1), data)
        Expect(err).ToNot(HaveOccurred())
        Backend.Commit()
        Expect(isSuccessful(tx)).To(BeTrue())
    })

    FIt("should increase TKN balance of the Wallet", func() {
        b, err := TKNBurner.BalanceOf(nil, WalletAddress)
        Expect(err).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal("39489134037"))
    })

    FIt("should decrease the ETH balance of the Wallet by 1 ETH", func() {
        b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
        Expect(e).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal(EthToWei(0).String()))
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
