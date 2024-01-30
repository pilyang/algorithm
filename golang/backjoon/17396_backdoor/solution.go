// https://www.acmicpc.net/problem/17396

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := *bufio.NewReader(os.Stdin)
	writer := *bufio.NewWriter(os.Stdout)

	// read graph size data
	var (
		vertexCount int
		edgeCount   int
	)
	fmt.Fscanf(&reader, "%d %d\n", &vertexCount, &edgeCount)

	blackListInput, _ := reader.ReadString('\n')
	isBlacklist := make([]bool, vertexCount)

	// read blacklist data which is for the skip vertex
	for index, input := range strings.Split(blackListInput, " ") {
		if input == "1" {
			isBlacklist[index] = true
		}
	}
	isBlacklist[vertexCount-1] = false

	// init graph
	graph := make([][]*Edge, vertexCount)

	for i := 0; i < edgeCount; i++ {
		var (
			v1   int
			v2   int
			cost int
		)
		fmt.Fscanf(&reader, "%d %d %d\n", &v1, &v2, &cost)
		if !isBlacklist[v1] && !isBlacklist[v2] {
			graph[v1] = append(graph[v1], &Edge{target: v2, distance: int64(cost)})
			graph[v2] = append(graph[v2], &Edge{target: v1, distance: int64(cost)})
		}
	}

	result := findMinPathLengthFromStartToEnd(graph)
	fmt.Fprintf(&writer, "%d\n", result)

	writer.Flush()
}

func findMinPathLengthFromStartToEnd(graph [][]*Edge) int64 {
	num := len(graph)

	dist := make([]int64, num)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0
	pq := PriorityQueue{&Edge{target: 0, distance: 0}}
	heap.Init(&pq)

	var current *Edge
	for len(pq) != 0 {
		current = heap.Pop(&pq).(*Edge)
		if dist[current.target] < current.distance {
			continue
		}

		for _, edge := range graph[current.target] {
			newDist := dist[current.target] + edge.distance
			if newDist < dist[edge.target] {
				dist[edge.target] = newDist
				heap.Push(&pq, &Edge{target: edge.target, distance: newDist})
			}
		}
	}

	if dist[num-1] == math.MaxInt64 {
		return -1
	}
	return dist[num-1]
}

type Edge struct {
	target   int
	distance int64

	// // recommended if it need to use Fix, Remove functions which needs the index
	// index int
}

// implement Priority Queue with heap
// cehck the heap interface and sort interface wich is contained to heap interface
type PriorityQueue []*Edge

// check the memory usage difference between pointer reciever and value reciever
func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].distance < (*pq)[j].distance
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x any) {
	edge := x.(*Edge)
	*pq = append(*pq, edge)
}

func (pq *PriorityQueue) Pop() any {
	oldQp := *pq
	n := len(*pq)
	edge := oldQp[n-1]
	oldQp[n-1] = nil // prevent memory leak

	*pq = oldQp[0 : n-1]
	return edge
}
