package fs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func fileExists(fileSystem fs.FS, fileName string) (bool, error) {
	_, err := fs.Stat(fileSystem, fileName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, fmt.Errorf("File %q does not exist.", fileName)
	}
	return false, err
}

func getOutputFileName(fileName, fileExtension string) string {
	return strings.TrimSuffix(filepath.Base(fileName), fileExtension) + "_cleaned" + fileExtension
}

func FileValid(fileSystem fs.FS, fileName string) (string, error) {
	fileExists, err := fileExists(fileSystem, fileName)
	if !fileExists || err != nil {
		return "", err
	}

	fileExtension := filepath.Ext(fileName)
	if fileExtension != ".vcf" {
		return "", fmt.Errorf("Extension %v not accepted. Please use a .vcf file.", fileExtension)
	}

	outFileName := getOutputFileName(fileName, fileExtension)

	return outFileName, nil
}
