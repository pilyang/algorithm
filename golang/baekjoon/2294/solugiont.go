// https://www.acmicpc.net/problem/2294

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	input := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(input[0])
	k, _ := strconv.Atoi(input[1])

	coins := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		coins[i], _ = strconv.Atoi(scanner.Text())
	}

	result := findRequiredMinCoin(&coins, k)
	fmt.Fprintf(writer, "%d\n", result)
}

// f(n) = min(f(n-coin1), f(n-coin2), ... ) + 1
func findRequiredMinCoin(coins *[]int, target int) int {
	var minCoin int
	coins, minCoin = distinctAndMin(coins)

	if target < minCoin {
		return -1
	}

	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	dp[minCoin] = 1
	for i := minCoin + 1; i <= target; i++ {
		fmin := math.MaxInt
		for _, coin := range *coins {
			if i-coin >= 0 && dp[i-coin] < fmin && dp[i-coin] >= 0 {
				fmin = dp[i-coin]
			}
		}
		if fmin != math.MaxInt {
			dp[i] = fmin + 1
		}
	}

	return dp[target]
}

func distinctAndMin(sl *[]int) (*[]int, int) {
	mark := make(map[int]bool)
	distinctedSlice := make([]int, 0, len(*sl))
	m := (*sl)[0]

	for _, value := range *sl {
		if _, exist := mark[value]; !exist {
			mark[value] = true
			distinctedSlice = append(distinctedSlice, value)
			if value < m {
				m = value
			}
		}
	}
	return &distinctedSlice, m
}
