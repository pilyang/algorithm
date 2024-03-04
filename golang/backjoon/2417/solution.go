// https://www.acmicpc.net/problem/2417
// math, binary-search

package main

import (
	"bufio"
	"math"
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

func readInt() int64 {
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return n
}

func binarySearch(target int64) int64 {
	if target == 0 {
		return 0
	}

	var l, h int64 = 0, int64(math.Sqrt(math.MaxInt64)) + 1
	var m int64
	for h-l > 1 {
		m = l + (h-l)/2
		if m*m >= target {
			h = m
		} else {
			l = m
		}
	}
	return h
}

func main() {
	defer writer.Flush()

	n := readInt()
	var ans int64 = binarySearch(n)

	writer.WriteString(strconv.FormatInt(ans, 10))
	writer.WriteByte('\n')
}
