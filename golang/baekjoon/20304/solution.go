// https://www.acmicpc.net/problem/20304

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
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	maxPassword := readInt()
	tc := readInt()
	triedNums := make([]int, tc)
}
