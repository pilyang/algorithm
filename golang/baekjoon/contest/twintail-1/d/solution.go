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

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	rows, columns := readInt(), readInt()

	arr := make([][]int, rows)
	for r := range arr {
		arr[r] = make([]int, columns)
	}

	var r, c int
	var direction [2]int
	var mirrorBaseline [2]int // {isColumnMirror, value}
	var maxNum int

	directionInput := readString()
	switch directionInput {
	case "U":
		r, c = 0, columns/2
		direction = [2]int{1, 0}
		mirrorBaseline = [2]int{1, columns / 2}
		maxNum = rows * (columns/2 + 1)
	case "D":
		r, c = rows-1, columns/2
		direction = [2]int{-1, 0}
		mirrorBaseline = [2]int{1, columns / 2}
		maxNum = rows * (columns/2 + 1)
	case "L":
		r, c = rows/2, 0
		direction = [2]int{0, 1}
		mirrorBaseline = [2]int{0, rows / 2}
		maxNum = (rows/2 + 1) * columns
	case "R":
		r, c = rows/2, columns-1
		direction = [2]int{0, -1}
		mirrorBaseline = [2]int{0, rows / 2}
		maxNum = (rows/2 + 1) * columns
	}

	// move is clock-wise direction
	for num := 1; num <= maxNum; num++ {
		arr[r][c] = num
		if mirrorBaseline[0] == 1 {
			arr[r][2*mirrorBaseline[1]-c] = num
		} else {
			arr[2*mirrorBaseline[1]-r][c] = num
		}

		// updirection
		if direction[0] == -1 {
			if r == 0 || arr[r-1][c] != 0 {
				direction = [2]int{0, 1}
			}

			// down direction
		} else if direction[0] == 1 {
			if r == rows-1 || arr[r+1][c] != 0 {
				direction = [2]int{0, -1}
			}

			// left direction
		} else if direction[1] == -1 {
			if c == 0 || arr[r][c-1] != 0 {
				direction = [2]int{-1, 0}
			}

			// right direction
		} else if direction[1] == 1 {
			if c == columns-1 || arr[r][c+1] != 0 {
				direction = [2]int{1, 0}
			}
		}

		r += direction[0]
		c += direction[1]
	}

	for _, row := range arr {
		for _, value := range row {
			writer.WriteString(strconv.Itoa(value))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
	}
}
