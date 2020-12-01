package mocks

import (
	"net/http"
)


type MockHTTPDo struct{
	DoFunc func( *http.Request) (*http.Response, error)
}

func (m *MockHTTPDo) Do(req *http.Request) (*http.Response, error){
	return m.DoFunc(req)
}

type MockHTTPCall struct{
	CallFunc func (string, []byte) (*http.Response, error)
}

func(m *MockHTTPCall) HTTPCall(path string, data []byte) (*http.Response, error){
	return m.CallFunc(path,data)
}

