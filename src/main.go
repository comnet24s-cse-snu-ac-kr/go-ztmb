package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var packetHex string
	var aesKey string

	fmt.Println("Packet string (hex): ")
	fmt.Scanln(&packetHex)
	fmt.Println("AES key (256): ")
	fmt.Scanln(&aesKey)

  b, err := hex.DecodeString(packetHex)
  if err != nil {
    fmt.Println("error:", err)
  }

	packet := new(DnsPacket)
	if err := packet.Marshal(b); err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := Encode0x20(packet); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("0x20:", packet.String())

	cipher, iv, err := EncryptAES(packet.Unmarshal(), []byte(aesKey))
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Cipher:", hex.EncodeToString(cipher))
	fmt.Println("IV:", hex.EncodeToString(iv))
}
