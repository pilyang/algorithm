package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		lessons *[]int
		cdCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				lessons: &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				cdCount: 3,
			},
			want: 17,
		},
		{
			name: "test case 2",
			args: args{
				lessons: &[]int{5, 9, 6, 8, 7, 7, 5},
				cdCount: 7,
			},
			want: 9,
		},
		{
			name: "test case 3",
			args: args{
				lessons: &[]int{99, 1, 99, 1},
				cdCount: 3,
			},
			want: 100,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := binarySearch(tc.args.lessons, tc.args.cdCount)
			assert.Equal(t, tc.want, actual)
		})
	}
}
