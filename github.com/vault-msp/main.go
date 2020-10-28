package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vault-msp/handlers"
)

// struct defn

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/pki", handlers.EnablePKI).Methods("POST")
	router.HandleFunc("/ca",handlers.IssueCA).Methods("POST")
	router.HandleFunc("/role",handlers.CreateRole).Methods("POST")
	router.HandleFunc("/issueCert",handlers.IssueCert).Methods("POST")

	http.ListenAndServe(":8000", router)

}