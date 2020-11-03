package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/vault-msp/httpreq"
	config"github.com/vault-msp/config"
	data"github.com/vault-msp/data"  //Role struct
)


//CreateRole handler to create a role for issuing certificates
func CreateRole(rw http.ResponseWriter, r *http.Request) {

	role := data.Role{}

	config,err := config.SetConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

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

	reqObj := httpreq.CreateRequest("POST","http://"+config.VaultURL+"v1/"+role.Path+"/roles/"+role.Roles,config.VaultToken,vaultData)
		resp, err := reqObj.HTTPCall()

		if err != nil {
			log.Fatal("could not send request! Server connection issue")
		}
		log.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 204 {
			log.Panicf("NON 204 STATUS CODE")
		}

		defer resp.Body.Close()
}