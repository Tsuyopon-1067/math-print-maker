/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"printmaker/cmdhelper"

	"github.com/spf13/cobra"
)

// quadraticCmd represents the quadratic command
var quadraticCmd = &cobra.Command{
	Use:   "quadratic",
	Short: "Generate quadratic equation problems and answers in pdf format",
	Long: `Generate quadratic equation problems and answers in pdf format.
	You can specify the number of problems to generate by passing the size flag and columns of generated pdf file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdhelper.CmdQuadratic(size, column)
	},
}

func init() {
	rootCmd.AddCommand(quadraticCmd)

	quadraticCmd.Flags().IntVarP(&size, "size", "s", 100, "Size flag")
	quadraticCmd.Flags().IntVarP(&column, "column", "c", 3, "Column flag")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quadraticCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quadraticCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
