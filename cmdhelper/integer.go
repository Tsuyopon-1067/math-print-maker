package cmdhelper

import (
	"fmt"
	"math/rand"
	"printmaker/tex"
	"printmaker/types"
)

func CmdInteger(size int, column int, option int, l int, r int) {
	// option: (0: normal, 1: plus, 2: minus)
	problems := generateIntegerProblemList(size, option, l, r)
	tex.GeneratePdf(problems, column)
}

func generateIntegerProblemList(size int, option int, l int, r int) []types.ProblemAnswer {
	problems := make([]types.ProblemAnswer, size)
	for i := 0; i < size; i++ {
		problems[i] = generateIntegerProblem(option, l, r)
	}
	return problems
}

func generateIntegerProblem(option int, l int, r int) types.ProblemAnswer {
	a := generateDigitInteger(l)
	b := generateDigitInteger(r)
	switch option {
	case 1:
		return generateIntegerAdditionSubtractionProblemAnswer(a, b, 1)
	case 2:
		if a < b {
			a, b = b, a
		}
		return generateIntegerAdditionSubtractionProblemAnswer(a, b, 2)
	default:
		rand := rand.Intn(2) + 1
		return generateIntegerAdditionSubtractionProblemAnswer(a, b, rand)
	}
}

func generateIntegerAdditionSubtractionProblemAnswer(a int, b int, addSub int) types.ProblemAnswer {
	var problem, answer string
	switch addSub {
	case 1:
		problem = fmt.Sprint(a) + "+" + fmt.Sprint(b)
		answer = fmt.Sprint(a + b)
	case 2:
		problem = fmt.Sprint(a) + "-" + fmt.Sprint(b)
		answer = fmt.Sprint(a - b)
	default:
		problem = fmt.Sprint(a) + "+" + fmt.Sprint(b)
		answer = fmt.Sprint(a + b)
	}
	return types.ProblemAnswer{Problem: problem, Answer: answer}
}

func generateDigitInteger(digit int) int {
	max := 1
	for i := 0; i < digit; i++ {
		max *= 10
	}
	min := max / 10
	return rand.Intn(max-min) + min
}
