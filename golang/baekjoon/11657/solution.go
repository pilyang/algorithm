// https://www.acmicpc.net/problem/11657
// spfa(shortest path fastest algorithm)
// improved Bellman-Ford

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readStrings(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input := readStrings(scanner)
	v, _ := strconv.Atoi(input[0])
	e, _ := strconv.Atoi(input[1])
	graph := make([][]*Edge, v)

	for i := 0; i < e; i++ {
		eInput := readStrings(scanner)
		from, _ := strconv.Atoi(eInput[0])
		from--
		to, _ := strconv.Atoi(eInput[1])
		to--
		cost, _ := strconv.Atoi(eInput[2])
		graph[from] = append(graph[from], &Edge{to: to, cost: int64(cost)})
	}

	dists, err := findMinDists(&graph)
	if err != nil {
		fmt.Fprintln(writer, "-1")
		return
	}
	for i := 1; i < v; i++ {
		fmt.Fprintf(writer, "%d\n", dists[i])
	}
}

type Edge struct {
	to   int
	cost int64
}

func findMinDists(graph *[][]*Edge) ([]int64, error) {
	vertexCount := len(*graph)

	dists := make([]int64, vertexCount)
	for i := range dists {
		dists[i] = math.MaxInt64
	}
	dists[0] = 0
	isInQue := make([]bool, vertexCount)
	queue := make([]int, 0, vertexCount)
	cnt := make([]int, vertexCount)

	queue = append(queue, 0)
	isInQue[0] = true
	for len(queue) != 0 {
		// pop
		current := queue[0]
		queue = queue[1:]

		// mark as visit
		isInQue[current] = false
		cnt[current]++
		if cnt[current] >= vertexCount {
			return nil, errors.New("impossible")
		}

		for _, next := range (*graph)[current] {
			newDist := dists[current] + next.cost
			if dists[next.to] > newDist {
				dists[next.to] = newDist
				if !isInQue[next.to] {
					queue = append(queue, next.to)
					isInQue[next.to] = true
				}
			}
		}

	}

	for i := 1; i < vertexCount; i++ {
		if dists[i] == math.MaxInt64 {
			dists[i] = -1
		}
	}
	return dists, nil
}
