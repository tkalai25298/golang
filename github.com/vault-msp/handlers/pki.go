package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	data "github.com/vault-msp/data" //Pki struct
)

//EnablePKI handler to create a pki engine to store certs
func (vault *Vault) EnablePKI(rw http.ResponseWriter, req *http.Request) {

	reqData := data.Pki{}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &reqData)
	if err != nil {
		log.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	
	err = reqData.Validate()

	if err != nil {
		log.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}

	vaultData, err := json.Marshal(reqData.Data)

	//Sending http request to vault server
	resp, err := vault.requestObject.HTTPCall("/v1/sys/mounts/"+reqData.Path,vaultData)

	if err != nil {
		log.Println("[ERROR] Could not send request! Server connection issue ", err)
		http.Error(rw, "Error Unbale to send Vault Server Request ", http.StatusBadGateway)
		return
	}
	
	log.Println("The Status Response ==>> ", resp.StatusCode)

	if resp.StatusCode != 204 {
		log.Println("[ERROR] Non 200 Status Code ", err)
		http.Error(rw, "Error Non 200 Status Code ", http.StatusBadGateway)
		return
	}

}
