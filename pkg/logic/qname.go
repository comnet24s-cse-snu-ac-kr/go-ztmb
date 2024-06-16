package logic

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

type QName struct {
	labels [][]byte
	length int
}

func (qn *QName) Marshal(bs []byte) error {
	qn.length = 0
	qn.labels = make([][]byte, 0)

	for i := 0; i < len(bs); i += (int(bs[i]) + 1) {
		qn.length += 1 + int(bs[i])
		if bs[i] == 0 {
			break
		}
		label := make([]byte, 0)
		for j := 0; j < int(bs[i]); j++ {
			label = append(label, bs[i+j+1])
		}
		if len(label) > 63 {
			return errors.New(fmt.Sprintf("QNAME label size limit exceeded (%dbytes)", len(label)))
		}
		qn.labels = append(qn.labels, label)
	}

	if qn.length > 255 {
		return errors.New(fmt.Sprintf("QNAME total size limit exceeded (%dbytes)", qn.length))
	}

	return nil
}

func (qn *QName) Unmarshal() []byte {
	out := make([]byte, 0)
	for _, label := range qn.labels {
		out = append(out, byte(len(label)))
		out = append(out, label...)
	}
	out = append(out, 0)
	return out
}

func (qn *QName) String() string {
	out := ""
	for _, label := range qn.labels {
		out += string(label) + "."
	}
	return out
}

func (qn *QName) Length() int {
	return qn.length
}

func (qn *QName) Encode0x20() error {
	// fit 255, fill 46
	wdot := []byte("." + qn.String())
	padding := bytes.Repeat([]byte{'.'}, 255-len(wdot))
	wdot = append(wdot, padding...)

	digest, err := poseidon.HashBytesX(wdot, 9)
	if err != nil {
		return err
	}

	idx := 0
	for i, label := range qn.labels {
		for j, c := range label {
			idx++
			if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
				if digest.Bit(idx) == 0 {
					qn.labels[i][j] = c | 0x20
				} else {
					qn.labels[i][j] = c &^ 0x20
				}
			}
		}
		idx++
	}

	return nil
}
