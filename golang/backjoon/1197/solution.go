// https://www.acmicpc.net/problem/1197
// mst (minimum spanning tree)
// union-find
// kruscal

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
	writer = bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
}

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func findRoot(v int) int {
	if root[v] == 0 {
		return v
	}
	root[v] = findRoot(root[v])
	return root[v]
}

func union(v1, v2 int) bool {
	r1, r2 := findRoot(v1), findRoot(v2)
	if r1 == r2 {
		// already union
		return false
	}
	root[r2] = r1
	return true
}

var (
	graph [][3]int // [v1, v2, weight]
	root  []int    // 0 for root
)

func main() {
	defer writer.Flush()

	v, e := readInt(), readInt()
	root = make([]int, v+1)
	graph = make([][3]int, e)

	for i := 0; i < e; i++ {
		graph[i] = [3]int{readInt(), readInt(), readInt()}
	}

	slices.SortFunc(graph, func(a, b [3]int) int {
		return a[2] - b[2]
	})

	var mstVal, cnt int

	for _, edge := range graph {
		if union(edge[0], edge[1]) {
			mstVal += edge[2]
			cnt++
		}
		if cnt == v-1 {
			break
		}
	}

	writer.WriteString(strconv.Itoa(mstVal))
	writer.WriteByte('\n')
}
