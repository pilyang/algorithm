// https://www.acmicpc.net/problem/2579
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
	res, _ := strconv.Atoi(scanner.Text())
	return res
}

// f(n) = max(dp[n,1], dp[n,2])
// dp[n,0] 해당 칸 안밟는 최대
// dp[n,1] 해당 칸이 연속으로 1번째인 경우의 최대
// dp[n,2] 해당칸이 연속으로 2번째인 경우 최대
func maxScoreSum(scores *[]int) int {
	l := len(*scores)
	dp := make([][3]int, l)

	dp[0][1] = (*scores)[0]
	for i := 1; i < l; i++ {
		dp[i][0] = maxInt(dp[i-1][1], dp[i-1][2])
		dp[i][1] = dp[i-1][0] + (*scores)[i]
		dp[i][2] = dp[i-1][1] + (*scores)[i]
	}
	return maxInt(dp[l-1][1], dp[l-1][2])
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	count := readInt()
	scores := make([]int, count)

	for i := 0; i < count; i++ {
		scores[i] = readInt()
	}

	res := maxScoreSum(&scores)
	writer.WriteString(strconv.Itoa(res) + "\n")
}
