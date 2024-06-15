package main

import (
	"log"
	"net"
)

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

	buffer := make([]byte, UDP_MAX_SIZE)
	for {
		// 2. Handle request
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("error:", err)
			continue
		}
		log.Printf("server: Received (%d bytes from %v)", n, remoteAddr)

		// 3. Response upstream
		rcvd, err := upstream(buffer[:n])
    if err != nil {
      log.Println("error:", err)
      continue
    }
    if _, err := conn.WriteToUDP(rcvd, remoteAddr); err != nil {
			log.Println("error:", err)
			continue
		}
		log.Printf("server: Response (%d bytes ,%v)", n, remoteAddr)
	}
}
