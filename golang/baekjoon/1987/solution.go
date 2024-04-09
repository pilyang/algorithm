// https://www.acmicpc.net/problem/1987
// dfs, back tracking

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

func scanRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

var drc = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func solve(graph [][]rune, rows, columns int) int {
	result := 0
	pathRunes := make([]bool, 26)
	// var bits uint32

	var dfs func(r, c, count int)
	dfs = func(r, c, count int) {
		for _, d := range drc {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= columns {
				continue
			}
			if pathRunes[graph[nr][nc]-'A'] {
				continue
			}
			// if bits&(1<<(graph[nr][nc]-'A')) != 0 {
			// 	continue
			// }

			pathRunes[graph[nr][nc]-'A'] = true
			// bits |= 1 << (graph[nr][nc] - 'A')
			dfs(nr, nc, count+1)
			pathRunes[graph[nr][nc]-'A'] = false
			// bits &^= 1 << (graph[nr][nc] - 'A')
		}

		if count > result {
			result = count
		}
	}

	pathRunes[graph[0][0]-'A'] = true
	// bits |= 1 << (graph[0][0] - 'A')
	dfs(0, 0, 1)

	return result
}

func main() {
	defer writer.Flush()

	rows, columns := scanInt(), scanInt()
	graph := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		graph[i] = scanRunes()
	}

	result := solve(graph, rows, columns)
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
