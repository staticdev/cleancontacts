package clean

import (
	"testing"

	"github.com/spf13/afero"
)

var DirtyCard = `BEGIN:VCARD
VERSION:3.0
END:VCARD`
var FakeFS = afero.NewMemMapFs()

func TestRun(t *testing.T) {

	fileNameIn := "dirty-contact.vcf"
	filePathOut := "./dirty-contact_cleaned.vcf"
	wantOut := ""

	afero.WriteFile(FakeFS, fileNameIn, []byte(DirtyCard), 0600)

	Run(FakeFS, fileNameIn, filePathOut)
	out, err := afero.ReadFile(FakeFS, "dirty-contact_cleaned.vcf")
	if err != nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, out, err)
	}
}
