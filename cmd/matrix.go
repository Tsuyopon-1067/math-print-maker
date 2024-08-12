/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"printmaker/cmdhelper"

	"github.com/spf13/cobra"
)

// matrixCmd represents the matrix command
var matrixCmd = &cobra.Command{
	Use:   "matrix",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdhelper.CmdMatrix(size, column, dim)
	},
}

var (
	dim int
)

func init() {
	rootCmd.AddCommand(matrixCmd)

	matrixCmd.Flags().IntVarP(&size, "size", "s", 100, "Size flag")
	matrixCmd.Flags().IntVarP(&column, "column", "c", 3, "Column flag")
	matrixCmd.Flags().IntVarP(&dim, "dim", "d", 3, "Dimension flag")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matrixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// matrixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
