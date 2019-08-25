package makerDAO_test

import (
	"context"
	"os"
    "fmt"
	"testing"
    "math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tokencard/contracts/v2/pkg/bindings/mocks/makerDAO"
    "github.com/tokencard/contracts/v2/pkg/bindings"
    . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    . "github.com/tokencard/contracts/v2/test/shared"
)

var SaiProxy *makerDAO.SaiProxyCreateAndExecute
var SaiProxyAddress common.Address

var SaiTub *makerDAO.SaiTub
var SaiTubAddress common.Address

var SaiVox *makerDAO.SaiVox
var SaiVoxAddress common.Address

var Dai *makerDAO.Dai
var DaiAddress common.Address

var Sin *makerDAO.Dai
var SinAddress common.Address

var Skr *makerDAO.Dai
var SkrAddress common.Address

var MKR *makerDAO.MKR
var MKRAddress common.Address

var WETH *makerDAO.WETH
var WETHAddress common.Address

var DSProxyFactory *makerDAO.DSProxyFactory
var DSProxyFactoryAddress common.Address

var ProxyRegistry *makerDAO.ProxyRegistry
var ProxyRegistryAddress common.Address

var Pip *makerDAO.Medianizer
var PipAddress common.Address

var Pep *makerDAO.MedianizerNew
var PepAddress common.Address

var GemPit *makerDAO.Dai
var GemPitAddress common.Address

var DSGuard *makerDAO.DSGuard
var DSGuardAddress common.Address

var Wallet *bindings.Wallet
var WalletAddress common.Address

func init() {
	TestRig.AddCoverageForContracts(
		"../../../build/mocks/makerDAO/SaiProxyCreateAndExecute/combined.json",
		"../../../contracts",
	)

    TestRig.AddCoverageForContracts(
		"../../../build/mocks/makerDAO/SaiTub/combined.json",
		"../../../contracts",
	)
}

func TestTokenWhitelistSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "makerDAO Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

    var tx *types.Transaction

	SaiProxyAddress, tx, SaiProxy, err = makerDAO.DeploySaiProxyCreateAndExecute(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    DSGuardAddress, tx, DSGuard, err = makerDAO.DeployDSGuard(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    DSProxyFactoryAddress, tx, DSProxyFactory, err = makerDAO.DeployDSProxyFactory(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    ProxyRegistryAddress, tx, ProxyRegistry, err = makerDAO.DeployProxyRegistry(BankAccount.TransactOpts(), Backend, DSProxyFactoryAddress)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    DSGuardAddress, tx, DSGuard, err = makerDAO.DeployDSGuard(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    DaiAddress, tx, Dai, err = makerDAO.DeployDai(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    SinAddress, tx, Sin, err = makerDAO.DeployDai(BankAccount.TransactOpts(), Backend)
    Expect(err).ToNot(HaveOccurred())
    Backend.Commit()
    Expect(isSuccessful(tx)).To(BeTrue())

    SkrAddress, tx, Skr, err = makerDAO.DeployDai(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    WETHAddress, tx, WETH, err = makerDAO.DeployWETH(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    MKRAddress, tx, MKR, err = makerDAO.DeployMKR(BankAccount.TransactOpts(), Backend)
    Expect(err).ToNot(HaveOccurred())
    Backend.Commit()
    Expect(isSuccessful(tx)).To(BeTrue())

    PipAddress, tx, Pip, err = makerDAO.DeployMedianizer(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    PepAddress, tx, Pep, err = makerDAO.DeployMedianizerNew(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    GemPitAddress, tx, GemPit, err = makerDAO.DeployDai(BankAccount.TransactOpts(), Backend)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    par := new(big.Int)
    par.SetString("1000000000000000000000000000",10)
    SaiVoxAddress, tx, SaiVox, err = makerDAO.DeploySaiVox(BankAccount.TransactOpts(), Backend, par)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    SaiTubAddress, tx, SaiTub, err = makerDAO.DeploySaiTub(BankAccount.TransactOpts(), Backend, DaiAddress, SinAddress, SkrAddress, WETHAddress, MKRAddress, PipAddress, PepAddress, SaiVoxAddress, GemPitAddress)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    WalletAddress, tx, Wallet, err = bindings.DeployWallet(BankAccount.TransactOpts(), Backend, Owner.Address(), true, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(1))
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})

var _ = AfterEach(func() {
    td := CurrentGinkgoTestDescription()

	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())

})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("mocks/makerDAO/SaiProxyCreateAndExecute.sol", 0.0)
	TestRig.PrintGasUsage(os.Stdout)
})

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}
