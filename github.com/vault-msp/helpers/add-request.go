package helpers

import (
	vaultinterface"github.com/vault-msp/vaultinterface"
	httpreq"github.com/vault-msp/httpreq"
	data"github.com/vault-msp/data"
)

//VaultInterface to create pki,ca,role,issue cert
type VaultInterface struct{
	Pki vaultinterface.PKI
	CA vaultinterface.RootCA
	Roles vaultinterface.Role
	Cert vaultinterface.Cert
}


//AddRequestObject to add request obj to the data from vaultInterface 
func AddRequestObject(vault *data.VaultComplete,requestObj httpreq.HTTPClient) *VaultInterface{
	vaultInterface := VaultInterface{}

	//Pki Data
	vaultInterface.Pki.Data = vault.PKI
	vaultInterface.Pki.Request = requestObj
 
	//CA Data
	vaultInterface.CA.Data = vault.CA
	vaultInterface.CA.Request = requestObj

	//Role Data
	vaultInterface.Roles.Data = vault.Roles
	vaultInterface.Roles.Request = requestObj

	//IssueCert Data
	vaultInterface.Cert.Data = vault.Certs
	vaultInterface.Cert.Request = requestObj

	return &vaultInterface
}