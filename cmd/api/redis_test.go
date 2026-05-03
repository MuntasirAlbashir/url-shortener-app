package main

import "testing"

type MockRedisClient struct {
}

func (m *MockRedisClient) Set(key string, value string) string {
	return key
}

func TestRedisService(t *testing.T) {
	t.Run("Should store the URL against a database", func(t *testing.T) {
		expected := "key"
		service := NewRedisService(&MockRedisClient{})
		got, _ := service.Register(URL{Key: "key"})
		if got != expected {
			t.Errorf("got %s, want %s", got, expected)
		}
	})
}
