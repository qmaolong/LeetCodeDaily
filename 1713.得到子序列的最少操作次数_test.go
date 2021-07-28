package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		target []int
		arr    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				target: []int{5, 1, 3},
				arr:    []int{9, 4, 2, 3, 4},
			},
			want: 2,
		},
		{
			name: "t2",
			args: args{
				target: []int{6, 4, 8, 1, 3, 2},
				arr:    []int{4, 7, 6, 2, 3, 8, 6, 1},
			},
			want: 3,
		},
		{
			name: "t3",
			args: args{
				target: []int{16, 7, 20, 11, 15, 13, 10, 14, 6, 8},
				arr:    []int{11, 14, 15, 7, 5, 5, 6, 10, 11, 6},
			},
			want: 6,
		},
		{
			name: "t4",
			args: args{
				target: []int{1, 3, 8},
				arr:    []int{2, 6},
			},
			want: 3,
		},
		{
			name: "t5",
			args: args{
				target: []int{17, 5, 7, 1, 2, 19, 4, 20, 10, 18},
				arr:    []int{2, 10, 5, 9, 9, 17, 8, 1, 19, 1},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
