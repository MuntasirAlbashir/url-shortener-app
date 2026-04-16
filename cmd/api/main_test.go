package main

import "testing"

func TestStartWebServer(t *testing.T) {
	t.Run("Should print server started", func(t *testing.T) {
		want := "Server started"
		got := StartWebServer()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
