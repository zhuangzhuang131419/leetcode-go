package leetcode

import "testing"

func Test_escapeGhosts(t *testing.T) {
	type args struct {
		ghosts [][]int
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: struct {
				ghosts [][]int
				target []int
			}{
				ghosts: [][]int{
					{1, 0},
					{0, 3},
				},
				target:[]int{
					0, 1,
				},
			},
			want: true,
		},
		{
			args: struct {
				ghosts [][]int
				target []int
			}{
				ghosts: [][]int{
					{1, 0},
				},
				target:[]int{
					2, 0,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeGhosts(tt.args.ghosts, tt.args.target); got != tt.want {
				t.Errorf("escapeGhosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
