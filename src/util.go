package main

import (
	"fmt"
	"strings"
)

func toStringSlice(bs []byte) []string {
	ret := make([]string, len(bs))
	for i, b := range bs {
		ret[i] = fmt.Sprintf("%d", b)
	}
	return ret
}

func prettyBytes(bs []byte, indent int) string {
	out := ""
	indentStr := strings.Repeat("  ", indent)
	for i := 0; i < len(bs); {
		out += fmt.Sprintf("%s%03d", indentStr, i)
		for j := 0; (j < 16) && (i+j < len(bs)); j++ {
			out += fmt.Sprintf(" 0x%02x", bs[i+j])
		}
		out += "\n"
		i += 16
	}
	return out
}
