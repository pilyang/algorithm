// https://www.acmicpc.net/problem/14940
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

func readdInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

func bfsAndGetDsitances(dist *[][]int, graph *[][]bool, startX, startY int) {
	var cx, cy, nx, ny int

	rowCount, columnCount := len(*graph), len((*graph)[0])

	queue := make([][2]int, 0, rowCount*columnCount)
	queue = append(queue, [2]int{startX, startY})
	// isVisitable
	(*graph)[startY][startX] = false

	for len(queue) != 0 {
		cx, cy = queue[0][0], queue[0][1]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nx, ny = cx+dx[i], cy+dy[i]
			if nx >= 0 && nx < columnCount && ny >= 0 && ny < rowCount && (*graph)[ny][nx] {
				queue = append(queue, [2]int{nx, ny})
				(*graph)[ny][nx] = false
				(*dist)[ny][nx] = (*dist)[cy][cx] + 1
			}
		}
	}
}

func main() {
	defer writer.Flush()

	rowCount, columnCount := readdInt(), readdInt()
	graph := make([][]bool, rowCount)
	dist := make([][]int, rowCount)

	var startX, startY int

	for y := 0; y < rowCount; y++ {
		graph[y] = make([]bool, columnCount)
		dist[y] = make([]int, columnCount)
		for x := 0; x < columnCount; x++ {
			switch readdInt() {
			case 1:
				graph[y][x] = true
				dist[y][x] = -1
			case 2:
				dist[y][x] = 0
				startX, startY = x, y
			}
		}
	}

	bfsAndGetDsitances(&dist, &graph, startX, startY)

	for _, distLine := range dist {
		for _, val := range distLine {
			writer.WriteString(strconv.Itoa(val) + " ")
		}
		writer.WriteString("\n")
	}
}
