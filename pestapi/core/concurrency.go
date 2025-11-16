package core

import (
	"strings"
	"pestapi/model"
)

func searchWorker(keyword string, pests []model.Pest, resultChan chan model.Pest, done chan bool) {
	for _, p := range pests {
		if strings.Contains(strings.ToLower(p.CommonName), strings.ToLower(keyword)){
			resultChan <- p
		}
	}
	done <- true
}

func SearchFast(keywoard string) []model.Pest {
	pests := PestStore.GetAll()

	chunkSize := (len(pests) + 3) / 4
	parts := [][]model.Pest{}
	for i := 0; i < len(pests); i+= chunkSize {
		end := i +chunkSize
		if end > len(pests) {
			end = len(pests)
		}
		parts = append(parts, pests[i:end])
	}

	resultChan := make(chan model.Pest)
	doneChan := make(chan bool)
	for _, part := range parts {
		go searchWorker(keywoard, part, resultChan, doneChan)
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