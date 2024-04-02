// https://www.acmicpc.net/problem/1931
// greedy algorithm

package main

import (
	"bufio"
	"os"
	"slices"
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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func findMaxMeeting(meetingIdx int) int {
	var result int
	return result
}

func main() {
	defer writer.Flush()

	n := readInt()
	meetings := make([][2]int, n)
	for i := 0; i < n; i++ {
		meetings[i] = [2]int{readInt(), readInt()}
	}

	slices.SortFunc(meetings, func(a, b [2]int) int {
		if a[1] == b[1] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	var cnt, current int
	for _, meeting := range meetings {
		if current <= meeting[0] {
			current = meeting[1]
			cnt++
		}
	}
	writer.WriteString(strconv.Itoa(cnt))
	writer.WriteByte('\n')
}
