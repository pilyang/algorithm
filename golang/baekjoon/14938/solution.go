// https://www.acmicpc.net/problem/14938
// floyd-warshall

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

const maxInt = math.MaxInt

func main() {
	defer writer.Flush()

	vertexCount, maxDistance, edgeCount := scanInt(), scanInt(), scanInt()

	items := make([]int, vertexCount+1)
	distances := make([][]int, vertexCount+1)
	for i := 1; i <= vertexCount; i++ {
		distances[i] = make([]int, vertexCount+1)
		for j := 1; j <= vertexCount; j++ {
			if i == j {
				distances[i][j] = 0
			} else {
				distances[i][j] = maxInt
			}
		}
	}

	for i := 1; i <= vertexCount; i++ {
		items[i] = scanInt()
	}

	for i := 0; i < edgeCount; i++ {
		from, to, distance := scanInt(), scanInt(), scanInt()
		distances[from][to] = distance
		distances[to][from] = distance
	}

	for k := 1; k <= vertexCount; k++ {
		for from := 1; from <= vertexCount; from++ {
			for to := 1; to <= vertexCount; to++ {
				if distances[from][k] == maxInt || distances[k][to] == maxInt {
					continue
				}
				if distances[from][to] > distances[from][k]+distances[k][to] {
					distances[from][to] = distances[from][k] + distances[k][to]
				}
			}
		}
	}

	maxItem := 0
	for i := 1; i <= vertexCount; i++ {
		itemCount := 0
		for j := 1; j <= vertexCount; j++ {
			if distances[i][j] <= maxDistance {
				itemCount += items[j]
			}
		}
		if itemCount > maxItem {
			maxItem = itemCount
		}
	}

	writer.WriteString(strconv.Itoa(maxItem))
	writer.WriteByte('\n')
}
