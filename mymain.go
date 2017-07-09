package main

import (
	"log"
  "fmt"

	"github.com/dynamike67/go-testing/command"
)

func main() {
  result := mydoit(7, 2)
  fmt.Println("7+2 =", result)
  fmt.Println(myhelloworld("HalloWelt!"))
	log.SetFlags(0)
	if err := command.RootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}

func mydoit(a int, b int) int {
  return a + b
}

func myhelloworld(words string) string {
  return words
}
