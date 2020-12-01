package handlers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	data"golang/vault-msp/data"  //CA struct
)


//IssueCA handler to issue root certificate 
func (vault *Vault) IssueCA (rw http.ResponseWriter,req *http.Request) {

	paths := [2]string{"CA","TLSCA"}

	defer req.Body.Close()
	ca := data.RootCAData{}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		vault.l.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody,&ca)

	if err != nil {
		vault.l.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	err = ca.Validate()

	if err != nil {
		vault.l.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}
	
	for _,path := range paths{

	pkiPath := ca.Organization + path	
	vaultData,err := json.Marshal(ca)
	if err != nil{
		vault.l.Println("[ERROR] Marshalling ca data ", err)
		http.Error(rw, "Error Marshalling ca data ", http.StatusBadGateway)
		return
	}

	resp, err := vault.requestObject.HTTPCall("/v1/"+pkiPath+"/root/generate/internal",vaultData)

		if err != nil {
			vault.l.Println("[ERROR] Could not send request! Server connection issue ", err)
			http.Error(rw, "Error Unbale to send Vault Server Request ", http.StatusBadGateway)
			return
		}

		vault.l.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 200 {
			vault.l.Println("[ERROR] Non 200 Status Code ", err)
			http.Error(rw, "Error Non 200 Status Code ", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
	}

	var data = Response{Response: "RootCA created! "}
	rw.Header().Set("Content-Type", "application/json")
	err = data.JSONResponse(rw)
	
	if err != nil {
		vault.l.Println("[ERROR] Could not Marshal response json ", err)
		http.Error(rw, "Error Unbale to marshal response json ", http.StatusBadGateway)
		return
	}

}