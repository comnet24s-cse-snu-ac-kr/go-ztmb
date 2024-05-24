package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

const (
  PACKET = "24d7010000010000000000003f6257466a4c545930514739775a57357a63326775593239744c48567459574d744d544934514739775a57357a63326775593239744c47687459574d746332683f684d6930794e545973614731685979317a614745794c5455784d69786f6257466a4c584e6f5954457361473168597931745a4455745a585274514739775a573f357a63326775593239744c47687459574d74636d6c775a57316b4d5459774c575630625542766347567563334e6f4c6d4e766253786f6257466a4c584e6f59145445744f5459745a585274514739775a57357a630138016601310531333934300674756e6e656c076578616d706c65036f72670000050001"
)

type DnsHeader struct {
  id        [2]byte
  flags     [2]byte
  qdcount   [2]byte
  ancount   [2]byte
  nscount   [2]byte
  arcount   [2]byte
}

type DnsQuestion struct {
  qname   []byte
  qtype   [2]byte
  qclass  [2]byte
}

type DnsPacket struct {
  header    DnsHeader
  question  DnsQuestion
}

func (dns *DnsPacket) PrintQname() string {
  out := ""
  b := dns.question.qname
  i := 0
  for ; i < len(b); {
    if b[i] == 0 {
      break
    }
    for j := 0; j < int(b[i]); j++ {
      out += fmt.Sprintf("%c", b[i + j + 1])
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

  reader := bytes.NewReader(byteSlice)

  if _, err := reader.Read(dns.header.id[:]); err != nil {return err }
  if _, err := reader.Read(dns.header.flags[:]); err != nil { return err }
  if _, err := reader.Read(dns.header.qdcount[:]); err != nil { return err }
  if _, err := reader.Read(dns.header.ancount[:]); err != nil { return err }
  if _, err := reader.Read(dns.header.nscount[:]); err != nil { return err }
  if _, err := reader.Read(dns.header.arcount[:]); err != nil { return err }

  dns.question.qname = make([]byte, len(byteSlice) - 16)
  if _, err := reader.Read(dns.question.qname); err != nil { return err }

  if _, err := reader.Read(dns.question.qtype[:]); err != nil { return err }
  if _, err := reader.Read(dns.question.qclass[:]); err != nil { return err }

  return nil
}

func (dns *DnsPacket) Encode0x20() error {
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

func main() {
  packet := new(DnsPacket)
  if err := packet.DecodeHexString(PACKET); err != nil {
    fmt.Println("error:", err)
  }

  if err := packet.Encode0x20(); err != nil {
    fmt.Println("error:", err)
  }

  fmt.Println(packet.PrintQname())
}
