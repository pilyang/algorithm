// stack, implemenatation

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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() int {
	if s.empty() {
		return -1
	}
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func main() {
	defer writer.Flush()

	s := make(stack, 0, 100)

	n := scanInt()

	instructions := make([]string, n)
	paramters := make([]int, n)
	for i := 0; i < n; i++ {
		instruction := scanString()
		instructions[i] = instruction
		param := -1
		switch instruction {
		case "PUSH", "IFZERO":
			param = scanInt()
		}
		paramters[i] = param
	}

	register := 0

	index := 0
	for instructions[index] != "DONE" {
		switch instructions[index] {
		case "PUSH":
			s.push(paramters[index])
		case "STORE":
			num := s.pop()
			register = num
		case "LOAD":
			s.push(register)
		case "PLUS":
			num1 := s.pop()
			num2 := s.pop()
			s.push(num1 + num2)
		case "TIMES":
			num1 := s.pop()
			num2 := s.pop()
			s.push(num1 * num2)
		case "IFZERO":
			num := s.pop()
			if num == 0 {
				index = paramters[index]
				continue
			}
		}
		index++
	}

	result := s.pop()
	writer.WriteString(strconv.Itoa(result))
}
