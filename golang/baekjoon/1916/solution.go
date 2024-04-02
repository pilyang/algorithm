// https://www.acmicpc.net/problem/1916

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func readLine() []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func main() {
	defer writer.Flush()
	cityCount, _ := strconv.Atoi(readLine()[0])
	busCount, _ := strconv.Atoi(readLine()[0])

	edges := make([][]*Edge, cityCount+1)
	for i := 0; i < busCount; i++ {
		input := readLine()
		from, _ := strconv.Atoi(input[0])
		to, _ := strconv.Atoi(input[1])
		cost, _ := strconv.Atoi(input[2])
		edges[from] = append(edges[from], &Edge{to, cost})
	}
	input := readLine()
	start, _ := strconv.Atoi(input[0])
	end, _ := strconv.Atoi(input[1])

	result := dijkstra(&edges, start, end)
	fmt.Fprintf(writer, "%d\n", result)
}

func dijkstra(edge *[][]*Edge, start, end int) int {
	cost := make([]int, len(*edge))
	for i := range cost {
		cost[i] = math.MaxInt32
	}
	cost[start] = 0

	pq := PriorityQueue{&Edge{to: start, cost: 0}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Edge)
		if cost[current.to] < current.cost {
			continue
		}

		for _, next := range (*edge)[current.to] {
			newCost := cost[current.to] + next.cost
			if newCost < cost[next.to] {
				cost[next.to] = newCost
				heap.Push(&pq, &Edge{to: next.to, cost: newCost})
			}
		}
	}

	return cost[end]
}

type Edge struct {
	to   int
	cost int
}

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

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n-1]
	(*pq)[n-1] = nil
	*pq = (*pq)[0 : n-1]
	return x
}
