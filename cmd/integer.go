/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"printmaker/cmdhelper"

	"github.com/spf13/cobra"
)

// integerCmd represents the integer command
var integerCmd = &cobra.Command{
	Use:   "integer",
	Short: "Generate integer addition and subtraction problems and answers in pdf format.",
	Long: `Generate integer addition and subtraction problems and answers in pdf format. You can specify the number of digits.
For example:
	Generate 300 integer addition and subtraction problems in 2 columns format.
	$printmaker integer --size 300 --column 2`,
	Run: func(cmd *cobra.Command, args []string) {
		option := 0
		if isAddition {
			option = 1
		} else if isSubtraction {
			option = 2
		}
		cmdhelper.CmdInteger(size, column, option, l, r)
	},
}

var (
	l int
	r int
)

func init() {
	rootCmd.AddCommand(integerCmd)
	integerCmd.Flags().IntVarP(&size, "size", "s", 100, "Size flag")
	integerCmd.Flags().IntVarP(&column, "column", "c", 3, "Column flag")
	integerCmd.Flags().BoolVarP(&isAddition, "plus", "p", false, "Generate addition problems")
	integerCmd.Flags().BoolVarP(&isSubtraction, "minus", "m", false, "Generate subtraction problems")
	integerCmd.Flags().IntVarP(&l, "left", "l", 2, "left digit")
	integerCmd.Flags().IntVarP(&r, "right", "r", 2, "right digit")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// integerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// integerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
