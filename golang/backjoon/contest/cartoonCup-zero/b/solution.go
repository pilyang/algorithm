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

	count, item := readInt(), readString()

	var res int
	for count > 0 {
		n, c := readString(), readInt()
		ns := strings.Split(n, "_")
		for _, s := range ns {
			if s == item {
				res += c
				break
			}
		}
		count--
	}

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}
