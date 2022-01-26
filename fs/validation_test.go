package fs

import (
	"testing"
	"testing/fstest"
)

var (
	fakeFS = fstest.MapFS{
		"readme.md":              {},
		"dirty-contacts.vcf":     {},
		"secret-folder":          {Mode: 0000},
		"secret-folder/cont.vcf": {},
	}
)

func TestFileValid(t *testing.T) {
	fileName := "dirty-contacts.vcf"
	wantOut := "dirty-contacts_cleaned.vcf"

	out, err := FileValid(fakeFS, fileName)
	if out != wantOut || err != nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}

func TestFileValidWrongExt(t *testing.T) {
	fileName := "readme.md"
	wantOut := ""

	out, err := FileValid(fakeFS, fileName)
	if out != wantOut || err == nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}

func TestFileValidInexistingFile(t *testing.T) {
	fileName := "contacts.vcf"
	wantOut := ""

	out, err := FileValid(fakeFS, fileName)
	if out != wantOut || err == nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}

// See issue: https://github.com/golang/go/issues/50787
// func TestFileValidPermissionDenied(t *testing.T) {
// 	fileName := "cont.vcf"
// 	fsys, _ := fs.Sub(fakeFS, "secret-folder")
// 	wantOut := ""

// 	out, err := FileValid(fsys, fileName)
// 	if out != wantOut || err == nil {
// 		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
// 	}
// }
