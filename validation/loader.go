package validation

import (
	"beneficiary-validate/model"
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfigFromFile(currency string) (*model.ConfigRule, error) {
	bytes, err := os.ReadFile("config/rules.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var jsonConfig map[string]model.ConfigRule
	if err := json.Unmarshal(bytes, &jsonConfig); err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	config, ok := jsonConfig[currency]
	if !ok {
		return nil, fmt.Errorf("currency not supported: %s", currency)
	}

	return &config, nil
}
