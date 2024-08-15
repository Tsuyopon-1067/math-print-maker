/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"printmaker/cmdhelper"

	"github.com/spf13/cobra"
)

// decimalCmd represents the decimal command
var decimalCmd = &cobra.Command{
	Use:   "decimal",
	Short: "Generate decimal addition and subtraction problems and answers in pdf format.",
	Long: `Generate decimal addition and subtraction problems and answers in pdf format.
For example:
	Generate 300 decimal addition and subtraction problems in 2 columns format.
	$printmaker decimal --size 300 --column 2`,
	Run: func(cmd *cobra.Command, args []string) {
		option := 0
		if isAddition {
			option = 1
		} else if isSubtraction {
			option = 2
		}
		cmdhelper.CmdDecimal(size, column, option)
	},
}

var (
	isAddition    bool
	isSubtraction bool
)

func init() {
	rootCmd.AddCommand(decimalCmd)

	decimalCmd.Flags().IntVarP(&size, "size", "s", 100, "Size flag")
	decimalCmd.Flags().IntVarP(&column, "column", "c", 3, "Column flag")
	decimalCmd.Flags().BoolVarP(&isAddition, "plus", "p", false, "Generate addition problems")
	decimalCmd.Flags().BoolVarP(&isSubtraction, "minus", "m", false, "Generate subtraction problems")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decimalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decimalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
