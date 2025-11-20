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

// FILTER BY PART
func FilterByPartHandler(w http.ResponseWriter, r *http.Request) {
    part := r.URL.Query().Get("part")
    result := core.FilterByPart(part)
    json.NewEncoder(w).Encode(result)
}

// FILTER BY TYPE
func FilterByTypeHandler(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("type")
    result := core.FilterByTypeValue(t)
    json.NewEncoder(w).Encode(result)
}

// SORT
func SortHandler(w http.ResponseWriter, r *http.Request) {
    field := r.URL.Query().Get("field")
    order := r.URL.Query().Get("order")

    pests := core.PestStore.GetAll()
    result := core.SortPests(pests, field, order)

    json.NewEncoder(w).Encode(result)
}

// DELETE
func DeletePestHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    core.DeletePest(id)
    json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
}

// UPDATE MUTABLE
func UpdatePestHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    var payload model.Pest
    json.NewDecoder(r.Body).Decode(&payload)

    core.UpdatePest(id, payload)

    json.NewEncoder(w).Encode(map[string]string{"message": "updated"})
}

// SEARCH SCIENTIFIC
func SearchScientificHandler(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("key")
    result := core.SearchScientific(keyword)
    json.NewEncoder(w).Encode(result)
}

// RANDOM
func RandomPestHandler(w http.ResponseWriter, r *http.Request) {
    p := core.RandomPest()
    json.NewEncoder(w).Encode(p)
}

// FULL STATS
func StatsFullHandler(w http.ResponseWriter, r *http.Request) {
    s := core.FullStats()
    json.NewEncoder(w).Encode(s)
}

// CONCURRENT OPTIMIZED
func SearchConcurrentHandler(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("keyword")
    result := core.SearchConcurrentOptimized(keyword)
    json.NewEncoder(w).Encode(result)
}

// PIPELINE ADVANCED
func PipelineAdvancedHandler(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("type")
    part := r.URL.Query().Get("part")
    sortField := r.URL.Query().Get("sort")
    order := r.URL.Query().Get("order")
    limitStr := r.URL.Query().Get("limit")

    limit := 0
    if limitStr != "" {
        limit, _ = strconv.Atoi(limitStr)
    }

    pests := core.PestStore.GetAll()
    result := core.PipelineAdvanced(pests, t, part, sortField, order, limit)

    json.NewEncoder(w).Encode(result)
}
