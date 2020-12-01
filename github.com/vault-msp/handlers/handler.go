package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/vault-msp/httpreq"
)

//Response to frontend
type Response struct{
	Response string `json:"response"`
}

//JSONResponse for Encoding struct to json and write to response writer 
func (response *Response) JSONResponse (rw io.Writer) error {
	err := json.NewEncoder(rw)
	return err.Encode(response)
}

//Vault to create http request object
type Vault struct{
	l *log.Logger
	requestObject httpreq.HTTPClient
}


//NewVaultRequest that returns requestObject
func NewVaultRequest(l *log.Logger,URL string,token string) *Vault{

	VaultURL,err := url.Parse(URL)  //parsing URL string to url.URL

	if err != nil {
		log.Fatalf(err.Error())
	}

	req := httpreq.CreateRequest(http.MethodPost,VaultURL,token)

	return &Vault{l:l,requestObject:req}
}
