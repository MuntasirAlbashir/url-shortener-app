package main

import (
	"testing"
)

type MockRedisClient struct {
	store map[string]string
}

func (m *MockRedisClient) Set(key string, value string) bool {
	m.store[key] = value
	return true
}

type FailedRedisClient struct {
}

func (f *FailedRedisClient) Set(key string, value string) bool {
	return false
}

func TestRedisService(t *testing.T) {
	t.Run("Should store the URL against a database", func(t *testing.T) {
		store := map[string]string{}
		service := NewRedisService(&MockRedisClient{store})
		got, _ := service.Register(URLRequest{Value: "https://example.com"})
		if got == "" {
			t.Errorf("Expected the key to be set, but got %s", got)
		}
		if store[got] != "https://example.com" {
			t.Errorf("got %s, want https://example.com", got)
		}
	})
	t.Run("Should return the same key for the same input", func(t *testing.T) {
		store := map[string]string{}

		service := NewRedisService(&MockRedisClient{store})
		firstCall, _ := service.Register(URLRequest{Value: "https://example.com"})
		secondCall, _ := service.Register(URLRequest{Value: "https://example.com"})

		if firstCall != secondCall {
			t.Errorf("expected keys to match: got %s first call and  %s second call", firstCall, secondCall)
		}
	})

	t.Run("Should throw an error if the datastore fails", func(t *testing.T) {
		service := NewRedisService(&FailedRedisClient{})
		got, err := service.Register(URLRequest{Value: "https://example.com"})
		if err == nil {
			t.Errorf("Expected an error")
		}
		if got != "" {
			t.Errorf("Expected empty string")
		}
	})
}
