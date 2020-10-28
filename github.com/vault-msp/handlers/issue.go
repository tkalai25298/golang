package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"bytes"
)

//IssueCert handler to issue certs by a role
func IssueCert(rw http.ResponseWriter,r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}
	client := &http.Client{}

	authToken := "myroot"
	res, _ := http.NewRequest("POST", "http://localhost:8200/v1/ordererCA/issue/msp", bytes.NewBuffer(reqBody))
	res.Header.Add("X-Vault-Token", authToken)

	resp, err := client.Do(res)

	if err != nil {
		log.Fatal("issue certs error !")
	}
	fmt.Print(resp)
}