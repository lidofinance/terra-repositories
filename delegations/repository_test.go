package delegations

import (
	"context"
	"testing"

	"github.com/lidofinance/terra-repositories/mocks"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/query"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/factory"

	"github.com/stretchr/testify/assert"
)

const (
	fcdTestDelegationsValidatedThreshold = 500
)

var (
	testValidatorsPaginationLimit = "500"
)

func TestGetDelegationsFromAddress(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		repo := New(&client.TerraRESTApis{Query: &mocks.TerraQueryServiceMock{}})
		delegations, err := repo.GetDelegationsFromAddress(context.Background(), mocks.TestDelegatorAddress)
		if assert.Nil(t, err) {
			assert.Equalf(t, 3, len(delegations), "delegations number mismatch")
			t.Logf("validating %d delegations from %s", len(delegations), mocks.TestDelegatorAddress)
			validateDelegations(t, delegations)
		}
	})

	t.Run("WithFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		terraClient := factory.NewDefaultClient()
		validators, err := terraClient.Query.Validators(&query.ValidatorsParams{
			Context:         context.Background(),
			PaginationLimit: &testValidatorsPaginationLimit,
		})
		if err != nil {
			t.Fatalf("failed to get network validators: %v", err)
		}
		repo := New(terraClient)
		var validatedDelegations int
		for _, validator := range validators.Payload.Validators {
			delegations, err := terraClient.Query.ValidatorDelegations(&query.ValidatorDelegationsParams{
				Context:       context.Background(),
				ValidatorAddr: validator.OperatorAddress,
			})
			if err != nil {
				t.Logf("failed to get validator %s delegations: %v", validator.OperatorAddress, err)
				continue
			}
			for _, delegationResp := range delegations.Payload.DelegationResponses {
				delegatorAddr := delegationResp.Delegation.DelegatorAddress
				delegations, err := repo.GetDelegationsFromAddress(context.Background(), delegatorAddr)
				if assert.Nilf(t, err, "failed to get delegations of %s", delegatorAddr) {
					t.Logf("validating %d delegations from %s", len(delegations), delegatorAddr)
					validateDelegations(t, delegations)
					validatedDelegations += len(delegations)
				}
				if validatedDelegations > fcdTestDelegationsValidatedThreshold {
					t.Logf("%d delegations have been validating in total", validatedDelegations)
					return
				}
			}
		}
		t.Fatalf("couldn't validate at least %d delegations", fcdTestDelegationsValidatedThreshold)
	})
}

func validateDelegations(t *testing.T, delegations []Delegation) {
	for _, delegation := range delegations {
		assert.NotEmptyf(t, delegation.DelegatorAddress, "delegation address should not be empty")
		assert.NotEmptyf(t, delegation.ValidatorAddress, "validator address should not be empty")
		assert.NotEmptyf(t, delegation.Shares, "shares should not be empty")
	}
}
