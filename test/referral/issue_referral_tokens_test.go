package referral_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
    "github.com/ethereum/go-ethereum/common"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Issue", func() {

	When("there are enough referral tokens", func() {
        BeforeEach(func() {
            tx, err := Referral.MintReferralTokens(Owner.TransactOpts(), big.NewInt(10))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        When("the Owner issues 5 referral tokens", func() {
            BeforeEach(func() {
                tx, err := Referral.IssueReferralTokens(Owner.TransactOpts(), RandomAccount.Address(), big.NewInt(5))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            It("should update the referral index", func(){
                ri, err := Referral.ReferralIndex(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(ri.String()).To(Equal(big.NewInt(5).String()))
            })

            It("should update the first 5 tokens' firstOwner", func(){
                fo, err := Referral.FirstOwner(nil,big.NewInt(0))
                Expect(err).ToNot(HaveOccurred())
                Expect(fo).To(Equal(RandomAccount.Address()))

                fo, err = Referral.FirstOwner(nil,big.NewInt(4))
                Expect(err).ToNot(HaveOccurred())
                Expect(fo).To(Equal(RandomAccount.Address()))

                fo, err = Referral.FirstOwner(nil,big.NewInt(5))
                Expect(err).ToNot(HaveOccurred())
                Expect(fo).To(Equal(common.HexToAddress("0x0")))
            })

            It("should emit a IssuedReferralTokens event", func(){
                it, err := Referral.FilterIssuedReferralTokens(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(it.Next()).To(BeTrue())
                evt := it.Event
                Expect(it.Next()).To(BeFalse())
                Expect(evt.From).To(Equal(Owner.Address()))
                Expect(evt.To).To(Equal(RandomAccount.Address()))
                Expect(evt.Amount.String()).To(Equal("5"))
            })

            When("the Owner tries to issue more tokes than the tokens minted so far", func() {

                BeforeEach(func() {
                    tx, err := Referral.IssueReferralTokens(Owner.TransactOpts(), RandomAccount.Address(), big.NewInt(5))
                    Expect(err).ToNot(HaveOccurred())
                    Backend.Commit()
                    Expect(isSuccessful(tx)).To(BeTrue())
                })

                It("should fail", func(){
                    tx, err := Referral.IssueReferralTokens(Owner.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), big.NewInt(0))
                    Expect(err).ToNot(HaveOccurred())
                    Backend.Commit()
                    Expect(isSuccessful(tx)).To(BeFalse())
                    Expect(TestRig.LastExecuted()).To(MatchRegexp(`"tokens exceed the current suppply!"\);`))
                })
            })

        })

        When("the Owner tries to issue 0 referral tokens", func() {
            BeforeEach(func() {
                tx, err := Referral.IssueReferralTokens(Owner.TransactOpts(), RandomAccount.Address(), big.NewInt(0))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            It("should issue 1", func(){
                ri, err := Referral.ReferralIndex(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(ri.String()).To(Equal(big.NewInt(1).String()))
            })

            It("should update the token's firstOwner", func(){
                fo, err := Referral.FirstOwner(nil,big.NewInt(0))
                Expect(err).ToNot(HaveOccurred())
                Expect(fo).To(Equal(RandomAccount.Address()))

                fo, err = Referral.FirstOwner(nil,big.NewInt(1))
                Expect(err).ToNot(HaveOccurred())
                Expect(fo).To(Equal(common.HexToAddress("0x0")))
            })

            It("should emit a IssuedReferralTokens event", func(){
                it, err := Referral.FilterIssuedReferralTokens(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(it.Next()).To(BeTrue())
                evt := it.Event
                Expect(it.Next()).To(BeFalse())
                Expect(evt.From).To(Equal(Owner.Address()))
                Expect(evt.To).To(Equal(RandomAccount.Address()))
                Expect(evt.Amount.String()).To(Equal("1"))
            })
        })

        When("the Owner tries to issue more tokens than the maximum amount giveaway per person", func() {
            It("should fail", func(){
                tx, err := Referral.IssueReferralTokens(Owner.TransactOpts(ethertest.WithGasLimit(80000)), RandomAccount.Address(), big.NewInt(6))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
                Expect(TestRig.LastExecuted()).To(MatchRegexp(`"too many referral tokens given!"\);`))
            })
        })

        When("a random account tries to mint", func() {
            It("should fail", func(){
                tx, err := Referral.MintReferralTokens(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), big.NewInt(5))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
            })
        })
	})

})
