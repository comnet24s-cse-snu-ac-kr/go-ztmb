package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"

	chacha "golang.org/x/crypto/chacha20poly1305"
)

// ---

const (
	AEAD_TAG_SIZE = 16
)

type AEAD interface {
	Encrypt(plaintext []byte) ([]byte, []byte, error)
	Print()
	Key() []byte
	Nonce() []byte
	PreCounterBlockSuffix() []byte
}

// ---

const (
	AES_KEY_BYTES   = 32
	AES_NONCE_BYTES = 12
)

type aesParam struct {
	key                   []byte
	nonce                 []byte
	preCounterBlockSuffix []byte
}

func (param *aesParam) Key() []byte                   { return param.key }
func (param *aesParam) Nonce() []byte                 { return param.nonce }
func (param *aesParam) PreCounterBlockSuffix() []byte { return param.preCounterBlockSuffix }

func (param *aesParam) Encrypt(plaintext []byte) ([]byte, []byte, error) {
	if len(param.key) != AES_KEY_BYTES {
		return nil, nil, errors.New(fmt.Sprintf("Key size mismatch (not %dbytes)", AES_KEY_BYTES))
	}

	if len(param.nonce) != AES_NONCE_BYTES {
		return nil, nil, errors.New(fmt.Sprintf("Nonce size mismatch (not %dbyte)", AES_NONCE_BYTES))
	}

	block, err := aes.NewCipher(param.key)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	// Do not use `AdditionalData` for simplicity
	c := gcm.Seal(nil, param.nonce, plaintext, nil)
	return c[:len(c)-AEAD_TAG_SIZE], c[len(c)-AEAD_TAG_SIZE:], nil
}

func (param *aesParam) Print() {
	fmt.Printf("Cipher (AES-256-GCM)\n")
	fmt.Printf("  Key:                    %s\n", hex.EncodeToString(param.key))
	fmt.Printf("  Nonce:                  %s\n", hex.EncodeToString(param.nonce))
	fmt.Printf("  PreCounterBlockSuffix:  %s\n", hex.EncodeToString(param.preCounterBlockSuffix))
}

// ---

type chachaPolyParam struct {
	key                   []byte
	nonce                 []byte
	preCounterBlockSuffix []byte
}

func (param *chachaPolyParam) Key() []byte                   { return param.key }
func (param *chachaPolyParam) Nonce() []byte                 { return param.nonce }
func (param *chachaPolyParam) PreCounterBlockSuffix() []byte { return param.preCounterBlockSuffix }

func (param *chachaPolyParam) Encrypt(plaintext []byte) ([]byte, []byte, error) {
	if len(param.key) != chacha.KeySize {
		return nil, nil, errors.New(fmt.Sprintf("Key size mismatch (not %dbytes)", chacha.KeySize))
	}

	if len(param.nonce) != chacha.NonceSize {
		return nil, nil, errors.New(fmt.Sprintf("Nonce size mismatch (not %dbyte)", chacha.NonceSize))
	}

	aead, err := chacha.New(param.key)
	if err != nil {
		return nil, nil, err
	}

	// Do not use `AdditionalData` for simplicity
	c := aead.Seal(nil, param.nonce, plaintext, nil)
	return c[:len(c)-AEAD_TAG_SIZE], c[len(c)-AEAD_TAG_SIZE:], nil
}

func (param *chachaPolyParam) Print() {
	fmt.Printf("Cipher (Chacha20Poly1305)\n")
	fmt.Printf("  Key:                    %s\n", hex.EncodeToString(param.key))
	fmt.Printf("  Nonce:                  %s\n", hex.EncodeToString(param.nonce))
	fmt.Printf("  PreCounterBlockSuffix:  %s\n", hex.EncodeToString(param.preCounterBlockSuffix))
}
