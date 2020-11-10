package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"github.com/vault-msp/data"
	"os"
	"log"
	"net/http"
	"github.com/vault-msp/mocks"
	"testing"
)

	

func TestPki(t *testing.T) {
	l := log.New(os.Stdout,"test",log.LstdFlags)

	empty := make([]byte, 1)

	mock := &mocks.MockHTTPCall{}
	mock.CallFunc = func(path string, data []byte) (*http.Response, error){
		return &http.Response{
			StatusCode: 204,
			Body: ioutil.NopCloser(bytes.NewReader(empty)),
		}, nil
	}
	req := &Vault{
		requestObject: mock,
		l: l,
	}
	pki := data.Pki{
		Path: "something",
	}
	reqBody,err := json.Marshal(pki)

	serverRes, err := http.NewRequest("POST","http://localhost:8000/pki",bytes.NewReader([]byte(reqBody)))

	if err!= nil {
		t.Fatalf("could not send request: %v",err)
	}
		
	writer := httptest.NewRecorder()

	req.EnablePKI(writer,serverRes)

	result :=  writer.Result()
	// defer result.Body.Close()

	if result.StatusCode != 200 {
		t.Errorf("expected 204 got %v",result.StatusCode)
	}

}