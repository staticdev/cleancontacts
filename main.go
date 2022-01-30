package main

import (
	"github.com/staticdev/cleancontacts/cmd"
)

func main() {
	command, file := cmd.RootCmd()
	cmd.Execute(command, file)
}
