package httpreq

import (
	"net/url"
	"net/http"
	"bytes"
	"io/ioutil"
	// "log"
)

//Request object to make http request
type Request struct {
	RequestObj *http.Request
	Client HTTPDo
}


// HTTPDo interface for Client.Do 
type HTTPDo interface {	
	Do(req *http.Request) (*http.Response, error)
} 

//HTTPClient interface for HTTPCall
type HTTPClient interface {	
	HTTPCall(path string, data []byte) (*http.Response, error)
} 

//CreateRequest fn to create the Request Object for HTTPCall()
func CreateRequest (method string,url *url.URL, token string) *Request {

	client := &http.Client{}
	
	reqObj := &http.Request{
		Method : method,
		URL : url,
		Header : map[string][]string {
			"X-Vault-Token" : {token},
			},
	}
	
	RequestObject := Request{
		RequestObj : reqObj,
		Client : client,
	}

	return &RequestObject
}

//HTTPCall to create client request to endpoint
func (req *Request) HTTPCall(path string, data []byte) (*http.Response,error) {

	request := req.RequestObj

	request.URL.Path = path

	requestBody := ioutil.NopCloser(bytes.NewReader([]byte(data)))
	request.Body = requestBody


	response,err := req.Client.Do(request)

	return response,err
} 
