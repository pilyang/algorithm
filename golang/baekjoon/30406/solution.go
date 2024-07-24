// greedy

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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

var gifts = make(map[int]int, 3)

var (
	score3Pair = [][2]int{{0, 3}, {1, 2}}
	score2Pair = [][2]int{{0, 2}, {1, 3}}
	score1Pair = [][2]int{{0, 1}, {2, 3}}
	// score0Pair = [][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}

	// scores = []int{3, 2, 1, 0}
	scores = []int{3, 2, 1}
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func processScorePair(pairs [][2]int, unitScore int) int {
	score := 0
	for _, pair := range pairs {
		minCount := minInt(gifts[pair[0]], gifts[pair[1]])
		score += minCount * unitScore

		gifts[pair[0]] -= minCount
		gifts[pair[1]] -= minCount
	}

	return score
}

func main() {
	defer writer.Flush()

	n := scanInt()
	for i := 0; i < n; i++ {
		gifts[scanInt()]++
	}

	// pairss := [][][2]int{score3Pair, score2Pair, score1Pair, score0Pair}
	pairss := [][][2]int{score3Pair, score2Pair, score1Pair}
	score := 0
	for i, pair := range pairss {
		score += processScorePair(pair, scores[i])
	}

	writer.WriteString(strconv.Itoa(score))
	writer.WriteByte('\n')
}
