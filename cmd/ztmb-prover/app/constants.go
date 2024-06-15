package main

import (
  "os"
)

const (
	// TODO: Must be handled w/ env
	UPSTREAM_PORT     = 853
	UPSTREAM_PROTOCOL = "tcp"

	SERVER_ADDR     = "0.0.0.0"
	SERVER_PORT     = 20053
	SERVER_PROTOCOL = "udp"

  TLS_MAX_SIZE = 16 * 1024
  UDP_MAX_SIZE = 65526
)

var (
	UPSTREAM_ADDR = os.Getenv("UPSTREAM_ADDR")
)
