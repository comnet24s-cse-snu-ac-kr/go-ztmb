package main

import (
	"log"
  "github.com/ztmb/ztmb/cmd/ztmb-prover/app"
)

func main() {
	if err := app.Server(); err != nil {
		log.Fatalln("error:", err)
	}
}
