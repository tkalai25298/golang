package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/vault-msp/httpreq"
)

//RootCA struct for request params body
type RootCA struct{
	Path string `json:"path"`
	Data CAData `json:"data"`
}

//CAData struct for vault data config to create root CA cert
type CAData struct {
	CommonName string `json:"common_name"`
	TTL string `json:"ttl"`
	KeyType string `json:"key_type"`
	KeyBits int `json:"key_bits"`
	Organization string `json:"organization"`
}

//IssueCA handler to issue root certificate 
func IssueCA (w http.ResponseWriter,r *http.Request) {

	ca := RootCA{}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}
	err = json.Unmarshal(reqBody,&ca)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	vaultData,err := json.Marshal(ca.Data)

	reqObj := httpreq.CreateRequest("POST","http://localhost:8200/v1/"+ca.Path+"/root/generate/internal","myroot",vaultData)
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