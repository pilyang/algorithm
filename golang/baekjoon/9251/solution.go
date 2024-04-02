// https://www.acmicpc.net/problem/9251
// dp, lcs (longest common subsequence)

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

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()

	first := readRunes()
	second := readRunes()

	dp := make([][]int, len(first)+1)
	dp[0] = make([]int, len(second)+1)
	for fi := 1; fi <= len(first); fi++ {
		dp[fi] = make([]int, len(second)+1)
		for si := 1; si <= len(second); si++ {
			if first[fi-1] == second[si-1] {
				dp[fi][si] = dp[fi-1][si-1] + 1
			} else {
				dp[fi][si] = maxInt(dp[fi-1][si], dp[fi][si-1])
			}
		}
	}

	res := dp[len(first)][len(second)]
	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}
