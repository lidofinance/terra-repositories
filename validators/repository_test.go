package validators

import (
	"context"
	"sync"
	"testing"

	"github.com/lidofinance/terra-repositories/mocks"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/factory"

	"github.com/stretchr/testify/assert"
)

const (
	hubContract                = "terra1mtwph2juhj0rvjz7dy92gvl6xvukaxu8rfv8ts"
	validatorsRegistryContract = "terra10wt548y4y3xeqfrqsgqlqh424lll8fqxp6dyed"

	bombayFCDHost   string = "bombay-fcd.terra.dev"
	bombayFCDScheme string = "https"
)

// bombayFCDEndpoint is the Full Client Daemon endpoint of the bombay network:
// https://bombay-fcd.terra.dev.
var bombayFCDEndpoint = factory.Endpoint{
	Host:    bombayFCDHost,
	Schemes: []string{bombayFCDScheme},
}

func TestV1Repository(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		repo := NewV1Repository(mocks.V1WhitelistedValidatorsContractAddress, &client.TerraRESTApis{
			Staking: &mocks.TerraStakingService{},
			Wasm:    &mocks.TerraWASMService{},
		})
		addresses, err := repo.GetValidatorsAddresses(context.Background())
		if err != nil {
			t.Fatalf("failed to get validator addresses: %v", err)
		}
		for _, validatorAddr := range addresses {
			valInfo, err := repo.GetValidatorInfo(context.Background(), validatorAddr)
			if assert.Nil(t, err) {
				t.Logf("validating %s validator info", validatorAddr)
				validateValidatorInfo(t, valInfo)
			}
		}
	})

	t.Run("WithFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		repo := NewV1Repository(hubContract, factory.NewDefaultClient())
		addresses, err := repo.GetValidatorsAddresses(context.Background())
		if err != nil {
			t.Fatalf("failed to get validator addresses: %v", err)
		}
		wg := &sync.WaitGroup{}
		for _, validatorAddr := range addresses {
			wg.Add(1)
			go func(addr string) {
				defer wg.Done()
				valInfo, err := repo.GetValidatorInfo(context.Background(), addr)
				if assert.Nil(t, err) {
					t.Logf("validating %s validator info", addr)
					validateValidatorInfo(t, valInfo)
				}
			}(validatorAddr)
		}
		wg.Wait()
	})
}

func TestV2Repository(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		repo := NewV2Repository(mocks.V2WhitelistedValidatorsContractAddress, &client.TerraRESTApis{
			Staking: &mocks.TerraStakingService{},
			Wasm:    &mocks.TerraWASMService{},
		})
		addresses, err := repo.GetValidatorsAddresses(context.Background())
		if err != nil {
			t.Fatalf("failed to get validator addresses: %v", err)
		}
		for _, validatorAddr := range addresses {
			valInfo, err := repo.GetValidatorInfo(context.Background(), validatorAddr)
			if assert.Nil(t, err) {
				t.Logf("validating %s validator info", validatorAddr)
				validateValidatorInfo(t, valInfo)
			}
		}
	})

	t.Run("WithBombayFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		repo := NewV2Repository(
			validatorsRegistryContract,
			factory.NewClient(bombayFCDEndpoint, client.DefaultBasePath),
		)
		addresses, err := repo.GetValidatorsAddresses(context.Background())
		if err != nil {
			t.Fatalf("failed to get validator addresses: %v", err)
		}
		wg := &sync.WaitGroup{}
		for _, validatorAddr := range addresses {
			wg.Add(1)
			go func(addr string) {
				defer wg.Done()
				valInfo, err := repo.GetValidatorInfo(context.Background(), addr)
				if assert.Nil(t, err) {
					t.Logf("validating %s validator info", addr)
					validateValidatorInfo(t, valInfo)
				}
			}(validatorAddr)
		}
		wg.Wait()
	})
}

func validateValidatorInfo(t *testing.T, valInfo ValidatorInfo) {
	assert.NotEmptyf(t, valInfo.Address, "validator address must not be empty")
	assert.NotEmptyf(t, valInfo.PubKey, "validator pub key must not be empty")
	assert.NotEmptyf(t, valInfo.Moniker, "validator moniker must not be empty")
}
