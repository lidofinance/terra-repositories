package validators

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lidofinance/terra-repositories/utils"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/staking"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/wasm"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

func NewV1ValidatorsRepository(hubContract string, apiClient *client.TerraRESTApis) *V1ValidatorsRepository {
	return &V1ValidatorsRepository{
		hubContract: hubContract,
		apiClient:   apiClient,
	}
}

type V1ValidatorsRepository struct {
	hubContract string
	apiClient   *client.TerraRESTApis
}

func (r *V1ValidatorsRepository) GetValidatorsAddresses(ctx context.Context) ([]string, error) {
	reqRaw, err := json.Marshal(&HubWhitelistedValidatorsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal HubWhitelistedValidators request: %w", err)
	}

	p := wasm.GetWasmContractsContractAddressStoreParams{}
	p.SetContext(ctx)
	p.SetContractAddress(r.hubContract)
	p.SetQueryMsg(string(reqRaw))

	resp, err := r.apiClient.Wasm.GetWasmContractsContractAddressStore(&p)
	if err != nil {
		return nil, fmt.Errorf("failed to process HubWhitelistedValidators request: %w", err)
	}

	if err := resp.GetPayload().Validate(nil); err != nil {
		return nil, fmt.Errorf("failed to validate ValidatorsWhitelist: %w", err)
	}

	hubResp := &HubWhitelistedValidatorsResponse{}
	err = utils.CastMapToStruct(resp.Payload.Result, hubResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HubWhitelistedValidators body interface: %w", err)
	}

	return hubResp.Validators, nil
}

func (r *V1ValidatorsRepository) GetValidatorInfo(ctx context.Context, address string) (ValidatorInfo, error) {
	validatorInfoResponse, err := r.apiClient.Staking.GetStakingValidatorsValidatorAddr(
		&staking.GetStakingValidatorsValidatorAddrParams{
			ValidatorAddr: address,
			Context:       ctx,
		},
	)
	if err != nil {
		return ValidatorInfo{}, fmt.Errorf("failed to GetStakingValidatorsValidatorAddr: %w", err)
	}

	if err := validatorInfoResponse.GetPayload().Validate(nil); err != nil {
		return ValidatorInfo{}, fmt.Errorf("failed to validate ValidatorInfo for validator %s: %w", address, err)
	}

	commissionRate, err := cosmostypes.NewDecFromStr(validatorInfoResponse.GetPayload().Result.Commission.CommissionRates.Rate)
	if err != nil {
		return ValidatorInfo{}, fmt.Errorf("failed to parse validator's comission rate: %w", err)
	}

	commissionRateValue, err := commissionRate.Float64()
	if err != nil {
		return ValidatorInfo{}, fmt.Errorf("failed to parse float validator's comission rate: %w", err)
	}

	consPubKeyAddress, err := GetPubKeyIdentifier(validatorInfoResponse.GetPayload().Result.ConsensusPubkey)
	if err != nil {
		return ValidatorInfo{}, fmt.Errorf("failed to extract identifier from payload: %w", err)
	}

	return ValidatorInfo{
		Address:        address,
		Moniker:        validatorInfoResponse.GetPayload().Result.Description.Moniker,
		PubKey:         consPubKeyAddress,
		CommissionRate: commissionRateValue,
		Jailed:         validatorInfoResponse.GetPayload().Result.Jailed,
	}, nil
}
