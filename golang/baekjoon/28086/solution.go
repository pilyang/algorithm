// math, implementation, string

package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func octalStringToInt(octalString string) int64 {
	result, _ := strconv.ParseInt(octalString, 8, 64)
	return result
}

func seperateFormula(input string) (num1 int64, op string, num2 int64) {
	var opIndex int
	for i := 1; i < len(input); i++ {
		if input[i] < '0' || input[i] > '7' {
			opIndex = i
			break
		}
	}
	num1 = octalStringToInt(input[:opIndex])
	num2 = octalStringToInt(input[opIndex+1:])
	op = string(input[opIndex])
	return
}

func calculate(formula string) (int64, error) {
	num1, operator, num2 := seperateFormula(formula)
	var result int64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		}
		result = num1 / num2
		if num1*num2 < 0 && num1%num2 != 0 {
			result--
		}
	}

	return result, nil
}

func main() {
	defer writer.Flush()

	formula := scanString()
	result, err := calculate(formula)
	if err != nil {
		writer.WriteString("invalid\n")
		return
	}
	writer.WriteString(strconv.FormatInt(result, 8) + "\n")
}
