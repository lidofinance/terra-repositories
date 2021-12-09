package validators

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/staking"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

const (
	Bech32TerraValConsPrefix = "terravalcons"
)

type baseRepository struct {
	apiClient *client.TerraRESTApis
}

func (r *baseRepository) GetValidatorInfo(ctx context.Context, address string) (ValidatorInfo, error) {
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

	consPubKeyAddress, err := getPubKeyIdentifier(validatorInfoResponse.GetPayload().Result.ConsensusPubkey)
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

func getPubKeyIdentifier(conspubkey *staking.GetStakingValidatorsValidatorAddrOKBodyResultConsensusPubkey) (string, error) {
	key, err := base64.StdEncoding.DecodeString(conspubkey.Value)
	if err != nil {
		return "", fmt.Errorf("failed to decode validator's ConsensusPubkey: %w", err)
	}

	pub := &ed25519.PubKey{Key: key}

	consPubKeyAddress, err := bech32.ConvertAndEncode(Bech32TerraValConsPrefix, pub.Address())
	if err != nil {
		return "", fmt.Errorf("failed to convert validator's ConsensusPubkeyAddress to bech32: %w", err)
	}
	return consPubKeyAddress, nil
}
