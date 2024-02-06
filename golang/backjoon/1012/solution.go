// https://www.acmicpc.net/problem/1012

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	tryCount, _ := strconv.Atoi(readInput(scanner)[0])
	results := make([]int, tryCount)
	for i := 0; i < tryCount; i++ {
		lineInput := readInput(scanner)
		column, _ := strconv.Atoi(lineInput[0])
		row, _ := strconv.Atoi(lineInput[1])
		k, _ := strconv.Atoi(lineInput[2])
		graph := make([][]bool, row)
		for gi := range graph {
			graph[gi] = make([]bool, column)
		}
		for ii := 0; ii < k; ii++ {
			input := readInput(scanner)
			x, _ := strconv.Atoi(input[0])
			y, _ := strconv.Atoi(input[1])
			graph[y][x] = true
		}

		results[i] = countComponent(&graph)
	}

	for _, result := range results {
		str := strconv.Itoa(result)
		writer.WriteString(str + "\n")
	}
}

func countComponent(graph *[][]bool) int {
	// initialize isVisit
	row, col := len(*graph), len((*graph)[0])
	isVisit := make([][]bool, row)
	for i := range isVisit {
		isVisit[i] = make([]bool, col)
	}

	count := 0

	for x := 0; x < col; x++ {
		for y := 0; y < row; y++ {
			if (*graph)[y][x] && !isVisit[y][x] {
				dfs(x, y, graph, &isVisit)
				count++
			}
		}
	}

	return count
}

func dfs(startX, startY int, graph *[][]bool, isVisit *[][]bool) {
	row, col := len(*graph), len((*graph)[0])
	stack := NewStack(row * col)
	stack.Push([2]int{startX, startY})
	for !stack.IsEmpty() {
		c, _ := stack.Pop()
		nexts := [4][2]int{
			{c[0], c[1] + 1},
			{c[0], c[1] - 1},
			{c[0] + 1, c[1]},
			{c[0] - 1, c[1]},
		}
		for _, n := range nexts {
			x, y := n[0], n[1]
			if x >= 0 && x < col && y >= 0 && y < row && (*graph)[y][x] && !(*isVisit)[y][x] {
				stack.Push(n)
				(*isVisit)[y][x] = true
			}
		}
	}
}

type Stack [][2]int

func NewStack(cap int) *Stack {
	s := Stack(make([][2]int, 0, cap))
	return &s
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val [2]int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() ([2]int, bool) {
	if s.IsEmpty() {
		return [2]int{}, false
	}
	index := len(*s) - 1
	el := (*s)[index]
	(*s) = (*s)[:index]
	return el, true
}
