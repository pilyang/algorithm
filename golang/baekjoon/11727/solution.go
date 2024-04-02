// https://www.acmicpc.net/problem/11727
// dp

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
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

const mod int = 10_007

func main() {
	defer writer.Flush()

	n := readInt()
	if n == 1 {
		writer.WriteString("1\n")
		return
	}
	if n == 2 {
		writer.WriteString("3\n")
		return
	}

	fn1, fn2 := 3, 1
	result := 0
	for i := 3; i <= n; i++ {
		result = (fn1 + fn2*2) % mod
		fn2, fn1 = fn1, result
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
