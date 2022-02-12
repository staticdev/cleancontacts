package clean

import (
	"io"
	"log"

	"github.com/emersion/go-vcard"
	"github.com/spf13/afero"
)

type CleanError struct {
	Msg string
}

func (err CleanError) Error() string {
	return err.Msg
}

func handleError(err error, errs []error) []error {
	if err != nil {
		return append(errs, CleanError{Msg: err.Error()})
	}
	return errs
}

type Clean struct{}

func (Clean) ContactClean(fileSystem afero.Fs, fileNameIn, filePathOut string) error {
	errs := make([]error, 0)
	in, err := fileSystem.Open(fileNameIn)
	errs = handleError(err, errs)
	defer in.Close()

	out, err := fileSystem.Create(filePathOut)
	errs = handleError(err, errs)
	defer out.Close()

	dec := vcard.NewDecoder(in)
	enc := vcard.NewEncoder(out)
	for {
		card, err := dec.Decode()
		if err == io.EOF {
			break
		} else {
			errs = handleError(err, errs)
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
		errs = handleError(err, errs)
		log.Printf("%s exported\n", cleanCard.PreferredValue(vcard.FieldFormattedName))
	}
	if len(errs) != 0 {
		return errs[0]
	}
	return nil
}
