package api

import (
	"encoding/json"
	"net/http"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func HandleIntake(w http.ResponseWriter, r *http.Request) {
	// TODO: customer intake agent
	w.WriteHeader(http.StatusNotImplemented)
}

func HandlePlan(w http.ResponseWriter, r *http.Request) {
	// TODO: planning committee agent
	w.WriteHeader(http.StatusNotImplemented)
}

func HandleSpawn(w http.ResponseWriter, r *http.Request) {
	// TODO: agent spawner
	w.WriteHeader(http.StatusNotImplemented)
}
