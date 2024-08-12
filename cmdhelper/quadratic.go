package cmdhelper

import (
	"math/rand"
	"printmaker/format"
	"printmaker/tex"
	"printmaker/types"
)

func CmdQuadratic(size int, column int) {
	problems := generateQuadraticProblemList(size)
	tex.GeneratePdf(problems, column)
}

func generateQuadraticProblemList(size int) []types.ProblemAnswer {
	problems := make([]types.ProblemAnswer, size)
	for i := 0; i < size; i++ {
		problems[i] = generateQuadraticProblem()
	}
	return problems
}

func generateQuadraticProblem() types.ProblemAnswer {
	alpha := randRange(-9, 10)
	beta := randRange(-9, 10)
	gamma := randRange(1, 3)

	return generateQuadraticProblemAnswer(alpha, beta, gamma)
}

func generateQuadraticProblemAnswer(alpha int, beta int, gamma int) types.ProblemAnswer {
	coefficientA := gamma
	coefficientB := -gamma * (alpha + beta)
	coefficientC := alpha * beta * gamma

	problem := format.CharacterExpression(coefficientA, "x^2", true) +
		format.CharacterExpression(coefficientB, "x", false) +
		format.CharacterExpression(coefficientC, "", false) +
		"=0"
	answer := format.AnswerExpression("x", alpha, beta)
	return types.ProblemAnswer{Problem: problem, Answer: answer}
}

func randRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
