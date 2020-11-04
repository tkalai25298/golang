package data

import (
	"github.com/go-playground/validator"
)

//Pki struct for request params body
type Pki struct {
	Path string `json:"path" validate:"required"`
	Data PkiData `json:"data"`
}

//PkiData struct for request params with data to be passed for vault
type PkiData struct {
	Type   string `json:"type"`

	Config struct {
		MaxLeaseTTL     string `json:"max_lease_ttl"`
		DefaultLeaseTTL string `json:"default_lease_ttl"`
	} `json:"config"`

	SealWrap bool `json:"seal_wrap"`
}


//Validate for Pki struct json validation
func (pki *Pki) Validate() error {

	pki.SetDefaultValues()

	validate := validator.New()
	return validate.Struct(pki)
}


//SetDefaultValues to assign missing values to be passed for vault server
func (pki *Pki) SetDefaultValues()  {

	if pki.Data.Type == "" {
		pki.Data.Type = "pki"
	}
	
	if pki.Data.Config.MaxLeaseTTL == "" {
		pki.Data.Config.MaxLeaseTTL = "87600h"
	}

	if pki.Data.SealWrap == false {
		pki.Data.SealWrap = true
	}
}




//Validate for Pki struct json validation
// func (pki *Pki) Validate() error {
// 	validate := validator.New()
// 	validate.RegisterValidation("type", SetType)

	

// 	return validate.Struct(pki)
// }

//SetType to set missed default values to be passed for Vault
// func SetType(fl validator.FieldLevel) bool {

// 	if fl.Field().IsZero() {
		
// 		return true
// 	}
// 	return false
// }