// https://www.acmicpc.net/problem/31434
// dp
// very hard...
// same with /backjoon/arena/c

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

func updateDP(t, s, carrot int) {
	if dp[t] == nil {
		dp[t] = make(map[int]int)
	}
	if prev, exist := dp[t][s]; !exist || prev < carrot {
		dp[t][s] = carrot
	}
}

var dp []map[int]int

func main() {
	defer writer.Flush()

	n, k := readInt(), readInt()

	items := make(map[int]int, n)
	for i := 0; i < n; i++ {
		cost, increase := readInt(), readInt()
		if prev := items[cost]; prev < increase {
			items[cost] = increase
		}
	}

	// [t]map[s]carrot
	dp = make([]map[int]int, k+1)
	dp[0] = make(map[int]int)
	dp[0][1] = 0
	for t := 0; t < k; t++ {
		for s, carrot := range dp[t] {
			// dp[t+1][s] = dp[t][s] + s
			updateDP(t+1, s, carrot+s)
			for cost, increase := range items {
				// dp[t+1][s+item.increase] = dp[t][s] - item.cost
				if cost <= carrot {
					updateDP(t+1, s+increase, carrot-cost)
				}
			}
		}
	}

	var result int
	for _, carrot := range dp[k] {
		if result < carrot {
			result = carrot
		}
	}
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
