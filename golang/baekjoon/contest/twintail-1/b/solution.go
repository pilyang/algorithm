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

	n := readInt()

	var sumA, sumB int

	for i := 0; i < n; i++ {
		sumA += readInt()
	}
	for i := 0; i < n; i++ {
		sumB += readInt()
	}

	writer.WriteString(strconv.Itoa(sumB))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(sumA))
	writer.WriteByte('\n')
}
