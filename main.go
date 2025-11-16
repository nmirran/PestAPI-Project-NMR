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
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
