// https://www.acmicpc.net/problem/14500

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

// [2]int{x, y}
type Tetromino struct {
	Base [][2]int
	Sub  [][][2]int
}

var tetrominos = []*Tetromino{
	{
		Base: [][2]int{{0, 0}, {1, 0}, {2, 0}},
		Sub:  [][][2]int{{{0, 1}}, {{0, -1}}, {{1, 1}}, {{1, -1}}, {{2, 1}}, {{2, -1}}, {{3, 0}}},
	},
	{
		Base: [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Sub:  [][][2]int{{{-1, 0}}, {{1, 0}}, {{-1, 1}}, {{1, 1}}, {{-1, 2}}, {{1, 2}}, {{0, 3}}},
	},
	{
		Base: [][2]int{{0, 0}, {1, 0}},
		Sub:  [][][2]int{{{0, 1}, {1, 1}}, {{1, 1}, {2, 1}}, {{-1, 1}, {0, 1}}},
	},
	{
		Base: [][2]int{{0, 0}, {0, 1}},
		Sub:  [][][2]int{{{1, -1}, {1, 0}}, {{1, 1}, {1, 2}}},
	},
}

type Paper struct {
	m      *[][]int
	rs, cs int
}

func (p *Paper) sumValues(x, y int, offSets *[][2]int) (int, bool) {
	sum := 0
	for _, offSet := range *offSets {
		nx, ny := x+offSet[0], y+offSet[1]
		if nx < 0 || nx >= p.cs || ny < 0 || ny >= p.rs {
			return 0, false
		}
		sum += (*p.m)[ny][nx]
	}
	return sum, true
}

func main() {
	defer writer.Flush()

	r, c := readInt(), readInt()
	m := make([][]int, r)
	for y := 0; y < r; y++ {
		m[y] = make([]int, c)
		for x := 0; x < c; x++ {
			m[y][x] = readInt()
		}
	}

	paper := Paper{m: &m, rs: r, cs: c}
	maxValue := 0
	for y := 0; y < r; y++ {
		for x := 0; x < c; x++ {
			for _, tet := range tetrominos {
				if baseSize, isPossible := paper.sumValues(x, y, &tet.Base); isPossible {
					for _, sub := range tet.Sub {
						if subSize, isSubPossible := paper.sumValues(x, y, &sub); isSubPossible {
							newSize := baseSize + subSize
							if maxValue < newSize {
								maxValue = newSize
							}
						}
					}
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(maxValue) + "\n")
}
