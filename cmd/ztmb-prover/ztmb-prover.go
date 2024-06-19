package main

import (
	"github.com/ztmb/ztmb/cmd/ztmb-prover/app"
	"log"
)

func main() {
	if err := app.Server(); err != nil {
		log.Fatalln("error:", err)
	}
}
