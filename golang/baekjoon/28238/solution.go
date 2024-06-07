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
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

type Days [5]bool

func main() {
	defer writer.Flush()

	n := scanInt()
	students := make([]Days, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 5; j++ {
			students[i][j] = scanString() == "1"
		}
	}

	maxCount := 0
	var resultDays Days
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			count := 0
			for _, studnet := range students {
				if studnet[i] && studnet[j] {
					count++
				}
			}
			if count > maxCount {
				maxCount = count
				resultDays = Days{false, false, false, false, false}
				resultDays[i] = true
				resultDays[j] = true
			}
		}
	}

	writer.WriteString(strconv.Itoa(maxCount) + "\n")
	if maxCount > 0 {
		for _, day := range resultDays {
			if day {
				writer.WriteString("1 ")
			} else {
				writer.WriteString("0 ")
			}
		}
		writer.WriteByte('\n')
	} else {
		writer.WriteString("1 1 0 0 0\n")
	}
}
