// 누적합?

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	n := scanInt()

	cumulativeMistakes := make([]int, n+1)
	prevDifficulty, currentDifficulty := 0, scanInt()

	for i := 1; i < n; i++ {
		prevDifficulty = currentDifficulty
		currentDifficulty = scanInt()

		cumulativeMistakes[i] = cumulativeMistakes[i-1]
		if prevDifficulty > currentDifficulty {
			cumulativeMistakes[i]++
		}
	}
	cumulativeMistakes[n] = cumulativeMistakes[n-1]

	for tc := scanInt(); tc > 0; tc-- {
		l, r := scanInt(), scanInt()
		writer.WriteString(strconv.Itoa(cumulativeMistakes[r-1] - cumulativeMistakes[l-1]))
		writer.WriteByte('\n')
	}
}
