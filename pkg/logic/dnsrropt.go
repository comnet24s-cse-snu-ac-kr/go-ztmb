package logic

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

const (
	OPT_CODE_PADDING = 12
)

type DnsResourceRecord interface {
	Marshal(b []byte) error
	Unmarshal() []byte
	Print()
	Length() int
}

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

func (rr *DnsRROPT) FillZero(size int) error {
	if size < 0 {
		return errors.New("Negative size")
	}
	binary.BigEndian.PutUint16(rr.optionCode[:], OPT_CODE_PADDING)
	binary.BigEndian.PutUint16(rr.optionLength[:], uint16(size))
	rr.padding = bytes.Repeat([]byte{0}, size)

	return nil
}

func (rr *DnsRROPT) Print() {
	fmt.Println("RR OPT")
	fmt.Printf("  OPTCODE:  0x%s\n", hex.EncodeToString(rr.optionCode[:]))
	fmt.Printf("  OPTLEN:   0x%s\n", hex.EncodeToString(rr.optionLength[:]))
	fmt.Printf("  PADDING:\n%s\n", prettyBytes(rr.padding[:], 2))
}

func (rr *DnsRROPT) Length() int {
	return len(rr.padding) + 4
}
