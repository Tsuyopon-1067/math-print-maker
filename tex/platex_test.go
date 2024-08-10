package tex

import (
	"printmaker/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePlatesSource(t *testing.T) {
	problemAnsList := []types.ProblemAnswer{
		{Problem: "p1", Answer: "a1"},
		{Problem: "p2", Answer: "a2"},
		{Problem: "p3", Answer: "a3"},
	}
	column := 3
	actual := platesSource(problemAnsList, column)
	expected := `\documentclass[11pt,a4paper,dvipdfmx]{jsarticle}
\usepackage{multicol}
\usepackage{enumerate}
\begin{document}
\section*{問題}
\begin{multicols}{3}
\begin{enumerate}[(1)]
\item $p1$
\item $p2$
\item $p3$
\end{enumerate}
\end{multicols}

\newpage
\section*{解答}
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
