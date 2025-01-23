package uecho

import (
	"fmt"
	"net"
)

// EchoServer represents a simple UDP Echo Server.
type EchoServer struct {
	Address string
	Port    int
	conn    *net.UDPConn
}

// NewEchoServer creates a new instance of EchoServer.
func NewEchoServer(address string, port int) *EchoServer {
	return &EchoServer{
		Address: address,
		Port:    port,
	}
}

// Start initializes the UDP connection and begins listening for messages.
func (s *EchoServer) Start() error {
	addr := net.UDPAddr{
		IP:   net.ParseIP(s.Address),
		Port: s.Port,
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	s.conn = conn

	fmt.Printf("EchoServer started at %s:%d\n", s.Address, s.Port)
	defer s.conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := s.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading from UDP client: %v\n", err)
			continue
		}

		go s.handleClient(clientAddr, buffer[:n])
	}
}

// handleClient echoes back the received message to the client.
func (s *EchoServer) handleClient(clientAddr *net.UDPAddr, message []byte) {
	fmt.Printf("Received message from %s: %s\n", clientAddr.String(), string(message))

	_, err := s.conn.WriteToUDP(message, clientAddr)
	if err != nil {
		fmt.Printf("Error sending response to %s: %v\n", clientAddr.String(), err)
	}
}

// Stop gracefully stops the server.
func (s *EchoServer) Stop() {
	if s.conn != nil {
		_ = s.conn.Close()
		fmt.Println("EchoServer stopped.")
	}
}

