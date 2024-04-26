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
	numCount   int
	indexStack []int
	nums       []int
)

func backTracking(depth, startIndex int) {
	if depth == m {
		for _, index := range indexStack {
			writer.WriteString(strconv.Itoa(nums[index]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	for i := startIndex; i < numCount; i++ {
		indexStack[depth] = i
		backTracking(depth+1, i)
	}
}

func main() {
	defer writer.Flush()

	n, m = scanInt(), scanInt()

	nums = make([]int, n)
	numSet := make(map[int]struct{}, n)
	indexStack = make([]int, m)

	numCount = 0
	for i := 0; i < n; i++ {
		num := scanInt()
		if _, ok := numSet[num]; ok {
			continue
		}
		nums[numCount] = num
		numSet[num] = struct{}{}
		numCount++
	}

	nums = nums[:numCount]
	slices.Sort(nums)

	backTracking(0, 0)
}
