package validators

import (
	"context"
	"fmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/staking"
)

type ActiveSetValidatorsRepository struct {
	baseRepository
}

func NewActiveSetValidatorsRepository(apiClient *client.TerraRESTApis) *ActiveSetValidatorsRepository {
	return &ActiveSetValidatorsRepository{
		baseRepository: baseRepository{
			apiClient: apiClient,
		},
	}
}

func (r *ActiveSetValidatorsRepository) GetValidatorsAddresses(ctx context.Context) ([]string, error) {
	var limit int64 = 130 // all validators
	var status string = "BOND_STATUS_BONDED"
	var buf = []string{}

	params := staking.GetStakingValidatorsParams{
		Limit:   &limit,
		Status:  &status,
		Context: ctx,
	}

	resp, err := r.apiClient.Staking.GetStakingValidators(&params)

	if err != nil {
		return nil, fmt.Errorf("failed to get active validators set page: %w", err)
	}

	for _, v := range resp.Payload.Result {
		buf = append(buf, v.OperatorAddress)
	}

	return buf, nil
}
