package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKet extracts an API Key from
// the headers of an HTTP request
// Example:
// Authorization: APIKey (inseert apikey here)
func GetApiKey(headers http.Header) (string, error){

	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("no auth token found")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("incorrect auth info")
	}

	return vals[1], nil
}