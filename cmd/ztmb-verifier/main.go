package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
)

func main() {
	// Server TLS configuration
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Println("Failed to load key pair:", err)
		return
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS13, // Ensure TLS 1.3
	}

	// Listen on TCP port 853
	ln, err := tls.Listen("tcp", ":853", config)
	if err != nil {
		fmt.Println("Failed to listen on port 853:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server listening on port 853")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Accepted new connection")

	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Failed to read data:", err)
			}
			return
		}
		fmt.Printf("Received: %s\n", string(buf[:n]))

		// Echo the data back to the client
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Failed to write data:", err)
			return
		}
	}
}
