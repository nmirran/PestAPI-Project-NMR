package core

import (
	"strings"
	"pestapi/model"
)

// Search Functional
func SearchKeyword(keyword string) []model.Pest {
	return PestStore.Filter(func(p model.Pest) bool {
		return strings.Contains(strings.ToLower(p.CommonName), strings.ToLower(keyword))
	})
}
