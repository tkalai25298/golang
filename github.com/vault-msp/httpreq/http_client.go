package httpreq

import (
	"net/http"
	"bytes"
)



// HTTPClient interface that performs similar to Client.Do 
type HTTPClient interface {	
	Do(req *http.Request) (*http.Response, error)
} 

//Client variable to set an HTTPClient instance
var Client HTTPClient	

func init() { 	//auto called when package is imported
	Client = &http.Client{}
}


//HTTPPost to create client request to endpoint
func (req *Request) HTTPPost() (*http.Response,error) {

	res, err := http.NewRequest(req.Method,req.URL,bytes.NewBuffer(req.Data))
	res.Header.Add("X-Vault-Token",req.Token)

	resp, err := Client.Do(res)		

	return resp,err
} 
