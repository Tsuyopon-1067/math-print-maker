package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstMinusOne(t *testing.T) {
	actual := CharacterExpression(-1, "x", true)
	expected := "-x"

	assert.Equal(t, expected, actual)
}

func TestFirstOne(t *testing.T) {
	actual := CharacterExpression(1, "x", true)
	expected := "x"

	assert.Equal(t, expected, actual)
}

func TestFirstPlusOne(t *testing.T) {
	actual := CharacterExpression(+1, "x", true)
	expected := "x"

	assert.Equal(t, expected, actual)
}

func TestSecondMinusOne(t *testing.T) {
	actual := CharacterExpression(-1, "x", false)
	expected := "-x"

	assert.Equal(t, expected, actual)
}

func TestSecondOne(t *testing.T) {
	actual := CharacterExpression(1, "x", false)
	expected := "+x"

	assert.Equal(t, expected, actual)
}

func TestSecondPlusOne(t *testing.T) {
	actual := CharacterExpression(+1, "x", false)
	expected := "+x"

	assert.Equal(t, expected, actual)
}

func TestFirstOtherNumber(t *testing.T) {
	actual := CharacterExpression(1234, "y", true)
	expected := "1234y"

	assert.Equal(t, expected, actual)
}

func TestSecondOtherNumber(t *testing.T) {
	actual := CharacterExpression(1234, "y", false)
	expected := "+1234y"

	assert.Equal(t, expected, actual)
}

func TestZero(t *testing.T) {
	actual := CharacterExpression(0, "x", true)
	expected := ""

	assert.Equal(t, expected, actual)
}
