// https://www.acmicpc.net/problem/2355
// sigma, bronze - 스트릭유지용...

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

func scanInt() int64 {
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return n
}

func main() {
	defer writer.Flush()

	a, b := scanInt(), scanInt()
	if a > b {
		a, b = b, a
	}
	sigma := (a + b) * (b - a + 1) / 2
	writer.WriteString(strconv.FormatInt(sigma, 10))
	writer.WriteByte('\n')
}
