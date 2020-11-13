package vaultinterface

import (
	"encoding/json"
	"net/http"
	"github.com/vault-msp/httpreq"
	"github.com/vault-msp/data"
)

//Role for vaultCompleteInterface httpcall()
type Role struct{
	Data data.Role
	Request httpreq.HTTPClient
}

//NewRole - Creating new Role object
func NewRole( request httpreq.HTTPClient) *Role{
	return &Role{Request: request}
}

//CreateRoles to create RootCA cert for vaultComplete Interface
func (role *Role) CreateRoles() *Errors{
	
	err := role.Data.Validate()

	if err != nil {
		return &Errors{ Message: "Error Request Json validation ", Status: http.StatusBadRequest}
	}

	vaultData, err := json.Marshal(role.Data.Data)

	//Sending http request to vault server
	resp, err := role.Request.HTTPCall("/v1/"+role.Data.Path+"/roles/"+role.Data.Roles,vaultData)

	if err != nil {
		return &Errors{ Message: "Error Unbale to send Vault Server Request ", Status: http.StatusBadGateway}
	}

	if resp.StatusCode != 204 {
		return &Errors{ Message: "Error Non 200 Status Code ", Status: http.StatusBadGateway}
	}
	return nil
}