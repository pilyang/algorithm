// https://www.acmicpc.net/problem/12865
// knapsack problem
// https://www.youtube.com/watch?v=rhda6lR5kyQ

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
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

type Item struct {
	weight, score int
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[n][w] : n번째 아이템까지 고려한 경우 w무게에서의 최대 코스트
// dp[n][w]
// = max
// | dp[n-1][w] : when the n'th item is not in knapsack
// | dp[n-1][w-item(n).w] + item(n).cost : when th n'th item is in the knapsack
//
// this case memorize the all cases throw the items which means it needs 2-d storage
func findMaxSum(items *[]*Item, maxWeight int) int {
	dp := make([][]int, len(*items)+1)
	dp[0] = make([]int, maxWeight+1)
	var n int
	for i, item := range *items {
		n = i + 1
		dp[n] = make([]int, maxWeight+1)
		for w := 1; w <= maxWeight; w++ {
			if w-item.weight >= 0 {
				dp[n][w] = maxInt(dp[n-1][w], dp[n-1][w-item.weight]+item.score)
				continue
			}
			dp[n][w] = dp[n-1][w]
		}
	}

	return dp[len(*items)][maxWeight]
}

// this case re-use the storage of dp[n][w] to dp[w]
// it means it needs 1-d storage
// it reduce the storage and time both, because it doesn't need to allocate the 2-d slice
// dp[w] : w무게에서의 최대 코스트
// dp[w]
// = max
// | dp[w] : when the n'th item is not in knapsack
// | dp[w-item(n).w] + item(n).cost : when th n'th item is in the knapsack
func findMaxSum2(items *[]*Item, maxWeight int) int {
	dp := make([]int, maxWeight+1)

	for _, item := range *items {
		for w := maxWeight; w >= item.weight; w-- {
			dp[w] = maxInt(dp[w], dp[w-item.weight]+item.score)
		}
	}

	return dp[maxWeight]
}

func main() {
	defer writer.Flush()

	num, maxWeight := readInt(), readInt()
	items := make([]*Item, num)
	for i := 0; i < num; i++ {
		items[i] = &Item{weight: readInt(), score: readInt()}
	}

	result := findMaxSum2(&items, maxWeight)
	writer.WriteString(strconv.Itoa(result) + "\n")
}
