package proposals

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-openapi/strfmt"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/query"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/governance"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/models"
)

var (
	VotingStatus  = "Voting"
	DepositStatus = "Deposit"
	votePageSize  = "1000"
)

func New(apiClient *client.TerraRESTApis) *Repository {
	return &Repository{apiClient: apiClient}
}

type Repository struct {
	apiClient *client.TerraRESTApis
}

func (r *Repository) Get(ctx context.Context, proposalID int) (*models.GetProposalResult, error) {
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
func (r *Repository) FetchVoting(ctx context.Context) ([]*models.GetProposalListResultProposals, error) {
	return r.fetch(ctx, VotingStatus)
}

// FetchDeposit fetches proposals with deposit status
func (r *Repository) FetchDeposit(ctx context.Context) ([]*models.GetProposalListResultProposals, error) {
	return r.fetch(ctx, DepositStatus)
}

func (r *Repository) fetch(ctx context.Context, status string) ([]*models.GetProposalListResultProposals, error) {
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

func (r *Repository) GetVotingProposalVotes(ctx context.Context, proposalID int) ([]*query.VotesOKBodyVotesItems0, error) {
	var paginationKey strfmt.Base64
	votes := make([]*query.VotesOKBodyVotesItems0, 0)
	for {
		resp, err := r.apiClient.Query.Votes(
			&query.VotesParams{
				PaginationKey:   &paginationKey,
				PaginationLimit: &votePageSize,
				ProposalID:      strconv.Itoa(proposalID),
				Context:         ctx,
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
		if resp.GetPayload().Pagination.NextKey.String() == "" {
			break
		}
		paginationKey = resp.GetPayload().Pagination.NextKey
	}
	return votes, nil

}