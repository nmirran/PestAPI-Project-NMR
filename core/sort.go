package core

import (
    "sort"
    "pestapi/model"
    "strings"
)

func SortPests(pests []model.Pest, field, order string) []model.Pest {
    sorted := append([]model.Pest{}, pests...)

    field = strings.ToLower(strings.TrimSpace(field))
    order = strings.ToLower(strings.TrimSpace(order))

    if field == "" {
        field = "common_name"
    }
    asc := order != "desc"

    sort.Slice(sorted, func(i, j int) bool {
        switch field {
        case "common_name","name":
            if asc{
                return strings.ToLower(sorted[i].CommonName) <
					strings.ToLower(sorted[j].CommonName)
            }
            return strings.ToLower(sorted[i].CommonName) >
				strings.ToLower(sorted[j].CommonName)

        case "symptoms":
            if asc {
				return len(sorted[i].Symptoms) < len(sorted[j].Symptoms)
			}
			return len(sorted[i].Symptoms) > len(sorted[j].Symptoms)

		default:
			if asc {
				return strings.ToLower(sorted[i].CommonName) <
					strings.ToLower(sorted[j].CommonName)
			}
			return strings.ToLower(sorted[i].CommonName) >
				strings.ToLower(sorted[j].CommonName)
		}
    })

    return sorted
}

func SortPestsFunc(field, order string) func([]model.Pest) []model.Pest {
    return func(pests []model.Pest) []model.Pest {
        return SortPests(pests, field, order)
    }
}
