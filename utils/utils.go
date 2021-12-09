package utils

import (
	"encoding/json"
	"fmt"
)

func CastMapToStruct(m interface{}, ret interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal body interface{}: %w", err)
	}

	err = json.Unmarshal(data, ret)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}
