package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

const (
  VRF_URL = "localhost"
  VRF_PORT = 853
  VRF_PROTOCOL = "tcp"
)

func send(data []byte) ([]byte, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS13,
	}

	conn, err := tls.Dial(VRF_PROTOCOL, fmt.Sprintf("%s:%d", VRF_URL, VRF_PORT), config)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
  log.Println(fmt.Sprintf("Connected to verifier: %s:%d/%s", VRF_URL, VRF_PORT, VRF_PROTOCOL))

	if _, err = conn.Write(data); err != nil {
		return nil, err
	}
  log.Println("Sent")

	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	log.Printf("Received")

  return buf[:n], nil
}

func main() {
  for {
    var input string
    if _, err := fmt.Scanln(&input); err != nil {
      log.Fatal(err)
    }
    send([]byte(input))
  }
}
