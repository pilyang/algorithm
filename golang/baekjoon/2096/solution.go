// https://www.acmicpc.net/problem/2096
// dp
// sliding window

package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer wr.Flush()

	num := readInt()
	road := make([][3]int, 0, num)
	var a, b, c int
	for i := 0; i < num; i++ {
		a, b, c = readInt(), readInt(), readInt()
		road = append(road, [3]int{a, b, c})
	}

	// seccond : 0 - max, 1 - min
	dp := [2][2][3]int{
		{
			road[0],
			road[0],
		},
	}
	prevIdx, currIdx := 0, 1
	for i := 1; i < num; i++ {
		// maxValues
		dp[currIdx][0][0] = maxInt(dp[prevIdx][0][0], dp[prevIdx][0][1]) + road[i][0]
		dp[currIdx][0][1] = maxInt(dp[prevIdx][0][0], maxInt(dp[prevIdx][0][1], dp[prevIdx][0][2])) + road[i][1]
		dp[currIdx][0][2] = maxInt(dp[prevIdx][0][1], dp[prevIdx][0][2]) + road[i][2]
		// minValues
		dp[currIdx][1][0] = minInt(dp[prevIdx][1][0], dp[prevIdx][1][1]) + road[i][0]
		dp[currIdx][1][1] = minInt(dp[prevIdx][1][0], minInt(dp[prevIdx][1][1], dp[prevIdx][1][2])) + road[i][1]
		dp[currIdx][1][2] = minInt(dp[prevIdx][1][1], dp[prevIdx][1][2]) + road[i][2]

		prevIdx, currIdx = currIdx, prevIdx
	}

	maxVal := maxInt(dp[prevIdx][0][0], maxInt(dp[prevIdx][0][1], dp[prevIdx][0][2]))
	minVal := minInt(dp[prevIdx][1][0], minInt(dp[prevIdx][1][1], dp[prevIdx][1][2]))
	wr.WriteString(strconv.Itoa(maxVal) + " " + strconv.Itoa(minVal) + "\n")
}
