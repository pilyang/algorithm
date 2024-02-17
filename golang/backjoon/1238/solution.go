// https://www.acmicpc.net/problem/1238
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
	writer = bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
}

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

type Edge struct {
	target, cost int
}

type PriorityQueue []*Edge

func NewPriorityQueue(capa int) *PriorityQueue {
	pq := make(PriorityQueue, 0, capa)
	return &pq
}

func (p *PriorityQueue) Len() int               { return len(*p) }
func (p *PriorityQueue) Less(i int, j int) bool { return (*p)[i].cost < (*p)[j].cost }
func (p *PriorityQueue) Swap(i int, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }

func (p *PriorityQueue) Push(x any) { *p = append(*p, x.(*Edge)) }
func (p *PriorityQueue) Pop() any {
	e := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return e
}

func dijkstra(graph *[][]*Edge, start int) *[]int {
	// init
	pq = pq[:0]
	costs := make([]int, len(*graph))
	for i := range costs {
		costs[i] = math.MaxInt
	}

	heap.Push(&pq, &Edge{target: start, cost: 0})
	costs[start] = 0

	var current *Edge
	for len(pq) > 0 {
		current = heap.Pop(&pq).(*Edge)

		if costs[current.target] < current.cost {
			continue
		}

		for _, next := range (*graph)[current.target] {
			newCost := current.cost + next.cost
			if newCost < costs[next.target] {
				costs[next.target] = newCost
				heap.Push(&pq, &Edge{target: next.target, cost: newCost})
			}
		}
	}

	return &costs
}

var (
	v, e, x int
	pq      PriorityQueue
)

func main() {
	defer writer.Flush()

	v, e, x = readInt(), readInt(), readInt()
	forward := make([][]*Edge, v+1)
	reverse := make([][]*Edge, v+1)

	var from, to, cost int
	for i := 0; i < e; i++ {
		from, to, cost = readInt(), readInt(), readInt()
		forward[from] = append(forward[from], &Edge{target: to, cost: cost})
		reverse[to] = append(reverse[to], &Edge{target: from, cost: cost})
	}

	pq = *NewPriorityQueue(v)
	forwardCosts := dijkstra(&forward, x)
	reverwseCosts := dijkstra(&reverse, x)

	var maxCost int
	for i := 1; i <= v; i++ {
		if i == x {
			continue
		}
		maxCost = maxInt(maxCost, (*forwardCosts)[i]+(*reverwseCosts)[i])
	}

	writer.WriteString(strconv.Itoa(maxCost))
	writer.WriteByte('\n')
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
