package httpreq

import (
	// "net/url"
	"net/http"
	"bytes"
)

//Request object to make http request
type Request struct {
	Method string 
	URL string
	Token string
	Data []byte
}

//CreateRequest fn to create the Request Object for HTTPCall()
func CreateRequest (method string,url string, token string,data []byte) (req *Request) {
	reqObj := Request{}
	reqObj.Method = method
	reqObj.URL = url
	reqObj.Token = token
	reqObj.Data = data
	
	return &reqObj

}

//HTTPCall to create client request to endpoint
func (req *Request) HTTPCall() (*http.Response,error) {

	client := &http.Client{}

	res, err := http.NewRequest(req.Method,req.URL,bytes.NewBuffer(req.Data))
	res.Header.Add("X-Vault-Token",req.Token)

	resp, err := client.Do(res)

	return resp,err
} 