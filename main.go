package main

import (
	"log"

	"github.com/findsam/tbot/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
