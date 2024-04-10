// https://www.acmicpc.net/problem/2448
// recursive

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
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

var board [][]bool

var dc = []int{-2, -1, 0, 1, 2}

func drawStar(r, c, size int) {
	if size == 3 {
		board[r][c] = true
		board[r+1][c-1] = true
		board[r+1][c+1] = true
		for _, d := range dc {
			board[r+2][c+d] = true
		}
		return
	}

	drawStar(r, c, size/2)
	drawStar(r+size/2, c-size/2, size/2)
	drawStar(r+size/2, c+size/2, size/2)
}

func main() {
	defer writer.Flush()

	line := scanInt()
	board = make([][]bool, line)
	for i := 0; i < line; i++ {
		board[i] = make([]bool, 2*line)
	}

	drawStar(0, line-1, line)

	for _, row := range board {
		for _, c := range row {
			if c {
				writer.WriteByte('*')
			} else {
				writer.WriteByte(' ')
			}
		}
		writer.WriteByte('\n')
	}
}
