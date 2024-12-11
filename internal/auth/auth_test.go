package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyMalformedt(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Apikeyabc123")
	authResult, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", "", authResult)
	}
}
