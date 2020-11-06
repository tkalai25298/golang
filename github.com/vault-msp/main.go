package main

import (
	"net/http"
	"log"
	"os"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/vault-msp/handlers"
	httpreq"github.com/vault-msp/httpreq"
	config"github.com/vault-msp/config"
)


func main() {

	l := log.New(os.Stdout, "vault-api", log.LstdFlags)

	config, err := config.SetConfig() //getting env variables for vault server
	if err != nil {
		log.Fatalf(err.Error())
	}
	VaultURL,err := url.Parse(config.VaultURL)

	requestObj := httpreq.CreateRequest(http.MethodPost,VaultURL,config.VaultToken)

	VaultRequest := handlers.NewPKI(l,requestObj)

	router := mux.NewRouter()

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/pki", VaultRequest.EnablePKI)
	postRouter.HandleFunc("/ca", VaultRequest.IssueCA)
	postRouter.HandleFunc("/role", VaultRequest.CreateRole)
	postRouter.HandleFunc("/issueCert", VaultRequest.IssueCert)

	http.ListenAndServe(":8000", router)
	
}