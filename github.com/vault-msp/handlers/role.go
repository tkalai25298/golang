package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	// "github.com/vault-msp/httpreq"
	// config"github.com/vault-msp/config"
	data"github.com/vault-msp/data"  //Role struct
)


//CreateRole handler to create a role for issuing certificates
func (enable *Enable) CreateRole(rw http.ResponseWriter, r *http.Request) {

	role := data.Role{}


	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}

	err = json.Unmarshal(reqBody,&role)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	err = role.Validate()

	if err != nil {
		log.Fatal("json validation error",err)
	}

	vaultData,err := json.Marshal(role.Data)

	resp, err := enable.requestObject.HTTPCall("/v1/"+role.Path+"/roles/"+role.Roles,vaultData)


		if err != nil {
			log.Fatal("could not send request! Server connection issue")
		}
		log.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 204 {
			log.Panicf("NON 204 STATUS CODE")
		}

		defer resp.Body.Close()
}