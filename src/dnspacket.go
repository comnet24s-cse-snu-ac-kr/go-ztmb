package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
  "github.com/iden3/go-iden3-crypto/poseidon"
)

// ---

type QName []byte

func (qn *QName) String() string {
	out := ""
	b := *qn

  for i := 0; i < len(b); i += int(b[i]) + 1 {
		if b[i] == 0 {
			break
		}
		for j := 0; j < int(b[i]); j++ {
			out += fmt.Sprintf("%c", b[i+j+1])
		}
		out += "."
	}

	return out
}

func (qn *QName) Encode0x20() error {
  b := *qn

  digest, err := poseidon.HashBytes(b)
  if err != nil {
    return err
  }

  for i:= 0; i < len(b); i += int(b[i]) + 1 {
    if b[i] == 0 {
      break
    }
    for j := 0; j < int(b[i]); j++ {
      c := b[i+j+1]
      if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
        if digest.Bit(i+j+1) == 0 {
          b[i+j+1] = c | 0x20
        } else {
          b[i+j+1] = c &^ 0x20
        }
      }
    }
  }

	return nil
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

	if _, err := reader.Read(h.id[:]); err != nil {
		return err
	}
	if _, err := reader.Read(h.flags[:]); err != nil {
		return err
	}
	if _, err := reader.Read(h.qdcount[:]); err != nil {
		return err
	}
	if _, err := reader.Read(h.ancount[:]); err != nil {
		return err
	}
	if _, err := reader.Read(h.nscount[:]); err != nil {
		return err
	}
	if _, err := reader.Read(h.arcount[:]); err != nil {
		return err
	}

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

type DnsPacket struct {
	header     DnsHeader
	question   []DnsQuestion
	answer     []DnsResourceRecord // Not used for simplicity
	authority  []DnsResourceRecord // Not used for simplicity
	additional []DnsResourceRecord
}

/**
 * @arg     []byte    Assume that the packet has header and question field only
 * @return  []byte
 */
func (p *DnsPacket) Marshal(b []byte) error {
	if err := p.header.Marshal(b[:11]); err != nil {
		return err
	}

	q := new(DnsQuestion)
	if err := q.Marshal(b[12:]); err != nil {
		return err
	}
	p.question = append(p.question, *q)

	return nil
}

func (p *DnsPacket) Unmarshal() []byte {
	buf := make([]byte, 0)

	buf = append(buf, p.header.Unmarshal()...)

	for _, q := range p.question {
		buf = append(buf, q.Unmarshal()...)
	}

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

	for i, q := range p.question {
		fmt.Printf("Question #%d\n", i)
		q.Print()
	}

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
	binary.BigEndian.PutUint16(p.header.arcount[:], arcount+1)
	p.additional = append(p.additional, rr)
}
