// https://www.acmicpc.net/problem/1504
// graph, floyd warshall

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	n, e := scanInt(), scanInt()
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i == j {
				continue
			}
			graph[i][j] = math.MaxInt
		}
	}

	for i := 0; i < e; i++ {
		from, to, cost := scanInt(), scanInt(), scanInt()
		graph[from][to] = int(cost)
		graph[to][from] = int(cost)
	}

	wp1, wp2 := scanInt(), scanInt()

	for k := 1; k <= n; k++ {
		for from := 1; from <= n; from++ {
			if from == k {
				continue
			}
			for to := 1; to <= n; to++ {
				if to == k || to == from {
					continue
				}
				if graph[from][k] == math.MaxInt || graph[k][to] == math.MaxInt {
					continue
				}
				graph[from][to] = min(graph[from][to], graph[from][k]+graph[k][to])
			}
		}
	}

	result := math.MaxInt
	for _, path := range [][2]int{{wp1, wp2}, {wp2, wp1}} {
		if graph[1][path[0]] == math.MaxInt || graph[path[0]][path[1]] == math.MaxInt || graph[path[1]][n] == math.MaxInt {
			continue
		}
		result = min(result, graph[1][path[0]]+graph[path[0]][path[1]]+graph[path[1]][n])
	}
	if result == math.MaxInt {
		result = -1
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
