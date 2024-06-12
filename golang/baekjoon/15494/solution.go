//  math

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
	n /= 364

	x, k := 0, 1
	if n > 103 {
		k += (n - 101) / 3
	}
	x = n - 3*k

	writer.WriteString(strconv.Itoa(x) + "\n" + strconv.Itoa(k) + "\n")
}
