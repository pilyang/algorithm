// https://www.acmicpc.net/problem/11659
// Prefix Sum

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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func main() {
	defer writer.Flush()

	num, tc := readInt(), readInt()

	prefixSum := make([]int, num+1)

	for i := 1; i <= num; i++ {
		prefixSum[i] = prefixSum[i-1] + readInt()
	}

	var s, e int
	for i := 0; i < tc; i++ {
		s, e = readInt(), readInt()
		writer.WriteString(strconv.Itoa(prefixSum[e] - prefixSum[s-1]))
		writer.WriteByte('\n')
	}
}
