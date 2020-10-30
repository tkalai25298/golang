package main

import (
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/vault-msp/handlers"
)


func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	ph := handlers.NewPKI(l)

	router := mux.NewRouter()


	router.HandleFunc("/pki", ph.EnablePKI).Methods("POST")
	router.HandleFunc("/ca",handlers.IssueCA).Methods("POST")
	router.HandleFunc("/role",handlers.CreateRole).Methods("POST")
	router.HandleFunc("/issueCert",handlers.IssueCert).Methods("POST")

	http.ListenAndServe(":8000", router)

}