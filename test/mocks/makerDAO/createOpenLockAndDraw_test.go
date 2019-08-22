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

            tx, err = Pip.SetAuthority(BankAccount.TransactOpts(), DSGuardAddress)
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

            tx, err = Pip.SetMin(BankAccount.TransactOpts(), big.NewInt(7))
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

            tx, err = SaiProxy.CreateOpenLockAndDraw(RandomAccount.TransactOpts(ethertest.WithValue(FinneyToWei(30))), ProxyRegistryAddress, SaiTubAddress, EthToWei(3))
            Expect(err).ToNot(HaveOccurred())
            Backend.Commit()
            Expect(isSuccessful(tx)).To(BeTrue())

        })

    	When("CreateOpenLockAndDraw is called by with 0.03 ETH (30 Finney)", func() {

            It("update the cap value", func() {
    			cap, err := SaiTub.Cap(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(cap.String()).To(Equal(EthToWei(10).String()))
    		})

            It("update the axe value", func() {
    			axe, err := SaiTub.Axe(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(axe.String()).To(Equal("1130000000000000000000000000"))
    		})

            It("update the tax value", func() {
    			tax, err := SaiTub.Tax(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(tax.String()).To(Equal("1000000000000000000000000000"))
    		})

            It("update the mat value", func() {
    			mat, err := SaiTub.Mat(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(mat.String()).To(Equal("1500000000000000000000000000"))
    		})

            It("update the stability fee", func() {
    			fee, err := SaiTub.Fee(nil)
                Expect(err).ToNot(HaveOccurred())
    			Expect(fee.String()).To(Equal("1000000005913228294456064283"))
    		})

            It("read the mock value", func() {
    			val, err := Pip.Read(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(val).To(Equal([32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 107, 168, 99, 12, 93, 118, 128, 0}))
    		})

            It("Should increase the total DAI supply", func() {
                ts, err := Dai.TotalSupply(nil)
                Expect(err).ToNot(HaveOccurred())
                Expect(ts.String()).To(Equal(EthToWei(3).String()))
    		})

    		It("Should increase the random acount's DAI balance", func() {
                bal, err := Dai.BalanceOf(nil, RandomAccount.Address())
                Expect(err).ToNot(HaveOccurred())
                Expect(bal.String()).To(Equal(EthToWei(3).String()))
    		})

    	})
    })
})
