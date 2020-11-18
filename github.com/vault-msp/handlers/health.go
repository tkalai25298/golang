package handlers

import (
	"fmt"
	"net/http"
)

//HealthCheck to check server health
func HealthCheck(rw http.ResponseWriter,req *http.Request) {
	fmt.Fprintf(rw,"Welcome to VAULT-GOLANG API")
}