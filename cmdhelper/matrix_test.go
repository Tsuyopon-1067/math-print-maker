package cmdhelper

import (
	"printmaker/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrixProblemList(t *testing.T) {
	a := types.CreateIntMatrix([][]int{{1, 2}, {3, 4}})
	b := types.CreateIntMatrix([][]int{{3, 4}, {5, 6}})
	expected := types.ProblemAnswer{
		Problem: `\begin{bmatrix}
1 & 2 \\
3 & 4
\end{bmatrix}
\times
\begin{bmatrix}
3 & 4 \\
5 & 6
\end{bmatrix}
`,
		Answer: `\begin{bmatrix}
13 & 16 \\
29 & 36
\end{bmatrix}
`,
	}

	actual := generateMatrixProductProblemAnswer(a, b)
	assert.Equal(t, expected, actual)
}
