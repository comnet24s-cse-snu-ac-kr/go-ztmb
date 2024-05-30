package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
  input := new(Input)
  if err := input.ReadJsonFile(); err != nil {
		fmt.Println("error:", err)
		return
  }

	b, err := hex.DecodeString(input.Packet)
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

	if err := packet.question[0].Encode0x20(); err != nil {
		fmt.Println("error:", err)
		return
	}

  packet.Print()

	cipher, err := EncryptAES256GCM([]byte(input.AesKey), []byte(input.Nonce), packet.Unmarshal(), []byte(input.AdditionalData))
	if err != nil {
		fmt.Println("error:", err)
		return
	}

  fmt.Println("Cipher")
	fmt.Printf("  Hex:     %s\n", hex.EncodeToString(cipher))
	fmt.Printf("  Length:  %d\n", len(cipher))
}
