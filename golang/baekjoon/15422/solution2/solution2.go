package main

import (
	"bufio"
	"container/heap"
	_ "container/heap"
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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

type Edge struct {
	to   [2]int
	cost int64
}

type PriorityQueue []*Edge

func (p *PriorityQueue) Len() int               { return len(*p) }
func (p *PriorityQueue) Less(i int, j int) bool { return (*p)[i].cost < (*p)[j].cost }
func (p *PriorityQueue) Swap(i int, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }

func (p *PriorityQueue) Push(x any) { *p = append(*p, x.(*Edge)) }
func (p *PriorityQueue) Pop() any {
	i := len(*p) - 1
	e := (*p)[i]
	(*p)[i] = nil
	*p = (*p)[:i]
	return e
}

// graph[FreeTiketUsageNum][CityNum] = Edges slice
type Graph [2][][]*Edge

func findMinCost(graph *Graph, start, end int) int64 {
	// init dists
	cityCount := len((*graph)[0])
	var costFromStart [2][]int64
	for i := range costFromStart {
		costFromStart[i] = make([]int64, cityCount)
		for j := range costFromStart[i] {
			costFromStart[i][j] = math.MaxInt64
		}
	}
	costFromStart[0][start] = 0

	pq := make(PriorityQueue, 0, 2*cityCount)
	heap.Push(&pq, &Edge{to: [2]int{0, start}, cost: 0})
	var curent *Edge
	for pq.Len() > 0 {
		curent = heap.Pop(&pq).(*Edge)
		if costFromStart[curent.to[0]][curent.to[1]] < curent.cost {
			continue
		}

		for _, next := range (*graph)[curent.to[0]][curent.to[1]] {
			newCost := costFromStart[curent.to[0]][curent.to[1]] + next.cost
			if newCost < costFromStart[next.to[0]][next.to[1]] {
				costFromStart[next.to[0]][next.to[1]] = newCost
				heap.Push(&pq, &Edge{to: next.to, cost: newCost})
			}
		}
	}

	return minInt(costFromStart[0][end], costFromStart[1][end])
}

func minInt(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	city, roadNum, freeNum, start, end := readInt(), readInt(), readInt(), readInt(), readInt()

	var graph Graph
	graph[0] = make([][]*Edge, city)
	graph[1] = make([][]*Edge, city)

	var a, b, cost int
	for i := 0; i < roadNum; i++ {
		a, b, cost = readInt(), readInt(), readInt()
		graph[0][a] = append(graph[0][a], &Edge{to: [2]int{0, b}, cost: int64(cost)})
		graph[0][b] = append(graph[0][b], &Edge{to: [2]int{0, a}, cost: int64(cost)})
		graph[1][a] = append(graph[1][a], &Edge{to: [2]int{1, b}, cost: int64(cost)})
		graph[1][b] = append(graph[1][b], &Edge{to: [2]int{1, a}, cost: int64(cost)})
	}
	for i := 0; i < freeNum; i++ {
		a, b = readInt(), readInt()
		graph[0][a] = append(graph[0][a], &Edge{to: [2]int{1, b}, cost: 0})
	}

	result := findMinCost(&graph, start, end)
	writer.WriteString(strconv.FormatInt(result, 10) + "\n")
}
