// greedy

package main

import (
	"bufio"
	"math"
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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	minColorCount := math.MaxInt

	nums := make([]int, 3)
	for i := range nums {
		nums[i] = scanInt()
		if minColorCount > nums[i] {
			minColorCount = nums[i]
		}
	}

	count := minColorCount
	for i := range nums {
		nums[i] -= minColorCount
		count += nums[i] / 3
		nums[i] %= 3
	}

	// 000,001,011 -> 1
	// 111 -> 1
	// 002 -> 1
	//
	// 022, 012, ... -> 2
	// 222 -> 2
	var sum int
	isAllSame := true
	prev := nums[0]
	for _, n := range nums {
		sum += n
		if prev != n {
			isAllSame = false
		}
	}
	count += sum / 3
	if !isAllSame {
		count++
	}

	writer.WriteString(strconv.Itoa(count))
}
