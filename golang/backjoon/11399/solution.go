// https://www.acmicpc.net/problem/11399
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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func main() {
	defer writer.Flush()

	n := readInt()
	arr := make([]int, n)
	for i := range arr {
		arr[i] = readInt()
	}

	slices.Sort(arr)

	var sum int
	for i := range arr {
		sum += arr[i] * (n - i)
	}

	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
