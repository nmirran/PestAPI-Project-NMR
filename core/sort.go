package core

import (
    "sort"
    "pestapi/model"
)

func SortPests(pests []model.Pest, field, order string) []model.Pest {
    sorted := append([]model.Pest{}, pests...)

    sort.Slice(sorted, func(i, j int) bool {
        switch field {
        case "name":
            if order == "desc" {
                return sorted[i].CommonName > sorted[j].CommonName
            }
            return sorted[i].CommonName < sorted[j].CommonName]

        case "symptoms":
            if order == "desc" {
                return len(sorted[i].Symptoms) > len(sorted[j].Symptoms)
            }
            return len(sorted[i].Symptoms) < len(sorted[j].Symptoms)
        }
        return false
    })

    return sorted
}
