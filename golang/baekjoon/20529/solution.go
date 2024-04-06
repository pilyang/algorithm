// https://www.acmicpc.net/problem/20529
// brute force, dfs (back tracking)

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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

func scanMbti() Mbti {
	var mbti Mbti
	for i, c := range scanString() {
		mbti[i] = c == mbtiTag[i]
	}
	return mbti
}

type Mbti [4]bool

var mbtiTag = [4]rune{'E', 'S', 'T', 'J'}

func calcDistance(a, b, c Mbti) int {
	distance := 0
	for i := 0; i < 4; i++ {
		if a[i] != b[i] {
			distance++
		}
		if a[i] != c[i] {
			distance++
		}
		if b[i] != c[i] {
			distance++
		}
	}
	return distance
}

func main() {
	defer writer.Flush()

	tc := scanInt()
	for tc > 0 {
		tc--

		mbtis := make([]Mbti, scanInt())
		for i := range mbtis {
			mbtis[i] = scanMbti()
		}

		result := solve(mbtis)
		writer.WriteString(strconv.Itoa(result))
		writer.WriteByte('\n')
	}
}

func solve(mbtis []Mbti) int {
	minDistance := 12

	stack := make([]Mbti, 0, 3)

	var dfs func(int)
	dfs = func(index int) {
		if minDistance == 0 {
			return
		}

		if len(stack) == 3 {
			distance := calcDistance(stack[0], stack[1], stack[2])
			if distance < minDistance {
				minDistance = distance
			}
			return
		}

		for i := index; i < len(mbtis); i++ {
			stack = append(stack, mbtis[i])
			dfs(i + 1)
			stack = stack[:len(stack)-1]
		}
	}
	dfs(0)

	return minDistance
}
