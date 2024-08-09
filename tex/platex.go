package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	content := `
	\documentclass[a4j]{jarticle}
	\usepackage{amsmath}
	\begin{document}
		\begin{align}
		  a &= b + c \\
		  d &= e + f
		\end{align}
	\end{document}
	`
	GeneratePdf(content)
}

func GeneratePdf(content string) {
	baseFileName := "tmp"
	tex := baseFileName + ".tex"
	aux := baseFileName + ".aux"
	dvi := baseFileName + ".dvi"
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
