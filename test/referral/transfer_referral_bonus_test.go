package referral_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
    "github.com/ethereum/go-ethereum/common"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Transfer Referral Bonus", func() {

	When("5 tokens are issued", func() {

        ReferralAccount1 := ethertest.NewAccount()
        ReferralAccount2 := ethertest.NewAccount()

        BeforeEach(func() {
            tx, err := Referral.MintReferralTokens(Owner.TransactOpts(), big.NewInt(10))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        BeforeEach(func() {
            tx, err := Referral.IssueReferralTokens(Owner.TransactOpts(), RandomAccount.Address(), big.NewInt(5))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        When("the new token owner transfers 2 tokens to other accounts", func() {

            BeforeEach(func() {
                tx, err := Referral.TransferReferralToken(RandomAccount.TransactOpts(), ReferralAccount1.Address(), big.NewInt(0))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            BeforeEach(func() {
                tx, err := Referral.TransferReferralToken(RandomAccount.TransactOpts(), ReferralAccount2.Address(), big.NewInt(1))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            It("should emit 2 ERC721 Transfer events", func(){
                from := []common.Address{RandomAccount.Address()}
				to := []common.Address{}
                tokenID := []*big.Int{big.NewInt(0), big.NewInt(1)}
                it, err := Referral.FilterTransfer(nil, from, to, tokenID)
                Expect(err).ToNot(HaveOccurred())
                Expect(it.Next()).To(BeTrue())
                evt := it.Event
                Expect(it.Next()).To(BeTrue())
                Expect(evt.From).To(Equal(RandomAccount.Address()))
                Expect(evt.To).To(Equal(ReferralAccount1.Address()))
                Expect(evt.TokenId.String()).To(Equal("0"))
                evt = it.Event
                Expect(it.Next()).To(BeFalse())
                Expect(evt.From).To(Equal(RandomAccount.Address()))
                Expect(evt.To).To(Equal(ReferralAccount2.Address()))
                Expect(evt.TokenId.String()).To(Equal("1"))
            })

            When("the original owner tries to re-transfer the tokens", func() {
                It("should fail", func(){
                    tx, err := Referral.TransferReferralToken(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), big.NewInt(0))
                    Expect(err).ToNot(HaveOccurred())
                    Backend.Commit()
                    Expect(isSuccessful(tx)).To(BeFalse())
                })
            })

            When("the contract owns enough TKN, the cards are activated and the bonus is sent", func() {
                BeforeEach(func() {
    				tx, err := TKN.Credit(BankAccount.TransactOpts(), ReferralAddress, big.NewInt(1000))
    				Expect(err).ToNot(HaveOccurred())
    				Backend.Commit()
    				Expect(isSuccessful(tx)).To(BeTrue())

    			})
                BeforeEach(func() {
                    tx, err := Referral.TransferReferralBonus(Owner.TransactOpts(), []*big.Int{big.NewInt(0), big.NewInt(1)})
                    Expect(err).ToNot(HaveOccurred())
                    Backend.Commit()
                    Expect(isSuccessful(tx)).To(BeTrue())
                })

                It("should emit 2 TransferredReferralBonus events", func(){
                    tokenID := []*big.Int{}
                    it, err := Referral.FilterTransferredReferralBonus(nil, tokenID)
                    Expect(err).ToNot(HaveOccurred())
                    Expect(it.Next()).To(BeTrue())
                    evt := it.Event
                    Expect(it.Next()).To(BeTrue())
                    Expect(evt.From).To(Equal(Owner.Address()))
                    Expect(evt.To).To(Equal(RandomAccount.Address()))
                    Expect(evt.TokenId.String()).To(Equal("0"))
                    Expect(evt.Amount.String()).To(Equal("10"))
                    evt = it.Event
                    Expect(it.Next()).To(BeFalse())
                    Expect(evt.From).To(Equal(Owner.Address()))
                    Expect(evt.To).To(Equal(RandomAccount.Address()))
                    Expect(evt.TokenId.String()).To(Equal("1"))
                    Expect(evt.Amount.String()).To(Equal("10"))
                })

                It("should increase the TKN balance of the original token ower", func() {
					b, err := TKN.BalanceOf(nil, RandomAccount.Address())
					Expect(err).ToNot(HaveOccurred())
					Expect(b.String()).To(Equal("20"))
				})

                It("should mark the tokens as activated", func() {
					a, err := Referral.Activated(nil, big.NewInt(0))
					Expect(err).ToNot(HaveOccurred())
					Expect(a).To(BeTrue())

                    a, err = Referral.Activated(nil, big.NewInt(1))
					Expect(err).ToNot(HaveOccurred())
					Expect(a).To(BeTrue())

                    a, err = Referral.Activated(nil, big.NewInt(2))
					Expect(err).ToNot(HaveOccurred())
					Expect(a).To(BeFalse())
				})

                When("the owner tries to resend the bonuses", func() {
                    BeforeEach(func() {
                        tx, err := Referral.TransferReferralBonus(Owner.TransactOpts(), []*big.Int{big.NewInt(0), big.NewInt(1)})
                        Expect(err).ToNot(HaveOccurred())
                        Backend.Commit()
                        Expect(isSuccessful(tx)).To(BeTrue())
                    })

                    It("should NOT increase the TKN balance of the original token ower", func() {
    					b, err := TKN.BalanceOf(nil, RandomAccount.Address())
    					Expect(err).ToNot(HaveOccurred())
    					Expect(b.String()).To(Equal("20"))
    				})

                    It("should NOT emit any new TransferReferralBonus event", func(){
                        tokenID := []*big.Int{}
                        it, err := Referral.FilterTransferredReferralBonus(nil, tokenID)
                        Expect(err).ToNot(HaveOccurred())
                        Expect(it.Next()).To(BeTrue())
                        Expect(it.Next()).To(BeTrue())
                        Expect(it.Next()).To(BeFalse())
                    })
                })

            })

        })

        When("non-owners try to transfer referral tokens", func() {
            It("should fail", func(){
                tx, err := Referral.TransferReferralToken(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), ReferralAccount1.Address(), big.NewInt(5))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
            })

            It("should fail", func(){
                tx, err := Referral.TransferReferralToken(Owner.TransactOpts(ethertest.WithGasLimit(80000)), ReferralAccount1.Address(), big.NewInt(0))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
            })
        })

    })
})
