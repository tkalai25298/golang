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


func TestCA(t *testing.T) {
	tt := []struct{
		name string
		CA data.RootCA
		statusCode int	//statusResponse for Vault server
		statusResponse int	//statusResponse of Pki handler
		err string 
	}{
		{	name: "Valid RootCA Request body & StatusOK",
			CA : data.RootCA{ Path:"sampleCA",
							  Data:data.CAData{CommonName:"sampleCA",Organization:"sample"},
							},
			statusCode : 204,	
			statusResponse : 200,
		},
		{
			name: "valid RootCA request body & StatusBadRequest",
			CA : data.RootCA{ Path:"sampleCA",
							  Data:data.CAData{CommonName:"sampleCA",Organization:"sample"},
							},
			statusCode : 400,
			statusResponse : 502,
		},
		{
			name:"Invalid RootCA Request body",
			CA : data.RootCA{Path:"sampleCA",},
			statusCode:400,
			statusResponse : 400,
		},
		{
			name: "http New Request error",
			CA : data.RootCA{ Path:"sampleCA",
							  Data:data.CAData{CommonName:"sampleCA",Organization:"sample"},
							},
			err : "Could not create http New Request Object",
			statusResponse : 502,
		},
	}

	for _,tc := range tt{
		t.Run(tc.name, func(t *testing.T) {
			l := log.New(os.Stdout,"Root CA-test",log.LstdFlags)

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
			CA := tc.CA

			err := CA.Validate()

			reqBody,err := json.Marshal(CA)
			if err != nil { t.Errorf("[Error] Marshal of CA struct failed!")}

			serverRes, err := http.NewRequest("POST","http://localhost:8000/ca",bytes.NewReader(reqBody))

			if err!= nil {
				t.Fatalf("could not send request: %v",err)
			}
			if tc.err != "" {
				if tc.statusResponse != http.StatusBadGateway{
					t.Errorf("[ERROR] %v",err)
				}
			}
				
			writer := httptest.NewRecorder()

			req.IssueCA(writer,serverRes)

			result :=  writer.Result()

			if result.StatusCode != tc.statusResponse {
				t.Fatalf("expected %v got %v",tc.statusResponse,result.StatusCode)
			}
		})

	}

}