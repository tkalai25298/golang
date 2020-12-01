package vaultinterface

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang/vault-msp/data"
	"golang/vault-msp/httpreq"
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

	paths := [2]string{"CA","TLSCA"}
	
	err := role.Data.Validate()

	if err != nil {
		return &Errors{ Message: fmt.Sprintf("Error Request Json validation: %v",err ), Status: http.StatusBadRequest}
	}

	for _,path := range paths{
	
		pkiPath := role.Data.Data.Organization + path
		
		if path == "TLSCA" {
			role.Data.Data.ClientFlag = true
			role.Data.Data.ServerFlag = true
		} else {
			role.Data.Data.ClientFlag = false
			role.Data.Data.ServerFlag = false
		}

		vaultData, err := json.Marshal(role.Data.Data)

		if err != nil{
			return &Errors{ Message: fmt.Sprintf("Error Unable to marshal role data :%v",err), Status: http.StatusBadGateway}
		}

		for _,rolename := range role.Data.Roles {

			//Sending http request to vault server
			resp, err := role.Request.HTTPCall("/v1/"+pkiPath+"/roles/"+rolename,vaultData)

			if err != nil {
				return &Errors{ Message: fmt.Sprintf("Error Unbale to send Vault Server Request :%v",err), Status: http.StatusBadGateway}
			}

			if resp.StatusCode != 204 {
				return &Errors{ Message: fmt.Sprintf("Error Non 200 Status Code creating the Role:got %v",resp.StatusCode), Status: http.StatusBadGateway}
			}
			
		}
	}
	return nil
}