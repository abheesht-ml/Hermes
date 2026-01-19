package main

import (
	"fmt"
	"sort"
	"sync"
)

type VectorRecord struct {
	ID     string    `json:"id"`
	Vector []float32 `json:"vector"`
}
type SearchResult struct {
	ID       string  `json:"id"`
	Distance float32 `json:"distance"`
}
type MemoryStore struct {
	mu   sync.RWMutex
	data []VectorRecord
}

func NewStore() *MemoryStore {
	return &MemoryStore{
		data: make([]VectorRecord, 0),
	}
}
func (s *MemoryStore) Insert(id string, vector []float32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	record := VectorRecord{
		ID:     id,
		Vector: vector,
	}
	s.data = append(s.data, record)
	fmt.Printf("Inserted '%d' vectors", len(s.data))
}
func (s *MemoryStore) Search(query []float32, k int) []SearchResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var results []SearchResult
	for _, record := range s.data {
		dist := EuclideanDistance(query, record.Vector)
		if dist >= 0 {
			results = append(results, SearchResult{
				ID:       record.ID,
				Distance: dist,
			})
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})
	if k > len(results) {
		k = len(results)
	}
	return results[:k]
}
