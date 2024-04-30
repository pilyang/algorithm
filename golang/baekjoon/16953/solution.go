// bfs, greedy

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

func findResultGreedy(from, to int) int {
	count := 1
	for from < to {
		if to%10 == 1 {
			to /= 10
		} else if to%2 == 0 {
			to /= 2
		} else {
			return -1
		}
		count++
	}

	if from == to {
		return count
	}
	return -1
}

func findResultBFS(from, to int) int {
	queue := make([][2]int, 0, 50)
	queue = append(queue, [2]int{from, 0})

	var current [2]int
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if current[0] == to {
			return current[1] + 1
		}
		if current[0] > to {
			continue
		}

		queue = append(queue, [2]int{current[0] * 2, current[1] + 1})
		queue = append(queue, [2]int{current[0]*10 + 1, current[1] + 1})
	}

	return -1
}

func main() {
	defer writer.Flush()

	from, to := scanInt(), scanInt()

	result := findResultGreedy(from, to)
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
