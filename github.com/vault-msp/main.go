package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vault-msp/handlers"
)


func main() {

	router := mux.NewRouter()

	pki := handlers.EnablePKI()

	router.HandleFunc("/pki", handlers.EnablePKI).Methods("POST")
	router.HandleFunc("/ca",handlers.IssueCA).Methods("POST")
	router.HandleFunc("/role",handlers.CreateRole).Methods("POST")
	router.HandleFunc("/issueCert",handlers.IssueCert).Methods("POST")

	http.ListenAndServe(":8000", router)

}