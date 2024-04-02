// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/description/?envType=daily-question&envId=2024-01-28
package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSubmatrixSumTarget(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0",
			args: args{
				matrix: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
				target: 0,
			},
			want: 4,
		},
		{
			name: "matrix = [[1,-1],[-1,1]], target = 0",
			args: args{
				matrix: [][]int{{1, -1}, {-1, 1}},
				target: 0,
			},
			want: 5,
		},
		{
			name: "matrix = [[904]], target = 0",
			args: args{
				matrix: [][]int{{904}},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := numSubmatrixSumTarget(tc.args.matrix, tc.args.target)
			assert.Equal(t, tc.want, result)
		})
	}
}
