package main

import (
	"bytes"
	"encoding/binary"
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
	fmt.Printf("  Flags:     %08b %08b\n", h.flags[0], h.flags[1])
	fmt.Printf("  QDCOUNT:   0x%s\n", hex.EncodeToString(h.qdcount[:]))
	fmt.Printf("  ANCOUNT:   0x%s\n", hex.EncodeToString(h.ancount[:]))
	fmt.Printf("  NSCOUNT:   0x%s\n", hex.EncodeToString(h.nscount[:]))
	fmt.Printf("  ARCOUNT:   0x%s\n", hex.EncodeToString(h.arcount[:]))
}

func (h *DnsHeader) Length() int {
  return 12
}

// ---

type DnsPacket struct {
	header     DnsHeader
	question   []DnsQuestion
	answer     []DnsResourceRecord // Not used for simplicity
	authority  []DnsResourceRecord // Not used for simplicity
  rem        []byte // Remaining (not-marshalled) bytes
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
  b = b[12:]

  qdcount := int(binary.BigEndian.Uint16(p.header.qdcount[:]))
  for i := 0; i < qdcount; i++ {
    q := new(DnsQuestion)
    if err := q.Marshal(b); err != nil {
      return err
    }
    b = b[q.Length():]
    p.question = append(p.question, *q)
  }
  p.rem = b

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

  buf = append(buf, p.rem...)

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

  fmt.Println("Remaining not-marshalled bytes")
  fmt.Println(prettyBytes(p.rem, 1))

	for i, rr := range p.additional {
		fmt.Printf("Additional Rerouces Record #%d\n", i)
		rr.Print()
	}
}

func (p *DnsPacket) Length() int {
  acc := p.header.Length() + len(p.rem)
  for _, qd := range p.question {
    acc += qd.Length()
  }
  for _, an := range p.answer {
    acc += an.Length()
  }
  for _, ns := range p.authority {
    acc += ns.Length()
  }
  for _, ar := range p.additional {
    acc += ar.Length()
  }
  return acc
}

func (p *DnsPacket) AppendAdditionalRR(rr DnsResourceRecord) {
	arcount := binary.BigEndian.Uint16(p.header.arcount[:])
	binary.BigEndian.PutUint16(p.header.arcount[:], arcount+1)
	p.additional = append(p.additional, rr)
}
