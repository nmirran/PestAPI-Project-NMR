package core

import (
	"encoding/json"
	"os"
	"pestapi/model"
)

func LoadPestsFromJSON(path string) ([]model.Pest, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pests []model.Pest
	err = json.Unmarshal(content, &pests)
	if err != nil {
		return nil, err
	}

	return pests, nil
}
