// https://www.acmicpc.net/problem/9375
// math, combination

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
	n, _ := strconv.Atoi(readString())
	return n
}

func solve() int {
	n := readInt()
	clothes := make(map[string]int)

	for i := 0; i < n; i++ {
		_, cat := readString(), readString()
		clothes[cat]++
	}

	result := 1
	for cat := range clothes {
		result *= clothes[cat] + 1
	}

	return result - 1
}

func main() {
	defer writer.Flush()

	t := readInt()
	for i := 0; i < t; i++ {
		result := solve()
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
	}
}
