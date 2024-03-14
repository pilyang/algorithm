// https://www.acmicpc.net/problem/11404
// floyd-warshall

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
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const maxCost = math.MaxInt

func main() {
	defer writer.Flush()

	v, e := readInt(), readInt()
	graph := make([][]int, v)
	for i := 0; i < v; i++ {
		graph[i] = make([]int, v)
		for j := 0; j < v; j++ {
			if i == j {
				continue
			}
			graph[i][j] = maxCost
		}
	}

	for i := 0; i < e; i++ {
		from, to, cost := readInt()-1, readInt()-1, readInt()
		graph[from][to] = minInt(graph[from][to], cost)
	}

	// floyd-warshall
	for k := 0; k < v; k++ {
		for from := 0; from < v; from++ {
			for to := 0; to < v; to++ {
				if graph[from][k] != maxCost && graph[k][to] != maxCost {
					graph[from][to] = minInt(graph[from][to], graph[from][k]+graph[k][to])
				}
			}
		}
	}

	for r := 0; r < v; r++ {
		for c := 0; c < v; c++ {
			if graph[r][c] == maxCost {
				graph[r][c] = 0
			}
			writer.WriteString(strconv.Itoa(graph[r][c]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
	}
}
