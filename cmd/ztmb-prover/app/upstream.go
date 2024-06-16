package app

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"log"
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
	buf := make([]byte, TLS_MAX_SIZE)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	log.Printf("upstream: Received (%s)", hex.EncodeToString(buf))

	return buf[:n], nil
}
