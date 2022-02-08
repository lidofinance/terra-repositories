package mocks

import (
	"encoding/json"
	"fmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/query"

	"github.com/go-openapi/runtime"
)

const (
	terraQueryServiceSamplesFolder = "terra_query_service"

	TestDelegatorAddress           = "terra1rq5eq7szjsj5zsqxzz595gc0z8hqsx7y9gdz03"
	TestSoundValidatorAddress      = "terravalcons1qqn3zhwavg7unykfjn5ggwh2m4njrwdx7lfxzh"
	TestTombstonedValidatorAddress = "terravalcons1xadsk662jwx669mjkxgdsc9s7a4fj33uy6yr9d"

	soundValidatorSigningInfoSampleFile      = "validator_signing_info_sound.json"
	tombstonedValidatorSigningInfoSampleFile = "validator_signing_info_tombstoned.json"
	delegatorDelegationsSampleFile           = "delegator_delegations_resp.json"
	queryProposalsVotesPage1                 = "gov_proposal_votes_page1.json"
	queryProposalsVotesPage2                 = "gov_proposal_votes_page2.json"
)

type TerraQueryServiceMock struct{}

func (s *TerraQueryServiceMock) Account(params *query.AccountParams, opts ...query.ClientOption) (*query.AccountOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Accounts(params *query.AccountsParams, opts ...query.ClientOption) (*query.AccountsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Actives(params *query.ActivesParams, opts ...query.ClientOption) (*query.ActivesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AggregatePrevote(params *query.AggregatePrevoteParams, opts ...query.ClientOption) (*query.AggregatePrevoteOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AggregatePrevotes(params *query.AggregatePrevotesParams, opts ...query.ClientOption) (*query.AggregatePrevotesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AggregateVote(params *query.AggregateVoteParams, opts ...query.ClientOption) (*query.AggregateVoteOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AggregateVotes(params *query.AggregateVotesParams, opts ...query.ClientOption) (*query.AggregateVotesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AllBalances(params *query.AllBalancesParams, opts ...query.ClientOption) (*query.AllBalancesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AllEvidence(params *query.AllEvidenceParams, opts ...query.ClientOption) (*query.AllEvidenceOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Allowance(params *query.AllowanceParams, opts ...query.ClientOption) (*query.AllowanceOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Allowances(params *query.AllowancesParams, opts ...query.ClientOption) (*query.AllowancesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AnnualProvisions(params *query.AnnualProvisionsParams, opts ...query.ClientOption) (*query.AnnualProvisionsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AppliedPlan(params *query.AppliedPlanParams, opts ...query.ClientOption) (*query.AppliedPlanOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) AuthParams(params *query.AuthParamsParams, opts ...query.ClientOption) (*query.AuthParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Balance(params *query.BalanceParams, opts ...query.ClientOption) (*query.BalanceOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) BankParams(params *query.BankParamsParams, opts ...query.ClientOption) (*query.BankParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ByteCode(params *query.ByteCodeParams, opts ...query.ClientOption) (*query.ByteCodeOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Channel(params *query.ChannelParams, opts ...query.ClientOption) (*query.ChannelOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ChannelClientState(params *query.ChannelClientStateParams, opts ...query.ClientOption) (*query.ChannelClientStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ChannelConsensusState(params *query.ChannelConsensusStateParams, opts ...query.ClientOption) (*query.ChannelConsensusStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Channels(params *query.ChannelsParams, opts ...query.ClientOption) (*query.ChannelsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ClientConnections(params *query.ClientConnectionsParams, opts ...query.ClientOption) (*query.ClientConnectionsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ClientParams(params *query.ClientParamsParams, opts ...query.ClientOption) (*query.ClientParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ClientState(params *query.ClientStateParams, opts ...query.ClientOption) (*query.ClientStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ClientStates(params *query.ClientStatesParams, opts ...query.ClientOption) (*query.ClientStatesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ClientStatus(params *query.ClientStatusParams, opts ...query.ClientOption) (*query.ClientStatusOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) CodeInfo(params *query.CodeInfoParams, opts ...query.ClientOption) (*query.CodeInfoOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) CommunityPool(params *query.CommunityPoolParams, opts ...query.ClientOption) (*query.CommunityPoolOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Connection(params *query.ConnectionParams, opts ...query.ClientOption) (*query.ConnectionOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ConnectionChannels(params *query.ConnectionChannelsParams, opts ...query.ClientOption) (*query.ConnectionChannelsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ConnectionClientState(params *query.ConnectionClientStateParams, opts ...query.ClientOption) (*query.ConnectionClientStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ConnectionConsensusState(params *query.ConnectionConsensusStateParams, opts ...query.ClientOption) (*query.ConnectionConsensusStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Connections(params *query.ConnectionsParams, opts ...query.ClientOption) (*query.ConnectionsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ConsensusState(params *query.ConsensusStateParams, opts ...query.ClientOption) (*query.ConsensusStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ConsensusStates(params *query.ConsensusStatesParams, opts ...query.ClientOption) (*query.ConsensusStatesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ContractInfo(params *query.ContractInfoParams, opts ...query.ClientOption) (*query.ContractInfoOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ContractStore(params *query.ContractStoreParams, opts ...query.ClientOption) (*query.ContractStoreOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) CurrentPlan(params *query.CurrentPlanParams, opts ...query.ClientOption) (*query.CurrentPlanOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Delegation(params *query.DelegationParams, opts ...query.ClientOption) (*query.DelegationOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DelegationRewards(params *query.DelegationRewardsParams, opts ...query.ClientOption) (*query.DelegationRewardsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DelegationTotalRewards(params *query.DelegationTotalRewardsParams, opts ...query.ClientOption) (*query.DelegationTotalRewardsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DelegatorDelegations(params *query.DelegatorDelegationsParams, opts ...query.ClientOption) (*query.DelegatorDelegationsOK, error) {
	if params.DelegatorAddr != TestDelegatorAddress {
		return s.defaultDelegatorDelegationsResponse(), nil
	}

	data, err := readSample(terraQueryServiceSamplesFolder, delegatorDelegationsSampleFile)
	if err != nil {
		return nil, err
	}
	resp := &query.DelegatorDelegationsOK{
		Payload: &query.DelegatorDelegationsOKBody{},
	}
	if err := json.Unmarshal(data, resp.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to resp: %w", err)
	}
	return resp, nil
}

func (s *TerraQueryServiceMock) DelegatorUnbondingDelegations(params *query.DelegatorUnbondingDelegationsParams, opts ...query.ClientOption) (*query.DelegatorUnbondingDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DelegatorValidator(params *query.DelegatorValidatorParams, opts ...query.ClientOption) (*query.DelegatorValidatorOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DelegatorValidators(params *query.DelegatorValidatorsParams, opts ...query.ClientOption) (*query.DelegatorValidatorsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DelegatorWithdrawAddress(params *query.DelegatorWithdrawAddressParams, opts ...query.ClientOption) (*query.DelegatorWithdrawAddressOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DenomMetadata(params *query.DenomMetadataParams, opts ...query.ClientOption) (*query.DenomMetadataOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DenomTrace(params *query.DenomTraceParams, opts ...query.ClientOption) (*query.DenomTraceOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DenomTraces(params *query.DenomTracesParams, opts ...query.ClientOption) (*query.DenomTracesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DenomsMetadata(params *query.DenomsMetadataParams, opts ...query.ClientOption) (*query.DenomsMetadataOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Deposit(params *query.DepositParams, opts ...query.ClientOption) (*query.DepositOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Deposits(params *query.DepositsParams, opts ...query.ClientOption) (*query.DepositsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) DistributionParams(params *query.DistributionParamsParams, opts ...query.ClientOption) (*query.DistributionParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Evidence(params *query.EvidenceParams, opts ...query.ClientOption) (*query.EvidenceOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ExchangeRate(params *query.ExchangeRateParams, opts ...query.ClientOption) (*query.ExchangeRateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ExchangeRates(params *query.ExchangeRatesParams, opts ...query.ClientOption) (*query.ExchangeRatesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) FeederDelegation(params *query.FeederDelegationParams, opts ...query.ClientOption) (*query.FeederDelegationOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) GovParams(params *query.GovParamsParams, opts ...query.ClientOption) (*query.GovParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Grants(params *query.GrantsParams, opts ...query.ClientOption) (*query.GrantsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) HistoricalInfo(params *query.HistoricalInfoParams, opts ...query.ClientOption) (*query.HistoricalInfoOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) IBCTransferParams(params *query.IBCTransferParamsParams, opts ...query.ClientOption) (*query.IBCTransferParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) IBCUpgradedConsensusState(params *query.IBCUpgradedConsensusStateParams, opts ...query.ClientOption) (*query.IBCUpgradedConsensusStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Indicators(params *query.IndicatorsParams, opts ...query.ClientOption) (*query.IndicatorsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Inflation(params *query.InflationParams, opts ...query.ClientOption) (*query.InflationOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) MarketParams(params *query.MarketParamsParams, opts ...query.ClientOption) (*query.MarketParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) MintParams(params *query.MintParamsParams, opts ...query.ClientOption) (*query.MintParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) MissCounter(params *query.MissCounterParams, opts ...query.ClientOption) (*query.MissCounterOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ModuleVersions(params *query.ModuleVersionsParams, opts ...query.ClientOption) (*query.ModuleVersionsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) NextSequenceReceive(params *query.NextSequenceReceiveParams, opts ...query.ClientOption) (*query.NextSequenceReceiveOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) OracleParams(params *query.OracleParamsParams, opts ...query.ClientOption) (*query.OracleParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) PacketAcknowledgement(params *query.PacketAcknowledgementParams, opts ...query.ClientOption) (*query.PacketAcknowledgementOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) PacketAcknowledgements(params *query.PacketAcknowledgementsParams, opts ...query.ClientOption) (*query.PacketAcknowledgementsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) PacketCommitment(params *query.PacketCommitmentParams, opts ...query.ClientOption) (*query.PacketCommitmentOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) PacketCommitments(params *query.PacketCommitmentsParams, opts ...query.ClientOption) (*query.PacketCommitmentsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) PacketReceipt(params *query.PacketReceiptParams, opts ...query.ClientOption) (*query.PacketReceiptOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Params(params *query.ParamsParams, opts ...query.ClientOption) (*query.ParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Pool(params *query.PoolParams, opts ...query.ClientOption) (*query.PoolOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Proposal(params *query.ProposalParams, opts ...query.ClientOption) (*query.ProposalOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Proposals(params *query.ProposalsParams, opts ...query.ClientOption) (*query.ProposalsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) RawStore(params *query.RawStoreParams, opts ...query.ClientOption) (*query.RawStoreOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Redelegations(params *query.RedelegationsParams, opts ...query.ClientOption) (*query.RedelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) RewardWeight(params *query.RewardWeightParams, opts ...query.ClientOption) (*query.RewardWeightOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) SeigniorageProceeds(params *query.SeigniorageProceedsParams, opts ...query.ClientOption) (*query.SeigniorageProceedsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) SigningInfo(params *query.SigningInfoParams, opts ...query.ClientOption) (*query.SigningInfoOK, error) {
	var filename string
	switch params.ConsAddress {
	case TestSoundValidatorAddress:
		filename = soundValidatorSigningInfoSampleFile
	case TestTombstonedValidatorAddress:
		filename = tombstonedValidatorSigningInfoSampleFile
	default:
		return nil, fmt.Errorf("Validator not found")
	}
	data, err := readSample(terraQueryServiceSamplesFolder, filename)
	if err != nil {
		return nil, err
	}
	resp := &query.SigningInfoOK{
		Payload: &query.SigningInfoOKBody{},
	}
	if err := json.Unmarshal(data, resp.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to resp: %w", err)
	}
	return resp, nil
}

func (s *TerraQueryServiceMock) SigningInfos(params *query.SigningInfosParams, opts ...query.ClientOption) (*query.SigningInfosOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) SlashingParams(params *query.SlashingParamsParams, opts ...query.ClientOption) (*query.SlashingParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) StakingDelegatorValidators(params *query.StakingDelegatorValidatorsParams, opts ...query.ClientOption) (*query.StakingDelegatorValidatorsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) StakingParams(params *query.StakingParamsParams, opts ...query.ClientOption) (*query.StakingParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) SupplyOf(params *query.SupplyOfParams, opts ...query.ClientOption) (*query.SupplyOfOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Swap(params *query.SwapParams, opts ...query.ClientOption) (*query.SwapOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TallyResult(params *query.TallyResultParams, opts ...query.ClientOption) (*query.TallyResultOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TaxCap(params *query.TaxCapParams, opts ...query.ClientOption) (*query.TaxCapOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TaxCaps(params *query.TaxCapsParams, opts ...query.ClientOption) (*query.TaxCapsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TaxProceeds(params *query.TaxProceedsParams, opts ...query.ClientOption) (*query.TaxProceedsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TaxRate(params *query.TaxRateParams, opts ...query.ClientOption) (*query.TaxRateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TerraPoolDelta(params *query.TerraPoolDeltaParams, opts ...query.ClientOption) (*query.TerraPoolDeltaOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TobinTax(params *query.TobinTaxParams, opts ...query.ClientOption) (*query.TobinTaxOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TobinTaxes(params *query.TobinTaxesParams, opts ...query.ClientOption) (*query.TobinTaxesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TotalSupply(params *query.TotalSupplyParams, opts ...query.ClientOption) (*query.TotalSupplyOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) TreasuryParams(params *query.TreasuryParamsParams, opts ...query.ClientOption) (*query.TreasuryParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) UnbondingDelegation(params *query.UnbondingDelegationParams, opts ...query.ClientOption) (*query.UnbondingDelegationOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) UnreceivedAcks(params *query.UnreceivedAcksParams, opts ...query.ClientOption) (*query.UnreceivedAcksOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) UnreceivedPackets(params *query.UnreceivedPacketsParams, opts ...query.ClientOption) (*query.UnreceivedPacketsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) UpgradedClientState(params *query.UpgradedClientStateParams, opts ...query.ClientOption) (*query.UpgradedClientStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) UpgradedConsensusState(params *query.UpgradedConsensusStateParams, opts ...query.ClientOption) (*query.UpgradedConsensusStateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Validator(params *query.ValidatorParams, opts ...query.ClientOption) (*query.ValidatorOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ValidatorCommission(params *query.ValidatorCommissionParams, opts ...query.ClientOption) (*query.ValidatorCommissionOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ValidatorDelegations(params *query.ValidatorDelegationsParams, opts ...query.ClientOption) (*query.ValidatorDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ValidatorOutstandingRewards(params *query.ValidatorOutstandingRewardsParams, opts ...query.ClientOption) (*query.ValidatorOutstandingRewardsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ValidatorSlashes(params *query.ValidatorSlashesParams, opts ...query.ClientOption) (*query.ValidatorSlashesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) ValidatorUnbondingDelegations(params *query.ValidatorUnbondingDelegationsParams, opts ...query.ClientOption) (*query.ValidatorUnbondingDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Validators(params *query.ValidatorsParams, opts ...query.ClientOption) (*query.ValidatorsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Vote(params *query.VoteParams, opts ...query.ClientOption) (*query.VoteOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) VoteTargets(params *query.VoteTargetsParams, opts ...query.ClientOption) (*query.VoteTargetsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) Votes(params *query.VotesParams, opts ...query.ClientOption) (*query.VotesOK, error) {
	filename := queryProposalsVotesPage1
	if params.PaginationKey.String() == "FCQ+djfCTlED/udx+oKXk7H1CrtY" {
		filename = queryProposalsVotesPage2
	}
	data, err := readSample(terraQueryServiceSamplesFolder, filename)
	if err != nil {
		return nil, err
	}
	resp := &query.VotesOK{
		Payload: &query.VotesOKBody{},
	}
	if err := json.Unmarshal(data, resp.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to resp: %w", err)
	}
	return resp, nil
}

func (s *TerraQueryServiceMock) WasmParams(params *query.WasmParamsParams, opts ...query.ClientOption) (*query.WasmParamsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraQueryServiceMock) SetTransport(transport runtime.ClientTransport) {}

// defaultDelegatorDelegationsResponse returns the DelegatorDelegations response for cases when
// the requested address has no delegations.
func (s *TerraQueryServiceMock) defaultDelegatorDelegationsResponse() *query.DelegatorDelegationsOK {
	return &query.DelegatorDelegationsOK{
		Payload: &query.DelegatorDelegationsOKBody{
			DelegationResponses: []*query.DelegatorDelegationsOKBodyDelegationResponsesItems0{},
			Pagination: &query.DelegatorDelegationsOKBodyPagination{
				NextKey: nil,
				Total:   "0",
			},
		},
	}
}
