// https://www.acmicpc.net/problem/11053
// dp

package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func main() {
	defer wr.Flush()

	n := readInt()
	arr := make([]int, n)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	var maxCount int
	for _, count := range dp {
		if count > maxCount {
			maxCount = count
		}
	}
	wr.WriteString(strconv.Itoa(maxCount) + "\n")
}
