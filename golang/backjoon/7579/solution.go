// https://www.acmicpc.net/problem/7579
// dp, knapsack

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
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[n][c] : n번째 앱까지 고려한 경우 c코스트 이하로 얻을 수 있는 최대 메모리
// dp[n][c]
// = max
// | dp[n-1][c] : when the n'th app is not installed
// | dp[n-1][c-cost(n)] + memory(n) : when the n'th app is installed
//
// reduce the 2-d storage to 1-d storage
// dp[c] : c코스트 이하로 얻을 수 있는 최대 메모리
// dp[c] = max(dp[c], dp[c-cost(n)]+memory(n))
func findMinCost(memories, costs []int, maxCost, reqMem int) int {
	dp := make([]int, maxCost+1)
	for item := 0; item < len(memories); item++ {
		for cost := maxCost; cost >= costs[item]; cost-- {
			dp[cost] = maxInt(dp[cost], dp[cost-costs[item]]+memories[item])
		}
	}

	for cost, mem := range dp {
		if mem >= reqMem {
			return cost
		}
	}
	return -1
}

func main() {
	defer writer.Flush()

	num, reqMem := readInt(), readInt()
	mems := make([]int, num)
	costs := make([]int, num)
	for i := 0; i < num; i++ {
		mems[i] = readInt()
	}
	var maxCost int
	for i := 0; i < num; i++ {
		costs[i] = readInt()
		maxCost += costs[i]
	}

	result := findMinCost(mems, costs, maxCost, reqMem)

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
