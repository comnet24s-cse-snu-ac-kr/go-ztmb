package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

const (
	AES_KEY_BYTES   = 32
	AES_NONCE_BYTES = 12
)

type AESGCM struct {
  key []byte
  nonce []byte
  preCounterBlockSuffix []byte
  cipher []byte
}

func (ag *AESGCM) Encrypt(plaintext []byte) error {
	if len(ag.key) != AES_KEY_BYTES {
		return errors.New(fmt.Sprintf("AES key size mismatch (not %dbytes)", AES_KEY_BYTES))
	}

	block, err := aes.NewCipher(ag.key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	if len(ag.nonce) != AES_NONCE_BYTES {
		return errors.New(fmt.Sprintf("Nonce size mismatch (not %dbyte)", AES_NONCE_BYTES))
	}

	// Do not use `AdditionalData` for simplicity
	ag.cipher = gcm.Seal(nil, ag.nonce, plaintext, nil)
  return nil
}

func (ag *AESGCM) Print() {
	fmt.Printf("Cipher\n")
	fmt.Printf("  Key:                    %s\n", hex.EncodeToString(ag.key))
	fmt.Printf("  Nonce:                  %s\n", hex.EncodeToString(ag.nonce))
	fmt.Printf("  PreCounterBlockSuffix:  %s\n", hex.EncodeToString(ag.preCounterBlockSuffix))
	fmt.Printf("  Hex:                    %s\n", hex.EncodeToString(ag.cipher))
	fmt.Printf("  Length:                 %d\n", len(ag.cipher))
}
