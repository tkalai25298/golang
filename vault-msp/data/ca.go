package data

import (
	"github.com/go-playground/validator"
)

//RootCAData struct for request params body
type RootCAData struct{
	Organization string `json:"organization" validate:"required"`
	CommonName string `json:"common_name"`
	TTL string `json:"ttl"`
	KeyType string `json:"key_type"`
	KeyBits int `json:"key_bits"`
}

//Validate for RootCA cert struct json validation
func (ca *RootCAData) Validate() error {
	ca.SetDefaultValues()

	validate := validator.New()
	return validate.Struct(ca)
}	

//SetDefaultValues to assign missing values to be passed for vault server
func (ca *RootCAData) SetDefaultValues() {
	if ca.TTL == "" {
		ca.TTL = "87600h"
	}

	if ca.KeyType == "" {
		ca.KeyType = "ec"
	}

	if ca.KeyBits == 0 {
		ca.KeyBits = 256
	}
	
	if ca.CommonName == "" {
		ca.CommonName = ca.Organization + "CA"
	}
}




