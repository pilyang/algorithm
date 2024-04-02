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
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	table := make([][]string, 10)

	isPossible := false
	columnPossible := make([]bool, 10)
	for i := 0; i < 10; i++ {
		columnPossible[i] = true
	}

	for i := 0; i < 10; i++ {
		table[i] = make([]string, 10)
		rowCheck := true
		for j := 0; j < 10; j++ {
			table[i][j] = readString()
			if table[i][j] != table[i][0] {
				rowCheck = false
			}
			if table[0][j] != table[i][j] {
				columnPossible[j] = false
			}
		}
		if rowCheck {
			isPossible = true
		}
	}

	if !isPossible {
		for i := 0; i < 10; i++ {
			if columnPossible[i] {
				isPossible = true
				break
			}
		}
	}

	if isPossible {
		writer.WriteString("1\n")
	} else {
		writer.WriteString("0\n")
	}
}
