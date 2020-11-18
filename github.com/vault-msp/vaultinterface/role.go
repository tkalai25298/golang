package vaultinterface

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vault-msp/data"
	"github.com/vault-msp/httpreq"
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
		return &Errors{ Message: fmt.Sprintf("Error Request Json validation: %v",err ), Status: http.StatusBadRequest}
	}

	vaultData, err := json.Marshal(role.Data.Data)

	for _,rolename := range role.Data.Roles {

		//Sending http request to vault server
		resp, err := role.Request.HTTPCall("/v1/"+role.Data.Path+"/roles/"+rolename,vaultData)

		if err != nil {
			return &Errors{ Message: fmt.Sprintf("Error Unbale to send Vault Server Request :%v",err), Status: http.StatusBadGateway}
		}

		if resp.StatusCode != 204 {
			return &Errors{ Message: fmt.Sprintf("Error Non 200 Status Code creating the Role:got %v",resp.StatusCode), Status: http.StatusBadGateway}
		}
		
	}
	return nil
}