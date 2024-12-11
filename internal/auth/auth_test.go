package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyMalformedt(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Apikeyabc123")
	_, err = GetAPIKey(req.Header)
	if err == nil {
		t.Fatalf("expected: %v, got: %v", errors.New("malformed authorization header"), err)
	}
}
