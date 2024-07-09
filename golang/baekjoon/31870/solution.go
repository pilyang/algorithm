// bubble sort

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func main() {
	defer writer.Flush()

	n := scanInt()

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = scanInt()
	}
	ascTarget := make([]int, n)
	copy(ascTarget, nums)

	ascCount := 0
	// bublesort num in ascending order
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if ascTarget[j] > ascTarget[j+1] {
				ascTarget[j], ascTarget[j+1] = ascTarget[j+1], ascTarget[j]
				ascCount++
			}
		}
	}

	descTarget := make([]int, n)
	copy(descTarget, nums)

	descCount := 0
	// bublesort num in descending order
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if descTarget[j] < descTarget[j+1] {
				descTarget[j], descTarget[j+1] = descTarget[j+1], descTarget[j]
				descCount++
			}
		}
	}
	descCount++

	minCount := ascCount
	if descCount < minCount {
		minCount = descCount
	}
	writer.WriteString(strconv.Itoa(minCount))
}
