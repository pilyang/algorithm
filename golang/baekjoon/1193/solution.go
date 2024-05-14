// math, implementation

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func findGroupAndIndex(num int) (group, index int) {
	group = 1
	for num > group {
		num -= group
		group++
	}
	return group, num
}

func main() {
	defer writer.Flush()

	num := scanInt()

	group, index := findGroupAndIndex(num)
	if group%2 == 0 {
		writer.WriteString(strconv.Itoa(index))
		writer.WriteByte('/')
		writer.WriteString(strconv.Itoa(group - index + 1))
		writer.WriteByte('\n')
	} else {
		writer.WriteString(strconv.Itoa(group - index + 1))
		writer.WriteByte('/')
		writer.WriteString(strconv.Itoa(index))
		writer.WriteByte('\n')
	}
}
