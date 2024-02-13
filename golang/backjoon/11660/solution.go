// https://www.acmicpc.net/problem/11660
// Prefix Sum

package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func main() {
	defer wr.Flush()

	n, m := readInt(), readInt()

	nums := make([][]int, n)
	prefixSum := make([][]int, n+1)
	prefixSum[0] = make([]int, n+1)
	for i := range nums {
		nums[i] = make([]int, n)
		prefixSum[i+1] = make([]int, n+1)
		for j := range nums[i] {
			nums[i][j] = readInt()
			prefixSum[i+1][j+1] = prefixSum[i][j+1] + prefixSum[i+1][j] + nums[i][j] - prefixSum[i][j]
		}
	}

	var x1, y1, x2, y2 int
	var size int
	for i := 0; i < m; i++ {
		x1, y1, x2, y2 = readInt()-1, readInt()-1, readInt()-1, readInt()-1
		size = prefixSum[x2+1][y2+1] - prefixSum[x1][y2+1] - prefixSum[x2+1][y1] + prefixSum[x1][y1]
		wr.WriteString(strconv.Itoa(size) + "\n")
	}
}
