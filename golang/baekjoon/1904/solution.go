// https://www.acmicpc.net/problem/1904

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

	result := countPossibleNum(num)

	fmt.Fprintf(writer, "%d\n", result)
}

const modNum = 15746

var dp = [1000001]int{0, 1, 2}

// f(n) = f(n-1) + f(n-2)
// if n == 4
// 001(1), 100(1), 111(1) -> which first digit is 1
// 11(00), 00(00) -> which first digit (and seccond digit also) 2
func countPossibleNum(num int) int {
	for i := 3; i <= num; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % modNum
	}

	return dp[num]
}
