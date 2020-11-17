package handlers

import (
	"net/http"
)

//HealthCheck to check server health
func HealthCheck(rw http.ResponseWriter,req *http.Request) {
	rw.Write([]byte("Welcome to Vault Certs Generation API"))
}