// https://www.acmicpc.net/problem/11657
// spfa(shortest path fastest algorithm)
// improved Bellman-Ford

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMinDists(t *testing.T) {
	type args struct {
		graph *[][]*Edge
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				graph: &[][]*Edge{
					{
						// from 0
						{to: 1, cost: 4},
						{to: 2, cost: 3},
					},
					{
						// from 1
						{to: 2, cost: -1},
					},
					{
						// from 2
						{to: 0, cost: -2},
					},
				},
			},
			want:    []int64{0, 4, 3},
			wantErr: false,
		},
		{
			name: "test 2",
			args: args{
				graph: &[][]*Edge{
					{
						// from 0
						{to: 1, cost: 4},
						{to: 2, cost: 3},
					},
					{
						// from 1
						{to: 2, cost: -4},
					},
					{
						// from 2
						{to: 0, cost: -2},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test 3",
			args: args{
				graph: &[][]*Edge{
					{
						// from 0
						{to: 1, cost: 4},
						{to: 1, cost: 3},
					},
					{
						// from 1
					},
					{
						// from 2
					},
				},
			},
			want:    []int64{0, 3, -1},
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := findMinDists(tc.args.graph)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
