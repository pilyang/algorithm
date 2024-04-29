// bfs, simul, implemenation

package main

import (
	"bufio"
	"math"
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

type pos struct {
	row, column int
}

type path struct {
	pos
	time int
}

type shark struct {
	size     int
	eatCount int
	pos
}

var (
	currentShark shark
	n            int
	board        [][]int
)

var drcs = [4]pos{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func upOrLeftPos(p1, p2 pos) pos {
	if p1.row < p2.row {
		return p1
	}
	if p1.row == p2.row && p1.column < p2.column {
		return p1
	}
	return p2
}

func eatNextFish() (int, bool) {
	isVisit := make([]bool, n*n)

	queue := make([]path, 0, n*n)

	queue = append(queue, path{currentShark.pos, 0})
	isVisit[currentShark.row*n+currentShark.column] = true

	var nextFishPos pos
	minTime := math.MaxInt

	var current path
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if board[current.row][current.column] != 0 && board[current.row][current.column] < currentShark.size {
			if minTime < current.time {
				continue
			}

			if minTime > current.time {
				minTime = current.time
				nextFishPos = current.pos
				continue
			}

			nextFishPos = upOrLeftPos(nextFishPos, current.pos)
			continue
		}

		for _, drc := range drcs {
			nr, nc := current.row+drc.row, current.column+drc.column

			if nr < 0 || nr >= n || nc < 0 || nc >= n {
				continue
			}
			if board[nr][nc] > currentShark.size {
				continue
			}
			if isVisit[nr*n+nc] {
				continue
			}

			isVisit[nr*n+nc] = true
			queue = append(queue, path{pos{nr, nc}, current.time + 1})
		}
	}

	if minTime == math.MaxInt {
		return 0, false
	}

	currentShark.pos = nextFishPos
	board[nextFishPos.row][nextFishPos.column] = 0

	currentShark.eatCount++
	if currentShark.eatCount == currentShark.size {
		currentShark.size++
		currentShark.eatCount = 0
	}

	return minTime, true
}

func main() {
	defer writer.Flush()

	n = scanInt()
	board = make([][]int, n)
	for r := 0; r < n; r++ {
		board[r] = make([]int, n)
		for c := 0; c < n; c++ {
			board[r][c] = scanInt()
			if board[r][c] == 9 {
				currentShark = shark{size: 2, eatCount: 0, pos: pos{r, c}}
				board[r][c] = 0
			}
		}
	}

	result := 0
	for {
		time, ok := eatNextFish()
		if !ok {
			break
		}
		result += time
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
