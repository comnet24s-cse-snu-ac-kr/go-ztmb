package app

import (
  "github.com/ztmb/ztmb/pkg/logic"
)

func preproxy(packet []byte) ([]byte, error) {
  dns := new(logic.DnsPacket)
  if err := dns.Marshal(packet); err != nil {
    return nil, err
  }

	// padding := new(logic.DnsRROPT)
	// padding.FillZero(512 - dns.Length() - 4)
  // dns.AppendAdditionalRR(padding)

	for _, q := range dns.Question() {
    if err := q.Qname().Encode0x20(); err != nil {
      return nil, err
    }
	}

  return dns.Unmarshal(), nil
}
