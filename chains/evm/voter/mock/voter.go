// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChainSafe/chainbridge-core/chains/evm/voter (interfaces: ChainClient,MessageHandler,BridgeContract)

// Package mock_voter is a generated GoMock package.
package mock_voter

import (
	context "context"
	big "math/big"
	reflect "reflect"

	evmclient "github.com/ChainSafe/chainbridge-core/chains/evm/calls/evmclient"
	transactor "github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor"
	proposal "github.com/ChainSafe/chainbridge-core/chains/evm/voter/proposal"
	message "github.com/ChainSafe/chainbridge-core/relayer/message"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	rpc "github.com/ethereum/go-ethereum/rpc"
	gomock "github.com/golang/mock/gomock"
)

// MockChainClient is a mock of ChainClient interface.
type MockChainClient struct {
	ctrl     *gomock.Controller
	recorder *MockChainClientMockRecorder
}

// MockChainClientMockRecorder is the mock recorder for MockChainClient.
type MockChainClientMockRecorder struct {
	mock *MockChainClient
}

// NewMockChainClient creates a new mock instance.
func NewMockChainClient(ctrl *gomock.Controller) *MockChainClient {
	mock := &MockChainClient{ctrl: ctrl}
	mock.recorder = &MockChainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainClient) EXPECT() *MockChainClientMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockChainClient) CallContract(arg0 context.Context, arg1 map[string]interface{}, arg2 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockChainClientMockRecorder) CallContract(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockChainClient)(nil).CallContract), arg0, arg1, arg2)
}

// CodeAt mocks base method.
func (m *MockChainClient) CodeAt(arg0 context.Context, arg1 common.Address, arg2 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockChainClientMockRecorder) CodeAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockChainClient)(nil).CodeAt), arg0, arg1, arg2)
}

// From mocks base method.
func (m *MockChainClient) From() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "From")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// From indicates an expected call of From.
func (mr *MockChainClientMockRecorder) From() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockChainClient)(nil).From))
}

// GetTransactionByHash mocks base method.
func (m *MockChainClient) GetTransactionByHash(arg0 common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", arg0)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash.
func (mr *MockChainClientMockRecorder) GetTransactionByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockChainClient)(nil).GetTransactionByHash), arg0)
}

// LockNonce mocks base method.
func (m *MockChainClient) LockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockNonce")
}

// LockNonce indicates an expected call of LockNonce.
func (mr *MockChainClientMockRecorder) LockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockNonce", reflect.TypeOf((*MockChainClient)(nil).LockNonce))
}

// RelayerAddress mocks base method.
func (m *MockChainClient) RelayerAddress() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelayerAddress")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// RelayerAddress indicates an expected call of RelayerAddress.
func (mr *MockChainClientMockRecorder) RelayerAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelayerAddress", reflect.TypeOf((*MockChainClient)(nil).RelayerAddress))
}

// SignAndSendTransaction mocks base method.
func (m *MockChainClient) SignAndSendTransaction(arg0 context.Context, arg1 evmclient.CommonTransaction) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignAndSendTransaction", arg0, arg1)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignAndSendTransaction indicates an expected call of SignAndSendTransaction.
func (mr *MockChainClientMockRecorder) SignAndSendTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignAndSendTransaction", reflect.TypeOf((*MockChainClient)(nil).SignAndSendTransaction), arg0, arg1)
}

// SubscribePendingTransactions mocks base method.
func (m *MockChainClient) SubscribePendingTransactions(arg0 context.Context, arg1 chan<- common.Hash) (*rpc.ClientSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribePendingTransactions", arg0, arg1)
	ret0, _ := ret[0].(*rpc.ClientSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribePendingTransactions indicates an expected call of SubscribePendingTransactions.
func (mr *MockChainClientMockRecorder) SubscribePendingTransactions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribePendingTransactions", reflect.TypeOf((*MockChainClient)(nil).SubscribePendingTransactions), arg0, arg1)
}

// TransactionByHash mocks base method.
func (m *MockChainClient) TransactionByHash(arg0 context.Context, arg1 common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", arg0, arg1)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockChainClientMockRecorder) TransactionByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockChainClient)(nil).TransactionByHash), arg0, arg1)
}

// UnlockNonce mocks base method.
func (m *MockChainClient) UnlockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockNonce")
}

// UnlockNonce indicates an expected call of UnlockNonce.
func (mr *MockChainClientMockRecorder) UnlockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockNonce", reflect.TypeOf((*MockChainClient)(nil).UnlockNonce))
}

// UnsafeIncreaseNonce mocks base method.
func (m *MockChainClient) UnsafeIncreaseNonce() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeIncreaseNonce")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsafeIncreaseNonce indicates an expected call of UnsafeIncreaseNonce.
func (mr *MockChainClientMockRecorder) UnsafeIncreaseNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeIncreaseNonce", reflect.TypeOf((*MockChainClient)(nil).UnsafeIncreaseNonce))
}

