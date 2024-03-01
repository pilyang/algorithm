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

var dxys = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func check(sx, sy int) bool {
	chocolate[sy][sx] = false

	// resotre chocolate
	defer func() {
		chocolate[sy][sx] = true
	}()

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	stack = stack[:0]
	fx, fy := findstart()

	stack = append(stack, [4]int{fx, fy, -1, -1})
	visited[fy][fx] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, dxy := range dxys {
			nx, ny := current[0]+dxy[0], current[1]+dxy[1]
			if ny < 0 || ny >= n || nx < 0 || nx >= n {
				continue
			}
			// if next is node
			if chocolate[ny][nx] {
				// and not visited yet
				if !visited[ny][nx] {
					stack = append(stack, [4]int{nx, ny, current[0], current[1]})
					visited[ny][nx] = true
					continue
				}

				if visited[ny][nx] && (nx != current[2] || ny != current[3]) {
					return false
				}
			}
		}
	}

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if chocolate[y][x] && !visited[y][x] {
				return false
			}
		}
	}

	return true
}

func findstart() (x, y int) {
	for x = 0; x < n; x++ {
		for y = 0; y < n; y++ {
			if chocolate[y][x] {
				return
			}
		}
	}
	return
}

func findResults() [][2]int {
	results := make([][2]int, 0, n*n)

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if chocolate[y][x] && check(x, y) {
				results = append(results, [2]int{y + 1, x + 1})
			}
		}
	}

	return results
}

var (
	chocolate [][]bool
	stack     [][4]int // x, y, px, py
	n         int
)

func main() {
	defer writer.Flush()

	n = readInt()
	chocolate = make([][]bool, n)
	for i := range chocolate {
		chocolate[i] = make([]bool, n)
		for j, c := range readRunes() {
			chocolate[i][j] = c == '#'
		}
	}
	stack = make([][4]int, 0, n*n)

	results := findResults()
	writer.WriteString(strconv.Itoa(len(results)))
	writer.WriteByte('\n')
	for _, result := range results {
		writer.WriteString(strconv.Itoa(result[0]))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(result[1]))
		writer.WriteByte('\n')
	}
}
