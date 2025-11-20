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

    stats := Stats{
        TotalPests: len(pests),
        ByType:     map[string]int{},
        TotalSymptoms: 0,
    }

    for _, p := range pests {
        stats.ByType[p.PestType]++
        stats.TotalSymptoms += len(p.Symptoms)
    }

    return stats
}