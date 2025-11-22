package core

import (
	"errors"
	"pestapi/model"
)

type PestRepo struct {
	GetAll   func() []model.Pest
	Add      func(model.Pest)
	FindByID func(int) (model.Pest, error)
	Filter   func(func(model.Pest) bool) []model.Pest
}
var PestStore PestRepo

func init() {

	pests, err := LoadPestsFromJSON("data/pests.json")
	if err != nil {
		panic("FAILED TO LOAD DATASET: " + err.Error())
	}

	// closure untuk menyimpan state secara private
	var pestState = append([]model.Pest(nil), pests...)

	// immutable return
	getAll := func() []model.Pest {
		copyData := make([]model.Pest, len(pestState))
		copy(copyData, pestState)
		return copyData
	}

	// side effect: add
	add := func(p model.Pest) {
		pestState = append(pestState, p)
	}

	// find with recursive
	var findRecursive func([]model.Pest, int, int) (model.Pest, error)

	findRecursive = func(arr []model.Pest, id int, i int) (model.Pest, error) {
		switch {
		case i >= len(arr):
			return model.Pest{}, errors.New("not found")
		case arr[i].ID == id:
			return arr[i], nil
		default: 
			return findRecursive(arr, id, i+1)
		}
	}

	findByID := func(id int) (model.Pest, error) {
		return findRecursive(getAll(), id, 0)
	}

	// hof-filter
	filter := func(pred func(model.Pest) bool) []model.Pest {
		result := []model.Pest{}
		for _, p := range getAll() {
			if pred(p) {
				result = append(result, p)
			}
		}
		return result
	}
	PestStore = PestRepo{
		GetAll:   getAll,
		Add:      add,
		FindByID: findByID,
		Filter:   filter,
	}
}

// Functional error handling
func (repo PestRepo) FindByID_Func(id int) Result[model.Pest] {
	p, err := repo.FindByID(id)
	if err != nil {
		return ErrResult[model.Pest](err)
	}
	return Ok(p)
}
