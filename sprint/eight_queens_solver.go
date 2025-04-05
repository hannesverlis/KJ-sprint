package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	var solutions []string
	solveEightQueens(&solutions, []int{}, 8)
	return strings.Join(solutions, "\n")
}

func solveEightQueens(solutions *[]string, currentQueens []int, n int) {
	if len(currentQueens) == n {
		solution := make([]string, n)
		for col, row := range currentQueens {
			solution[col] = strconv.Itoa(row + 1) // Convert row index to string (starting from 1)
		}
		*solutions = append(*solutions, strings.Join(solution, ""))
		return
	}

	row := len(currentQueens)
	for col := 0; col < n; col++ {
		if isSafe(currentQueens, row, col) {
			solveEightQueens(solutions, append(currentQueens, col), n)
		}
	}
}

func isSafe(currentQueens []int, row, col int) bool {
	for prevRow, prevCol := range currentQueens {
		if prevCol == col || prevRow-prevCol == row-col || prevRow+prevCol == row+col {
			return false
		}
	}
	return true
}