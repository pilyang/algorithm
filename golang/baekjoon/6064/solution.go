// https://www.acmicpc.net/problem/6064
// lcm, gcd, math, 정수론, 중국인의 나머지 정리?, brute force

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()
	tc := scanInt()
	for tc > 0 {
		tc--
		m, n, x, y := scanInt(), scanInt(), scanInt(), scanInt()
		result := solve(m, n, x, y)
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func solve(m, n, x, y int) int {
	maxYear := lcm(m, n)
	year := x

	for year <= maxYear {
		if (year-y)%n == 0 {
			return year
		}
		year += m
	}

	return -1
}
