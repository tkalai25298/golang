package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	data "github.com/vault-msp/data" //Pki struct
	"github.com/vault-msp/httpreq"
)

//Enable to create httpClient request object
type Enable struct{
	l *log.Logger
	requestObject httpreq.HTTPClient
}
//NewPKI that returns requestObject
func NewPKI(l *log.Logger,req httpreq.HTTPClient) *Enable{
	return &Enable{l:l,requestObject:req}
}

//EnablePKI handler to create a pki engine to store certs
func (enable *Enable) EnablePKI(rw http.ResponseWriter, req *http.Request) {

	reqData := data.Pki{}

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

	//Sending http request to vault server
	resp, err := enable.requestObject.HTTPCall("/v1/sys/mounts/"+reqData.Path,vaultData)

	if err != nil {
		log.Fatal("could not send request! Server connection issue")
	}
	
	log.Println("The Status Response ==>> ", resp.StatusCode)

	if resp.StatusCode != 204 {
		log.Panicf("NON 204 STATUS CODE")
	}

	defer resp.Body.Close()
}
