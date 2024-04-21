// bfs

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

func solve(start, dest int) (result, count int) {
	if start == dest {
		return 0, 1
	}
	if start > dest {
		return start - dest, 1
	}

	isVisit := make([]bool, 100_001)
	queue := make([][2]int, 0, 100_001)

	queue = append(queue, [2]int{start, 0})

	var current [2]int
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		isVisit[current[0]] = true

		if count != 0 && current[1] > result {
			break
		}

		if current[0] == dest {
			result = current[1]
			count++
			continue
		}

		if count != 0 && current[1] == result {
			continue
		}

		if current[0]-1 >= 0 && !isVisit[current[0]-1] {
			queue = append(queue, [2]int{current[0] - 1, current[1] + 1})
		}
		if current[0]+1 <= 100_000 && !isVisit[current[0]+1] {
			queue = append(queue, [2]int{current[0] + 1, current[1] + 1})
		}
		if current[0]*2 <= 100_000 && !isVisit[current[0]*2] {
			queue = append(queue, [2]int{current[0] * 2, current[1] + 1})
		}
	}

	return
}

func main() {
	defer writer.Flush()

	result, count := solve(scanInt(), scanInt())

	writer.WriteString(strconv.Itoa(result) + "\n")
	writer.WriteString(strconv.Itoa(count) + "\n")
}
