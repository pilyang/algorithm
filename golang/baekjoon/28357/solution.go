// binary search
package main

import (
	"bufio"
	"os"
	"slices"
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

func bruteforceSolve(scores []int, k int) int {
	slices.Sort(scores)

	candyCount := 0
	currentIdx := len(scores) - 1
	for currentIdx > 0 && scores[currentIdx] == scores[currentIdx-1] {
		currentIdx--
	}
	x := scores[currentIdx]

	// first move step by step
	for {
		// find NextIdx
		nextIdx := currentIdx - 1
		for nextIdx >= 0 && scores[nextIdx] == scores[currentIdx] {
			nextIdx--
		}
		if nextIdx < 0 {
			break
		}

		diff := scores[currentIdx] - scores[nextIdx]
		candyToAdd := (len(scores) - nextIdx - 1) * diff

		if candyToAdd+candyCount > k {
			break
		}

		candyCount += candyToAdd
		currentIdx = nextIdx
		x = scores[currentIdx]
	}

	// Calculate the maximum possible diff directly
	numberOfStudents := len(scores) - currentIdx
	maxDiff := (k - candyCount) / numberOfStudents

	// Decrease x by maxDiff
	x -= maxDiff
	if x < 0 {
		x = 0
	}
	return x
}

func isPossible(scores []int, k, x int) bool {
	count := 0
	for _, score := range scores {
		diff := score - x
		if diff > 0 {
			count += diff
		}
	}
	return count <= k
}

func binarySearchSolve(scores []int, k int, maxScore int) int {
	low, high := 0, maxScore
	for high > low {
		mid := (high + low) / 2
		if isPossible(scores, k, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return high
}

func main() {
	defer writer.Flush()

	n, k := scanInt(), scanInt()
	scores := make([]int, n)
	maxScore := 0
	for i := range scores {
		score := scanInt()
		scores[i] = score

		if maxScore < score {
			maxScore = score
		}

	}

	// x := bruteforceSolve(scores, k)
	x := binarySearchSolve(scores, k, maxScore)

	writer.WriteString(strconv.Itoa(x))
	writer.WriteByte('\n')
}
