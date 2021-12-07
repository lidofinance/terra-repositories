package proposals

import (
	"context"
	"testing"

	"github.com/lidofinance/terra-repositories/mocks"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/factory"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/models"

	"github.com/stretchr/testify/assert"
)

const (
	fcdTestDepositProposalID  = mocks.TestGetProposalDepositID
	fcdTestVotingProposalID   = mocks.TestGetProposalVotingID
	fcdTestPassedProposalID   = mocks.TestGetProposalPassedID
	fcdTestRejectedProposalID = mocks.TestGetProposalRejectedID
)

func TestGetProposal(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		repo := New(&client.TerraRESTApis{Governance: &mocks.TerraGovernanceServiceMock{}})
		for _, proposalID := range []int{
			mocks.TestGetProposalDepositID,
			mocks.TestGetProposalVotingID,
			mocks.TestGetProposalPassedID,
			mocks.TestGetProposalRejectedID,
		} {
			proposal, err := repo.Get(context.Background(), proposalID)
			if assert.Nil(t, err) {
				t.Logf("validating proposal %d", proposalID)
				validateGetProposalResult(t, proposal)
			}
		}
	})

	t.Run("WithFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		repo := New(factory.NewDefaultClient())
		for _, proposalID := range []int{
			fcdTestDepositProposalID,
			fcdTestVotingProposalID,
			fcdTestPassedProposalID,
			fcdTestRejectedProposalID,
		} {
			proposal, err := repo.Get(context.Background(), proposalID)
			if assert.Nil(t, err) {
				t.Logf("validating proposal %d", proposalID)
				validateGetProposalResult(t, proposal)
			}
		}
	})
}

func TestFetchVoting(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		repo := New(&client.TerraRESTApis{Governance: &mocks.TerraGovernanceServiceMock{}})
		votings, err := repo.FetchVoting(context.Background())
		if assert.Nil(t, err) {
			validateVotingProposals(t, votings)
		}
	})

	t.Run("WithFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		repo := New(factory.NewDefaultClient())
		votings, err := repo.FetchVoting(context.Background())
		if assert.Nil(t, err) {
			validateVotingProposals(t, votings)
		}
	})
}

func TestFetchDeposit(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		repo := New(&client.TerraRESTApis{Governance: &mocks.TerraGovernanceServiceMock{}})
		deposits, err := repo.FetchDeposit(context.Background())
		if assert.Nil(t, err) {
			validateDepositProposals(t, deposits)
		}
	})

	t.Run("WithFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		repo := New(factory.NewDefaultClient())
		deposits, err := repo.FetchDeposit(context.Background())
		if assert.Nil(t, err) {
			validateDepositProposals(t, deposits)
		}
	})
}

func TestGetVotes(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		repo := New(&client.TerraRESTApis{Governance: &mocks.TerraGovernanceServiceMock{}})
		votes, err := repo.GetVotes(context.Background(), mocks.TestGetProposalVotesID)
		if assert.Nil(t, err) {
			t.Logf("validating %d votes of proposal %d", len(votes), mocks.TestGetProposalVotesID)
			validateGetVotesResult(t, votes)
		}
	})

	t.Run("WithFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		repo := New(factory.NewDefaultClient())
		for _, proposalID := range []int{
			fcdTestVotingProposalID,
			fcdTestPassedProposalID,
			fcdTestRejectedProposalID,
		} {
			votes, err := repo.GetVotes(context.Background(), proposalID)
			if assert.Nil(t, err) {
				t.Logf("validating %d votes of proposal %d", len(votes), proposalID)
				validateGetVotesResult(t, votes)
			}
		}
	})
}

// validateGetProposalResult validates the proposal fields presence in accordance with the
// specification defined required fields list. Just the same as the validateGetProposalListResult,
// but has a different parameter type.
func validateGetProposalResult(t *testing.T, proposal *models.GetProposalResult) {
	assert.NotNilf(t, proposal.ID, "proposal ID must not be empty")
	if assert.NotNilf(t, proposal.Proposer, "proposer must not be empty") {
		assert.NotNilf(t, proposal.Proposer.AccountAddress, "proposer account address must not be empty")
	}
	assert.NotNilf(t, proposal.Type, "proposal type must not be empty")
	assert.NotNilf(t, proposal.Status, "proposal status must not be empty")
	assert.NotNilf(t, proposal.SubmitTime, "proposal submit time must not be empty")
	assert.NotNilf(t, proposal.Title, "proposal title must not be empty")
	assert.NotNilf(t, proposal.Description, "proposal description must not be empty")
	assert.NotNilf(t, proposal.Deposit, "proposal deposit must not be empty")
	if assert.NotNilf(t, proposal.Vote, "proposal vote must not be empty") {
		assert.NotNilf(t, proposal.Vote.ID, "vote ID must not be empty")
		if assert.NotNilf(t, proposal.Vote.Distribution, "proposal votes distribution must not be empty") {
			assert.NotNilf(t, proposal.Vote.Distribution.Yes, "proposal YES votes distribution must not be empty")
			assert.NotNilf(t, proposal.Vote.Distribution.No, "proposal NO votes distribution must not be empty")
			assert.NotNilf(t, proposal.Vote.Distribution.Abstain, "proposal ABSTAIN votes distribution must not be empty")
			assert.NotNilf(t, proposal.Vote.Distribution.NoWithVeto, "proposal NO WITH VETO votes distribution must not be empty")
		}
		if assert.NotNilf(t, proposal.Vote.Count, "proposal votes count must not be empty") {
			assert.NotNilf(t, proposal.Vote.Count.Yes, "proposal YES votes count must not be empty")
			assert.NotNilf(t, proposal.Vote.Count.No, "proposal NO votes count must not be empty")
			assert.NotNilf(t, proposal.Vote.Count.Abstain, "proposal ABSTAIN votes count must not be empty")
			assert.NotNilf(t, proposal.Vote.Count.NoWithVeto, "proposal NO WITH VETO votes count must not be empty")
		}
		assert.NotNilf(t, proposal.Vote.Total, "proposal total votes must not be empty")
		assert.NotNilf(t, proposal.Vote.VotingEndTime, "proposal voting end time must not be empty")
		assert.NotNilf(t, proposal.Vote.StakedLuna, "proposal staked luna must not be empty")
	}
}

