package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	data "github.com/vault-msp/data" //Pki struct
)

//EnablePKI handler to create a pki engine to store certs
func (vault *Vault) EnablePKI(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	pki := data.PkiData{}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		vault.l.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &pki)
	if err != nil {
		vault.l.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	
	err = pki.Validate()

	if err != nil {
		vault.l.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}

	pkiPath:= pki.Organization+"CA"
	vaultData, err := json.Marshal(pki)

	//Sending http request to vault server
	resp, err := vault.requestObject.HTTPCall("/v1/sys/mounts/"+pkiPath,vaultData)
	// vault.l.Printf("%v",resp)

	responseBody,err := ioutil.ReadAll(resp.Body)

	vault.l.Printf("Response from vault: %+v ",string(responseBody))
	
	if err != nil {
		vault.l.Println("[ERROR] Could not send request! Server connection issue ", err)
		http.Error(rw, "Error Unbale to send Vault Server Request ", http.StatusBadGateway)
		return
	}
	
	vault.l.Println("The Status Response ==>> ", resp.StatusCode)

	if resp.StatusCode != 204 {
		vault.l.Println("[ERROR] Non 200 Status Code ", err)
		http.Error(rw, "Error Non 200 Status Code ", http.StatusBadGateway)
		return
	}

}
