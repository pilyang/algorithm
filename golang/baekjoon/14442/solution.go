// https://www.acmicpc.net/problem/14442
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

var (
	graph [][]bool
	visit [][][]bool
)

type coordinate struct {
	row, col, layer, dist int
}

type Queue struct {
	c int
	v []coordinate
}

func newQueue(c int) *Queue {
	return &Queue{
		c: c,
		v: make([]coordinate, 0, c),
	}
}

func (q *Queue) len() int {
	return len(q.v)
}

func (q *Queue) push(e coordinate) {
	q.v = append(q.v, e)
}

func (q *Queue) pop() coordinate {
	e := q.v[0]
	q.v = q.v[1:]
	if len(q.v) == 0 {
		q.v = make([]coordinate, 0, q.c)
	}
	return e
}

func main() {
	defer writer.Flush()

	r, c, maxBroken := readInt(), readInt(), readInt()
	graph = make([][]bool, r)
	visit = make([][][]bool, r)
	for i := range graph {
		graph[i] = make([]bool, c)
		visit[i] = make([][]bool, c)
		for j, ch := range readRunes() {
			visit[i][j] = make([]bool, maxBroken+1)
			graph[i][j] = ch == '0'
		}
	}

	queue := newQueue(r * c)
	queue.push(coordinate{0, 0, 0, 1})

	visit[0][0][0] = true
	for queue.len() > 0 {
		current := queue.pop()

		if current.row == r-1 && current.col == c-1 {
			writer.WriteString(strconv.Itoa(current.dist))
			writer.WriteByte('\n')
			return
		}

		for _, drc := range drcs {
			nr, nc := current.row+drc[0], current.col+drc[1]

			if nc < 0 || nc >= c || nr < 0 || nr >= r {
				continue
			}

			// current layer
			if graph[nr][nc] && !visit[nr][nc][current.layer] {
				visit[nr][nc][current.layer] = true
				queue.push(coordinate{nr, nc, current.layer, current.dist + 1})
			}

			// next layer
			nl := current.layer + 1
			if !graph[nr][nc] && current.layer < maxBroken && !visit[nr][nc][nl] {
				visit[nr][nc][nl] = true
				queue.push(coordinate{nr, nc, nl, current.dist + 1})
			}
		}
	}

	writer.WriteString("-1\n")
}
