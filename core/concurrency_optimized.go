package core

import (
    "sync"
    "strings"
    "pestapi/model"
)

func SearchConcurrentOptimized(keyword string) []model.Pest {
    pests := PestStore.GetAll()
    workers := 4
    if len(pests) < workers {
        workers = 1
    }

    chunkSize := (len(pests) + workers - 1) / workers
    loweredKeyword := strings.ToLower(keyword)

    var wg sync.WaitGroup
    resultChan := make(chan model.Pest, len(pests))

    for i := 0; i < workers; i++ {
        start := i * chunkSize
        end := start + chunkSize
        if end > len(pests) {
            end = len(pests)
        }

        part := pests[start:end]
        wg.Add(1)

        go func(ps []model.Pest) {
            defer wg.Done()
            for _, p := range ps {
                if strings.Contains(strings.ToLower(p.CommonName), loweredKeyword) {
                    resultChan <- p
                }
            }
        }(part)
    }

    wg.Wait()
    close(resultChan)

    results := []model.Pest{}
    for r := range resultChan {
        results = append(results, r)
    }

    return results
}
