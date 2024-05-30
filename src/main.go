package main

import (
	"fmt"
)

func main() {
	// 1. Input
	input := new(InputJson)
	packet, aes, err := input.ReadFile()
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
		if err := q.Encode0x20(); err != nil {
			fmt.Println("error:", err)
			return
		}
	}
	packet.Print()

	// 4. Encrypt w/ AES_256_GCM
	if err := aes.Encrypt(packet.Unmarshal()); err != nil {
		fmt.Println("error:", err)
		return
	}
	aes.Print()

	// 5. Output
	output := new(OutputJson)
	if err := output.WriteFile(paddingOnly, aes.cipher, aes); err != nil {
		fmt.Println("error:", err)
		return
	}
}
