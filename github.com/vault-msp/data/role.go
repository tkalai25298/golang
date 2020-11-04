package data

import (
	"github.com/go-playground/validator"
)


//Role struct for creating role request obj
type Role struct{
	Path string `json:"path" validate:"required"`
	Roles string `json:"roles" validate:"required"`
	Data RoleData `json:"data"`
}
//RoleData struct for vault config data to create role
type RoleData struct{
	ServerFlag bool `json:"server_flag"`
	ClientFlag bool `json:"client_flag"`
	KeyType string `json:"key_type"`
	KeyBits int `json:"key_bits"`
	KeyUsage []string `json:"key_usage"`
	MaxTTL string `json:"max_ttl"` 
	GenerateLease bool `json:"generate_lease"`
	AllowAnyName bool `json:"allow_any_name" validate:"required"`
	OU string `json:"ou" validate:"required"`
	Organization string `json:"organization" validate:"required"`
	AllowedDomains string `json:"allowed_domains" validate:"required"`
	AllowSubdomains bool `json:"allow_subdomains" validate:"required"`
	BasicConstraintsValidForNonCA bool `json:"basic_constraints_valid_for_non_ca"`
}



//Validate for Role struct json validation
func (role *Role) Validate() error {
	role.SetDefaultValues()

	validate := validator.New()
	return validate.Struct(role)
}	

//SetDefaultValues to assign missing values to be passed for vault server
func (role *Role) SetDefaultValues() {
	if role.Data.KeyType == "" {
		role.Data.KeyType = "ec"
	}	

	if role.Data.KeyBits == 0 {
		role.Data.KeyBits = 256
	}

	if len(role.Data.KeyUsage) == 0 {
		role.Data.KeyUsage = []string{"DigitalSignature"}
	}

	if role.Data.MaxTTL == "" {
		role.Data.MaxTTL = "3000h"
	}

	if role.Data.GenerateLease == false {
		role.Data.GenerateLease = true
	}

	if role.Data.BasicConstraintsValidForNonCA == false {
		role.Data.BasicConstraintsValidForNonCA = true
	}

}