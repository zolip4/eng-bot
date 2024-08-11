package controllers

import (
	"encoding/json"
	"engbot/models"
	"net/http"
	"strconv"
	"sync"
)

var (
	records   = make(map[int]models.Record)
	nextID    = 1
	recordsMu sync.Mutex
)

func HandleRecords(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createRecord(w, r)
	case http.MethodPut:
		updateRecord(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	var record models.Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recordsMu.Lock()
	record.ID = nextID
	records[nextID] = record
	nextID++
	recordsMu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var record models.Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recordsMu.Lock()
	if _, exists := records[id]; !exists {
		recordsMu.Unlock()
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}

	record.ID = id
	records[id] = record
	recordsMu.Unlock()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(record)
}
