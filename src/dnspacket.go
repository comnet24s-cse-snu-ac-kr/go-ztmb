package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
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
  Print()
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

func (h *DnsHeader) Marshal(b []byte) error {
  reader := bytes.NewReader(b)

	if _, err := reader.Read(h.id[:]); err != nil { return err }
	if _, err := reader.Read(h.flags[:]); err != nil { return err }
	if _, err := reader.Read(h.qdcount[:]); err != nil { return err }
	if _, err := reader.Read(h.ancount[:]); err != nil { return err }
	if _, err := reader.Read(h.nscount[:]); err != nil { return err }
	if _, err := reader.Read(h.arcount[:]); err != nil { return err }

  return nil
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
  fmt.Printf("  ID:        0x%s\n", hex.EncodeToString(h.id[:]))
  fmt.Printf("  Flags:     %b %b\n", h.flags[0], h.flags[1])
  fmt.Printf("  QDCOUNT:   0x%s\n", hex.EncodeToString(h.qdcount[:]))
  fmt.Printf("  ANCOUNT:   0x%s\n", hex.EncodeToString(h.ancount[:]))
  fmt.Printf("  NSCOUNT:   0x%s\n", hex.EncodeToString(h.nscount[:]))
  fmt.Printf("  ARCOUNT:   0x%s\n", hex.EncodeToString(h.arcount[:]))
}

// ---

type DnsQuestion struct {
	qname  QName
	qtype  [2]byte
	qclass [2]byte
}

func (q *DnsQuestion) Marshal(b []byte) error {
  reader := bytes.NewReader(b)

	q.qname = make([]byte, len(b)-4)
	if _, err := reader.Read(q.qname); err != nil {
		return err
	}

	if _, err := reader.Read(q.qtype[:]); err != nil {
		return err
	}

	if _, err := reader.Read(q.qclass[:]); err != nil {
		return err
	}

	return nil
}

func (q *DnsQuestion) Unmarshal() []byte {
	buf := make([]byte, 0)

	buf = append(buf, q.qname[:]...)
	buf = append(buf, q.qtype[:]...)
	buf = append(buf, q.qclass[:]...)

	return buf
}

func (q *DnsQuestion) Print() {
  fmt.Println("Question")
  fmt.Printf("  QNMAE:     %s\n", q.qname.String())
  fmt.Printf("  QTYPE:     0x%s\n", hex.EncodeToString(q.qtype[:]))
  fmt.Printf("  QCLASS:    0x%s\n", hex.EncodeToString(q.qclass[:]))
}

// ---

type DnsPacket struct {
	header     DnsHeader
	question   DnsQuestion
	answer     []DnsResourceRecord // Not used for simplicity
	authority  []DnsResourceRecord // Not used for simplicity
	additional []DnsResourceRecord
}

/**
 * @arg     []byte    Assume that the packet has header and question field only
 * @return  []byte
 */
func (p *DnsPacket) Marshal(b []byte) error {
  if err := p.header.Marshal(b[:11]); err != nil { return err }
  if err := p.question.Marshal(b[12:]); err != nil { return err }
	return nil
}

func (p *DnsPacket) Unmarshal() []byte {
	buf := make([]byte, 0)

	buf = append(buf, p.header.Unmarshal()...)
	buf = append(buf, p.question.Unmarshal()...)

	for _, rr := range p.answer {
		buf = append(buf, rr.Unmarshal()...)
	}

	for _, rr := range p.authority {
		buf = append(buf, rr.Unmarshal()...)
	}

	for _, rr := range p.additional {
		buf = append(buf, rr.Unmarshal()...)
	}

	return buf
}

func (p *DnsPacket) Print() {
  p.header.Print()
  p.question.Print()

	for i, rr := range p.answer {
    fmt.Printf("Answer Rerouces Record #%d\n", i)
    rr.Print()
	}

	for i, rr := range p.authority {
    fmt.Printf("Authority Rerouces Record #%d\n", i)
    rr.Print()
	}

	for i, rr := range p.additional {
    fmt.Printf("Additional Rerouces Record #%d\n", i)
    rr.Print()
	}
}

func (p *DnsPacket) AppendAdditionalRR(rr DnsResourceRecord) {
  arcount := binary.BigEndian.Uint16(p.header.arcount[:])
  binary.BigEndian.PutUint16(p.header.arcount[:], arcount + 1)
  p.additional = append(p.additional, rr)
}
