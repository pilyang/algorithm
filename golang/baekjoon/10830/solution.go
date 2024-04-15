// https://www.acmicpc.net/problem/10830
// math, fast exponentiation

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

const mod = 1_000

func multiply(matrix1, matrix2 [][]int) [][]int {
	result := make([][]int, len(matrix1))
	for rowIndex, mat1Row := range matrix1 {
		result[rowIndex] = make([]int, len(matrix2[0]))
		for colIndex := range matrix2[0] {
			for k := range mat1Row {
				result[rowIndex][colIndex] += matrix1[rowIndex][k] * matrix2[k][colIndex]
			}
			result[rowIndex][colIndex] %= mod
		}
	}

	return result
}

func pow(matrix [][]int, b int) [][]int {
	result := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		result[i] = make([]int, len(matrix))
		result[i][i] = 1
	}

	for b > 0 {
		if b&1 == 1 {
			result = multiply(result, matrix)
		}
		matrix = multiply(matrix, matrix)
		b >>= 1
	}

	return result
}

func main() {
	defer writer.Flush()

	n, b := scanInt(), scanInt()
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = scanInt()
		}
	}

	result := pow(matrix, b)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			writer.WriteString(strconv.Itoa(result[i][j]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
	}
}
