package cmdhelper

import (
	"printmaker/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateIntegerAddition22(t *testing.T) {
	a := 23
	b := 12
	problem := "23+12"
	answer := "35"
	expected := types.ProblemAnswer{Problem: problem, Answer: answer}

	actual := generateIntegerAdditionSubtractionProblemAnswer(a, b, 1)
	assert.Equal(t, expected, actual)
}

func TestGenerateIntegerSubtraction22(t *testing.T) {
	a := 23
	b := 12
	problem := "23-12"
	answer := "11"
	expected := types.ProblemAnswer{Problem: problem, Answer: answer}

	actual := generateIntegerAdditionSubtractionProblemAnswer(a, b, 2)
	assert.Equal(t, expected, actual)
}
