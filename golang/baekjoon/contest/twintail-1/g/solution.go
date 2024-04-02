// fail...
package main

import (
	"bufio"
	"fmt"
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

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

const modNum = 7

func modularPow(a, n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for n > 0 {
		if n%2 == 1 {
			result = (result * a) % modNum
		}
		// calculate a^2
		a = (a * a) % modNum
		n /= 2
	}
	result %= modNum
	return result
}

func main() {
	defer writer.Flush()

	a, b, c := readInt(), readInt(), readInt()
	k, l := readInt(), readInt()

	bc := modularPow(b, c%(modNum-1))
	// a^b^c
	abc := modularPow(a, bc)
	// b^c/a
	bca := ((bc % modNum) * modularPow(a, modNum-2)) % modNum

	abcWeekday := (k + abc) % modNum
	bcaWeekday := (k + bca) % modNum

	fmt.Printf("abc: %d, bca: %d\n", abc, bca)
	fmt.Printf("abcWeekday: %d, bcaWeekday: %d\n", abcWeekday, bcaWeekday)

	if abcWeekday == l && bcaWeekday == l {
		writer.WriteString("3\n")
	} else if bcaWeekday == l {
		writer.WriteString("2\n")
	} else if abcWeekday == l {
		writer.WriteString("1\n")
	} else {
		writer.WriteString("0\n")
	}
}
