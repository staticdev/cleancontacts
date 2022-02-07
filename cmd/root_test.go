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

type FailingFileIO struct{}

func (FailingFileIO) GetOutputFileName(fileSystem afero.Fs, fileName string) (string, error) {
	return "", cmd.CommandError{"some file io error"}
}

type FakeClean struct{}

func (FakeClean) ContactClean(fileSystem afero.Fs, fileNameIn, filePathOut string) {}

func TestExecute(t *testing.T) {
	testCases := []struct {
		name        string
		args        []string
		expectedErr error
		fileIoer    cmd.FileIoer
	}{
		{
			name:     "happy-path",
			args:     []string{"contacts.vcf"},
			fileIoer: FakeFileIO{},
		},
		{
			name:        "no-args",
			expectedErr: cmd.CommandError{Msg: "Contact file argument not provided."},
			fileIoer:    FakeFileIO{},
		},
		{
			name:        "fileio-error",
			args:        []string{"contacts.vcf"},
			expectedErr: cmd.CommandError{"some file io error"},
			fileIoer:    FailingFileIO{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := cmd.RootCmd(tc.fileIoer, FakeClean{})
			root.SetArgs(tc.args)
			err := cmd.Execute(root)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want (%v) got (%v)", tc.expectedErr, err)
			}
		})
	}
}
