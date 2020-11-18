package main

import (
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
	gohandlers"github.com/gorilla/handlers"
	"github.com/vault-msp/handlers"
	config"github.com/vault-msp/config"
)


func main() {
	
	l := log.New(os.Stdout, "vault-api", log.LstdFlags)  //creating a logger

	config, err := config.SetConfig() //getting env variables for vault server URL & Token
	if err != nil {
		log.Fatalf(err.Error())
	}

	VaultRequest := handlers.NewVaultRequest(l,config.VaultURL,config.VaultToken)	//handler for vault server object

	router := mux.NewRouter() //Gorilla mux router

	postRouter := router.Methods(http.MethodPost).Subrouter()	//Router for post methods

	postRouter.HandleFunc("/pki", VaultRequest.EnablePKI)	//HandlerFunc to register handlers
	postRouter.HandleFunc("/ca", VaultRequest.IssueCA)
	postRouter.HandleFunc("/role", VaultRequest.CreateRole)
	postRouter.HandleFunc("/issueCert", VaultRequest.IssueCert)
	postRouter.HandleFunc("/msp",VaultRequest.VaultInterface)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/health",handlers.HealthCheck)

	//CORS
	origins := gohandlers.CORS(gohandlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")}))

	port := os.Getenv("PORT")
	log.Println("Server running on port",port)
	err = http.ListenAndServe(":"+port, origins(router))

	if err != nil {
        log.Fatal("ListenAndServe: ", err)
	}
}