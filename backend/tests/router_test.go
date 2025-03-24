package test

import (
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
