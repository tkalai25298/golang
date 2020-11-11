package handlers

import (
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
	tt := []struct{
		name string
		Pki data.Pki
		statusCode int	//statusResponse for Vault server
		statusResponse int	//statusResponse of Pki handler
		err string 
	}{
		{	name: "Valid Pki Request body & StatusOK",
			Pki : data.Pki{ Path:"sampleCA",},
			statusCode : 204,	
			statusResponse : 200,
		},
		{
			name: "valid Pki request body & StatusBadRequest",
			Pki : data.Pki{ Path:"sampleCA",},
			statusCode : 400,
			statusResponse : 502,
		},
		{
			name:"Invalid Pki Request body",
			Pki : data.Pki{},
			statusCode:400,
			statusResponse : 400,
		},
		{
			name: "http New Request error",
			Pki : data.Pki{ Path:"sampleCA",},
			err : "Could not create http New Request Object",
			statusResponse : 502,
		},
	}

	for _,tc := range tt{
		t.Run(tc.name, func(t *testing.T) {
			l := log.New(os.Stdout,"pki-test",log.LstdFlags)

			mock := &mocks.MockHTTPCall{}
			mock.CallFunc = func(path string, data []byte) (*http.Response, error){
				return &http.Response{
					StatusCode: tc.statusCode,
				}, nil
			}

			req := &Vault{
				requestObject: mock,
				l: l,
			}
			pki := tc.Pki

			err := pki.Validate()

			reqBody,err := json.Marshal(pki)
			if err != nil { t.Errorf("[Error] Marshal of pki struct failed!")}

			serverRes, err := http.NewRequest("POST","http://localhost:8000/pki",bytes.NewReader(reqBody))

			if err!= nil {
				t.Fatalf("could not send request: %v",err)
			}
			if tc.err != "" {
				if tc.statusResponse != http.StatusBadGateway{
					t.Errorf("[ERROR] %v",err)
				}
			}
				
			writer := httptest.NewRecorder()

			req.EnablePKI(writer,serverRes)

			result :=  writer.Result()

			if result.StatusCode != tc.statusResponse {
				t.Fatalf("expected %v got %v",tc.statusResponse,result.StatusCode)
			}
		})

	}

}