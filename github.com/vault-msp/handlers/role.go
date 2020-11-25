package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	data"github.com/vault-msp/data"  //Role struct
)


//CreateRole handler to create a role for issuing certificates
func (vault *Vault) CreateRole(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	role := data.Role{}


	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody,&role)

	if err != nil {
		log.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	err = role.Validate()

	if err != nil {
		log.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}

	pkiPath := role.Data.Organization + "CA"
	vaultData,err := json.Marshal(role.Data)

	for _,rolename := range role.Roles{

		resp, err := vault.requestObject.HTTPCall("/v1/"+pkiPath+"/roles/"+rolename,vaultData)

		if err != nil {
			log.Println("[ERROR] Could not send request! Server connection issue ", err)
			http.Error(rw, "Error Unbale to send Vault Server Request ", http.StatusBadGateway)
			return
		}

		log.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 204 {
			log.Println("[ERROR] Non 200 Status Code ", err)
			http.Error(rw, "Error Non 200 Status Code ", http.StatusBadGateway)
			return
		}

		defer resp.Body.Close()
	}
}