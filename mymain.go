package main

import (
	"log"

	"github.com/dynamike67/go-testing/command"
)

func main() {
	log.SetFlags(0)
	if err := command.RootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
