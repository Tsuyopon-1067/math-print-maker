/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"math/rand"
	"printmaker/format"
	"printmaker/tex"
	"printmaker/types"

	"github.com/spf13/cobra"
)

// quadraticCmd represents the quadratic command
var quadraticCmd = &cobra.Command{
	Use:   "quadratic",
	Short: "Generate quadratic equation problems and answers in pdf format",
	Long: `Generate quadratic equation problems and answers in pdf format.
	You can specify the number of problems to generate by passing the size flag and columns of generated pdf file.`,
	Run: func(cmd *cobra.Command, args []string) {
		problems := generateProblemList(size)
		tex.GeneratePdf(problems, column)
	},
}

var (
	size   int
	column int
)

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

func generateProblemList(size int) []types.ProblemAnswer {
	problems := make([]types.ProblemAnswer, size)
	for i := 0; i < size; i++ {
		problems[i] = generateProblem()
	}
	return problems
}

func generateProblem() types.ProblemAnswer {
	alpaha := randRange(-9, 10)
	beta := randRange(-9, 10)
	gamma := randRange(1, 3)

	coefficientA := gamma
	coefficientB := -gamma * (alpaha + beta)
	coefficientC := alpaha * beta * gamma

	problem := format.CharacterExpression(coefficientA, "x^2", true) +
		format.CharacterExpression(coefficientB, "x", false) +
		format.CharacterExpression(coefficientC, "", false) +
		"=0"
	answer := format.AnswerExpression("x", alpaha, beta)
	return types.ProblemAnswer{Problem: problem, Answer: answer}
}

func randRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
