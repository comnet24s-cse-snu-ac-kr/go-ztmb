package main

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"
)

// TODO: env config
const (
	SERVER_PORT     = 853
	SERVER_PROTOCOL = "tcp"

	UPSTREAM_ADDR     = "127.0.0.1"
	UPSTREAM_PORT     = 53
	UPSTREAM_PROTOCOL = "udp"

  TLS_MAX_SIZE = 16 * 1024
  UDP_MAX_SIZE = 65535

  TLS_KEY_FILE = "proxy.key"
  TLS_CRT_FILE = "proxy.crt"
)

func server() error {
	// 1. TLS configuration
	cert, err := tls.LoadX509KeyPair(TLS_CRT_FILE, TLS_KEY_FILE)
	if err != nil {
		return err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS13,
	}

	// 2. Listen on TCP port 853
	ln, err := tls.Listen(SERVER_PROTOCOL, fmt.Sprintf(":%d", SERVER_PORT), config)
	if err != nil {
		return err
	}
	defer ln.Close()
	log.Printf("server: Started (%d/%s)\n", SERVER_PORT, SERVER_PROTOCOL)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("error:", err)
			continue
		}
		go handleServer(conn)
	}
}

func handleServer(conn net.Conn) {
	defer conn.Close()
	log.Println("handleServer: Connection accepted")

	msg := make([]byte, TLS_MAX_SIZE)
	for {
		// Receive data
		n, err := conn.Read(msg)
		if err != nil {
			log.Println("error:", err)
			return
		}
		msg = msg[:n]
		log.Printf("handleServer: Received (%s)\n", hex.EncodeToString(msg))

		upstreamResponse, err := handleUpstream(msg)
		if err != nil {
			log.Println("error:", err)
			return
		}

		// Echo the data back to the client
		_, err = conn.Write(upstreamResponse)
		if err != nil {
			log.Println("error:", err)
			return
		}
	}
}

func handleUpstream(msg []byte) ([]byte, error) {
	// Setup upstream
	udpAddr, err := net.ResolveUDPAddr(UPSTREAM_PROTOCOL, fmt.Sprintf("%s:%d", UPSTREAM_ADDR, UPSTREAM_PORT))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP(UPSTREAM_PROTOCOL, nil, udpAddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Send the message
	_, err = conn.Write(msg)
	if err != nil {
		return nil, err
	}
	log.Printf("handleUpstream: Sent (%s)\n", hex.EncodeToString(msg))

	// Set a read deadline
	if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return nil, err
	}

	// Receive the response
	response := make([]byte, UDP_MAX_SIZE)
	n, _, err := conn.ReadFromUDP(response)
	if err != nil {
		return nil, err
	}
	response = response[:n]
	log.Printf("handleUpstream: Received (%s)\n", hex.EncodeToString(response))

	return response, nil
}

func main() {
	if err := server(); err != nil {
		log.Fatalln("error", err)
	}
}
