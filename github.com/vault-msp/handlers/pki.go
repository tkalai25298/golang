package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"bytes"
	"encoding/json"
)

//Pki struct for request params body
type Pki struct {
	Path string `json:"path"`
	Data Data `json:"data"`
}
//Data struct for request params with data to be passed for vault
type Data struct {
	Type string `json:"type"`
	Config Config `json:"config"`
	SealWrap bool `json:"seal_wrap"`
}
//Config struct to be passed for Data struct
type Config struct {
	MaxLeaseTTL string `json:"max_lease_ttl"`
	DefaultLeaseTTL string `json:"default_lease_ttl"`
}


//EnablePKI handler to create a pki engine to store certs
func EnablePKI(rw http.ResponseWriter,r *http.Request){
	pki := Pki{}

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
	log.Printf("%v",vaultData)

	client := &http.Client{}

		// auth := "Bearer "+"myroot"
		res, _ := http.NewRequest("POST","http://localhost:8200/v1/sys/mounts/"+pki.Path,bytes.NewBuffer(vaultData))
		res.Header.Add("X-Vault-Token","myroot")

		resp, err := client.Do(res)

		if err != nil {
			log.Fatal("pki creation error !")
		}
		fmt.Print(resp.Body)

}

// func (pki *Pki) Decoder(r io.Reader) error {
// 	e := json.NewDecoder(r)
// 	return e.Decode(pki)
// }