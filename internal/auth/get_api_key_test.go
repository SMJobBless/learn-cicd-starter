package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		it          string
		headers     http.Header
		expectedKey string
		expectErr   bool
	}{
		{
			it:        "no authorization header returns error",
			headers:   http.Header{},
			expectErr: true,
		},
		{
			it:        "malformed header missing ApiKey prefix",
			headers:   http.Header{"Authorization": []string{"Bearer sometoken"}},
			expectErr: true,
		},
		{
			it:        "malformed header with only one part",
			headers:   http.Header{"Authorization": []string{"ApiKey"}},
			expectErr: true,
		},
		{
			it:          "valid ApiKey header returns key",
			headers:     http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			expectedKey: "my-secret-key",
			expectErr:   false,
		},
	}

	for _, c := range cases {
		actual, err := GetAPIKey(c.headers)
		if c.expectErr {
			if err != nil {
				t.Logf("✅ it %v", c.it)
			} else {
				t.Errorf("❌ it %v: expected error but got none", c.it)
			}
		} else {
			if err != nil {
				t.Errorf("❌ it %v: unexpected error: %v", c.it, err)
			} else if actual == c.expectedKey {
				t.Logf("✅ it %v", c.it)
			} else {
				t.Errorf("❌ it %v: expected %q, got %q", c.it, c.expectedKey, actual)
			}
		}
	}
}
