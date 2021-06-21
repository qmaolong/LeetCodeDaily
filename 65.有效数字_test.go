package main

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{
				s: "2",
			},
			want: true,
		},
		{
			name: "t2",
			args: args{
				s: "3e+7",
			},
			want: true,
		},
		{
			name: "t3",
			args: args{
				s: "0089",
			},
			want: true,
		},
		{
			name: "t4",
			args: args{
				s: "-0.1",
			},
			want: true,
		},
		{
			name: "t5",
			args: args{
				s: "+3.14",
			},
			want: true,
		},
		{
			name: "t6",
			args: args{
				s: "4.",
			},
			want: true,
		},
		{
			name: "t7",
			args: args{
				s: "-.9",
			},
			want: true,
		},
		{
			name: "t8",
			args: args{
				s: "2e10",
			},
			want: true,
		},
		{
			name: "t9",
			args: args{
				s: "-90E3",
			},
			want: true,
		},
		{
			name: "t10",
			args: args{
				s: "+6e-1",
			},
			want: true,
		},
		{
			name: "t11",
			args: args{
				s: "53.5e93",
			},
			want: true,
		},
		{
			name: "t12",
			args: args{
				s: "-123.456e789",
			},
			want: true,
		},
		{
			name: "t13",
			args: args{
				s: "abc",
			},
			want: false,
		},
		{
			name: "t14",
			args: args{
				s: "1a",
			},
			want: false,
		},
		{
			name: "t15",
			args: args{
				s: "1e",
			},
			want: false,
		},
		{
			name: "t16",
			args: args{
				s: "e3",
			},
			want: false,
		},
		{
			name: "t17",
			args: args{
				s: "99e2.5",
			},
			want: false,
		},
		{
			name: "t18",
			args: args{
				s: "--6",
			},
			want: false,
		},
		{
			name: "t19",
			args: args{
				s: "-+3",
			},
			want: false,
		},
		{
			name: "t20",
			args: args{
				s: "95a54e53",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
