package core

import (
    "strings"
    "pestapi/model"
)

func FilterByPart(part string) func([]model.Pest) []model.Pest {
    return func (pests []model.Pest) []model.Pest {
        result := []model.Pest{}
        for _, p := range pests {
            for _, af := range p.AffectedParts {
                if strings.EqualFold(af, part) {
                    result = append(result, p)
                    break
                }
            }
        }
        return result
    }
}

func FilterByTypeValue(t string) func([]model.Pest) []model.Pest {
    return func (pests []model.Pest) []model.Pest {
        result := []model.Pest{}
        for _, p := range pests {
            if strings.EqualFold(p.PestType, t) {
                    result = append(result, p)
            }
        }
        return result
    }
}
