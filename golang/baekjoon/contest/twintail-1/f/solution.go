package main

import (
	"bufio"
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
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

type Edge struct {
	to, cost int
}

var tree [][]Edge

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfsChildSum(current, parent int) int {
	// for leaf-node
	if len(tree[current]) == 1 && parent != -1 {
		return tree[current][0].cost
	}

	var sum int
	for _, child := range tree[current] {
		if child.to == parent {
			continue
		}
		sum += minInt(child.cost, dfsChildSum(child.to, current))
	}
	return sum
}

func main() {
	defer writer.Flush()

	vertex := readInt()
	tree = make([][]Edge, vertex+1)

	for i := 0; i < vertex-1; i++ {
		from, to, cost := readInt(), readInt(), readInt()
		tree[from] = append(tree[from], Edge{to, cost})
		tree[to] = append(tree[to], Edge{from, cost})
	}

	result := dfsChildSum(1, -1)
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
