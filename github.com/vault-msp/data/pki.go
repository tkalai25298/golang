package data

import (
	"github.com/go-playground/validator"
)

//PkiData struct for request params with data to be passed for vault
type PkiData struct {
	Organization string `json:"organization" validate:"required"`
	Type   string `json:"type"`
	Config Config `json:"config"`
	SealWrap bool `json:"seal_wrap"`
}

//Config for Pki vault config
type Config struct {
	MaxLeaseTTL     string `json:"max_lease_ttl"`
	DefaultLeaseTTL string `json:"default_lease_ttl"`
}


//Validate for Pki struct json validation
func (pki *PkiData) Validate() error {

	pki.SetDefaultValues()

	validate := validator.New()
	return validate.Struct(pki)
}


//SetDefaultValues to assign missing values to be passed for vault server
func (pki *PkiData) SetDefaultValues()  {

	if pki.Type == "" {
		pki.Type = "pki"
	}
	
	if pki.Config.MaxLeaseTTL == "" {
		pki.Config.MaxLeaseTTL = "87600h"
	}

	if pki.SealWrap == false {
		pki.SealWrap = true
	}
}

