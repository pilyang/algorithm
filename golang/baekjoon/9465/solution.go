// https://www.acmicpc.net/problem/9465

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// scanner -> index out of range
// func readInt(scanner *bufio.Scanner) int {
// 	scanner.Scan()
// 	num, _ := strconv.Atoi(scanner.Text())
// 	return num
// }

func readInt(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(input))
	return num
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	probNum := readInt(reader)
	results := make([]int, probNum)

	for i := 0; i < probNum; i++ {
		columnCount := readInt(reader)
		values := make([][2]int, columnCount)
		for j := 0; j < 2; j++ {
			input, _ := reader.ReadString('\n')
			inputSl := strings.Split(strings.TrimSpace(input), " ")
			for k, str := range inputSl {
				values[k][j], _ = strconv.Atoi(str)
			}
		}
		results[i] = getMaxScore(&values)
	}

	for _, result := range results {
		fmt.Fprintf(writer, "%d\n", result)
	}
}

// f(n) = max(dp[n][0], dp[n][1], dp[n][2])
// actually dp[n][0] could not be the result -> only for dp memorization
// the largest score (f(n)), the last column will be filled
// seccond index of dp
// 0 : n'th column with not selected
// 1 : n'th column with 0 row selected
// 2 : n'th column with 1 row selected
func getMaxScore(values *[][2]int) int {
	target := len(*values)
	dp := make([][3]int, target)
	dp[0][0] = 0
	dp[0][1] = (*values)[0][0]
	dp[0][2] = (*values)[0][1]

	for i := 1; i < target; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1], dp[i-1][2]))
		dp[i][1] = max(dp[i-1][1], max(dp[i-1][0], dp[i-1][2])+(*values)[i][0])
		dp[i][2] = max(dp[i-1][2], max(dp[i-1][0], dp[i-1][1])+(*values)[i][1])
	}

	return max(dp[target-1][1], dp[target-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
