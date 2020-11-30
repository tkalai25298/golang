package handlers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	data"github.com/vault-msp/data"  //Role struct
)


//CreateRole handler to create a role for issuing certificates
func (vault *Vault) CreateRole(rw http.ResponseWriter, req *http.Request) {

	paths := [2]string{"CA","TLSCA"}

	defer req.Body.Close()
	role := data.Role{}


	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		vault.l.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody,&role)

	if err != nil {
		vault.l.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	err = role.Validate()

	if err != nil {
		vault.l.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}

	for _,path := range paths{

		pkiPath := role.Data.Organization + path
		vaultData,err := json.Marshal(role.Data)

		if err != nil{
			vault.l.Println("[ERROR] Marshalling role data ", err)
			http.Error(rw, "Error Marshalling role data ", http.StatusBadGateway)
			return
		}

		for _,rolename := range role.Roles{

			resp, err := vault.requestObject.HTTPCall("/v1/"+pkiPath+"/roles/"+rolename,vaultData)

			if err != nil {
				vault.l.Println("[ERROR] Could not send request! Server connection issue ", err)
				http.Error(rw, "Error Unbale to send Vault Server Request ", http.StatusBadGateway)
				return
			}

			vault.l.Println("The Status Response ==>> ",resp.StatusCode)

			if resp.StatusCode != 204 {
				vault.l.Println("[ERROR] Non 200 Status Code ", err)
				http.Error(rw, "Error Non 200 Status Code ", http.StatusBadGateway)
				return
			}

			defer resp.Body.Close()
		}
	}
	var data = Response{Response: "Roles created! "}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}