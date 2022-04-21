package proposals

import (
	"context"
	"testing"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/query"

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

var proposalStatusInputToOutput = map[string]string{
	"1": "PROPOSAL_STATUS_DEPOSIT_PERIOD",
	"2": "PROPOSAL_STATUS_VOTING_PERIOD",
	"3": "PROPOSAL_STATUS_PASSED",
	"4": "PROPOSAL_STATUS_REJECTED",
}

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
		repo := New(&client.TerraRESTApis{Governance: &mocks.TerraGovernanceServiceMock{}, Query: &mocks.TerraQueryServiceMock{}})
		votes, err := repo.FetchVoting(context.Background())
		if assert.Nil(t, err) {
			validateVotingProposals(t, votes)
		}
	})

	t.Run("WithFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		repo := New(factory.NewDefaultClient())
		_, err := repo.FetchVoting(context.Background())
		assert.Nil(t, err)
	})
}

func TestFetchDeposit(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		repo := New(&client.TerraRESTApis{Governance: &mocks.TerraGovernanceServiceMock{}, Query: &mocks.TerraQueryServiceMock{}})
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
		repo := New(&client.TerraRESTApis{Query: &mocks.TerraQueryServiceMock{}})
		votes, err := repo.GetVotingProposalVotes(context.Background(), mocks.TestGetProposalVotesID)
		if assert.Nil(t, err) {
			t.Logf("validating %d votes of proposal %d", len(votes), mocks.TestGetProposalVotesID)
			validateGetVotesResult(t, votes)
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

func validateVotingProposals(t *testing.T, proposals []*query.ProposalsOKBodyProposalsItems0) {
	if assert.NotEmptyf(t, proposals, "voting proposals amount must be more than 0") {
		t.Logf("validating %d voting proposals", len(proposals))
		for _, proposal := range proposals {
			t.Logf("validating voting proposal %s", proposal.ProposalID)
			assert.Equalf(t, proposalStatusInputToOutput[VotingStatus], proposal.Status, "FetchVoting list result must be of Voting status")
			validateGetProposalListResult(t, proposal)
		}
	}
}

func validateDepositProposals(t *testing.T, proposals []*query.ProposalsOKBodyProposalsItems0) {
	if assert.NotEmptyf(t, proposals, "deposit proposals amount must be more than 0") {
		t.Logf("validating %d deposit proposals", len(proposals))
		for _, proposal := range proposals {
			t.Logf("validating deposit proposal %s", proposal.ProposalID)
			assert.Equalf(t, proposalStatusInputToOutput[DepositStatus], proposal.Status, "FetchDeposit list result must be of Deposit status")
			validateGetProposalListResult(t, proposal)
		}
	}
}

// validateGetProposalListResult validates the proposal fields presence in accordance with the
// specification defined required fields list. Just the same as the validateGetProposalResult, but
// has a different parameter type.
func validateGetProposalListResult(t *testing.T, proposal *query.ProposalsOKBodyProposalsItems0) {
	assert.NotNilf(t, proposal.ProposalID, "proposal proposalID must not be empty")

	if assert.NotNilf(t, proposal.Content, "proposal content must not be empty") {
		assert.NotNilf(t, proposal.Content.AtType, "proposal atType must not be empty")
		assert.NotNilf(t, proposal.Content.Title, "proposal title must not be empty")
		assert.NotNilf(t, proposal.Content.Description, "proposal description must not be empty")
	}

	assert.NotNilf(t, proposal.Status, "proposal status must not be empty")

	if assert.NotNilf(t, proposal.FinalTallyResult, "proposal final tally result must not be empty") {
		assert.NotNilf(t, proposal.FinalTallyResult.Yes, "proposal final tally yes result must not be empty")
		assert.NotNilf(t, proposal.FinalTallyResult.No, "proposal final tally no result must not be empty")
		assert.NotNilf(t, proposal.FinalTallyResult.Abstain, "proposal final tally abstain result must not be empty")
		assert.NotNilf(t, proposal.FinalTallyResult.NoWithVeto, "proposal final tally noWithVeto result must not be empty")
	}

	assert.NotNilf(t, proposal.SubmitTime, "proposal submitTime must not be empty")
	assert.NotNilf(t, proposal.DepositEndTime, "proposal depositEndTime must not be empty")

	if assert.NotNilf(t, proposal.TotalDeposit, "proposal totalDeposit must not be empty") {
		assert.NotNilf(t, proposal.TotalDeposit[0].Denom, "proposal totalDeposit must not be empty")
		assert.NotNilf(t, proposal.TotalDeposit[0].Amount, "proposal totalDeposit amount must not be empty")
	}

	assert.NotNilf(t, proposal.VotingStartTime, "proposal votingStartTime must not be empty")
	assert.NotNilf(t, proposal.VotingEndTime, "proposal votingEndTime must not be empty")
}

// validateGetVotesResult validates each vote fields presence in accordance with the specification
// defined required fields list.
func validateGetVotesResult(t *testing.T, votes []*query.VotesOKBodyVotesItems0) {
	if len(votes) != 200 {
		t.Error("votes amount must be equal 200")
		return
	}
	for _, vote := range votes {
		assert.NotNilf(t, vote.ProposalID, "vote answer must not be empty")
		assert.NotNilf(t, vote.Voter, "voter must not be empty")
	}
}
