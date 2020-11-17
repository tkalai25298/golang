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


//SplitRequest to split corresponding request objects for fn calls
func SplitRequest(vault *data.VaultComplete,requestObj httpreq.HTTPClient) *VaultInterface{
	vaultInterface := VaultInterface{}

	//Pki Data
	vaultInterface.Pki.Data = data.Pki{
		Path : vault.PKI.Path,
		Data : data.PkiData{
			Type : vault.PKI.Data.Type,
			Config : data.Config{
				MaxLeaseTTL : vault.PKI.Data.Config.MaxLeaseTTL,
				DefaultLeaseTTL : vault.PKI.Data.Config.DefaultLeaseTTL,
			},
			SealWrap : vault.PKI.Data.SealWrap,
		},
	}
	vaultInterface.Pki.Request = requestObj
 
	//CA Data
	vaultInterface.CA.Data = data.RootCA{
		Path : vault.CA.Path,
		Data : data.CAData{
			CommonName : vault.CA.Data.CommonName,
			TTL : vault.CA.Data.TTL,
			KeyBits : vault.CA.Data.KeyBits,
			KeyType : vault.CA.Data.KeyType,
			Organization : vault.CA.Data.Organization,
		},
	}
	vaultInterface.CA.Request = requestObj

	//Role Data
	vaultInterface.Roles.Data = data.Role{
		Path : vault.Roles.Path,
		Roles : vault.Roles.Roles,
		Data : data.RoleData{
			ServerFlag : vault.Roles.Data.ServerFlag,
			ClientFlag : vault.Roles.Data.ClientFlag,
			KeyBits : vault.Roles.Data.KeyBits,
			KeyType : vault.Roles.Data.KeyType,
			KeyUsage : vault.Roles.Data.KeyUsage,
			MaxTTL : vault.Roles.Data.MaxTTL,
			GenerateLease : vault.Roles.Data.GenerateLease,
			AllowAnyName : vault.Roles.Data.AllowAnyName,
			OU : vault.Roles.Data.OU,
			Organization : vault.Roles.Data.Organization,
			AllowedDomains : vault.Roles.Data.AllowedDomains,
			AllowSubdomains : vault.Roles.Data.AllowSubdomains,
			BasicConstraintsValidForNonCA : vault.Roles.Data.BasicConstraintsValidForNonCA,	
		},
	}
	vaultInterface.Roles.Request = requestObj

	//IssueCert Data
	vaultInterface.Cert.Data = data.Cert{
		Path : vault.Certs.Path,
		Roles : vault.Certs.Roles,
		Data : data.IssueCertData{
			CommonName : vault.Certs.Data.CommonName,
			TTL : vault.Certs.Data.TTL,
			AltNames : vault.Certs.Data.AltNames,
		},
	}
	vaultInterface.Cert.Request = requestObj

	return &vaultInterface
}