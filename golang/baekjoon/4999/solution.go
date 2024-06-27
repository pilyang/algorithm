package main

import (
	"bufio"
	"os"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func scnaString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	possible, required := scnaString(), scnaString()
	if len(possible) >= len(required) {
		writer.WriteString("go\n")
	} else {
		writer.WriteString("no\n")
	}
}
