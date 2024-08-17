package cmdhelper

import (
	"printmaker/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateModOk(t *testing.T) {
	a := 43
	b := 21
	problem := "43 \\div 21"
	answer := "2 あまり 1"
	expected := types.ProblemAnswer{Problem: problem, Answer: answer}

	actual := generateModAdditionSubtractionProblemAnswer(a, b)
	assert.Equal(t, expected, actual)
}
