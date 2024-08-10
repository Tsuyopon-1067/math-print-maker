package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSucNormal(t *testing.T) {
	actual := AnswerExpression("x", 1, 2, 3)
	expected := "x = 1, 2, 3"

	assert.Equal(t, expected, actual)
}

func TestSucOneItem(t *testing.T) {
	actual := AnswerExpression("x", 1)
	expected := "x = 1"

	assert.Equal(t, expected, actual)
}

func TestSucDouble(t *testing.T) {
	actual := AnswerExpression("x", 1, 1)
	expected := "x = 1"

	assert.Equal(t, expected, actual)
}

func TestSucDouble3(t *testing.T) {
	actual := AnswerExpression("x", 1, 1, 2, 3, 3, 4, 5)
	expected := "x = 1, 2, 3, 4, 5"

	assert.Equal(t, expected, actual)
}
