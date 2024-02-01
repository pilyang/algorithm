// https://www.acmicpc.net/problem/1922

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMinNetworkCost(t *testing.T) {
	type args struct {
		networks    *[]*Edge
		vertexCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				networks: &[]*Edge{
					{1, 2, 5},
					{1, 3, 4},
					{2, 3, 2},
					{2, 4, 7},
					{3, 4, 6},
					{3, 5, 11},
					{4, 5, 3},
					{4, 6, 8},
					{5, 6, 8},
				},
				vertexCount: 6,
			},
			want: 23,
		},
		{
			name: "test 2",
			args: args{
				networks: &[]*Edge{
					{1, 2, 3},
					{2, 3, 3},
					{3, 4, 3},
					{4, 5, 3},
					{2, 3, 2},
					{3, 4, 2},
				},
				vertexCount: 5,
			},
			want: 10,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := findMinNetworkCost(tc.args.networks, tc.args.vertexCount)
			assert.Equal(t, tc.want, actual)
		})
	}
}
