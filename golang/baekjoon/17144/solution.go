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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

var (
	room, temp       [][]int
	rows, columns, t int
	upCleaner        int
	downCleaner      int
)

var drcs = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func spread() {
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if room[r][c] <= 0 {
				continue
			}

			count := 0
			spreading := room[r][c] / 5

			for _, drc := range drcs {
				nr, nc := r+drc[0], c+drc[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= columns || room[nr][nc] == -1 {
					continue
				}

				temp[nr][nc] += spreading
				count++
			}
			temp[r][c] += room[r][c] - spreading*count
			room[r][c] = 0
		}
	}

	room, temp = temp, room
}

func moveValues(cleanerRow, endRow int, direction int) {
	for r := cleanerRow + direction; r != endRow; r += direction {
		room[r][0] = room[r+direction][0]
	}

	for c := 1; c < columns; c++ {
		room[endRow][c-1] = room[endRow][c]
	}

	for r := endRow; r != cleanerRow; r -= direction {
		room[r][columns-1] = room[r-direction][columns-1]
	}

	for c := columns - 1; c > 1; c-- {
		room[cleanerRow][c] = room[cleanerRow][c-1]
	}

	room[cleanerRow][1] = 0
}

func clean() {
	// Clean up part
	moveValues(upCleaner, 0, -1)

	// Clean down part
	moveValues(downCleaner, rows-1, 1)
}

func main() {
	defer writer.Flush()

	rows, columns, t = scanInt(), scanInt(), scanInt()
	room = make([][]int, rows)
	temp = make([][]int, rows)
	for i := 0; i < rows; i++ {
		room[i] = make([]int, columns)
		temp[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			num := scanInt()
			room[i][j] = num
			if num == -1 {
				temp[i][j] = -1
				if upCleaner == 0 {
					upCleaner = i
				} else {
					downCleaner = i
				}
			}
		}
	}

	for t > 0 {
		spread()
		clean()
		t--
	}

	var sum int
	for _, row := range room {
		for _, dust := range row {
			if dust > 0 {
				sum += dust
			}
		}
	}

	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
