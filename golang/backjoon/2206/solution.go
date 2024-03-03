// https://www.acmicpc.net/problem/2206
// bfs

package main

import (
	"bufio"
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

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

// r, c, broken, cost
var drcs = [][2]int{
	// for current layer
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

const maxBroken = 1

var (
	graph [][]bool
	dist  [][][2]int
)

type coordinate struct {
	row, col, layer int
}

func main() {
	defer writer.Flush()

	r, c := readInt(), readInt()
	graph = make([][]bool, r)
	dist = make([][][2]int, r)
	for i := range graph {
		graph[i] = make([]bool, c)
		dist[i] = make([][2]int, c)
		for j, ch := range readRunes() {
			graph[i][j] = ch == '0'
		}
	}

	queue := make([]coordinate, 0, 2*r*c)
	queue = append(queue, coordinate{0, 0, 0})
	dist[0][0][0] = 1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, drc := range drcs {
			nr, nc := current.row+drc[0], current.col+drc[1]

			if nc < 0 || nc >= c || nr < 0 || nr >= r {
				continue
			}

			// current layer
			if graph[nr][nc] && dist[nr][nc][current.layer] == 0 {
				dist[nr][nc][current.layer] = dist[current.row][current.col][current.layer] + 1
				queue = append(queue, coordinate{nr, nc, current.layer})
			}

			// next layer
			nl := current.layer + 1
			if !graph[nr][nc] && current.layer < maxBroken && dist[nr][nc][nl] == 0 {
				dist[nr][nc][nl] = dist[current.row][current.col][current.layer] + 1
				queue = append(queue, coordinate{nr, nc, nl})
			}
		}
	}

	minDist := math.MaxInt
	for i := 0; i < 2; i++ {
		if dist[r-1][c-1][i] != 0 && minDist > dist[r-1][c-1][i] {
			minDist = dist[r-1][c-1][i]
		}
	}
	if minDist == math.MaxInt {
		minDist = -1
	}
	writer.WriteString(strconv.Itoa(minDist))
	writer.WriteByte('\n')
}
