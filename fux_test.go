package fux

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSchemeMatchers(t *testing.T) {
	f := New()

	f.Get("/", func(w http.ResponseWriter, request *http.Request) {
		_, err := w.Write([]byte("http response"))
		assert.Nil(t, err)
	}).Schemes("http")

	f.Get("/", func(w http.ResponseWriter, request *http.Request) {
		_, err := w.Write([]byte("https response"))
		assert.Nil(t, err)
	}).Schemes("https")

	assertResponseBody := func(t *testing.T, s *httptest.Server, expectedResponse string) {
		response, err := s.Client().Get(s.URL)
		assert.Nil(t, err)
		assert.Equal(t, 200, response.StatusCode)
		body, err := io.ReadAll(response.Body)
		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, string(body))
	}

	t.Run("httpServer", func(t *testing.T) {
		s := httptest.NewServer(f.Router)
		defer s.Close()
		assertResponseBody(t, s, "http response")
	})
	t.Run("httpsServer", func(t *testing.T) {
		s := httptest.NewTLSServer(f.Router)
		defer s.Close()
		assertResponseBody(t, s, "https response")
	})
}
