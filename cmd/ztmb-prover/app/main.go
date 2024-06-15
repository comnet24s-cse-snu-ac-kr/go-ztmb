package main

import (
	"log"
)

func main() {
	if err := server(); err != nil {
		log.Fatalln("error:", err)
	}
}
