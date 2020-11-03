package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	config "github.com/vault-msp/config"
	data "github.com/vault-msp/data" //Pki struct
	"github.com/vault-msp/httpreq"
)

//EnablePKI handler to create a pki engine to store certs
func EnablePKI(rw http.ResponseWriter, req *http.Request) {

	reqData := data.Pki{}

	config, err := config.SetConfig() //getting env variables for vault server
	if err != nil {
		log.Fatalf(err.Error())
	}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}

	err = json.Unmarshal(reqBody, &reqData)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	
	err = reqData.Validate()

	if err != nil {
		log.Fatal("json validation error",err)
	}

	vaultData, err := json.Marshal(reqData.Data)
	log.Printf("%s", vaultData)

	//Creating the Request body object
	reqObj := httpreq.CreateRequest(http.MethodPost, "http://"+config.VaultURL+"/v1/sys/mounts/"+reqData.Path, config.VaultToken, vaultData)
	//Sending http request to vault server
	resp, err := reqObj.HTTPCall()

	if err != nil {
		log.Fatal("could not send request! Server connection issue")
	}
	
	log.Println("The Status Response ==>> ", resp.StatusCode)

	if resp.StatusCode != 204 {
		log.Panicf("NON 204 STATUS CODE")
	}

	defer resp.Body.Close()
}
