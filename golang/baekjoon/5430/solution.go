// https://www.acmicpc.net/problem/5430

package main

import (
	"bufio"
	"errors"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func readString() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt() int {
	res, _ := strconv.Atoi(readString())
	return res
}

func readCommands() []rune {
	return []rune(readString())
}

func readNums() []string {
	input := readString()
	sl := strings.Split(input[1:len(input)-1], ",")
	if len(sl) == 1 && sl[0] == "" {
		return []string{}
	}
	return sl
}

func writeResult(res *[]string) {
	writer.WriteString("[" + strings.Join(*res, ",") + "]\n")
}

func runCommands(commands *[]rune, nums *[]string) error {
	// true : from first, false : from last
	flag := true

	for _, command := range *commands {
		switch command {
		case 'R':
			flag = !flag
		case 'D':
			if len(*nums) == 0 {
				return errors.New("error")
			}
			if flag {
				*nums = (*nums)[1:]
			} else {
				*nums = (*nums)[:len(*nums)-1]
			}
		}
	}
	if !flag {
		slices.Reverse(*nums)
	}

	return nil
}

func main() {
	defer writer.Flush()

	count := readInt()
	for i := 0; i < count; i++ {
		commands := readCommands()
		readInt()
		nums := readNums()

		if err := runCommands(&commands, &nums); err != nil {
			writer.WriteString("error\n")
		} else {
			writeResult(&nums)
		}
	}
}
