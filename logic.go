package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

const (
  PACKET = "24d7010000010000000000003f6257466a4c545930514739775a57357a63326775593239744c48567459574d744d544934514739775a57357a63326775593239744c47687459574d746332683f684d6930794e545973614731685979317a614745794c5455784d69786f6257466a4c584e6f5954457361473168597931745a4455745a585274514739775a573f357a63326775593239744c47687459574d74636d6c775a57316b4d5459774c575630625542766347567563334e6f4c6d4e766253786f6257466a4c584e6f59145445744f5459745a585274514739775a57357a630138016601310531333934300674756e6e656c076578616d706c65036f72670000050001"
  AES_256 = "12345678901234567890123456789012"
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
  raw       []byte
  header    DnsHeader
  question  DnsQuestion
}

func (dns *DnsPacket) GetStringQname() string {
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

  dns.raw = byteSlice
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

// pad applies PKCS7 padding to the plaintext to make it a multiple of the block size
func pad(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// EncryptAES encrypts a byte slice using AES-256 in CBC mode
func EncryptAES(plaintext, key []byte) ([]byte, []byte, error) {
	// Create a new AES cipher with the given key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	// Generate a new IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	// Pad the plaintext
	plaintext = pad(plaintext, aes.BlockSize)

	// Create a new CBC encrypter
	mode := cipher.NewCBCEncrypter(block, iv)

	// Encrypt the plaintext
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, iv, nil
}

func main() {
  packet := new(DnsPacket)
  if err := packet.DecodeHexString(PACKET); err != nil {
    fmt.Println("error:", err)
  }

  if err := Encode0x20(packet); err != nil {
    fmt.Println("error:", err)
  }

  fmt.Println(packet.GetStringQname())

  cipher, iv, err := EncryptAES(packet.raw, []byte(AES_256))
  if err != nil {
    fmt.Println("error:", err)
  }

  fmt.Println("Cipher:", cipher)
  fmt.Println("IV:", iv)
}
