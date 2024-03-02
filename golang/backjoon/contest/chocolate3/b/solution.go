package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func getPalindrome(d int) string {
	if d == 1 {
		return "0"
	}
	mid := strings.Repeat("2", d-2)
	return "1" + mid + "1"
}

func main() {
	defer writer.Flush()

	n := readInt()

	for n > 0 {
		n--
		writer.WriteString(getPalindrome(readInt()))
		writer.WriteByte('\n')
	}
}
