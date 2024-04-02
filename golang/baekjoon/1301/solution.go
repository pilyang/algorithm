// https://www.acmicpc.net/problem/1301
// backtracking, dp

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
}

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

var (
	n, total          int
	beeds, containCnt [6]int
	dp                [6][6][11][11][11][11][11]int // dp(current, prev, 종류별 사용한 구슬들 숫자...) = 가능한 수..!!
)

func initDp() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			for k := 0; k < 11; k++ {
				for l := 0; l < 11; l++ {
					for m := 0; m < 11; m++ {
						for n := 0; n < 11; n++ {
							for o := 0; o < 11; o++ {
								dp[i][j][k][l][m][n][o] = -1
							}
						}
					}
				}
			}
		}
	}
}

func backTracking(depth, current, prev int) int {
	if depth == total {
		return 1
	}

	if dpVal := dp[current][prev][containCnt[1]][containCnt[2]][containCnt[3]][containCnt[4]][containCnt[5]]; dpVal != -1 {
		return dpVal
	}

	var res int
	for num, count := range beeds {
		if num == 0 {
			continue
		}
		// 이 조건에서 dp를 적용해야할라나..?
		if num != current && num != prev && containCnt[num] < count {
			// add
			containCnt[num]++

			sub := backTracking(depth+1, num, current)
			dp[num][current][containCnt[1]][containCnt[2]][containCnt[3]][containCnt[4]][containCnt[5]] = sub
			res += sub

			// remove
			containCnt[num]--
		}
	}
	return res
}

func main() {
	defer writer.Flush()

	n = readInt()

	for i := 1; i < n+1; i++ {
		input := readInt()
		beeds[i] = input
		total += input
	}
	initDp()

	result := backTracking(0, 0, 0)
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
