package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
)

// ---

type InputJson struct {
	Packet string `json:"packet"`
	Key    string `json:"key"`
	Nonce  string `json:"nonce"`

	// Note that the word "couter" indicates suffix for PreCounterBlock (J0)
	// which is 0x00000001 for 12-byte nonce (IV).
	// See NIST SP 800-38D, section 7.1.
	Counter string `json:"counter"`
}

func (input *InputJson) ReadFile() (*DnsPacket, AEAD, error) {
	if len(os.Args) != 2 {
		return nil, nil, errors.New("Input JSON file not provided")
	}

	raw, err := os.ReadFile(os.Args[1])
	if err != nil {
		return nil, nil, err
	}

	if err := json.Unmarshal(raw, input); err != nil {
		return nil, nil, err
	}

	dns := new(DnsPacket)
	if packet, err := hex.DecodeString(input.Packet); err != nil {
		return nil, nil, err
	} else {
		dns.Marshal(packet)
	}

	// aead := new(aesParam)
	aead := new(chachaPolyParam)

	if aead.key, err = hex.DecodeString(input.Key); err != nil {
		return nil, nil, err
	}

	if aead.nonce, err = hex.DecodeString(input.Nonce); err != nil {
		return nil, nil, err
	}

	if aead.preCounterBlockSuffix, err = hex.DecodeString(input.Counter); err != nil {
		return nil, nil, err
	}

	return dns, aead, nil
}

// ---

type OutputJson struct {
	Key        []string `json:"key"`
	Nonce      []string `json:"nonce"`
	Packet     []string `json:"packet"`
	CipherText []string `json:"ciphertext"`

	// Same as `inputJson` struct.
	Counter []string `json:"counter"`
}

func (output *OutputJson) WriteFile(packet, cipher []byte, aead AEAD) error {
	output.Key = toStringSlice(aead.Key())
	output.Nonce = toStringSlice(aead.Nonce())
	output.Counter = toStringSlice(aead.PreCounterBlockSuffix())

	output.Packet = toStringSlice(packet)
	output.CipherText = toStringSlice(cipher)

	dat, err := json.Marshal(output)
	if err != nil {
		return err
	}

	if err := os.WriteFile("result.json", dat, 0644); err != nil {
		return err
	}

	return nil
}
