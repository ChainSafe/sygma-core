package erc721

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/flags"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/utils"
	"github.com/ChainSafe/chainbridge-core/chains/evm/evmclient"
	"github.com/ChainSafe/chainbridge-core/chains/evm/evmgaspricer"
	"github.com/ChainSafe/chainbridge-core/chains/evm/evmtransaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var mintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Mint token on an ERC721 mintable contract",
	Long:  "Mint token on an ERC721 mintable contract",
	RunE: func(cmd *cobra.Command, args []string) error {
		txFabric := evmtransaction.NewTransaction
		return MintCmd(cmd, args, txFabric, &evmgaspricer.LondonGasPriceDeterminant{})
	},
	Args: func(cmd *cobra.Command, args []string) error {
		err := ValidateMintFlags(cmd, args)
		if err != nil {
			return err
		}

		err = ProcessMintFlags(cmd, args)
		return err
	},
}

func init() {
	BindMintFlags()
}

func BindMintFlags() {
	mintCmd.Flags().StringVar(&Erc721Address, "contract-address", "", "address of contract")
	mintCmd.Flags().StringVar(&DstAddress, "destination-address", "", "address of recipient")
	mintCmd.Flags().StringVar(&TokenId, "tokenId", "", "ERC721 token ID")
	mintCmd.Flags().StringVar(&Metadata, "metadata", "", "ERC721 token metadata")
	flags.MarkFlagsAsRequired(mintCmd, "contract-address", "destination-address", "tokenId", "metadata", "contract-address")
}

func ValidateMintFlags(cmd *cobra.Command, args []string) error {
	if !common.IsHexAddress(Erc721Address) {
		return fmt.Errorf("invalid ERC721 contract address %s", Erc721Address)
	}
	if !common.IsHexAddress(DstAddress) {
		return fmt.Errorf("invalid recipient address %s", DstAddress)
	}
	return nil
}

func ProcessMintFlags(cmd *cobra.Command, args []string) error {
	var err error
	erc721Addr = common.HexToAddress(Erc721Address)

	if !common.IsHexAddress(DstAddress) {
		dstAddress = senderKeyPair.CommonAddress()
	} else {
		dstAddress = common.HexToAddress(DstAddress)
	}

	var ok bool
	if tokenId, ok = big.NewInt(0).SetString(TokenId, 10); !ok {
		return fmt.Errorf("invalid token id value")
	}

	if Metadata[0:2] == "0x" {
		Metadata = Metadata[2:]
	}
	metadata, err = hex.DecodeString(Metadata)
	return err
}

func MintCmd(cmd *cobra.Command, args []string, txFabric calls.TxFabric, gasPricer utils.GasPricerWithPostConfig) error {
	ethClient, err := evmclient.NewEVMClientFromParams(
		url, senderKeyPair.PrivateKey())
	if err != nil {
		log.Error().Err(fmt.Errorf("eth client intialization error: %v", err))
		return err
	}

	gasPricer.SetClient(ethClient)
	gasPricer.SetOpts(&evmgaspricer.GasPricerOpts{UpperLimitFeePerGas: gasPrice})

	_, err = calls.ERC721Mint(ethClient, txFabric, gasPricer.(calls.GasPricer), gasLimit, tokenId, metadata, erc721Addr, dstAddress)
	if err != nil {
		return err
	}

	log.Info().Msgf("%v token minted", tokenId)
	return err
}
