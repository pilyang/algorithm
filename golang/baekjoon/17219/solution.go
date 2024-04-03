// https://www.acmicpc.net/problem/17219
// map

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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

func main() {
	defer writer.Flush()

	n, m := scanInt(), scanInt()

	passwords := make(map[string]string, n)
	for i := 0; i < n; i++ {
		s := scanString()
		passwords[s] = scanString()
	}

	for i := 0; i < m; i++ {
		s := scanString()
		writer.WriteString(passwords[s])
		writer.WriteByte('\n')
	}
}
