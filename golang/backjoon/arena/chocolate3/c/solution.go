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

func countPossibleKnight(xs, ys, dx, dy int) int {
	var count int
	// init
	grid := make([][]bool, ys)
	for i := range grid {
		grid[i] = make([]bool, xs)
	}

	for x := 0; x < xs; x++ {
		for y := 0; y < ys; y++ {
			if !grid[y][x] {
				grid[y][x] = true
				nx, ny := x+dx, y+dy
				if nx >= 0 && nx < xs && ny >= 0 && ny < ys {
					grid[ny][nx] = true
				}
				count++
				continue
			}
		}
	}

	return count
}

func main() {
	defer writer.Flush()

	n := readInt()
	for n > 0 {
		n--
		result := countPossibleKnight(readInt(), readInt(), readInt(), readInt())
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
	}
}
