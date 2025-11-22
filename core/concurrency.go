package core

import (
	"strings"
	"pestapi/model"
)

func searchWorker(keyword string, pests []model.Pest, resultChan chan<- model.Pest, done chan<- struct{}) {
	for _, p := range pests {
		if strings.Contains(strings.ToLower(p.CommonName), strings.ToLower(keyword)){
			resultChan <- p
		}
	}
	done <- struct{}{}
}

func SearchFast(keyword string) []model.Pest {
	pests := PestStore.GetAll()

	// Divide into 4 chunks 
	chunkSize := (len(pests) + 3) / 4
	parts := [][]model.Pest{}
	for i := 0; i < len(pests); i+= chunkSize {
		end := i +chunkSize
		if end > len(pests) {
			end = len(pests)
		}
		parts = append(parts, pests[i:end])
	}

	resultChan := make(chan model.Pest, 100)
	doneChan := make(chan struct{}, len(parts))

	for _, part := range parts {
		go searchWorker(keyword, part, resultChan, doneChan)
	}
	results := []model.Pest{}
	doneCount := 0
	totalWorkers := len(parts)
	
	for {
		select {
		case p := <-resultChan:
			results = append(results, p)
		case <-doneChan:
			doneCount++
			if doneCount == totalWorkers {
				close(resultChan)
				close(doneChan)
				return results
			}
		}
	}
}