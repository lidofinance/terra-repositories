package mocks

import (
	"encoding/json"
	"fmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/staking"

	"github.com/go-openapi/runtime"
)

const (
	terraStakingServiceSamplesFolder = "terra_staking_service"

	stakingFundValidatorAddr  = "terravaloper123gn6j23lmexu0qx5qhmgxgunmjcqsx8gmsyse"
	everstakeoneValidatorAddr = "terravaloper13g7z3qq6f00qww3u4mpcs3xw5jhqwraswraapc"
	p2pValidatorAddr          = "terravaloper144l7c3uph5a7h62xd8u5et3rqvj3dqtvvka2fu"
	blockscapeValidatorAddr   = "terravaloper1542ek7muegmm806akl0lam5vlqlph7spflfcun"

	stakingFundValidatorInfoSampleFile  = "staking_fund.json"
	everstakeoneValidatorInfoSampleFile = "everstakeone.json"
	p2pValidatorInfoSampleFile          = "p2p.json"
	blockscapeValidatorInfoSampleFile   = "blockscape.json"
)

type TerraStakingService struct{}

func (s *TerraStakingService) GetStakingDelegatorsDelegatorAddrDelegations(params *staking.GetStakingDelegatorsDelegatorAddrDelegationsParams, opts ...staking.ClientOption) (*staking.GetStakingDelegatorsDelegatorAddrDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddr(params *staking.GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrParams, opts ...staking.ClientOption) (*staking.GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingDelegatorsDelegatorAddrUnbondingDelegations(params *staking.GetStakingDelegatorsDelegatorAddrUnbondingDelegationsParams, opts ...staking.ClientOption) (*staking.GetStakingDelegatorsDelegatorAddrUnbondingDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddr(params *staking.GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrParams, opts ...staking.ClientOption) (*staking.GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingDelegatorsDelegatorAddrValidators(params *staking.GetStakingDelegatorsDelegatorAddrValidatorsParams, opts ...staking.ClientOption) (*staking.GetStakingDelegatorsDelegatorAddrValidatorsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddr(params *staking.GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams, opts ...staking.ClientOption) (*staking.GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingParameters(params *staking.GetStakingParametersParams, opts ...staking.ClientOption) (*staking.GetStakingParametersOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingPool(params *staking.GetStakingPoolParams, opts ...staking.ClientOption) (*staking.GetStakingPoolOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingRedelegations(params *staking.GetStakingRedelegationsParams, opts ...staking.ClientOption) (*staking.GetStakingRedelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingValidators(params *staking.GetStakingValidatorsParams, opts ...staking.ClientOption) (*staking.GetStakingValidatorsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingValidatorsValidatorAddr(params *staking.GetStakingValidatorsValidatorAddrParams, opts ...staking.ClientOption) (*staking.GetStakingValidatorsValidatorAddrOK, error) {
	var filename string
	switch params.ValidatorAddr {
	case stakingFundValidatorAddr:
		filename = stakingFundValidatorInfoSampleFile
	case everstakeoneValidatorAddr:
		filename = everstakeoneValidatorInfoSampleFile
	case p2pValidatorAddr:
		filename = p2pValidatorInfoSampleFile
	case blockscapeValidatorAddr:
		filename = blockscapeValidatorInfoSampleFile
	default:
		return nil, fmt.Errorf("Invalid validator address")
	}
	data, err := readSample(terraStakingServiceSamplesFolder, filename)
	if err != nil {
		return nil, err
	}
	resp := &staking.GetStakingValidatorsValidatorAddrOK{
		Payload: &staking.GetStakingValidatorsValidatorAddrOKBody{},
	}
	if err := json.Unmarshal(data, resp.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to resp: %w", err)
	}
	return resp, nil
}

func (s *TerraStakingService) GetStakingValidatorsValidatorAddrDelegations(params *staking.GetStakingValidatorsValidatorAddrDelegationsParams, opts ...staking.ClientOption) (*staking.GetStakingValidatorsValidatorAddrDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetStakingValidatorsValidatorAddrUnbondingDelegations(params *staking.GetStakingValidatorsValidatorAddrUnbondingDelegationsParams, opts ...staking.ClientOption) (*staking.GetStakingValidatorsValidatorAddrUnbondingDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1Staking(params *staking.GetV1StakingParams, opts ...staking.ClientOption) (*staking.GetV1StakingOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1StakingAccount(params *staking.GetV1StakingAccountParams, opts ...staking.ClientOption) (*staking.GetV1StakingAccountOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1StakingReturn(params *staking.GetV1StakingReturnParams, opts ...staking.ClientOption) (*staking.GetV1StakingReturnOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1StakingReturnOperatorAddr(params *staking.GetV1StakingReturnOperatorAddrParams, opts ...staking.ClientOption) (*staking.GetV1StakingReturnOperatorAddrOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1StakingValidators(params *staking.GetV1StakingValidatorsParams, opts ...staking.ClientOption) (*staking.GetV1StakingValidatorsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1StakingValidatorsOperatorAddr(params *staking.GetV1StakingValidatorsOperatorAddrParams, opts ...staking.ClientOption) (*staking.GetV1StakingValidatorsOperatorAddrOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1StakingValidatorsOperatorAddrClaims(params *staking.GetV1StakingValidatorsOperatorAddrClaimsParams, opts ...staking.ClientOption) (*staking.GetV1StakingValidatorsOperatorAddrClaimsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1StakingValidatorsOperatorAddrDelegations(params *staking.GetV1StakingValidatorsOperatorAddrDelegationsParams, opts ...staking.ClientOption) (*staking.GetV1StakingValidatorsOperatorAddrDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) GetV1StakingValidatorsOperatorAddrDelegators(params *staking.GetV1StakingValidatorsOperatorAddrDelegatorsParams, opts ...staking.ClientOption) (*staking.GetV1StakingValidatorsOperatorAddrDelegatorsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) PostStakingDelegatorsDelegatorAddrDelegations(params *staking.PostStakingDelegatorsDelegatorAddrDelegationsParams, opts ...staking.ClientOption) (*staking.PostStakingDelegatorsDelegatorAddrDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) PostStakingDelegatorsDelegatorAddrRedelegations(params *staking.PostStakingDelegatorsDelegatorAddrRedelegationsParams, opts ...staking.ClientOption) (*staking.PostStakingDelegatorsDelegatorAddrRedelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) PostStakingDelegatorsDelegatorAddrUnbondingDelegations(params *staking.PostStakingDelegatorsDelegatorAddrUnbondingDelegationsParams, opts ...staking.ClientOption) (*staking.PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraStakingService) SetTransport(transport runtime.ClientTransport) {}
