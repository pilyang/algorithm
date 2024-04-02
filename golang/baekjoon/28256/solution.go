// https://www.acmicpc.net/problem/28256
// bfs dfs

package main

import (
	"bufio"
	"os"
	"slices"
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

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

var dxys = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func dfs(x, y int) int {
	stack = stack[:0]

	stack = append(stack, [2]int{x, y})
	graph[x][y] = false

	size := 1
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, dxy := range dxys {
			nx, ny := current[0]+dxy[0], current[1]+dxy[1]
			if nx < 0 || nx >= 3 || ny < 0 || ny >= 3 || !graph[nx][ny] {
				continue
			}

			stack = append(stack, [2]int{nx, ny})
			graph[nx][ny] = false
			size++
		}
	}

	return size
}

func findConnections() []int {
	result := make([]int, 0, 4)

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if graph[x][y] {
				result = append(result, dfs(x, y))
			}
		}
	}

	slices.Sort(result)

	return result
}

func assert(expected, actual []int) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}

	return true
}

var (
	graph [][]bool
	stack [][2]int
)

func main() {
	defer writer.Flush()

	graph = make([][]bool, 3)
	for i := range graph {
		graph[i] = make([]bool, 3)
	}
	stack = make([][2]int, 0, 8)

	tc := readInt()
	for tc > 0 {
		tc--
		for i := range graph {
			for j, r := range readRunes() {
				graph[i][j] = r == 'O'
			}
		}
		n := readInt()
		expected := make([]int, n)
		for i := range expected {
			expected[i] = readInt()
		}
		actual := findConnections()

		if assert(expected, actual) {
			writer.WriteString("1\n")
		} else {
			writer.WriteString("0\n")
		}
	}
}
