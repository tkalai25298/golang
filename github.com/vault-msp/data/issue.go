package data

import (
	"github.com/go-playground/validator"
)

//Cert for req Params obj
type Cert struct{
	Path string `json:"path" validate:"required"`
	Roles string `json:"roles" validate:"required"`
	Data IssueCertData `json:"data"`
}

//IssueCertData to pass vault data config to issue certificates by a role
type IssueCertData struct {
	CommonName string `json:"common_name" validate:"required"`
	TTL string `json:"ttl"`
	AltNames string `json:"alt_names"`
}


//Validate for Role struct json validation
func (cert *Cert) Validate() error {
	cert.SetDefaultValues()

	validate := validator.New()
	return validate.Struct(cert)
}	

//SetDefaultValues to assign missing values to be passed for vault server
func (cert *Cert) SetDefaultValues() {
	if cert.Data.TTL == "" {
		cert.Data.TTL = "2400h"
	}

}