package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "replace a sequence of bytes in a file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		frombin, err := hex.DecodeString(fromhex)
		if err != nil {
			fmt.Println("Unable decode sequence for replacing: ", err.Error())
			return
		}

		tobin, err := hex.DecodeString(tohex)
		if err != nil {
			fmt.Println("Unable decode sequence to replace with: ", err.Error())
			return
		}

		if len(frombin) != len(tobin) {
			fmt.Println("Replacing sequence and replacement sequence must be of equal length")
			return
		}

		if !isFileExist(source) {
			fmt.Println("Source file does not exist")
			return
		}

		if isFileExist(target) && !force {
			fmt.Println("Target file already exists. Please use --force to overwrite")
			return
		}

		input, err := ioutil.ReadFile(source)
		if err != nil {
			fmt.Println("Unable to read source file: ", err.Error())
			return
		}
		output := input

		fmt.Printf("Original sequence: %x \n", frombin)
		fmt.Printf("New sequence: %x \n", tobin)

		for i := 0; i < len(input)-len(frombin); i++ {
			if input[i] == frombin[0] {
				if isSequenceEqual(input[i:i+len(frombin)], frombin) {
					fmt.Printf("Found sequence at position %d \n", i)
					for j := 0; j < len(frombin); j++ {
						output[i+j] = tobin[j]
					}
				}
			}
		}

		if dryrun {
			fmt.Println("Dry run. No changes made")
			return
		}

		out, err := os.Create(target)
		if err != nil {
			fmt.Println("Unable to create target file: ", err.Error())
			return
		}
		defer out.Close()

		_, err = io.Copy(out, bytes.NewReader(output))
		if err != nil {
			fmt.Println("Unable to write to target file: ", err.Error())
			return
		}
		fmt.Println("Target file written successfully")
	},
}

func init() {
	replaceCmd.Flags().StringVarP(&source, "source", "i", "", "source (input) binary file")
	replaceCmd.Flags().StringVarP(&target, "target", "o", "", "target (output) binary file")
	replaceCmd.Flags().StringVarP(&fromhex, "from", "f", "", "binary sequence to replace")
	replaceCmd.Flags().StringVarP(&tohex, "to", "t", "", "binary sequence to replace with")
	replaceCmd.Flags().BoolVarP(&force, "force", "", false, "overwrite target file if it exists")
	replaceCmd.MarkFlagRequired("source")
	replaceCmd.MarkFlagRequired("target")
	replaceCmd.MarkFlagRequired("from")
	replaceCmd.MarkFlagRequired("to")
	rootCmd.AddCommand(replaceCmd)
}
