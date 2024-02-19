// https://www.acmicpc.net/problem/1918
// stack

package main

import (
	"bufio"
	"os"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

type CalcStack []rune

func (s *CalcStack) push(r rune) {
	*s = append(*s, r)
}

func (s *CalcStack) pop() rune {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *CalcStack) top() rune {
	if len(*s) == 0 {
		return ' '
	}
	return (*s)[len(*s)-1]
}

func main() {
	defer writer.Flush()

	runes := readRunes()

	result := make([]rune, 0, len(runes))
	calckStack := make(CalcStack, 0, len(runes))
	// var bracketCnt int

	for _, r := range runes {
		switch r {
		case '(':
			calckStack.push(r)
		case ')':
			for calckStack[len(calckStack)-1] != '(' {
				result = append(result, calckStack.pop())
			}
			calckStack.pop()
		case '+', '-':
			for len(calckStack) > 0 && calckStack.top() != '(' {
				result = append(result, calckStack.pop())
			}
			calckStack.push(r)
		case '*', '/':
			if r2 := calckStack.top(); r2 == '*' || r2 == '/' {
				result = append(result, calckStack.pop())
			}
			calckStack.push(r)
		default:
			result = append(result, r)
		}
	}

	for len(calckStack) > 0 {
		result = append(result, calckStack.pop())
	}

	for _, r := range result {
		writer.WriteRune(r)
	}
	writer.WriteByte('\n')
}
