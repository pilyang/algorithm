// ad hoc

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	n := scanInt()

	if n == 1 {
		writer.WriteString("1\n1 1\n")
		return
	}

	result := 2*n - 2
	writer.WriteString(strconv.Itoa(result) + "\n")

	row := 1
	for col := 1; col <= n; col++ {
		writer.WriteString(strconv.Itoa(row) + " " + strconv.Itoa(col) + "\n")
	}

	row = n
	for col := 2; col <= n-1; col++ {
		writer.WriteString(strconv.Itoa(row) + " " + strconv.Itoa(col) + "\n")
	}
}
