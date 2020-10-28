package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"bytes"
	"fmt"
)

//CreateRole handler to create a role for issuing certificates
func CreateRole(rw http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}
	client := &http.Client{}

	authToken := "myroot"
	res, _ := http.NewRequest("POST", "http://localhost:8200/v1/ordererCA/roles/msp", bytes.NewBuffer(reqBody))
	res.Header.Add("X-Vault-Token", authToken)

	resp, err := client.Do(res)

	if err != nil {
		log.Fatal("role creation error !")
	}
	fmt.Print(resp)
}