// https://www.acmicpc.net/problem/6987
// bruteforce
// backtracking

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

// firstTeam vs seccondTeam
// firstTeamNum < seccondTeamNum
func isPossible(playCnt, firstTeamNum, nextTeamNum int) bool {
	if playCnt == 15 {
		return true
	}

	if nextTeamNum == 6 {
		return isPossible(playCnt, firstTeamNum+1, firstTeamNum+2)
	}

	if scores[firstTeamNum][0] > 0 && scores[nextTeamNum][2] > 0 {
		// team1 winning case
		scores[firstTeamNum][0]--
		scores[nextTeamNum][2]--

		if isPossible(playCnt+1, firstTeamNum, nextTeamNum+1) {
			return true
		}

		scores[firstTeamNum][0]++
		scores[nextTeamNum][2]++

	}
	if scores[firstTeamNum][1] > 0 && scores[nextTeamNum][1] > 0 {
		// team1 tie case
		scores[firstTeamNum][1]--
		scores[nextTeamNum][1]--

		if isPossible(playCnt+1, firstTeamNum, nextTeamNum+1) {
			return true
		}

		scores[firstTeamNum][1]++
		scores[nextTeamNum][1]++

	}
	if scores[firstTeamNum][2] > 0 && scores[nextTeamNum][0] > 0 {
		// team1 lose case
		scores[firstTeamNum][2]--
		scores[nextTeamNum][0]--

		if isPossible(playCnt+1, firstTeamNum, nextTeamNum+1) {
			return true
		}

		scores[firstTeamNum][2]++
		scores[nextTeamNum][0]++
	}

	return false
}

var scores [][3]int

func main() {
	defer writer.Flush()

	scores = make([][3]int, 6)

	for i := 0; i < 4; i++ {
		teamPlayCountFlag := true
		for j := 0; j < 6; j++ {
			win, tie, lose := readInt(), readInt(), readInt()
			scores[j] = [3]int{win, tie, lose}

			if teamPlayCountFlag {
				teamPlayCountFlag = win+tie+lose == 5
			}
		}

		if teamPlayCountFlag && isPossible(0, 0, 1) {
			writer.WriteString("1 ")
		} else {
			writer.WriteString("0 ")
		}
	}
	writer.WriteByte('\n')
}
