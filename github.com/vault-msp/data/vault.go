package data


//VaultComplete request body to perform create pki,ca,role,issue cert
type VaultComplete struct {
	PKI   Pki    `json:"pki"`
	CA    RootCA `json:"ca"`
	Roles Role   `json:"roles"`
	Certs Cert   `json:"cert"`
}