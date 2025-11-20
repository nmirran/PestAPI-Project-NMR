package core

import (
	"pestapi/model"
	"sort"
)

type PestSimple struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Pipeline[T any](data T, steps ...func(T) T) T {
	result := data
	for _, step := range steps {
		result = step(result)
	}
	return result
}

func FilterByType(t string) func([]model.Pest) []model.Pest {
	return func(pests []model.Pest) []model.Pest {
		result := []model.Pest{}
		for _, p := range pests {
			if p.PestType == t {
				result = append(result, p)
			}
		}
		return result
	}
}

func SortByName(pests []model.Pest) []model.Pest {
	copyData := append([]model.Pest{}, pests...)
	sort.Slice(copyData, func(i, j int) bool {
		return copyData[i].CommonName < copyData[j].CommonName
	})
	return copyData
}

func Limit(n int) func([]model.Pest) []model.Pest {
	return func(pests []model.Pest) []model.Pest {
		if len(pests) <= n {
			return pests
		}
		return pests[:n]
	}
}

func MapToSimple(pests []model.Pest) []PestSimple {
	result := []PestSimple{}
	for _, p := range pests {
		result = append(result, PestSimple{
			ID:   p.ID,
			Name: p.CommonName,
		})
	}
	return result
}

func ReduceSymptoms(pests []model.Pest) int {
	total := 0
	for _, p := range pests {
		total += len(p.Symptoms)
	}
	return total
}

func PipelineAdvanced(pests []model.Pest, typeVal, part, sortField, order string, limit int) []model.Pest {

    // filter by pest type
    if typeVal != "" {
        pests = FilterByTypeValue(typeVal)
    }

    // filter by plant part
    if part != "" {
        pests = FilterByPart(part)
    }

    // sort
    if sortField != "" {
        pests = SortPests(pests, sortField, order)
    }

    // limit
    if limit > 0 && limit < len(pests) {
        pests = pests[:limit]
    }

    return pests
}
