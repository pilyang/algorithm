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

	n, m := scanInt(), scanInt()
	friends := make([]int, n+1)

	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		friends[a]++
		friends[b]++
	}

	for i := 1; i <= n; i++ {
		writer.WriteString(strconv.Itoa(friends[i]))
		writer.WriteByte('\n')
	}
}
