package vaultinterface

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vault-msp/data"
	"github.com/vault-msp/httpreq"
)

//PKI for vaultCompleteInterface httppkill()
type PKI struct{
	Data data.PkiData
	Request httpreq.HTTPClient
}

//NewPKI - Creating new PKI object
func NewPKI( request httpreq.HTTPClient) *PKI{
	return &PKI{Request: request}
}

//EnablePKI to create PKI cert for vaultComplete Interface
func (pki *PKI) EnablePKI() *Errors{
	
	paths := [2]string{"CA","TLSCA"}
	err := pki.Data.Validate()

	if err != nil {
		return &Errors{ Message: fmt.Sprintf("Error Request Json validation: %v",err ), Status: http.StatusBadRequest}
	}

	for _,path := range paths{

		pkiPath := pki.Data.Organization + path
		vaultData, err := json.Marshal(pki.Data)
		
		//Sending http request to vault server
		resp, err := pki.Request.HTTPCall("/v1/sys/mounts/"+pkiPath,vaultData)

		if err != nil {
			log.Println(err)
			return &Errors{ Message: fmt.Sprintf("Error Unbale to send Vault Server Request :%v",err), Status: http.StatusBadGateway}
		}

		if resp.StatusCode != 204 {
			return &Errors{ Message: fmt.Sprintf("Error Non 204 Status Code for pki: got %v",resp.StatusCode), Status: http.StatusBadGateway}
		}
	}
	return nil
}