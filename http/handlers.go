package httpapi

import (
	"pestapi/core"
	"pestapi/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
	
func GetAllPestsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, core.PestStore.GetAll())
}

func GetCommonNamesHandler(w http.ResponseWriter, r *http.Request) {
	pests := core.PestStore.GetAll()

	names := core.Map(pests, func(p model.Pest) string {
		return p.CommonName
	})
	
	writeJSON(w, http.StatusOK, names)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		http.Error(w, "keyword required", http.StatusBadRequest)
		return
	}
	result := core.SearchKeyword(keyword)
	writeJSON(w, http.StatusOK, result)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	total := core.TotalSymptoms()
	writeJSON(w, http.StatusOK, map[string]int{"total_symptoms": total})
}

func PureNamesHandler(w http.ResponseWriter, r *http.Request) {
	pests := core.PestStore.GetAll()
	names := core.ExtractCommonNames(pests)
	writeJSON(w, http.StatusOK, names)
}

func AddPestHandler(w http.ResponseWriter, r *http.Request) {
	var newPest model.Pest

	if err := json.NewDecoder(r.Body).Decode(&newPest); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest) 
		return
	}

	core.PestStore.Add(newPest) // SIDE EFFECT

	writeJSON(w, http.StatusCreated, map[string]string{
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
	writeJSON(w, http.StatusOK, pest)
}

func SearchFastHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		http.Error(w, "Keyword parameter is required", http.StatusBadRequest)
		return
	}
	results := core.SearchFast(keyword)
	writeJSON(w, http.StatusOK, results)
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
	writeJSON(w, http.StatusOK, map[string]interface{}{
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

	writeJSON(w, http.StatusOK, result.Value)
}
func ImmutableDemoHandler(w http.ResponseWriter, r *http.Request) {
	updated := core.UpdateCommonNameImmutable(1, "New Name")
	removed := core.RemovePestImmutable(1)
	deepcopy := core.DeepCopyPests(core.PestStore.GetAll())
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"immutable_update": updated,
		"immutable_remove": removed,
		"deep_copy":        deepcopy,
	})
}

// FILTER BY PART
func FilterByPartHandler(w http.ResponseWriter, r *http.Request) {
    part := r.URL.Query().Get("part")
	if part == "" {
		http.Error(w, "part parameter is required", http.StatusBadRequest)
		return
	}
    pests := core.PestStore.GetAll()
	result := core.FilterByPart(part)(pests)

	writeJSON(w, http.StatusOK, result)
}

// FILTER BY TYPE
func FilterByTypeHandler(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("type")
	if t == "" {
		http.Error(w, "type parameter is required", http.StatusBadRequest)
		return
	}
    pests := core.PestStore.GetAll()
	result := core.FilterByTypeValue(t)(pests)

	writeJSON(w, http.StatusOK, result)
}

// SORT
func SortHandler(w http.ResponseWriter, r *http.Request) {
    field := r.URL.Query().Get("field")
    order := r.URL.Query().Get("order")

    pests := core.PestStore.GetAll()
    result := core.SortPests(pests, field, order)

	writeJSON(w, http.StatusOK, result)
}

// DELETE
func DeletePestHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
    id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

    core.DeletePest(id)
	writeJSON(w, http.StatusOK, map[string]string{"message": "deleted"})
}

// UPDATE MUTABLE
func UpdatePestHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

    var payload model.Pest
   	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest) 
		return
	}

    core.UpdatePest(id, payload)

	writeJSON(w, http.StatusOK, map[string]string{"message": "updated"})
}

// SEARCH SCIENTIFIC
func SearchScientificHandler(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		http.Error(w, "keyword parameter is required", http.StatusBadRequest)
		return
	}
    result := core.SearchScientific(keyword)
	writeJSON(w, http.StatusOK, result)
}

// RANDOM
func RandomPestHandler(w http.ResponseWriter, r *http.Request) {
    p := core.RandomPest()
	if p.ID == 0 {
		http.Error(w, "no pests available", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, p)
}

// FULL STATS
func StatsFullHandler(w http.ResponseWriter, r *http.Request) {
    s := core.FullStats()
	writeJSON(w, http.StatusOK, s)
}

// CONCURRENT OPTIMIZED
func SearchConcurrentHandler(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		http.Error(w, "keyword parameter is required", http.StatusBadRequest)
		return
	}
    result := core.SearchConcurrentOptimized(keyword)
	writeJSON(w, http.StatusOK, result)
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
        n, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}
		limit = n
    }

    pests := core.PestStore.GetAll()
    result := core.PipelineAdvanced(pests, t, part, sortField, order, limit)

	writeJSON(w, http.StatusOK, result)
}
