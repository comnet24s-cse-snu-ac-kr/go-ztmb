package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func check(err error) {
  if err != nil {
    fmt.Println("error:", err)
    os.Exit(1)
  }
}

func main() {
	// 1. Input
	input := new(InputJson)
	packet, aead, err := input.ReadFile()
  check(err)

	// 2. Add EDNS0 padding opt
	padding := new(DnsRROPT)
	padding.FillZero(512 - len(input.Packet)/2 - 4)
	packet.AppendAdditionalRR(padding)
	paddingOnly := packet.Unmarshal()

	// 3. Encode 0x20
	for _, q := range packet.question {
		check(q.qname.Encode0x20())
	}
	packet.Print()

	// 4. Encrypt w/ AES_256_GCM
	cipher, tag, err := aead.Encrypt(packet.Unmarshal())
  check(err)
	aead.Print()

	fmt.Printf("  Length:                 %d\n", len(cipher))
	fmt.Printf("  Tag:                    %s\n", hex.EncodeToString(tag))
	fmt.Printf("  Hex:\n%s\n", prettyBytes(cipher, 2))

	// 5. Output
	output := new(OutputJson)
	check(output.WriteFile(paddingOnly, cipher, aead))
}
