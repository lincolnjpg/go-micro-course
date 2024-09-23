package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(request *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func Test_Authenticate(t *testing.T) {
	jsonToReturn := `
{
	"error": false,
	"message": "some message"
}
`

	client := NewTestClient(func(request *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client

	postBody := map[string]interface{}{
		"email":    "me@here.com",
		"password": "very secret",
	}

	body, _ := json.Marshal(postBody)

	request, _ := http.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	handler.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusAccepted {
		t.Errorf("expected http.StatusAccepted, but got %d", responseRecorder.Code)
	}
}
