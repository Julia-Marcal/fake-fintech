package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Julia-Marcal/fake-fintech/internal/http/router"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := router.StartRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
func TestLoginRoute(t *testing.T) {
	r := router.StartRouter()

	tests := []struct {
		name         string
		payload      interface{}
		expectedCode int
		expectedBody string
	}{
		{
			name: "successful login",
			payload: map[string]interface{}{
				"email":    "gmail@gmail.com",
				"password": "123456",
			},
			expectedCode: http.StatusOK,
			expectedBody: `"token"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadBytes, err := json.Marshal(tt.payload)
			if err != nil {
				t.Fatalf("Failed to marshal payload: %v", err)
			}
			payloadReader := bytes.NewReader(payloadBytes)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/api/login", payloadReader)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}
