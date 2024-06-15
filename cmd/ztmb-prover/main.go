package main

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
)

const (
  // TODO: Must be handled w/ env
  UPSTREAM_PORT = 853
  UPSTREAM_PROTOCOL = "tcp"

  SERVER_ADDR = "0.0.0.0"
  SERVER_PORT = 20053
  SERVER_PROTOCOL = "udp"
)

var (
  UPSTREAM_ADDR = os.Getenv("UPSTREAM_ADDR")
)

func upstream(data []byte) ([]byte, error) {
  // 1. Connect TLS
	config := &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS13,
	}
	conn, err := tls.Dial(UPSTREAM_PROTOCOL, fmt.Sprintf("%s:%d", UPSTREAM_ADDR, UPSTREAM_PORT), config)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
  log.Printf("upstream: Connected (%s:%d/%s)", UPSTREAM_ADDR, UPSTREAM_PORT, UPSTREAM_PROTOCOL)

  // 2. Send to upstream
	if _, err = conn.Write(data); err != nil {
		return nil, err
	}
  log.Printf("upstream: Sent (%s)", hex.EncodeToString(data))

  // 3. Response
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
  log.Printf("upstream: Received (%s)", hex.EncodeToString(buf))

  return buf[:n], nil
}

func server() error {
  // 1. Init
	addr := net.UDPAddr{
		Port: SERVER_PORT,
		IP:   net.ParseIP(SERVER_ADDR),
	}
	conn, err := net.ListenUDP(SERVER_PROTOCOL, &addr)
	if err != nil {
    return err
	}
	defer conn.Close()
  log.Printf("server: Started (%d/%s)", addr.Port, SERVER_PROTOCOL)

	buffer := make([]byte, 1024)
	for {
    // 2. Handle request
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
    log.Printf("server: Received (%d bytes, %v, %s)", n, remoteAddr, buffer[:n])

    // 3. Response upstream
    rcvd, err := upstream(buffer[:n])
		_, err = conn.WriteToUDP(rcvd, remoteAddr)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
    log.Printf("server: Response (%d bytes ,%v, %s)", n, remoteAddr, buffer[:n])
	}
}

func main() {
  if err := server(); err != nil {
    log.Fatalln("error:", err)
  }
}
