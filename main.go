package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var db = NewStore()

type SearchRequest struct {
	Vector []float32 `json:"vector"`
	K      int       `json:"k"`
}
type InsertRequest struct {
	ID     string    `json:"id"`
	Vector []float32 `json:"vector"`
}

func handleInsertion(w http.ResponseWriter, r *http.Request) {
	var request InsertRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	if len(request.Vector) == 0 {
		http.Error(w, "Vector cannot be empty", http.StatusBadRequest)
		return
	}
	db.Insert(request.ID, request.Vector)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Vector saved to RAM"))
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	var request SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	if request.K <= 0 {
		request.K = 1
	}
	start := time.Now()
	results := db.Search(request.Vector, request.K)
	duration := time.Since(start)
	response := map[string]interface{}{
		"results": results,
		"count":   len(results),
		"latency": duration.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/insert", handleInsertion)
	http.HandleFunc("/search", HandleSearch)
	fmt.Println("Hermes")
	fmt.Println("A hybrid C++ and Go")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
