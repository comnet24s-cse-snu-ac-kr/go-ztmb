package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// 1. Input
	input := new(InputJson)
	packet, aead, err := input.ReadFile()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// 2. Add EDNS0 padding opt
	padding := new(DnsRROPT)
	padding.FillZero(512 - len(input.Packet)/2 - 4)
	packet.AppendAdditionalRR(padding)
	paddingOnly := packet.Unmarshal()

	// 3. Encode 0x20
	for _, q := range packet.question {
		if err := q.qname.Encode0x20(); err != nil {
			fmt.Println("error:", err)
			return
		}
	}
	packet.Print()

	// 4. Encrypt w/ AES_256_GCM
	cipher, tag, err := aead.Encrypt(packet.Unmarshal())
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	aead.Print()

	fmt.Printf("  Length:                 %d\n", len(cipher))
	fmt.Printf("  Tag:                    %s\n", hex.EncodeToString(tag))
	fmt.Printf("  Hex:\n%s\n", prettyBytes(cipher, 2))

	// 5. Output
	output := new(OutputJson)
	if err := output.WriteFile(paddingOnly, cipher, aead); err != nil {
		fmt.Println("error:", err)
		return
	}
}
