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

var rowColMask = [4][2]int{{0, 1}, {1, 0}, {1, 1}}

func main() {
	defer writer.Flush()

	n := scanInt()
	targetRowDiff, targetColDiff := scanInt(), scanInt()

	points := make(map[int]map[int]struct{})
	for i := 0; i < n; i++ {
		row, col := scanInt(), scanInt()
		if _, ok := points[row]; !ok {
			points[row] = make(map[int]struct{})
		}
		points[row][col] = struct{}{}
	}

	count := 0

	// row, col for left top
	for row, cols := range points {
		for col := range cols {
			flag := true
			for _, mask := range rowColMask {
				targetRow, targetCol := row+mask[0]*targetRowDiff, col+mask[1]*targetColDiff
				if cols, rowOk := points[targetRow]; rowOk {
					if _, colOk := cols[targetCol]; colOk {
						continue
					}
				}
				flag = false
				break
			}

			if flag {
				count++
			}
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
