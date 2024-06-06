package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

type DnsQuestion struct {
	qname  QName
	qtype  [2]byte
	qclass [2]byte
}

func (q *DnsQuestion) Marshal(b []byte) error {
	if err := q.qname.Marshal(b); err != nil {
		return err
	}

	reader := bytes.NewReader(b[q.qname.length:])

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

	buf = append(buf, q.qname.Unmarshal()...)
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

func (q *DnsQuestion) Length() int {
	return q.qname.Length() + 4
}
