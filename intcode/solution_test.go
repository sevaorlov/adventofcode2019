package intcode

import (
	"reflect"
	"testing"
)

func extendArray(a []int64, n int) []int64 {
	new := make([]int64, len(a)*100)
	return append(a, new...)
}

func TestSolve(t *testing.T) {
	type args struct {
		a                 []int64
		inputInstructions []int64
		index             int
		stopOnOutput      bool
	}
	tests := []struct {
		name      string
		args      args
		want      []int64
		wantIndex int
		wantErr   bool
	}{
		{
			name: "",
			args: args{
				a:                 extendArray([]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, 10),
				inputInstructions: []int64{},
			},
			want:      []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			wantIndex: -1,
			wantErr:   false,
		},
		{
			name: "",
			args: args{
				a:                 []int64{104, 1125899906842624, 99},
				inputInstructions: []int64{0},
			},
			want:      []int64{1125899906842624},
			wantIndex: -1,
			wantErr:   false,
		},
		{
			name: "",
			args: args{
				a:                 []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
				inputInstructions: []int64{0},
			},
			want:      []int64{1219070632396864},
			wantIndex: -1,
			wantErr:   false,
		},
		{
			name: "",
			args: args{
				a:                 []int64{
					3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,
					104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
				},
				inputInstructions: []int64{0},
			},
			want:      []int64{999},
			wantIndex: -1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Solve(tt.args.a, tt.args.inputInstructions, tt.args.index, tt.args.stopOnOutput)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantIndex {
				t.Errorf("Solve() got1 = %v, want %v", got1, tt.wantIndex)
			}
		})
	}
}
