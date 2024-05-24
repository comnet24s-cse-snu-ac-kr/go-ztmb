package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

type DnsHeader struct {
	id      [2]byte
	flags   [2]byte
	qdcount [2]byte
	ancount [2]byte
	nscount [2]byte
	arcount [2]byte
}

type DnsQuestion struct {
	qname  []byte
	qtype  [2]byte
	qclass [2]byte
}

type DnsPacket struct {
	raw      []byte
	header   DnsHeader
	question DnsQuestion
}

func (dns *DnsPacket) GetStringQname() string {
	out := ""
	b := dns.question.qname
	i := 0
	for i < len(b) {
		if b[i] == 0 {
			break
		}
		for j := 0; j < int(b[i]); j++ {
			out += fmt.Sprintf("%c", b[i+j+1])
		}
		out += "."
		i += int(b[i]) + 1
	}

	return out
}

func (dns *DnsPacket) DecodeHexString(input string) error {
	byteSlice, err := hex.DecodeString(input)
	if err != nil {
		return err
	}

	dns.raw = byteSlice
	reader := bytes.NewReader(byteSlice)

	if _, err := reader.Read(dns.header.id[:]); err != nil {
		return err
	}
	if _, err := reader.Read(dns.header.flags[:]); err != nil {
		return err
	}
	if _, err := reader.Read(dns.header.qdcount[:]); err != nil {
		return err
	}
	if _, err := reader.Read(dns.header.ancount[:]); err != nil {
		return err
	}
	if _, err := reader.Read(dns.header.nscount[:]); err != nil {
		return err
	}
	if _, err := reader.Read(dns.header.arcount[:]); err != nil {
		return err
	}

	dns.question.qname = make([]byte, len(byteSlice)-16)
	if _, err := reader.Read(dns.question.qname); err != nil {
		return err
	}

	if _, err := reader.Read(dns.question.qtype[:]); err != nil {
		return err
	}
	if _, err := reader.Read(dns.question.qclass[:]); err != nil {
		return err
	}

	return nil
}
