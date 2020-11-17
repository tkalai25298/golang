package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	data"github.com/vault-msp/data"  //CA struct
)


//IssueCA handler to issue root certificate 
func (vault *Vault) IssueCA (rw http.ResponseWriter,req *http.Request) {
	defer req.Body.Close()
	ca := data.RootCA{}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(reqBody,&ca)

	if err != nil {
		log.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	err = ca.Validate()

	if err != nil {
		log.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}


	vaultData,err := json.Marshal(ca.Data)

	resp, err := vault.requestObject.HTTPCall("/v1/"+ca.Path+"/root/generate/internal",vaultData)

		if err != nil {
			log.Println("[ERROR] Could not send request! Server connection issue ", err)
			http.Error(rw, "Error Unbale to send Vault Server Request ", http.StatusBadGateway)
			return
		}
		log.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 200 {
			log.Println("[ERROR] Non 200 Status Code ", err)
			http.Error(rw, "Error Non 200 Status Code ", http.StatusBadGateway)
			return
			
		}

		defer resp.Body.Close()
}