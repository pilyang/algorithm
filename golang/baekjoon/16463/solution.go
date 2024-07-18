// implementation, simulation, brute-force

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

var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	defer writer.Flush()

	targetYear := scanInt()

	// date from 2019.01.01
	// 0 for 2019.01.01
	currentDate := 0
	count := 0

	for year := 2019; year <= targetYear; year++ {
		for m, daysInMonth := range days {
			if m == 1 && isLeapYear(year) {
				daysInMonth++
			}
			// check if this month has 13 friday
			if (currentDate+12)%7 == 3 {
				count++
			}

			// move to next month
			currentDate += daysInMonth
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
