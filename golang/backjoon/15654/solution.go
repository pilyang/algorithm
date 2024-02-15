// https://www.acmicpc.net/problem/15654
// backtracking dfs

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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func backtracking(val int) {
	// push
	stack = append(stack, val)

	if len(stack) == m+1 {
		for i := 1; i <= m; i++ {
			writer.WriteString(strconv.Itoa(stack[i]) + " ")
		}
		writer.WriteString("\n")
		return
	}

	for nextIdx, nextVal := range nums {
		if !isContain[nextIdx] {
			isContain[nextIdx] = true
			backtracking(nextVal)
			stack = stack[:len(stack)-1]
			isContain[nextIdx] = false
		}
	}
}

func backtrackingDepth(depth int) {
	if depth == m {
		for _, num := range stack {
			writer.WriteString(strconv.Itoa(num))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	for i, val := range nums {
		if !isContain[i] {
			stack = append(stack, val)
			isContain[i] = true

			backtrackingDepth(depth + 1)

			stack = stack[:len(stack)-1]
			isContain[i] = false
		}
	}
}

var (
	n, m        int
	nums, stack []int
	isContain   []bool
)

func main() {
	defer writer.Flush()

	n, m = readInt(), readInt()
	nums = make([]int, 0, n)
	stack = make([]int, 0, m)
	isContain = make([]bool, n)
	for i := 0; i < n; i++ {
		nums = append(nums, readInt())
	}

	slices.Sort(nums)

	// backtracking(-1)
	backtrackingDepth(0)
}
