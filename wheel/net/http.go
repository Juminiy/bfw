package net

import (
	"bytes"
	"errors"
	"net/http"
	"strings"
)

const (
	requestContentType = "application/json"
)

var (
	unsupportedHTTPMethodError = errors.New("unsupported HTTP Method")
)

func RequestByURL(method, url string) (*http.Response, error) {
	method = strings.ToUpper(method)
	switch method {
	case http.MethodGet:
		{
			return http.Get(url)
		}
	case http.MethodPost:
		{
			var buf bytes.Buffer
			return http.Post(url, requestContentType, &buf)
		}
	default:
		{
			panic(unsupportedHTTPMethodError)
		}
	}
}
