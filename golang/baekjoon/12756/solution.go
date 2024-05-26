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

func checkDeathTurn(attack, health int) int {
	turn := health / attack
	if health%attack != 0 {
		turn++
	}
	return turn
}

func main() {
	defer writer.Flush()

	aAttack, aHealth := scanInt(), scanInt()
	bAttack, bHealth := scanInt(), scanInt()

	aDeathTurn := checkDeathTurn(bAttack, aHealth)
	bDeathTurn := checkDeathTurn(aAttack, bHealth)

	if aDeathTurn > bDeathTurn {
		writer.WriteString("PLAYER A\n")
	} else if aDeathTurn < bDeathTurn {
		writer.WriteString("PLAYER B\n")
	} else {
		writer.WriteString("DRAW\n")
	}
}
