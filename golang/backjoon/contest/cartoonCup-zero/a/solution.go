package main

import (
	"bufio"
	"os"
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

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	s := readString()

	if s[0] != '"' || s[len(s)-1] != '"' || len(s) == 1 {
		writer.WriteString("CE\n")
		return
	}
	s = s[1 : len(s)-1]
	if len(strings.TrimSpace(s)) == 0 {
		writer.WriteString("CE\n")
		return
	}
	writer.WriteString(s)
	writer.WriteByte('\n')
}
