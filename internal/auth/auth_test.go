package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("missing header", func(t *testing.T) {
		headers := http.Header{}
		_, err := auth.GetAPIKey(headers)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("malformed auth header", func(t *testing.T) {
		headers := http.Header{
			"Authorization": []string{"NotCorrectHeader"},
		}
		_, err := auth.GetAPIKey(headers)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("correct header", func(t *testing.T) {
		headers := http.Header{
			"Authorization": []string{"ApiKey some-key"},
		}
		value, err := auth.GetAPIKey(headers)
		if err != nil {
			t.Errorf("expected no error, got: %s", err)
			return
		}

		if value != "some-key" {
			t.Errorf("wrong api key, got=%s", value)
		}
	})
}
