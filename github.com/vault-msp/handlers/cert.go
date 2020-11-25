package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"

	data"github.com/vault-msp/data"  //Issue struct
)


//IssueCert handler to issue certs by a role
func (vault *Vault) IssueCert(rw http.ResponseWriter,req *http.Request) {
	defer req.Body.Close()
	cert := data.Cert{}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}
	
	err = json.Unmarshal(reqBody,&cert)

	if err != nil {
		log.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	err = cert.Validate()

	if err != nil {
		log.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}

	pkiPath := cert.Data.Organization+"CA"
	vaultData,err := json.Marshal(cert.Data)


	resp, err := vault.requestObject.HTTPCall("/v1/"+pkiPath+"/issue/"+cert.Roles,vaultData)
	defer resp.Body.Close()


		if err != nil {
			log.Println("[ERROR] Could not send request! Server connection issue ", err)
			http.Error(rw, "Error Unbale to send Vault Server Request ", http.StatusBadGateway)
			return
		}
		log.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 200 {
			log.Println("[ERROR] Non 200 Status Code ", err)
			http.Error(rw, "Error Non 200 Status Code ", http.StatusBadGateway)
			return
		}

	
	
}