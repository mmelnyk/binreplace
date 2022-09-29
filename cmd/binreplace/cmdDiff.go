package main

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "find differences between two binary files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !isFileExist(source) {
			fmt.Println("Source file does not exist")
			return
		}
		if !isFileExist(target) {
			fmt.Println("Target file does not exist")
			return
		}
		input, err := ioutil.ReadFile(source)
		if err != nil {
			fmt.Println("Unable to read source file: ", err.Error())
			return
		}
		output, err := ioutil.ReadFile(target)
		if err != nil {
			fmt.Println("Unable to read target file: ", err.Error())
			return
		}
		if len(input) != len(output) {
			fmt.Println("Source and target files are of different length")
			return
		}
		for i := 0; i < len(input); i++ {
			if input[i] != output[i] {
				fmt.Printf("Difference found at position %d \n", i)
				j := 0
				dist := 0
				for ; i+j < len(input); j++ {
					if input[i+j] == output[i+j] {
						dist++
						if dist > mindistance {
							break
						}
					} else {
						dist = 0
					}
				}
				if dist > 0 { // reduce distance by 1 to compensate for the last increment
					dist--
				}
				from := input[i : i+j-dist]
				to := output[i : i+j-dist]
				fmt.Printf(" From sequence: %x \n", from)
				fmt.Printf(" To sequence: %x \n", to)
				i += j
			}
		}
	},
}

func init() {
	diffCmd.Flags().StringVarP(&source, "source", "i", "", "source (input) binary file for comparison")
	diffCmd.Flags().StringVarP(&target, "target", "o", "", "target (output) binary file for comparison")
	diffCmd.Flags().IntVarP(&mindistance, "mindistance", "m", 4, "minimum distance between two differences")
	diffCmd.MarkFlagRequired("source")
	diffCmd.MarkFlagRequired("target")
	rootCmd.AddCommand(diffCmd)
}
