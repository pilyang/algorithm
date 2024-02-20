// https://www.acmicpc.net/problem/1967
// dfs or bfs
// tree

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
	scanner.Split(bufio.ScanWords)
}

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

type Edge struct {
	to, dist int
}

var (
	tree  [][]*Edge
	stack []int
)

func findFarthestNode(start int) (to, dist int) {
	stack = stack[:0]
	isVisit := make([]bool, len(tree))
	dists := make([]int, len(tree))

	stack = append(stack, start)
	isVisit[start] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if dist < dists[current] {
			to = current
			dist = dists[current]
		}

		for _, next := range tree[current] {
			if !isVisit[next.to] {
				isVisit[next.to] = true
				stack = append(stack, next.to)
				dists[next.to] = dists[current] + next.dist
			}
		}
	}
	return to, dist
}

func main() {
	defer writer.Flush()

	num := readInt()

	tree = make([][]*Edge, num+1)
	stack = make([]int, 0, num)

	var a, b, d int
	for i := 0; i < num-1; i++ {
		a, b, d = readInt(), readInt(), readInt()
		tree[a] = append(tree[a], &Edge{to: b, dist: d})
		tree[b] = append(tree[b], &Edge{to: a, dist: d})
	}

	start, _ := findFarthestNode(1)
	_, diameter := findFarthestNode(start)

	writer.WriteString(strconv.Itoa(diameter))
	writer.WriteByte('\n')
}
