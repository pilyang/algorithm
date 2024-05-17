// dp, brute-force

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
	board     [][]int
	pathCount [][]int
)

const mod = 1_000_000_007

func main() {
	defer writer.Flush()

	rows, columns, _ := scanInt(), scanInt(), scanInt()

	board = make([][]int, rows)
	pathCount = make([][]int, rows)

	for r := range board {
		board[r] = make([]int, columns)
		pathCount[r] = make([]int, columns)
		for c := range board[r] {
			board[r][c] = scanInt()
		}
	}

	pathCount[0][0] = 1

	for r := 0; r < rows-1; r++ {
		for c := 0; c < columns-1; c++ {
			for nr := r + 1; nr < rows; nr++ {
				for nc := c + 1; nc < columns; nc++ {
					if board[r][c] != board[nr][nc] {
						pathCount[nr][nc] += pathCount[r][c]
						pathCount[nr][nc] %= mod
					}
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(pathCount[rows-1][columns-1]))
	writer.WriteByte('\n')
}
