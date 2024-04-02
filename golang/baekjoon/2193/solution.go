// https://www.acmicpc.net/problem/2193

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

	result := pinaryNumberCount(num)
	fmt.Fprintf(writer, "%d\n", result)
}

var dp = [91]int{0, 1, 1}

func pinaryNumberCount(digit int) int {
	// f(n) = f(n-1) + f(n-2))
	// if n = 5
	// 10000, 10010, 10100 -> 이친구들은 f(n-1)과 같음 -> 첫자리가 0 인 경우
	// 10001, 10101        -> 이친구들은 f(n-2)와 같음 -> 첫자리가 1인경우 -> 그 다음 2번째 자리는 0이여야함
	if digit <= 2 {
		return 1
	}
	for i := 3; i <= digit; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[digit]
}
