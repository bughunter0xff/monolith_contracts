package referral_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Mint", func() {

	When("the Owner mints", func() {
        BeforeEach(func() {
            tx, err := Referral.MintReferralTokens(Owner.TransactOpts())
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())
        })

        It("should update the number of total tokens minted", func(){
            tm, err := Referral.MintedTokens(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(tm.String()).To(Equal(big.NewInt(10).String()))
        })

        It("should amit a MintedReferralTokens event", func(){
            it, err := Referral.FilterMintedReferralTokens(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(it.Next()).To(BeTrue())
            evt := it.Event
            Expect(it.Next()).To(BeFalse())
            Expect(evt.Amount.String()).To(Equal("10"))
            Expect(evt.NewSupply.String()).To(Equal("10"))
        })

	})

    When("a random account tries to mint", func() {
        It("should fail", func(){
            tx, err := Referral.MintReferralTokens(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
        })
    })

})
