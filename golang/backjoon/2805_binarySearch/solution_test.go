package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaximumHegith(t *testing.T) {
	type args struct {
		trees    *[]int
		required int
	}
	testCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				trees:    &[]int{20, 15, 10, 17},
				required: 7,
			},
			want: 15,
		},
		{
			name: "test case 2",
			args: args{
				trees:    &[]int{4, 42, 40, 26, 46},
				required: 20,
			},
			want: 36,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			answer := findMaximumHeight(tc.args.trees, tc.args.required)
			assert.Equal(t, tc.want, answer)
		})
	}
}
