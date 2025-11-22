package core

import "pestapi/model"

type Stats struct {
    TotalPests int            `json:"total_pests"`
    ByType     map[string]int `json:"by_type"`
    TotalSymptoms int         `json:"total_symptoms"`
}

func TotalSymptoms() int {
	pests := PestStore.GetAll()
	return Reduce(pests, 0, func(acc int, p model.Pest) int {
		return acc + len(p.Symptoms)
	})
}

func FullStats() Stats {
    pests := PestStore.GetAll()

    byType := Reduce(pests, map[string]int{}, func(acc map[string]int, p model.Pest) map[string]int {
		newMap := make(map[string]int, len(acc)+1)
		for k, v := range acc {
			newMap[k] = v
		}
		newMap[p.PestType]++
		return newMap
	})

	totalSymptoms := Reduce(pests, 0, func(acc int, p model.Pest) int {
		return acc + len(p.Symptoms)
	})

	return Stats{
		TotalPests:    len(pests),
		ByType:        byType,
		TotalSymptoms: totalSymptoms,
	}
}