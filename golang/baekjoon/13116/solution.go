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

	t := scanInt()
	for t > 0 {
		t--
		a, b := scanInt(), scanInt()
		for a != b {
			if a > b {
				a /= 2
			} else {
				b /= 2
			}
		}
		writer.WriteString(strconv.Itoa(a * 10))
		writer.WriteByte('\n')
	}
}
