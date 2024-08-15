package cmdhelper

import (
	"fmt"
	"math/rand"
	"printmaker/tex"
	"printmaker/types"

	"github.com/shopspring/decimal"
)

func CmdDecimal(size int, column int, option int) {
	// 0: normal, 1: plus, 2: minus
	problems := generateDecimalProblemList(size, option)
	tex.GeneratePdf(problems, column)
}

func generateDecimalProblemList(size int, option int) []types.ProblemAnswer {
	problems := make([]types.ProblemAnswer, size)
	for i := 0; i < size; i++ {
		problems[i] = generateDecimalProblem(option)
	}
	return problems
}

func generateDecimalProblem(option int) types.ProblemAnswer {
	a := generateDecimal(10, false)
	b := generateDecimal(10, false)
	switch option {
	case 1:
		return generateDecimalProblemAnswer(a, b, 1)
	case 2:
		return generateDecimalProblemAnswer(a, b, 2)
	default:
		rand := rand.Intn(2) + 1
		return generateDecimalProblemAnswer(a, b, rand)
	}
}

func generateDecimalProblemAnswer(a decimal.Decimal, b decimal.Decimal, addsub int) types.ProblemAnswer {
	c := a.Add(b)
	if addsub == 1 {
		problem := a.String() + "+" + b.String()
		answer := c.String()
		return types.ProblemAnswer{Problem: problem, Answer: answer}
	} else {
		problem := c.String() + "-" + a.String()
		answer := b.String()
		return types.ProblemAnswer{Problem: problem, Answer: answer}

	}
}

func generateDecimal(max int, isMinus bool) decimal.Decimal {
	max -= 1
	if max < 0 {
		return decimal.NewFromInt(0)
	}
	var n int
	if isMinus {
		n = rand.Intn(2*max-1) - (max - 1)
	} else {
		n = rand.Intn(max)

	}
	m := rand.Intn(100)
	str := fmt.Sprintf("%d.%d", n, m)
	res, _ := decimal.NewFromString(str)
	return res
}
