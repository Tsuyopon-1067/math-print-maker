package cmdhelper

import (
	"printmaker/types"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAddition(t *testing.T) {
	a := decimal.NewFromFloat(1.23)
	b := decimal.NewFromFloat(4.56)
	problem := "1.23+4.56"
	answer := "5.79"
	expected := types.ProblemAnswer{Problem: problem, Answer: answer}

	actual := generateDecimalProblemAnswer(a, b, 1)
	assert.Equal(t, expected, actual)
}

func TestGenerateSubtraction(t *testing.T) {
	a := decimal.NewFromFloat(1.23)
	b := decimal.NewFromFloat(4.56)
	problem := "5.79-1.23"
	answer := "4.56"
	expected := types.ProblemAnswer{Problem: problem, Answer: answer}

	actual := generateDecimalProblemAnswer(a, b, 2)
	assert.Equal(t, expected, actual)
}
