package core

import (
    "sync"
    "strings"
    "pestapi/model"
)

func SearchConcurrentOptimized(keyword string) []model.Pest {
    pests := PestStore.GetAll()
    workers := 4
    chunkSize := len(pests) / workers

    var wg sync.WaitGroup
    resultChan := make(chan model.Pest, len(pests))

    for i := 0; i < workers; i++ {
        start := i * chunkSize
        end := start + chunkSize
        if i == workers-1 {
            end = len(pests)
        }

        wg.Add(1)
        go func(part []model.Pest) {
            defer wg.Done()
            for _, p := range part {
                if strings.Contains(strings.ToLower(p.CommonName), strings.ToLower(keyword)) {
                    resultChan <- p
                }
            }
        }(pests[start:end])
    }

    wg.Wait()
    close(resultChan)

    results := []model.Pest{}
    for r := range resultChan {
        results = append(results, r)
    }

    return results
}
