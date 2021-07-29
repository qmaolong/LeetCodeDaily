package main

import (
	"reflect"
	"testing"
)

func Test_pathInZigZagTree(t *testing.T) {
	type args struct {
		label int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1",
			args: args{
				label: 14,
			},
			want: []int{1, 3, 4, 14},
		},
		{
			name: "t2",
			args: args{
				label: 26,
			},
			want: []int{1, 2, 6, 10, 26},
		},
		{
			name: "t3",
			args: args{
				label: 1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathInZigZagTree(tt.args.label); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathInZigZagTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
