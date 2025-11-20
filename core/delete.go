package core

func DeletePest(id int) {
    pests := PestStore.GetAll()
    newList := []model.Pest{}

    for _, p := range pests {
        if p.ID != id {
            newList = append(newList, p)
        }
    }

    // side effect: mutate internal closure state
    PestStore.Add = func(_ model.Pest) {} // disable temporarily
    var newState = newList
    PestStore.GetAll = func() []model.Pest {
        return append([]model.Pest{}, newState...)
    }
}
