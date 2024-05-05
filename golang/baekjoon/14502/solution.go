// dfs

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
	rows, columns int
	board         [][]int
	isVisit       [][]bool
	answer        int
	startPoints   [][2]int
)

func convertNumberToCoordinate(n int) (row, column int) {
	row = n / columns
	column = n % columns
	return
}

func initIsVisit() {
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			isVisit[row][column] = false
		}
	}
}

var (
	stack [][2]int
	drcs  = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
)

func dfs() {
	stack = stack[:0]
	initIsVisit()

	for _, startPoint := range startPoints {
		stack = append(stack, startPoint)
		isVisit[startPoint[0]][startPoint[1]] = true

		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, drc := range drcs {
				nextRow, nextColumn := current[0]+drc[0], current[1]+drc[1]
				if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns {
					continue
				}

				if isVisit[nextRow][nextColumn] {
					continue
				}

				if board[nextRow][nextColumn] == 0 {
					stack = append(stack, [2]int{nextRow, nextColumn})
					isVisit[nextRow][nextColumn] = true
				}
			}
		}
	}

	count := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if board[row][column] == 0 && !isVisit[row][column] {
				count++
			}
		}
	}
	if count > answer {
		answer = count
	}
}

func makeWall(wallCount, startNum int) {
	if wallCount == 3 {
		dfs()
		return
	}

	for num := startNum; num < rows*columns; num++ {
		row, column := convertNumberToCoordinate(num)
		if board[row][column] == 0 {
			board[row][column] = 1
			makeWall(wallCount+1, num+1)
			board[row][column] = 0
		}
	}
}

func main() {
	defer writer.Flush()

	rows, columns = scanInt(), scanInt()
	board = make([][]int, rows)
	for row := 0; row < rows; row++ {
		board[row] = make([]int, columns)
		for column := 0; column < columns; column++ {
			board[row][column] = scanInt()
			if board[row][column] == 2 {
				startPoints = append(startPoints, [2]int{row, column})
			}
		}
	}

	isVisit = make([][]bool, rows)
	for row := 0; row < rows; row++ {
		isVisit[row] = make([]bool, columns)
	}

	stack = make([][2]int, 0, rows*columns)

	makeWall(0, 0)

	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
