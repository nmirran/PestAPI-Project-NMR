package core

import "pestapi/model"

func TotalSymptoms() int {
	pests := PestStore.GetAll()
	return Reduce(pests, 0, func(acc int, p model.Pest) int {
		return acc + len(p.Symptoms)
	})
}
