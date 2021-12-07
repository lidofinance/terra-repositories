package signinfo

import (
	"context"
	"sync"
	"testing"

	"github.com/lidofinance/terra-repositories/mocks"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/query"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/factory"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	testSigningInfosPaginationLimit = "500"
)

func TestRepositoryPipeline(t *testing.T) {
	t.Run("WithMock", func(t *testing.T) {
		t.Run("Sound", func(t *testing.T) {
			repo := NewRepositoryCol5(&client.TerraRESTApis{Query: &mocks.TerraQueryServiceMock{}}, logrus.New())
			err := repo.Init(context.Background(), mocks.TestSoundValidatorAddress)
			if assert.Nil(t, err) {
				t.Logf("validating validator %s signing info", mocks.TestSoundValidatorAddress)
				assert.Falsef(t, repo.GetTombstoned(), "sound validator expected not to be tombstoned")
				assert.Equalf(t, mocks.TestSoundValidatorAddress, repo.GetAddress(), "validator address mismatch")
			}
		})

		t.Run("Tombstoned", func(t *testing.T) {
			repo := NewRepositoryCol5(&client.TerraRESTApis{Query: &mocks.TerraQueryServiceMock{}}, logrus.New())
			err := repo.Init(context.Background(), mocks.TestTombstonedValidatorAddress)
			if assert.Nil(t, err) {
				t.Logf("validating validator %s signing info", mocks.TestTombstonedValidatorAddress)
				assert.Truef(t, repo.GetTombstoned(), "tombstoned validator expected be tombstoned")
				assert.Equalf(t, mocks.TestTombstonedValidatorAddress, repo.GetAddress(), "validator address mismatch")
			}
		})
	})

	t.Run("WithFCD", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping test in short mode.")
		}
		terraClient := factory.NewDefaultClient()
		signingInfoParams := &query.SigningInfosParams{
			Context:         context.Background(),
			PaginationLimit: &testSigningInfosPaginationLimit,
		}
		signingInfo, err := terraClient.Query.SigningInfos(signingInfoParams)
		if assert.Nil(t, err) {
			if assert.NotEmptyf(t, signingInfo.Payload.Info, "signing infos amount must be more than 0") {
				wg := &sync.WaitGroup{}
				t.Logf("validating %d validators signing info", len(signingInfo.Payload.Info))
				for _, info := range signingInfo.Payload.Info {
					wg.Add(1)
					go func(addr string) {
						defer wg.Done()
						repo := NewRepositoryCol5(terraClient, logrus.New())
						assert.Nil(t, repo.Init(context.Background(), addr))
						assert.Equalf(t, addr, repo.GetAddress(), "validator address mismatch")
					}(info.Address)
				}
				wg.Wait()
			}
		}
	})
}
