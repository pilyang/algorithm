// ad-hoc

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

	r, c, k := scanInt(), scanInt(), scanInt()

	if r+c-1 > k {
		writer.WriteString("NO\n")
		return
	}

	writer.WriteString("YES\n")
	for row := 1; row <= r; row++ {
		for column := 0; column < c; column++ {
			writer.WriteString(strconv.Itoa(row + column))
			writer.WriteRune(' ')
		}
		writer.WriteRune('\n')
	}
}
