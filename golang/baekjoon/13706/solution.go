package main

import (
	"bufio"
	"math/big"
	"os"
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

var (
	one = big.NewInt(1)
	two = big.NewInt(2)
)

func main() {
	defer writer.Flush()

	num, _ := new(big.Int).SetString(scanString(), 10)

	low, high := big.NewInt(0), new(big.Int).Div(num, two)
	mid := new(big.Int)

	for low.Cmp(high) < 0 {
		mid.Div(new(big.Int).Add(low, high), two)
		exp := new(big.Int).Exp(mid, two, nil)
		if num.Cmp(exp) == 0 {
			break
		} else if num.Cmp(exp) < 0 {
			high.Set(mid)
		} else {
			low.Add(mid, one)
		}
	}

	writer.WriteString(mid.String())
}
