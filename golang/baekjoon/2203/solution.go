// https://www.acmicpc.net/problem/2003
// two pointer

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

	num, target := readInt(), readInt()
	numbers := make([]int, 0, num)
	for i := 0; i < num; i++ {
		numbers = append(numbers, readInt())
	}

	var sum, count, p1, p2 int
	for {
		if sum < target {
			p2++
			if p2 > num {
				break
			}
			sum += numbers[p2-1]
		} else {
			sum -= numbers[p1]
			p1++
		}
		if sum == target {
			count++
		}
	}

	wr.WriteString(strconv.Itoa(count) + "\n")
}
