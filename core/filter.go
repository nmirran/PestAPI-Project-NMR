package core

import (
    "strings"
    "pestapi/model"
)

func FilterByPart(part string) []model.Pest {
    return PestStore.Filter(func(p model.Pest) bool {
        for _, af := range p.AffectedParts {
            if strings.EqualFold(af, part) {
                return true
            }
        }
        return false
    })
}

func FilterByTypeValue(t string) []model.Pest {
    return PestStore.Filter(func(p model.Pest) bool {
        return strings.EqualFold(p.PestType, t)
    })
}
