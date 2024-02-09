// https://www.acmicpc.net/problem/1260
// dfs and bfs

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
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

func writeIntSlice(sl *[]int) {
	for _, val := range *sl {
		writer.WriteString(strconv.Itoa(val) + " ")
	}
	writer.WriteString("\n")
}

func dfs(graph *[][]int, start int) *[]int {
	l := len(*graph) - 1
	visitOrder := make([]int, 0, l)
	isVisited := make([]bool, l+1)

	// init start point
	stack := make([]int, 0, l)
	stack = append(stack, start)

	for len(stack) != 0 {
		// pop
		ci := len(stack) - 1
		current := stack[ci]
		stack = stack[:ci]
		if isVisited[current] {
			continue
		}
		visitOrder = append(visitOrder, current)
		isVisited[current] = true

		// push stack from big to small
		for i := len((*graph)[current]) - 1; i >= 0; i-- {
			n := (*graph)[current][i]
			if !isVisited[n] {
				stack = append(stack, n)
			}
		}
	}

	return &visitOrder
}

func bfs(graph *[][]int, start int) *[]int {
	l := len(*graph) - 1
	visitOrder := make([]int, 0, l)
	isVisited := make([]bool, l+1)

	// init start point
	queue := make([]int, 0, l)
	queue = append(queue, start)

	for len(queue) != 0 {
		// pop
		current := queue[0]
		queue = queue[1:]
		if isVisited[current] {
			continue
		}
		visitOrder = append(visitOrder, current)
		isVisited[current] = true

		// push
		for _, n := range (*graph)[current] {
			if !isVisited[n] {
				queue = append(queue, n)
			}
		}
	}

	return &visitOrder
}

func main() {
	defer writer.Flush()

	vc, ec, st := readInt(), readInt(), readInt()
	graph := make([][]int, vc+1)
	for i := 0; i < ec; i++ {
		a, b := readInt(), readInt()
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for i := range graph {
		slices.Sort(graph[i])
	}

	dfsOrder := dfs(&graph, st)
	bfsOrder := bfs(&graph, st)
	writeIntSlice(dfsOrder)
	writeIntSlice(bfsOrder)
}
