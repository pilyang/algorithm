// https://www.acmicpc.net/problem/1874
// stack

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func solve(target []int) string {
	var builder strings.Builder

	stack := make([]int, 0, len(target))
	num := 1
	for _, t := range target {
		for num <= t {
			stack = append(stack, num)
			builder.WriteString("+\n")
			num++
		}

		if stack[len(stack)-1] != t {
			return "NO\n"
		}

		stack = stack[:len(stack)-1]
		builder.WriteString("-\n")
	}

	return builder.String()
}

func main() {
	defer writer.Flush()

	n := scanInt()
	target := make([]int, n)
	for i := range target {
		target[i] = scanInt()
	}
	result := solve(target)
	writer.WriteString(result)
}
