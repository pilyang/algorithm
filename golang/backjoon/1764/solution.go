// https://www.acmicpc.net/problem/1764
// set

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

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func main() {
	defer writer.Flush()

	n, m := readInt(), readInt()

	lset := make(map[string]struct{}, n)
	lr := make([]string, 0, m)

	for i := 0; i < n; i++ {
		lset[readString()] = struct{}{}
	}

	for i := 0; i < m; i++ {
		r := readString()
		if _, exist := lset[r]; exist {
			lr = append(lr, r)
		}
	}

	slices.Sort(lr)
	writer.WriteString(strconv.Itoa(len(lr)))
	writer.WriteByte('\n')
	for _, name := range lr {
		writer.WriteString(name)
		writer.WriteByte('\n')
	}
}
