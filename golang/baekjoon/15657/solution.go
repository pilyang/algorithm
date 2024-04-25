// back-tracking, sort

package main

import (
	"bufio"
	"os"
	"slices"
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

var (
	n, m       int
	indexStack []int
	nums       []int
)

func backTracking(depth int) {
	if depth == m {
		for _, index := range indexStack {
			writer.WriteString(strconv.Itoa(nums[index]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	startIndex := 0
	if depth > 0 {
		startIndex = indexStack[depth-1]
	}

	for i := startIndex; i < n; i++ {
		indexStack[depth] = i
		backTracking(depth + 1)
	}
}

func main() {
	defer writer.Flush()

	n, m = scanInt(), scanInt()

	nums = make([]int, n)
	indexStack = make([]int, m)

	for i := 0; i < n; i++ {
		nums[i] = scanInt()
	}

	slices.Sort(nums)

	backTracking(0)
}
