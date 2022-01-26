package main

import (
	"log"

	"github.com/staticdev/cleancontacts/cmd"
)

func main() {
	command, file := cmd.RootCmd()
	err := cmd.Execute(command, file)
	if err != nil {
		log.Fatal(err)
	}
}
