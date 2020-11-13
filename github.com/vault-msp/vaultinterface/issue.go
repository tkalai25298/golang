package vaultinterface

import (
	"encoding/json"
	"net/http"

	"github.com/vault-msp/httpreq"
	"github.com/vault-msp/data"
)

//Cert for vaultCompleteInterface httpcall()
type Cert struct{
	Data data.Cert
	Request httpreq.HTTPClient
}

//NewCert - Creating new IssueCert object
func NewCert(request httpreq.HTTPClient) *Cert{
	return &Cert{Request: request}
}

//IssueCert to create cert for vaultComplete Interface
func (cert *Cert) IssueCert() *Errors{
	
	err := cert.Data.Validate()

	if err != nil {
		return &Errors{ Message: "Error Request Json validation ", Status: http.StatusBadRequest}
	}

	vaultData, err := json.Marshal(cert.Data.Data)

	//Sending http request to vault server
	resp, err := cert.Request.HTTPCall("/v1/"+cert.Data.Path+"/issue/"+cert.Data.Roles,vaultData)

	if err != nil {
		return &Errors{ Message: "Error Unbale to send Vault Server Request ", Status: http.StatusBadGateway}
	}

	if resp.StatusCode != 200 {
		return &Errors{ Message: "Error Non 200 Status Code ", Status: http.StatusBadGateway}
	}
	return nil
}