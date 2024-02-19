// https://www.acmicpc.net/problem/1865
// spfa (shortest path fastest algorithm) improved Bellman-Ford
// bellman-Ford algoritm -> not spfa
// reason
// spfa algorithm is only search for the connected edges from current node
// but bellman-ford algorithm can search all edges in graph
// spfa should search for the other starting nodes

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
	from, to, cost int
}

const maxTime = 10_000

func isContainNegativeLoop(edges *[]*Edge, vertexCount int) bool {
	dist := make([]int, vertexCount+1)
	for i := range dist {
		dist[i] = maxTime + 1
	}

	var isUpdated bool

	for i := 1; i < vertexCount; i++ {
		isUpdated = false
		for _, edge := range *edges {
			if dist[edge.to] > dist[edge.from]+edge.cost {
				dist[edge.to] = dist[edge.from] + edge.cost
				isUpdated = true
			}
		}

		if !isUpdated {
			break
		}
	}

	if isUpdated {
		for _, edge := range *edges {
			if dist[edge.to] > dist[edge.from]+edge.cost {
				return true
			}
		}
	}

	return false
}

func main() {
	defer writer.Flush()

	tc := readInt()

	var vertexCount, edgeCount, warpCount int
	var s, e, t int
	var edges []*Edge
	for i := 0; i < tc; i++ {
		vertexCount, edgeCount, warpCount = readInt(), readInt(), readInt()
		edges = make([]*Edge, 0, 2*edgeCount+warpCount)

		// read road
		for j := 0; j < edgeCount; j++ {
			s, e, t = readInt(), readInt(), readInt()
			edges = append(edges, &Edge{from: s, to: e, cost: t})
			edges = append(edges, &Edge{from: e, to: s, cost: t})
		}

		// read warp
		for j := 0; j < warpCount; j++ {
			s, e, t = readInt(), readInt(), readInt()
			edges = append(edges, &Edge{from: s, to: e, cost: -t})
		}

		if result := isContainNegativeLoop(&edges, vertexCount); result {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}
