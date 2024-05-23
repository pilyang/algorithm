// implementation

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

func solve() int {
	n := scanInt()
	count := make([]int, 1001)

	maxNum := 0
	for i := 0; i < n; i++ {
		num := scanInt()
		count[num]++

		if count[num] > count[maxNum] {
			maxNum = num
		} else if count[num] == count[maxNum] && num < maxNum {
			maxNum = num
		}
	}

	return maxNum
}

func main() {
	defer writer.Flush()

	tc := scanInt()
	for tc > 0 {
		tc--
		writer.WriteString(strconv.Itoa(solve()))
		writer.WriteByte('\n')
	}
}
