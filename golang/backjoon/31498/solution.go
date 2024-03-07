// https://www.acmicpc.net/problem/31498
// same with cartooncup d
// binary search

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

var distanceT, speedT, distanceD, speedD, k int

func movingDistanceD(t int) int {
	return speedD * t
}

func movingDistnaceT(t int) int {
	return speedT*t - k*((t*(t-1))/2)
}

func main() {
	defer writer.Flush()

	distanceT, speedT = readInt(), readInt()
	distanceD, speedD = readInt(), readInt()
	k = readInt()

	distanceD += distanceT

	var maxStep int
	if k != 0 {
		maxStep = speedT / k
		lastMovingDistanceT := movingDistnaceT(maxStep)

		// check if could not reach to goal
		if lastMovingDistanceT < distanceT {
			writer.WriteString("-1\n")
			return
		}
	} else {
		maxStep = distanceT / speedT
		if distanceT%speedT != 0 {
			maxStep++
		}
	}

	// binary search
	low, high := 0, maxStep
	var result int
	for high > low {
		result = (low + high) / 2
		if movingDistnaceT(result) >= distanceT {
			high = result
		} else {
			result++
			low = result
		}
	}

	if movingDistanceD(result) >= distanceD {
		writer.WriteString("-1\n")
		return
	}

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
