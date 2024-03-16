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
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	n := readInt()
	arr := make([]int, n)
	arr[n-1] = n
	cnt := 1
	for i := n - 2; i >= 0; i-- {
		arr[i] = cnt
		cnt++
	}

	for _, v := range arr {
		writer.WriteString(strconv.Itoa(v))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
