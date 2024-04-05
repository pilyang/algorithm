// https://www.acmicpc.net/problem/17626
// brute force, dp

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	n := scanInt()

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			if dp[i] > dp[i-j*j]+1 {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.WriteByte('\n')
}
