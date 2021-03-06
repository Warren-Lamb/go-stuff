package main

import (
   "bytes"
   "io/ioutil"
   "net/http"
   "testing"
)

type MockHttpClient struct {}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
   response := &http.Response{
      Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test Mock Response"))),
   }

   return response, nil
}

func TestSendWithValidResponse(t *testing.T) {
   httpClient := &MockHttpClient{}
   err := send(httpClient, "IT_JUST_WORKS!")

   if err != nil {
      t.Errorf("Shouldn't have received an error with a valid MockHttpClient, got %s", err)
   }
}