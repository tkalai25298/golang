package data


//VaultComplete request body to perform create pki,ca,role,issue cert
type VaultComplete struct {
	PKI   PkiData    `json:"pki"`
	CA    RootCAData `json:"ca"`
	Roles Role   `json:"roles"`
	Certs Cert   `json:"cert"`
}

