// sort

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

func scanPositiveInt() int {
	scanner.Scan()
	ans := 0
	for _, c := range scanner.Bytes() {
		ans = ans*10 + int(c-'0')
	}
	return ans
}

func main() {
	defer writer.Flush()

	n, k := scanPositiveInt(), scanPositiveInt()
	scores := make([]int, n)

	for i := 0; i < n; i++ {
		scores[i] = scanPositiveInt()
	}

	slices.SortFunc(scores, func(a, b int) int {
		return b - a
	})

	sum := 0
	for i := 0; i < k; i++ {
		sum += scores[i]
	}
	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
