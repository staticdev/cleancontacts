package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/staticdev/cleancontacts/clean"
	"github.com/staticdev/cleancontacts/fs"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func RootCmd() (rootCmd *cobra.Command, fileName string) {
	var file string

	cmd := &cobra.Command{
		Use:   "cleancontacts <filepath>.vcf",
		Short: "Cleanup your phone contacts to prevent apps for having access to all details of your contacts.",
		Long: `Do not want to share all your contact info to mobile apps? This software is for you!

Export your contacts in VCard format and run the program. BANG! You have a new VCard file with cleaned contacts with just their names and telephones.`,
		Version: "0.1.0",
		Args: func(cmd *cobra.Command, args []string) error {
			if file == "" && len(args) < 1 {
				return errors.New("needs 1 arg")
			}
			return nil
		},
		Example: `cleancontacts contacts.vcf
cleancontacts /path/contacts.vcf`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(file)
			var filePath string

			if file != "" {
				filePath = file
			} else {
				filePath = args[0]
			}

			dir, fileName := filepath.Split(filePath)
			fsys := afero.NewOsFs()
			fileNameOut, err := fs.FileValid(fsys, fileName)
			if err != nil {
				return err
			}
			filePathOut := filepath.Join(dir, fileNameOut)

			clean.Run(fsys, fileName, filePathOut)
			return nil
		},
	}
	return cmd, file
}

func Execute(cmd *cobra.Command, file string) error {
	return cmd.Execute()
}
