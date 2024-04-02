// https://www.acmicpc.net/problem/2469
// implementation

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func applyLadder(entry *[]rune, ladder *[]rune) {
	for i, l := range *ladder {
		if l == '-' {
			(*entry)[i], (*entry)[i+1] = (*entry)[i+1], (*entry)[i]
		}
	}
}

func findEmtpyLadder(start, end *[]rune) (*[]rune, bool) {
	result := make([]rune, len(*start)-1)

	for i := 0; i < len(result); i++ {
		if (*start)[i] == (*end)[i] {
			result[i] = '*'
		} else {
			if (*start)[i] == (*end)[i+1] && (*start)[i+1] == (*end)[i] {
				result[i] = '-'
				if i+1 < len(result) {
					result[i+1] = '*'
				}
				i++
			} else {
				return nil, false
			}
		}
	}

	return &result, true
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	defer writer.Flush()

	k, n := readInt(), readInt()

	start := []rune(alphabet[:k])
	end := readRunes()

	var cnt int
	for {
		preLadder := readRunes()
		cnt++
		if preLadder[0] == '?' {
			break
		}
		applyLadder(&start, &preLadder)
	}

	postLadder := make([][]rune, n-cnt)
	for i := range postLadder {
		postLadder[i] = readRunes()
	}
	for i := len(postLadder) - 1; i >= 0; i-- {
		applyLadder(&end, &postLadder[i])
	}

	if result, isPossible := findEmtpyLadder(&start, &end); isPossible {
		for _, r := range *result {
			writer.WriteRune(r)
		}
	} else {
		writer.WriteString(strings.Repeat("x", k-1))
	}
	writer.WriteByte('\n')
}
