package main

import (
	"fmt"
)

func toStringSlice(bs []byte) []string {
	ret := make([]string, len(bs))
	for i, b := range bs {
		ret[i] = fmt.Sprintf("%d", b)
	}
	return ret
}

func prettyBytes(bs []byte) string {
  out := ""
  for i := 0; i < len(bs); {
    out += fmt.Sprintf("%03d", i)
    for j := 0; (j < 16) && (i+j < len(bs)); j++ {
      out += fmt.Sprintf(" %02x", bs[i+j])
    }
    out += "\n"
    i += 16
  }
  return out
}
