package logic

import "bytes"

type SnortSignature struct {
  msg string
  content []byte
  offset  int
}

var (
  IodineSigs = []*SnortSignature{
    {msg: "Iodine signature (1): covert iodine tunnel request", content: []byte{0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, offset: 2},
    {msg: "Iodine signature (2): covert iodine tunnel request IP packet encapsulated", content: []byte{0x45, 0x10}, offset: 13},
    {msg: "Iodine signature (3): covert iodine tunnel request", content: []byte{0x00, 0x00, 0x29, 0x10, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00}, offset: -1},
  }
)

func (s *SnortSignature) Check(mainSlice []byte) bool {
	mainLen, sigLen := len(mainSlice), len(s.content)

	if sigLen == 0 || sigLen > mainLen {
		return false
	}

  if s.offset == -1 {
    for i := 0; i <= mainLen-sigLen; i++ {
      if bytes.Compare(mainSlice[i:i+sigLen], s.content) == 0 {
        return true
      }
    }
  } else {
    if bytes.Compare(mainSlice[s.offset:s.offset+sigLen], s.content) == 0 {
      return true
    }
  }

	return false
}