func validateVotingProposals(t *testing.T, proposals []*models.GetProposalListResultProposals) {
	if assert.NotEmptyf(t, proposals, "voting proposals amount must be more than 0") {
		t.Logf("validating %d voting proposals", len(proposals))
		for _, proposal := range proposals {
			t.Logf("validating voting proposal %s", *proposal.ID)
			assert.Equalf(t, "Voting", *proposal.Status, "FetchVoting list result must be of Voting status")
			validateGetProposalListResult(t, proposal)
		}
	}
}

func validateDepositProposals(t *testing.T, proposals []*models.GetProposalListResultProposals) {
	if assert.NotEmptyf(t, proposals, "deposit proposals amount must be more than 0") {
		t.Logf("validating %d deposit proposals", len(proposals))
		for _, proposal := range proposals {
			t.Logf("validating deposit proposal %s", *proposal.ID)
			assert.Equalf(t, "Deposit", *proposal.Status, "FetchDeposit list result must be of Deposit status")
			validateGetProposalListResult(t, proposal)
		}
	}
}

// validateGetProposalListResult validates the proposal fields presence in accordance with the
// specification defined required fields list. Just the same as the validateGetProposalResult, but
// has a different parameter type.
func validateGetProposalListResult(t *testing.T, proposal *models.GetProposalListResultProposals) {
	assert.NotNilf(t, proposal.ID, "proposal ID must not be empty")
	if assert.NotNilf(t, proposal.Proposer, "proposer must not be empty") {
		assert.NotNilf(t, proposal.Proposer.AccountAddress, "proposer account address must not be empty")
	}
	assert.NotNilf(t, proposal.Type, "proposal type must not be empty")
	assert.NotNilf(t, proposal.Status, "proposal status must not be empty")
	assert.NotNilf(t, proposal.SubmitTime, "proposal submit time must not be empty")
	assert.NotNilf(t, proposal.Title, "proposal title must not be empty")
	assert.NotNilf(t, proposal.Description, "proposal description must not be empty")
	assert.NotNilf(t, proposal.Deposit, "proposal deposit must not be empty")
	if assert.NotNilf(t, proposal.Vote, "proposal vote must not be empty") {
		assert.NotNilf(t, proposal.Vote.ID, "vote ID must not be empty")
		if assert.NotNilf(t, proposal.Vote.Distribution, "proposal votes distribution must not be empty") {
			assert.NotNilf(t, proposal.Vote.Distribution.Yes, "proposal YES votes distribution must not be empty")
			assert.NotNilf(t, proposal.Vote.Distribution.No, "proposal NO votes distribution must not be empty")
			assert.NotNilf(t, proposal.Vote.Distribution.Abstain, "proposal ABSTAIN votes distribution must not be empty")
			assert.NotNilf(t, proposal.Vote.Distribution.NoWithVeto, "proposal NO WITH VETO votes distribution must not be empty")
		}
		if assert.NotNilf(t, proposal.Vote.Count, "proposal votes count must not be empty") {
			assert.NotNilf(t, proposal.Vote.Count.Yes, "proposal YES votes count must not be empty")
			assert.NotNilf(t, proposal.Vote.Count.No, "proposal NO votes count must not be empty")
			assert.NotNilf(t, proposal.Vote.Count.Abstain, "proposal ABSTAIN votes count must not be empty")
			assert.NotNilf(t, proposal.Vote.Count.NoWithVeto, "proposal NO WITH VETO votes count must not be empty")
		}
		assert.NotNilf(t, proposal.Vote.Total, "proposal total votes must not be empty")
		assert.NotNilf(t, proposal.Vote.VotingEndTime, "proposal voting end time must not be empty")
		assert.NotNilf(t, proposal.Vote.StakedLuna, "proposal staked luna must not be empty")
	}
}

// validateGetVotesResult validates each vote fields presence in accordance with the specification
// defined required fields list.
func validateGetVotesResult(t *testing.T, votes []*models.GetProposalVotesResultVotes) {
	if len(votes) == 0 {
		t.Error("votes amount must be more than 0")
		return
	}
	for _, vote := range votes {
		assert.NotNilf(t, vote.Answer, "vote answer must not be empty")
		if assert.NotNilf(t, vote.Voter, "voter must not be empty") {
			assert.NotNilf(t, vote.Voter.AccountAddress, "voter account address must not be empty")
			if vote.Voter.OperatorAddress != "" {
				assert.NotNilf(t, vote.Voter.Moniker, "operator voter moniker must not be empty")
			}
		}
	}
}
