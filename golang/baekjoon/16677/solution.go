// string, implementation

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
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

var original string

func countAddedRune(target string) (int, bool) {
	if len(original) >= len(target) {
		return 0, false
	}

	var targetIdx int
	for _, r := range original {
		for target[targetIdx] != byte(r) {
			targetIdx++
			if targetIdx >= len(target) {
				return 0, false
			}
		}
	}

	return len(target) - len(original), true
}

func main() {
	defer writer.Flush()

	original = scanString()

	n := scanInt()

	var max float64
	var result string

	for i := 0; i < n; i++ {
		target, score := scanString(), scanInt()

		addedCount, ok := countAddedRune(target)
		if !ok {
			continue
		}

		correct := float64(score) / float64(addedCount)

		if correct > max {
			max = correct
			result = target
		}
	}

	if result == "" {
		writer.WriteString("No Jam")
	} else {
		writer.WriteString(result)
	}
	writer.WriteByte('\n')
}
