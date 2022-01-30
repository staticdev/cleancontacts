package fs

import (
	"testing"

	"github.com/spf13/afero"
)

// Problem with MapFS, using afero
// See issue: https://github.com/golang/go/issues/50787
// var (
// 	FakeFS = fstest.MapFS{
// 		"readme.md":              {},
// 		"dirty-contacts.vcf":     {},
// 		"secret-folder":          {Mode: 0000},
// 		"secret-folder/cont.vcf": {},
// 	}
// )

// Problem with NewMemMapFs
// See issue: https://github.com/spf13/afero/issues/335
// var FakeFS = afero.NewMemMapFs()

var FakeFS = afero.NewOsFs()

func TestFileValid(t *testing.T) {
	fileName := "dirty-contacts.vcf"
	afero.WriteFile(FakeFS, fileName, []byte(""), 0600)
	wantOut := "dirty-contacts_cleaned.vcf"

	out, err := FileValid(FakeFS, fileName)
	if out != wantOut || err != nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}

func TestFileValidWrongExt(t *testing.T) {
	fileName := "readme.md"
	afero.WriteFile(FakeFS, fileName, []byte(""), 0600)
	wantOut := ""

	out, err := FileValid(FakeFS, fileName)
	if out != wantOut || err == nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}

func TestFileValidInexistingFile(t *testing.T) {
	fileName := "contacts.vcf"
	wantOut := ""

	out, err := FileValid(FakeFS, fileName)
	if out != wantOut || err == nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}

func TestFileValidPermissionDenied(t *testing.T) {
	folderName := "secret"
	fileName := "secret/cont.vcf"
	err := FakeFS.Mkdir(folderName, 0700) // temporary permission to write
	if err != nil {
		t.Errorf("%v", err)
	}
	err = afero.WriteFile(FakeFS, fileName, []byte(""), 0700)
	if err != nil {
		t.Errorf("%v", err)
	}
	FakeFS.Chmod(folderName, 0000)
	if err != nil {
		t.Errorf("%v", err)
	}
	wantOut := ""

	out, err := FileValid(FakeFS, fileName)
	if out != wantOut || err == nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}
