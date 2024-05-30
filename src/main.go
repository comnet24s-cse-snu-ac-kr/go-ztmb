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

	output := new(Output)

	packet := new(DnsPacket)
	if err := packet.Marshal(input.Packet); err != nil {
		fmt.Println("error:", err)
		return
	}

	padding := new(DnsRROPT)
	padding.FillZero(512 - len(input.Packet) - 4)
	packet.AppendAdditionalRR(padding)

	if err := packet.question[0].Encode0x20(); err != nil {
		fmt.Println("error:", err)
		return
	}

	output.Packet = packet.Unmarshal()

	packet.Print()

	output.Key = input.AesKey
	output.Nonce = input.Nonce
	cipher, err := EncryptAES256GCM(input.AesKey, input.Nonce, packet.Unmarshal())
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	output.CipherText = cipher

	if err := output.WriteJsonFile(); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Cipher")
	fmt.Printf("  Hex:     %s\n", hex.EncodeToString(cipher))
	fmt.Printf("  Length:  %d\n", len(cipher))
}
