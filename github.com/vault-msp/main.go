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
	getRouter.HandleFunc("/roles",handlers.ListRoles)

	//CORS
	c := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedHeaders([]string{"Authorization","Content-Type","Access-Control-Allow-Origin"}),
		gohandlers.AllowedMethods([]string{"GET","POST"}),
	)
  

	port := os.Getenv("PORT")
	log.Println("Server running on port",port)
	server := http.Server{
		Addr: ":"+port,
		Handler: c(router),
	}
	err = server.ListenAndServe()
	if err != nil{
		log.Fatal("Err:" ,err)
	}
	
}