package vaultinterface

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"io"
	"github.com/vault-msp/httpreq"
	"github.com/vault-msp/data"
)

type PKI struct{
	Data data.Pki
	Request httpreq.HTTPClient
}

func NewPKI( request httpreq.HTTPClient) *PKI{
	return &PKI{Request: request}
}

func (p *PKI) ExecuteRequest( body io.ReadCloser) *Errors{
	defer body.Close()
	reqBody, err := ioutil.ReadAll(body)

	if err != nil {
		return &Errors{ Message: "Error Reading Request body", Status: http.StatusBadRequest}
	}

	err = json.Unmarshal(reqBody, &p.Data)

	if err != nil {
		return &Errors{ Message: "Error Decoding Request body  ", Status: http.StatusBadRequest}
	}

	
	err = p.Data.Validate()

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