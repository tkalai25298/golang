package test

import "net/http"

// MockClient is the mock client for HTTPClient interface
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}


// Do is the mock client's `Do` func for HTTPClient interface
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
		return m.DoFunc(req)
}