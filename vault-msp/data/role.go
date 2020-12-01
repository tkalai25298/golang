package data

import (
	"github.com/go-playground/validator"
)

//Role for creating a role request obj
type Role struct{
	Roles []string `json:"roles" validate:"required" `
	Data RoleData `json:"data"`
}


//RoleData struct for vault config data to create role
type RoleData struct {
	Organization string `json:"organization" validate:"required"`
	ServerFlag bool `json:"server_flag"`
	ClientFlag bool `json:"client_flag"`
	KeyType string `json:"key_type"`
	KeyBits int `json:"key_bits"`
	KeyUsage []string `json:"key_usage"`
	MaxTTL string `json:"max_ttl"` 
	GenerateLease bool `json:"generate_lease"`
	AllowAnyName bool `json:"allow_any_name" `
	OU string `json:"ou"`
	AllowLocalhost string `json:"allow_localhost" `
	AllowedDomains string `json:"allowed_domains"`
	AllowSubdomains bool `json:"allow_subdomains"`
	BasicConstraintsValidForNonCA bool `json:"basic_constraints_valid_for_non_ca"`
}



//Validate for Role struct json validation
func (role *Role) Validate() error {
	role.SetDefaultValues()

	validate := validator.New()

	for _,roles := range role.Roles {
	//validating the role names to be one of the 4types
	err := validate.Var(roles,"oneof=client admin orderer peer")
	if err != nil {
		return err
	}
	}
	
	//Validating the Role struct 
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

	if !role.Data.GenerateLease {
		role.Data.GenerateLease = true
	}
	if !role.Data.AllowSubdomains {
		role.Data.AllowSubdomains = true
	}

	if role.Data.AllowedDomains == "" {
		role.Data.AllowedDomains = "service.consul"
	}

	if !role.Data.BasicConstraintsValidForNonCA{
		role.Data.BasicConstraintsValidForNonCA = true
	}
}