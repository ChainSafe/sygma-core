// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package voter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls"
	"github.com/ChainSafe/chainbridge-core/chains/evm/voter/proposal"
	"github.com/ChainSafe/chainbridge-core/relayer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

type ChainClient interface {
	LatestBlock() (*big.Int, error)
	RelayerAddress() common.Address
	CallContract(ctx context.Context, callArgs map[string]interface{}, blockNumber *big.Int) ([]byte, error)
	ChainID(ctx context.Context) (*big.Int, error)
	calls.ClientDispatcher
}

type MessageHandler interface {
	HandleMessage(m *relayer.Message) (*proposal.Proposal, error)
}

type EVMVoter struct {
	mh             MessageHandler
	client         ChainClient
	fabric         calls.TxFabric
	gasPriceClient calls.GasPricer
}

func NewVoter(mh MessageHandler, client ChainClient, fabric calls.TxFabric, gasPriceClient calls.GasPricer) *EVMVoter {
	return &EVMVoter{
		mh:             mh,
		client:         client,
		fabric:         fabric,
		gasPriceClient: gasPriceClient,
	}
}

func (v *EVMVoter) VoteProposal(m *relayer.Message) error {
	prop, err := v.mh.HandleMessage(m)
	if err != nil {
		return err
	}
	ps, err := calls.ProposalStatus(v.client, prop)
	if err != nil {
		return fmt.Errorf("error getting proposal: %+v status %w", prop, err)
	}
	votedByTheRelayer, err := calls.IsProposalVotedBy(v.client, v.client.RelayerAddress(), prop)
	if err != nil {
		return err
	}
	// if this relayer had not voted for proposal and proposal is in Active or Inactive status
	// we need to vote for it
	if !votedByTheRelayer && (ps == relayer.ProposalStatusActive || ps == relayer.ProposalStatusInactive) {
		results := make(chan *proposal.Proposal)
		errors := make(chan error)
		go calls.SimulateVoteProposal(v.client, prop, results, errors, 5)

		select {
		case err := <-errors:
			return fmt.Errorf("voting failed. Err: %w", err)
		case res := <-results:
			hash, err := calls.VoteProposal(v.client, v.fabric, v.gasPriceClient, res)
			if err != nil {
				return fmt.Errorf("voting failed. Err: %w", err)
			}
			log.Debug().Str("hash", hash.String()).Uint64("nonce", res.DepositNonce).Msgf("Voted")
			return nil
		}
	}
	return nil
}
