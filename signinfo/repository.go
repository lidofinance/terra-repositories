package signinfo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/query"

	"github.com/sirupsen/logrus"
)

type Repository interface {
	Init(ctx context.Context, consAddr string) error
	GetMissedBlockCounter() float64
	GetTombstoned() bool
	GetAddress() string
}

func New(apiClient *client.TerraRESTApis, logger *logrus.Logger) *BaseRepository {
	return &BaseRepository{
		apiClient: apiClient,
		logger:    logger,
	}
}

type BaseRepository struct {
	apiClient   *client.TerraRESTApis
	logger      *logrus.Logger
	signingInfo *query.SigningInfoOKBodyValSigningInfo
}

func (s *BaseRepository) Init(ctx context.Context, consAddr string) error {
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

func (s *BaseRepository) GetMissedBlockCounter() float64 {
	if s.signingInfo == nil || s.signingInfo.MissedBlocksCounter == "" { // no blocks is sent as "", not as "0".
		return 0
	}
	// If the current block is greater than minHeight and the validator's MissedBlocksCounter is
	// greater than maxMissed, they will be slashed. So numMissedBlocks > 0 does not mean that we
	// are already slashed, but is alarming. Note: Liveness slashes do NOT lead to a tombstoning.
	// https://docs.terra.money/dev/spec-slashing.html#begin-block
	numMissedBlocks, err := strconv.ParseFloat(s.signingInfo.MissedBlocksCounter, 64)
	if err != nil {
		s.logger.Errorf("failed to Parse `missed_blocks_counter:`: %s", err)
	}
	return numMissedBlocks
}

func (s *BaseRepository) GetTombstoned() bool {
	if s.signingInfo != nil {
		return s.signingInfo.Tombstoned
	}
	return false
}

func (s *BaseRepository) GetAddress() string {
	if s.signingInfo != nil {
		return s.signingInfo.Address
	}
	return ""
}
