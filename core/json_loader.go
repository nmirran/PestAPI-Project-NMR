package core

import (
	"encoding/json"
	"os"
	"pestapi/model"
	"fmt"
)

func LoadPestsFromJSON(path string) ([]model.Pest, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	var pests []model.Pest
	if err := json.Unmarshal(content, &pests); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return pests, nil
}
