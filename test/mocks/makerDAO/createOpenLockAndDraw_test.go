package makerDAO_test

import (
    "encoding/hex"
    "math/big"
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

            tx, err = Dai.SetAuthority(BankAccount.TransactOpts(), DSGuardAddress)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

            //keccak256(mint(address,uint256)) = 0x40c10f19
            mintID := common.Hex2Bytes("40c10f19")
            var mintID32 = [32]byte{}
            copy(mintID32[:4], mintID)

            tx, err = DSGuard.Permit(BankAccount.TransactOpts(), SaiTubAddress, SkrAddress, mintID32)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

            tx, err = DSGuard.Permit(BankAccount.TransactOpts(), SaiTubAddress, DaiAddress, mintID32)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

            //set axe
            param := common.Hex2Bytes(hex.EncodeToString([]byte("axe")))
            var param32 = [32]byte{}
            copy(param32[:3], param)

            newVal := new(big.Int)
            newVal, ok := newVal.SetString("1130000000000000000000000000",10)
            Expect(ok).To(BeTrue())

            tx, err = SaiTub.Mold(BankAccount.TransactOpts(), param32, newVal)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

            //set cap
            param = common.Hex2Bytes(hex.EncodeToString([]byte("cap")))
            copy(param32[:3], param)

            tx, err = SaiTub.Mold(BankAccount.TransactOpts(), param32, EthToWei(10))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

            //set mat
            param = common.Hex2Bytes(hex.EncodeToString([]byte("mat")))
            copy(param32[:3], param)

            newVal, ok = newVal.SetString("1500000000000000000000000000",10)
            Expect(ok).To(BeTrue())

            tx, err = SaiTub.Mold(BankAccount.TransactOpts(), param32, newVal)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

            //set stability fee
            param = common.Hex2Bytes(hex.EncodeToString([]byte("fee")))
            copy(param32[:3], param)

            newVal, ok = newVal.SetString("1000000005913228294456064283",10)
            Expect(ok).To(BeTrue())

            tx, err = SaiTub.Mold(BankAccount.TransactOpts(), param32, newVal)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

        })

    	When("CreateOpenLockAndDraw is called by with 0.03 ETH (30 Finney)", func() {

            It("Should succeed", func() {
    			cap, err := SaiTub.Cap(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(cap.String()).To(Equal(EthToWei(10).String()))
    		})

            It("Should succeed", func() {
    			axe, err := SaiTub.Axe(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(axe.String()).To(Equal("1130000000000000000000000000"))
    		})

            It("Should succeed", func() {
    			axe, err := SaiTub.Tax(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(axe.String()).To(Equal("1000000000000000000000000000"))
    		})

            It("Should succeed", func() {
    			mat, err := SaiTub.Mat(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(mat.String()).To(Equal("1500000000000000000000000000"))
    		})

            It("Should succeed", func() {
    			fee, err := SaiTub.Fee(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(fee.String()).To(Equal("1000000005913228294456064283"))
    		})

    		FIt("Should succeed", func() {
    			tx, err := SaiProxy.CreateOpenLockAndDraw(RandomAccount.TransactOpts(ethertest.WithValue(FinneyToWei(30))), ProxyRegistryAddress, SaiTubAddress, EthToWei(3))
                Expect(err).ToNot(HaveOccurred())
    			Backend.Commit()
    			Expect(isSuccessful(tx)).To(BeTrue())
    		})

    	})
    })
})
