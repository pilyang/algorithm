// https://www.acmicpc.net/problem/15422

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := *bufio.NewReader(os.Stdin)
	writer := *bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var numOfCity, numOfRoad, numOfFree, startCity, endCity int
	fmt.Fscanf(&reader, "%d %d %d %d %d\n", &numOfCity, &numOfRoad, &numOfFree, &startCity, &endCity)

	graph := make([][]*Edge, numOfCity)
	for i := range graph {
		graph[i] = make([]*Edge, 0)
	}
	for i := 0; i < numOfRoad; i++ {
		var a, b int
		var cost int64
		fmt.Fscanf(&reader, "%d %d %d\n", &a, &b, &cost)
		graph[a] = append(graph[a], &Edge{b, cost})
		graph[b] = append(graph[b], &Edge{a, cost})
	}

	freeTickets := make([][2]int, numOfFree)
	for i := 0; i < numOfFree; i++ {
		var from, to int
		fmt.Fscanf(&reader, "%d %d\n", &from, &to)
		freeTickets = append(freeTickets, [2]int{from, to})
	}

	resutl := findMinCost(&graph, &freeTickets, startCity, endCity)
	fmt.Fprintf(&writer, "%d\n", resutl)
}

func findMinCost(graph *[][]*Edge, freeTickets *[][2]int, startCity, endCity int) int64 {
	if len(*freeTickets) == 0 {
		return dijkstra(graph, startCity, endCity)
	}
	minCost := int64(math.MaxInt64)
	for _, ticket := range *freeTickets {
		old := (*graph)[ticket[0]]
		(*graph)[ticket[0]] = append(old, &Edge{ticket[1], 0})
		if newCost := dijkstra(graph, startCity, endCity); newCost < minCost {
			minCost = newCost
		}
		(*graph)[ticket[0]] = old
	}

	return int64(minCost)
}

func dijkstra(graph *[][]*Edge, startCity int, endCity int) int64 {
	dist := make([]int64, len(*graph))
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[startCity] = 0

	// init priority queue with starting position
	// priority queue has Edge from startCity
	pq := &PriorityQueue{&Edge{startCity, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Edge)
		if dist[current.target] < current.cost {
			continue
		}

		for _, edge := range (*graph)[current.target] {
			newDist := dist[current.target] + edge.cost
			if newDist < dist[edge.target] {
				dist[edge.target] = newDist
				heap.Push(pq, &Edge{edge.target, newDist})
			}
		}
	}
	return dist[endCity]
}

// Edge struct
type Edge struct {
	target int
	cost   int64
}

// priority queue for dijkstra
type PriorityQueue []*Edge

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].cost < (*pq)[j].cost
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
