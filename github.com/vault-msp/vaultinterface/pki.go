package vaultinterface

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/vault-msp/httpreq"
	"github.com/vault-msp/data"
)

//PKI for vaultCompleteInterface httppkill()
type PKI struct{
	Data data.Pki
	Request httpreq.HTTPClient
}

//NewPKI - Creating new PKI object
func NewPKI( request httpreq.HTTPClient) *PKI{
	return &PKI{Request: request}
}

//EnablePKI to create PKI cert for vaultComplete Interface
func (pki *PKI) EnablePKI() *Errors{
	
	err := pki.Data.Validate()

	if err != nil {
		return &Errors{ Message: "Error Request Json validation ", Status: http.StatusBadRequest}
	}

	vaultData, err := json.Marshal(pki.Data.Data)
	
	//Sending http request to vault server
	resp, err := pki.Request.HTTPCall("/v1/sys/mounts/"+pki.Data.Path,vaultData)

	if err != nil {
		log.Println(err)
		return &Errors{ Message: "Error Unbale to send Vault Server Request ", Status: http.StatusBadGateway}
	}

	if resp.StatusCode != 204 {
		return &Errors{ Message: "Error Non 204 Status Code for pki ", Status: http.StatusBadGateway}
	}
	return nil
}