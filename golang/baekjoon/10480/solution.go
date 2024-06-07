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
	for i := 0; i < n; i++ {
		num := scanInt()
		writer.WriteString(strconv.Itoa(num) + " is ")
		if num%2 == 0 {
			writer.WriteString("even\n")
		} else {
			writer.WriteString("odd\n")
		}
	}
}
