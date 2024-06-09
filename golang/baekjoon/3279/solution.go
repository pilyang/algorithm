// greedy

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
	rates := make([]float64, n)
	for i := 0; i < n; i++ {
		rates[i] = float64(scanInt())
	}

	usd := 100.0
	marks := 0.0
	for i := 0; i < n-1; i++ {
		if rates[i] < rates[i+1] {
			usd += marks / rates[i]
			marks = 0
		} else {
			marks += usd * rates[i]
			usd = 0
		}
	}

	if marks > 0 {
		usd += marks / rates[n-1]
	}

	writer.WriteString(strconv.FormatFloat(usd, 'f', 2, 64))
	writer.WriteByte('\n')
}
