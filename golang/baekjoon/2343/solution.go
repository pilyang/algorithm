package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acmicpc.net/problem/2343

func main() {
	reader := *bufio.NewReader(os.Stdin)
	writer := *bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int
	fmt.Fscanf(&reader, "%d %d\n", &n, &m)

	lessons := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(&reader, "%d ", &lessons[i])
	}

	result := binarySearch(&lessons, m)
	fmt.Fprintf(&writer, "%d\n", result)
}

func binarySearch(lessons *[]int, cdCount int) int {
	var min, max int
	for _, lesson := range *lessons {
		max += lesson
		if min < lesson {
			min = lesson
		}
	}
	min--
	max++

	for max-min > 1 {
		mid := (min + max) / 2
		if isPossible(lessons, cdCount, mid) {
			max = mid
		} else {
			min = mid
		}
	}

	return max
}

func isPossible(lessons *[]int, cdCount int, cdSize int) bool {
	sum := 0
	for _, lesson := range *lessons {
		if sum+lesson > cdSize {
			cdCount--
			sum = lesson
		} else {
			sum += lesson
		}
		if cdCount == 0 {
			return false
		}
	}

	return true
}
