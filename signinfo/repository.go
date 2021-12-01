package signinfo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lidofinance/terra-fcd-rest-client/v5/client"
	"github.com/lidofinance/terra-fcd-rest-client/v5/client/query"

	"github.com/sirupsen/logrus"
)

type Repository interface {
	Init(ctx context.Context, consAddr string) error
	GetMissedBlockCounter() float64
	GetTombstoned() bool
	GetAddress() string
}

func NewRepositoryCol5(apiClient *client.TerraRESTApis, logger *logrus.Logger) *RepositoryColumbus5 {
	return &RepositoryColumbus5{
		apiClient: apiClient,
		logger:    logger,
	}
}

type RepositoryColumbus5 struct {
	apiClient   *client.TerraRESTApis
	logger      *logrus.Logger
	signingInfo *query.SigningInfoOKBodyValSigningInfo
}

func (s *RepositoryColumbus5) Init(ctx context.Context, consAddr string) error {
	signingInfoResponse, err := s.apiClient.Query.SigningInfo(
		&query.SigningInfoParams{
			ConsAddress: consAddr,
			Context:     ctx,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to GetSlashingSigningInfos for validator's consaddr %s: %w", consAddr, err)
	}
	if err := signingInfoResponse.GetPayload().Validate(nil); err != nil {
		return fmt.Errorf("failed to validate SignInfo for validator %s: %w", consAddr, err)
	}
	s.signingInfo = signingInfoResponse.GetPayload().ValSigningInfo
	return nil
}

func (s *RepositoryColumbus5) GetMissedBlockCounter() float64 {
	if s.signingInfo != nil {
		// No blocks is sent as "", not as "0".
		if s.signingInfo.MissedBlocksCounter != "" {
			// If the current block is greater than minHeight and the validator's MissedBlocksCounter is
			// greater than maxMissed, they will be slashed. So numMissedBlocks > 0 does not mean that we
			// are already slashed, but is alarming. Note: Liveness slashes do NOT lead to a tombstoning.
			// https://docs.terra.money/dev/spec-slashing.html#begin-block
			numMissedBlocks, err := strconv.ParseInt(s.signingInfo.MissedBlocksCounter, 10, 64)
			if err != nil {
				s.logger.Errorf("failed to Parse `missed_blocks_counter:`: %s", err)
			} else {
				if numMissedBlocks > 0 {
					return float64(numMissedBlocks)
				}
			}
		}
	}
	return 0
}

func (s *RepositoryColumbus5) GetTombstoned() bool {
	if s.signingInfo != nil {
		return s.signingInfo.Tombstoned
	}
	return false
}

func (s *RepositoryColumbus5) GetAddress() string {
	if s.signingInfo != nil {
		return s.signingInfo.Address
	}
	return ""
}
