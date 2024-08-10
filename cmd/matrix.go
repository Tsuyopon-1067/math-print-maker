/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"printmaker/tex"
	"printmaker/types"

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
		problems := generateMatrixProblemList(size)
		tex.GeneratePdf(problems, column)
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

func generateMatrixProblemList(size int) []types.ProblemAnswer {
	problems := make([]types.ProblemAnswer, size)
	for i := 0; i < size; i++ {
		problems[i] = generateMatrixProblem()
	}
	return problems
}

func generateMatrixProblem() types.ProblemAnswer {
	a := generateMatrix(dim)
	b := generateMatrix(dim)

	return generateMatrixProductProblemAnswer(a, b)
}

func generateMatrixProductProblemAnswer(a types.IntMatrix, b types.IntMatrix) types.ProblemAnswer {
	c := a.Product(b)
	problem := a.ToTexString() + "\\times\n" + b.ToTexString()
	answer := c.ToTexString()
	return types.ProblemAnswer{Problem: problem, Answer: answer}
}

func generateMatrix(dim int) types.IntMatrix {
	array := make([][]int, dim)
	for i := range array {
		array[i] = make([]int, dim)
		for j := range array {
			array[i][j] = randRange(0, 10)
		}
	}
	return types.CreateIntMatrix(array)
}
