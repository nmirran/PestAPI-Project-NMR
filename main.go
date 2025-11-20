package main

import (
	"fmt"
	"log"
	"net/http"

	httpapi "pestapi/http"
)

func main() {

	http.HandleFunc("/pests", httpapi.GetAllPestsHandler)
	http.HandleFunc("/pests/names", httpapi.GetCommonNamesHandler)
	http.HandleFunc("/pests/search", httpapi.SearchHandler)
	http.HandleFunc("/pests/stats", httpapi.StatsHandler)
	http.HandleFunc("/pests/pure/names", httpapi.PureNamesHandler)
	http.HandleFunc("/pests/add", httpapi.AddPestHandler)
	http.HandleFunc("/pests/get", httpapi.GetPestByIDHandler)
	http.HandleFunc("/pests/search-fast", httpapi.SearchFastHandler)
	http.HandleFunc("/pests/pipeline", httpapi.PipelineDemoHandler)
	http.HandleFunc("/pests/get-func", httpapi.GetPestByID_FuncHandler)
	http.HandleFunc("/pests/immutable-demo", httpapi.ImmutableDemoHandler)
	http.HandleFunc("/pests/by-part", httpapi.FilterByPartHandler)
	http.HandleFunc("/pests/by-type", httpapi.FilterByTypeHandler)
	http.HandleFunc("/pests/sorted", httpapi.SortHandler)
	http.HandleFunc("/pests/delete", httpapi.DeletePestHandler)
	http.HandleFunc("/pests/update", httpapi.UpdatePestHandler)
	http.HandleFunc("/pests/random", httpapi.RandomPestHandler)
	http.HandleFunc("/pests/search-scientific", httpapi.SearchScientificHandler)
	http.HandleFunc("/pests/stats/full", httpapi.StatsFullHandler)
	http.HandleFunc("/pests/search-concurrent", httpapi.SearchConcurrentHandler)
	http.HandleFunc("/pests/pipeline-advanced", httpapi.PipelineAdvancedHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
