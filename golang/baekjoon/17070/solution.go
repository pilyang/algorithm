// dp

package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

var (
	board             [][]int
	n                 int
	prevDp, currentDp [3][]int // 0: diagonal, 1: horizontal, 2: vertical
)

func checkDiagonal(r, c int) bool {
	return board[r-1][c] == 0 && board[r][c-1] == 0
}

func main() {
	defer writer.Flush()
	n = scanInt()

	board = make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
		for j := 0; j < n; j++ {
			board[i][j] = scanInt()
		}
	}

	for i := 0; i < 3; i++ {
		prevDp[i] = make([]int, n)
		currentDp[i] = make([]int, n)
	}

	for i := 1; i < n; i++ {
		if board[0][i] == 1 {
			break
		}
		prevDp[1][i] = 1
	}

	for row := 1; row < n; row++ {
		for column := 2; column < n; column++ {
			if board[row][column] == 1 {
				continue
			}

			// diagonal
			if checkDiagonal(row, column) {
				currentDp[0][column] = prevDp[0][column-1] + prevDp[1][column-1] + prevDp[2][column-1]
			} else {
				currentDp[0][column] = 0
			}

			// horizontal
			currentDp[1][column] = currentDp[0][column-1] + currentDp[1][column-1]

			// verticla
			currentDp[2][column] = prevDp[0][column] + prevDp[2][column]
		}

		prevDp, currentDp = currentDp, prevDp
		for i := 0; i < 3; i++ {
			for c := 0; c < n; c++ {
				currentDp[i][c] = 0
			}
		}
	}

	result := prevDp[0][n-1] + prevDp[1][n-1] + prevDp[2][n-1]
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
