package main

import (
	"bytes"
	"fmt"
)

// ---

type QName []byte

func (qn *QName) String() string {
	out := ""
	b := *qn
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

// ---

type DnsResourceRecord interface {
	Marshal(b []byte) error
	Unmarshal() []byte
}

// ---

type DnsHeader struct {
	id      [2]byte
	flags   [2]byte
	qdcount [2]byte
	ancount [2]byte
	nscount [2]byte
	arcount [2]byte
}

func (h *DnsHeader) Marshal(b []byte) {
  copy(h.id[:], b[0:1])
  copy(h.flags[:], b[2:3])
  copy(h.qdcount[:], b[4:5])
  copy(h.ancount[:], b[6:7])
  copy(h.nscount[:], b[8:9])
  copy(h.arcount[:], b[10:11])

  h.Print()
}

func (h *DnsHeader) Unmarshal() []byte {
	buf := make([]byte, 0)

	buf = append(buf, h.id[:]...)
	buf = append(buf, h.flags[:]...)
	buf = append(buf, h.qdcount[:]...)
	buf = append(buf, h.ancount[:]...)
	buf = append(buf, h.nscount[:]...)
	buf = append(buf, h.arcount[:]...)

	return buf
}

func (h *DnsHeader) Print() {
  fmt.Println("Header")
  fmt.Printf("  ID:        0x%X\n", h.id)
  fmt.Printf("  Flags:     %b\n", h.flags)
  fmt.Printf("  QDCOUNT:   0x%X\n", h.qdcount)
  fmt.Printf("  ANCOUNT:   0x%X\n", h.ancount)
  fmt.Printf("  NSCOUNT:   0x%X\n", h.nscount)
  fmt.Printf("  ARCOUNT:   0x%X\n", h.arcount)
}

// ---

type DnsQuestion struct {
	qname  QName
	qtype  [2]byte
	qclass [2]byte
}

type DnsPacket struct {
	header     DnsHeader
	question   DnsQuestion
	answer     []DnsResourceRecord // Not used for simplicity
	authority  []DnsResourceRecord // Not used for simplicity
	additional []DnsResourceRecord
}

func (dns *DnsPacket) Unmarshal() []byte {
	buf := make([]byte, 0)

	buf = append(buf, dns.header.Unmarshal()...)

	buf = append(buf, dns.question.qname[:]...)
	buf = append(buf, dns.question.qtype[:]...)
	buf = append(buf, dns.question.qclass[:]...)

	for _, a := range dns.additional {
		buf = append(buf, a.Unmarshal()...)
	}

	return buf
}

/**
 * @arg     []byte    Assume that the packet has header and question field only
 * @return  []byte
 */
func (dns *DnsPacket) Marshal(b []byte) error {

  dns.header.Marshal(b[:11])

  reader := bytes.NewReader(b[12:])
	dns.question.qname = make([]byte, len(b)-4)
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
