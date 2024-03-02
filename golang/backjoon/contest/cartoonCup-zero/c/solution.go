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

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func query(s string) int {
	writer.WriteString("? ")
	writer.WriteString(s)
	writer.WriteByte('\n')
	writer.Flush()

	return readInt()
}

func answer(s string) {
	writer.WriteString("! ")
	writer.WriteString(s)
	writer.WriteByte('\n')
	writer.Flush()
}

func main() {
	defer writer.Flush()

	n := readInt()
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = readString()
	}

	for tc := 0; tc < 2; tc++ {
		for i := 0; i < n; i++ {
			if query(names[i]) == 1 {
				answer(names[i])
				return
			}
		}
	}
}
