package handler

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "grill-master-service"})
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "created", "message": "order fired to the grill"})
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id, "status": "in_progress"})
}

func ListStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stations := []map[string]interface{}{
		{"name": "grill-1", "type": "grill", "active": true},
		{"name": "grill-2", "type": "grill", "active": true},
		{"name": "fryer-1", "type": "fryer", "active": true},
		{"name": "prep-1", "type": "prep", "active": true},
	}
	json.NewEncoder(w).Encode(stations)
}
