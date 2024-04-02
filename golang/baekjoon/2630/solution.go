// https://www.acmicpc.net/problem/2630
// divide and conquer
// recursive

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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func whichColor(x, y, size int) (color, isSame bool) {
	firstColor := paper[x][y]
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			if paper[i][j] != firstColor {
				return false, false
			}
		}
	}
	return firstColor, true
}

func findColor(x, y, size int) (blue, white int) {
	if color, isSame := whichColor(x, y, size); isSame {
		if color {
			return 1, 0
		} else {
			return 0, 1
		}
	}

	halfSize := size / 2
	offsets := [][2]int{
		{0, 0},
		{0, halfSize},
		{halfSize, 0},
		{halfSize, halfSize},
	}

	for _, offset := range offsets {
		x, y := x+offset[0], y+offset[1]
		if color, isSame := whichColor(x, y, halfSize); isSame {
			if color {
				blue++
			} else {
				white++
			}
		} else {
			b, w := findColor(x, y, halfSize)
			blue += b
			white += w
		}
	}

	return
}

var paper [][]bool

// true : blue
// false : white
func main() {
	defer writer.Flush()

	n := readInt()
	paper = make([][]bool, n)
	for i := range paper {
		paper[i] = make([]bool, n)
		for j := range paper[i] {
			paper[i][j] = readInt() == 1
		}
	}

	blue, white := findColor(0, 0, n)
	writer.WriteString(strconv.Itoa(white))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(blue))
	writer.WriteByte('\n')
}
