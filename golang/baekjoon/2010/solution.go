// math

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

func main() {
	defer writer.Flush()

	n := scanInt()

	result := 0

	for i := 0; i < n; i++ {
		result += scanInt()
	}
	result -= n - 1

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
