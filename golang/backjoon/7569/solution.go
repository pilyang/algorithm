// https://www.acmicpc.net/problem/7569
// bfs

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
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

var (
	dx = [6]int{1, -1, 0, 0, 0, 0}
	dy = [6]int{0, 0, 1, -1, 0, 0}
	dz = [6]int{0, 0, 0, 0, 1, -1}
)

func bfsCount(graph *[][][]int8, initial *[][3]int) int {
	count := -1

	hs, rs, cs := len(*graph), len((*graph)[0]), len((*graph)[0][0])
	maxQSize := maxInt(hs*rs, maxInt(rs*cs, cs*hs))

	queues := make([][][3]int, 2)
	queues[0] = make([][3]int, 0, maxQSize)
	queues[1] = make([][3]int, 0, maxQSize)
	queues[0] = *initial
	currentIndex := 0
	nextIndex := 1

	for len(queues[currentIndex]) != 0 {
		queues[nextIndex] = queues[nextIndex][:0]

		for _, current := range queues[currentIndex] {
			for i := 0; i < 6; i++ {
				nx, ny, nz := current[0]+dx[i], current[1]+dy[i], current[2]+dz[i]
				if nx >= 0 && nx < cs && ny >= 0 && ny < rs && nz >= 0 && nz < hs && (*graph)[nz][ny][nx] == 0 {
					queues[nextIndex] = append(queues[nextIndex], [3]int{nx, ny, nz})
					// set as visit
					(*graph)[nz][ny][nx] = 1
				}
			}
		}
		count++
		currentIndex, nextIndex = nextIndex, currentIndex
	}

	for z := range *graph {
		for y := range (*graph)[z] {
			for x := range (*graph)[z][y] {
				if (*graph)[z][y][x] == 0 {
					return -1
				}
			}
		}
	}

	return count
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	column, row, height := readInt(), readInt(), readInt()
	graph := make([][][]int8, height)
	initial := [][3]int{}
	for h := range graph {
		graph[h] = make([][]int8, row)
		for r := range graph[h] {
			graph[h][r] = make([]int8, column)
			for c := range graph[h][r] {
				val := int8(readInt())
				graph[h][r][c] = val
				if val == int8(1) {
					initial = append(initial, [3]int{c, r, h})
				}
			}
		}
	}

	res := bfsCount(&graph, &initial)
	writer.WriteString(strconv.Itoa(res) + "\n")
}
