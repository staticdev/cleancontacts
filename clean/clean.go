package clean

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/emersion/go-vcard"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetOutputPath(fileName, fileExtension string) string {
	return strings.TrimSuffix(filepath.Base(fileName), fileExtension) + "_cleaned" + fileExtension
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Run(fileIn, fileOut string) {
	in, err := os.Open(fileIn)
	handleError(err)
	defer in.Close()

	out, err := os.Create(fileOut)
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
		if len(ns) == 0 {
			continue
		}
		var versions = card.Values("VERSION")
		var fns = card.Values("FN")
		var tels = card.Values("TEL")
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
