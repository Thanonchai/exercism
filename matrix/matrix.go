package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m Matrix) Rows() (r [][]int) {
	r = make([][]int, len(m))
	for i, a := range m {
		r[i] = make([]int, len(a))
		copy(r[i], a)
	}
	return r
}

func (m Matrix) Cols() [][]int {
	rowNum := len(m)
	colNum := len(m[0])
	result := make([][]int, colNum)
	for i := 0; i < colNum; i++ {
		result[i] = make([]int, rowNum)
		for j := 0; j < rowNum; j++ {
			result[i][j] = m[j][i]
		}
	}
	return result
}

func (m Matrix) Set(row, col, val int) bool {
	switch {
	case row < 0:
		return false
	case col < 0:
		return false
	case row >= len(m):
		return false
	case col >= len(m[0]):
		return false
	}
	m[row][col] = val
	return true
}

func New(input string) (Matrix, error) {
	lines := strings.Split(input, "\n")
	col := len(strings.Split(lines[0], " "))
	m := make([][]int, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		nums := strings.Split(line, " ")
		if len(nums) != col {
			return [][]int{}, errors.New("Inequal dimensions.")
		}
		m[i] = make([]int, len(nums))
		for j, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				return [][]int{}, errors.New("Not a number.")
			}
			m[i][j] = val
		}
	}
	return Matrix(m), nil
}
