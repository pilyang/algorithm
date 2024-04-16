// https://www.acmicpc.net/problem/11054
// dp

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

	dpAsc := make([]int, n+2)
	dpDesc := make([]int, n+2)

	// arr num will be placed in index 1 ~ n
	// index 0, n+1 for 0 value which prevents out-of-rance access when checking final result
	// index n+1 is for ascending only (dpDesc[n+1] = 0)
	// index 0 is for descending only (dpAsc[0] = 0)
	arr := make([]int, n+2)
	for i := 1; i <= n; i++ {
		arr[i] = scanInt()
	}

	for i := 1; i <= n; i++ {
		// for ascending dp
		dpAsc[i] = 1
		for j := 1; j < i; j++ {
			if arr[j] < arr[i] && dpAsc[j]+1 > dpAsc[i] {
				dpAsc[i] = dpAsc[j] + 1
			}
		}

		// for descending dp
		descIdx := n - i + 1
		dpDesc[descIdx] = 1
		for j := n; j > descIdx; j-- {
			if arr[j] < arr[descIdx] && dpDesc[j]+1 > dpDesc[descIdx] {
				dpDesc[descIdx] = dpDesc[j] + 1
			}
		}
	}

	result := 0
	for i := 0; i <= n; i++ {
		for j := i + 1; j <= n+1; j++ {
			if dpAsc[i]+dpDesc[j] > result && arr[i] != arr[j] {
				result = dpAsc[i] + dpDesc[j]
			}
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
