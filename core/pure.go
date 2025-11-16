package core

import "pestapi/model"

func ExtractCommonNames(pests []model.Pest) []string {
	result := make([]string, len(pests))
	for i, p := range pests {
		result[i] = p.CommonName
	}
	return result
}

func CountTotalSymptoms(pests []model.Pest) int {
	total := 0
	for _, p := range pests {
		total += len(p.Symptoms)
	}
	return total
}
