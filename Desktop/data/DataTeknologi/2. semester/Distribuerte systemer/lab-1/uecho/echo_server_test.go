package uecho

import (
	"testing"
)

func TestEchoServer_Start(t *testing.T) {
	server := NewEchoServer("127.0.0.1", 8080)

	go func() {
		if err := server.Start(); err != nil {
			t.Errorf("Failed to start server: %v", err)
		}
	}()

	// Add additional tests to verify server behavior if needed.
	server.Stop()
}

