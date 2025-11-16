package httpapi

import (
	"pestapi/core"
	"pestapi/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetAllPestsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(core.PestStore.GetAll())
}

func GetCommonNamesHandler(w http.ResponseWriter, r *http.Request) {
	pests := core.PestStore.GetAll()

	names := core.Map(pests, func(p model.Pest) string {
		return p.CommonName
	})

	json.NewEncoder(w).Encode(names)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	result := core.SearchKeyword(keyword)
	json.NewEncoder(w).Encode(result)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	total := core.TotalSymptoms()
	json.NewEncoder(w).Encode(map[string]int{"total_symptoms": total})
}

func PureNamesHandler(w http.ResponseWriter, r *http.Request) {
	pests := core.PestStore.GetAll()
	names := core.ExtractCommonNames(pests)
	json.NewEncoder(w).Encode(names)
}

func AddPestHandler(w http.ResponseWriter, r *http.Request) {
	var newPest model.Pest

	json.NewDecoder(r.Body).Decode(&newPest)

	core.PestStore.Add(newPest) // SIDE EFFECT

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Pest added successfully",
	})
}

func GetPestByIDHandler(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("id")
	if ids == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(ids)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	pest, err := core.PestStore.FindByID(id)
	if err != nil {
		http.Error(w, "Pest not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pest)
}

func SearchFastHandler(w http.ResponseWriter, r *http.Request) {
	keywoard := r.URL.Query().Get("keyword")
	if keywoard == "" {
		http.Error(w, "Keyword parameter is required", http.StatusBadRequest)
		return
	}
	results := core.SearchFast(keywoard)
	json.NewEncoder(w).Encode(results)
}

func PipelineDemoHandler(w http.ResponseWriter, r *http.Request) {
	pests := core.PestStore.GetAll()
	t := r.URL.Query().Get("type")
	if t == "" {
		t ="insect"
	}
	limitStr := r.URL.Query().Get("limit")
	limit := 5
	if limitStr != "" {
		n, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = n
		}
	}
	processed := core.Pipeline(
		pests,
		core.FilterByType(t),
		core.SortByName,
		core.Limit(limit),
	)
	output := core.MapToSimple(processed)
	totalSymptoms := core.ReduceSymptoms(processed)
	json.NewEncoder(w).Encode(map[string]interface{}{
	"data":           output,
	"total_symptoms": totalSymptoms,
	})
}

func GetPestByID_FuncHandler(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("id")
	if ids == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(ids)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	result := core.PestStore.FindByID_Func(id)

	if result.IsErr() {
		http.Error(w, result.Err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result.Value)
}
func ImmutableDemoHandler(w http.ResponseWriter, r *http.Request) {
	updated := core.UpdateCommonNameImmutable(1, "New Name")
	removed := core.RemovePestImmutable(1)
	deepcopy := core.DeepCopyPests(core.PestStore.GetAll())
		json.NewEncoder(w).Encode(map[string]interface{}{
		"immutable_update": updated,
		"immutable_remove": removed,
		"deep_copy":        deepcopy,
	})
}