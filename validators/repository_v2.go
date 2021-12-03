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

func NewV2ValidatorsRepository(valRegistryContract string, apiClient *client.TerraRESTApis) *V2ValidatorsRepository {
	return &V2ValidatorsRepository{
		validatorsRegistryContract: valRegistryContract,
		apiClient:                  apiClient,
	}
}

type V2ValidatorsRepository struct {
	validatorsRegistryContract string
	apiClient                  *client.TerraRESTApis
}

func (r *V2ValidatorsRepository) GetValidatorsAddresses(ctx context.Context) ([]string, error) {
	reqRaw, err := json.Marshal(&ValidatorRegistryValidatorsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ValidatorRegistryValidatorsRequest request: %w", err)
	}

	p := wasm.GetWasmContractsContractAddressStoreParams{}
	p.SetContext(ctx)
	p.SetContractAddress(r.validatorsRegistryContract)
	p.SetQueryMsg(string(reqRaw))

	resp, err := r.apiClient.Wasm.GetWasmContractsContractAddressStore(&p)
	if err != nil {
		return nil, fmt.Errorf("failed to process ValidatorRegistryValidatorsRequest request: %w", err)
	}

	if err := resp.GetPayload().Validate(nil); err != nil {
		return nil, fmt.Errorf("failed to validate ValidatorsWhitelist: %w", err)
	}

	valResp := make(ValidatorRegistryValidatorsResponse, 0)
	err = utils.CastMapToStruct(resp.Payload.Result, &valResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ValidatorRegistryValidatorsResponse body interface: %w", err)
	}

	valAddresses := make([]string, len(valResp))
	for i, val := range valResp {
		valAddresses[i] = val.Address
	}
	return valAddresses, nil
}

func (r *V2ValidatorsRepository) GetValidatorInfo(ctx context.Context, address string) (ValidatorInfo, error) {
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
