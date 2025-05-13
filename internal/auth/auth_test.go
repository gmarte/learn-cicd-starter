// internal/auth/auth_test.go
//package auth

//import (
	"errors"
	"net/http"
	"testing"
)

// TestGetAPIKey_Valid tests the successful retrieval of an API key.
func TestGetAPIKey_Valid(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey mysecretkey123")

	expectedKey := "mysecretkey123"
	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	if apiKey != expectedKey {
		t.Errorf("expected key %q, but got %q", expectedKey, apiKey)
	}
}

// TestGetAPIKey_NoAuthHeader tests the case where the Authorization header is missing.
func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{} // Empty headers

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error, but got nil")
	}

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected error %q, but got %q", ErrNoAuthHeaderIncluded, err)
	}
}

// TestGetAPIKey_MalformedHeader tests a malformed Authorization header.
func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer somekey") // Incorrect scheme

	expectedErrorMessage := "malformed authorization header"
	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error for malformed header, but got nil")
	}

	if err.Error() != expectedErrorMessage {
		t.Errorf("expected error message %q, but got %q", expectedErrorMessage, err.Error())
	}

	// You could add another case for malformed (e.g. too few parts)
	headers = http.Header{}
	headers.Add("Authorization", "ApiKey") // Not enough parts
	_, err = GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error for malformed header (too few parts), but got nil")
	}
	if err.Error() != expectedErrorMessage {
		t.Errorf("expected error message %q for too few parts, but got %q", expectedErrorMessage, err.Error())
	}
}
