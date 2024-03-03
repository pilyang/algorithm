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

func main() {
	defer writer.Flush()

	n := readInt()
	dp := make([]int, n+1)
	dp[2] = 1

	for i := 4; i <= n; i += 2 {
		dp[i] = dp[i-2] * (i - 1)
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.WriteByte('\n')
}
