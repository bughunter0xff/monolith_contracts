package kyber_test

import (
    // "math/big"
    // "strings"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// . "github.com/tokencard/contracts/v2/test/shared"
	// "github.com/ethereum/go-ethereum/accounts/abi"
    // "github.com/ethereum/go-ethereum/common"
)

var _ = Describe("kyber setup test", func() {

    It("should set the correct kyber network in proxy", func() {
        kn, err := KyberNetworkProxy.KyberNetworkContract(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(kn).To(Equal(KyberNetworkAddress))
    })

    It("should set the right kncPerEthRatePrecision in FeeBurner", func() {
        prec, err := FeeBurner.KncPerEthRatePrecision(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(prec.String()).To(Equal("1273294871580578838478"))
    })

    It("should set the correct kyber network proxy", func() {
        knp, err := KyberNetwork.KyberNetworkProxyContract(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(knp).To(Equal(KyberNetworkProxyAddress))
    })

    It("should set the correct fee burner contract", func() {
        fb, err := KyberNetwork.FeeBurnerContract(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(fb).To(Equal(FeeBurnerAddress))
    })

    It("should set the correct expected rate contract", func() {
        er, err := KyberNetwork.ExpectedRateContract(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(er).To(Equal(ExpectedRateAddress))
    })

    It("should set the right max gas price", func() {
        mgs, err := KyberNetwork.MaxGasPriceValue(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(mgs.String()).To(Equal("100000000000"))
    })

    It("should set the right negligible rate diff", func() {
        nrd, err := KyberNetwork.NegligibleRateDiff(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(nrd.String()).To(Equal("20"))
    })

    It("should enable KyberNetwork", func() {
        en, err := KyberNetwork.Enabled(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(en).To(BeTrue())
    })


})