// UnsafeNonce mocks base method.
func (m *MockChainClient) UnsafeNonce() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeNonce")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsafeNonce indicates an expected call of UnsafeNonce.
func (mr *MockChainClientMockRecorder) UnsafeNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeNonce", reflect.TypeOf((*MockChainClient)(nil).UnsafeNonce))
}

// WaitAndReturnTxReceipt mocks base method.
func (m *MockChainClient) WaitAndReturnTxReceipt(arg0 common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitAndReturnTxReceipt", arg0)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitAndReturnTxReceipt indicates an expected call of WaitAndReturnTxReceipt.
func (mr *MockChainClientMockRecorder) WaitAndReturnTxReceipt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAndReturnTxReceipt", reflect.TypeOf((*MockChainClient)(nil).WaitAndReturnTxReceipt), arg0)
}

// MockMessageHandler is a mock of MessageHandler interface.
type MockMessageHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMessageHandlerMockRecorder
}

// MockMessageHandlerMockRecorder is the mock recorder for MockMessageHandler.
type MockMessageHandlerMockRecorder struct {
	mock *MockMessageHandler
}

// NewMockMessageHandler creates a new mock instance.
func NewMockMessageHandler(ctrl *gomock.Controller) *MockMessageHandler {
	mock := &MockMessageHandler{ctrl: ctrl}
	mock.recorder = &MockMessageHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageHandler) EXPECT() *MockMessageHandlerMockRecorder {
	return m.recorder
}

// HandleMessage mocks base method.
func (m *MockMessageHandler) HandleMessage(arg0 *message.Message) (*proposal.Proposal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0)
	ret0, _ := ret[0].(*proposal.Proposal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockMessageHandlerMockRecorder) HandleMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockMessageHandler)(nil).HandleMessage), arg0)
}

// MockBridgeContract is a mock of BridgeContract interface.
type MockBridgeContract struct {
	ctrl     *gomock.Controller
	recorder *MockBridgeContractMockRecorder
}

// MockBridgeContractMockRecorder is the mock recorder for MockBridgeContract.
type MockBridgeContractMockRecorder struct {
	mock *MockBridgeContract
}

// NewMockBridgeContract creates a new mock instance.
func NewMockBridgeContract(ctrl *gomock.Controller) *MockBridgeContract {
	mock := &MockBridgeContract{ctrl: ctrl}
	mock.recorder = &MockBridgeContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBridgeContract) EXPECT() *MockBridgeContractMockRecorder {
	return m.recorder
}

// GetThreshold mocks base method.
func (m *MockBridgeContract) GetThreshold() (byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThreshold")
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThreshold indicates an expected call of GetThreshold.
func (mr *MockBridgeContractMockRecorder) GetThreshold() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThreshold", reflect.TypeOf((*MockBridgeContract)(nil).GetThreshold))
}

// IsProposalVotedBy mocks base method.
func (m *MockBridgeContract) IsProposalVotedBy(arg0 common.Address, arg1 *proposal.Proposal) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProposalVotedBy", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsProposalVotedBy indicates an expected call of IsProposalVotedBy.
func (mr *MockBridgeContractMockRecorder) IsProposalVotedBy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProposalVotedBy", reflect.TypeOf((*MockBridgeContract)(nil).IsProposalVotedBy), arg0, arg1)
}

// ProposalStatus mocks base method.
func (m *MockBridgeContract) ProposalStatus(arg0 *proposal.Proposal) (message.ProposalStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProposalStatus", arg0)
	ret0, _ := ret[0].(message.ProposalStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposalStatus indicates an expected call of ProposalStatus.
func (mr *MockBridgeContractMockRecorder) ProposalStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposalStatus", reflect.TypeOf((*MockBridgeContract)(nil).ProposalStatus), arg0)
}

// SimulateVoteProposal mocks base method.
func (m *MockBridgeContract) SimulateVoteProposal(arg0 *proposal.Proposal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulateVoteProposal", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SimulateVoteProposal indicates an expected call of SimulateVoteProposal.
func (mr *MockBridgeContractMockRecorder) SimulateVoteProposal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateVoteProposal", reflect.TypeOf((*MockBridgeContract)(nil).SimulateVoteProposal), arg0)
}

// VoteProposal mocks base method.
func (m *MockBridgeContract) VoteProposal(arg0 *proposal.Proposal, arg1 transactor.TransactOptions) (*common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VoteProposal", arg0, arg1)
	ret0, _ := ret[0].(*common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VoteProposal indicates an expected call of VoteProposal.
func (mr *MockBridgeContractMockRecorder) VoteProposal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VoteProposal", reflect.TypeOf((*MockBridgeContract)(nil).VoteProposal), arg0, arg1)
}
