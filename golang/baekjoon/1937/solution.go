// https://www.acmicpc.net/problem/1927
// min heap

package main

import (
	"bufio"
	"container/heap"
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
}

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

type PriorityQueue []int

func (p *PriorityQueue) Len() int               { return len(*p) }
func (p *PriorityQueue) Less(i int, j int) bool { return (*p)[i] < (*p)[j] }
func (p *PriorityQueue) Swap(i int, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }

func (p *PriorityQueue) Push(x any) { *p = append(*p, x.(int)) }
func (p *PriorityQueue) Pop() any {
	e := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return e
}

func main() {
	defer writer.Flush()

	n := readInt()
	pq := make(PriorityQueue, 0, n)

	for i := 0; i < n; i++ {
		if input := readInt(); input == 0 {
			if pq.Len() == 0 {
				writer.WriteString("0\n")
			} else {
				writer.WriteString(strconv.Itoa(heap.Pop(&pq).(int)))
				writer.WriteByte('\n')
			}
		} else {
			heap.Push(&pq, input)
		}
	}
}
