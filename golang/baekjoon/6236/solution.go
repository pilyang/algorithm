package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acmicpc.net/problem/6236

func main() {
	reader := *bufio.NewReader(os.Stdin)
	writer := *bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int
	fmt.Fscanf(&reader, "%d %d\n", &n, &m)

	moneyPlan := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(&reader, "%d\n", &moneyPlan[i])
	}

	result := findMinMoney(moneyPlan, m)
	fmt.Printf("%d\n", result)
}

func findMinMoney(moneyPlan []int, m int) int {
	var min, max int
	for _, money := range moneyPlan {
		if money > min {
			min = money
		}
		max += money
	}
	min--
	max++
	for max-min > 1 {
		mid := (min + max) / 2
		if isPossibleWith(moneyPlan, m, mid) {
			max = mid
		} else {
			min = mid
		}
	}

	return max
}

func isPossibleWith(moneyPlan []int, m, k int) bool {
	var count, sum int
	for _, money := range moneyPlan {
		if sum+money > k {
			count++
			sum = money
		} else {
			sum += money
		}
		if count == m {
			return false
		}
	}

	return true
}
