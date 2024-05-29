package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func EncryptAES(plain, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	if len(plain) != 512 {
		return nil, nil, errors.New("Packet size mismatch (not 512byte)")
	}

	enc := cipher.NewCBCEncrypter(block, iv)
	cipher := make([]byte, len(plain))
	enc.CryptBlocks(cipher, plain)

	return cipher, iv, nil
}
