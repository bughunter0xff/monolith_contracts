package kyber_test

import (
    "context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tokencard/contracts/v2/test/shared"
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

    It("should set the right Bps KNC wallet correspondence in FeeBurner", func() {
        bps, err := FeeBurner.ReserveFeesInBps(nil, KyberReserveAddress)
        Expect(err).ToNot(HaveOccurred())
        Expect(bps.String()).To(Equal("25"))
    })

    It("should set the right reserve/KNCWallet correspondence in FeeBurner", func() {
        kncw, err := FeeBurner.ReserveKNCWallet(nil, KyberReserveAddress)
        Expect(err).ToNot(HaveOccurred())
        Expect(kncw).To(Equal(KNCWallet.Address()))
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

    It("should enable trades", func() {
        te, err := KyberReserve.TradeEnabled(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(te).To(BeTrue())
    })

    It("should enable trades", func() {
        te, err := KyberReserve.TradeEnabled(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(te).To(BeTrue())
    })

    It("should increase the ETH balance of the reserve contract by 100 ETH", func() {
        b, e := Backend.BalanceAt(context.Background(), KyberReserveAddress, nil)
        Expect(e).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal(EthToWei(100).String()))
    })

    It("should increase TKN balance of the reserve", func() {
        b, err := TKNBurner.BalanceOf(nil, KyberReserveAddress)
        Expect(err).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal("38000"))
    })

    It("should increase KNC balance of the KNC wallet", func() {
        b, err := KNCBurner.BalanceOf(nil, KNCWallet.Address())
        Expect(err).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal("38000"))
    })
})
