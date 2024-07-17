// combination, dp

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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

var combDp [50_001][4]int64

func comb(n, r int) int64 {
	if n < r {
		return 0
	}
	if r == 1 {
		return int64(n)
	}

	if combDp[n][r] == 0 {
		combDp[n][r] = comb(n-1, r-1) + comb(n-1, r)
	}
	return combDp[n][r]
}

func main() {
	defer writer.Flush()

	teamName := scanString()
	teamRunes := make(map[rune]int, 3)
	for _, r := range teamName {
		teamRunes[r]++
	}

	playerNum := scanInt()
	playerFirst := make(map[rune]int, playerNum)
	for i := 0; i < playerNum; i++ {
		name := scanString()
		playerFirst[rune(name[0])]++
	}

	var count int64
	count = 1
	for r, num := range teamRunes {
		count *= comb(playerFirst[r], num)
	}
	writer.WriteString(strconv.FormatInt(count, 10))
	writer.WriteByte('\n')
}
