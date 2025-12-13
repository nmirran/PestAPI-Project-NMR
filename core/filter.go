package core

import (
    "strings"
    "pestapi/model"
)

func FilterByPart(part string) func([]model.Pest) []model.Pest {
    normalized := strings.ToLower(strings.TrimSpace(part))

    return func (pests []model.Pest) []model.Pest {
        result := []model.Pest{}
        for _, p := range pests {
            for _, ap := range p.AffectedParts {
                apNorm := strings.ToLower(strings.TrimSpace(ap))
                if strings.Contains(apNorm, normalized) {
                    result = append(result, p)
                    break
                }
            }
        }
        return result
    }
}

func FilterByTypeValue(t string) func([]model.Pest) []model.Pest {
    normalized := strings.ToLower(strings.TrimSpace(t))

    return func (pests []model.Pest) []model.Pest {
        result := []model.Pest{}
        for _, p := range pests {
            if strings.ToLower(strings.TrimSpace(p.PestType)) == normalized {
                    result = append(result, p)
            }
        }
        return result
    }
}
