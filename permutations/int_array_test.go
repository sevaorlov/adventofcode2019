package permutations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateInt(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				a: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateInt(tt.args.a)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
