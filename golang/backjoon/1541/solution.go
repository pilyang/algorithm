// https://www.acmicpc.net/problem/1541

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

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

func extravtNumber(runes []rune, start int) (int, int) {
	num := 0
	for i := start; i < len(runes); i++ {
		if runes[i] == '+' || runes[i] == '-' {
			return num, i
		}
		num = num*10 + int(runes[i]-'0')
	}
	return num, len(runes)
}

func main() {
	defer writer.Flush()

	runes := readRunes()
	var cursor, num, sum int
	isMeatMinus := false
	for {
		num, cursor = extravtNumber(runes, cursor)

		if isMeatMinus {
			sum -= num
		} else {
			sum += num
		}

		if cursor >= len(runes) {
			break
		}

		if !isMeatMinus && runes[cursor] == '-' {
			isMeatMinus = true
		}
		cursor++
	}

	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
