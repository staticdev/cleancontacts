package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

type ValidationError struct {
	Msg string
}

func (err ValidationError) Error() string {
	return err.Msg
}

type FileIO struct{}

func fileExists(fileSystem afero.Fs, fileName string) (bool, error) {
	_, err := fileSystem.Stat(fileName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, ValidationError{Msg: fmt.Sprintf("File %q does not exist.", fileName)}
	}
	return false, err
}

func validExtension(gotExtension string) error {
	wantExtension := ".vcf"
	if gotExtension != wantExtension {
		return ValidationError{Msg: fmt.Sprintf("Extension %q not accepted, please use a %q file.", gotExtension, wantExtension)}
	}
	return nil
}

func getOutputFileName(fileName, fileExtension string) string {
	return strings.TrimSuffix(filepath.Base(fileName), fileExtension) + "_cleaned" + fileExtension
}

func (FileIO) GetOutputFileName(fileSystem afero.Fs, fileName string) (string, error) {
	fileExists, err := fileExists(fileSystem, fileName)
	if !fileExists || err != nil {
		return "", err
	}

	fileExtension := filepath.Ext(fileName)
	err = validExtension(fileExtension)
	if err != nil {
		return "", err
	}

	outFileName := getOutputFileName(fileName, fileExtension)

	return outFileName, nil
}
