package handlers

import (
	"net/url"
	"github.com/vault-msp/httpreq"
	"log"
	"net/http"
)

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
