package main

// https://www.acmicpc.net/problem/2805

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := *bufio.NewReader(os.Stdin)
	writer := *bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanf(&reader, "%d %d\n", &n, &m)
	input, _ := reader.ReadString('\n')
	treeInputSlice := strings.Split(strings.TrimSpace(input), " ")

	trees := make([]int, n)
	for i, input := range treeInputSlice {
		trees[i], _ = strconv.Atoi(input)
	}

	result := findMaximumHeight(&trees, m)

	fmt.Fprintf(&writer, "%d\n", result)
	writer.Flush()
}

func findMaximumHeight(trees *[]int, required int) int {
	low, high := 0, 0
	for _, treeHeight := range *trees {
		high += treeHeight
	}

	// binary search
	var mid int
	for high-low > 1 {
		mid = (low + high) / 2
		switch sum := addSum(trees, mid); {
		case sum == required:
			return mid
		case sum < required:
			high = mid
		case sum > required:
			low = mid
		}
	}

	return low
}

func addSum(trees *[]int, height int) int {
	sum := 0
	for _, treeHeight := range *trees {
		if treeHeight > height {
			sum += treeHeight - height
		}
	}

	return sum
}
