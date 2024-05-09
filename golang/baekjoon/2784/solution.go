// brute-force

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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

var docs []string

func solve() ([]string, bool) {
	result := make([]string, 3)

	var backTracking func(int, int) bool
	backTracking = func(depth, index int) bool {
		if depth == 3 {
			return validate(result)
		}

		result[depth] = docs[index]

		for i := 0; i < len(docs); i++ {
			if i == index {
				continue
			}

			if backTracking(depth+1, i) {
				return true
			}
		}
		return false
	}

	for startIdx := 0; startIdx < len(docs); startIdx++ {
		if possible := backTracking(0, startIdx); possible {
			return result, true
		}
	}

	return nil, false
}

func validate(value []string) bool {
	isCheckedRow := make([]bool, 3)
	isCheckedCol := make([]bool, 3)

	for _, doc := range docs {
		checked := false

		for i := 0; i < 3; i++ {
			if !isCheckedRow[i] && value[i] == doc {
				checked = true
				isCheckedRow[i] = true
				break
			}

			if !isCheckedCol[i] && value[0][i] == doc[0] && value[1][i] == doc[1] && value[2][i] == doc[2] {
				checked = true
				isCheckedCol[i] = true
				break
			}
		}

		if !checked {
			return false
		}
	}

	return true
}

func main() {
	defer writer.Flush()

	docs = make([]string, 6)
	for i := 0; i < 6; i++ {
		docs[i] = scanString()
	}

	if result, ok := solve(); ok {
		for _, line := range result {
			writer.WriteString(line)
			writer.WriteByte('\n')
		}
	} else {
		writer.WriteString("0\n")
	}
}
