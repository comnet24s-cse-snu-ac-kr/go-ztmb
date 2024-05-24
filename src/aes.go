package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func pad(src []byte, size int) []byte {
	padtext := bytes.Repeat([]byte{0}, size-len(src))
	return append(src, padtext...)
}

func EncryptAES(plain, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	padSize := len(plain) + (aes.BlockSize - len(plain)%aes.BlockSize)
	plain = pad(plain, padSize)

	enc := cipher.NewCBCEncrypter(block, iv)
	cipher := make([]byte, len(plain))
	enc.CryptBlocks(cipher, plain)

	cipher = pad(cipher, 512)

	return cipher, iv, nil
}
