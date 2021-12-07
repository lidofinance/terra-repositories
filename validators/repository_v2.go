package validators

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lidofinance/terra-repositories/utils"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/wasm"
)

func NewV2Repository(valRegistryContract string, apiClient *client.TerraRESTApis) *V2Repository {
	return &V2Repository{
		baseRepository: baseRepository{
			apiClient: apiClient,
		},
		validatorsRegistryContract: valRegistryContract,
	}
}

type V2Repository struct {
	baseRepository
	validatorsRegistryContract string
}

func (r *V2Repository) GetValidatorsAddresses(ctx context.Context) ([]string, error) {
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
