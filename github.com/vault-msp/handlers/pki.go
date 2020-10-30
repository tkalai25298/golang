package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"

	"github.com/vault-msp/httpreq"
	config"github.com/vault-msp/config"
	data"github.com/vault-msp/data"
)
//PKI struct for http.handler
type PKI struct {
	l *log.Logger
}

// NewPKI creates a pki handler for new pki engine with logger
func NewPKI(l *log.Logger) *PKI {
	return &PKI{l}
}

//EnablePKI handler to create a pki engine to store certs
func (p *PKI) EnablePKI(rw http.ResponseWriter,r *http.Request){

	pki := data.Pki{}

	config,err := config.SetConfig() //getting env variables for vault server
	if err != nil {
		log.Fatalf(err.Error())
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("error reading request body", err)
	}
	err = json.Unmarshal(reqBody,&pki)

	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	// log.Printf("Received: %+v\n", pki.Data)

	vaultData,err := json.Marshal(pki.Data)
	p.l.Printf("%s",vaultData)

		reqObj := httpreq.CreateRequest("POST","http://"+config.VaultURL+"/v1/sys/mounts/"+pki.Path,config.VaultToken,vaultData)
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
