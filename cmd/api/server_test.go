package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockURLService struct {
}

func (m *MockURLService) Register(url URLRequest) (string, error) {
	return url.Value, nil
}

type FailedMockURLService struct {
}

func (f *FailedMockURLService) Register(url URLRequest) (string, error) {
	return "", errors.New("failed to register url in datastore")
}

func TestRegisterURL(t *testing.T) {
	t.Run("Should return the key in the body", func(t *testing.T) {

		service := &MockURLService{}
		server := NewUrlServer(service)
		req := httptest.NewRequest(http.MethodPost, "/?value=ABC", nil)
		res := httptest.NewRecorder()

		server.RegisterURL(res, req)

		if res.Body.String() != "ABC" {
			t.Errorf("Expected the key to be 123, got %s", res.Body.String())
		}
	})
	t.Run("Should return a statusBadRequest if the url is missing", func(t *testing.T) {
		service := &MockURLService{}
		server := NewUrlServer(service)
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		server.RegisterURL(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("Expected the status to be 400, got %d", res.Code)
		}
	})

	t.Run("Should return an error when the url service fails", func(t *testing.T) {
		server := NewUrlServer(&FailedMockURLService{})
		req := httptest.NewRequest(http.MethodPost, "/?value=ABC", nil)
		res := httptest.NewRecorder()

		server.RegisterURL(res, req)
		if res.Code != http.StatusInternalServerError {
			t.Errorf("Expected the status to be 500, got %d", res.Code)
		}
	})
}
