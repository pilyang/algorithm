// recursive

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

var result1, result2 int

func count(a, b int) {
	if a == 0 || b == 0 {
		return
	}
	result1++
	count(a-1, b)
	count(a, b-1)
}

func main() {
	defer writer.Flush()

	n := scanInt()

	count(n, n)
	result2 = n * n

	writer.WriteString(strconv.Itoa(result1))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(result2))
	writer.WriteByte('\n')
}
