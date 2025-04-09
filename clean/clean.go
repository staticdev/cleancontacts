package clean

import (
	"io"
	"log"

	"github.com/emersion/go-vcard"
	"github.com/spf13/afero"
)

type CleanerError struct {
	Msg string
}

func (err CleanerError) Error() string {
	return err.Msg
}

func handleError(err error, errs []error) []error {
	if err != nil {
		return append(errs, CleanerError{Msg: err.Error()})
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
		}
		errs = handleError(err, errs)
		ns := card.Values("N")
		tels := card.Values("TEL")

		// skip contacts with no name or tel
		if len(ns) == 0 || len(tels) == 0 {
			continue
		}
		versions := card.Values("VERSION")
		fns := card.Values("FN")
		cleanCard := make(vcard.Card)
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
