// https://www.acmicpc.net/problem/1546

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

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	n := readInt()
	scores := make([]int, n)
	maxScore := 0
	for i := 0; i < n; i++ {
		scores[i] = readInt()
		if maxScore < scores[i] {
			maxScore = scores[i]
		}
	}

	var sum, result float64
	for _, score := range scores {
		sum += float64(score) / float64(maxScore) * 100
	}
	result = sum / float64(n)

	writer.WriteString(strconv.FormatFloat(result, 'f', 3, 64))
	writer.WriteByte('\n')
}
