package vaultinterface

import (
	"encoding/json"
	"net/http"
	"github.com/vault-msp/httpreq"
	"github.com/vault-msp/data"
)

//PKI for vaultCompleteInterface httpcall()
type PKI struct{
	Data data.Pki
	Request httpreq.HTTPClient
}

//NewPKI - Creating new PKI object
func NewPKI( request httpreq.HTTPClient) *PKI{
	return &PKI{Request: request}
}

//EnablePki to enable pki engine for vaultComplete Interface
func (p *PKI) EnablePki() *Errors {

	err := p.Data.Validate()

	if err != nil {
		return &Errors{ Message: "Error Request Json validation ", Status: http.StatusBadRequest}
	}

	vaultData, err := json.Marshal(p.Data.Data)

	//Sending http request to vault server
	resp, err := p.Request.HTTPCall("/v1/sys/mounts/"+p.Data.Path,vaultData)

	if err != nil {
		return &Errors{ Message: "Error Unbale to send Vault Server Request ", Status: http.StatusBadGateway}
	}

	if resp.StatusCode != 204 {
		return &Errors{ Message: "Error Non 200 Status Code ", Status: http.StatusBadGateway}
	}
	return nil
}