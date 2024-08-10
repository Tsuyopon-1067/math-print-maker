package tex

import (
	"fmt"
	"os"
	"os/exec"
	"printmaker/types"
)

func GeneratePdf(problemList []types.ProblemAnswer, column int) {
	content := platesSource(problemList, column)

	baseFileName := "tmp"
	tex := baseFileName + ".tex"
	dvi := baseFileName + ".dvi"
	aux := baseFileName + ".aux"
	log := baseFileName + ".log"

	writeFile(tex, content)
	execCommand("platex", "-interaction=batchmode", tex)
	execCommand("xdvipdfmx", dvi)
	execCommand("rm", aux, dvi, log, tex)
}

func writeFile(fileName string, content string) {
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("writing file error: %s", err)
		return
	}
}

func execCommand(command string, args ...string) {
	logFlag := false
	platexCmd := exec.Command(command, args...)

	output, err := platexCmd.CombinedOutput()
	if logFlag {
		fmt.Printf("%s command output:\n%s\n", command, output)
	}
	if err != nil {
		fmt.Printf("%s command error: %s\n", command, err)
		return
	}
}

func multiCols(contentList []string, column int) string {
	content := ""
	for i, c := range contentList {
		content += "\\item " + "$" + c + "$"
		if i != len(contentList)-1 {
			content += "\n"
		}
	}

	res := fmt.Sprintf("\\begin{multicols}{%d}\n", column) +
		"\\begin{enumerate}[(1)]" + "\n"
	res += content + "\n"
	res += "\\end{enumerate}\n"
	res += "\\end{multicols}\n"
	return res
}

func platesSource(problemAnswerList []types.ProblemAnswer, column int) string {
	problemList := make([]string, len(problemAnswerList))
	ansList := make([]string, len(problemAnswerList))
	for i, pa := range problemAnswerList {
		problemList[i] = pa.Problem
		ansList[i] = pa.Answer
	}

	problem := multiCols(problemList, column)
	ans := multiCols(ansList, column)

	return "\\documentclass[11pt,a4paper,dvipdfmx]{jsarticle}\n" +
		"\\usepackage{multicol}\n" +
		"\\usepackage{enumerate}\n" +
		"\\usepackage{amsmath}\n" +
		"\\begin{document}\n" +
		"\\section*{問題}\n" +
		problem + "\n" +
		"\\newpage\n" +
		"\\section*{解答}\n" +
		ans + "\n" +
		"\\end{document}\n"
}
