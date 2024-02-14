// https://www.acmicpc.net/problem/12865
// knapsack problem
// https://www.youtube.com/watch?v=rhda6lR5kyQ

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxSum(t *testing.T) {
	type args struct {
		items     *[]*Item
		maxWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				items: &[]*Item{
					{weight: 6, score: 13},
					{weight: 4, score: 8},
					{weight: 3, score: 6},
					{weight: 5, score: 12},
				},
				maxWeight: 7,
			},
			want: 14,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := findMaxSum(tc.args.items, tc.args.maxWeight)
			assert.Equal(t, tc.want, actual)
		})
	}
}

func Test_findMaxSum2(t *testing.T) {
	type args struct {
		items     *[]*Item
		maxWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				items: &[]*Item{
					{weight: 6, score: 13},
					{weight: 4, score: 8},
					{weight: 3, score: 6},
					{weight: 5, score: 12},
				},
				maxWeight: 7,
			},
			want: 14,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := findMaxSum(tc.args.items, tc.args.maxWeight)
			assert.Equal(t, tc.want, actual)
		})
	}
}
