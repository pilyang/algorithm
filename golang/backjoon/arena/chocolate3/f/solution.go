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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

var (
	graph [2][]int
	dp    [3][]int
)

func solution() {
	n := readInt()
	graph[0] = make([]int, n)
	graph[1] = make([]int, n)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)
	dp[2] = make([]int, n)
	for y := range dp {
		for x := range dp[y] {
			dp[y][x] = math.MinInt
		}
	}

	for y := range graph {
		for x := range graph[y] {
			graph[y][x] = readInt()
		}
	}

	sx, sy, ex, ey := readInt(), readInt(), readInt(), readInt()
	dp[sy][sx] = graph[sy][sx]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	tc := readInt()
	for tc > 0 {
		tc--
		solution()
	}
}
