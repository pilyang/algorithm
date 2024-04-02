// https://www.acmicpc.net/problem/11725
// Tree Parent
// BFS, DFS -> BFS가 좀 더 빠름
// 이유 : 큐 최대 사이즈가 확실히 나와있고, 무한히 크지 않아서, for 내부에서 메모리 재할당을 최소화하고 구현 가능
// 큰 차이는 없긴 함..^^

package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func findParents(graph *[][]int) *[]int {
	// use parents as vivists slice
	parents := make([]int, len(*graph))
	parents[1] = -1

	q := make([]int, len(*graph))
	q = append(q, 1)

	var current int
	for len(q) > 0 {
		current = q[0]
		q = q[1:]

		for _, n := range (*graph)[current] {
			if parents[n] == 0 {
				parents[n] = current
				q = append(q, n)
			}
		}
	}

	return &parents
}

func main() {
	defer wr.Flush()

	// node num starts from 1
	n := readInt()
	graph := make([][]int, n+1)

	var a, b int
	for i := 0; i < n-1; i++ {
		a, b = readInt(), readInt()
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	result := findParents(&graph)
	for i := 2; i < len(*result); i++ {
		wr.WriteString(strconv.Itoa((*result)[i]) + "\n")
	}
}
