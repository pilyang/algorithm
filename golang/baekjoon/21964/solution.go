package main

import (
	"bufio"
	"os"
)

var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func scanString() string {
	s, _ := reader.ReadString('\n')
	return s[:len(s)-1]
}

func main() {
	defer writer.Flush()

	_ = scanString()
	s := scanString()

	writer.WriteString(s[len(s)-5:])
}
