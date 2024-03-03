// https://www.acmicpc.net/problem/1629
// math, pow

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

func nextInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

// a ^ n % modNum
// (a x b) mod m = (a mod m x b mod m) mod m
// (a^n) mod m = (a^(n/2) x a^(n/2)) mod m
func main() {
	defer writer.Flush()

	a, n, modNum := nextInt(), nextInt(), nextInt()

	result := 1
	for n > 0 {
		if n%2 == 1 {
			result = (result * a) % modNum
		}
		// calculate a^2
		a = (a * a) % modNum
		n /= 2
	}
	result %= modNum

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
