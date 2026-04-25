package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockURLService struct {
}

func (m *MockURLService) Register(url URL) (string, error) {
	return url.Key, nil
}

func TestRegisterURL(t *testing.T) {
	t.Run("Should return the key in the body", func(t *testing.T) {

		service := &MockURLService{}
		server := NewUrlServer(service)
		req := httptest.NewRequest(http.MethodPost, "/?key=ABC", nil)
		res := httptest.NewRecorder()

		server.RegisterURL(res, req)

		if res.Body.String() != "ABC" {
			t.Errorf("Expected the key to be 123, got %s", res.Body.String())
		}
	})
}
