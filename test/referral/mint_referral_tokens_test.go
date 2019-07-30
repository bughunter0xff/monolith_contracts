package referral_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Mint", func() {

	When("the owner mints", func() {
        BeforeEach(func() {
            tx, err := Referral.MintReferralTokens(Owner.TransactOpts(), big.NewInt(50))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("the total supply has to be fixed", func(){
            ts, err := Referral.TotalSupply(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(ts.String()).To(Equal(big.NewInt(60).String()))
        })

        It("should update the number of total tokens minted", func(){
            mt, err := Referral.MintedTokens(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(mt.String()).To(Equal(big.NewInt(50).String()))
        })

        It("should emit a MintedReferralTokens event", func(){
            it, err := Referral.FilterMintedReferralTokens(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(it.Next()).To(BeTrue())
            evt := it.Event
            Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(Owner.Address()))
            Expect(evt.Amount.String()).To(Equal("50"))
            Expect(evt.NewSupply.String()).To(Equal("50"))
        })

        When("the mints the rest of the tokens", func() {
            BeforeEach(func() {
                ts, _ := Referral.TotalSupply(nil)
                mt, _ := Referral.MintedTokens(nil)
                supply := big.NewInt(0)
                supply.Add(supply,ts.Sub(ts,mt)) //e.g. supplyPlusOne = total supply(60) - minted(50)  = 10
                tx, err := Referral.MintReferralTokens(Owner.TransactOpts(), supply)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            It("should update the number of total tokens minted", func(){
                mt, err := Referral.MintedTokens(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(mt.String()).To(Equal(big.NewInt(60).String()))
            })
        })

        When("the owner tries to mint more than the predetermined supply", func() {
            It("should fail", func(){
                ts, _ := Referral.TotalSupply(nil)
                mt, _ := Referral.MintedTokens(nil)
                supplyPlusOne := big.NewInt(1)
                supplyPlusOne.Add(supplyPlusOne,ts.Sub(ts,mt)) //e.g. supplyPlusOne = total supply(60) - minted(50) + 1 = 11
                tx, err := Referral.MintReferralTokens(Owner.TransactOpts(ethertest.WithGasLimit(800000)), supplyPlusOne)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
                Expect(TestRig.LastExecuted()).To(MatchRegexp(`"total supply exceeded"\);`))
            })
        })

        When("the owner tries to mint a huge amount of tokens leading to overflow", func() {
            It("should fail", func(){
                tx, err := Referral.MintReferralTokens(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), big.NewInt(-1))
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeFalse())
                Expect(TestRig.LastExecuted()).To(MatchRegexp(`"overflow or 0 input"\);`))
            })
        })
	})

    When("the owner tries to mint 0 tokens", func() {
        It("should fail", func(){
            tx, err := Referral.MintReferralTokens(Owner.TransactOpts(ethertest.WithGasLimit(1000000)), big.NewInt(0))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
            Expect(TestRig.LastExecuted()).To(MatchRegexp(`"overflow or 0 input"\);`))
        })
    })

    When("a random account tries to mint", func() {
        It("should fail", func(){
            tx, err := Referral.MintReferralTokens(RandomAccount.TransactOpts(ethertest.WithGasLimit(1000000)), big.NewInt(50))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
        })
    })

})
