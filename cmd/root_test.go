package cmd_test

import (
	"errors"
	"testing"

	"github.com/spf13/afero"
	"github.com/staticdev/cleancontacts/cmd"
)

type FakeFileIO struct{}

func (FakeFileIO) GetOutputFileName(_ afero.Fs, _ string) (string, error) {
	return "", nil
}

type FailingFileIO struct{}

func (FailingFileIO) GetOutputFileName(_ afero.Fs, _ string) (string, error) {
	return "", cmd.CommandError{Msg: "some file io error"}
}

type FakeClean struct{}

func (FakeClean) ContactClean(_ afero.Fs, _, _ string) error {
	return nil
}

type FailingClean struct{}

func (FailingClean) ContactClean(_ afero.Fs, _, _ string) error {
	return cmd.CommandError{Msg: "some clean error"}
}

func TestExecute(t *testing.T) {
	testCases := []struct {
		name        string
		args        []string
		expectedErr error
		fileIoer    cmd.FileIoer
		cleaner     cmd.ContactCleaner
	}{
		{
			name:     "happy-path",
			args:     []string{"contacts.vcf"},
			fileIoer: FakeFileIO{},
			cleaner:  FakeClean{},
		},
		{
			name:        "no-args",
			expectedErr: cmd.CommandError{Msg: "Contact file argument not provided."},
			fileIoer:    FakeFileIO{},
			cleaner:     FakeClean{},
		},
		{
			name:        "fileio-error",
			args:        []string{"contacts.vcf"},
			expectedErr: cmd.CommandError{Msg: "some file io error"},
			fileIoer:    FailingFileIO{},
			cleaner:     FakeClean{},
		},
		{
			name:        "clean-error",
			args:        []string{"contacts.vcf"},
			expectedErr: cmd.CommandError{Msg: "some clean error"},
			fileIoer:    FakeFileIO{},
			cleaner:     FailingClean{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := cmd.RootCmd(tc.fileIoer, tc.cleaner)
			root.SetArgs(tc.args)
			err := cmd.Execute(root)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want (%v) got (%v)", tc.expectedErr, err)
			}
		})
	}
}
