package cmd

import (
	"fmt"
	"io/fs"
	"testing"
	"testing/fstest"
)

var (
	fakeFS = fstest.MapFS{
		"root-folder":                    {Mode: fs.ModeDir},
		"root-folder/file-1.md":          {Data: []byte("Wrong format file")},
		"root-folder/dirty-contacts.vcf": {},
	}
)

// https://dev.to/albertodeago88/learn-golang-basics-by-creating-a-file-counter-50f1
// tests := []struct {
//     args     []string
//     function func() func(cmd *cobra.Command, args []string) error
//     output   string
// }

func TestExecute(t *testing.T) {
	fmt.Print(fakeFS)
}
