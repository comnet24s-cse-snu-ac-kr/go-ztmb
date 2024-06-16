package app

import (
	"log"

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
    if n, err := q.Qname().Encode0x20(); err != nil {
      return nil, err
    } else {
      for i, sig := range logic.IodineSigs {
        if sig.Check(packet) {
          log.Printf("preproxy: Iodine signature detected (ID:%d)", i)
        }
      }
      log.Printf("preproxy: Modified byte count (%d/%d)", n, q.Length())
    }
	}

  return dns.Unmarshal(), nil
}
