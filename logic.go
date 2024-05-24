package main

import (
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

const (
  PACKET = "24d7010000010000000000003f6257466a4c545930514739775a57357a63326775593239744c48567459574d744d544934514739775a57357a63326775593239744c47687459574d746332683f684d6930794e545973614731685979317a614745794c5455784d69786f6257466a4c584e6f5954457361473168597931745a4455745a585274514739775a573f357a63326775593239744c47687459574d74636d6c775a57316b4d5459774c575630625542766347567563334e6f4c6d4e766253786f6257466a4c584e6f59145445744f5459745a585274514739775a57357a630138016601310531333934300674756e6e656c076578616d706c65036f72670000050001"
)

func stringToBigInts(s string) []*big.Int {
  ints := make([]*big.Int, 0, len(s))

  for _, char := range s {
    charInt := big.NewInt(int64(char))
    ints = append(ints, charInt)
  }

  return ints
}

func logic(inputBI []*big.Int) (*big.Int, error) {
	hash, err := poseidon.Hash(inputBI)
  return hash, err
}

func main() {
	inputString := PACKET
  inputBI := stringToBigInts(inputString)
  outputBI, err:= logic(inputBI)
  if err != nil {
		fmt.Println("error:", err)
    return
  }

  fmt.Println("Poseidon hash:", outputBI)
}
