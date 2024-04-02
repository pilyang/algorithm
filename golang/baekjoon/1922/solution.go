// https://www.acmicpc.net/problem/1922

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// iniate reader and writer
	// reader := *bufio.NewReader(os.Stdin)
	scanner := *bufio.NewScanner(os.Stdin)
	writer := *bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// read data
	var vertexCount, edgeCount int
	scanner.Scan()
	vertexCount, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	edgeCount, _ = strconv.Atoi(scanner.Text())
	// fmt.Fscanf(&reader, "%d\n%d\n", &vertexCount, &edgeCount)
	networks := make([]*Edge, edgeCount)
	for i := range networks {
		var from, to, cost int
		// fmt.Fscanf(&reader, "%d %d %d\n", &from, &to, &cost)
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		from, _ = strconv.Atoi(input[0])
		to, _ = strconv.Atoi(input[1])
		cost, _ = strconv.Atoi(input[2])
		networks[i] = &Edge{from, to, cost}
	}

	if result, err := findMinNetworkCost(&networks, vertexCount); err != nil {
		fmt.Fprintln(&writer, "-1")
	} else {
		fmt.Fprintln(&writer, result)
	}
}

type Edge struct {
	From int
	To   int
	Cost int
}

func findMinNetworkCost(networks *[]*Edge, vertexCount int) (int, error) {
	sort.Slice(*networks, func(i, j int) bool {
		return (*networks)[i].Cost < (*networks)[j].Cost
	})
	union := newUnion(vertexCount)

	var sum, count int
	for _, network := range *networks {
		a, b := network.From-1, network.To-1
		if !union.isUnion(a, b) {
			union.merge(a, b)
			sum += network.Cost
			count++
		}
		if count == vertexCount-1 {
			return sum, nil
		}
	}

	return -1, errors.New("impossible to connect all networks")
}

type Union []int

func newUnion(n int) *Union {
	u := Union(make([]int, n))
	for i := range u {
		u[i] = -1
	}
	return &u
}

func (u *Union) find(a int) int {
	if (*u)[a] < 0 {
		return a
	}
	(*u)[a] = u.find((*u)[a])
	return (*u)[a]
}

func (u *Union) isUnion(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *Union) merge(a, b int) {
	aRoot, bRoot := u.find(a), u.find(b)
	(*u)[aRoot] = bRoot
}
