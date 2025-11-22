package core

import (
    "strings"
    "pestapi/model"
)

func SearchScientific(keyword string) []model.Pest {
    kw := strings.ToLower(keyword)
    return PestStore.Filter(func(p model.Pest) bool {
        return strings.Contains(strings.ToLower(p.ScientificName),kw)
    })
}
