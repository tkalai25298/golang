package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/vault-msp/httpreq"
	config"github.com/vault-msp/config"
	data"github.com/vault-msp/data"  //CA struct
)


//IssueCA handler to issue root certificate 
func IssueCA (w http.ResponseWriter,r *http.Request) {

	ca := data.RootCA{}

	config,err := config.SetConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

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

	reqObj := httpreq.CreateRequest("POST","http://"+config.VaultURL+"/v1/"+ca.Path+"/root/generate/internal",config.VaultToken,vaultData)
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