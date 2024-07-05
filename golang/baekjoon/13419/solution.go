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

	tc := scanInt()
	for tc > 0 {
		tc--

		s := scanString()
		res1, res2 := make([]rune, 0, len(s)), make([]rune, 0, len(s))
		for i, r := range s {
			if i%2 == 0 {
				res1 = append(res1, r)
			} else {
				res2 = append(res2, r)
			}
		}
		if len(s)%2 == 1 {
			res1, res2 = append(res1, res2...), append(res2, res1...)
		}

		writer.WriteString(string(res1) + "\n" + string(res2) + "\n")
	}
}
