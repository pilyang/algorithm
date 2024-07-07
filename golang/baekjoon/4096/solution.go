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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func stringToIntSlice(s string) []int {
	sl := make([]int, len(s))

	for i, c := range s {
		sl[i] = int(c - '0')
	}
	return sl
}

func pow10(n int) int {
	num := 10
	res := 1
	for n > 0 {
		if (n & 1) == 1 {
			res *= num
		}
		num = num * num
		n >>= 1
	}
	return res
}

func carryOver(nums []int, startIdx int) {
	for i := startIdx; i >= 0; i-- {
		if nums[i] >= 10 {
			nums[i] -= 10
			nums[i-1]++
		}
	}
}

func main() {
	defer writer.Flush()

	for {
		input := scanString()
		if input == "0" {
			break
		}

		nums := stringToIntSlice(input)
		l := len(nums)
		result := 0

		for i := 0; i < l/2; i++ {

			diff := nums[i] - nums[l-i-1]

			if diff == 0 {
				continue
			}

			// update the num slice
			nums[l-i-1] += diff
			// carry over
			if diff < 0 {
				diff = 10 + diff
				nums[l-i-2]++

				carryOver(nums, l-i-2)
				if nums[i] != nums[l-i-1] {
					diff++
					nums[l-i-1] = nums[i]
				}
			}

			// add diff to result
			result += diff * pow10(i)
		}

		writer.WriteString(strconv.Itoa(result) + "\n")
	}
}
