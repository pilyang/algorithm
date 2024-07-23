// dfs, backTracking

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

var board = [5][5]int{}

var drcs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func backTracking(row, column int, distance, score int) bool {
	if distance == 4 {
		return false
	}

	if board[row][column] == 1 {
		score++
	}
	if score == 2 {
		return true
	}

	isPossible := false

	original := board[row][column]
	board[row][column] = -1
	for _, drc := range drcs {
		nextRow, nextColumn := row+drc[0], column+drc[1]

		if nextRow < 0 || nextRow > 4 || nextColumn < 0 || nextColumn > 4 || board[nextRow][nextColumn] == -1 {
			continue
		}

		if backTracking(nextRow, nextColumn, distance+1, score) {
			isPossible = true
			break
		}

	}
	board[row][column] = original

	return isPossible
}

func main() {
	defer writer.Flush()

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			board[r][c] = scanInt()
		}
	}

	startRow, startColumn := scanInt(), scanInt()
	isPossible := backTracking(startRow, startColumn, 0, 0)

	if isPossible {
		writer.WriteString("1\n")
	} else {
		writer.WriteString("0\n")
	}
}
