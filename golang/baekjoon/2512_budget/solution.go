package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acmicpc.net/problem/2512

func main() {
	reader := *bufio.NewReader(os.Stdin)
	writer := *bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanf(&reader, "%d\n", &n)

	var budgets []int = make([]int, n)
	for index := 0; index < n; index++ {
		fmt.Fscanf(&reader, "%d ", &budgets[index])
	}

	var maxBudget int
	fmt.Fscanf(&reader, "%d\n", &maxBudget)

	result := findMaximumBudget(&budgets, maxBudget)

	fmt.Fprintf(&writer, "%d\n", result)
	writer.Flush()
}

func findMaximumBudget(budgets *[]int, maxBudget int) int {
	low, high := 0, 0
	for _, budget := range *budgets {
		if budget > high {
			high = budget
		}
	}

	if makeSum(budgets, high) <= maxBudget {
		return high
	}

	for high-low > 1 {
		mid := (low + high) / 2
		switch sum := makeSum(budgets, mid); {
		case sum == maxBudget:
			return mid
		case sum < maxBudget:
			low = mid
		case sum > maxBudget:
			high = mid
		}
	}

	return low
}

func makeSum(budgets *[]int, limit int) int {
	sum := 0
	for _, budget := range *budgets {
		if budget > limit {
			sum += limit
		} else {
			sum += budget
		}
	}

	return sum
}
