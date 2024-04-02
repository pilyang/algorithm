// https://www.acmicpc.net/problem/1697
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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

var moveForward = func(current int) int {
	return current + 1
}

var moveBackwrad = func(current int) int {
	return current - 1
}

var warp = func(current int) int {
	return current * 2
}

var nextFuncs = []func(int) int{moveForward, moveBackwrad, warp}

const mapSize int = 100_000

func main() {
	defer writer.Flush()

	start, to := readInt(), readInt()

	isVisit := make([]bool, mapSize+1)
	// [2]int{to, cnt}
	queue := make([][2]int, 0, mapSize+1)
	isVisit[start] = true

	queue = append(queue, [2]int{start, 0})
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current[0] == to {
			writer.WriteString(strconv.Itoa(current[1]))
			break
		}

		for _, nextFunc := range nextFuncs {
			next := nextFunc(current[0])
			if next >= 0 && next <= mapSize && !isVisit[next] {
				queue = append(queue, [2]int{next, current[1] + 1})
				isVisit[next] = true
			}
		}

	}
	writer.WriteByte('\n')
}
