package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vault-msp/data"
	"github.com/vault-msp/helpers"
)

//VaultInterface Handler to create pki,create,role,issue cert
func (vault *Vault) VaultInterface(rw http.ResponseWriter,req *http.Request) {
	defer req.Body.Close()

	reqData := data.VaultComplete{}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		vault.l.Println("[ERROR] Reading request body: ", err)
		http.Error(rw, "Error Reading Request body", http.StatusBadRequest)
		return
	}
	
	err = json.Unmarshal(reqBody, &reqData)
	if err != nil {
		vault.l.Println("[ERROR] Decoding Request body:  ", err)
		http.Error(rw, "Error Decoding Request body  ", http.StatusBadRequest)
		return
	}


	//seperating the request object to Pki,ca,role,issue request
	vaultInterface := helpers.AddRequestObject(&reqData,vault.requestObject)

	vault.l.Println("===>>>Creating Pki Engine...")
	executeErr := vaultInterface.Pki.EnablePKI()
	if executeErr != nil {
		vault.l.Println("[ERROR] in PKI: ",executeErr)
		http.Error(rw,executeErr.Message,executeErr.Status)
		return
	}

	vault.l.Println("===>>>Creating RootCA cert...")
	executeErr = vaultInterface.CA.IssueRootCA()
	if executeErr != nil {
		vault.l.Println("[ERROR] in CA: ",executeErr)
		http.Error(rw,executeErr.Message,executeErr.Status)
		return
	}

	vault.l.Println("===>>>Creating Roles to issue the certs...")
	executeErr = vaultInterface.Roles.CreateRoles()
	if executeErr != nil {
		vault.l.Println("[ERROR] in Roles: ",executeErr)
		http.Error(rw,executeErr.Message,executeErr.Status)
		return
	}

	vault.l.Println("===>>>Issuing the Certs...")
	executeErr = vaultInterface.Cert.IssueCert()
	if executeErr != nil {
		vault.l.Println("[ERROR] in Cert: ",executeErr)
		http.Error(rw,executeErr.Message,executeErr.Status)
		return
	}
	vault.l.Println("Completed!")

	fmt.Fprintf(rw,"Creation of PKI,Roles,RootCA,Issue certs Completed! ")

		
}