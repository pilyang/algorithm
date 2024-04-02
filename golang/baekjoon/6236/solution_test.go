package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMinMoney(t *testing.T) {
	type args struct {
		moneyPlan []int
		m         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				moneyPlan: []int{100, 400, 300, 100, 500, 101, 400},
				m:         5,
			},
			want: 500,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := findMinMoney(tc.args.moneyPlan, tc.args.m)
			assert.Equal(t, tc.want, actual)
		})
	}
}
