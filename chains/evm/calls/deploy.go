package calls

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ChainSafe/chainbridge-core/chains/evm/evmclient"
	"github.com/ChainSafe/chainbridge-core/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

const DefaultGasLimit = 2000000

//type TxFabric func(nonce uint64, to *common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) evmclient.CommonTransaction
type TxFabric func(chainId *big.Int, nonce uint64, to *common.Address, amount *big.Int, gasLimit uint64, gasPricer evmclient.GasPricer, data []byte) (evmclient.CommonTransaction, error)

func DeployErc20(c *evmclient.EVMClient, gasPricer evmclient.GasPricer, txFabric TxFabric, name, symbol string) (common.Address, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PresetMinterPauserABI))
	if err != nil {
		return common.Address{}, err
	}
	address, err := deployContract(c, gasPricer, parsed, common.FromHex(ERC20PresetMinterPauserBin), txFabric, name, symbol)
	if err != nil {
		return common.Address{}, err
	}
	return address, nil
}

func DeployBridge(c *evmclient.EVMClient, gasPricer evmclient.GasPricer, txFabric TxFabric, chainID uint8, relayerAddrs []common.Address, initialRelayerThreshold *big.Int) (common.Address, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, err
	}
	address, err := deployContract(c, gasPricer, parsed, common.FromHex(BridgeBin), txFabric, chainID, relayerAddrs, initialRelayerThreshold, big.NewInt(0), big.NewInt(100))
	if err != nil {
		return common.Address{}, err
	}
	return address, nil
}

func DeployErc20Handler(c *evmclient.EVMClient, gasPricer evmclient.GasPricer, txFabric TxFabric, bridgeAddress common.Address) (common.Address, error) {
	log.Debug().Msgf("Deployng ERC20 Handler with params: %s", bridgeAddress.String())
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, err
	}
	address, err := deployContract(c, gasPricer, parsed, common.FromHex(ERC20HandlerBin), txFabric, bridgeAddress, [][32]byte{}, []common.Address{}, []common.Address{})
	if err != nil {
		return common.Address{}, err
	}
	return address, nil
}

func deployContract(client ChainClient, gasPricer evmclient.GasPricer, abi abi.ABI, bytecode []byte, txFabric TxFabric, params ...interface{}) (common.Address, error) {
	client.LockNonce()
	n, err := client.UnsafeNonce()
	if err != nil {
		return common.Address{}, err
	}
	input, err := abi.Pack("", params...)
	if err != nil {
		return common.Address{}, err
	}

	tx, err := txFabric(nil, n.Uint64(), nil, big.NewInt(0), config.DefaultGasLimit, gasPricer, append(bytecode, input...))
	if err != nil {
		return common.Address{}, err
	}

	hash, err := client.SignAndSendTransaction(context.TODO(), tx)
	if err != nil {
		return common.Address{}, err
	}
	time.Sleep(2 * time.Second)
	_, err = client.WaitAndReturnTxReceipt(tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	log.Debug().Str("hash", hash.String()).Uint64("nonce", n.Uint64()).Msgf("Contract deployed")
	address := crypto.CreateAddress(client.From(), n.Uint64())
	err = client.UnsafeIncreaseNonce()
	if err != nil {
		return common.Address{}, err
	}
	client.UnlockNonce()
	// checks bytecode at address
	// nil is latest block
	if code, err := client.CodeAt(context.Background(), address, nil); err != nil {
		return common.Address{}, err
	} else if len(code) == 0 {
		return common.Address{}, fmt.Errorf("no code at provided address %s", address.String())
	}
	return address, nil
}