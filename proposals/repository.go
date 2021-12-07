package proposals

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/governance"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/models"
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

func New(apiClient *client.TerraRESTApis) *BaseRepository {
	return &BaseRepository{apiClient: apiClient}
}

type BaseRepository struct {
	apiClient *client.TerraRESTApis
}

func (r *BaseRepository) Get(ctx context.Context, proposalID int) (*models.GetProposalResult, error) {
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
func (r *BaseRepository) FetchVoting(ctx context.Context) ([]*models.GetProposalListResultProposals, error) {
	return r.fetch(ctx, VotingStatus)
}

// FetchDeposit fetches proposals with deposit status
func (r *BaseRepository) FetchDeposit(ctx context.Context) ([]*models.GetProposalListResultProposals, error) {
	return r.fetch(ctx, DepositStatus)
}

func (r *BaseRepository) fetch(ctx context.Context, status string) ([]*models.GetProposalListResultProposals, error) {
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

func (r *BaseRepository) GetVotes(ctx context.Context, proposalID int) ([]*models.GetProposalVotesResultVotes, error) {
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
