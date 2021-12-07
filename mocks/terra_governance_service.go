package mocks

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/governance"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/models"

	"github.com/go-openapi/runtime"
)

const (
	terraGovernanceServiceSamplesFolder = "terra_governance_service"

	TestGetProposalDepositID  = 150
	TestGetProposalVotingID   = 149
	TestGetProposalPassedID   = 4
	TestGetProposalRejectedID = 5

	TestGetProposalVotesID = 4

	proposalByIDDesopitSampleFile  = "v1_gov_proposals_proposal_deposit.json"
	proposalByIDVotingSampleFile   = "v1_gov_proposals_proposal_voting.json"
	proposalByIDPassedSampleFile   = "v1_gov_proposals_proposal_passed.json"
	proposalByIDRejectedSampleFile = "v1_gov_proposals_proposal_rejected.json"

	depositProposalsSampleFile = "v1_gov_proposals_deposit.json"
	votingProposalsSampleFile  = "v1_gov_proposals_voting.json"

	proposalVotesSamlpeFile = "v1_gov_proposal_votes.json"
)

type TerraGovernanceServiceMock struct{}

func (s *TerraGovernanceServiceMock) GetGovParametersDeposit(params *governance.GetGovParametersDepositParams, opts ...governance.ClientOption) (*governance.GetGovParametersDepositOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovParametersTallying(params *governance.GetGovParametersTallyingParams, opts ...governance.ClientOption) (*governance.GetGovParametersTallyingOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovParametersVoting(params *governance.GetGovParametersVotingParams, opts ...governance.ClientOption) (*governance.GetGovParametersVotingOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovProposals(params *governance.GetGovProposalsParams, opts ...governance.ClientOption) (*governance.GetGovProposalsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovProposalsProposalID(params *governance.GetGovProposalsProposalIDParams, opts ...governance.ClientOption) (*governance.GetGovProposalsProposalIDOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovProposalsProposalIDDeposits(params *governance.GetGovProposalsProposalIDDepositsParams, opts ...governance.ClientOption) (*governance.GetGovProposalsProposalIDDepositsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovProposalsProposalIDDepositsDepositor(params *governance.GetGovProposalsProposalIDDepositsDepositorParams, opts ...governance.ClientOption) (*governance.GetGovProposalsProposalIDDepositsDepositorOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovProposalsProposalIDProposer(params *governance.GetGovProposalsProposalIDProposerParams, opts ...governance.ClientOption) (*governance.GetGovProposalsProposalIDProposerOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovProposalsProposalIDTally(params *governance.GetGovProposalsProposalIDTallyParams, opts ...governance.ClientOption) (*governance.GetGovProposalsProposalIDTallyOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovProposalsProposalIDVotes(params *governance.GetGovProposalsProposalIDVotesParams, opts ...governance.ClientOption) (*governance.GetGovProposalsProposalIDVotesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetGovProposalsProposalIDVotesVoter(params *governance.GetGovProposalsProposalIDVotesVoterParams, opts ...governance.ClientOption) (*governance.GetGovProposalsProposalIDVotesVoterOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetV1GovProposals(params *governance.GetV1GovProposalsParams, opts ...governance.ClientOption) (*governance.GetV1GovProposalsOK, error) {
	var filename string
	switch *params.Status {
	case "Voting":
		filename = votingProposalsSampleFile
	case "Deposit":
		filename = depositProposalsSampleFile
	default:
		return s.defaultGetV1GovProposalsResponse(), nil
	}
	data, err := readSample(terraGovernanceServiceSamplesFolder, filename)
	if err != nil {
		return nil, err
	}
	resp := &governance.GetV1GovProposalsOK{
		Payload: &models.GetProposalListResult{},
	}
	if err := json.Unmarshal(data, resp.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to resp: %w", err)
	}
	return resp, nil
}

func (s *TerraGovernanceServiceMock) GetV1GovProposalsProposalID(params *governance.GetV1GovProposalsProposalIDParams, opts ...governance.ClientOption) (*governance.GetV1GovProposalsProposalIDOK, error) {
	var filename string
	switch params.ProposalID {
	case strconv.Itoa(TestGetProposalDepositID):
		filename = proposalByIDDesopitSampleFile
	case strconv.Itoa(TestGetProposalVotingID):
		filename = proposalByIDVotingSampleFile
	case strconv.Itoa(TestGetProposalPassedID):
		filename = proposalByIDPassedSampleFile
	case strconv.Itoa(TestGetProposalRejectedID):
		filename = proposalByIDRejectedSampleFile
	default:
		return nil, fmt.Errorf("Proposal not found")
	}
	data, err := readSample(terraGovernanceServiceSamplesFolder, filename)
	if err != nil {
		return nil, err
	}
	resp := &governance.GetV1GovProposalsProposalIDOK{
		Payload: &models.GetProposalResult{},
	}
	if err := json.Unmarshal(data, resp.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to resp: %w", err)
	}
	return resp, nil
}

func (s *TerraGovernanceServiceMock) GetV1GovProposalsProposalIDDeposits(params *governance.GetV1GovProposalsProposalIDDepositsParams, opts ...governance.ClientOption) (*governance.GetV1GovProposalsProposalIDDepositsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) GetV1GovProposalsProposalIDVotes(params *governance.GetV1GovProposalsProposalIDVotesParams, opts ...governance.ClientOption) (*governance.GetV1GovProposalsProposalIDVotesOK, error) {
	if params.ProposalID != strconv.Itoa(TestGetProposalVotesID) {
		return nil, fmt.Errorf("Proposal not found")
	}
	data, err := readSample(terraGovernanceServiceSamplesFolder, proposalVotesSamlpeFile)
	if err != nil {
		return nil, err
	}
	resp := &governance.GetV1GovProposalsProposalIDVotesOK{
		Payload: &models.GetProposalVotesResult{},
	}
	if err := json.Unmarshal(data, resp.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to resp: %w", err)
	}
	return resp, nil
}

func (s *TerraGovernanceServiceMock) PostGovProposals(params *governance.PostGovProposalsParams, opts ...governance.ClientOption) (*governance.PostGovProposalsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) PostGovProposalsParamChange(params *governance.PostGovProposalsParamChangeParams, opts ...governance.ClientOption) (*governance.PostGovProposalsParamChangeOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) PostGovProposalsProposalIDDeposits(params *governance.PostGovProposalsProposalIDDepositsParams, opts ...governance.ClientOption) (*governance.PostGovProposalsProposalIDDepositsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) PostGovProposalsProposalIDVotes(params *governance.PostGovProposalsProposalIDVotesParams, opts ...governance.ClientOption) (*governance.PostGovProposalsProposalIDVotesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraGovernanceServiceMock) SetTransport(transport runtime.ClientTransport) {}

// defaultGetV1GovProposalsResponse returns the GetV1GovProposals response for cases when no
// proposals with the given status have been found.
func (s *TerraGovernanceServiceMock) defaultGetV1GovProposalsResponse() *governance.GetV1GovProposalsOK {
	denomVal := "uluna"
	amountVal := "50000000"
	maxDepositPeriodVal := "1209600000000000"
	votingPeriodVal := "604800000000000"
	return &governance.GetV1GovProposalsOK{
		Payload: &models.GetProposalListResult{
			MinDeposit: []*models.GetProposalListResultMinDeposit{{
				Denom:  &denomVal,
				Amount: &amountVal,
			}},
			MaxDepositPeriod: &maxDepositPeriodVal,
			VotingPeriod:     &votingPeriodVal,
			Proposals:        []*models.GetProposalListResultProposals{},
		},
	}
}
