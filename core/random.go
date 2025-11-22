package core

import (
    "math/rand"
    "time"
    "pestapi/model"
)

func RandomPest() model.Pest {
    pests := PestStore.GetAll()
    
    if len(pests) == 0 {
        return model.Pest{}
    }

    rand.Seed(time.Now().UnixNano())
    return pests[rand.Intn(len(pests))]
}
