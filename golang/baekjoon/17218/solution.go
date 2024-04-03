// https://www.acmicpc.net/problem/17218
// lcs, dp

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(r1, r2 []rune) []rune {
	lcsdp := make([][]int, len(r1)+1)
	lcsdp[0] = make([]int, len(r2)+1)

	for i1 := 1; i1 <= len(r1); i1++ {
		lcsdp[i1] = make([]int, len(r2)+1)
		for i2 := 1; i2 <= len(r2); i2++ {
			if r1[i1-1] == r2[i2-1] {
				lcsdp[i1][i2] = lcsdp[i1-1][i2-1] + 1
			} else {
				lcsdp[i1][i2] = max(lcsdp[i1-1][i2], lcsdp[i1][i2-1])
			}
		}
	}

	result := make([]rune, lcsdp[len(r1)][len(r2)])
	cursor1, cursor2 := len(r1), len(r2)
	for i := len(result) - 1; i >= 0; {
		if lcsdp[cursor1][cursor2] == lcsdp[cursor1-1][cursor2] {
			cursor1--
			continue
		}

		if lcsdp[cursor1][cursor2] == lcsdp[cursor1][cursor2-1] {
			cursor2--
			continue
		}

		result[i] = r1[cursor1-1]
		cursor1--
		cursor2--
		i--
	}

	return result
}

func main() {
	defer writer.Flush()

	r1, r2 := scanRunes(), scanRunes()

	result := solve(r1, r2)
	for _, rr := range result {
		writer.WriteRune(rr)
	}
	writer.WriteByte('\n')
}
