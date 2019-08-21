package makerDAO_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
	"github.com/tokencard/ethertest"
    "github.com/ethereum/go-ethereum/common"
)

var _ = Describe("CreateOpenLockAndDraw", func() {

    When("the authority and the persmissions have been set", func() {

        BeforeEach(func() {
            tx, err := Skr.SetAuthority(BankAccount.TransactOpts(), DSGuardAddress)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

            mintID := common.Hex2Bytes("40c10f1900000000000000000000000000000000000000000000000000000000")
            var mintID32 = [32]byte{}
            copy(mintID32[:], mintID)
            tx, err = DSGuard.Permit(BankAccount.TransactOpts(), SaiTubAddress, SkrAddress, mintID32)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

        })

    	When("CreateOpenLockAndDraw is called by with 0.03 ETH (30 Finney)", func() {

            It("Should succeed", func() {
                tx, err := SaiProxy.CreateAndOpen(RandomAccount.TransactOpts(ethertest.WithGasLimit(2000000)), ProxyRegistryAddress, SaiTubAddress)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

            It("Should succeed", func() {
                tx, err := SaiProxy.CreateOpenAndLock(RandomAccount.TransactOpts(ethertest.WithGasLimit(2000000), ethertest.WithValue(FinneyToWei(30))), ProxyRegistryAddress, SaiTubAddress)
                Expect(err).ToNot(HaveOccurred())
                Backend.Commit()
                Expect(isSuccessful(tx)).To(BeTrue())
            })

    		// It("Should succeed", func() {
    		// 	tx, err := SaiProxy.CreateOpenLockAndDraw(RandomAccount.TransactOpts(ethertest.WithValue(FinneyToWei(30))), ProxyRegistryAddress, SaiTubAddress, EthToWei(3))
            //     Expect(err).ToNot(HaveOccurred())
    		// 	Backend.Commit()
    		// 	Expect(isSuccessful(tx)).To(BeTrue())
    		// })
    	})
    })
})
