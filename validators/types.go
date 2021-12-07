package validators

type ValidatorInfo struct {
	Address        string // terravaloper
	PubKey         string // terravalcons
	Moniker        string
	CommissionRate float64
	Jailed         bool
}

type HubWhitelistedValidatorsRequest struct {
	WhitelistedValidators struct{} `json:"whitelisted_validators"`
}

type HubWhitelistedValidatorsResponse struct {
	Validators []string `json:"validators"`
}

type ValidatorRegistryValidatorsRequest struct {
	WhitelistedValidators struct{} `json:"get_validators_for_delegation"`
}

type ValidatorRegistryValidatorsResponse = []Validator

type Validator struct {
	Address        string `json:"address"`
	TotalDelegated string `json:"total_delegated"` // Uint128
}
