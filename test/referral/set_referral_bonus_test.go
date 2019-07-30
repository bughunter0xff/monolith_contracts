package referral_test

import (
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/test/shared"
	"github.com/tokencard/ethertest"
)

var _ = Describe("Set Referral Bonus", func() {


    It("the initial bonus should be 10", func(){
        b, err := Referral.TKNReferralBonus(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal("10"))
    })

	When("the owner sets the bonus to 15 TKN", func() {

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

        BeforeEach(func() {
			tx, err := TKN.Credit(BankAccount.TransactOpts(), ReferralAddress, big.NewInt(1000))
			Expect(err).ToNot(HaveOccurred())
			Backend.Commit()
			Expect(isSuccessful(tx)).To(BeTrue())

		})

        BeforeEach(func() {
            tx, err := Referral.SetReferralBonus(Owner.TransactOpts(), big.NewInt(15))
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

        It("should emit a SetBonus event", func(){
            it, err := Referral.FilterSetBonus(nil)
            Expect(err).ToNot(HaveOccurred())
            Expect(it.Next()).To(BeTrue())
            evt := it.Event
            Expect(it.Next()).To(BeFalse())
            Expect(evt.From).To(Equal(Owner.Address()))
            Expect(evt.NewBonus.String()).To(Equal("15"))
        })

        It("should increase the TKN balance of the original token according to the new bonus", func() {
			b, err := TKN.BalanceOf(nil, RandomAccount.Address())
			Expect(err).ToNot(HaveOccurred())
			Expect(b.String()).To(Equal("30"))
		})
    })

    When("a random account tries to set a new bonus", func() {
        It("should fail", func(){
            tx, err := Referral.SetReferralBonus(RandomAccount.TransactOpts(ethertest.WithGasLimit(80000)), big.NewInt(15))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeFalse())
        })
    })

})
