// https://www.acmicpc.net/problem/11266
// cut vertex, graph

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
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findCutVertex(graph [][]int) *[]bool {
	isCutVertex := make([]bool, len(graph))
	discover := make([]int, len(graph))

	var cnt int
	var dfs func(int, bool) int
	dfs = func(current int, isRoot bool) int {
		cnt++
		discover[current] = cnt
		num := discover[current]

		var childCnt int
		for _, next := range graph[current] {
			if discover[next] != 0 {
				num = minInt(num, discover[next])
				continue
			}

			childCnt++
			childNum := dfs(next, false)
			if !isRoot && discover[current] <= childNum {
				isCutVertex[current] = true
			}
			num = minInt(num, childNum)
		}

		if isRoot && childCnt > 1 {
			isCutVertex[current] = true
		}

		return num
	}

	for i := range graph {
		if discover[i] == 0 {
			dfs(i, true)
		}
	}

	return &isCutVertex
}

func main() {
	defer writer.Flush()

	v, e := readInt(), readInt()
	graph := make([][]int, v)
	for i := 0; i < e; i++ {
		a, b := readInt()-1, readInt()-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	results := findCutVertex(graph)

	var cnt int
	for _, isCutVertex := range *results {
		if isCutVertex {
			cnt++
		}
	}
	writer.WriteString(strconv.Itoa(cnt))
	writer.WriteByte('\n')

	for i, isCutVertex := range *results {
		if isCutVertex {
			writer.WriteString(strconv.Itoa(i + 1))
			writer.WriteByte(' ')
		}
	}
	writer.WriteByte('\n')
}
