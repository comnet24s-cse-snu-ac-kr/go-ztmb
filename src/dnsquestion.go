package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

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

func (q *DnsQuestion) Encode0x20() error {
	digest, err := poseidon.HashBytes(q.qname)
	if err != nil {
		return err
	}

	for i, b := range q.qname {
		if ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z') {
			if digest.Bit(i) == 0 {
				q.qname[i] = b | 0x20
			} else {
				q.qname[i] = b &^ 0x20
			}
		}
	}

	return nil
}
