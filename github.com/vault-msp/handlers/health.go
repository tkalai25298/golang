package handlers

import (
	"fmt"
	"net/http"
)

//HealthCheck to check server health
func HealthCheck(rw http.ResponseWriter,req *http.Request) {
	
	fmt.Println("health check hit")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	rw.Write([]byte("Welcome to VAULT-GOLANG API"))
}