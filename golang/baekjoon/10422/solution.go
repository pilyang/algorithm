// dp, catalan number

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

const mod = 1_000_000_007

var dp [2501]int

func findResult(n int) int {
	if dp[n] != 0 {
		return dp[n]
	}
	result := 0
	for i := 0; i < n; i++ {
		result += findResult(i) * findResult(n-i-1)
		result %= mod
	}
	dp[n] = result
	return dp[n]
}

func main() {
	defer writer.Flush()

	dp[0] = 1
	dp[1] = 1

	tc := scanInt()
	for ; tc > 0; tc-- {
		n := scanInt()
		if n%2 == 1 {
			writer.WriteString("0\n")
			continue
		}
		writer.WriteString(strconv.Itoa(findResult(n / 2)))
		writer.WriteByte('\n')
	}
}
