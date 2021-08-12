package main

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				s: "bbbab",
			},
			want: 4,
		},
		{
			name: "t2",
			args: args{
				s: "cbbd",
			},
			want: 2,
		},
		{
			name: "t3",
			args: args{
				s: "a",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
