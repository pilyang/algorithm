// dfs, back-tracking

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

var (
	n, m  int
	stack []int
)

func dfs(depth int) {
	if depth == m {
		for _, v := range stack {
			writer.WriteString(strconv.Itoa(v))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	start := 1
	if depth > 0 {
		start = stack[depth-1]
	}
	for i := start; i <= n; i++ {
		stack[depth] = i
		dfs(depth + 1)
	}
}

func main() {
	defer writer.Flush()

	n, m = scanInt(), scanInt()

	stack = make([]int, m)

	dfs(0)
}
