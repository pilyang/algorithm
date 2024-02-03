// https://www.acmicpc.net/problem/1904

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPossibleNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{num: 1},
			want: 1,
		},
		// {
		// 	name: "test 2",
		// 	args: args{num: 2},
		// 	want: 2,
		// },
		// {
		// 	name: "test 2",
		// 	args: args{num: 4},
		// 	want: 5,
		// },
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := countPossibleNum(tc.args.num)
			assert.Equal(t, tc.want, actual)
		})
	}
}
