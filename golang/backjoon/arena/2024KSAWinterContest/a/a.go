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

func findNonPrime(numbers *[]int) (int, bool) {
	for _, num := range *numbers {
		if num == 0 {
			return 0, true
		}
		if num == 1 {
			return 1, true
		}
		if num%2 == 0 && num != 2 {
			return num, true
		}

		return num*10 + num, true
	}

	return 0, false
}

func main() {
	defer writer.Flush()

	n := readInt()
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = readInt()
	}

	if num, exist := findNonPrime(&numbers); exist {
		writer.WriteString("YES\n")
		writer.WriteString(strconv.Itoa(num))
	} else {
		writer.WriteString("NO\n")
	}
	writer.WriteByte('\n')
}
