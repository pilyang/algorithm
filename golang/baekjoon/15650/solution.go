// https://www.acmicpc.net/problem/15650
// brute force
// stack + dfs
// backtracking

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
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func backtracking(stack *[]int, num int) {
	// push
	*stack = append(*stack, num)
	// print if stack is full
	if len(*stack) == m {
		for _, val := range *stack {
			writer.WriteString(strconv.Itoa(val) + " ")
		}
		writer.WriteString("\n")
		return
	}

	// try next numbers
	for next := num + 1; next <= n; next++ {
		backtracking(stack, next)
		// pop
		*stack = (*stack)[:len(*stack)-1]
	}
}

var n, m int

func main() {
	defer writer.Flush()

	n, m = readInt(), readInt()

	stack := make([]int, 0, m)

	for i := 1; i <= n-m+1; i++ {
		backtracking(&stack, i)
		stack = stack[:0]
	}
}
