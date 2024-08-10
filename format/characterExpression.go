package format

import "fmt"

func CharacterExpression(coefficient int, character string, isFirst bool) string {
	if coefficient == 0 {
		return ""
	}

	if coefficient == 1 {
		if isFirst {
			return character
		}
		return "+" + character
	}
	if coefficient == -1 {
		return "-" + character
	}

	num := fmt.Sprint(coefficient)
	if coefficient < 0 {
		return num + character
	}
	if isFirst {
		return num + character
	} else {
		return "+" + num + character
	}
}
