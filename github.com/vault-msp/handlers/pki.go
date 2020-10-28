package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"

	"github.com/vault-msp/httpreq"
)

//Pki struct for request params body
type Pki struct {
	Path string `json:"path"`
	Data PkiData `json:"data"`
}

//PkiData struct for request params with data to be passed for vault
type PkiData struct {
	Type string `json:"type"`
	Config Config `json:"config"`
	SealWrap bool `json:"seal_wrap"`
}
//Config struct to be passed for Data struct
type Config struct {
	MaxLeaseTTL string `json:"max_lease_ttl"`
	DefaultLeaseTTL string `json:"default_lease_ttl"`
}


//EnablePKI handler to create a pki engine to store certs
func EnablePKI(rw http.ResponseWriter,r *http.Request){
	pki := Pki{}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}
	err = json.Unmarshal(reqBody,&pki)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	// log.Printf("Received: %+v\n", pki.Data)

	vaultData,err := json.Marshal(pki.Data)
	log.Printf("%v",vaultData)

		reqObj := httpreq.CreateRequest("POST","http://localhost:8200/v1/sys/mounts/"+pki.Path,"myroot",vaultData)
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
