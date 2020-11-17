package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/vault-msp/data"
	"github.com/vault-msp/mocks"
)
func TestVault (t *testing.T) {
		mock := &mocks.MockHTTPCall{}
			mock.CallFunc = func(path string, data []byte) (*http.Response, error){
				return &http.Response{
					StatusCode: 200,
				}, nil
			}

			l := log.New(os.Stdout,"pki-test",log.LstdFlags)
			req := &Vault{
				requestObject: mock,
				l: l,
			}

			reqData := data.VaultComplete{
				PKI: data.Pki{
					Path:"sampleCA",
				},
			}
			reqBody,err := json.Marshal(reqData)
			if err != nil { t.Errorf("[Error] Marshal of pki struct failed!")}

			serverRes, err := http.NewRequest("POST","http://localhost:8000/pki",bytes.NewReader(reqBody))

			if err!= nil {
				t.Fatalf("could not send request: %v",err)
			}

			writer := httptest.NewRecorder()

			req.VaultInterface(writer,serverRes)

			result :=  writer.Result()

			if result.StatusCode != 200 {
				t.Fatalf("expected 200 but got %v",result.StatusCode)
			}

}