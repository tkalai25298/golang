package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"

	"github.com/vault-msp/httpreq"
)

//Cert for req Params obj
type Cert struct{
	Path string `json:"path"`
	Roles string `json:"roles"`
	Data IssueCertData `json:"data"`
}

//IssueCertData to pass vault data config to issue certificates by a role
type IssueCertData struct {
	CommonName string `json:"common_name"`
	TTL string `json:"ttl"`
	AltNames string `json:"alt_names"`
}

//IssueCert handler to issue certs by a role
func IssueCert(rw http.ResponseWriter,r *http.Request) {

	cert := Cert{}
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}
	
	err = json.Unmarshal(reqBody,&cert)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	vaultData,err := json.Marshal(cert.Data)

	reqObj := httpreq.CreateRequest("POST","http://localhost:8200/v1/"+cert.Path+"/issue/"+cert.Roles,"myroot",vaultData)
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