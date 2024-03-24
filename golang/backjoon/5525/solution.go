// https://www.acmicpc.net/problem/5525
// string

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func readString() string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt() int {
	n, _ := strconv.Atoi(readString())
	return n
}

func readRunes() []rune {
	return []rune(readString())
}

func main() {
	defer writer.Flush()

	n := readInt()
	m := readInt()
	s := readRunes()

	var cursor, count, result int

	for cursor < m-2 {
		if s[cursor] == 'I' && s[cursor+1] == 'O' && s[cursor+2] == 'I' {
			count++
			if count == n {
				result++
				count--
			}

			cursor += 2
			continue
		}

		count = 0
		cursor++
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
