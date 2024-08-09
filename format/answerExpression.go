package format

import (
	"fmt"
	"slices"
)

func AnswerExpression(character string, answer ...int) string {
	slices.Sort(answer)
	unique := slices.Compact(answer)

	res := character + " = "
	for _, ans := range unique {
		res += fmt.Sprint(ans) + ", "
	}
	return res[:len(res)-2]
}
