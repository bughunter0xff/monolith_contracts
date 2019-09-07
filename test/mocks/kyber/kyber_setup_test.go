package kyber_test

import (
    "context"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
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

    It("should set the right fee in Bps for the KNC wallet in FeeBurner", func() {
        bps, err := FeeBurner.WalletFeesInBps(nil, KNCWallet.Address())
        Expect(err).ToNot(HaveOccurred())
        Expect(bps.String()).To(Equal("25"))
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

    It("should set the correct kyber network in proxy", func() {
        nr, err := KyberNetwork.GetNumReserves(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(nr.String()).To(Equal("1"))
    })

    It("should enable trades", func() {
        te, err := KyberReserve.TradeEnabled(nil)
        Expect(err).ToNot(HaveOccurred())
        Expect(te).To(BeTrue())
    })

    It("should increase kyber reserve's TKN allowance ", func() {
        b, e := KyberReserve.GetBalance(nil, TKNBurnerAddress)
        Expect(e).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal("3700000000000"))
    })

    It("should increase kyber reserve's ETH balance ", func() {
        b, e := KyberReserve.GetBalance(nil, common.HexToAddress("0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"))
        Expect(e).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal(EthToWei(100).String()))
    })

    It("should set the TKN wallet ", func() {
        tw, e := KyberReserve.TokenWallet(nil, TKNBurnerAddress)
        Expect(e).ToNot(HaveOccurred())
        Expect(tw).To(Equal(TKNWallet.Address()))
    })

    It("should increase the ETH balance of the reserve contract by 100 ETH", func() {
        b, e := Backend.BalanceAt(context.Background(), KyberReserveAddress, nil)
        Expect(e).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal(EthToWei(100).String()))
    })

    It("should increase the TKN balance TKN wallet", func() {
        b, err := TKNBurner.BalanceOf(nil, TKNWallet.Address())
        Expect(err).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal("3800000000000"))
    })

    It("should increase the KNC balance of the KNC wallet", func() {
        b, err := KNCBurner.BalanceOf(nil, KNCWallet.Address())
        Expect(err).ToNot(HaveOccurred())
        Expect(b.String()).To(Equal(EthToWei(38000).String()))
    })

    It("should get the right conversion rate", func() {
        cr, err := KyberReserve.GetConversionRate(nil, common.HexToAddress("0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), TKNBurnerAddress, EthToWei(1), big.NewInt(1))
        Expect(err).ToNot(HaveOccurred())
        Expect(cr.String()).To(Equal("394891340376052539795"))
    })

    It("should get the right conversion rate", func() {
        cr, err := KyberReserve.GetConversionRate(nil, common.HexToAddress("0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), TKNBurnerAddress, EthToWei(1), big.NewInt(1))
        Expect(err).ToNot(HaveOccurred())
        Expect(cr.String()).To(Equal("394891340376052539795"))
    })

    It("should get the right dest quantity", func() {
        cr, err := KyberReserve.GetConversionRate(nil, common.HexToAddress("0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), TKNBurnerAddress, EthToWei(1), big.NewInt(1))
        qty, err := KyberReserve.GetDestQty(nil, common.HexToAddress("0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), TKNBurnerAddress, EthToWei(1), cr)
        Expect(err).ToNot(HaveOccurred())
        Expect(qty.String()).To(Equal("39489134037"))
    })


    It("should find", func() {
        br, err := KyberNetwork.FindBestRate(nil, common.HexToAddress("0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), TKNBurnerAddress, EthToWei(1))
        Expect(err).ToNot(HaveOccurred())
        Expect(br.Obsolete.String()).To(Equal("0"))
        Expect(br.Rate.String()).To(Equal("394891340370000000000"))
    })

    It("should find", func() {
        reserve, rate, err := KyberNetwork.SearchBestRate(nil, common.HexToAddress("0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), TKNBurnerAddress, EthToWei(1), false)
        Expect(err).ToNot(HaveOccurred())
        Expect(reserve).To(Equal(KyberReserveAddress))
        Expect(rate.String()).To(Equal("394891340376052539795"))
    })

})
