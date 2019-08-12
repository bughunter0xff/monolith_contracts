#!/bin/bash

set -e -o pipefail

SOLC_LATEST="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.5.10 --optimize /=/"
SOLC_0_4_25="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.4.25 --optimize /=/"
SOLC_0_4_19="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.4.19 --optimize /=/"
SOLC_0_4_10="docker run --rm -u `id -u` -v $PWD:/solidity --workdir /solidity/contracts ethereum/solc:0.4.11 --optimize /=/"


compile_solidity() {
  echo "compiling ${2}"
  ${1} --overwrite --bin --abi ${2}.sol -o /solidity/build/${2} --combined-json bin-runtime,srcmap-runtime,ast,srcmap,bin
}

contract_sources=(
  'wallet'
  'oracle'
  'licence'
  'holder'
  'controller'
  'tokenWhitelist'
  'walletDeployer'
  'mocks/token'
  'mocks/burnerToken'
  'mocks/nonCompliantToken'
  'mocks/base64Exporter'
  'mocks/oraclize'
  'mocks/bytesUtilsExporter'
  'mocks/parseIntScientificExporter'
  'mocks/tokenWhitelistableExporter'
  'internals/tokenWhitelistable'
  'internals/parseIntScientific'
  'externals/ens/PublicResolver'
  'externals/ens/ENSRegistry'
)

makerDAO_sources=(
    'mocks/makerDAO/SaiProxyCreateAndExecute'
    'mocks/makerDAO/ProxyRegistry'
    'mocks/makerDAO/Medianizer'
    'mocks/makerDAO/SaiTub'
    'mocks/makerDAO/SaiVox'
    'mocks/makerDAO/MedianizerNew'
    'mocks/makerDAO/MKR'
    'mocks/makerDAO/Dai'
    'mocks/makerDAO/PETH'
    'mocks/makerDAO/WETH'
)

for c in "${contract_sources[@]}"
do
    compile_solidity "${SOLC_LATEST}" $c
done

for c in "${makerDAO_sources[@]:0:2}"
do
    compile_solidity "$SOLC_0_4_25" $c
done

for c in "${makerDAO_sources[2]}"
do
    compile_solidity "$SOLC_0_4_10" $c 1>/dev/null
done

for c in "${makerDAO_sources[@]:3}"
do
    compile_solidity "$SOLC_0_4_19" $c
done

GE_PATH=${PWD}/vendor/github.com/ethereum/go-ethereum

# Generate Go bindings from solidity contracts.
ABIGEN="docker run --rm -u `id -u` --workdir /go/src/github/tokencard/contracts -e GOPATH=/go -v $GE_PATH:/go/src/github.com/ethereum/go-ethereum -v $PWD:/go/src/github/tokencard/contracts ethereum/client-go:alltools-v1.8.21 abigen"

generate_binding() {
  contract=$(echo $1 | awk '{print $1}')
  go_source=$(echo $1 | awk '{print $2}')
  go_type=$(echo $1 | awk '{print $3}')
  package=$(echo $1 | awk '{print $4}')
  echo "Generating binding for ${go_type} (${contract})"
  ${ABIGEN} --abi ./build/${contract}.abi  --bin ./build/${contract}.bin --pkg ${package} --type=${go_type} --out ./pkg/bindings/${go_source}
}

contracts=(
  "wallet/Wallet wallet.go Wallet bindings"
  "oracle/Oracle oracle.go Oracle bindings"
  "licence/Licence licence.go Licence bindings"
  "holder/Holder holder.go Holder bindings"
  "controller/Controller controller.go Controller bindings"
  "tokenWhitelist/TokenWhitelist tokenWhitelist.go TokenWhitelist bindings"
  "walletDeployer/WalletDeployer walletDeployer.go WalletDeployer bindings"
  "mocks/token/Token mocks/token.go Token mocks"
  "mocks/burnerToken/BurnerToken mocks/burnerToken.go BurnerToken mocks"
  "mocks/nonCompliantToken/NonCompliantToken mocks/nonCompliantToken.go NonCompliantToken mocks"
  "mocks/base64Exporter/Base64Exporter mocks/base64Exporter.go Base64Exporter mocks"
  "mocks/oraclize/OraclizeConnector mocks/oraclizeConnector.go OraclizeConnector mocks"
  "mocks/oraclize/OraclizeAddrResolver mocks/oraclizeAddrResolver.go OraclizeAddrResolver mocks"
  "mocks/bytesUtilsExporter/BytesUtilsExporter mocks/bytesUtilsExporter.go BytesUtilsExporter mocks"
  "mocks/parseIntScientificExporter/ParseIntScientificExporter mocks/parseIntScientificExporter.go ParseIntScientificExporter mocks"
  "mocks/tokenWhitelistableExporter/TokenWhitelistableExporter mocks/tokenWhitelistableExporter.go TokenWhitelistableExporter mocks"
  "internals/tokenWhitelistable/TokenWhitelistable internals/tokenWhitelistable.go TokenWhitelistable internals"
  "internals/parseIntScientific/ParseIntScientific internals/parseIntScientific.go ParseIntScientific internals"
  "externals/ens/ENSRegistry/ENSRegistry externals/ens/ENSRegistry.go ENSRegistry ens"
  "externals/ens/PublicResolver/PublicResolver externals/ens/PublicResolver.go PublicResolver ens"

)

makerDAO_contracts=(
    "mocks/makerDAO/SaiProxyCreateAndExecute/SaiProxyCreateAndExecute mocks/makerDAO/SaiProxyCreateAndExecute.go SaiProxyCreateAndExecute makerDAO"
    "mocks/makerDAO/ProxyRegistry/ProxyRegistry mocks/makerDAO/ProxyRegistry.go ProxyRegistry makerDAO"
    "mocks/makerDAO/ProxyRegistry/DSProxyFactory mocks/makerDAO/DSProxyFactory.go DSProxyFactory makerDAO"
    "mocks/makerDAO/ProxyRegistry/DSProxyFactory mocks/makerDAO/DSProxyFactory.go DSProxyFactory makerDAO"
    "mocks/makerDAO/SaiVox/SaiVox mocks/makerDAO/SaiVox.go SaiVox makerDAO"
    "mocks/makerDAO/Dai/Dai mocks/makerDAO/Dai.go Dai makerDAO"
    "mocks/makerDAO/MKR/MKR mocks/makerDAO/MKR.go MKR makerDAO"
    "mocks/makerDAO/Medianizer/Medianizer mocks/makerDAO/Medianizer.go Medianizer makerDAO"
    "mocks/makerDAO/MedianizerNew/Medianizer mocks/makerDAO/MedianizerNew.go MedianizerNew makerDAO"
    "mocks/makerDAO/SaiTub/SaiTub mocks/makerDAO/SaiTub.go SaiTub makerDAO"
    "mocks/makerDAO/WETH/WETH mocks/makerDAO/WETH.go WETH makerDAO"
    "mocks/makerDAO/PETH/PETH mocks/makerDAO/PETH.go PETH makerDAO"
)

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

for c in "${makerDAO_contracts[@]}"
do
    generate_binding "$c"
done

echo "done"
