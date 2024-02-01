// https://www.acmicpc.net/problem/15422

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
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
		freeTickets[i] = [2]int{from, to}
	}

	resutl := findMinCost(&graph, &freeTickets, startCity, endCity)
	fmt.Fprintf(&writer, "%d\n", resutl)
}

type dataSender struct {
	ch chan<- int64
	wg *sync.WaitGroup
}

func findMinCost(graph *[][]*Edge, freeTickets *[][2]int, startCity, endCity int) int64 {
	if len(*freeTickets) == 0 {
		freeTickets = &[][2]int{{startCity, startCity}}
	}

	ch := make(chan int64, len(*freeTickets))
	ds := &dataSender{ch, &sync.WaitGroup{}}
	for _, ticket := range *freeTickets {
		ds.wg.Add(1)
		go dijkstraWithFreeTicket(graph, ticket, startCity, endCity, ds)
	}

	go func() {
		ds.wg.Wait()
		close(ch)
	}()

	minCost := int64(math.MaxInt64)
	for newCost := range ch {
		if newCost < minCost {
			minCost = newCost
		}
	}

	return int64(minCost)
}

func dijkstraWithFreeTicket(graph *[][]*Edge, ticket [2]int, startCity int, endCity int, ds *dataSender) {
	defer ds.wg.Done()

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

		if current.target == ticket[0] {
			if dist[current.target] < dist[ticket[1]] {
				dist[ticket[1]] = dist[current.target]
				heap.Push(pq, &Edge{ticket[1], dist[current.target]})
			}
		}
		for _, edge := range (*graph)[current.target] {
			newDist := dist[current.target] + edge.cost
			if newDist < dist[edge.target] {
				dist[edge.target] = newDist
				heap.Push(pq, &Edge{edge.target, newDist})
			}
		}
	}
	ds.ch <- dist[endCity]
	runtime.Gosched()
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
