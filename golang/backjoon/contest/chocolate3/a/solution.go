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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

func main() {
	defer writer.Flush()

	n := readInt()
	for n > 0 {
		n--
		var notCount, num int
		rs := readRunes()
		for _, r := range rs {
			if r == '!' {
				notCount++
			} else {
				num = int(r - '0')
				break
			}
		}

		if len(rs) > notCount+1 {
			num = 1
		}
		notCount %= 2
		if notCount == 1 {
			num ^= 1
		}
		writer.WriteString(strconv.Itoa(num))
		writer.WriteByte('\n')
	}
}
