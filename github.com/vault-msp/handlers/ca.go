package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	// "github.com/vault-msp/httpreq"
	data"github.com/vault-msp/data"  //CA struct
)


//IssueCA handler to issue root certificate 
func (enable *Enable) IssueCA (w http.ResponseWriter,r *http.Request) {

	ca := data.RootCA{}


	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}
	err = json.Unmarshal(reqBody,&ca)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	err = ca.Validate()

	if err != nil {
		log.Fatal("json validation error",err)
	}


	vaultData,err := json.Marshal(ca.Data)

	resp, err := enable.requestObject.HTTPCall("/v1/"+ca.Path+"/root/generate/internal",vaultData)

		if err != nil {
			log.Fatal("could not send request! Server connection issue")
		}
		log.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 200 {
			log.Panicf("NON 200 STATUS CODE")
		}

		defer resp.Body.Close()
}