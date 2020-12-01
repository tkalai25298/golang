package handlers

import (
	"io/ioutil"
	"net/http"
	"encoding/json"

	data"github.com/vault-msp/data"  //Issue struct
)


//IssueCert handler to issue certs by a role
func (vault *Vault) IssueCert(rw http.ResponseWriter,req *http.Request) {

	paths := [2]string{"CA","TLSCA"}

	defer req.Body.Close()
	cert := data.Cert{}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		vault.l.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}
	
	err = json.Unmarshal(reqBody,&cert)

	if err != nil {
		vault.l.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}

	// vault.l.Println("req params",cert)

	err = cert.Validate()

	if err != nil {
		vault.l.Println("[ERROR] Request Json validation  ", err)
		http.Error(rw, "Error Request Json validation ", http.StatusBadRequest)
		return
	}

	for _,path := range paths{

		pkiPath := cert.Data.Organization + path
		vaultData,err := json.Marshal(cert.Data)

		if err != nil{
			vault.l.Println("[ERROR] Marshalling cert data ", err)
			http.Error(rw, "Error Marshalling cert data ", http.StatusBadGateway)
			return
		}

		resp, err := vault.requestObject.HTTPCall("/v1/"+pkiPath+"/issue/"+cert.Roles,vaultData)
		
		
		// Body, err := ioutil.ReadAll(resp.Body)
		// vault.l.Println("response: ",string(Body))

		if err != nil {
			vault.l.Println("[ERROR] Could not send request! Server connection issue ", err)
			http.Error(rw, "Error Unbale to send Vault Server Request ", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		vault.l.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 200 {
			vault.l.Println("[ERROR] Non 200 Status Code ", err)
			http.Error(rw, "Error Non 200 Status Code ", http.StatusBadGateway)
			return
		}
	}

	var data = Response{Response: "Certs issued! "}
	rw.Header().Set("Content-Type", "application/json")
	err = data.JSONResponse(rw)
	
	if err != nil {
		vault.l.Println("[ERROR] Could not Marshal response json ", err)
		http.Error(rw, "Error Unbale to marshal response json ", http.StatusBadGateway)
		return
	}
}