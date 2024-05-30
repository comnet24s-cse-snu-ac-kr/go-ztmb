package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	nonce := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ad := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

	padding := new(DnsRROPT)
	padding.FillZero(512 - len(b) - 4)
  packet.AppendAdditionalRR(padding)

	if err := Encode0x20(packet); err != nil {
		fmt.Println("error:", err)
		return
	}

  packet.Print()

	cipher, err := EncryptAES256GCM([]byte(aesKey), nonce, packet.Unmarshal(), ad)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Cipher:", hex.EncodeToString(cipher))
	fmt.Println("Cipher len:", len(cipher))
}
