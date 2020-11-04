package data

//Cert for req Params obj
type Cert struct{
	Path string `json:"path"`
	Roles string `json:"roles"`
	Data IssueCertData `json:"data"`
}

//IssueCertData to pass vault data config to issue certificates by a role
type IssueCertData struct {
	CommonName string `json:"common_name"`
	TTL string `json:"ttl"`
	AltNames string `json:"alt_names"`
}


// //Validate for Role struct json validation
// func (role *Role) Validate() error {
// 	role.SetDefaultValues()

// 	validate := validator.New()
// 	return validate.Struct(role)
// }	

// //SetDefaultValues to assign missing values to be passed for vault server
// func (role *Role) SetDefaultValues() {
// 	if role.Data.KeyType == "" {
// 		role.Data.KeyType = "ec"
// 	}	

// 	if role.Data.KeyBits == 0 {
// 		role.Data.KeyBits = 256
// 	}

// 	if role.Data.KeyUsage == "" {
// 		role.Data.KeyUsage = ["DigitalSignature"]
// 	}

// 	if role.Data.MaxTTL == "" {
// 		role.Data.MaxTTL = "3000h"
// 	}

// 	if role.Data.GenerateLease == false {
// 		role.Data.GenerateLease = true
// 	}

// 	if role.Data.BasicConstraintsValidForNonCA == false {
// 		role.Data.BasicConstraintsValidForNonCA = true
// 	}

// }