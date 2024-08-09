package tex

import (
	"printmaker/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePlatesSource(t *testing.T) {
	problemAnsList := []types.ProblemAnswer{
		types.ProblemAnswer{"p1", "a1"},
		types.ProblemAnswer{"p2", "a2"},
		types.ProblemAnswer{"p3", "a3"},
	}
	column := 3
	actual := platesSource(problemAnsList, column)
	expected := `\documentclass[11pt,a4paper,dvipdfmx]{jsarticle}
\usepackage{multicol}
\usepackage{enumerate}
\begin{document}
\begin{multicols}{3}
\begin{enumerate}[(1)]
\item $p1$
\item $p2$
\item $p3$
\end{enumerate}
\end{multicols}

\newpage
\begin{multicols}{3}
\begin{enumerate}[(1)]
\item $a1$
\item $a2$
\item $a3$
\end{enumerate}
\end{multicols}

\end{document}
`

	assert.Equal(t, expected, actual)
}
