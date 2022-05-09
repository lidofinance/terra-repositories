package validators

import (
	"context"
	"fmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/staking"
)

type FullSetValidatorsRepository struct {
	baseRepository
}

func NewFullSetValidatorsRepository(apiClient *client.TerraRESTApis) *FullSetValidatorsRepository {
	return &FullSetValidatorsRepository{
		baseRepository: baseRepository{
			apiClient: apiClient,
		},
	}
}

func (r *FullSetValidatorsRepository) GetValidatorsAddresses(ctx context.Context) ([]string, error) {
	params := staking.GetV1StakingValidatorsParams{Context: ctx}
	resp, err := r.apiClient.Staking.GetV1StakingValidators(&params)
	if err != nil {
		return nil, fmt.Errorf("failed to get full validators set page: %w", err)
	}

	var buf = []string{}
	for _, v := range resp.Payload {
		buf = append(buf, *v.OperatorAddress)
	}
	return buf, nil
}
