package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"

	"github.com/vault-msp/httpreq"
	config"github.com/vault-msp/config"
	data"github.com/vault-msp/data"  //Issue struct
)


//IssueCert handler to issue certs by a role
func IssueCert(rw http.ResponseWriter,r *http.Request) {

	cert := data.Cert{}

	config,err := config.SetConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}
	
	err = json.Unmarshal(reqBody,&cert)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	// err = cert.Validate()

	if err != nil {
		log.Fatal("json validation error",err)
	}

	vaultData,err := json.Marshal(cert.Data)

	reqObj := httpreq.CreateRequest("POST","http://"+config.VaultURL+"v1/"+cert.Path+"/issue/"+cert.Roles,config.VaultToken,vaultData)
		resp, err := reqObj.HTTPCall()

		if err != nil {
			log.Fatal("could not send request! Server connection issue")
		}
		log.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 200 {
			log.Panicf("NON 200 STATUS CODE")
		}

		defer resp.Body.Close()
}