package main

import (
	"fmt"
)

func main() {
  packet := new(DnsPacket)
  if err := packet.DecodeHexString(PACKET); err != nil {
    fmt.Println("error:", err)
  }

  if err := Encode0x20(packet); err != nil {
    fmt.Println("error:", err)
  }

  fmt.Println(packet.GetStringQname())

  cipher, iv, err := EncryptAES(packet.raw, []byte(AES_256))
  if err != nil {
    fmt.Println("error:", err)
  }

  fmt.Println("Cipher:", cipher)
  fmt.Println("IV:", iv)
}
