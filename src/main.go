package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// 1. Input
	input := new(Input)
	if err := input.ReadJsonFile(); err != nil {
		fmt.Println("error:", err)
		return
	}

	packet := new(DnsPacket)
	if err := packet.Marshal(input.Packet); err != nil {
		fmt.Println("error:", err)
		return
	}

	// 2. Add EDNS0 padding opt
	output := new(Output)
	padding := new(DnsRROPT)
	padding.FillZero(512 - len(input.Packet) - 4)
	packet.AppendAdditionalRR(padding)
	output.Packet = packet.Unmarshal()

	// 3. Encode 0x20
	if err := packet.question[0].Encode0x20(); err != nil {
		fmt.Println("error:", err)
		return
	}
	packet.Print()

	// 4. Encrypt w/ AES_256_GCM
	output.Key = input.AesKey
	output.Nonce = input.Nonce
	output.PreCounterBlockSuffix = input.PreCounterBlockSuffix
	cipher, err := EncryptAES256GCM(input.AesKey, input.Nonce, packet.Unmarshal())
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	output.CipherText = cipher

	// 5. Output
	if err := output.WriteJsonFile(); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Cipher")
	fmt.Printf("  Hex:     %s\n", hex.EncodeToString(cipher))
	fmt.Printf("  Length:  %d\n", len(cipher))
}
