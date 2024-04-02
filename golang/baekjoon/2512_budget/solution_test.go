package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaximumBudget(t *testing.T) {
	type args struct {
		budgets   *[]int
		maxBudget int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				budgets:   &[]int{120, 110, 140, 150},
				maxBudget: 485,
			},
			want: 127,
		},
		{
			name: "test2",
			args: args{
				budgets:   &[]int{70, 80, 30, 40, 100},
				maxBudget: 450,
			},
			want: 100,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := findMaximumBudget(tc.args.budgets, tc.args.maxBudget)
			assert.Equal(t, tc.want, actual)
		})
	}
}
