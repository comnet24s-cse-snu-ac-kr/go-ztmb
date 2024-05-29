package main

import (
	"bytes"
	"encoding/binary"
)

const (
	OPT_CODE_PADDING = 12
)

type DnsRROPT struct {
	optionCode   [2]byte
	optionLength [2]byte
	padding      []byte
}

func (rr *DnsRROPT) Marshal(b []byte) error {
	reader := bytes.NewReader(b)

	if _, err := reader.Read(rr.optionCode[:]); err != nil {
		return err
	}
	if _, err := reader.Read(rr.optionLength[:]); err != nil {
		return err
	}

	rr.padding = make([]byte, len(b)-4)
	if _, err := reader.Read(rr.padding); err != nil {
		return err
	}

	return nil
}

func (rr *DnsRROPT) Unmarshal() []byte {
	buf := make([]byte, 0)

	buf = append(buf, rr.optionCode[:]...)
	buf = append(buf, rr.optionLength[:]...)

	buf = append(buf, rr.padding...)

	return buf
}

func (rr *DnsRROPT) FillZero(size int) {
	binary.LittleEndian.PutUint16(rr.optionCode[:], OPT_CODE_PADDING)
	binary.LittleEndian.PutUint16(rr.optionLength[:], uint16(size))
	rr.padding = bytes.Repeat([]byte{0}, size)
}
