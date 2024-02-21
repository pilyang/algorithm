// https://www.acmicpc.net/problem/1074
// divide and conquer, recursive

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

func IntPow(a, n int) int {
	if n == 0 {
		return 1
	}
	ans := 1
	for i := 0; i < n; i++ {
		ans *= a
	}
	return ans
}

// pos divide
// 0 1
// 2 3
func z(n, r, c int) int {
	if n == 1 {
		return findArea(n, r, c)
	}

	quarterSize := IntPow(4, n-1)
	area := findArea(n, r, c)
	mid := (IntPow(2, n) - 1) / 2
	if r > mid {
		r = r - mid - 1
	}
	if c > mid {
		c = c - mid - 1
	}

	return area*quarterSize + z(n-1, r, c)
}

func findArea(n, r, c int) int {
	switch mid := (IntPow(2, n) - 1) / 2; {
	case r <= mid && c <= mid:
		return 0
	case r <= mid && c > mid:
		return 1
	case r > mid && c <= mid:
		return 2
	default:
		return 3
	}
}

func main() {
	defer writer.Flush()

	n, r, c := readInt(), readInt(), readInt()

	res := z(n, r, c)

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}
