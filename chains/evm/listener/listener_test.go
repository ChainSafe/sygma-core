package listener_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ChainSafe/chainbridge-core/chains/evm/listener"
	mock_listener "github.com/ChainSafe/chainbridge-core/chains/evm/listener/mock"
	"github.com/ChainSafe/chainbridge-core/relayer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

var errIncorrectCalldataLen = errors.New("invalid calldata length: less than 84 bytes")

type ListenerTestSuite struct {
	suite.Suite
	mockEventHandler *mock_listener.MockEventHandler
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(ListenerTestSuite))
}

func (s *ListenerTestSuite) SetupSuite()    {}
func (s *ListenerTestSuite) TearDownSuite() {}
func (s *ListenerTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.mockEventHandler = mock_listener.NewMockEventHandler(gomockController)
}
func (s *ListenerTestSuite) TearDownTest() {}

func (s *ListenerTestSuite) TestErc20HandleEvent() {
	// 0xf1e58fb17704c2da8479a533f9fad4ad0993ca6b
	recipientByteSlice := []byte{241, 229, 143, 177, 119, 4, 194, 218, 132, 121, 165, 51, 249, 250, 212, 173, 9, 147, 202, 107}

	// construct ERC20 deposit data
	// follows behavior of solidity tests
	// https://github.com/ChainSafe/chainbridge-solidity/blob/develop/test/contractBridge/depositERC20.js#L46-L50
	var calldata []byte
	calldata = append(calldata, math.PaddedBigBytes(big.NewInt(2), 32)...)
	calldata = append(calldata, math.PaddedBigBytes(big.NewInt(int64(len(recipientByteSlice))), 32)...)
	calldata = append(calldata, recipientByteSlice...)

	depositLog := &listener.DepositLogs{
		DestinationID:   0,
		ResourceID:      [32]byte{0},
		DepositNonce:    1,
		SenderAddress:   common.HexToAddress("0x4CEEf6139f00F9F4535Ad19640Ff7A0137708485"),
		Calldata:        calldata,
		HandlerResponse: []byte{},
	}

	sourceID := uint8(1)
	amountParsed := calldata[:32]
	recipientAddressParsed := calldata[65:]

	expected := &relayer.Message{
		Source:       sourceID,
		Destination:  depositLog.DestinationID,
		DepositNonce: depositLog.DepositNonce,
		ResourceId:   depositLog.ResourceID,
		Type:         relayer.FungibleTransfer,
		Payload: []interface{}{
			amountParsed,
			recipientAddressParsed,
		},
	}

	message, err := listener.Erc20EventHandler(sourceID, depositLog.DestinationID, depositLog.DepositNonce, depositLog.ResourceID, depositLog.Calldata, depositLog.HandlerResponse)

	s.Nil(err)

	s.NotNil(message)

	s.Equal(message, expected)
}

func (s *ListenerTestSuite) TestErc20HandleEventIncorrectCalldataLen() {
	// 0xf1e58fb17704c2da8479a533f9fad4ad0993ca6b
	recipientByteSlice := []byte{241, 229, 143, 177, 119, 4, 194, 218, 132, 121, 165, 51, 249, 250, 212, 173, 9, 147, 202, 107}

	// construct ERC20 deposit data
	// follows behavior of solidity tests
	// https://github.com/ChainSafe/chainbridge-solidity/blob/develop/test/contractBridge/depositERC20.js#L46-L50
	var calldata []byte
	calldata = append(calldata, math.PaddedBigBytes(big.NewInt(2), 16)...)
	calldata = append(calldata, math.PaddedBigBytes(big.NewInt(int64(len(recipientByteSlice))), 16)...)
	calldata = append(calldata, recipientByteSlice...)

	depositLog := &listener.DepositLogs{
		DestinationID:   0,
		ResourceID:      [32]byte{0},
		DepositNonce:    1,
		SenderAddress:   common.HexToAddress("0x4CEEf6139f00F9F4535Ad19640Ff7A0137708485"),
		Calldata:        calldata,
		HandlerResponse: []byte{},
	}

	sourceID := uint8(1)

	message, err := listener.Erc20EventHandler(sourceID, depositLog.DestinationID, depositLog.DepositNonce, depositLog.ResourceID, depositLog.Calldata, depositLog.HandlerResponse)

	s.Nil(message)

	s.EqualError(err, errIncorrectCalldataLen.Error())
}
