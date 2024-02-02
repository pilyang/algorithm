// https://www.acmicpc.net/problem/1463

package main

import (
	"bufio"
	"fmt"
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

func main() {
	defer writer.Flush()

	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := getMinOperationNum(num)
	fmt.Fprintf(writer, "%d\n", result)
}

func getMinOperationNum(num int) int {
	// f(n) = min(f(n-1), f(n/2), f(n/3)) + 1
	// f(1) = 0
	// f(2) = 1
	// f(3) = 1
	if num == 1 {
		return 0
	}
	if num <= 3 {
		return 1
	}
	dp := make([]int, num+1)
	for i := 1; i <= 3; i++ {
		dp[i] = 1
	}
	// try bottom up
	for i := 4; i <= num; i++ {
		f1, f2, f3 := num, num, num
		f1 = dp[i-1]
		if i%2 == 0 {
			f2 = dp[i/2]
		}
		if i%3 == 0 {
			f3 = dp[i/3]
		}
		dp[i] = minValue(f1, f2, f3) + 1
	}

	return dp[num]
}

func minValue(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
