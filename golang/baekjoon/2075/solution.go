// https://www.acmicpc.net/problem/2075

package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

type PriorityQueue []int

func (p *PriorityQueue) Len() int               { return len(*p) }
func (p *PriorityQueue) Less(i int, j int) bool { return (*p)[i] < (*p)[j] }
func (p *PriorityQueue) Swap(i int, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }
func (p *PriorityQueue) Push(x any)             { *p = append(*p, x.(int)) }
func (p *PriorityQueue) Pop() any {
	i := p.Len() - 1
	e := (*p)[i]
	*p = (*p)[:i]
	return e
}

func main() {
	defer wr.Flush()
	n := readInt()
	pq := make(PriorityQueue, 0, n+1)
	for i := 0; i < n*n; i++ {
		heap.Push(&pq, readInt())
		if pq.Len() > n {
			heap.Pop(&pq)
		}
	}
	wr.WriteString(strconv.Itoa(heap.Pop(&pq).(int)) + "\n")
}
