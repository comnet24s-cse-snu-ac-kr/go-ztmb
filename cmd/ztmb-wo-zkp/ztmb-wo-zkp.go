package main

import (
	"encoding/hex"
	"fmt"
	ztmb "github.com/ztmb/ztmb/pkg/logic"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func main() {
	// 1. Input
	input := new(ztmb.InputJson)
	packet, aead, err := input.ReadFile()
	check(err)

	// 2. Add EDNS0 padding opt
	padding := new(ztmb.DnsRROPT)
	padding.FillZero(512 - packet.Length() - 4) // Decrease 4 for OPT_RR metadata
	packet.AppendAdditionalRR(padding)
	paddingOnly := packet.Unmarshal()

	// 3. Encode 0x20
	totalLength := 0
	totalModified := 0
	for _, q := range packet.Question() {
		n, err := q.Qname().Encode0x20()
		check(err)
		totalLength += q.Qname().Length()
		totalModified += n
	}
	packet.Print()
	fmt.Printf("  0x20 Modified/Total:    %d/%d\n", totalModified, totalLength)

	// 4. Encrypt w/ AES_256_GCM
	cipher, tag, err := aead.Encrypt(packet.Unmarshal())
	check(err)
	aead.Print()

	fmt.Printf("  Length:                 %d\n", len(cipher))
	fmt.Printf("  Tag:                    %s\n", hex.EncodeToString(tag))
	fmt.Printf("  Hex:\n%s\n", ztmb.PrettyBytes(cipher, 2))

	// 5. Output
	output := new(ztmb.OutputJson)
	check(output.WriteFile(paddingOnly, cipher, aead))
}
