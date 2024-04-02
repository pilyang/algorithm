// https://www.acmicpc.net/problem/2294

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRequiredMinCoun(t *testing.T) {
	type args struct {
		coins  *[]int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				coins:  &[]int{1, 5, 12},
				target: 15,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				coins:  &[]int{3, 5, 6},
				target: 7,
			},
			want: -1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := findRequiredMinCoin(tc.args.coins, tc.args.target)
			assert.Equal(t, tc.want, actual)
		})
	}
}
