package validators

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lidofinance/terra-repositories/utils"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/wasm"
)

func NewV1Repository(hubContract string, apiClient *client.TerraRESTApis) *V1Repository {
	return &V1Repository{
		baseRepository: baseRepository{
			apiClient: apiClient,
		},
		hubContract: hubContract,
	}
}

type V1Repository struct {
	baseRepository
	hubContract string
}

func (r *V1Repository) GetValidatorsAddresses(ctx context.Context) ([]string, error) {
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
