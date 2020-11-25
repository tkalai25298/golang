package vaultinterface

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vault-msp/data"
	"github.com/vault-msp/httpreq"
)

//Cert for vaultCompleteInterface httpcall()
type Cert struct {
	Data    data.Cert
	Request httpreq.HTTPClient
}

//NewCert - Creating new IssueCert object
func NewCert(request httpreq.HTTPClient) *Cert {
	return &Cert{Request: request}
}

//IssueCert to create cert for vaultComplete Interface
func (cert *Cert) IssueCert() *Errors {

	paths := [2]string{"CA","TLSCA"}

	err := cert.Data.Validate()

	if err != nil {
		return &Errors{ Message: fmt.Sprintf("Error Request Json validation: %v",err ), Status: http.StatusBadRequest}
	}

	for _,path := range paths{

		pkiPath := cert.Data.Data.Organization + path
		vaultData, err := json.Marshal(cert.Data.Data)

		//Sending http request to vault server
		resp, err := cert.Request.HTTPCall("/v1/"+pkiPath+"/issue/"+cert.Data.Roles, vaultData)

		if err != nil {
			return &Errors{Message: fmt.Sprintf("Error Unbale to send Vault Server Request: %v",err), Status: http.StatusBadGateway}
		}

		if resp.StatusCode != 200 {
			return &Errors{Message: fmt.Sprintf("Error Non 200 Status Code for Issue cert:got %v",resp.StatusCode), Status: http.StatusBadGateway}
		}
	}

	return nil
}
