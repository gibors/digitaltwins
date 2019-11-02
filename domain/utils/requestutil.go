package utils

import (
	"bytes"
	"net/http"
)

func GenerateRequest(headers map[string]string, url string, method string, token string, body []byte) *http.Request {

	var request *http.Request

	if body != nil {
		request, _ = http.NewRequest(method, url, bytes.NewBuffer(body))
	} else {
		request, _ = http.NewRequest(method, url, nil)
	}

	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}
	request.Header.Set("Content-type", "application/json")

	if headers != nil {
		for m := range headers {
			request.Header.Set(m, headers[m])
		}
	}

	return request
}
