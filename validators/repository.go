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

	NetworkGenerationColumbus4 = "columbus-4"
	NetworkGenerationColumbus5 = "columbus-5"
)

type ValidatorsRepository interface {
	GetValidatorsAddresses(ctx context.Context) ([]string, error)
	GetValidatorInfo(ctx context.Context, address string) (ValidatorInfo, error)
}

func GetPubKeyIdentifier(networkGeneration string, pubkey interface{}) (string, error) {
	switch networkGeneration {
	case NetworkGenerationColumbus4:
		// columbus4 ConsensusPubkey is just a string
		pk, ok := pubkey.(string)
		if !ok {
			return "", fmt.Errorf("failed to cast pubkey interface to string: %+v", pubkey)
		}
		return pk, nil
	case NetworkGenerationColumbus5:
		// columbus5 ConsensusPubkey is a struct
		// "consensus_pubkey": {
		//      "type": "tendermint/PubKeyEd25519",
		//      "value": "EAI7kGuMo6BG1poseFcoMiSa4vHmXcYM4VCpFeIMncw="
		//    }
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

	default:
		panic("unknown network generation. available variants: columbus-4 or columbus-5")
	}
}

// GetValConsAddr - get valconsaddr from pubkeyidentifier
// pubkeyidentifier is either valconspub for columbus4 or valcons for columbus5
func GetValConsAddr(networkGeneration string, pubkeyidentifier string) (string, error) {
	switch networkGeneration {
	case NetworkGenerationColumbus4:
		valconsAddr, err := ValConsPubToAddr(pubkeyidentifier)
		if err != nil {
			return "", fmt.Errorf("failed to convert valconspub(%s) to valconsaddr: %w", pubkeyidentifier, err)
		}
		return valconsAddr, nil
	case NetworkGenerationColumbus5:
		valconsAddr, err := ValConsToAddr(pubkeyidentifier)
		if err != nil {
			return "", fmt.Errorf("failed to convert valcons(%s) to valconsaddr: %w", pubkeyidentifier, err)
		}
		return valconsAddr, nil
	default:
		panic("unknown network generation. available variants: columbus-4 or columbus-5")
	}
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
