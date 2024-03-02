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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func readRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

var redDxy = [][2]int{
	{0, 1},
	{1, 1},
}

var blueDxy = [][2]int{
	{1, 0},
	{1, 1},
}

func checkColor(x, y int, dxys [][2]int, color rune) bool {
	isChecked[y][x] = true

	for _, dxy := range dxys {
		nx, ny := x+dxy[0], y+dxy[1]
		if ny < 0 || ny >= len(chocolate) || nx < 0 || nx >= len(chocolate[ny]) {
			return false
		}
		if chocolate[ny][nx] != color || isChecked[ny][nx] {
			return false
		}
		isChecked[ny][nx] = true
	}

	return true
}

func isPossible() bool {
	for y := range chocolate {
		for x := range chocolate[y] {
			if isChecked[y][x] {
				continue
			}
			switch chocolate[y][x] {
			case 'R':
				if !checkColor(x, y, redDxy, 'R') {
					return false
				}
			case 'B':
				if !checkColor(x, y, blueDxy, 'B') {
					return false
				}
			}
		}
	}
	return true
}

var (
	chocolate [][]rune
	isChecked [][]bool
)

func main() {
	defer writer.Flush()

	n := readInt()
	chocolate = make([][]rune, n)
	isChecked = make([][]bool, n)
	for i := 0; i < n; i++ {
		chocolate[i] = readRunes()
		isChecked[i] = make([]bool, i+1)
	}

	if isPossible() {
		writer.WriteString("1\n")
	} else {
		writer.WriteString("0\n")
	}
}
