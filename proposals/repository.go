package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lidofinance/terra-fcd-rest-client/v5/client"
	"github.com/lidofinance/terra-fcd-rest-client/v5/client/governance"
	"github.com/lidofinance/terra-fcd-rest-client/v5/models"
)

var (
	VotingStatus  = "Voting"
	DepositStatus = "Deposit"
	votePageSize  = 100.0
)

type Repository interface {
	Get(ctx context.Context, proposalID int) (*models.GetProposalResult, error)
	FetchVoting(ctx context.Context) ([]*models.GetProposalListResultProposals, error)
	FetchDeposit(ctx context.Context) ([]*models.GetProposalListResultProposals, error)
	GetVotes(ctx context.Context, proposalID int) ([]*models.GetProposalVotesResultVotes, error)
}

func NewFCDRepo(apiClient *client.TerraRESTApis) *FCDRepo {
	return &FCDRepo{
		apiClient: apiClient,
	}
}

type FCDRepo struct {
	apiClient *client.TerraRESTApis
}

func (r *FCDRepo) Get(ctx context.Context, proposalID int) (*models.GetProposalResult, error) {
	resp, err := r.apiClient.Governance.GetV1GovProposalsProposalID(
		&governance.GetV1GovProposalsProposalIDParams{
			ProposalID: strconv.Itoa(proposalID),
			Context:    ctx,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get proposal id = %d: %w", proposalID, err)
	}

	err = resp.GetPayload().Validate(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to validate proposal id = %d response: %w", proposalID, err)
	}

	return resp.GetPayload(), nil
}

// FetchVoting fetches proposals with voting status
func (r *FCDRepo) FetchVoting(ctx context.Context) ([]*models.GetProposalListResultProposals, error) {
	return r.fetch(ctx, VotingStatus)
}

// FetchDeposit fetches proposals with voting status
func (r *FCDRepo) FetchDeposit(ctx context.Context) ([]*models.GetProposalListResultProposals, error) {
	return r.fetch(ctx, DepositStatus)
}

func (r *FCDRepo) fetch(ctx context.Context, status string) ([]*models.GetProposalListResultProposals, error) {
	resp, err := r.apiClient.Governance.GetV1GovProposals(
		&governance.GetV1GovProposalsParams{
			Status:  &status,
			Context: ctx,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get proposals with {%s} status: %w", status, err)
	}

	err = resp.GetPayload().Validate(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to validate proposals with {%s} status response: %w", status, err)
	}

	return resp.GetPayload().Proposals, nil
}

func (r *FCDRepo) GetVotes(ctx context.Context, proposalID int) ([]*models.GetProposalVotesResultVotes, error) {
	page := 1.0
	votes := make([]*models.GetProposalVotesResultVotes, 0)
	for {
		resp, err := r.apiClient.Governance.GetV1GovProposalsProposalIDVotes(
			&governance.GetV1GovProposalsProposalIDVotesParams{
				ProposalID: strconv.Itoa(proposalID),
				Limit:      &votePageSize,
				Page:       &page,
				Context:    ctx,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get votes for proposal id = %d: %w", proposalID, err)
		}

		err = resp.GetPayload().Validate(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to validate votes response for proposal id = %d: %w", proposalID, err)
		}

		votes = append(votes, resp.GetPayload().Votes...)
		if *resp.GetPayload().TotalCnt <= *resp.GetPayload().Page*votePageSize {
			break
		}
		page++
	}
	return votes, nil
}
