// https://www.acmicpc.net/problem/1149

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	costs := make([][3]int, num)
	for i := range costs {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		for j := 0; j < 3; j++ {
			costs[i][j], _ = strconv.Atoi(input[j])
		}
	}

	result := findMinCost(&costs)
	writer.WriteString(strconv.Itoa(result))
	writer.WriteString("\n")
}

// f(n) = min(f(n,0), f(n,1), f(n,2))
// f(n, k) = min( f(n-1, (k+1)%3), f(n-1, (k+1)%3) ) + costs[n][k]
func findMinCost(costs *[][3]int) int {
	// init dp slice
	l := len(*costs)
	dp := make([][3]int, l)
	dp[0][0] = (*costs)[0][0]
	dp[0][1] = (*costs)[0][1]
	dp[0][2] = (*costs)[0][2]

	for i := 1; i < l; i++ {
		for k := 0; k < 3; k++ {
			dp[i][k] = min(dp[i-1][(k+1)%3], dp[i-1][(k+2)%3]) + (*costs)[i][k]
		}
	}

	return min(dp[l-1][0], min(dp[l-1][1], dp[l-1][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
