// https://www.acmicpc.net/problem/9095
// dp

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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func main() {
	defer writer.Flush()

	results := [10]int{1, 2, 4}
	for i := 3; i < 10; i++ {
		results[i] = results[i-1] + results[i-2] + results[i-3]
	}

	t := readInt()
	for t > 0 {
		t--
		n := readInt()
		writer.WriteString(strconv.Itoa(results[n-1]))
		writer.WriteByte('\n')
	}
}
