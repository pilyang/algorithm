package outofboundarypaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	testCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "m: 2, n: 2, maxMove: 2, startRow: 0, startColumn: 0",
			args: args{m: 2, n: 2, maxMove: 2, startRow: 0, startColumn: 0},
			want: 6,
		},
		{
			name: "m: 1, n: 3, maxMove: 3, startRow: 0, startColumn: 1",
			args: args{m: 1, n: 3, maxMove: 3, startRow: 0, startColumn: 1},
			want: 12,
		},
		{
			name: "m: 36, n: 5, maxMove: 50, startRow: 15, startColumn: 3",
			args: args{m: 36, n: 5, maxMove: 50, startRow: 15, startColumn: 3},
			want: 390153306,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findPaths(tc.args.m, tc.args.n, tc.args.maxMove, tc.args.startRow, tc.args.startColumn)
			assert.Equal(t, tc.want, result)
		})
	}
}
