// https://www.acmicpc.net/problem/24831

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
	scanner.Split(bufio.ScanWords)
}

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	tc := readInt()
	for ; tc > 0; tc-- {
		x, y, a, b := readInt(), readInt(), readInt(), readInt()
		distance := y - x
		step := a + b
		if distance%step == 0 {
			writer.WriteString(strconv.Itoa(distance / step))
			writer.WriteByte('\n')
		} else {
			writer.WriteString("-1\n")
		}
	}
}
