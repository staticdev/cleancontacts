package clean

import (
	"testing"

	"github.com/spf13/afero"
)

var DirtyCard = `BEGIN:VCARD
VERSION:3.0
FN:Alcione Marques Dos Santos
N:Santos;Alcione Marques;Dos;;
item1.EMAIL;TYPE=INTERNET:alcimarque32@gmail.com
item1.X-ABLabel:
item2.EMAIL;TYPE=INTERNET:alcimarques32@hotmail.com
item2.X-ABLabel:
TEL;TYPE=CELL:+55 31984136833
item3.ADR:;;991 Avenida Mauro Nunes Moreira;;Ibirité - MG;;BR;991 Avenida M
auro Nunes Moreira\nIbirité - MG\nBR
item3.X-ABLabel:
NOTE:CPF\: 039377886-03\n\nDados bancários\:\nBanco Santander\nAg\: 4177\nC
ta\: 1065478-0
CATEGORIES:myContacts
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
