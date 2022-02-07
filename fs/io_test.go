package fs_test

import (
	"errors"
	"runtime"
	"testing"

	"github.com/spf13/afero"
	"github.com/staticdev/cleancontacts/fs"
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

var FileIO = fs.FileIO{}
var FakeFS = afero.NewOsFs()

func TestGetOutputFileName(t *testing.T) {
	testCases := []struct {
		name        string
		fileName    string
		want        string
		expectedErr error
	}{
		{
			name:     "happy-path",
			fileName: "dirty-contacts.vcf",
			want:     "dirty-contacts_cleaned.vcf",
		},
		{
			name:        "file-does-not-exist",
			fileName:    "file-does-not-exist.vcf",
			expectedErr: fs.ValidationError{Msg: "File \"file-does-not-exist.vcf\" does not exist."},
		},
		{
			name:        "wrong-extension",
			fileName:    "readme.md",
			expectedErr: fs.ValidationError{Msg: "Extension \".md\" not accepted, please use a \".vcf\" file."},
		},
	}
	for _, tc := range testCases {
		defer func() {
			FakeFS.Remove("dirty-contacts.vcf")
			FakeFS.RemoveAll("readme.md")
		}()
		afero.WriteFile(FakeFS, "dirty-contacts.vcf", []byte(""), 0600)
		afero.WriteFile(FakeFS, "readme.md", []byte(""), 0600)

		t.Run(tc.name, func(t *testing.T) {
			got, err := FileIO.GetOutputFileName(FakeFS, tc.fileName)
			if got != tc.want {
				t.Errorf("want (%q) got (%q)", tc.want, got)
			}
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want (%v) got (%v)", tc.expectedErr, err)
			}
		})
	}
}

func TestGetOutputFileNamePermissionDenied(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skipf("Unix based only")
	}

	folderName := "secret"
	fileName := "secret/cont.vcf"
	defer func() {
		FakeFS.Chmod(folderName, 0700)
		FakeFS.RemoveAll(folderName)
	}()
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

	out, err := FileIO.GetOutputFileName(FakeFS, fileName)
	if out != wantOut || err == nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}
