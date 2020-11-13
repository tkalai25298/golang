package vaultinterface

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/vault-msp/httpreq"
	"github.com/vault-msp/data"
)

//RootCA for vaultCompleteInterface httpcall()
type RootCA struct{
	Data data.RootCA
	Request httpreq.HTTPClient
}

//NewRootCA - Creating new RootCA object
func NewRootCA( request httpreq.HTTPClient) *RootCA{
	return &RootCA{Request: request}
}

//IssueRootCA to create RootCA cert for vaultComplete Interface
func (ca *RootCA) IssueRootCA() *Errors{
	
	err := ca.Data.Validate()

	if err != nil {
		return &Errors{ Message: "Error Request Json validation ", Status: http.StatusBadRequest}
	}

	vaultData, err := json.Marshal(ca.Data.Data)

	
	//Sending http request to vault server
	resp, err := ca.Request.HTTPCall("/v1/"+ca.Data.Path+"/root/generate/internal",vaultData)

	if err != nil {
		log.Println(err)
		return &Errors{ Message: "Error Unbale to send Vault Server Request ", Status: http.StatusBadGateway}
	}

	if resp.StatusCode != 200 {
		return &Errors{ Message: "Error Non 200 Status Code ", Status: http.StatusBadGateway}
	}
	return nil
}