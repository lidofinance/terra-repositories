package mocks

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client/wasm"

	"github.com/go-openapi/runtime"
)

const (
	terraWASMServiceSamplesFolder = "terra_wasm_service"

	V1WhitelistedValidatorsContractAddress = "terra1mtwph2juhj0rvjz7dy92gvl6xvukaxu8rfv8ts"
	V2WhitelistedValidatorsContractAddress = "terra10wt548y4y3xeqfrqsgqlqh424lll8fqxp6dyed"

	v1WhitelistedValidatorsQuery = "whitelisted_validators"
	v2WhitelistedValidatorsQuery = "get_validators_for_delegation"

	v1WhitelistedValidatorsSampleFile = "v1_whitelisted_validators.json"
	v2WhitelistedValidatorsSampleFile = "v2_whitelisted_validators.json"
)

type TerraWASMService struct{}

func (s *TerraWASMService) GetV1WasmCodeCodeID(params *wasm.GetV1WasmCodeCodeIDParams, opts ...wasm.ClientOption) (*wasm.GetV1WasmCodeCodeIDOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) GetV1WasmContractContractAddress(params *wasm.GetV1WasmContractContractAddressParams, opts ...wasm.ClientOption) (*wasm.GetV1WasmContractContractAddressOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) GetV1WasmContracts(params *wasm.GetV1WasmContractsParams, opts ...wasm.ClientOption) (*wasm.GetV1WasmContractsOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) GetWasmCodesCodeID(params *wasm.GetWasmCodesCodeIDParams, opts ...wasm.ClientOption) (*wasm.GetWasmCodesCodeIDOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) GetWasmContractsContractAddress(params *wasm.GetWasmContractsContractAddressParams, opts ...wasm.ClientOption) (*wasm.GetWasmContractsContractAddressOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) GetWasmContractsContractAddressStore(params *wasm.GetWasmContractsContractAddressStoreParams, opts ...wasm.ClientOption) (*wasm.GetWasmContractsContractAddressStoreOK, error) {
	var filename string
	switch {
	case strings.Contains(params.QueryMsg, v1WhitelistedValidatorsQuery) && params.ContractAddress == V1WhitelistedValidatorsContractAddress:
		filename = v1WhitelistedValidatorsSampleFile
	case strings.Contains(params.QueryMsg, v2WhitelistedValidatorsQuery) && params.ContractAddress == V2WhitelistedValidatorsContractAddress:
		filename = v2WhitelistedValidatorsSampleFile
	default:
		return nil, fmt.Errorf("contract not found")
	}
	data, err := readSample(terraWASMServiceSamplesFolder, filename)
	if err != nil {
		return nil, err
	}
	resp := &wasm.GetWasmContractsContractAddressStoreOK{
		Payload: &wasm.GetWasmContractsContractAddressStoreOKBody{},
	}
	if err := json.Unmarshal(data, resp.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to resp: %w", err)
	}
	return resp, nil
}

func (s *TerraWASMService) GetWasmContractsContractAddressStoreRaw(params *wasm.GetWasmContractsContractAddressStoreRawParams, opts ...wasm.ClientOption) (*wasm.GetWasmContractsContractAddressStoreRawOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) GetWasmParameters(params *wasm.GetWasmParametersParams, opts ...wasm.ClientOption) (*wasm.GetWasmParametersOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) PostWasmCodes(params *wasm.PostWasmCodesParams, opts ...wasm.ClientOption) (*wasm.PostWasmCodesOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) PostWasmCodesCodeID(params *wasm.PostWasmCodesCodeIDParams, opts ...wasm.ClientOption) (*wasm.PostWasmCodesCodeIDOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) PostWasmCodesCodeIDMigrate(params *wasm.PostWasmCodesCodeIDMigrateParams, opts ...wasm.ClientOption) (*wasm.PostWasmCodesCodeIDMigrateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) PostWasmContractsContractAddress(params *wasm.PostWasmContractsContractAddressParams, opts ...wasm.ClientOption) (*wasm.PostWasmContractsContractAddressOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) PostWasmContractsContractAddressAdminClear(params *wasm.PostWasmContractsContractAddressAdminClearParams, opts ...wasm.ClientOption) (*wasm.PostWasmContractsContractAddressAdminClearOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) PostWasmContractsContractAddressAdminUpdate(params *wasm.PostWasmContractsContractAddressAdminUpdateParams, opts ...wasm.ClientOption) (*wasm.PostWasmContractsContractAddressAdminUpdateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) PostWasmContractsContractAddressMigrate(params *wasm.PostWasmContractsContractAddressMigrateParams, opts ...wasm.ClientOption) (*wasm.PostWasmContractsContractAddressMigrateOK, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *TerraWASMService) SetTransport(transport runtime.ClientTransport) {}
