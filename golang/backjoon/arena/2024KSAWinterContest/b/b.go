package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func readRunes() *[]rune {
	val, _ := reader.ReadString('\n')
	rune := []rune(strings.TrimSpace(val))
	return &rune
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findResult(runes *[]rune) int {
	n := len(*runes)

	ksa := [3]rune{'K', 'S', 'A'}
	var subSize [3]int
	var nextKSA [3]int = [3]int{0, 1, 2}
	for _, n := range *runes {
		for i := 0; i < 3; i++ {
			if n == ksa[nextKSA[i]] {
				subSize[i]++
				nextKSA[i] = (nextKSA[i] + 1) % 3
			}
		}
	}

	var results [3]int
	for i := 0; i < 3; i++ {
		results[i] = (n - subSize[i]) * 2
		if n-subSize[i] < i {
			results[i] += 2
		}
	}

	return minInt(minInt(results[0], results[1]), results[2])
}

func main() {
	defer writer.Flush()

	input := readRunes()
	n := len(*input)
	var sub [3][]rune
	for i := range sub {
		sub[i] = make([]rune, 0, n)
	}

	writer.WriteString(strconv.Itoa(findResult(input)))
	writer.WriteByte('\n')
}
