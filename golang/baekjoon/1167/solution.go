// https://www.acmicpc.net/problem/1167
// https://blog.myungwoo.kr/112
// tree -> !!!!!not graph!!!!!
// reverse-dijkstrap? X
// backtracking, dfs

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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

type Edge struct {
	target, dist int
}

var stack []*Edge

func findMaxDistNode(graph *[][]*Edge, start int) (target int, dist int) {
	stack = stack[:0]
	stack = append(stack, &Edge{target: start, dist: 0})

	isVisit := make([]bool, len(*graph))
	isVisit[start] = true

	var current *Edge
	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if dist < current.dist {
			target = current.target
			dist = current.dist
		}

		for _, next := range (*graph)[current.target] {
			if !isVisit[next.target] {
				isVisit[next.target] = true
				stack = append(stack, &Edge{target: next.target, dist: current.dist + next.dist})
			}
		}
	}

	return target, dist
}

func main() {
	defer writer.Flush()

	v := readInt()
	graph := make([][]*Edge, v+1)

	var cv, nv, dist, cnt int
	for i := 0; i < v; i++ {
		cv = readInt()

		cnt = 0
		for nv = readInt(); nv != -1; nv = readInt() {
			dist = readInt()
			graph[cv] = append(graph[cv], &Edge{target: nv, dist: dist})
			cnt++
		}

		// add start point only for leaf node of tree
		if cnt == 1 {
			graph[0] = append(graph[0], &Edge{target: cv, dist: 0})
		}
	}

	stack := make([]*Edge, 0, v)
	stack = append(stack, &Edge{target: 0, dist: 0})

	midNode, _ := findMaxDistNode(&graph, 1)
	_, maxDist := findMaxDistNode(&graph, midNode)

	writer.WriteString(strconv.Itoa(maxDist))
	writer.WriteByte('\n')
}
