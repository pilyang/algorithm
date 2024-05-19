// math, brute force

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
}

func scanDigints() []int {
	scanner.Scan()
	bin := scanner.Bytes()
	n := len(bin)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = int(bin[i] - '0')
	}
	return b
}

func binaryToDecimal(b []int) int {
	n := len(b)
	decimal := 0
	for i := 0; i < n; i++ {
		decimal = decimal*2 + b[i]
	}
	return decimal
}

func DecimalToTerinary(decimal int) []int {
	t := make([]int, 0, 20)
	for decimal > 0 {
		t = append(t, decimal%3)
		decimal /= 3
	}

	n := len(t)
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}

	return t
}

func ternaryToDecimal(t []int) int {
	n := len(t)
	decimal := 0
	for i := 0; i < n; i++ {
		decimal = decimal*3 + t[i]
	}
	return decimal
}

// return the list of Decimal Value for the possible original binary values
func possibleBinaryOrigins(wrongBinary []int) []int {
	possibleOrigins := make([]int, len(wrongBinary)-1)
	for idx := 1; idx < len(wrongBinary); idx++ {
		wrongBinary[idx]++
		wrongBinary[idx] %= 2
		possibleOrigins[idx-1] = binaryToDecimal(wrongBinary)
		wrongBinary[idx]++
		wrongBinary[idx] %= 2
	}
	return possibleOrigins
}

func isOnlyOneDecimalDifferent(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	diffCount := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffCount++
			if diffCount > 1 {
				return false
			}
		}
	}

	return diffCount == 1
}

// 원래의 수로 가능한 모든 10진수를 2진수를 통해 구함
// 구한 수를 3진수로 변환하여 주어진 3진수와 비교
// 1자리 수가 다른 경우가 존재하면 해당 수가 답
func solution1(binary, ternary []int) int {
	if binary[0] == 0 {
		binary[0] = 1
		return binaryToDecimal(binary)
	}

	for _, decimal := range possibleBinaryOrigins(binary) {
		ter := DecimalToTerinary(decimal)
		if isOnlyOneDecimalDifferent(ter, ternary) {
			return decimal
		}
	}

	return 0
}

// 모든 2진수와 3진수를 10진수로 변환하여 비교
func solution2(binary, ternary []int) int {
	if binary[0] == 0 {
		binary[0] = 1
		return binaryToDecimal(binary)
	}

	valueSet := make(map[int]struct{})
	for _, decimal := range possibleBinaryOrigins(binary) {
		valueSet[decimal] = struct{}{}
	}

	for idx := range ternary {
		origin := ternary[idx]
		for i := 0; i < 2; i++ {
			ternary[idx]++
			ternary[idx] %= 3
			decimal := ternaryToDecimal(ternary)
			if _, ok := valueSet[decimal]; ok {
				return decimal
			}
		}
		ternary[idx] = origin
	}

	return 0
}

func main() {
	defer writer.Flush()

	binary := scanDigints()
	ternary := scanDigints()

	solution := solution1(binary, ternary)
	writer.WriteString(strconv.Itoa(solution))
	writer.WriteByte('\n')
}
