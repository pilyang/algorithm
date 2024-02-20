// https://www.acmicpc.net/problem/1932
// dp or brute force

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
	scanner.Split(bufio.ScanWords)
}

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

// dp[n, k] = max(dp[n-1, k-1], dp[n-1, k]) + val(n, k)
// concept
func findMax(triangle *[][]int) int {
	num := len(*triangle)

	var dp [2][]int
	dp[0] = make([]int, num+1)
	dp[1] = make([]int, num+1)

	currentIdx := 0
	prevIdx := 1

	for i, val := range *triangle {
		for k := 1; k <= i+1; k++ {
			dp[currentIdx][k] = maxInt(dp[prevIdx][k-1], dp[prevIdx][k]) + val[k]
		}
		currentIdx, prevIdx = prevIdx, currentIdx
	}

	var maxVal int
	for i := 1; i <= num; i++ {
		if maxVal < dp[prevIdx][i] {
			maxVal = dp[prevIdx][i]
		}
	}
	return maxVal
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	n := readInt()

	var dp [2][]int
	dp[0] = make([]int, n+1)
	dp[1] = make([]int, n+1)

	currentIdx := 0
	prevIdx := 1

	var triangleVal, maxVal int
	for i := 0; i < n; i++ {
		for k := 1; k <= i+1; k++ {
			triangleVal = readInt()
			dp[currentIdx][k] = maxInt(dp[prevIdx][k-1], dp[prevIdx][k]) + triangleVal
			if maxVal < dp[currentIdx][k] {
				maxVal = dp[currentIdx][k]
			}
		}
		currentIdx, prevIdx = prevIdx, currentIdx
	}

	writer.WriteString(strconv.Itoa(maxVal))
	writer.WriteByte('\n')
}
