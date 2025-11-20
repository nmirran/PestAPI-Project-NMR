package core

import (
    "math/rand"
    "time"
    "pestapi/model"
)

func RandomPest() model.Pest {
    pests := PestStore.GetAll()
    rand.Seed(time.Now().UnixNano())
    return pests[rand.Intn(len(pests))]
}
