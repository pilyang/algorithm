// https://www.acmicpc.net/problem/7662
// heap : priority queue

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
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	scanner.Scan()
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

const maxCommand = 1_000_000

type Item struct {
	val, lIndex, sIndex int
}

func NewItem(val int) *Item {
	return &Item{val: val, lIndex: -1, sIndex: -1}
}

type LargePQ []*Item

func (lq *LargePQ) Len() int           { return len(*lq) }
func (lq *LargePQ) Less(i, j int) bool { return (*lq)[i].val > (*lq)[j].val }
func (lq *LargePQ) Swap(i, j int) {
	itemI, itemJ := (*lq)[i], (*lq)[j]
	itemI.lIndex, itemJ.lIndex = itemJ.lIndex, itemI.lIndex
	(*lq)[i], (*lq)[j] = itemJ, itemI
}

func (lq *LargePQ) Push(x any) {
	i := lq.Len()
	item := x.(*Item)
	item.lIndex = i
	*lq = append(*lq, item)
}

func (lq *LargePQ) Pop() any {
	i := len(*lq) - 1
	e := (*lq)[i]
	(*lq)[i] = nil
	*lq = (*lq)[:i]
	return e
}

type SmallPQ []*Item

func (sq *SmallPQ) Len() int           { return len(*sq) }
func (sq *SmallPQ) Less(i, j int) bool { return (*sq)[i].val < (*sq)[j].val }
func (sq *SmallPQ) Swap(i, j int) {
	itemI, itemJ := (*sq)[i], (*sq)[j]
	itemI.sIndex, itemJ.sIndex = itemJ.sIndex, itemI.sIndex
	(*sq)[i], (*sq)[j] = itemJ, itemI
}

func (sq *SmallPQ) Push(x any) {
	i := sq.Len()
	item := x.(*Item)
	item.sIndex = i
	*sq = append(*sq, item)
}

func (sq *SmallPQ) Pop() any {
	i := len(*sq) - 1
	e := (*sq)[i]
	(*sq)[i] = nil
	*sq = (*sq)[:i]
	return e
}

type PrioriyQueue struct {
	*LargePQ
	*SmallPQ
}

func NewPQ(capa int) *PrioriyQueue {
	lq := make(LargePQ, 0, capa)
	sq := make(SmallPQ, 0, capa)
	return &PrioriyQueue{LargePQ: &lq, SmallPQ: &sq}
}

func (pq *PrioriyQueue) Push(item *Item) {
	heap.Push(pq.LargePQ, item)
	heap.Push(pq.SmallPQ, item)
}

func (pq *PrioriyQueue) PopLarge() *Item {
	large := heap.Pop(pq.LargePQ).(*Item)
	heap.Remove(pq.SmallPQ, large.sIndex)
	large.lIndex, large.sIndex = -1, -1
	return large
}

func (pq *PrioriyQueue) PopSmall() *Item {
	small := heap.Pop(pq.SmallPQ).(*Item)
	heap.Remove(pq.LargePQ, small.lIndex)
	small.lIndex, small.sIndex = -1, -1
	return small
}

func (pq *PrioriyQueue) Len() int {
	return len(*pq.LargePQ)
}

func (pq *PrioriyQueue) Init() {
	*pq.LargePQ = (*pq.LargePQ)[:0]
	*pq.SmallPQ = (*pq.SmallPQ)[:0]
}

func main() {
	defer writer.Flush()

	count := readInt()

	pq := NewPQ(maxCommand)
	for i := 0; i < count; i++ {
		pq.Init()

		commandCount := readInt()
		for j := 0; j < commandCount; j++ {
			switch readString() {
			case "I":
				pq.Push(NewItem(readInt()))
			case "D":
				commandNum := readInt()
				if pq.Len() == 0 {
					continue
				}
				switch commandNum {
				case 1:
					pq.PopLarge()
				case -1:
					pq.PopSmall()
				}
			}
		}

		if pq.Len() == 0 {
			writer.WriteString("EMPTY")
		} else if pq.Len() == 1 {
			val := strconv.Itoa(pq.PopLarge().val)
			writer.WriteString(val + " " + val)
		} else {
			largeVal := strconv.Itoa(pq.PopLarge().val)
			smallVal := strconv.Itoa(pq.PopSmall().val)
			writer.WriteString(largeVal + " " + smallVal)
		}
		writer.WriteString("\n")
	}
}
