// https://www.acmicpc.net/problem/31461
// implementation

// 맞왜틀?? 술마시고 풀어서 그런가..?

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
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(graph [][]int, sc, sr, ec, er int) int {
	var lowColumn, lowRow, highColumn, highRow int
	if sc < ec {
		lowColumn, lowRow, highColumn, highRow = sc, sr, ec, er
	} else {
		lowColumn, lowRow, highColumn, highRow = ec, er, sc, sr
	}

	var lowMax, midMax, highMax int

	var lowSum int
	for c := lowColumn - 1; c >= 0; c-- {
		tempMax := maxInt(lowSum+graph[0][c], lowSum+graph[1][c])
		lowMax = maxInt(lowMax, tempMax)

		lowSum += graph[0][c] + graph[1][c]
		lowMax = maxInt(lowMax, lowSum)
	}

	var highSum int
	for c := highColumn + 1; c < len(graph[0]); c++ {
		tempMax := maxInt(highSum+graph[0][c], highSum+graph[1][c])
		highMax = maxInt(highMax, tempMax)

		highSum += graph[0][c] + graph[1][c]
		highMax = maxInt(highMax, highSum)
	}

	midMax = graph[lowRow][lowColumn] + graph[highRow][highColumn]
	if lowColumn == highColumn {
		if lowMax < 0 && highMax < 0 {
			return midMax
		}
		return midMax + maxInt(lowMax, highMax)
	}

	var sum int

	lowOposite := graph[(lowRow+1)%2][lowColumn]
	if lowOposite >= 0 {
		midMax += lowOposite
		if lowMax > 0 {
			sum += lowMax
		}
	} else if lowMax+lowOposite > 0 {
		midMax += lowOposite
		sum += lowMax
	}

	highOposite := graph[(highRow+1)%2][highColumn]
	if highOposite >= 0 {
		midMax += highOposite
		if highMax > 0 {
			sum += highMax
		}
	} else if highMax+highOposite > 0 {
		midMax += highOposite
		sum += highMax
	}

	for c := lowColumn + 1; c < highColumn; c++ {
		midMax += maxInt(maxInt(graph[0][c], graph[1][c]), graph[0][c]+graph[1][c])
	}

	return sum + midMax
}

func main() {
	defer writer.Flush()

	tc := readInt()

	for tc > 0 {
		tc--
		columns := readInt()
		graph := make([][]int, 2)
		for r := 0; r < 2; r++ {
			graph[r] = make([]int, columns)
			for c := 0; c < columns; c++ {
				graph[r][c] = readInt()
			}
		}
		sc, sr, ec, er := readInt()-1, readInt()-1, readInt()-1, readInt()-1

		writer.WriteString(strconv.Itoa(solve(graph, sc, sr, ec, er)))
		writer.WriteByte('\n')
	}
}
