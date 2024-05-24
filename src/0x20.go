package main

import (
	"github.com/iden3/go-iden3-crypto/poseidon"
)

func Encode0x20(dns *DnsPacket) error {
	digest, err := poseidon.HashBytes(dns.question.qname)
	if err != nil {
		return err
	}

	for i, b := range dns.question.qname {
		if ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z') {
			if digest.Bit(i) == 0 {
				dns.question.qname[i] = b | 0x20
			} else {
				dns.question.qname[i] = b &^ 0x20
			}
		}
	}

	return nil
}
