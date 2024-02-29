// https://www.acmicpc.net/problem/28255
// implementation

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

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func rev(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isTripleLayer(s string) bool {
	x := (len(s) + 2) / 3
	sdot := s[0:x]
	rsdot := rev(sdot)

	switch m := len(s) % 3; m {
	// s = s' + rev(s') + s'
	case 0:
		return s[x:x*2] == rsdot && s[x*2:] == sdot
	// s = s' + tail(rev(s')) + tail(s')
	case 1:
		return s[x:x*2-1] == rsdot[1:] && s[x*2-1:] == sdot[1:]
	// s = s' + rev(s') + tail(s')
	// s = s' + tail(rev(s')) + s'
	case 2:
		return (s[x:x*2] == rsdot && s[x*2:] == sdot[1:]) ||
			(s[x:x*2-1] == rsdot[1:] && s[x*2-1:] == sdot)
	}
	return false
}

func main() {
	defer writer.Flush()

	n := readInt()

	for n > 0 {
		n--
		if isTripleLayer(readString()) {
			writer.WriteString("1\n")
		} else {
			writer.WriteString("0\n")
		}
	}
}
