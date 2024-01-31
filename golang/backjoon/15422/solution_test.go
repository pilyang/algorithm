// https://www.acmicpc.net/problem/15422

package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMinCost(t *testing.T) {
	type args struct {
		graph       *[][]*Edge
		freeTickets *[][2]int
		startCity   int
		endCity     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				graph: makeGraphFromInput(8,
					[]string{
						"0 1 10",
						"0 2 10",
						"1 2 10",
						"2 6 40",
						"6 7 10",
						"5 6 10",
						"3 5 15",
						"3 6 40",
						"3 4 20",
						"1 4 20",
						"1 3 20",
					}),
				freeTickets: makeFreeTickets([]string{"4 7"}),
				startCity:   0,
				endCity:     5,
			},
			want: 45,
		},
		{
			name: "test2",
			args: args{
				graph: makeGraphFromInput(8,
					[]string{
						"0 1 10",
						"0 2 10",
						"1 2 10",
						"2 6 40",
						"6 7 10",
						"5 6 10",
						"3 5 15",
						"3 6 40",
						"3 4 20",
						"1 4 20",
						"1 3 30",
					}),
				freeTickets: makeFreeTickets([]string{"4 7"}),
				startCity:   0,
				endCity:     5,
			},
			want: 50,
		},
		{
			name: "test3",
			args: args{
				graph:       makeGraphFromInput(2, []string{"0 1 10"}),
				freeTickets: makeFreeTickets([]string{}),
				startCity:   0,
				endCity:     1,
			},
			want: 10,
		},
		{
			name: "test4",
			args: args{
				graph:       makeGraphFromInput(2, []string{"0 1 10"}),
				freeTickets: makeFreeTickets([]string{"0 1"}),
				startCity:   1,
				endCity:     0,
			},
			want: 10,
		},
		{
			name: "test5",
			args: args{
				graph: makeGraphFromInput(6,
					[]string{
						"0 1 30",
						"1 4 10",
						"0 2 10",
						"2 3 10",
						"3 4 10",
						"4 5 50",
					}),
				freeTickets: makeFreeTickets([]string{"0 1", "4 5"}),
				startCity:   0,
				endCity:     5,
			},
			want: 30,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := findMinCost(tc.args.graph, tc.args.freeTickets, tc.args.startCity, tc.args.endCity)
			assert.Equal(t, tc.want, actual)
		})
	}
}

func makeGraphFromInput(numOfCity int, flights []string) *[][]*Edge {
	graph := make([][]*Edge, numOfCity)
	for i := range graph {
		graph[i] = make([]*Edge, 0)
	}
	for _, flight := range flights {
		info := strings.Split(flight, " ")
		a, _ := strconv.Atoi(info[0])
		b, _ := strconv.Atoi(info[1])
		cost, _ := strconv.ParseInt(info[2], 10, 64)
		graph[a] = append(graph[a], &Edge{b, cost})
		graph[b] = append(graph[b], &Edge{a, cost})
	}
	return &graph
}

func makeFreeTickets(freeFlights []string) *[][2]int {
	freeTickets := make([][2]int, 0)
	for _, flight := range freeFlights {
		info := strings.Split(flight, " ")
		from, _ := strconv.Atoi(info[0])
		to, _ := strconv.Atoi(info[1])
		freeTickets = append(freeTickets, [2]int{from, to})
	}
	return &freeTickets
}
