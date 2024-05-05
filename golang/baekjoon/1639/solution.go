// brute force, prefix sum
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

func scanRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

func main() {
	defer writer.Flush()

	runes := scanRunes()

	prefixSum := make([]int, len(runes)+1)
	for i, rune := range runes {
		prefixSum[i+1] = int(rune-'0') + prefixSum[i]
	}

	result := 0

	for from := 1; from < len(prefixSum); from++ {
		for to := from + 1; to < len(prefixSum); to += 2 {
			currentLen := to - from + 1
			if prefixSum[to]-prefixSum[to-currentLen/2] == prefixSum[to-currentLen/2]-prefixSum[from-1] {
				result = max(result, currentLen)
			}
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
