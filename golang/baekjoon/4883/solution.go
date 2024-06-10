// brute force

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
	line := make([]int, n)
	set := make(map[int]struct{})

	for i := 0; i < n; i++ {
		num := scanInt()
		line[i] = num
		set[num] = struct{}{}
	}

	maxCount := 1
	for excludeNum := range set {
		prev := -1
		count := 1
		for _, num := range line {
			if num == excludeNum {
				continue
			}

			if prev == num {
				count++
				if count > maxCount {
					maxCount = count
				}
			} else {
				prev = num
				count = 1
			}
		}
	}

	writer.WriteString(strconv.Itoa(maxCount))
	writer.WriteByte('\n')
}
