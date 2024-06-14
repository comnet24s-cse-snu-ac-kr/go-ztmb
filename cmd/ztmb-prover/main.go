package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	// Client TLS configuration
	config := &tls.Config{
		InsecureSkipVerify: true, // For demonstration purposes only
		MinVersion:         tls.VersionTLS13, // Ensure TLS 1.3
	}

	// Connect to the server
	conn, err := tls.Dial("tcp", "localhost:853", config)
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// Send a message to the server
	message := "Hello, secure world!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Failed to send message:", err)
		return
	}
	fmt.Printf("Sent: %s\n", message)

	// Read the response from the server
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}
	fmt.Printf("Received: %s\n", string(buf[:n]))
}
