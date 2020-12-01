package handlers

import (
	"log"
	"net/http"
)

//HealthCheck to check server health
func HealthCheck(rw http.ResponseWriter,req *http.Request) {

	var data = Response{Response: "Welcome to VAULT-GOLANG API"}
	rw.Header().Set("Content-Type", "application/json")
	
	err := data.JSONResponse(rw)
	if err != nil {
		log.Println("[ERROR] Could not Marshal response json ", err)
		http.Error(rw, "Error Unbale to marshal response json ", http.StatusBadGateway)
		return
	}
}