// greedy

package main

import (
	"bufio"
	"os"
	"slices"
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

	n := scanInt()
	houses := make([]int, n)
	for i := 0; i < n; i++ {
		houses[i] = scanInt()
	}

	slices.Sort(houses)

	index := (n - 1) / 2

	writer.WriteString(strconv.Itoa(houses[index]))
	writer.WriteByte('\n')
}
