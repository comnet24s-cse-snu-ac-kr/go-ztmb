package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Input struct {
  Packet string `json:"packet"`
  AesKey string `json:"aes-key"`
  Nonce string `json:"nonce"`
  AdditionalData string `json:"additional-data"`
}

func (input *Input) ReadJsonFile() error {
  if len(os.Args) != 2 {
    return errors.New("Input JSON file not provided")
  }

  raw, err := os.ReadFile(os.Args[1])
  if err != nil {
    return err
  }

  if err := json.Unmarshal(raw, input); err != nil {
    return err
  }

  return nil
}
