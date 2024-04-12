// https://www.acmicpc.net/problem/2638
// graph, dfs

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

type point struct {
	r, c int
}

const (
	empty = iota
	cheese
	air
)

var (
	graph       [][]int // 0: empty, 1: cheese, 2: air
	airStack    []point
	cheeseStack []point
)

var drcs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func isConnectedAir(r, c int) bool {
	for _, drc := range drcs {
		nr, nc := r+drc[0], c+drc[1]
		if nr < 0 || nr >= len(graph) || nc < 0 || nc >= len(graph[0]) {
			continue
		}
		if graph[nr][nc] == air {
			return true
		}
	}
	return false
}

func airDfs(r, c int) {
	airStack = airStack[:0]

	airStack = append(airStack, point{r: r, c: c})
	graph[r][c] = air

	var current point
	for len(airStack) > 0 {
		current, airStack = airStack[len(airStack)-1], airStack[:len(airStack)-1]

		for _, drc := range drcs {
			nr, nc := current.r+drc[0], current.c+drc[1]
			if nr < 0 || nr >= len(graph) || nc < 0 || nc >= len(graph[0]) {
				continue
			}
			if graph[nr][nc] == empty {
				airStack = append(airStack, point{r: nr, c: nc})
				graph[nr][nc] = air
			}
		}
	}
}

func isMelting(r, c int) bool {
	cnt := 0
	for _, drc := range drcs {
		nr, nc := r+drc[0], c+drc[1]
		if nr < 0 || nr >= len(graph) || nc < 0 || nc >= len(graph[0]) {
			continue
		}
		if graph[nr][nc] == air {
			cnt++
		}
	}
	return cnt >= 2
}

func solve() int {
	airDfs(0, 0)

	result := 0
	var current point
	for {

		// add melting targets
		for r := range graph {
			for c := range graph[r] {
				if graph[r][c] == cheese && isMelting(r, c) {
					cheeseStack = append(cheeseStack, point{r: r, c: c})
				}
			}
		}

		if len(cheeseStack) == 0 {
			break
		}

		// melt cheese
		for len(cheeseStack) > 0 {
			current, cheeseStack = cheeseStack[len(cheeseStack)-1], cheeseStack[:len(cheeseStack)-1]
			graph[current.r][current.c] = air

			for _, drc := range drcs {
				nr, nc := current.r+drc[0], current.c+drc[1]
				if nr < 0 || nr >= len(graph) || nc < 0 || nc >= len(graph[0]) {
					continue
				}

				if graph[nr][nc] == empty && isConnectedAir(nr, nc) {
					airDfs(nr, nc)
				}
			}
		}
		result++
	}

	return result
}

func main() {
	defer writer.Flush()

	r, c := scanInt(), scanInt()

	cnt := 0

	graph = make([][]int, r)
	for i := range graph {
		graph[i] = make([]int, c)
		for j := range graph[i] {
			graph[i][j] = scanInt()
			if graph[i][j] == 1 {
				cnt++
			}
		}
	}

	airStack = make([]point, 0, r*c)
	cheeseStack = make([]point, 0, cnt)

	result := solve()

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
