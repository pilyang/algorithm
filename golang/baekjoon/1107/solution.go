// https://www.acmicpc.net/problem/1107
// brute force

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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func isPossible(num int) bool {
	if num == 0 {
		return isWork[0]
	}
	for num > 0 {
		if !isWork[num%10] {
			return false
		}
		num /= 10
	}
	return true
}

var isWork = []bool{true, true, true, true, true, true, true, true, true, true}

func main() {
	defer writer.Flush()

	target := readInt()
	brokenCount := readInt()
	for i := 0; i < brokenCount; i++ {
		isWork[readInt()] = false
	}

	result := abs(target - 100)

	if brokenCount == 10 {
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
		return
	}

	for i := target; i <= 1000000; i++ {
		if isPossible(i) {
			result = min(result, abs(target-i)+len(strconv.Itoa(i)))
			break
		}
	}
	for i := target; i >= 0; i-- {
		if isPossible(i) {
			result = min(result, abs(target-i)+len(strconv.Itoa(i)))
			break
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
