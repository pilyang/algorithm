// https://www.acmicpc.net/problem/14503
// implementation

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

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

var (
	dr = []int{-1, 0, 1, 0}
	dc = []int{0, 1, 0, -1}
)

func main() {
	defer writer.Flush()

	// 0: not cleaned, 1 : wall, 2 : cleaned
	var board [][]int8
	rows, columns := readInt(), readInt()

	row, column, direction := readInt(), readInt(), readInt()

	board = make([][]int8, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]int8, columns)
		for j := 0; j < columns; j++ {
			board[i][j] = int8(readInt())
		}
	}

	var result int

	for {
		if board[row][column] == 0 {
			board[row][column] = 2
			result++
		}
		flag := false
		for i := 0; i < 4; i++ {
			direction = (direction + 3) % 4
			if row+dr[direction] < 0 || row+dr[direction] >= rows || column+dc[direction] < 0 || column+dc[direction] >= columns {
				continue
			}
			if board[row+dr[direction]][column+dc[direction]] == 0 {
				row += dr[direction]
				column += dc[direction]
				flag = true
				break
			}
		}
		if !flag {
			row -= dr[direction]
			column -= dc[direction]
			if row < 0 || row >= rows || column < 0 || column >= columns || board[row][column] == 1 {
				break
			}
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
