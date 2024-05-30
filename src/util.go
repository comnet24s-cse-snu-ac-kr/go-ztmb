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
