// https://www.acmicpc.net/problem/11266
// cut vertex, graph

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCutVertex(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want *[]bool
	}{
		// {
		// 	name: "case 1",
		// 	args: args{
		// 		graph: [][]int{
		// 			{3, 4, 5},
		// 			{6},
		// 			{6},
		// 			{0, 4},
		// 			{0, 3},
		// 			{0, 6},
		// 			{1, 2, 5},
		// 		},
		// 	},
		// 	want: &[]bool{true, false, false, false, false, true, true},
		// },
		// {
		// 	name: "case 2",
		// 	args: args{
		// 		graph: [][]int{
		// 			{1, 4},
		// 			{0, 2, 5},
		// 			{1, 3},
		// 			{2, 4},
		// 			{0, 3, 5},
		// 			{1, 4},
		// 		},
		// 	},
		// 	want: &[]bool{false, true, false, false, false, false},
		// },
		{
			name: "case 3",
			args: args{
				graph: [][]int{
					{1},
					{0, 2},
					{1},
				},
			},
			want: &[]bool{false, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findCutVertex(tt.args.graph)
			assert.Equal(t, *tt.want, *actual)
		})
	}
}
