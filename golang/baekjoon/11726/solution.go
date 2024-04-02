// https://www.acmicpc.net/problem/11726

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	result := countPossibleTileCount(num)

	fmt.Fprintf(writer, "%d\n", result)
}

const modNum = 10_007

// f(n) = f(n-1) + f(n-2)
// f(1) = 1
// f(2) = 2
func countPossibleTileCount(num int) int {
	if num < 3 {
		return num
	}
	dp := make([]int, num+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= num; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % modNum
	}

	return dp[num]
}
