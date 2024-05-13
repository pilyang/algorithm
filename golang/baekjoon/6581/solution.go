// scan until EOF

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
	scanner.Split(bufio.ScanWords)

	writer = bufio.NewWriter(os.Stdout)
}

func scanString() (string, bool) {
	if scanner.Scan() {
		return scanner.Text(), true
	}
	return "", false
}

const (
	space = " "
	eol   = "\n"
	tab   = "\t"

	tagBr = "<br>"
	tagHr = "<hr>"
)

var header = strings.Repeat("-", 80)

func main() {
	defer writer.Flush()

	input := make([]string, 0)
	for {
		word, ok := scanString()
		if !ok {
			break
		}
		// if word == "end" {
		// 	break
		// }
		input = append(input, word)
	}

	lineCount := 0
	for _, word := range input {
		word = strings.TrimSpace(word)
		switch word {
		case tagBr:
			writer.WriteString(eol)
			lineCount = 0
		case tagHr:
			if lineCount > 0 {
				writer.WriteString(eol)
			}
			writer.WriteString(header)
			writer.WriteString(eol)
			lineCount = 0
		case space, tab, eol:
			continue
		default:
			if lineCount+len(word)+1 > 80 {
				writer.WriteString(eol)
				lineCount = 0
			}
			if lineCount > 0 {
				writer.WriteString(space)
				lineCount++
			}
			writer.WriteString(word)
			lineCount += len(word)
		}
	}
	if lineCount > 0 {
		writer.WriteString(eol)
	}
}
