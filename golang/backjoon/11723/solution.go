// https://www.acmicpc.net/problem/11723
// set

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

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func main() {
	defer writer.Flush()

	n := readInt()

	set := make([]bool, 21)
	for i := 0; i < n; i++ {
		switch cmd := readString(); cmd {
		case "add":
			set[readInt()] = true
		case "remove":
			set[readInt()] = false
		case "check":
			if set[readInt()] {
				writer.WriteString("1\n")
			} else {
				writer.WriteString("0\n")
			}
		case "toggle":
			num := readInt()
			set[num] = !set[num]
		case "all":
			for i := 1; i <= 20; i++ {
				set[i] = true
			}
		case "empty":
			for i := 1; i <= 20; i++ {
				set[i] = false
			}
		}
	}
}
