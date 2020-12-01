package httpreq_test

import (
	"net/url"
	"github.com/vault-msp/httpreq"
	"net/http"
	"testing"
	"github.com/vault-msp/mocks"
)

func TestPki(t *testing.T) {

	mock := &mocks.MockHTTPDo{}
	mock.DoFunc = func(req * http.Request) (*http.Response, error){
		return &http.Response{
			StatusCode: 200,
		}, nil
	}
	
	reqUrl,_ := url.ParseRequestURI("localhost:8080")
	req := &httpreq.Request{
		RequestObj: &http.Request{
			URL: reqUrl,
		},
		Client: mock,
	}
	
	reqp, err := req.HTTPCall("/path",make([]byte,10))
	if err != nil {
		t.Fatalf("Did'nt expect error but got %v",err)
	}

	if reqp.StatusCode != 200{
		t.Fatalf("Expected 200 but got %d",reqp.StatusCode)
	}
}