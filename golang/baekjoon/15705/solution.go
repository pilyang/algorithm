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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

var (
	maxRow, maxColumn int
	board             [][]rune
	target            string
	directions        = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
)

func isTarget(row, column int, direction []int) bool {
	lastRow := row + direction[0]*len(target)
	lastColumn := column + direction[1]*len(target)

	if lastRow < 0 || lastRow >= maxRow || lastColumn < 0 || lastColumn >= maxColumn {
		return false
	}

	r, c := row, column
	for _, rune := range target {
		if rune != board[r][c] {
			return false
		}
		r += direction[0]
		c += direction[1]
	}
	return true
}

func hasTarget() bool {
	for row := 0; row < maxRow; row++ {
		for column := 0; column < maxColumn; column++ {
			for _, dir := range directions {
				if isTarget(row, column, dir) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	defer writer.Flush()

	target = scanString()

	maxRow, maxColumn = scanInt(), scanInt()
	board = make([][]rune, maxRow)
	for r := 0; r < maxRow; r++ {
		board[r] = []rune(scanString())
	}

	if hasTarget() {
		writer.WriteString("1\n")
	} else {
		writer.WriteString("0\n")
	}
}
