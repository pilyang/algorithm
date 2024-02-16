// https://www.acmicpc.net/problem/15663
// backtracing

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

var (
	n, m      int
	nums      []int
	stack     []string
	isContain []bool
)

func backtracing(depth int) {
	if depth == m {
		for _, s := range stack {
			writer.WriteString(s)
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		return
	}

	prev := -1
	for i, num := range nums {
		if num == prev {
			continue
		}
		if !isContain[i] {
			isContain[i] = true
			stack = append(stack, strconv.Itoa(num))

			backtracing(depth + 1)

			stack = stack[:len(stack)-1]
			isContain[i] = false
			prev = num
		}
	}
}

func main() {
	defer writer.Flush()

	n, m = readInt(), readInt()
	nums = make([]int, n)
	stack = make([]string, 0, m)
	isContain = make([]bool, n)

	for i := range nums {
		nums[i] = readInt()
	}
	slices.Sort(nums)

	backtracing(0)
}
