/*
Copyright © 2022 staticdev
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/staticdev/cleancontacts/clean"

	"github.com/spf13/cobra"
)

var File string

var rootCmd = &cobra.Command{
	Use:   "cleancontacts",
	Short: "Cleanup your phone contacts to prevent apps for having access to all details of your contacts.",
	Long: `Do not want to share all your contact info to mobile apps? This software is for you!

Export your contacts in VCard format and run the program. BANG! You have a new VCard file with cleaned contacts with just their names and telephones.`,
	Version: "0.1.0",
	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) < 1 {
			return errors.New("needs 1 arg")
		}
		return nil
	},
	Example: `cleancontacts contacts.vcf
cleancontacts /Downloads/contacts.vcf`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(File)
		var fileIn string
		var argument string

		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}

		fileExists, err := clean.FileExists(argument)
		if err != nil {
			fmt.Println(err)
		}
		if fileExists {
			fileIn, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())

			}
		} else {
			fmt.Printf("File %v does not exist.", argument)
			return
		}

		fileExtension := filepath.Ext(fileIn)
		if fileExtension != ".vcf" {
			fmt.Printf("Extension %v not accepted. Please use a .vcf file.", fileExtension)
			return
		}

		fileOut := clean.GetOutputPath(fileIn, fileExtension)

		clean.Run(fileIn, fileOut)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "a file name/path to VCard contacts")
}
