package main

import (
	"bytes"
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

func (dns *DnsPacket) Unmarshal() []byte {
	buf := make([]byte, 0)

	buf = append(buf, dns.header.id[:]...)
	buf = append(buf, dns.header.flags[:]...)
	buf = append(buf, dns.header.qdcount[:]...)
	buf = append(buf, dns.header.ancount[:]...)
	buf = append(buf, dns.header.nscount[:]...)
	buf = append(buf, dns.header.arcount[:]...)

	buf = append(buf, dns.question.qname[:]...)
	buf = append(buf, dns.question.qtype[:]...)
	buf = append(buf, dns.question.qclass[:]...)

	return buf
}

func (dns *DnsPacket) Marshal(b []byte) error {
	reader := bytes.NewReader(b)

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

	dns.question.qname = make([]byte, len(b)-16)
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
