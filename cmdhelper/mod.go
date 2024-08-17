package cmdhelper

import (
	"fmt"
	"printmaker/tex"
	"printmaker/types"
)

func CmdMod(size int, column int, l int, r int) {
	problems := generateModProblemList(size, l, r)
	tex.GeneratePdf(problems, column)
}

func generateModProblemList(size int, l int, r int) []types.ProblemAnswer {
	problems := make([]types.ProblemAnswer, size)
	for i := 0; i < size; i++ {
		problems[i] = generateModProblem(l, r)
	}
	return problems
}

func generateModProblem(l int, r int) types.ProblemAnswer {
	a := generateDigitInteger(l)
	b := generateDigitInteger(r)
	if a < b {
		a, b = b, a
	}
	return generateModAdditionSubtractionProblemAnswer(a, b)
}

func generateModAdditionSubtractionProblemAnswer(a int, b int) types.ProblemAnswer {
	mod := a % b
	quotient := a / b
	problem := fmt.Sprint(a) + " \\div " + fmt.Sprint(b)
	answer := fmt.Sprint(quotient) + " \\cdots " + fmt.Sprint(mod)
	return types.ProblemAnswer{Problem: problem, Answer: answer}
}
