// https://www.acmicpc.net/problem/16928
// bfs

package main

import (
	"bufio"
	"errors"
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
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

func minDiceCount(sl *map[int]int) (int, error) {
	count := 0

	isVisit := make([]bool, 100)

	queue := make([]int, 0, 6)
	queue = append(queue, 0)
	isVisit[0] = true

	for len(queue) != 0 {
		currents := make([]int, len(queue))
		copy(currents, queue)
		queue = make([]int, 0, 6)
		count++

		for _, current := range currents {
			for i := 1; i <= 6; i++ {
				next := current + i
				if val, exist := (*sl)[next]; exist {
					next = val
				}

				if next == 99 {
					return count, nil
				}

				if !isVisit[next] {
					queue = append(queue, next)
					isVisit[next] = true
				}
			}
		}
	}
	return 0, errors.New("impossible")
}

func main() {
	defer writer.Flush()

	l, s := readInt(), readInt()
	snakeLadder := make(map[int]int)

	count := s + l
	for i := 0; i < count; i++ {
		start, end := readInt()-1, readInt()-1
		snakeLadder[start] = end
	}

	res, err := minDiceCount(&snakeLadder)
	if err != nil {
		writer.WriteString("-1\n")
		return
	}
	writer.WriteString(strconv.Itoa(res) + "\n")
}
