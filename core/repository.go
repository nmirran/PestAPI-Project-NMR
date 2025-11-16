package core

import (
	"errors"
	"pestapi/data"
	"pestapi/model"
)

type PestRepo struct {
	GetAll   func() []model.Pest
	Add      func(model.Pest)
	FindByID func(int) (model.Pest, error)
	Filter   func(func(model.Pest) bool) []model.Pest
}

var PestStore = NewPestRepo(data.Pests)

func NewPestRepo(initial []model.Pest) PestRepo {
	// immutable slice
	pests := append([]model.Pest(nil), initial...)
	// non-recursive
	getAll := func() []model.Pest {
		return append([]model.Pest(nil), pests...)
	}
	add := func(p model.Pest) {
		pests = append(pests, p)
	}

	// recursvie
	var findRecursive func([]model.Pest, int, int) (model.Pest, error)
	findRecursive = func(arr []model.Pest, id int, i int) (model.Pest, error) {
		if i >= len(arr) {
			return model.Pest{}, errors.New("not found")
		}
		if arr[i].ID == id {
			return arr[i], nil
		}
		return findRecursive(arr, id, i+1)
	}

	findByID := func(id int) (model.Pest, error) {
		return findRecursive(getAll(), id, 0)
	}

	// hof
	filter := func(pred func(model.Pest) bool) []model.Pest {
		result := []model.Pest{}
		for _, p := range getAll() {
			if pred(p) {
				result = append(result, p)
			}
		}
		return result
	}

	return PestRepo{
		GetAll:   getAll,
		Add:      add,
		FindByID: findByID,
		Filter:   filter,
	}
}

func (repo PestRepo) FindByID_Func(id int) Result[model.Pest] {
	p, err := repo.FindByID(id)
	if err != nil {
		return Err[model.Pest](err)
	}
	return Ok(p)
}