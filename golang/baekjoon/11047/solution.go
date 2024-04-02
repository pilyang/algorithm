// https://www.acmicpc.net/problem/11047
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

	n, k := scanInt(), scanInt()
	coins := make([]int, n)
	startingPoint := 0
	for i := 0; i < n; i++ {
		coins[i] = scanInt()
		if coins[i] <= k {
			startingPoint = i
		}
	}

	count := 0
	for i := startingPoint; i >= 0; i-- {
		count += k / coins[i]
		k %= coins[i]
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
