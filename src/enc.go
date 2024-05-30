package main

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

const (
	AES_KEY_BYTES   = 32
	AES_NONCE_BYTES = 12
)

func EncryptAES256GCM(key, nonce, plaintext []byte) ([]byte, error) {
	if len(key) != AES_KEY_BYTES {
		return nil, errors.New(fmt.Sprintf("AES key size mismatch (not %dbit)", AES_KEY_BYTES))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(nonce) != AES_NONCE_BYTES {
		return nil, errors.New(fmt.Sprintf("Nonce size mismatch (not %dbyte)", AES_NONCE_BYTES))
	}

  // Do not use `AdditionalData` for simplicity
	return gcm.Seal(nil, nonce, plaintext, nil), nil
}
