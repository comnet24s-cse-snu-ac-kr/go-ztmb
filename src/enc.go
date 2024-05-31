package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"

	chacha "golang.org/x/crypto/chacha20poly1305"
)

const (
	ENC_KEY_BYTES   = 32
	ENC_NONCE_BYTES = 12
)

// ---

type AEAD interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Print()
	Key() []byte
	Nonce() []byte
	PreCounterBlockSuffix() []byte
}

// ---

type aesParam struct {
	key                   []byte
	nonce                 []byte
	preCounterBlockSuffix []byte
}

func (param *aesParam) Key() []byte                   { return param.key }
func (param *aesParam) Nonce() []byte                 { return param.nonce }
func (param *aesParam) PreCounterBlockSuffix() []byte { return param.preCounterBlockSuffix }

func (param *aesParam) Encrypt(plaintext []byte) ([]byte, error) {
	if len(param.key) != ENC_KEY_BYTES {
		return nil, errors.New(fmt.Sprintf("Key size mismatch (not %dbytes)", ENC_KEY_BYTES))
	}

	if len(param.nonce) != ENC_NONCE_BYTES {
		return nil, errors.New(fmt.Sprintf("Nonce size mismatch (not %dbyte)", ENC_NONCE_BYTES))
	}

	block, err := aes.NewCipher(param.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Do not use `AdditionalData` for simplicity
	return gcm.Seal(nil, param.nonce, plaintext, nil), nil
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

func (param *chachaPolyParam) Encrypt(plaintext []byte) ([]byte, error) {
	if len(param.key) != ENC_KEY_BYTES {
		return nil, errors.New(fmt.Sprintf("Key size mismatch (not %dbytes)", ENC_KEY_BYTES))
	}

	if len(param.nonce) != ENC_NONCE_BYTES {
		return nil, errors.New(fmt.Sprintf("Nonce size mismatch (not %dbyte)", ENC_NONCE_BYTES))
	}

	aead, err := chacha.New(param.key)
	if err != nil {
		return nil, err
	}

	// Do not use `AdditionalData` for simplicity
	return aead.Seal(nil, param.nonce, plaintext, nil), nil
}

func (param *chachaPolyParam) Print() {
	fmt.Printf("Cipher (Chacha20Poly1305)\n")
	fmt.Printf("  Key:                    %s\n", hex.EncodeToString(param.key))
	fmt.Printf("  Nonce:                  %s\n", hex.EncodeToString(param.nonce))
	fmt.Printf("  PreCounterBlockSuffix:  %s\n", hex.EncodeToString(param.preCounterBlockSuffix))
}
