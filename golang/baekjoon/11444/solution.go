// https://www.acmicpc.net/problem/11444
// fast fibonacci, math

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

func fibonacci(n int64) int {
	if val, ok := dp[int64(n)]; ok {
		return val
	}

	if n%2 == 0 {
		dp[n] = fibonacci(n/2) * (fibonacci(n/2+1) + fibonacci(n/2-1)) % mod
	} else {
		dp[n] = (fibonacci(n/2)*fibonacci(n/2) + fibonacci(n/2+1)*fibonacci(n/2+1)) % mod
	}

	return dp[n]
}

const mod = 1_000_000_007

var dp map[int64]int

func main() {
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	dp = make(map[int64]int)
	dp[1] = 1
	dp[2] = 1

	writer.WriteString(strconv.Itoa(fibonacci(n)))
	writer.WriteByte('\n')
}
