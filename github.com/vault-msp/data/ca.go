package data

import (
	"github.com/go-playground/validator"
)

//RootCA struct for request params body
type RootCA struct{
	Path string `json:"path" validate:"required"`
	Data CAData `json:"data"`
}

//CAData struct for vault data config to create root CA cert
type CAData struct {
	CommonName string `json:"common_name" validate:"required"`
	TTL string `json:"ttl"`
	KeyType string `json:"key_type"`
	KeyBits int `json:"key_bits"`
	Organization string `json:"organization" validate:"required"`
}


//Validate for RootCA cert struct json validation
func (ca *RootCA) Validate() error {
	ca.SetDefaultValues()

	validate := validator.New()
	return validate.Struct(ca)
}	

//SetDefaultValues to assign missing values to be passed for vault server
func (ca *RootCA) SetDefaultValues() {
	if ca.Data.TTL == "" {
		ca.Data.TTL = "87600h"
	}

	if ca.Data.KeyType == "" {
		ca.Data.KeyType = "ec"
	}

	if ca.Data.KeyBits == 0 {
		ca.Data.KeyBits = 256
	}
}




