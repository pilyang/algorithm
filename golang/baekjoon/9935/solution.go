// https://www.acmicpc.net/problem/9935
// string, stack?

package main

import (
	"bufio"
	"os"
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

// read rune slice in one line
func readRunes() []rune {
	str, _ := reader.ReadString('\n')
	return []rune(strings.TrimSpace(str))
}

func main() {
	defer writer.Flush()

	origin := readRunes()
	bomb := readRunes()

	result := make([]rune, 0, len(origin))
	for _, r := range origin {
		result = append(result, r)

		if len(result) < len(bomb) {
			continue
		}

		match := true
		for i, b := range bomb {
			if result[len(result)-len(bomb)+i] != b {
				match = false
				break
			}
		}

		if match {
			result = result[:len(result)-len(bomb)]
		}
	}

	if len(result) == 0 {
		writer.WriteString("FRULA\n")
	} else {
		for _, r := range result {
			writer.WriteRune(r)
		}
		writer.WriteByte('\n')
	}
}
