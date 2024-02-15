package fux

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	f := New()
	f.Get("/test", func(w http.ResponseWriter, request *http.Request) {
		Response(w).Status(http.StatusOK)
		_, err := w.Write([]byte("test response"))
		assert.Nil(t, err)
	})

	assertResponseBody := func(t *testing.T, s *httptest.Server, expectedResponse string) {
		response, err := s.Client().Get(fmt.Sprintf("%s/test", s.URL))
		assert.Nil(t, err)
		assert.Equal(t, 200, response.StatusCode)
		body, err := io.ReadAll(response.Body)
		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, string(body))
	}

	t.Run("get request", func(t *testing.T) {
		s := httptest.NewServer(f.Router)
		defer s.Close()
		assertResponseBody(t, s, "test response")
	})
}
