// https://www.acmicpc.net/problem/11403
// floyd-warshall

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

func main() {
	defer writer.Flush()

	n := readInt()
	graph := make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			graph[i][j] = readInt() == 1
		}
	}

	for k := 0; k < n; k++ {
		for from := 0; from < n; from++ {
			for to := 0; to < n; to++ {
				if graph[from][k] && graph[k][to] {
					graph[from][to] = true
				}
			}
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if graph[r][c] {
				writer.WriteString("1 ")
			} else {
				writer.WriteString("0 ")
			}
		}
		writer.WriteByte('\n')
	}
}
