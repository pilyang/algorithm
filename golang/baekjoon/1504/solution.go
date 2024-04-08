// https://www.acmicpc.net/problem/1504
// graph, floyd warshall -> dijkstra

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

type Edge struct {
	to, cost int
}

type PriorityQueue []Edge

func (p *PriorityQueue) Len() int               { return len(*p) }
func (p *PriorityQueue) Less(i int, j int) bool { return (*p)[i].cost < (*p)[j].cost }
func (p *PriorityQueue) Swap(i int, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }
func (p *PriorityQueue) Push(x any)             { *p = append(*p, x.(Edge)) }
func (p *PriorityQueue) Pop() any {
	e := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return e
}

var (
	pq       PriorityQueue
	graph    [][]Edge
	distance []int
)

func dijkstra(from, to int) int {
	for i := range distance {
		distance[i] = math.MaxInt
	}
	pq = pq[:0]

	distance[from] = 0
	heap.Push(&pq, Edge{to: from, cost: 0})
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(Edge)
		if distance[current.to] < current.cost {
			continue
		}

		for _, next := range graph[current.to] {
			newDist := distance[current.to] + next.cost
			if newDist < distance[next.to] {
				distance[next.to] = newDist
				heap.Push(&pq, Edge{to: next.to, cost: newDist})
			}
		}
	}

	return distance[to]
}

func main() {
	defer writer.Flush()

	n, e := scanInt(), scanInt()
	graph = make([][]Edge, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]Edge, 0)
	}
	distance = make([]int, n+1)
	pq = make(PriorityQueue, 0, n)

	for i := 0; i < e; i++ {
		from, to, cost := scanInt(), scanInt(), scanInt()
		graph[from] = append(graph[from], Edge{to: to, cost: cost})
		graph[to] = append(graph[to], Edge{to: from, cost: cost})
	}

	wp1, wp2 := scanInt(), scanInt()

	result := math.MaxInt
	for _, testCase := range [][][]int{
		{{1, wp1}, {wp1, wp2}, {wp2, n}},
		{{1, wp2}, {wp2, wp1}, {wp1, n}},
	} {
		flag := false
		dist := 0
		for _, path := range testCase {
			temp := dijkstra(path[0], path[1])
			if temp == math.MaxInt {
				flag = true
				break
			}
			dist += temp
		}
		if flag {
			continue
		}
		result = minInt(result, dist)
	}
	if result == math.MaxInt {
		result = -1
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
