package handle

import (
	"encoding/json"
	"github.com/zerodays/sistem-users/internal/database"
	"net/http"
)

var KillSwitch = false

type livenessResponse struct {
	Status     string `json:"status"`
	KillSwitch string `json:"kill_switch"`
}

type readinessResponse struct {
	Status   string `json:"status"`
	DBStatus string `json:"db_status"`
}

// Returns weather microservice is alive.
func LivenessHandle(w http.ResponseWriter, _ *http.Request) {
	resp := livenessResponse{}
	ok := true

	// Check kill switch
	if !KillSwitch {
		resp.KillSwitch = "UP"
	} else {
		ok = false
		resp.KillSwitch = "DOWN"
	}

	// Set global status
	if ok {
		resp.Status = "UP"
		w.WriteHeader(http.StatusOK)
	} else {
		resp.Status = "DOWN"
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	res, _ := json.Marshal(resp)
	_, _ = w.Write(res)
}

// Returns weather microservice is ready to accept connections.
func ReadinessHandle(w http.ResponseWriter, _ *http.Request) {
	resp := readinessResponse{}
	ok := true

	// Check if database is ready.
	dbErr := database.DB.Ping()
	if dbErr == nil {
		resp.DBStatus = "UP"
	} else {
		ok = false
		resp.DBStatus = "DOWN: " + dbErr.Error()
	}

	// Set global status
	if ok {
		resp.Status = "UP"
		w.WriteHeader(http.StatusOK)
	} else {
		resp.Status = "DOWN"
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	res, _ := json.Marshal(resp)
	_, _ = w.Write(res)
}

// Handles changing of kill switch
func KillSwitchHandle(w http.ResponseWriter, _ *http.Request) {
	KillSwitch = true
}
