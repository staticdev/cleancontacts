package cmd_test

import (
	"errors"
	"testing"

	"github.com/spf13/afero"
	"github.com/staticdev/cleancontacts/cmd"
)

type FakeFileIO struct{}

func (FakeFileIO) GetOutputFileName(fileSystem afero.Fs, fileName string) (string, error) {
	return "", nil
}

// type FailingFileIO struct{}

// func (FailingFileIO) GetOutputFileName(fileSystem afero.Fs, fileName string) (string, error) {
// 	return "", errors.New("some file io error")
// }

type FakeClean struct{}

func (FakeClean) ContactClean(fileSystem afero.Fs, fileNameIn, filePathOut string) {}

func TestExecute(t *testing.T) {
	testCases := []struct {
		name        string
		args        []string
		expectedErr error
	}{
		{
			name: "happy-path",
			args: []string{"contacts.vcf"},
		},
		{
			name:        "no-args",
			expectedErr: cmd.CommandError{Msg: "Contact file argument not provided."},
		},
		// TODO: simulate idiomatically how FileIO should throw an error
		// {
		// 	name:        "fileio-error",
		// 	expectedErr: errors.New("some file io error"),
		// 	fileIoer:    FailingFileIO{},
		// },
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := cmd.RootCmd(FakeFileIO{}, FakeClean{})
			root.SetArgs(tc.args)
			err := cmd.Execute(root)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want (%v) got (%v)", tc.expectedErr, err)
			}
		})
	}
}
