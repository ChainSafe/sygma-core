package itx_test

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/consts"
	"github.com/ChainSafe/chainbridge-core/chains/evm/transactor"
	"github.com/ChainSafe/chainbridge-core/chains/evm/transactor/itx"
	mock_itx "github.com/ChainSafe/chainbridge-core/chains/evm/transactor/itx/mock"
	"github.com/ChainSafe/chainbridge-core/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type TransactTestSuite struct {
	suite.Suite
	forwarder   *mock_itx.MockForwarder
	relayCaller *mock_itx.MockRelayCaller
	transactor  *itx.ITXTransactor
	kp          *secp256k1.Keypair
}

func TestRunTransactTestSuite(t *testing.T) {
	suite.Run(t, new(TransactTestSuite))
}

func (s *TransactTestSuite) SetupSuite()    {}
func (s *TransactTestSuite) TearDownSuite() {}
func (s *TransactTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.kp, _ = secp256k1.NewKeypairFromPrivateKey(common.Hex2Bytes("e8e0f5427111dee651e63a6f1029da6929ebf7d2d61cefaf166cebefdf2c012e"))
	s.forwarder = mock_itx.NewMockForwarder(gomockController)
	s.relayCaller = mock_itx.NewMockRelayCaller(gomockController)
	s.transactor = itx.NewITXTransactor(s.relayCaller, s.forwarder, s.kp)
	s.forwarder.EXPECT().ChainId().Return(big.NewInt(5))
}
func (s *TransactTestSuite) TearDownTest() {}

func (s *TransactTestSuite) TestTransact_FailedFetchingForwarderData() {
	to := common.HexToAddress("0x04005C8A516292af163b1AFe3D855b9f4f4631B5")
	data := []byte{}
	opts := transactor.TransactOptions{
		GasLimit: big.NewInt(200000),
		GasPrice: big.NewInt(1),
		Priority: "slow",
		Value:    big.NewInt(0),
		ChainID:  big.NewInt(5),
	}
	s.forwarder.EXPECT().ForwarderData(to, data, opts).Return(nil, errors.New("error"))

	_, err := s.transactor.Transact(to, data, opts)

	s.NotNil(err)
}

func (s *TransactTestSuite) TestTransact_FailedSendTransaction() {
	to := common.HexToAddress("0x04005C8A516292af163b1AFe3D855b9f4f4631B5")
	data := []byte{}
	opts := transactor.TransactOptions{
		GasLimit: big.NewInt(200000),
		GasPrice: big.NewInt(1),
		Priority: "slow",
		Value:    big.NewInt(0),
		ChainID:  big.NewInt(5),
	}
	s.forwarder.EXPECT().ForwarderData(to, data, opts).Return([]byte{}, nil)
	s.forwarder.EXPECT().ForwarderAddress().Return(to)
	s.relayCaller.EXPECT().CallContext(
		context.Background(),
		gomock.Any(),
		"relay_sendTransaction",
		gomock.Any(),
		gomock.Any(),
	).Return(errors.New("error"))

	_, err := s.transactor.Transact(to, data, opts)

	s.NotNil(err)
}

func (s *TransactTestSuite) TestTransact_SuccessfulSend() {
	to := common.HexToAddress("0x04005C8A516292af163b1AFe3D855b9f4f4631B5")
	data := []byte{}
	opts := transactor.TransactOptions{
		GasLimit: big.NewInt(200000),
		GasPrice: big.NewInt(1),
		Priority: "slow",
		Value:    big.NewInt(0),
		ChainID:  big.NewInt(5),
	}
	expectedSig := "0x68ad089b7daeabcdd76520377822cc32eba0b41ea702358bc8f7249bc296d408781eb60366a3bb6ad9fc62dca08bdf440a7c4f02e3680aa0b477a2dd5423d5af01"

	s.forwarder.EXPECT().ForwarderData(to, data, opts).Return([]byte{}, nil)
	s.forwarder.EXPECT().ForwarderAddress().Return(to)
	s.relayCaller.EXPECT().CallContext(
		context.Background(),
		gomock.Any(),
		"relay_sendTransaction",
		gomock.Any(),
		expectedSig,
	).Return(nil)

	hash, err := s.transactor.Transact(to, data, opts)

	s.Nil(err)
	s.NotNil(hash)
}

func (s *TransactTestSuite) TestTransact_SuccessfulSendWithDefaultOpts() {
	to := common.HexToAddress("0x04005C8A516292af163b1AFe3D855b9f4f4631B5")
	data := []byte{}
	expectedOpts := transactor.TransactOptions{
		GasLimit: big.NewInt(consts.DefaultGasLimit),
		GasPrice: big.NewInt(1),
		Priority: "slow",
		Value:    big.NewInt(0),
		ChainID:  big.NewInt(5),
	}
	expectedSig := "0x62ffe54c569cf70e7f62c410eff8f200dd46f0209bfa134484655c0621cc87c46c300cebc8e401b52acb04a5e00fa697be410458175d600c2471d5d7f986dd8501"

	s.forwarder.EXPECT().ForwarderData(to, data, expectedOpts).Return([]byte{}, nil)
	s.forwarder.EXPECT().ForwarderAddress().Return(to)
	s.relayCaller.EXPECT().CallContext(
		context.Background(),
		gomock.Any(),
		"relay_sendTransaction",
		gomock.Any(),
		expectedSig,
	).Return(nil)

	hash, err := s.transactor.Transact(to, data, transactor.TransactOptions{})

	s.Nil(err)
	s.NotNil(hash)
}
