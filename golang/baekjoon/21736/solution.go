// https://www.acmicpc.net/problem/21736
// bfs/dfs, graph

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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

var dxy = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func main() {
	defer writer.Flush()

	rows, columns := scanInt(), scanInt()
	graph := make([][]rune, rows)
	isVisited := make([][]bool, rows)

	// {r, c}
	var current [2]int

	for r := 0; r < rows; r++ {
		graph[r] = readRunes()
		isVisited[r] = make([]bool, columns)

		for c := 0; c < columns; c++ {
			if graph[r][c] == 'I' {
				current = [2]int{r, c}
			}
		}
	}

	stack := make([][2]int, 0, rows*columns)

	isVisited[current[0]][current[1]] = true
	stack = append(stack, current)

	var count int

	for len(stack) > 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]

		for _, d := range dxy {
			nr, nc := current[0]+d[0], current[1]+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= columns || isVisited[nr][nc] || graph[nr][nc] == 'X' {
				continue
			}

			if graph[nr][nc] == 'P' {
				count++
			}
			stack = append(stack, [2]int{nr, nc})
			isVisited[nr][nc] = true
		}
	}

	if count == 0 {
		writer.WriteString("TT")
	} else {
		writer.WriteString(strconv.Itoa(count))
	}
	writer.WriteByte('\n')
}
