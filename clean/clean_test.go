package clean_test

import (
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/staticdev/cleancontacts/clean"
)

var Clean = clean.Clean{}
var FakeFS = afero.NewMemMapFs()

func TestClean(t *testing.T) {
	testCases := []struct {
		name    string
		contact string
		want    string
	}{
		{
			name: "happy-path",
			contact: `BEGIN:VCARD
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
`,
			want: `BEGIN:VCARD
VERSION:3.0
FN:This Is A Full Name
N:Name;This is A;Full;;
TEL:+40 547984080
END:VCARD
`,
		},
		{
			name: "skip-no-tel",
			contact: `BEGIN:VCARD
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
`,
			want: ``,
		},
	}
	fileNameIn := "dirty-contacts.vcf"
	filePathOut := "./dirty-contact_cleaned.vcf"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			afero.WriteFile(FakeFS, fileNameIn, []byte(tc.contact), 0600)

			Clean.ContactClean(FakeFS, fileNameIn, filePathOut)
			out, _ := afero.ReadFile(FakeFS, "dirty-contact_cleaned.vcf")
			outStr := strings.Replace(string(out), "\r\n", "\n", -1)
			if outStr != tc.want {
				t.Errorf("want (%q), got (%q)", tc.want, outStr)
			}
		})
	}
}
