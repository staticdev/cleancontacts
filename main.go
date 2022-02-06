package main

import (
	"github.com/staticdev/cleancontacts/clean"
	"github.com/staticdev/cleancontacts/cmd"
	"github.com/staticdev/cleancontacts/fs"
)

func main() {
	fileIo := fs.FileIO{}
	cleaner := clean.Clean{}
	command := cmd.RootCmd(fileIo, cleaner)
	cmd.Execute(command)
}
