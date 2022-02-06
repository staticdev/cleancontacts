package clean

import (
	"io"
	"log"

	"github.com/emersion/go-vcard"
	"github.com/spf13/afero"
)

type Clean struct{}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (Clean) ContactClean(fileSystem afero.Fs, fileNameIn, filePathOut string) {
	in, err := fileSystem.Open(fileNameIn)
	handleError(err)
	defer in.Close()

	out, err := fileSystem.Create(filePathOut)
	handleError(err)
	defer out.Close()

	dec := vcard.NewDecoder(in)
	enc := vcard.NewEncoder(out)
	for {
		card, err := dec.Decode()
		if err == io.EOF {
			break
		} else {
			handleError(err)
		}

		var ns = card.Values("N")
		var tels = card.Values("TEL")
		// skip contacts with no name or tel
		if len(ns) == 0 || len(tels) == 0 {
			continue
		}
		var versions = card.Values("VERSION")
		var fns = card.Values("FN")
		var cleanCard vcard.Card = make(vcard.Card)
		for _, version := range versions {
			cleanCard.AddValue("VERSION", version)
		}
		for _, n := range ns {
			cleanCard.AddValue("N", n)
		}
		for _, fn := range fns {
			cleanCard.AddValue("FN", fn)
		}
		for _, tel := range tels {
			cleanCard.AddValue("TEL", tel)
		}
		err = enc.Encode(cleanCard)
		handleError(err)
		log.Printf("%s exported\n", cleanCard.PreferredValue(vcard.FieldFormattedName))
	}
}
