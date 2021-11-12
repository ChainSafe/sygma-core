package bridge

import (
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/flags"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var registerGenericResourceCmd = &cobra.Command{
	Use:   "register-generic-resource",
	Short: "Register a generic resource ID",
	Long:  "Register a resource ID with a contract address for a generic handler",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.LoggerMetadata(cmd.Name(), cmd.Flags())
	},
	Run: registerGenericResource,
}

func BindRegisterGenericResourceFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&Handler, "handler", "", "handler contract address")
	cmd.Flags().StringVar(&ResourceID, "resourceId", "", "resource ID to query")
	cmd.Flags().StringVar(&Bridge, "bridge", "", "bridge contract address")
	cmd.Flags().StringVar(&Target, "target", "", "contract address to be registered") // TODO change the description (target is not necessary a contract address, could be hash storage)
	cmd.Flags().StringVar(&Deposit, "deposit", "0x00000000", "deposit function signature")
	cmd.Flags().StringVar(&Execute, "execute", "0x00000000", "execute proposal function signature")
	cmd.Flags().BoolVar(&Hash, "hash", false, "treat signature inputs as function prototype strings, hash and take the first 4 bytes")
	flags.MarkFlagsAsRequired(cmd, "handler", "resourceId", "bridge", "target")
}

func init() {
	BindRegisterGenericResourceFlags(registerGenericResourceCmd)
}

func registerGenericResource(cmd *cobra.Command, args []string) {
	log.Debug().Msgf(`
Registering generic resource
Handler address: %s
Resource ID: %s
Bridge address: %s
Target address: %s
Deposit: %s
Execute: %s
Hash: %v
`, Handler, ResourceID, Bridge, Target, Deposit, Execute, Hash)
}

/*
func registerGenericResource(cctx *cli.Context) error {
	url := cctx.String("url")
	gasLimit := cctx.Int64("gasLimit")
	gasPrice := cctx.Int64("gasPrice")

	depositSig := cctx.String("deposit")
	depositSigBytes := common.Hex2Bytes(depositSig)
	depositSigBytesArr := utils.SliceTo4Bytes(depositSigBytes)

	executeSig := cctx.String("execute")
	executeSigBytes := common.Hex2Bytes(executeSig)
	executeSigBytesArr := utils.SliceTo4Bytes(executeSigBytes)

	if cctx.Bool("hash") {
		depositSigBytesArr = utils.GetSolidityFunctionSig(depositSig)
		executeSigBytesArr = utils.GetSolidityFunctionSig(executeSig)
	}

	sender, err := cliutils.DefineSender(cctx)
	if err != nil {
		return err
	}

	bridgeAddress, err := cliutils.DefineBridgeAddress(cctx)
	if err != nil {
		return err
	}

	handler := cctx.String("handler")
	if !common.IsHexAddress(handler) {
		return fmt.Errorf("invalid handler address %s", handler)
	}
	handlerAddress := common.HexToAddress(handler)
	targetContract := cctx.String("targetContract")
	if !common.IsHexAddress(targetContract) {
		return fmt.Errorf("invalid targetContract address %s", targetContract)
	}
	targetContractAddress := common.HexToAddress(targetContract)
	resourceId := cctx.String("resourceId")
	resourceIdBytes := common.Hex2Bytes(resourceId)
	resourceIdBytesArr := utils.SliceTo32Bytes(resourceIdBytes)

	log.Info().Msgf("Registering contract %s with resource ID %s on handler %s", targetContract, resourceId, handler)
	ethClient, err := client.NewClient(url, false, sender, big.NewInt(gasLimit), big.NewInt(gasPrice), big.NewFloat(1))
	if err != nil {
		return err
	}
	err = utils.RegisterGenericResource(ethClient, bridgeAddress, handlerAddress, resourceIdBytesArr, targetContractAddress, depositSigBytesArr, executeSigBytesArr)
	if err != nil {
		return err
	}
	fmt.Println("Resource registered")
	return nil
}
*/
