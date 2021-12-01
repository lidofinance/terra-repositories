package validators

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
)

const (
	Bech32TerraValConsPrefix = "terravalcons"
)

type ValidatorsRepository interface {
	GetValidatorsAddresses(ctx context.Context) ([]string, error)
	GetValidatorInfo(ctx context.Context, address string) (ValidatorInfo, error)
}

func GetPubKeyIdentifier(pubkey interface{}) (string, error) {
	pk, ok := pubkey.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to cast pubkey interface to map[string]string: %+v", pubkey)
	}
	pubkeyValue, ok := pk["value"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get pubkey's value from data struct: %+v", pk)
	}
	key, err := base64.StdEncoding.DecodeString(pubkeyValue)
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

func GetValConsAddr(valcons string) (string, error) {
	valconsAddr, err := ValConsToAddr(valcons)
	if err != nil {
		return "", fmt.Errorf("failed to convert valcons(%s) to valconsaddr: %w", valcons, err)
	}
	return valconsAddr, nil
}

func ValConsToAddr(valcons string) (string, error) {
	_, addr, err := bech32.DecodeAndConvert(valcons)
	return strings.ToUpper(hex.EncodeToString(addr)), err
}

func ValConsPubToAddr(valconspub string) (string, error) {
	_, data, err := bech32.DecodeAndConvert(valconspub)
	if err != nil {
		return "", fmt.Errorf("failed to decode terravalconspub: %w", err)
	}
	cdc := amino.NewCodec()
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKey{},
		ed25519.PubKeyName, nil)
	var pubKey crypto.PubKey
	err = cdc.UnmarshalBinaryBare(data, &pubKey)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal binary data to ed25519 pubkey: %w", err)
	}
	return pubKey.Address().String(), nil
}
