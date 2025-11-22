package core

import (
    "errors"
    "pestapi/model"
)

// MUTABLE update
func UpdatePest(id int, payload model.Pest) error {
    pests := PestStore.GetAll()
    newList := []model.Pest{}

    found := false
    for _, p := range pests {
        if p.ID == id {
            found = true
            newList = append(newList, payload)
        } else {
            newList = append(newList, p)
        }
    }

    if !found {
        return errors.New("pest not found")
    }

    // mutate state
    PestStore.Add = func(_ model.Pest) {} 
    var ns = newList
    PestStore.GetAll = func() []model.Pest {
        return append([]model.Pest{}, ns...)
    }

    return nil
}
