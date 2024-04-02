// https://www.acmicpc.net/problem/24832
// code force
// palindrome, implementation

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
	writer = bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
}

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func isPalindrome(s []rune) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	defer writer.Flush()

	n, m := readInt(), readInt()

	runes := make([][]rune, n)
	palindromeIdx := -1

	for i := 0; i < n; i++ {
		runes[i] = readRunes()
		if palindromeIdx == -1 && isPalindrome(runes[i]) {
			palindromeIdx = i
		}
	}
}
