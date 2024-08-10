package types

import "fmt"

type IntMatrix struct {
	value [][]int
}

func (m IntMatrix) Get(row, col int) int {
	return m.value[row][col]
}

func (m IntMatrix) Set(row, col, value int) {
	m.value[row][col] = value
}

func (m IntMatrix) Product(other IntMatrix) IntMatrix {
	if len(m.value[0]) != len(other.value) {
		panic("invalid matrix size")
	}

	res := make([][]int, len(m.value))
	for i := 0; i < len(m.value); i++ {
		res[i] = make([]int, len(other.value[0]))
		for j := 0; j < len(other.value[0]); j++ {
			res[i][j] = 0
			for k := 0; k < len(m.value[0]); k++ {
				res[i][j] += m.value[i][k] * other.value[k][j]
			}
		}
	}
	return IntMatrix{value: res}
}

func (m IntMatrix) ToTexString() string {
	str := "\\begin{bmatrix}\n"
	for i := range m.value {
		for j := range m.value[0] {
			str += fmt.Sprint(m.value[i][j])
			if j < len(m.value[0])-1 {
				str += " & "
			}
		}
		if i < len(m.value)-1 {
			str += " \\\\"
		}
		str += "\n"
	}
	str += "\\end{bmatrix}\n"
	return str
}

func CreateIntMatrix(args [][]int) IntMatrix {
	return IntMatrix{value: args}
}
