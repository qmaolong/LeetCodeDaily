package main

import (
	"reflect"
	"testing"
)

func Test_smallestK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1",
			args: args{
				arr: []int{1, 3, 5, 7, 2, 4, 6, 8},
				k:   4,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestK(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestK() = %v, want %v", got, tt.want)
			}
		})
	}
}
