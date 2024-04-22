// dijkstra

package main

import (
	"bufio"
	"container/heap"
	"math"
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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

type Node struct {
	to   int
	cost int
}

type PriorityQueue []Node

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less(i int, j int) bool {
	return (*p)[i].cost < (*p)[j].cost
}

func (p *PriorityQueue) Swap(i int, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(Node))
}

func (p *PriorityQueue) Pop() any {
	e := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return e
}

func main() {
	defer writer.Flush()

	nodeCount := scanInt()
	edgeCount := scanInt()

	graph := make([][][2]int, nodeCount+1)
	for i := 1; i <= nodeCount; i++ {
		graph[i] = make([][2]int, 0, nodeCount+1)
	}
	for i := 0; i < edgeCount; i++ {
		from := scanInt()
		to := scanInt()
		cost := scanInt()
		graph[from] = append(graph[from], [2]int{to, cost})
	}

	start, dest := scanInt(), scanInt()

	dist := make([]int, nodeCount+1)
	prev := make([]int, nodeCount+1)

	for i := 1; i <= nodeCount; i++ {
		dist[i] = math.MaxInt
	}

	dist[start] = 0

	pq := make(PriorityQueue, 0, nodeCount)
	heap.Push(&pq, Node{to: start, cost: 0})

	for len(pq) > 0 {
		current := heap.Pop(&pq).(Node)

		if current.cost > dist[current.to] {
			continue
		}

		for _, next := range graph[current.to] {
			newCost := current.cost + next[1]
			if newCost < dist[next[0]] {
				heap.Push(&pq, Node{to: next[0], cost: newCost})

				dist[next[0]] = newCost
				prev[next[0]] = current.to
			}
		}
	}

	path := make([]int, 0, nodeCount)
	for i := dest; i != 0; i = prev[i] {
		path = append(path, i)
	}

	writer.WriteString(strconv.Itoa(dist[dest]))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(len(path)))
	writer.WriteByte('\n')
	for i := len(path) - 1; i >= 0; i-- {
		writer.WriteString(strconv.Itoa(path[i]))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
