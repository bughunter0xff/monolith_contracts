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

    When("ETH is swapped with TKN ", func() {

        //When the wallet has 1 ETH
        BeforeEach(func() {
            BankAccount.MustTransfer(Backend, WalletAddress, EthToWei(1))
        })

        BeforeEach(func() {
            minConversionRate := new(big.Int)
            minConversionRate.SetString("384252315663000000000",10)
            maxDestAmount := new(big.Int)
            maxDestAmount.SetString("8000000000000000000000000000000000000000000000000000000000000000",16)
            hint := common.Hex2Bytes("5045524d")

            a, err := abi.JSON(strings.NewReader(TRADEWITHHINT_ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("tradeWithHint",
                common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"),
                EthToWei(1),
                TKNBurnerAddress,
                WalletAddress,
                maxDestAmount,
                minConversionRate,
                KNCWallet.Address(),
                hint,
            )
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteTransaction(Owner.TransactOpts(), KyberNetworkProxyAddress, EthToWei(1), data)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("should increase the TKN balance of the Wallet", func() {
            b, err := TKNBurner.BalanceOf(nil, WalletAddress)
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("39489134037"))
        })

        // It("should decrease the KNC balance of the KNC wallet", func() {
        //     b, err := KNCBurner.BalanceOf(nil, KNCWallet.Address())
        //     Expect(err).ToNot(HaveOccurred())
        //     Expect(b.String()).ToNot(Equal(EthToWei(38000).String()))
        // })

        It("should decrease the ETH balance of the Wallet by 1 ETH", func() {
            b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
            Expect(e).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal(EthToWei(0).String()))
        })
    })

    When("TKN is swapped with ETH ", func() {

        //Wallet has 38 TKN and approves kyber proxy
        BeforeEach(func() {
            tx, err := TKNBurner.Mint(Owner.TransactOpts(), WalletAddress, big.NewInt(3800000000))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        BeforeEach(func() {
            a, err := abi.JSON(strings.NewReader(ERC20_APPROVE_ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("approve", KyberNetworkProxyAddress, big.NewInt(3800000000))
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        BeforeEach(func() {
            minConversionRate := new(big.Int)
            minConversionRate.SetString("2436462325766737",10)
            maxDestAmount := new(big.Int)
            maxDestAmount.SetString("8000000000000000000000000000000000000000000000000000000000000000",16)
            hint := common.Hex2Bytes("5045524d")

            a, err := abi.JSON(strings.NewReader(TRADEWITHHINT_ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("tradeWithHint",
                TKNBurnerAddress,
                big.NewInt(3800000000),
                common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"),
                WalletAddress,
                maxDestAmount,
                minConversionRate,
                KNCWallet.Address(),
                hint,
            )
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteTransaction(Owner.TransactOpts(), KyberNetworkProxyAddress, big.NewInt(0), data)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("should decrease the TKN balance of the Wallet", func() {
            b, err := TKNBurner.BalanceOf(nil, WalletAddress)
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("0"))
        })

        It("should increase the ETH balance of the Wallet", func() {
            b, e := Backend.BalanceAt(context.Background(), WalletAddress, nil)
            Expect(e).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("95382270831578348"))
        })
    })

    When("TKN is swapped with DAI ", func() {

        //Wallet has 38 TKN and approves kyber proxy
        BeforeEach(func() {
            tx, err := TKNBurner.Mint(Owner.TransactOpts(), WalletAddress, big.NewInt(3800000000))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        BeforeEach(func() {
            a, err := abi.JSON(strings.NewReader(ERC20_APPROVE_ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("approve", KyberNetworkProxyAddress, big.NewInt(3800000000))
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteTransaction(Owner.TransactOpts(), TKNBurnerAddress, big.NewInt(0), data)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        BeforeEach(func() {
            minConversionRate := new(big.Int)
            minConversionRate.SetString("2436462325766737",10)
            maxDestAmount := new(big.Int)
            maxDestAmount.SetString("8000000000000000000000000000000000000000000000000000000000000000",16)
            hint := common.Hex2Bytes("5045524d")

            a, err := abi.JSON(strings.NewReader(TRADEWITHHINT_ABI))
            Expect(err).ToNot(HaveOccurred())
            data, err := a.Pack("tradeWithHint",
                TKNBurnerAddress,
                big.NewInt(3800000000),
                DAIAddress,
                WalletAddress,
                maxDestAmount,
                minConversionRate,
                KNCWallet.Address(),
                hint,
            )
            Expect(err).ToNot(HaveOccurred())

            tx, err := Wallet.ExecuteTransaction(Owner.TransactOpts(), KyberNetworkProxyAddress, big.NewInt(0), data)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("should decrease the TKN balance of the Wallet", func() {
            b, err := TKNBurner.BalanceOf(nil, WalletAddress)
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("0"))
        })

        It("should increase the DAI balance of the Wallet", func() {
            b, err := DAI.BalanceOf(nil, WalletAddress)
            Expect(err).ToNot(HaveOccurred())
            Expect(b.String()).To(Equal("3778500074"))
        })
    })


})
const ERC20_APPROVE_ABI = `[
    {
        "constant": false,
        "inputs": [
            {
                "name": "_spender",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "approve",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`

const TRADEWITHHINT_ABI = `[
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
