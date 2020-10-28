package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/vault-msp/httpreq"
)
//Role struct for creating role request obj
type Role struct{
	Path string `json:"path"`
	Roles string `json:"roles"`
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
	AllowAnyName bool `json:"allow_any_name"`
	OU string `json:"ou"`
	Organization string `json:"organization"`
	AllowedDomains string `json:"allowed_domains"`
	AllowSubdomains bool `json:"allow_subdomains"`
	BasicConstraintsValidForNonCA bool `json:"basic_constraints_valid_for_non_ca"`
}

//CreateRole handler to create a role for issuing certificates
func CreateRole(rw http.ResponseWriter, r *http.Request) {

	role := Role{}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}

	err = json.Unmarshal(reqBody,&role)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	vaultData,err := json.Marshal(role.Data)

	reqObj := httpreq.CreateRequest("POST","http://localhost:8200/v1/"+role.Path+"/roles/"+role.Roles,"myroot",vaultData)
		resp, err := reqObj.HTTPCall()

		if err != nil {
			log.Fatal("could not send request! Server connection issue")
		}
		log.Println("The Status Response ==>> ",resp.StatusCode)

		if resp.StatusCode != 204 {
			log.Panicf("NON 204 STATUS CODE")
		}

		defer resp.Body.Close()
}