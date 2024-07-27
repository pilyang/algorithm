// prefix sum

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

	row, column := scanInt(), scanInt()
	prefixSum := make([]int, column+1)
	prevSum := make([]int, column+1)
	for r := 0; r < row; r++ {
		for c := 1; c <= column; c++ {
			num := scanInt()
			prefixSum[c] = num + prefixSum[c-1] + prevSum[c] - prevSum[c-1]
		}
		prefixSum, prevSum = prevSum, prefixSum
	}
	prefixSum = prevSum

	a := scanInt()

	maxClap := 0
	for c := 0; c <= column-a; c++ {
		sum := prefixSum[c+a] - prefixSum[c]
		if maxClap < sum {
			maxClap = sum
		}
	}

	writer.WriteString(strconv.Itoa(maxClap))
	writer.WriteByte('\n')
}
