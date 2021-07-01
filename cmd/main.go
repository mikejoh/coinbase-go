package main

import (
	"log"

	"github.com/mikejoh/coinbase-go/cmd/cb"
)

func main() {
	err := cb.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
