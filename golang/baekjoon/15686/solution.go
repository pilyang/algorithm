// back-tracking, brute-force

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

var (
	house, chicken [][2]int
	result         int
	stack          []int // stack of chicken index
	n, m           int
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minDistance(h [2]int) int {
	min := 2 * n
	for _, c := range stack {
		dist := abs(h[0]-chicken[c][0]) + abs(h[1]-chicken[c][1])
		if dist < min {
			min = dist
		}
	}
	return min
}

func backTracking(depth int) {
	// calc distance and update result
	if depth == m {
		sum := 0
		for _, h := range house {
			sum += minDistance(h)
		}
		if sum < result {
			result = sum
		}

		return
	}

	nextIndexStart := 0
	if depth > 0 {
		nextIndexStart = stack[depth-1] + 1
	}

	for nexIndex := nextIndexStart; nexIndex < len(chicken); nexIndex++ {
		stack[depth] = nexIndex
		backTracking(depth + 1)
	}
}

func main() {
	defer writer.Flush()

	n, m = scanInt(), scanInt()

	house = make([][2]int, 0, 2*n)
	chicken = make([][2]int, 0, 13)
	stack = make([]int, m)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch scanInt() {
			case 1:
				house = append(house, [2]int{i, j})
			case 2:
				chicken = append(chicken, [2]int{i, j})
			}
		}
	}

	result = math.MaxInt
	backTracking(0)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
