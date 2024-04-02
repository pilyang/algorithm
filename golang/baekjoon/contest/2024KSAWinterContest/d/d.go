package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func printVals(command string, nums *[]int) {
	fmt.Print(command)
	for _, num := range *nums {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func useFunc(nums *[]int) {
	printVals("? ", nums)
}

func answer(nums *[]int) {
	printVals("! ", nums)
}

// const modNum = 998244353

func main() {
	n := readInt()

	// nums := make([]int, n)
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	fResults := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = j*n + (i + 1)
		}

		useFunc(&mat[i])
		temp := readInt()

		fResults[i] = temp
	}

	results := make([]int, n)

	answer(&results)
}
