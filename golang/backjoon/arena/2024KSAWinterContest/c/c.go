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

func backTracking(time, carrotCnt, s int) int {
	if time == k {
		return carrotCnt
	}

	if val, exist := dp[dpKey{time, carrotCnt, s}]; exist {
		return val
	}

	var result int
	// click
	if val := backTracking(time+1, carrotCnt+s, s); val > result {
		result = val
	}

	// use items
	for i := 0; i < n; i++ {
		if carrotCnt >= items[i][0] { // can use item
			if val := backTracking(time+1, carrotCnt-items[i][0], s+items[i][1]); val > result {
				result = val
			}
		}
	}

	dp[dpKey{time, carrotCnt, s}] = result
	return result
}

type dpKey struct {
	time, carrotCnt, s int
}

var (
	items [][2]int
	n, k  int
	dp    map[dpKey]int
)

func main() {
	defer writer.Flush()

	n, k = readInt(), readInt()
	items = make([][2]int, n)
	for i := range items {
		items[i] = [2]int{readInt(), readInt()}
	}

	dp = make(map[dpKey]int)

	result := backTracking(0, 0, 1)
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
