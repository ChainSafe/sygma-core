package erc721

import (
	"fmt"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/contracts/erc721"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/initialize"
	"github.com/ChainSafe/chainbridge-core/chains/evm/evmtransaction"
	"github.com/ChainSafe/chainbridge-core/util"
	"math/big"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/flags"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var approveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve token in an ERC721 contract for transfer.",
	Long:  "Approve token in an ERC721 contract for transfer.",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.LoggerMetadata(cmd.Name(), cmd.Flags())
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return util.CallPersistentPreRun(cmd, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := initialize.InitializeClient(url, senderKeyPair)
		if err != nil {
			return err
		}
		t, err := initialize.InitializeTransactor(gasPrice, evmtransaction.NewTransaction, c)
		if err != nil {
			return err
		}
		return ApproveCmd(cmd, args, erc721.NewErc721Contract(c, erc721Addr, t))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		err := ValidateApproveFlags(cmd, args)
		if err != nil {
			return err
		}

		err = ProcessApproveFlags(cmd, args)
		return err
	},
}

func BindApproveCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&Erc721Address, "contract-address", "", "address of contract")
	cmd.Flags().StringVar(&Recipient, "recipient", "", "address of recipient")
	cmd.Flags().StringVar(&TokenId, "tokenId", "", "ERC721 token ID")
	flags.MarkFlagsAsRequired(cmd, "contract-address", "recipient", "tokenId")
}

func init() {
	BindApproveCmdFlags(approveCmd)
}

func ValidateApproveFlags(cmd *cobra.Command, args []string) error {
	if !common.IsHexAddress(Erc721Address) {
		return fmt.Errorf("invalid ERC721 contract address %s", Erc721Address)
	}
	if !common.IsHexAddress(Recipient) {
		return fmt.Errorf("invalid recipient address")
	}
	return nil
}

func ProcessApproveFlags(cmd *cobra.Command, args []string) error {
	recipientAddr = common.HexToAddress(Recipient)
	erc721Addr = common.HexToAddress(Erc721Address)

	var ok bool
	if tokenId, ok = big.NewInt(0).SetString(TokenId, 10); !ok {
		return fmt.Errorf("invalid token id value")
	}
	return nil
}

func ApproveCmd(cmd *cobra.Command, args []string, erc721Contract *erc721.ERC721Contract) error {
	_, err = erc721Contract.Approve(
		tokenId, recipientAddr, transactor.TransactOptions{GasLimit: gasLimit},
	)
	if err != nil {
		return err
	}

	log.Info().Msgf("%v token approved", tokenId)
	return err
}
