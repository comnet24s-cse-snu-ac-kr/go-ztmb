package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

// Convert a string to a slice of big.Int
func stringToBigInts(s string) []*big.Int {
	hash := sha256.Sum256([]byte(s))
	ints := make([]*big.Int, len(hash)/32)

	for i := 0; i < len(hash)/32; i++ {
		ints[i] = new(big.Int).SetBytes(hash[i*32 : (i+1)*32])
	}

	return ints
}

func main() {
	// Example input string
	inputString := "Hello, Poseidon!"

	// Convert the string to a slice of big.Int
	inputInts := stringToBigInts(inputString)

	// Calculate Poseidon hash
	hash, err := poseidon.Hash(inputInts)
	if err != nil {
		fmt.Println("Error calculating Poseidon hash:", err)
		return
	}

	// Print the result
	fmt.Println("Poseidon Hash:", hash)
}
