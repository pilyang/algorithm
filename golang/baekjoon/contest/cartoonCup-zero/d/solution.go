package main

import (
	"bufio"
	"math"
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

func quadraticFormula(a, b, c int) ([2]float64, bool) {
	// fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return [2]float64{}, false
	}
	x1 := (-float64(b) - float64(math.Sqrt(float64(discriminant)))) / (2 * float64(a))
	x2 := (-float64(b) + float64(math.Sqrt(float64(discriminant)))) / (2 * float64(a))
	return [2]float64{x1, x2}, true
}

func main() {
	defer writer.Flush()

	td, ts := readInt(), readInt()
	dd, ds := readInt(), readInt()
	k := readInt()

	dd += td

	// kn^2 + (-k-2ts)n + 2td = 0
	res, ok := quadraticFormula(k, -1*k-2*ts, 2*td)
	// fmt.Printf("res: %v, ok: %v\n", res, ok)
	if !ok {
		writer.WriteString("-1\n")
		return
	}

	step := int(math.Ceil(res[0]))
	if ds*step >= dd {
		writer.WriteString("-1\n")
		return
	}

	writer.WriteString(strconv.Itoa(step))
	writer.WriteByte('\n')
}
