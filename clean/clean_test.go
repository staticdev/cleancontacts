package clean

import (
	"strings"
	"testing"

	"github.com/spf13/afero"
)

var FakeFS = afero.NewMemMapFs()

func TestRun(t *testing.T) {
	var dirtyCard = `BEGIN:VCARD
VERSION:3.0
FN:This Is A Full Name
N:Name;This is A;Full;;
item1.EMAIL;TYPE=INTERNET:some@email.com
item1.X-ABLabel:
TEL;TYPE=CELL:+40 547984080
item3.ADR:;;911 Omg Straat;;Hakooken - PR;;SW;911 Omg Straat\nHakooken - PR\nSW
item3.X-ABLabel:
NOTE:Some notes\n\nmore notes
CATEGORIES:myContacts
END:VCARD
`
	fileNameIn := "dirty-contact.vcf"
	filePathOut := "./dirty-contact_cleaned.vcf"
	wantOut := `BEGIN:VCARD
VERSION:3.0
FN:This Is A Full Name
N:Name;This is A;Full;;
TEL:+40 547984080
END:VCARD
`

	afero.WriteFile(FakeFS, fileNameIn, []byte(dirtyCard), 0600)

	Run(FakeFS, fileNameIn, filePathOut)
	out, err := afero.ReadFile(FakeFS, "dirty-contact_cleaned.vcf")
	outStr := strings.Replace(string(out), "\r\n", "\n", -1)
	if outStr != wantOut || err != nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, outStr, err)
	}
}

func TestRunSkipNoTel(t *testing.T) {
	var dirtyCard = `BEGIN:VCARD
VERSION:3.0
FN:This Is A Full Name
N:Name;This is A;Full;;
item1.EMAIL;TYPE=INTERNET:some@email.com
item1.X-ABLabel:
item3.ADR:;;911 Omg Straat;;Hakooken - PR;;SW;911 Omg Straat\nHakooken - PR\nSW
item3.X-ABLabel:
NOTE:Some notes\n\nmore notes
CATEGORIES:myContacts
END:VCARD
`
	fileNameIn := "dirty-contact.vcf"
	filePathOut := "./dirty-contact_cleaned.vcf"
	wantOut := ``

	afero.WriteFile(FakeFS, fileNameIn, []byte(dirtyCard), 0600)

	Run(FakeFS, fileNameIn, filePathOut)
	out, err := afero.ReadFile(FakeFS, "dirty-contact_cleaned.vcf")
	outStr := strings.Replace(string(out), "\r\n", "\n", -1)
	if outStr != wantOut || err != nil {
		t.Errorf("want %q, nil got %q, %v", wantOut, outStr, err)
	}
}
