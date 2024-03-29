// https://www.acmicpc.net/problem/9461
// dp

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
	writer = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

var dp [101]int

func solve(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return 1
	}
	if n == 4 || n == 5 {
		return 2
	}

	if dp[n] != 0 {
		return dp[n]
	}

	dp[n] = solve(n-1) + solve(n-5)
	return dp[n]
}

func main() {
	defer writer.Flush()

	tc := readInt()
	for i := 0; i < tc; i++ {
		n := readInt()
		writer.WriteString(strconv.Itoa(solve(n)))
		writer.WriteByte('\n')
	}
}
