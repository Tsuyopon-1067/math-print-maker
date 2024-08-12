package cmdhelper

import (
	"printmaker/tex"
	"printmaker/types"
)

func CmdMatrix(size int, column int, dim int) {
	problems := generateMatrixProblemList(size, dim)
	tex.GeneratePdf(problems, column)
}

func generateMatrixProblemList(size int, dim int) []types.ProblemAnswer {
	problems := make([]types.ProblemAnswer, size)
	for i := 0; i < size; i++ {
		problems[i] = generateMatrixProblem(dim)
	}
	return problems
}

func generateMatrixProblem(dim int) types.ProblemAnswer {
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
