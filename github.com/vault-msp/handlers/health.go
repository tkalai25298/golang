package handlers

import (
	"net/http"
	"encoding/json"
)

//HealthCheck to check server health
func HealthCheck(rw http.ResponseWriter,req *http.Request) {

	var data = Response{Response: "Welcome to VAULT-GOLANG API"}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}