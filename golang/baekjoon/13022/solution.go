// string, implementation

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

func scanRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

var (
	w          = 'w'
	olfPattern = []rune{'o', 'l', 'f'}
)

func isUnitWolf(runes []rune, startIdx int) (int, bool) {
	if startIdx+3 > len(runes) {
		return 0, false
	}

	index := startIdx
	count := 0
	for index < len(runes) && runes[index] == w {
		count++
		index++
	}

	if count == 0 {
		return 0, false
	}

	for _, nextPattern := range olfPattern {
		tempCount := 0
		for index < len(runes) && runes[index] == nextPattern {
			tempCount++
			index++
		}
		if tempCount != count {
			return 0, false
		}
	}

	return index, true
}

func isWolf(runes []rune) bool {
	startIdx := 0
	var isUnit bool
	for startIdx < len(runes) {
		startIdx, isUnit = isUnitWolf(runes, startIdx)
		if !isUnit {
			return false
		}
	}

	return true
}

func main() {
	defer writer.Flush()

	wolf := scanRunes()
	if isWolf(wolf) {
		writer.WriteString("1\n")
	} else {
		writer.WriteString("0\n")
	}
}
