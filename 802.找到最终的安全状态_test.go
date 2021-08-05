package main

import (
	"reflect"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1",
			args: args{
				graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}},
			},
			want: []int{2, 4, 5, 6},
		},
		{
			name: "t2",
			args: args{
				graph: [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}},
			},
			want: []int{4},
		},
		{
			name: "t3",
			args: args{
				graph: [][]int{{0}, {2, 3, 4}, {3, 4}, {0, 4}, {}},
			},
			want: []int{4},
		},
		{
			name: "t4",
			args: args{
				graph: [][]int{{1, 2, 3, 4}, {2, 3, 4}, {0}, {0, 4}, {4}},
			},
			want: []int{},
		},
		//[1,2,3,4],[2,3,4],[0],[0,4],[4]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualSafeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
