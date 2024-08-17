/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"printmaker/cmdhelper"

	"github.com/spf13/cobra"
)

// modCmd represents the mod command
var modCmd = &cobra.Command{
	Use:   "mod",
	Short: "Generate division problems with a remainder and answers in pdf format.",
	Long: `Generate division problems with a remainder and answers in pdf format.
For example:
	Generate 300 division problems with a remainder in 2 columns format.
	$printmaker mod --size 300 --column 2`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdhelper.CmdMod(size, column, l, r)
	},
}

func init() {
	rootCmd.AddCommand(modCmd)
	modCmd.Flags().IntVarP(&size, "size", "s", 100, "Size flag")
	modCmd.Flags().IntVarP(&column, "column", "c", 3, "Column flag")
	modCmd.Flags().IntVarP(&l, "left", "l", 2, "left digit")
	modCmd.Flags().IntVarP(&r, "right", "r", 2, "right digit")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
