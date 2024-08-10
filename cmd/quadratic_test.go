package cmd

import (
	"printmaker/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type cofficients struct {
	alpha, beta, gamma int
}

func TestGenerateProblemList(t *testing.T) {
	coefficient := []cofficients{
		{alpha: 2, beta: -3, gamma: 3},
		{alpha: -6, beta: -4, gamma: 1},
		{alpha: 5, beta: 7, gamma: 1},
		{alpha: -8, beta: 9, gamma: 1},
		{alpha: 9, beta: 9, gamma: 1},
	}
	expected := []types.ProblemAnswer{
		{Problem: "3x^2+3x-18=0", Answer: "x = -3, 2"},
		{Problem: "x^2+10x+24=0", Answer: "x = -6, -4"},
		{Problem: "x^2-12x+35=0", Answer: "x = 5, 7"},
		{Problem: "x^2-x-72=0", Answer: "x = -8, 9"},
		{Problem: "x^2-18x+81=0", Answer: "x = 9"},
	}

	actual := make([]types.ProblemAnswer, len(expected))
	for i := range coefficient {
		alpha, beta, gamma := coefficient[i].alpha, coefficient[i].beta, coefficient[i].gamma
		actual[i] = generateProblemAnswer(alpha, beta, gamma)
	}
	assert.Equal(t, expected, actual)
}
