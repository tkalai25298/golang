package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	data "github.com/vault-msp/data" //Issue struct
)

//Wallet handler to generate identity using blockchain-tx-api
func (vault *Vault) Wallet(rw http.ResponseWriter, req *http.Request) {

	paths := [2]string{"CA","TLSCA"}

	defer req.Body.Close()
	cert := data.Cert{}
	response := data.IssueCertResponse{}

	reqBody, err := ioutil.ReadAll(req.Body)
	
	if err != nil {
		vault.l.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}
	
	err = json.Unmarshal(reqBody, &cert)

	if err != nil {
		vault.l.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	err = cert.Validate()

	if err != nil {
		vault.l.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}

	for _,path := range paths{

		pkiPath := cert.Data.Organization + path
		vaultData, err := json.Marshal(cert.Data)
		//issue certs
		res, err := vault.requestObject.HTTPCall("/v1/"+pkiPath+"/issue/"+cert.Roles, vaultData)
		defer res.Body.Close()

		if err != nil {
			vault.l.Println("[ERROR] Bad request ", err)
			http.Error(rw, "Error Bad request ", http.StatusBadRequest)
			return
		}

		vault.l.Println("The Vault server Status Response ==>> ", res.StatusCode)

		if res.StatusCode != 200 {
			vault.l.Println("[ERROR] Non 200 Status Code from vault-server ", err)
			http.Error(rw, "Error Non 200 Status Code from vault-server ", http.StatusBadGateway)
			return
		}

		resp, err := ioutil.ReadAll(res.Body)

		if err != nil {
			vault.l.Println("[ERROR] Reading Vault Response body: ", err)
			http.Error(rw, "Error Reading Vault Response body", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(resp, &response)

		response.Data.Organization = cert.Data.Organization

		identityRequest,err := json.Marshal(response.Data)


		//generating identity
		result,err := http.Post("http://35.242.187.129:3000/Identity","application/json",bytes.NewBuffer(identityRequest))
		identityResult,err := ioutil.ReadAll(result.Body)

		vault.l.Printf("Identity response: %v ",string(identityResult));

		if err != nil {
			vault.l.Println("[ERROR] Could not send request! Server connection issue ", err)
			http.Error(rw, "Error Unbale to send Transaction API Server Request ", http.StatusBadGateway)
			return
		}
		vault.l.Println("The Identity Status Response ==>> ", result.StatusCode)

		if result.StatusCode != 200 {
			vault.l.Println("[ERROR] Non 200 Status Code in transaction-api", err)
			http.Error(rw, "Error Non 200 Status Code in transaction-api ", http.StatusBadGateway)
			return
		}
	}
	
	var data = Response{Response: "Identity generated ! "}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

}
